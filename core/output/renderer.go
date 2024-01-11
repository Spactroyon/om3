package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/andreazorzetto/yh/highlight"
	"github.com/fatih/color"
	tabwriter "github.com/juju/ansiterm"
	"github.com/opensvc/om3/util/render"
	"github.com/opensvc/om3/util/render/palette"
	"github.com/opensvc/om3/util/unstructured"
	"k8s.io/client-go/util/jsonpath"
	"sigs.k8s.io/yaml"
)

type (
	// RenderFunc is the protype of human format renderer functions.
	RenderFunc func() string

	// Renderer hosts the renderer options and data, and exposes the rendering
	// method.
	Renderer struct {
		DefaultOutput string
		Output        string
		Color         string
		Data          any
		Items         any
		HumanRenderer RenderFunc
		Colorize      *palette.ColorPaletteFunc
		Stream        bool
	}

	renderer interface {
		Render() string
	}
)

var (
	indent              = "    "
	regexpJSONKey       = regexp.MustCompile(`(".+":)`)
	regexpJSONReference = regexp.MustCompile(`({[\w.#-_:]+})`)
	regexpJSONScope     = regexp.MustCompile(`(@.+)(":)`)
	regexpJSONErrors    = regexp.MustCompile(`(")(down|stdby down|err|error)(")`)
	regexpJSONOptimal   = regexp.MustCompile(`(")(up|stdby up|ok)(")`)
	regexpJSONWarning   = regexp.MustCompile(`(")(warn)(")`)
	regexpJSONSecondary = regexp.MustCompile(`(")(n/a)(")`)
)

// Sprint returns the string representation of the data in one of the
// supported format (json, flat, human, ...).
//
// The human format needs a RenderFunc to be passed.
func (t Renderer) Sprint() (string, error) {
	var (
		options, format string
	)
	if t.DefaultOutput != "" {
		if t.Output == "auto" {
			t.Output = t.DefaultOutput
		}
		if strings.HasPrefix(t.Output, "+") {
			t.Output = t.DefaultOutput + "," + t.Output[1:]
		}
	}
	if i := strings.Index(t.Output, "="); i > 0 {
		options = t.Output[i+1:]
		format = t.Output[:i]
	} else {
		format = t.Output
	}
	formatID := toID[format]

	render.SetColor(t.Color)
	if t.Colorize == nil {
		t.Colorize = palette.DefaultFuncPalette()
	}
	switch data := t.Data.(type) {
	case []string:
		if data == nil {
			// JSON Marshal renders "null" for unallocated empty slices
			t.Data = make([]string, 0)
		}
	}
	switch formatID {
	case Flat:
		b, err := json.Marshal(t.Data)
		if err != nil {
			panic(err)
		}
		if color.NoColor {
			return SprintFlat(b), nil
		} else {
			return SprintFlatColor(b, t.Colorize), nil
		}
	case JSON:
		b, err := json.MarshalIndent(t.Data, "", indent)
		if err != nil {
			return "", err
		}
		s := string(b) + "\n"
		s = regexpJSONKey.ReplaceAllString(s, t.Colorize.Primary("$1"))
		s = regexpJSONReference.ReplaceAllString(s, t.Colorize.Optimal("$1"))
		s = regexpJSONScope.ReplaceAllString(s, t.Colorize.Error("$1")+"$2")
		s = regexpJSONErrors.ReplaceAllString(s, "$1"+t.Colorize.Error("$2")+"$3")
		s = regexpJSONOptimal.ReplaceAllString(s, "$1"+t.Colorize.Optimal("$2")+"$3")
		s = regexpJSONWarning.ReplaceAllString(s, "$1"+t.Colorize.Warning("$2")+"$3")
		s = regexpJSONSecondary.ReplaceAllString(s, "$1"+t.Colorize.Secondary("$2")+"$3")
		return s, nil
	case JSONLine:
		b, err := json.Marshal(t.Data)
		if err != nil {
			return "", err
		}
		return string(b) + "\n", nil
	case YAML:
		var sep string
		if t.Stream {
			sep = "---\n"
		}
		b, err := yaml.Marshal(t.Data)
		if err != nil {
			return "", err
		}
		if color.NoColor {
			return string(b) + sep, nil
		} else {
			buf := bytes.NewBuffer(b)
			s, err := highlight.Highlight(buf)
			if err != nil {
				return "", err
			}
			return s + sep, nil
		}
	case Tab:
		s, err := t.renderTab(options)
		if err != nil {
			return "", err
		}
		return s, nil
	default:
		if t.HumanRenderer != nil {
			return t.HumanRenderer(), nil
		}
		if r, ok := t.Data.(renderer); ok {
			return r.Render(), nil
		}
		b, err := json.MarshalIndent(t.Data, "", indent)
		if err != nil {
			return "", err
		}
		return string(b) + "\n", nil
	}
}

var jsonRegexp = regexp.MustCompile(`^\{\.?([^{}]+)\}$|^\.?([^{}]+)$`)

// RelaxedJSONPathExpression attempts to be flexible with JSONPath expressions, it accepts:
//   - metadata.name (no leading '.' or curly braces '{...}'
//   - {metadata.name} (no leading '.')
//   - .metadata.name (no curly braces '{...}')
//   - {.metadata.name} (complete expression)
//
// And transforms them all into a valid jsonpath expression:
//
//	{.metadata.name}
func RelaxedJSONPathExpression(pathExpression string) (string, error) {
	if len(pathExpression) == 0 {
		return pathExpression, nil
	}
	submatches := jsonRegexp.FindStringSubmatch(pathExpression)
	if submatches == nil {
		return "", fmt.Errorf("unexpected path string, expected a 'name1.name2' or '.name1.name2' or '{name1.name2}' or '{.name1.name2}'")
	}
	if len(submatches) != 3 {
		return "", fmt.Errorf("unexpected submatch list: %v", submatches)
	}
	var fieldSpec string
	if len(submatches[1]) != 0 {
		fieldSpec = submatches[1]
	} else {
		fieldSpec = submatches[2]
	}
	return fmt.Sprintf("{.%s}", fieldSpec), nil
}

func (t Renderer) renderTab(options string) (string, error) {
	var (
		hasHeader bool
		builder   strings.Builder
	)
	w := tabwriter.NewTabWriter(&builder, 1, 1, 1, ' ', 0)
	jsonPaths := make([]*jsonpath.JSONPath, 0)
	headers := make([]string, 0)
	for _, option := range strings.Split(options, ",") {
		l := strings.Split(option, ":")
		var header, jp string
		switch len(l) {
		case 2:
			header = l[0]
			jp = l[1]
		case 1:
			jp = option
		default:
			continue
		}
		if rjp, err := RelaxedJSONPathExpression(jp); err != nil {
			return "", err
		} else {
			jp = rjp
		}
		jsonPath := jsonpath.New(option)
		if err := jsonPath.Parse(jp); err != nil {
			return "", err
		}
		headers = append(headers, header+"\t")
		jsonPaths = append(jsonPaths, jsonPath)
		if header != "" {
			hasHeader = true
		}
	}
	if hasHeader {
		fmt.Fprintf(w, strings.Join(headers, "")+"\n")
	}
	var data any
	if t.Items != nil {
		data = t.Items
	} else {
		data = t.Data
	}
	unstructuredData, err := unstructured.NewListWithData(data)
	if err != nil {
		return "", err
	}
	for _, line := range unstructuredData {
		for _, jsonPath := range jsonPaths {
			values, err := jsonPath.FindResults(line)
			if err != nil {
				fmt.Fprintf(w, "<%s>\t", err)
				continue
			}
			valueStrings := []string{}
			if len(values) == 0 || len(values[0]) == 0 {
				fmt.Fprintf(w, "<none>\t")
				continue
			}
			for arrIx := range values {
				for valIx := range values[arrIx] {
					valueStrings = append(valueStrings, fmt.Sprintf("%v", values[arrIx][valIx].Interface()))
				}
			}
			value := strings.Join(valueStrings, ",")
			fmt.Fprintf(w, value+"\t")
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	return builder.String(), nil
}

// Print prints the representation of the data in one of the
// supported format (json, flat, human, ...).
//
// The human format needs a RenderFunc to be passed.
func (t Renderer) Print() {
	if s, err := t.Sprint(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Print(s)
	}
}

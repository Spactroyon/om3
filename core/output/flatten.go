package output

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Flatten accepts a nested struct and returns a flat struct with key like a.'b/c'.d[0].e
func Flatten(inputJSON map[string]interface{}) map[string]interface{} {
	var lkey = ""
	var flattened = make(map[string]interface{})
	flatten(inputJSON, lkey, &flattened)
	return flattened
}

// SprintFlat accepts a JSON formated byte array and returns the sorted
// "key = val" buffer
func SprintFlat(b []byte) string {
	var s string
	var data map[string]interface{}
	json.Unmarshal(b, &data)
	flattened := Flatten(data)
	keys := make([]string, 0)
	for key := range flattened {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		s += fmt.Sprintln(k, "=", flattened[k])
	}
	return s
}

// PrintFlat accepts a JSON formated byte array and prints to stdout the sorted
// "key = val"
func PrintFlat(b []byte) {
	var data map[string]interface{}
	json.Unmarshal(b, &data)
	flattened := Flatten(data)
	keys := make([]string, 0)
	for key := range flattened {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, "=", flattened[k])
	}
}

func flatten(value interface{}, lkey string, flattened *map[string]interface{}) {
	v := reflect.ValueOf(value)
	if value == nil {
		return
	}
	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < len(value.([]interface{})); i++ {
			k := fmt.Sprintf("%s[%d]", lkey, i)
			flatten(value.([]interface{})[i], k, flattened)
		}
	case reflect.Map:
		for rkey, rval := range value.(map[string]interface{}) {
			if strings.ContainsAny(rkey, ".#$/") {
				rkey = fmt.Sprintf("'%s'", rkey)
			}
			k := fmt.Sprintf("%s.%s", lkey, rkey)
			flatten(rval, k, flattened)
		}
	default:
		b, _ := json.Marshal(value)
		(*flattened)[lkey] = string(b)
	}
}

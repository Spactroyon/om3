package commands

import (
	"context"
	"fmt"

	"github.com/opensvc/om3/core/client"
	"github.com/opensvc/om3/core/output"
	"github.com/opensvc/om3/core/rawconfig"
	"github.com/opensvc/om3/daemon/api"
)

type (
	CmdObjectInstanceLs struct {
		OptsGlobal
		NodeSelector string
	}
)

func (t *CmdObjectInstanceLs) Run(selector, kind string) error {
	mergedSelector := mergeSelector(selector, t.ObjectSelector, kind, "")

	c, err := client.New(client.WithURL(t.Server))
	if err != nil {
		return err
	}
	params := api.GetInstancesParams{Path: &mergedSelector}
	if t.NodeSelector != "" {
		params.Node = &t.NodeSelector
	}
	resp, err := c.GetInstancesWithResponse(context.Background(), &params)
	if err != nil {
		return fmt.Errorf("api: %w", err)
	}
	var pb *api.Problem
	switch resp.StatusCode() {
	case 200:
		output.Renderer{
			DefaultOutput: "tab=OBJECT:meta.object,NODE:meta.node,AVAIL:data.status.avail",
			Output:        t.Output,
			Color:         t.Color,
			Data:          resp.JSON200,
			Items:         resp.JSON200.Items,
			Colorize:      rawconfig.Colorize,
		}.Print()
		return nil
	case 400:
		pb = resp.JSON400
	case 401:
		pb = resp.JSON401
	case 403:
		pb = resp.JSON403
	case 500:
		pb = resp.JSON500
	}
	return fmt.Errorf("%s", pb)
}

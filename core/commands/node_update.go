package commands

import (
	"context"
	"fmt"

	"github.com/opensvc/om3/core/actioncontext"
	"github.com/opensvc/om3/core/client"
	"github.com/opensvc/om3/core/clientcontext"
	"github.com/opensvc/om3/core/keyop"
	"github.com/opensvc/om3/core/nodeselector"
	"github.com/opensvc/om3/core/object"
	"github.com/opensvc/om3/daemon/api"
	"github.com/opensvc/om3/util/key"
)

type (
	CmdNodeUpdate struct {
		OptsGlobal
		OptsLock
		Delete []string
		Set    []string
		Unset  []string
	}
)

func (t *CmdNodeUpdate) Run() error {
	if t.Local {
		return t.doLocal()
	}
	if t.NodeSelector != "" {
		return t.doRemote()
	}
	if !clientcontext.IsSet() {
		return t.doLocal()
	}
	return fmt.Errorf("--node must be specified")
}

func (t *CmdNodeUpdate) doRemote() error {
	c, err := client.New()
	if err != nil {
		return err
	}
	params := api.PostNodeConfigUpdateParams{}
	params.Set = &t.Set
	params.Unset = &t.Unset
	params.Delete = &t.Delete
	nodenames, err := nodeselector.Expand(t.NodeSelector)
	if err != nil {
		return err
	}
	for _, nodename := range nodenames {
		response, err := c.PostNodeConfigUpdateWithResponse(context.Background(), nodename, &params)
		if err != nil {
			return err
		}
		switch response.StatusCode() {
		case 200:
		case 400:
			return fmt.Errorf("%s: %s", nodename, *response.JSON400)
		case 401:
			return fmt.Errorf("%s: %s", nodename, *response.JSON401)
		case 403:
			return fmt.Errorf("%s: %s", nodename, *response.JSON403)
		case 500:
			return fmt.Errorf("%s: %s", nodename, *response.JSON500)
		default:
			return fmt.Errorf("%s: unexpected response: %s", nodename, response.Status())
		}
	}
	return nil
}

func (t *CmdNodeUpdate) doLocal() error {
	o, err := object.NewNode()
	if err != nil {
		return err
	}
	ctx := context.Background()
	ctx = actioncontext.WithLockDisabled(ctx, t.Disable)
	ctx = actioncontext.WithLockTimeout(ctx, t.Timeout)
	return o.Update(ctx, t.Delete, key.ParseStrings(t.Unset), keyop.ParseOps(t.Set))
}

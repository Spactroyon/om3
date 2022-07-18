package commands

import (
	"github.com/spf13/cobra"
	"opensvc.com/opensvc/core/flag"
	"opensvc.com/opensvc/core/nodeaction"
	"opensvc.com/opensvc/core/object"
	"opensvc.com/opensvc/util/xstrings"
)

type (
	// CmdNodeComplianceShowModuleset is the cobra flag set of the sysreport command.
	CmdNodeComplianceShowModuleset struct {
		OptsGlobal
		OptModuleset
	}
)

// Init configures a cobra command and adds it to the parent command.
func (t *CmdNodeComplianceShowModuleset) Init(parent *cobra.Command) {
	cmd := t.cmd()
	parent.AddCommand(cmd)
	flag.Install(cmd, t)
}

func (t *CmdNodeComplianceShowModuleset) cmd() *cobra.Command {
	return &cobra.Command{
		Use:     "moduleset",
		Short:   "Show compliance moduleset and modules attached to this node.",
		Aliases: []string{"modulese", "modules", "module", "modul", "modu", "mod", "mo"},
		Run: func(_ *cobra.Command, _ []string) {
			t.run()
		},
	}
}

func (t *CmdNodeComplianceShowModuleset) run() {
	nodeaction.New(
		nodeaction.WithLocal(t.Local),
		nodeaction.WithRemoteNodes(t.NodeSelector),
		nodeaction.WithFormat(t.Format),
		nodeaction.WithColor(t.Color),
		nodeaction.WithServer(t.Server),
		nodeaction.WithRemoteAction("compliance show moduleset"),
		nodeaction.WithRemoteOptions(map[string]interface{}{
			"format":    t.Format,
			"moduleset": t.Moduleset,
		}),
		nodeaction.WithLocalRun(func() (interface{}, error) {
			n, err := object.NewNode()
			if err != nil {
				return nil, err
			}
			comp, err := n.NewCompliance()
			if err != nil {
				return nil, err
			}
			modsets := xstrings.Split(t.Moduleset, ",")
			data, err := comp.GetData(modsets)
			if err != nil {
				return nil, err
			}
			tree := data.ModulesetsTree()
			return tree, nil
		}),
	).Do()
}

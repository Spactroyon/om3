package commands

import (
	"github.com/spf13/cobra"
	"opensvc.com/opensvc/core/object"
)

type (
	// CmdObjectStatus is the cobra flag set of the status command.
	CmdObjectStatus struct {
		flagSetGlobal
		flagSetObject
	}
)

// Init configures a cobra command and adds it to the parent command.
func (t *CmdObjectStatus) Init(kind string, parent *cobra.Command, selector *string) {
	cmd := t.cmd(kind, selector)
	parent.AddCommand(cmd)
	t.flagSetGlobal.init(cmd)
	t.flagSetObject.init(cmd)
}

func (t *CmdObjectStatus) cmd(kind string, selector *string) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Print selected service and instance status",
		Long: `Resources Flags:

(1) R   Running,           . Not Running
(2) M   Monitored,         . Not Monitored
(3) D   Disabled,          . Enabled
(4) O   Optional,          . Not Optional
(5) E   Encap,             . Not Encap
(6) P   Not Provisioned,   . Provisioned
(7) S   Standby,           . Not Standby
(8) <n> Remaining Restart, + if more than 10,   . No Restart

`,
		Run: func(cmd *cobra.Command, args []string) {
			t.run(selector, kind)
		},
	}
}

func (t *CmdObjectStatus) run(selector *string, kind string) {
	mergedSelector := mergeSelector(*selector, t.ObjectSelector, kind, "")
	selection := object.NewSelection(mergedSelector)
	selection.SetServer(t.Server)
	selection.SetLocal(true)
	selection.Do(object.Action{
		Method: "Status",
	})
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"opensvc.com/opensvc/core/entrypoints/monitor"
)

var (
	daemonStatusWatchFlag    bool
	daemonStatusSelectorFlag string
)

var daemonStatusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Print the cluster status",
	Long:    monitor.CmdLong,
	Aliases: []string{"statu"},
	Run:     daemonStatusCmdRun,
}

func init() {
	daemonCmd.AddCommand(daemonStatusCmd)
	daemonStatusCmd.Flags().BoolVarP(&daemonStatusWatchFlag, "watch", "w", false, "Watch the monitor changes")
	daemonStatusCmd.Flags().StringVarP(&daemonStatusSelectorFlag, "selector", "s", "**", "Select opensvc objects (ex: **/db*,*/svc/db*)")
}

func daemonStatusCmdRun(_ *cobra.Command, _ []string) {
	m := monitor.New()
	m.SetColor(colorFlag)
	m.SetFormat(formatFlag)

	cli, err := newClient()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}
	if daemonStatusWatchFlag {
		statusGetter := cli.NewGetDaemonStatus().SetSelector(daemonStatusSelectorFlag)
		eventsGetter := cli.NewGetEvents().SetSelector(daemonStatusSelectorFlag)
		_ = m.DoWatch(statusGetter, eventsGetter, os.Stdout)
	} else {
		getter := cli.NewGetDaemonStatus().SetSelector(daemonStatusSelectorFlag)
		m.Do(getter, os.Stdout)
	}
}

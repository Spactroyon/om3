package object

import (
	"time"
)

type (
	// OptsGlobal contains options accepted by all actions
	OptsGlobal struct {
		Color          string `flag:"color"`
		Format         string `flag:"format"`
		Server         string `flag:"server"`
		Local          bool   `flag:"local"`
		NodeSelector   string `flag:"node"`
		ObjectSelector string `flag:"object"`
		DryRun         bool   `flag:"dry-run"`
	}

	// OptsResourceSelector contains options needed to initialize a
	// resourceselector.Options struct
	OptsResourceSelector struct {
		RID    string `flag:"rid"`
		Subset string `flag:"subset"`
		Tag    string `flag:"tags"`
	}

	// OptsLock contains options accepted by all actions using an action lock
	OptsLock struct {
		Disable bool          `flag:"nolock"`
		Timeout time.Duration `flag:"waitlock"`
	}

	// OptsAsync contains options accepted by all actions having an orchestration
	OptsAsync struct {
		Watch bool          `flag:"watch"`
		Wait  bool          `flag:"wait"`
		Time  time.Duration `flag:"time"`
	}

	// OptDisableRollback contains the disable-rollback option
	OptDisableRollback struct {
		DisableRollback bool `flag:"disable-rollback"`
	}

	// OptForce contains the force option
	OptForce struct {
		Force bool `flag:"force"`
	}

	// OptAttach contains the force option
	OptAttach struct {
		Attach bool `flag:"attach"`
	}

	// OptModule contains the module option
	OptModule struct {
		Module string `flag:"module"`
	}

	// OptModuleset contains the moduleset option
	OptModuleset struct {
		Moduleset string `flag:"moduleset"`
	}

	// OptRuleset contains the ruleset option
	OptRuleset struct {
		Ruleset string `flag:"ruleset"`
	}

	// OptConfirm contains the confirm option
	OptConfirm struct {
		Confirm bool `flag:"confirm"`
	}

	// OptCron contains the cron option
	OptCron struct {
		Cron bool `flag:"cron"`
	}

	// OpTo sets a barrier when iterating over a resource lister
	OptTo struct {
		To     string `flag:"to"`
		UpTo   string `flag:"upto"`   // Deprecated
		DownTo string `flag:"downto"` // Deprecated
	}

	//
	// OptLeader is used by the provision and unprovision action to trigger
	// allocation of shared resources on the leader node only.
	//
	// This option is usually set by the daemon, who is responsible for the
	// leader detection.
	//
	OptLeader struct {
		Leader bool `flag:"leader"`
	}

	OptsCreate struct {
		OptsGlobal
		OptsAsync
		OptsLock
		OptsResourceSelector
		OptTo
		OptForce
		Template    string   `flag:"template"`
		Config      string   `flag:"config"`
		Keywords    []string `flag:"kwops"`
		Env         string   `flag:"env"`
		Interactive bool     `flag:"interactive"`
		Provision   bool     `flag:"provision"`
		Restore     bool     `flag:"restore"`
		Namespace   string   `flag:"createnamespace"`
	}
)

var (
	defaultOptsGlobal = OptsGlobal{
		Color:  "auto",
		Format: "auto",
	}
	defaultOptsLock = OptsLock{
		Timeout: 5 * time.Second,
	}
)

func (t OptsResourceSelector) ResourceSelectorRID() string {
	return t.RID
}
func (t OptsResourceSelector) ResourceSelectorTag() string {
	return t.Tag
}
func (t OptsResourceSelector) ResourceSelectorSubset() string {
	return t.Subset
}
func (t OptDisableRollback) IsRollbackDisabled() bool {
	return t.DisableRollback
}
func (t OptCron) IsCron() bool {
	return t.Cron
}
func (t OptConfirm) IsConfirm() bool {
	return t.Confirm
}
func (t OptForce) IsForce() bool {
	return t.Force
}
func (t OptTo) ToStr() string {
	return t.To
}
func (t OptLeader) IsLeader() bool {
	return t.Leader
}
func (t OptsGlobal) IsDryRun() bool {
	return t.DryRun
}
func (t OptsLock) IsLockDisabled() bool {
	return t.Disable
}
func (t OptsLock) LockTimeout() time.Duration {
	return t.Timeout
}

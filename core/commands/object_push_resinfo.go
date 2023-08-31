package commands

import (
	"context"

	"github.com/opensvc/om3/core/actioncontext"
	"github.com/opensvc/om3/core/object"
	"github.com/opensvc/om3/core/objectaction"
	"github.com/opensvc/om3/core/path"
)

type (
	CmdObjectPushResInfo struct {
		OptsGlobal
		OptsLock
		OptsResourceSelector
		OptTo
	}
)

func (t *CmdObjectPushResInfo) Run(selector, kind string) error {
	mergedSelector := mergeSelector(selector, t.ObjectSelector, kind, "")
	return objectaction.New(
		objectaction.WithObjectSelector(mergedSelector),
		objectaction.WithRID(t.RID),
		objectaction.WithTag(t.Tag),
		objectaction.WithSubset(t.Subset),
		objectaction.WithLocal(true),
		objectaction.WithFormat(t.Output),
		objectaction.WithColor(t.Color),
		objectaction.WithRemoteNodes(t.NodeSelector),
		objectaction.WithRemoteAction("push resinfo"),
		objectaction.WithLocalRun(func(ctx context.Context, p path.T) (any, error) {
			o, err := object.NewActor(p,
				object.WithConsoleLog(t.Log != ""),
				object.WithConsoleColor(t.Color != "no"),
			)

			if err != nil {
				return nil, err
			}
			ctx = actioncontext.WithLockDisabled(ctx, t.Disable)
			ctx = actioncontext.WithLockTimeout(ctx, t.Timeout)
			ctx = actioncontext.WithTo(ctx, t.To)
			return o.PushResInfo(ctx)
		}),
	).Do()
}

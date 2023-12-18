package daemonapi

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/opensvc/om3/core/keyop"
	"github.com/opensvc/om3/daemon/rbac"
)

func assertGrant(ctx echo.Context, grants ...rbac.Grant) (bool, error) {
	if !grantsFromContext(ctx).HasGrant(grants...) {
		return false, JSONForbiddenMissingGrant(ctx, grants...)
	}
	return true, nil
}

func assertRole(ctx echo.Context, roles ...rbac.Role) (bool, error) {
	if !grantsFromContext(ctx).HasRole(roles...) {
		return false, JSONForbiddenMissingRole(ctx, roles...)
	}
	return true, nil
}

func keyopStringRbac(op string) error {
	kop := keyop.Parse(op)
	if kop == nil {
		return fmt.Errorf("invalid op: %s", op)
	}
	return keyopRbac(*kop)
}

func keyopRbac(op keyop.T) error {
	if strings.HasSuffix(op.Key.Option, "_trigger") {
		return fmt.Errorf("%s", op.Key)
	}
	drvGroup := strings.Split(op.Key.Section, "#")[0]
	switch drvGroup {
	case "app", "task":
		switch op.Key.Option {
		case "script", "start", "stop", "check", "info":
			return fmt.Errorf("%s", op.Key)
		}
	}
	return nil
}

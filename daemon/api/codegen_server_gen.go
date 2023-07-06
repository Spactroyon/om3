// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /auth/token)
	PostAuthToken(ctx echo.Context, params PostAuthTokenParams) error

	// (GET /daemon/dns/dump)
	GetDaemonDNSDump(ctx echo.Context) error

	// (GET /daemon/events)
	GetDaemonEvents(ctx echo.Context, params GetDaemonEventsParams) error

	// (POST /daemon/join)
	PostDaemonJoin(ctx echo.Context, params PostDaemonJoinParams) error

	// (POST /daemon/leave)
	PostDaemonLeave(ctx echo.Context, params PostDaemonLeaveParams) error

	// (POST /daemon/logs/control)
	PostDaemonLogsControl(ctx echo.Context) error

	// (GET /daemon/running)
	GetDaemonRunning(ctx echo.Context) error

	// (GET /daemon/status)
	GetDaemonStatus(ctx echo.Context, params GetDaemonStatusParams) error

	// (POST /daemon/stop)
	PostDaemonStop(ctx echo.Context) error

	// (POST /daemon/sub/action)
	PostDaemonSubAction(ctx echo.Context) error

	// (POST /instance/status)
	PostInstanceStatus(ctx echo.Context) error

	// (GET /networks)
	GetNetworks(ctx echo.Context, params GetNetworksParams) error

	// (GET /node/backlogs)
	GetNodeBacklogs(ctx echo.Context, params GetNodeBacklogsParams) error

	// (POST /node/clear)
	PostNodeClear(ctx echo.Context) error

	// (GET /node/drbd/allocation)
	GetNodeDRBDAllocation(ctx echo.Context) error

	// (GET /node/drbd/config)
	GetNodeDRBDConfig(ctx echo.Context, params GetNodeDRBDConfigParams) error

	// (POST /node/drbd/config)
	PostNodeDRBDConfig(ctx echo.Context, params PostNodeDRBDConfigParams) error

	// (GET /node/logs)
	GetNodeLogs(ctx echo.Context, params GetNodeLogsParams) error

	// (POST /node/monitor)
	PostNodeMonitor(ctx echo.Context) error

	// (GET /nodes/info)
	GetNodesInfo(ctx echo.Context) error

	// (POST /object/abort)
	PostObjectAbort(ctx echo.Context) error

	// (GET /object/backlogs)
	GetObjectBacklogs(ctx echo.Context, params GetObjectBacklogsParams) error

	// (POST /object/clear)
	PostObjectClear(ctx echo.Context) error

	// (GET /object/config)
	GetObjectConfig(ctx echo.Context, params GetObjectConfigParams) error

	// (GET /object/file)
	GetObjectFile(ctx echo.Context, params GetObjectFileParams) error

	// (GET /object/logs)
	GetObjectLogs(ctx echo.Context, params GetObjectLogsParams) error

	// (POST /object/monitor)
	PostObjectMonitor(ctx echo.Context) error

	// (POST /object/progress)
	PostObjectProgress(ctx echo.Context) error

	// (GET /object/selector)
	GetObjectSelector(ctx echo.Context, params GetObjectSelectorParams) error

	// (POST /object/switchTo)
	PostObjectSwitchTo(ctx echo.Context) error

	// (GET /pools)
	GetPools(ctx echo.Context, params GetPoolsParams) error

	// (GET /public/openapi)
	GetSwagger(ctx echo.Context) error

	// (GET /relay/message)
	GetRelayMessage(ctx echo.Context, params GetRelayMessageParams) error

	// (POST /relay/message)
	PostRelayMessage(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostAuthToken converts echo context to params.
func (w *ServerInterfaceWrapper) PostAuthToken(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostAuthTokenParams
	// ------------- Optional query parameter "role" -------------

	err = runtime.BindQueryParameter("form", true, false, "role", ctx.QueryParams(), &params.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter role: %s", err))
	}

	// ------------- Optional query parameter "duration" -------------

	err = runtime.BindQueryParameter("form", true, false, "duration", ctx.QueryParams(), &params.Duration)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter duration: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostAuthToken(ctx, params)
	return err
}

// GetDaemonDNSDump converts echo context to params.
func (w *ServerInterfaceWrapper) GetDaemonDNSDump(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDaemonDNSDump(ctx)
	return err
}

// GetDaemonEvents converts echo context to params.
func (w *ServerInterfaceWrapper) GetDaemonEvents(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDaemonEventsParams
	// ------------- Optional query parameter "duration" -------------

	err = runtime.BindQueryParameter("form", true, false, "duration", ctx.QueryParams(), &params.Duration)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter duration: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDaemonEvents(ctx, params)
	return err
}

// PostDaemonJoin converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonJoin(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostDaemonJoinParams
	// ------------- Required query parameter "node" -------------

	err = runtime.BindQueryParameter("form", true, true, "node", ctx.QueryParams(), &params.Node)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter node: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonJoin(ctx, params)
	return err
}

// PostDaemonLeave converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonLeave(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostDaemonLeaveParams
	// ------------- Required query parameter "node" -------------

	err = runtime.BindQueryParameter("form", true, true, "node", ctx.QueryParams(), &params.Node)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter node: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonLeave(ctx, params)
	return err
}

// PostDaemonLogsControl converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonLogsControl(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonLogsControl(ctx)
	return err
}

// GetDaemonRunning converts echo context to params.
func (w *ServerInterfaceWrapper) GetDaemonRunning(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDaemonRunning(ctx)
	return err
}

// GetDaemonStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetDaemonStatus(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDaemonStatusParams
	// ------------- Optional query parameter "namespace" -------------

	err = runtime.BindQueryParameter("form", true, false, "namespace", ctx.QueryParams(), &params.Namespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter namespace: %s", err))
	}

	// ------------- Optional query parameter "relatives" -------------

	err = runtime.BindQueryParameter("form", true, false, "relatives", ctx.QueryParams(), &params.Relatives)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter relatives: %s", err))
	}

	// ------------- Optional query parameter "selector" -------------

	err = runtime.BindQueryParameter("form", true, false, "selector", ctx.QueryParams(), &params.Selector)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter selector: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDaemonStatus(ctx, params)
	return err
}

// PostDaemonStop converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonStop(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonStop(ctx)
	return err
}

// PostDaemonSubAction converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonSubAction(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonSubAction(ctx)
	return err
}

// PostInstanceStatus converts echo context to params.
func (w *ServerInterfaceWrapper) PostInstanceStatus(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostInstanceStatus(ctx)
	return err
}

// GetNetworks converts echo context to params.
func (w *ServerInterfaceWrapper) GetNetworks(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNetworksParams
	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNetworks(ctx, params)
	return err
}

// GetNodeBacklogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodeBacklogs(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNodeBacklogsParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Required query parameter "paths" -------------

	err = runtime.BindQueryParameter("form", true, true, "paths", ctx.QueryParams(), &params.Paths)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter paths: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodeBacklogs(ctx, params)
	return err
}

// PostNodeClear converts echo context to params.
func (w *ServerInterfaceWrapper) PostNodeClear(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostNodeClear(ctx)
	return err
}

// GetNodeDRBDAllocation converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodeDRBDAllocation(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodeDRBDAllocation(ctx)
	return err
}

// GetNodeDRBDConfig converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodeDRBDConfig(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNodeDRBDConfigParams
	// ------------- Required query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, true, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodeDRBDConfig(ctx, params)
	return err
}

// PostNodeDRBDConfig converts echo context to params.
func (w *ServerInterfaceWrapper) PostNodeDRBDConfig(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostNodeDRBDConfigParams
	// ------------- Required query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, true, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostNodeDRBDConfig(ctx, params)
	return err
}

// GetNodeLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodeLogs(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNodeLogsParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Required query parameter "paths" -------------

	err = runtime.BindQueryParameter("form", true, true, "paths", ctx.QueryParams(), &params.Paths)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter paths: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodeLogs(ctx, params)
	return err
}

// PostNodeMonitor converts echo context to params.
func (w *ServerInterfaceWrapper) PostNodeMonitor(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostNodeMonitor(ctx)
	return err
}

// GetNodesInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodesInfo(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodesInfo(ctx)
	return err
}

// PostObjectAbort converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectAbort(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectAbort(ctx)
	return err
}

// GetObjectBacklogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectBacklogs(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectBacklogsParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Required query parameter "paths" -------------

	err = runtime.BindQueryParameter("form", true, true, "paths", ctx.QueryParams(), &params.Paths)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter paths: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectBacklogs(ctx, params)
	return err
}

// PostObjectClear converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectClear(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectClear(ctx)
	return err
}

// GetObjectConfig converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectConfig(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectConfigParams
	// ------------- Required query parameter "path" -------------

	err = runtime.BindQueryParameter("form", true, true, "path", ctx.QueryParams(), &params.Path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter path: %s", err))
	}

	// ------------- Optional query parameter "evaluate" -------------

	err = runtime.BindQueryParameter("form", true, false, "evaluate", ctx.QueryParams(), &params.Evaluate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter evaluate: %s", err))
	}

	// ------------- Optional query parameter "impersonate" -------------

	err = runtime.BindQueryParameter("form", true, false, "impersonate", ctx.QueryParams(), &params.Impersonate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter impersonate: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectConfig(ctx, params)
	return err
}

// GetObjectFile converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectFile(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectFileParams
	// ------------- Required query parameter "path" -------------

	err = runtime.BindQueryParameter("form", true, true, "path", ctx.QueryParams(), &params.Path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter path: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectFile(ctx, params)
	return err
}

// GetObjectLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectLogs(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectLogsParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Required query parameter "paths" -------------

	err = runtime.BindQueryParameter("form", true, true, "paths", ctx.QueryParams(), &params.Paths)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter paths: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectLogs(ctx, params)
	return err
}

// PostObjectMonitor converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectMonitor(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectMonitor(ctx)
	return err
}

// PostObjectProgress converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectProgress(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectProgress(ctx)
	return err
}

// GetObjectSelector converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectSelector(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectSelectorParams
	// ------------- Required query parameter "selector" -------------

	err = runtime.BindQueryParameter("form", true, true, "selector", ctx.QueryParams(), &params.Selector)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter selector: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectSelector(ctx, params)
	return err
}

// PostObjectSwitchTo converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectSwitchTo(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectSwitchTo(ctx)
	return err
}

// GetPools converts echo context to params.
func (w *ServerInterfaceWrapper) GetPools(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPoolsParams
	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPools(ctx, params)
	return err
}

// GetSwagger converts echo context to params.
func (w *ServerInterfaceWrapper) GetSwagger(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSwagger(ctx)
	return err
}

// GetRelayMessage converts echo context to params.
func (w *ServerInterfaceWrapper) GetRelayMessage(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRelayMessageParams
	// ------------- Optional query parameter "nodename" -------------

	err = runtime.BindQueryParameter("form", true, false, "nodename", ctx.QueryParams(), &params.Nodename)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter nodename: %s", err))
	}

	// ------------- Optional query parameter "cluster_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "cluster_id", ctx.QueryParams(), &params.ClusterId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter cluster_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRelayMessage(ctx, params)
	return err
}

// PostRelayMessage converts echo context to params.
func (w *ServerInterfaceWrapper) PostRelayMessage(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRelayMessage(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/auth/token", wrapper.PostAuthToken)
	router.GET(baseURL+"/daemon/dns/dump", wrapper.GetDaemonDNSDump)
	router.GET(baseURL+"/daemon/events", wrapper.GetDaemonEvents)
	router.POST(baseURL+"/daemon/join", wrapper.PostDaemonJoin)
	router.POST(baseURL+"/daemon/leave", wrapper.PostDaemonLeave)
	router.POST(baseURL+"/daemon/logs/control", wrapper.PostDaemonLogsControl)
	router.GET(baseURL+"/daemon/running", wrapper.GetDaemonRunning)
	router.GET(baseURL+"/daemon/status", wrapper.GetDaemonStatus)
	router.POST(baseURL+"/daemon/stop", wrapper.PostDaemonStop)
	router.POST(baseURL+"/daemon/sub/action", wrapper.PostDaemonSubAction)
	router.POST(baseURL+"/instance/status", wrapper.PostInstanceStatus)
	router.GET(baseURL+"/networks", wrapper.GetNetworks)
	router.GET(baseURL+"/node/backlogs", wrapper.GetNodeBacklogs)
	router.POST(baseURL+"/node/clear", wrapper.PostNodeClear)
	router.GET(baseURL+"/node/drbd/allocation", wrapper.GetNodeDRBDAllocation)
	router.GET(baseURL+"/node/drbd/config", wrapper.GetNodeDRBDConfig)
	router.POST(baseURL+"/node/drbd/config", wrapper.PostNodeDRBDConfig)
	router.GET(baseURL+"/node/logs", wrapper.GetNodeLogs)
	router.POST(baseURL+"/node/monitor", wrapper.PostNodeMonitor)
	router.GET(baseURL+"/nodes/info", wrapper.GetNodesInfo)
	router.POST(baseURL+"/object/abort", wrapper.PostObjectAbort)
	router.GET(baseURL+"/object/backlogs", wrapper.GetObjectBacklogs)
	router.POST(baseURL+"/object/clear", wrapper.PostObjectClear)
	router.GET(baseURL+"/object/config", wrapper.GetObjectConfig)
	router.GET(baseURL+"/object/file", wrapper.GetObjectFile)
	router.GET(baseURL+"/object/logs", wrapper.GetObjectLogs)
	router.POST(baseURL+"/object/monitor", wrapper.PostObjectMonitor)
	router.POST(baseURL+"/object/progress", wrapper.PostObjectProgress)
	router.GET(baseURL+"/object/selector", wrapper.GetObjectSelector)
	router.POST(baseURL+"/object/switchTo", wrapper.PostObjectSwitchTo)
	router.GET(baseURL+"/pools", wrapper.GetPools)
	router.GET(baseURL+"/public/openapi", wrapper.GetSwagger)
	router.GET(baseURL+"/relay/message", wrapper.GetRelayMessage)
	router.POST(baseURL+"/relay/message", wrapper.PostRelayMessage)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

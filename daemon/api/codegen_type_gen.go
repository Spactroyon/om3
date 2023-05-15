// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	BasicAuthScopes  = "basicAuth.Scopes"
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for PostDaemonLogsControlLevel.
const (
	PostDaemonLogsControlLevelDebug PostDaemonLogsControlLevel = "debug"
	PostDaemonLogsControlLevelError PostDaemonLogsControlLevel = "error"
	PostDaemonLogsControlLevelFatal PostDaemonLogsControlLevel = "fatal"
	PostDaemonLogsControlLevelInfo  PostDaemonLogsControlLevel = "info"
	PostDaemonLogsControlLevelNone  PostDaemonLogsControlLevel = "none"
	PostDaemonLogsControlLevelPanic PostDaemonLogsControlLevel = "panic"
	PostDaemonLogsControlLevelWarn  PostDaemonLogsControlLevel = "warn"
)

// Defines values for Orchestrate.
const (
	OrchestrateHa    Orchestrate = "ha"
	OrchestrateNo    Orchestrate = "no"
	OrchestrateStart Orchestrate = "start"
)

// Defines values for Placement.
const (
	PlacementLoadAvg    Placement = "load avg"
	PlacementNodesOrder Placement = "nodes order"
	PlacementNone       Placement = "none"
	PlacementScore      Placement = "score"
	PlacementShift      Placement = "shift"
	PlacementSpread     Placement = "spread"
)

// Defines values for PostDaemonSubActionAction.
const (
	PostDaemonSubActionActionStart PostDaemonSubActionAction = "start"
	PostDaemonSubActionActionStop  PostDaemonSubActionAction = "stop"
)

// Defines values for Provisioned.
const (
	False Provisioned = "false"
	Mixed Provisioned = "mixed"
	Na    Provisioned = "n/a"
	True  Provisioned = "true"
)

// Defines values for Role.
const (
	Admin          Role = "admin"
	Blacklistadmin Role = "blacklistadmin"
	Guest          Role = "guest"
	Heartbeat      Role = "heartbeat"
	Root           Role = "root"
	Squatter       Role = "squatter"
)

// Defines values for Topology.
const (
	Failover Topology = "failover"
	Flex     Topology = "flex"
)

// AuthToken defines model for AuthToken.
type AuthToken struct {
	Token         string    `json:"token"`
	TokenExpireAt time.Time `json:"token_expire_at"`
}

// DNSZone defines model for DNSZone.
type DNSZone = []DnsRecord

// DRBDAllocation defines model for DRBDAllocation.
type DRBDAllocation struct {
	ExpireAt time.Time          `json:"expire_at"`
	Id       openapi_types.UUID `json:"id"`
	Minor    int                `json:"minor"`
	Port     int                `json:"port"`
}

// DRBDConfig defines model for DRBDConfig.
type DRBDConfig struct {
	Data []byte `json:"data"`
}

// DaemonRunning defines model for DaemonRunning.
type DaemonRunning struct {
	Data []struct {
		Data     bool   `json:"data"`
		Endpoint string `json:"endpoint"`
	} `json:"data"`
	Entrypoint string `json:"entrypoint"`
	Status     int    `json:"status"`
}

// DaemonStatus defines model for DaemonStatus.
type DaemonStatus struct {
	Cluster Cluster `json:"cluster"`
	Daemon  Daemon  `json:"daemon"`
}

// EventList responseEventList is a list of sse
type EventList = openapi_types.File

// NetworkStatus defines model for NetworkStatus.
type NetworkStatus struct {
	Errors  *[]string           `json:"errors,omitempty"`
	Ips     *[]NetworkStatusIp  `json:"ips,omitempty"`
	Name    *string             `json:"name,omitempty"`
	Network *string             `json:"network,omitempty"`
	Type    *string             `json:"type,omitempty"`
	Usage   *NetworkStatusUsage `json:"usage,omitempty"`
}

// NetworkStatusIp defines model for NetworkStatusIp.
type NetworkStatusIp struct {
	Ip   string `json:"ip"`
	Node string `json:"node"`
	Path string `json:"path"`
	Rid  string `json:"rid"`
}

// NetworkStatusList defines model for NetworkStatusList.
type NetworkStatusList = []NetworkStatus

// NetworkStatusUsage defines model for NetworkStatusUsage.
type NetworkStatusUsage struct {
	Free int     `json:"free"`
	Pct  float32 `json:"pct"`
	Size int     `json:"size"`
	Used int     `json:"used"`
}

// PostDaemonLogsControl defines model for PostDaemonLogsControl.
type PostDaemonLogsControl struct {
	Level PostDaemonLogsControlLevel `json:"level"`
}

// PostDaemonLogsControlLevel defines model for PostDaemonLogsControl.Level.
type PostDaemonLogsControlLevel string

// App defines model for app.
type App = string

// Cluster defines model for cluster.
type Cluster struct {
	Config ClusterConfig `json:"config"`
	Node   ClusterNode   `json:"node"`
	Object ClusterObject `json:"object"`
	Status ClusterStatus `json:"status"`
}

// ClusterConfig defines model for clusterConfig.
type ClusterConfig = map[string]interface{}

// ClusterNode defines model for clusterNode.
type ClusterNode = map[string]interface{}

// ClusterObject defines model for clusterObject.
type ClusterObject = map[string]interface{}

// ClusterStatus defines model for clusterStatus.
type ClusterStatus = map[string]interface{}

// Daemon defines model for daemon.
type Daemon struct {
	Collector DaemonCollector `json:"collector"`
	Dns       DaemonDNS       `json:"dns"`
	Hb        DaemonHb        `json:"hb"`
	Listener  DaemonListener  `json:"listener"`
	Monitor   DaemonMonitor   `json:"monitor"`
	Routines  int             `json:"routines"`
	Scheduler DaemonScheduler `json:"scheduler"`
}

// DaemonCollector defines model for daemonCollector.
type DaemonCollector = DaemonSubsystemStatus

// DaemonDNS defines model for daemonDNS.
type DaemonDNS = DaemonSubsystemStatus

// DaemonHb defines model for daemonHb.
type DaemonHb struct {
	Modes   []DaemonHbMode   `json:"modes"`
	Streams []DaemonHbStream `json:"streams"`
}

// DaemonHbMode defines model for daemonHbMode.
type DaemonHbMode struct {
	// Mode the type of hb message used by node except when Type is patch where it is the patch queue length
	Mode string `json:"mode"`

	// Node a cluster node
	Node string `json:"node"`

	// Type the heartbeat message type used by node
	Type string `json:"type"`
}

// DaemonHbStream defines model for daemonHbStream.
type DaemonHbStream struct {
	Alerts     []DaemonSubsystemAlert `json:"alerts"`
	Beating    bool                   `json:"beating"`
	Configured time.Time              `json:"configured"`
	Created    time.Time              `json:"created"`
	Id         string                 `json:"id"`
	Last       time.Time              `json:"last"`
	State      string                 `json:"state"`

	// Type hb stream type
	Type string `json:"type"`
}

// DaemonHbStreamPeer defines model for daemonHbStreamPeer.
type DaemonHbStreamPeer struct {
	Beating bool      `json:"beating"`
	Last    time.Time `json:"last"`
}

// DaemonHbStreamType defines model for daemonHbStreamType.
type DaemonHbStreamType struct {
	// Type hb stream type
	Type string `json:"type"`
}

// DaemonListener defines model for daemonListener.
type DaemonListener = DaemonSubsystemStatus

// DaemonMonitor defines model for daemonMonitor.
type DaemonMonitor = DaemonSubsystemStatus

// DaemonScheduler defines model for daemonScheduler.
type DaemonScheduler = DaemonSubsystemStatus

// DaemonSubsystemAlert defines model for daemonSubsystemAlert.
type DaemonSubsystemAlert struct {
	Message  string   `json:"message"`
	Severity Severity `json:"severity"`
}

// DaemonSubsystemStatus defines model for daemonSubsystemStatus.
type DaemonSubsystemStatus struct {
	Alerts     []DaemonSubsystemAlert `json:"alerts"`
	Configured time.Time              `json:"configured"`
	Created    time.Time              `json:"created"`
	Id         string                 `json:"id"`
	State      string                 `json:"state"`
}

// DnsRecord defines model for dnsRecord.
type DnsRecord struct {
	Class string `json:"class"`
	Data  string `json:"data"`
	Name  string `json:"name"`
	Ttl   int    `json:"ttl"`
	Type  string `json:"type"`
}

// InstanceStatus defines model for instanceStatus.
type InstanceStatus struct {
	App         *App          `json:"app,omitempty"`
	Avail       Status        `json:"avail"`
	Children    *PathRelation `json:"children,omitempty"`
	Constraints *bool         `json:"constraints,omitempty"`
	Csum        *string       `json:"csum,omitempty"`
	Drp         *bool         `json:"drp,omitempty"`
	Env         *string       `json:"env,omitempty"`
	FlexMax     *int          `json:"flex_max,omitempty"`
	FlexMin     *int          `json:"flex_min,omitempty"`
	FlexTarget  *int          `json:"flex_target,omitempty"`
	Frozen      time.Time     `json:"frozen"`
	Kind        Kind          `json:"kind"`
	Optional    *Status       `json:"optional,omitempty"`
	Orchestrate *Orchestrate  `json:"orchestrate,omitempty"`
	Overall     Status        `json:"overall"`
	Parents     *PathRelation `json:"parents,omitempty"`

	// Placement object placement policy
	Placement *Placement `json:"placement,omitempty"`

	// Preserved preserve is true if this status has not been updated due to a
	// heartbeat downtime covered by a maintenance period.
	// when the maintenance period ends, the status should be unchanged,
	// and preserve will be set to false.
	Preserved *bool `json:"preserved,omitempty"`

	// Priority scheduling priority of an object instance on a its node
	Priority *Priority `json:"priority,omitempty"`

	// Provisioned service, instance or resource provisioned state
	Provisioned Provisioned              `json:"provisioned"`
	Resources   *[]ResourceExposedStatus `json:"resources,omitempty"`
	Running     *[]string                `json:"running,omitempty"`
	Scale       *int                     `json:"scale,omitempty"`
	Slaves      *PathRelation            `json:"slaves,omitempty"`
	StatusGroup *string                  `json:"status_group,omitempty"`

	// Subsets subset properties
	Subsets *map[string]struct {
		Parallel bool `json:"parallel"`
	} `json:"subsets,omitempty"`

	// Topology object topology
	Topology *Topology `json:"topology,omitempty"`
	Updated  time.Time `json:"updated"`
}

// Kind defines model for kind.
type Kind = string

// MonitorUpdateQueued defines model for monitorUpdateQueued.
type MonitorUpdateQueued struct {
	OrchestrationId openapi_types.UUID `json:"orchestration_id"`
}

// NodeInfo defines model for nodeInfo.
type NodeInfo struct {
	// Labels labels is the list of node labels.
	Labels []NodeLabel `json:"labels"`

	// Nodename nodename is the name of the node where the labels and paths are coming from.
	Nodename string `json:"nodename"`

	// Paths paths is the list of node to storage array san paths.
	Paths []SanPath `json:"paths"`
}

// NodeLabel defines model for nodeLabel.
type NodeLabel struct {
	// Name name is the label name.
	Name string `json:"name"`

	// Value value is the label value.
	Value string `json:"value"`
}

// NodesInfo defines model for nodesInfo.
type NodesInfo = []NodeInfo

// ObjectConfig defines model for objectConfig.
type ObjectConfig struct {
	Data  map[string]interface{} `json:"data"`
	Mtime time.Time              `json:"mtime"`
}

// ObjectFile defines model for objectFile.
type ObjectFile struct {
	Data  []byte    `json:"data"`
	Mtime time.Time `json:"mtime"`
}

// ObjectPath defines model for objectPath.
type ObjectPath = string

// ObjectSelector defines model for objectSelector.
type ObjectSelector = []ObjectPath

// Orchestrate defines model for orchestrate.
type Orchestrate string

// PathRelation defines model for pathRelation.
type PathRelation = []string

// Placement object placement policy
type Placement string

// PostDaemonSubAction defines model for postDaemonSubAction.
type PostDaemonSubAction struct {
	Action PostDaemonSubActionAction `json:"action"`

	// Subs daemon component list
	Subs []string `json:"subs"`
}

// PostDaemonSubActionAction defines model for PostDaemonSubAction.Action.
type PostDaemonSubActionAction string

// PostInstanceStatus defines model for postInstanceStatus.
type PostInstanceStatus struct {
	Path   string         `json:"path"`
	Status InstanceStatus `json:"status"`
}

// PostNodeDRBDConfigRequestBody defines model for postNodeDRBDConfigRequestBody.
type PostNodeDRBDConfigRequestBody struct {
	AllocationId openapi_types.UUID `json:"allocation_id"`
	Data         []byte             `json:"data"`
}

// PostNodeMonitor defines model for postNodeMonitor.
type PostNodeMonitor struct {
	GlobalExpect *string `json:"global_expect,omitempty"`
	LocalExpect  *string `json:"local_expect,omitempty"`
	State        *string `json:"state,omitempty"`
}

// PostObjectAbort defines model for postObjectAbort.
type PostObjectAbort struct {
	Path string `json:"path"`
}

// PostObjectClear defines model for postObjectClear.
type PostObjectClear struct {
	Path string `json:"path"`
}

// PostObjectMonitor defines model for postObjectMonitor.
type PostObjectMonitor struct {
	GlobalExpect *string `json:"global_expect,omitempty"`
	LocalExpect  *string `json:"local_expect,omitempty"`
	Path         string  `json:"path"`
	State        *string `json:"state,omitempty"`
}

// PostObjectProgress defines model for postObjectProgress.
type PostObjectProgress struct {
	IsPartial *bool  `json:"is_partial,omitempty"`
	Path      string `json:"path"`
	SessionId string `json:"session_id"`
	State     string `json:"state"`
}

// PostObjectSwitchTo defines model for postObjectSwitchTo.
type PostObjectSwitchTo struct {
	Destination []string `json:"destination"`
	Path        string   `json:"path"`
}

// PostRelayMessage defines model for postRelayMessage.
type PostRelayMessage struct {
	ClusterId   string `json:"cluster_id"`
	ClusterName string `json:"cluster_name"`
	Msg         string `json:"msg"`
	Nodename    string `json:"nodename"`
}

// Priority scheduling priority of an object instance on a its node
type Priority = int

// Problem defines model for problem.
type Problem struct {
	// Detail A human-readable explanation specific to this occurrence of the
	// problem.
	Detail string `json:"detail"`

	// Status The HTTP status code ([RFC7231], Section 6) generated by the
	// origin server for this occurrence of the problem.
	Status int `json:"status"`

	// Title A short, human-readable summary of the problem type.  It SHOULD
	// NOT change from occurrence to occurrence of the problem, except
	// for purposes of localization (e.g., using proactive content
	// negotiation; see [RFC7231], Section 3.4).
	Title string `json:"title"`
}

// Provisioned service, instance or resource provisioned state
type Provisioned string

// RelayMessage defines model for relayMessage.
type RelayMessage struct {
	Addr        string    `json:"addr"`
	ClusterId   string    `json:"cluster_id"`
	ClusterName string    `json:"cluster_name"`
	Msg         string    `json:"msg"`
	Nodename    string    `json:"nodename"`
	Updated     time.Time `json:"updated"`
}

// RelayMessageList defines model for relayMessageList.
type RelayMessageList = []RelayMessage

// RelayMessages defines model for relayMessages.
type RelayMessages struct {
	Messages RelayMessageList `json:"messages"`
}

// ResourceExposedStatus defines model for resourceExposedStatus.
type ResourceExposedStatus struct {
	// Disable hints the resource ignores all state transition actions
	Disable *bool `json:"disable,omitempty"`

	// Encap indicates that the resource is handled by the encapsulated agents,
	// and ignored at the hypervisor level
	Encap *bool `json:"encap,omitempty"`

	// Info key-value pairs providing interesting information to collect
	// site-wide about this resource
	Info  *map[string]interface{} `json:"info,omitempty"`
	Label string                  `json:"label"`
	Log   *[]struct {
		Level   string `json:"level"`
		Message string `json:"message"`
	} `json:"log,omitempty"`

	// Monitor tells the daemon if it should trigger a monitor action when the
	// resource is not up
	Monitor *bool `json:"monitor,omitempty"`

	// Optional is resource status aggregated into Overall instead of Avail instance status.
	// Errors in optional resource don't stop a state transition action
	Optional    *bool                    `json:"optional,omitempty"`
	Provisioned *ResourceProvisionStatus `json:"provisioned,omitempty"`
	Restart     *int                     `json:"restart,omitempty"`
	Rid         ResourceId               `json:"rid"`

	// Standby resource should always be up, even after a stop state transition action
	Standby *bool  `json:"standby,omitempty"`
	Status  Status `json:"status"`

	// Subset the name of the subset this resource is assigned to
	Subset *string   `json:"subset,omitempty"`
	Tags   *[]string `json:"tags,omitempty"`
	Type   string    `json:"type"`
}

// ResourceId defines model for resourceId.
type ResourceId = string

// ResourceProvisionStatus defines model for resourceProvisionStatus.
type ResourceProvisionStatus struct {
	Mtime *time.Time `json:"mtime,omitempty"`

	// State service, instance or resource provisioned state
	State Provisioned `json:"state"`
}

// Role defines model for role.
type Role string

// SanPath defines model for sanPath.
type SanPath struct {
	// Initiator initiator is the host side san path endpoint.
	Initiator SanPathInitiator `json:"initiator"`

	// Target target is the storage array side san path endpoint.
	Target SanPathTarget `json:"target"`
}

// SanPathInitiator initiator is the host side san path endpoint.
type SanPathInitiator struct {
	// Name name is a worldwide unique path endpoint identifier.
	Name *string `json:"name,omitempty"`

	// Type type is the endpoint type.
	Type *string `json:"type,omitempty"`
}

// SanPathTarget target is the storage array side san path endpoint.
type SanPathTarget struct {
	// Name name is a worldwide unique path endpoint identifier.
	Name *string `json:"name,omitempty"`

	// Type type is a the endpoint type.
	Type *string `json:"type,omitempty"`
}

// Severity defines model for severity.
type Severity = string

// Status defines model for status.
type Status = string

// Topology object topology
type Topology string

// DRBDConfigName defines model for DRBDConfigName.
type DRBDConfigName = string

// QueryDuration defines model for queryDuration.
type QueryDuration = string

// QueryEventFilter defines model for queryEventFilter.
type QueryEventFilter = []string

// QueryLimit defines model for queryLimit.
type QueryLimit = int64

// QueryNamespaceOptional defines model for queryNamespaceOptional.
type QueryNamespaceOptional = string

// QueryObjectPath defines model for queryObjectPath.
type QueryObjectPath = string

// QueryObjectSelector defines model for queryObjectSelector.
type QueryObjectSelector = string

// QueryRelativesOptional defines model for queryRelativesOptional.
type QueryRelativesOptional = bool

// QueryRelayClusterId defines model for queryRelayClusterId.
type QueryRelayClusterId = string

// QueryRelayNodename defines model for queryRelayNodename.
type QueryRelayNodename = string

// QueryRoles defines model for queryRoles.
type QueryRoles = []Role

// QuerySelectorOptional defines model for querySelectorOptional.
type QuerySelectorOptional = string

// N200 defines model for 200.
type N200 = Problem

// N400 defines model for 400.
type N400 = Problem

// N401 defines model for 401.
type N401 = Problem

// N403 defines model for 403.
type N403 = Problem

// N500 defines model for 500.
type N500 = Problem

// N503 defines model for 503.
type N503 = Problem

// PostAuthTokenParams defines parameters for PostAuthToken.
type PostAuthTokenParams struct {
	// Role list of api role
	Role *QueryRoles `form:"role,omitempty" json:"role,omitempty"`

	// Duration max token duration, maximum value 24h
	Duration *string `form:"duration,omitempty" json:"duration,omitempty"`
}

// GetDaemonEventsParams defines parameters for GetDaemonEvents.
type GetDaemonEventsParams struct {
	// Duration max duration
	Duration *QueryDuration `form:"duration,omitempty" json:"duration,omitempty"`

	// Limit limit items count
	Limit *QueryLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Filter list of event filter
	Filter *QueryEventFilter `form:"filter,omitempty" json:"filter,omitempty"`
}

// PostDaemonJoinParams defines parameters for PostDaemonJoin.
type PostDaemonJoinParams struct {
	// Node The node to add to cluster nodes
	Node string `form:"node" json:"node"`
}

// PostDaemonLeaveParams defines parameters for PostDaemonLeave.
type PostDaemonLeaveParams struct {
	// Node The leaving cluster node
	Node string `form:"node" json:"node"`
}

// GetDaemonStatusParams defines parameters for GetDaemonStatus.
type GetDaemonStatusParams struct {
	// Namespace namespace
	Namespace *QueryNamespaceOptional `form:"namespace,omitempty" json:"namespace,omitempty"`

	// Relatives relatives
	Relatives *QueryRelativesOptional `form:"relatives,omitempty" json:"relatives,omitempty"`

	// Selector selector
	Selector *QuerySelectorOptional `form:"selector,omitempty" json:"selector,omitempty"`
}

// GetNetworksParams defines parameters for GetNetworks.
type GetNetworksParams struct {
	// Name the name of a cluster backend network
	Name *string `form:"name,omitempty" json:"name,omitempty"`
}

// GetNodeDRBDConfigParams defines parameters for GetNodeDRBDConfig.
type GetNodeDRBDConfigParams struct {
	// Name the full path of the file is deduced from the name
	Name DRBDConfigName `form:"name" json:"name"`
}

// PostNodeDRBDConfigParams defines parameters for PostNodeDRBDConfig.
type PostNodeDRBDConfigParams struct {
	// Name the full path of the file is deduced from the name
	Name DRBDConfigName `form:"name" json:"name"`
}

// GetObjectConfigParams defines parameters for GetObjectConfig.
type GetObjectConfigParams struct {
	// Path object path
	Path QueryObjectPath `form:"path" json:"path"`

	// Evaluate evaluate
	Evaluate *bool `form:"evaluate,omitempty" json:"evaluate,omitempty"`

	// Impersonate impersonate the evaluation as node
	Impersonate *string `form:"impersonate,omitempty" json:"impersonate,omitempty"`
}

// GetObjectFileParams defines parameters for GetObjectFile.
type GetObjectFileParams struct {
	// Path object path
	Path QueryObjectPath `form:"path" json:"path"`
}

// GetObjectSelectorParams defines parameters for GetObjectSelector.
type GetObjectSelectorParams struct {
	// Selector object selector
	Selector QueryObjectSelector `form:"selector" json:"selector"`
}

// GetRelayMessageParams defines parameters for GetRelayMessage.
type GetRelayMessageParams struct {
	// Nodename the nodename component of the slot id on the relay
	Nodename *QueryRelayNodename `form:"nodename,omitempty" json:"nodename,omitempty"`

	// ClusterId the cluster id component of the slot id on the relay
	ClusterId *QueryRelayClusterId `form:"cluster_id,omitempty" json:"cluster_id,omitempty"`
}

// PostDaemonLogsControlJSONRequestBody defines body for PostDaemonLogsControl for application/json ContentType.
type PostDaemonLogsControlJSONRequestBody = PostDaemonLogsControl

// PostDaemonSubActionJSONRequestBody defines body for PostDaemonSubAction for application/json ContentType.
type PostDaemonSubActionJSONRequestBody = PostDaemonSubAction

// PostInstanceStatusJSONRequestBody defines body for PostInstanceStatus for application/json ContentType.
type PostInstanceStatusJSONRequestBody = PostInstanceStatus

// PostNodeDRBDConfigJSONRequestBody defines body for PostNodeDRBDConfig for application/json ContentType.
type PostNodeDRBDConfigJSONRequestBody = PostNodeDRBDConfigRequestBody

// PostNodeMonitorJSONRequestBody defines body for PostNodeMonitor for application/json ContentType.
type PostNodeMonitorJSONRequestBody = PostNodeMonitor

// PostObjectAbortJSONRequestBody defines body for PostObjectAbort for application/json ContentType.
type PostObjectAbortJSONRequestBody = PostObjectAbort

// PostObjectClearJSONRequestBody defines body for PostObjectClear for application/json ContentType.
type PostObjectClearJSONRequestBody = PostObjectClear

// PostObjectMonitorJSONRequestBody defines body for PostObjectMonitor for application/json ContentType.
type PostObjectMonitorJSONRequestBody = PostObjectMonitor

// PostObjectProgressJSONRequestBody defines body for PostObjectProgress for application/json ContentType.
type PostObjectProgressJSONRequestBody = PostObjectProgress

// PostObjectSwitchToJSONRequestBody defines body for PostObjectSwitchTo for application/json ContentType.
type PostObjectSwitchToJSONRequestBody = PostObjectSwitchTo

// PostRelayMessageJSONRequestBody defines body for PostRelayMessage for application/json ContentType.
type PostRelayMessageJSONRequestBody = PostRelayMessage

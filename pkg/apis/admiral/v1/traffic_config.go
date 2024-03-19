package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TrafficConfig is the Schema for the trafficconfigs API
type TrafficConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TrafficConfigSpec `json:"spec,omitempty"`
	// +optional
	Status TrafficConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficConfigList contains a list of TrafficConfig
type TrafficConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficConfig `json:"items"`
}

type TrafficConfigSpec struct {
	WorkloadEnv []string     `json:"workloadEnvs"`
	EdgeService *EdgeService `json:"edgeService"`
	QuotaGroup  *QuotaGroup  `json:"quotaGroup"`
}

type EdgeService struct {
	DynamicRouting []*DynamicRouting `json:"dynamicRouting,omitempty"`
	Filters        []*Filter         `json:"filters"`
	Routes         []*Route          `json:"routes"`
	Targets        []*Target         `json:"targets"`
	TargetGroups   []*TargetGroup    `json:"targetGroups"`
}

type Target struct {
	Name          string `json:"name"`
	DNS           string `json:"dns,omitempty"`
	MeshDNS       string `json:"meshDNS"`
	Port          int    `json:"port"`
	SocketTimeout int    `json:"socketTimeout"`
}

type TargetGroup struct {
	Name         string         `json:"name"`
	Weights      []*Weight      `json:"weights"`
	AppOverrides []*AppOverride `json:"appOverrides,omitempty"`
}

type AppOverride struct {
	AssetAlias string    `json:"assetAlias"`
	AssetID    string    `json:"assetID"` // assetID is just a UUID string
	Weights    []*Weight `json:"weights"`
}

type Weight struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

type QuotaGroup struct {
	TotalQuotaGroup []*TotalQuotaGroup `json:"totalQuotaGroups"`
	AppQuotaGroups  []*AppQuotaGroup   `json:"appQuotaGroups,omitempty"`
}

type Route struct {
	Name                 string    `json:"name"`
	Inbound              string    `json:"inbound"`
	Outbound             string    `json:"outbound"`
	FilterSelector       string    `json:"filterSelector"`
	WorkloadEnvSelectors []string  `json:"workloadEnvSelectors"`
	Timeout              int       `json:"timeout"`
	Config               []*Config `json:"config"`
}

type Config struct {
	TargetGroupSelector string `json:"targetGroupSelector"`
	TargetSelector      string `json:"targetSelector"`
}

type Filter struct {
	Name    string   `json:"name"`
	Retries Retry    `json:"retries,omitempty"`
	Options []string `json:"options,omitempty"`
}

type Retry struct {
	Attempts      int    `json:"attempts"`
	PerTryTimeout string `json:"perTryTimeout"`
}

type DynamicRouting struct {
	Name              string `json:"name"`
	Url               string `json:"url"`
	CacheKeyAlgorithm string `json:"cacheKeyAlgorithm"`
	TtlSec            int    `json:"ttlSec"`
	Local             bool   `json:"local"`
}

type TotalQuotaGroup struct {
	Name                 string               `json:"name"`
	Description          string               `json:"description"`
	Quotas               []*Quota             `json:"quotas"`
	WorkloadEnvSelectors []string             `json:"workloadEnvSelectors"`
	RegionLevelLimit     bool                 `json:"regionLevelLimit,omitempty"`
	CPULimit             *int                 `json:"cpuLimit,omitempty"`
	MemoryLimit          *int                 `json:"memoryLimit,omitempty"`
	PodLevelThreshold    *int                 `json:"podLevelThreshold,omitempty"`
	FailureModeBehaviour string               `json:"failureModeBehaviour,omitempty"`
	AdaptiveConcurrency  *AdaptiveConcurrency `json:"adaptiveConcurrency,omitempty"`
}
type AppQuotaGroup struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Quotas               []*Quota `json:"quotas"`
	AssociatedApps       []string `json:"associatedApps"`
	WorkloadEnvSelectors []string `json:"workloadEnvSelectors"`
}

type AdaptiveConcurrency struct {
	LatencyThreshold          string   `json:"latencyThreshold"`
	SkippedURLs               []string `json:"skippedURLs"`
	SampleAggregatePercentile int      `json:"sampleAggregatePercentile"`
	ConcurrencyUpdateInterval string   `json:"concurrencyUpdateInterval"`
	MinRTTCalInterval         string   `json:"minRTTCalInterval"`
	MinRTTCalJitter           int      `json:"minRTTCalJitter"`
	MinRTTCalRequestCount     int      `json:"minRTTCalRequestCount"`
	MinRTTCalMinConcurrency   int      `json:"minRTTCalMinConcurrency"`
	Enabled                   bool     `json:"enabled"`
}

type Quota struct {
	Name       string    `json:"name"`
	TimePeriod string    `json:"timePeriod"`
	MaxAmount  int       `json:"maxAmount"`
	KeyType    string    `json:"keyType,omitempty"`
	Algorithm  string    `json:"algorithm,omitempty"`
	Behaviour  string    `json:"behaviour,omitempty"`
	Rule       string    `json:"rule,omitempty"`
	Path       string    `json:"path,omitempty"`
	Methods    []string  `json:"methods,omitempty"`
	Headers    []*Header `json:"headers,omitempty"`
}

type Header struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Condition string `json:"condition"` // EQUALS, PREFIX, CONTAINS, REGEX
}

// TrafficConfigStatus defines the observed state of TrafficConfig
type TrafficConfigStatus struct {
	Message                  string      `json:"message"`
	LastAppliedConfigVersion string      `json:"lastAppliedConfigVersion"`
	LastUpdateTime           metav1.Time `json:"lastUpdateTime"`
	Status                   bool        `json:"status"`
}

package config

import "regexp"

// Config is the vCluster config. This struct describes valid helm values for vCluster as well as configuration used by the vCluster binary itself.
type Config struct {
	// ExportKubeConfig describes how vCluster should export the vCluster kube config
	ExportKubeConfig ExportKubeConfig `json:"exportKubeConfig,omitempty"`

	// ControlPlane holds options how to configure the vCluster control-plane
	ControlPlane ControlPlane `json:"controlPlane,omitempty"`

	// Sync describes how to sync resources from the vCluster to host cluster and back
	Sync Sync `json:"sync,omitempty"`

	// Observability holds options to proxy metrics from the host cluster into the vCluster
	Observability Observability `json:"observability,omitempty"`

	// Networking are networking options related to the vCluster
	Networking Networking `json:"networking,omitempty"`

	// Plugins define what vCluster plugins to load.
	Plugins map[string]Plugins `json:"plugins,omitempty"`

	// Policies defines policies to enforce for the vCluster deployment as well as within the vCluster
	Policies Policies `json:"policies,omitempty"`

	// RBAC are role based access control options for the vCluster
	RBAC RBAC `json:"rbac,omitempty"`

	// Experimental are alpha features for vCluster. Configuration here might change, so be careful with this.
	Experimental Experimental `json:"experimental,omitempty"`

	// Telemetry is the configuration related to telemetry gathered about vCluster usage.
	Telemetry Telemetry `json:"telemetry,omitempty"`

	// Platform holds options how vCluster should connect to vCluster platform.
	Platform Platform `json:"platform,omitempty"`

	// ServiceCIDR holds the service cidr for the vCluster. Please do not use that option anymore.
	ServiceCIDR string `json:"serviceCIDR,omitempty"`

	// Pro specifies if vCluster pro should be used. This is automatically inferred in newer versions. Please do not use that option anymore.
	Pro bool `json:"pro,omitempty"`

	// Plugin specifies what vCluster plugins to enable. Please use "plugins" instead. Please do not use that option anymore.
	Plugin map[string]Plugin `json:"plugin,omitempty"`
}

// ExportKubeConfig describes how vCluster should export the vCluster kube config
type ExportKubeConfig struct {
	// Context is the name of the context within the generated kube config to use.
	Context string `json:"context"`

	// Server can be used to override the default https://localhost:8443 and specify a custom hostname for the
	// generated kube-config.
	Server string `json:"server"`

	// Secret defines in which secret in the host cluster the generated kube-config should be stored.
	// If this is not defined, vCluster will only create it at `vc-NAME`. If another name is specified here
	// vCluster will also create the config in this other secret.
	Secret ExportKubeConfigSecretReference `json:"secret,omitempty"`
}

// ExportKubeConfigSecretReference defines in which secret in the host cluster the generated kube-config should be stored.
// If this is not defined, vCluster will only create it at `vc-NAME`. If another name is specified here
// vCluster will also create the config in this other secret.
type ExportKubeConfigSecretReference struct {
	// Name is the name of the secret where the kube config should get stored.
	Name string `json:"name,omitempty"`

	// Namespace defines the namespace where the kube config secret should get stored. If this is not equal to the namespace
	// where the vCluster is deployed, you need to make sure vCluster has access to this other namespace.
	Namespace string `json:"namespace,omitempty"`
}

type Sync struct {
	// ToHost configures what resources should get synced from the vCluster to the host cluster.
	ToHost SyncToHost `json:"toHost,omitempty"`

	// FromHost configures what resources should get purely synced from the host cluster to the vCluster.
	FromHost SyncFromHost `json:"fromHost,omitempty"`
}

type SyncToHost struct {
	// Services defines if services created within the vCluster should get synced to the host cluster.
	Services EnableSwitch `json:"services,omitempty"`
	// Endpoints defines if endpoints created within the vCluster should get synced to the host cluster.
	Endpoints EnableSwitch `json:"endpoints,omitempty"`
	// Ingresses defines if ingresses created within the vCluster should get synced to the host cluster.
	Ingresses EnableSwitch `json:"ingresses,omitempty"`
	// PriorityClasses defines if priority classes created within the vCluster should get synced to the host cluster.
	PriorityClasses EnableSwitch `json:"priorityClasses,omitempty"`
	// NetworkPolicies defines if network policies created within the vCluster should get synced to the host cluster.
	NetworkPolicies EnableSwitch `json:"networkPolicies,omitempty"`
	// VolumeSnapshots defines if volume snapshots created within the vCluster should get synced to the host cluster.
	VolumeSnapshots EnableSwitch `json:"volumeSnapshots,omitempty"`
	// PodDisruptionBudgets defines if pod disruption budgets created within the vCluster should get synced to the host cluster.
	PodDisruptionBudgets EnableSwitch `json:"podDisruptionBudgets,omitempty"`
	// ServiceAccounts defines if service accounts created within the vCluster should get synced to the host cluster.
	ServiceAccounts EnableSwitch `json:"serviceAccounts,omitempty"`
	// StorageClasses defines if storage classes created within the vCluster should get synced to the host cluster.
	StorageClasses EnableSwitch `json:"storageClasses,omitempty"`
	// PersistentVolumes defines if persistent volumes created within the vCluster should get synced to the host cluster.
	PersistentVolumes EnableSwitch `json:"persistentVolumes,omitempty"`
	// PersistentVolumeClaims defines if persistent volume claims created within the vCluster should get synced to the host cluster.
	PersistentVolumeClaims EnableSwitch `json:"persistentVolumeClaims,omitempty"`
	// ConfigMaps defines if config maps created within the vCluster should get synced to the host cluster.
	ConfigMaps SyncAllResource `json:"configMaps,omitempty"`
	// Secrets defines if secrets created within the vCluster should get synced to the host cluster.
	Secrets SyncAllResource `json:"secrets,omitempty"`
	// Pods defines if pods created within the vCluster should get synced to the host cluster.
	Pods SyncPods `json:"pods,omitempty"`
}

type SyncFromHost struct {
	// CSIDrivers defines if csi drivers should get synced from the host cluster to the vCluster, but not back.
	CSIDrivers EnableSwitch `json:"csiDrivers,omitempty"`
	// CSINodes defines if csi nodes should get synced from the host cluster to the vCluster, but not back.
	CSINodes EnableSwitch `json:"csiNodes,omitempty"`
	// CSIStorageCapacities defines if csi storage capacities should get synced from the host cluster to the vCluster, but not back.
	CSIStorageCapacities EnableSwitch `json:"csiStorageCapacities,omitempty"`
	// IngressClasses defines if ingress classes should get synced from the host cluster to the vCluster, but not back.
	IngressClasses EnableSwitch `json:"ingressClasses,omitempty"`
	// Events defines if events should get synced from the host cluster to the vCluster, but not back.
	Events EnableSwitch `json:"events,omitempty"`
	// StorageClasses defines if storage classes should get synced from the host cluster to the vCluster, but not back.
	StorageClasses EnableSwitch `json:"storageClasses,omitempty"`
	// Nodes defines if nodes should get synced from the host cluster to the vCluster, but not back.
	Nodes SyncNodes `json:"nodes,omitempty"`
}

type EnableSwitch struct {
	// Enabled defines if this option should be enabled.
	Enabled bool `json:"enabled,omitempty"`
}

type SyncAllResource struct {
	// Enabled defines if this option should be enabled.
	Enabled bool `json:"enabled,omitempty"`

	// All defines if all resources of that type should get synced or only the necessary ones that are needed.
	All bool `json:"all,omitempty"`
}

type SyncPods struct {
	// Enabled defines if pod syncing should be enabled.
	Enabled bool `json:"enabled,omitempty"`

	// TranslateImage maps an image to another image that should be used instead. For example this can be used to rewrite
	// a certain image that is used within the vCluster to be another image on the host cluster
	TranslateImage map[string]string `json:"translateImage,omitempty"`

	// EnforceTolerations will add the specified tolerations to all pods synced by the vCluster.
	EnforceTolerations []string `json:"enforceTolerations,omitempty"`

	// UseSecretsForSATokens will use secrets to save the generated service account tokens by vCluster instead of using a
	// pod annotation.
	UseSecretsForSATokens bool `json:"useSecretsForSATokens,omitempty"`

	// RewriteHosts is a special option needed to rewrite statefulset containers to allow the correct FQDN. vCluster will add
	// a small container to each stateful set pod that will initially rewrite the /etc/hosts file to match the FQDN expected by
	// the vCluster.
	RewriteHosts SyncRewriteHosts `json:"rewriteHosts,omitempty"`
}

type SyncRewriteHosts struct {
	// Enabled specifies if rewriting stateful set pods should be enabled.
	Enabled bool `json:"enabled,omitempty"`

	// InitContainerImage is the image vCluster should use to rewrite this FQDN.
	InitContainerImage string `json:"initContainerImage,omitempty"`
}

type SyncNodes struct {
	// Enabled specifies if syncing real nodes should be enabled. If this is disabled, vCluster will create fake nodes instead.
	Enabled bool `json:"enabled,omitempty"`

	// SyncAll specifies if all nodes should get synced by vCluster from the host to the vCluster or only the ones where pods are assigned to.
	SyncAll bool `json:"syncAll,omitempty"`

	// SyncLabelsTaints enables syncing labels and taints from the vCluster to the host cluster. If this is enabled someone within the vCluster will be able to change the labels and taints of the host cluster node.
	SyncLabelsTaints bool `json:"syncLabelsTaints,omitempty"`

	// ClearImageStatus will erase the image status when syncing a node. This allows to hide images that are pulled by the node.
	ClearImageStatus bool `json:"clearImageStatus,omitempty"`

	// Selector can be used to define more granular what nodes should get synced from the host cluster to the vCluster.
	Selector SyncNodeSelector `json:"selector,omitempty"`
}

type SyncNodeSelector struct {
	// Labels are the node labels used to sync nodes from host cluster to vCluster. This will also set the node selector when syncing a pod from vCluster to host cluster to the same value.
	Labels map[string]string `json:"labels,omitempty"`
}

type Observability struct {
	// Metrics allows to proxy metrics server apis from host to vCluster.
	Metrics ObservabilityMetrics `json:"metrics,omitempty"`
}

type ControlPlaneObservability struct {
	// ServiceMonitor can be used to automatically create a service monitor for vCluster deployment itself.
	ServiceMonitor ServiceMonitor `json:"serviceMonitor,omitempty"`
}

type ServiceMonitor struct {
	// Enabled configures if helm should create the service monitor.
	Enabled bool `json:"enabled,omitempty"`

	// Labels are the extra labels to add to the service monitor.
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations are the extra annotations to add to the service monitor.
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ObservabilityMetrics struct {
	// Proxy holds the configuration what metrics-server apis should get proxied.
	Proxy MetricsProxy `json:"proxy,omitempty"`
}

type MetricsProxy struct {
	// Nodes defines if metrics-server nodes api should get proxied from host to vCluster.
	Nodes EnableSwitch `json:"nodes,omitempty"`

	// Pods defines if metrics-server pods api should get proxied from host to vCluster.
	Pods EnableSwitch `json:"pods,omitempty"`
}

type Networking struct {
	// ReplicateServices allows replicating services from the host within the vCluster or the other way around.
	ReplicateServices ReplicateServices `json:"replicateServices,omitempty"`

	// ResolveServices allows to define extra DNS rules. This only works if embedded coredns is configured.
	ResolveServices []ResolveServices `json:"resolveServices,omitempty"`

	// Advanced holds advanced network options.
	Advanced NetworkingAdvanced `json:"advanced,omitempty"`
}

type ReplicateServices struct {
	// ToHost defines the services that should get synced from vCluster to the host cluster. If services are
	// synced to a different namespace than the vCluster is in, additional permissions for the other namespace
	// are required.
	ToHost []ServiceMapping `json:"toHost,omitempty"`

	// FromHost defines the services that should get synced from the host to the vCluster.
	FromHost []ServiceMapping `json:"fromHost,omitempty"`
}

type ServiceMapping struct {
	// From is the service that should get synced. Can be either in the form name or namespace/name.
	From string `json:"from,omitempty"`
	// To is the target service that it should get synced to. Can be either in the form name or namespace/name.
	To string `json:"to,omitempty"`
}

type ResolveServices struct {
	Service string               `json:"service,omitempty"`
	Target  ResolveServiceTarget `json:"target,omitempty"`
}

type ResolveServiceTarget struct {
	VCluster ResolveServiceService  `json:"vcluster,omitempty"`
	Host     ResolveServiceService  `json:"host,omitempty"`
	External ResolveServiceHostname `json:"external,omitempty"`
}

type ResolveServiceService struct {
	Service string `json:"service,omitempty"`
}

type ResolveServiceHostname struct {
	Hostname string `json:"hostname,omitempty"`
}

type NetworkingAdvanced struct {
	// ClusterDomain is the Kubernetes cluster domain to use within the vCluster.
	ClusterDomain string `json:"clusterDomain,omitempty"`

	// FallbackHostCluster allows to fallback dns to the host cluster. This is useful if you want to reach host services without
	// any other modification. You will need to provide a namespace for the service, e.g. my-other-service.my-other-namespace
	FallbackHostCluster bool `json:"fallbackHostCluster,omitempty"`

	// ProxyKubelets allows rewriting certain metrics and stats from the Kubelet to "fake" this for applications such as
	// prometheus or other node exporters.
	ProxyKubelets NetworkProxyKubelets `json:"proxyKubelets,omitempty"`
}

type NetworkProxyKubelets struct {
	// ByHostname will add a special vCluster hostname to the nodes where the node can be reached at. This doesn't work
	// for all applications, e.g. prometheus requires a node ip.
	ByHostname bool `json:"byHostname,omitempty"`

	// ByIP will create a separate service in the host cluster for every node that will point to vCluster and will be used to
	// route traffic.
	ByIP bool `json:"byIP,omitempty"`
}

type Plugin struct {
	Plugins `json:",inline"`

	// Version is the plugin version, this is only needed for legacy plugins.
	Version        string                 `json:"version,omitempty"`
	Env            []interface{}          `json:"env,omitempty"`
	EnvFrom        []interface{}          `json:"envFrom,omitempty"`
	Lifecycle      map[string]interface{} `json:"lifecycle,omitempty"`
	LivenessProbe  map[string]interface{} `json:"livenessProbe,omitempty"`
	ReadinessProbe map[string]interface{} `json:"readinessProbe,omitempty"`
	StartupProbe   map[string]interface{} `json:"startupProbe,omitempty"`
	WorkingDir     string                 `json:"workingDir,omitempty"`
	Optional       bool                   `json:"optional,omitempty"`
}

type Plugins struct {
	// Name is the name of the init-container and NOT the plugin name
	Name string `json:"name,omitempty"`
	// Command is the command that should be used for the init container
	Command []string `json:"command,omitempty"`
	// Args are the arguments that should be used for the init container
	Args []string `json:"args,omitempty"`
	// Image is the container image that should be used for the plugin
	Image string `json:"image,omitempty"`
	// ImagePullPolicy is the pull policy to use for the container image
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`
	// Config is the plugin config to use. This can be arbitrary config used for the plugin.
	Config map[string]interface{} `json:"config,omitempty"`
	// SecurityContext is the container security context used for the init container
	SecurityContext map[string]interface{} `json:"securityContext,omitempty"`
	// Resources are the container resources used for the init container
	Resources map[string]interface{} `json:"resources,omitempty"`
	// VolumeMounts are extra volume mounts for the init container
	VolumeMounts []interface{} `json:"volumeMounts,omitempty"`
	// RBAC holds additional rbac configuration for the plugin
	RBAC PluginsRBAC `json:"rbac,omitempty"`
}

type PluginsRBAC struct {
	// Role holds extra vCluster role permissions for the plugin
	Role PluginsExtraRules `json:"role,omitempty"`
	// ClusterRole holds extra vCluster cluster role permissions required for the plugin
	ClusterRole PluginsExtraRules `json:"clusterRole,omitempty"`
}

type PluginsExtraRules struct {
	// ExtraRules are extra rbac permissions roles that will be added to role or cluster role
	ExtraRules []RBACPolicyRule `json:"extraRules,omitempty"`
}

type RBACPolicyRule struct {
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds contained in this rule. '*' represents all verbs.
	Verbs []string `json:"verbs"`

	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of
	// the enumerated resources in any API group will be allowed. "" represents the core API group and "*" represents all API groups.
	APIGroups []string `json:"apiGroups,omitempty"`
	// Resources is a list of resources this rule applies to. '*' represents all resources.
	Resources []string `json:"resources,omitempty"`
	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
	ResourceNames []string `json:"resourceNames,omitempty"`

	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path
	// Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding.
	// Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
	NonResourceURLs []string `json:"nonResourceURLs,omitempty"`
}

type ControlPlane struct {
	// Distro holds vCluster related distro options.
	Distro Distro `json:"distro,omitempty"`
	// BackingStore defines which backing store to use for vCluster. If not defined will fallback to the default distro backing store.
	BackingStore BackingStore `json:"backingStore,omitempty"`
	// CoreDNS defines everything coredns related.
	CoreDNS CoreDNS `json:"coredns,omitempty"`
	// Proxy defines options for the vCluster control plane proxy that is used to do authentication and intercept requests.
	Proxy ControlPlaneProxy `json:"proxy,omitempty"`
	// Service defines options for the vCluster service deployed by helm.
	Service ControlPlaneService `json:"service,omitempty"`
	// Ingress defines options for the vCluster ingress deployed by helm.
	Ingress ControlPlaneIngress `json:"ingress,omitempty"`
	// StatefulSet defines options for the vCluster statefulSet deployed by helm.
	StatefulSet ControlPlaneStatefulSet `json:"statefulSet,omitempty"`
	// HostPathMapper defines if vCluster should rewrite host paths.
	HostPathMapper HostPathMapper `json:"hostPathMapper,omitempty"`
	// Observability defines if a service monitor should get deployed by helm.
	Observability ControlPlaneObservability `json:"observability,omitempty"`
	// Advanced holds additional configuration for the vCluster control plane.
	Advanced ControlPlaneAdvanced `json:"advanced,omitempty"`
}

type ControlPlaneStatefulSet struct {
	LabelsAndAnnotations `json:",inline"`

	// Image is the image for the controlPlane statefulSet container
	Image Image `json:"image,omitempty"`
	// ImagePullPolicy is the policy how to pull the image.
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`
	// WorkingDir specifies in what folder the main process should get started.
	WorkingDir string `json:"workingDir,omitempty"`
	// Command allows you to override the main command.
	Command []string `json:"command,omitempty"`
	// Args allows you to override the main arguments.
	Args []string `json:"args,omitempty"`
	// Env are additional environment variables for the statefulSet container.
	Env []map[string]interface{} `json:"env,omitempty"`
	// Pods are additional labels or annotations for the statefulSet pod.
	Pods LabelsAndAnnotations `json:"pods,omitempty"`

	// Probes enables or disables the main container probes.
	Probes ControlPlaneProbes `json:"probes,omitempty"`
	// Security defines pod or container security context.
	Security ControlPlaneSecurity `json:"security,omitempty"`
	// Persistence defines options around persistence for the statefulSet.
	Persistence ControlPlanePersistence `json:"persistence,omitempty"`
	// Scheduling holds options related to scheduling.
	Scheduling ControlPlaneScheduling `json:"scheduling,omitempty"`
	// HighAvailability holds options related to high availability.
	HighAvailability ControlPlaneHighAvailability `json:"highAvailability,omitempty"`
	// Resources are the resource requests and limits for the statefulSet container.
	Resources Resources `json:"resources,omitempty"`
}

type Distro struct {
	// K3S holds k3s relevant configuration.
	K3S DistroK3s `json:"k3s,omitempty"`
	// K0S holds k0s relevant configuration.
	K0S DistroK0s `json:"k0s,omitempty"`
	// K8S holds k8s relevant configuration.
	K8S DistroK8s `json:"k8s,omitempty"`
	// EKS holds eks relevant configuration.
	EKS DistroK8s `json:"eks,omitempty"`
}

type DistroK3s struct {
	// Enabled specifies if the k3s distro should be enabled. Only one distro can be enabled at the same time.
	Enabled bool `json:"enabled,omitempty"`
	// Token is the k3s token to use. If empty, vCluster will choose one.
	Token string `json:"token,omitempty"`

	DistroCommon    `json:",inline"`
	DistroContainer `json:",inline"`
}

type DistroK8s struct {
	// Enabled specifies if the k8s distro should be enabled. Only one distro can be enabled at the same time.
	Enabled bool `json:"enabled,omitempty"`

	// APIServer holds configuration specific to starting the api server.
	APIServer DistroContainerDisabled `json:"apiServer,omitempty"`
	// ControllerManager holds configuration specific to starting the scheduler.
	ControllerManager DistroContainerDisabled `json:"controllerManager,omitempty"`
	// Scheduler holds configuration specific to starting the scheduler.
	Scheduler DistroContainer `json:"scheduler,omitempty"`

	DistroCommon `json:",inline"`
}

type DistroK0s struct {
	// Enabled specifies if the k0s distro should be enabled. Only one distro can be enabled at the same time.
	Enabled bool `json:"enabled,omitempty"`
	// Config allows you to override the k0s config passed to the k0s binary.
	Config string `json:"config,omitempty"`

	DistroCommon    `json:",inline"`
	DistroContainer `json:",inline"`
}

type DistroCommon struct {
	// Env are extra environment variables to use for the main container.
	Env []map[string]interface{} `json:"env,omitempty"`
	// SecurityContext can be used for the distro init container
	SecurityContext map[string]interface{} `json:"securityContext,omitempty"`
	// Resources are the resources for the distro init container
	Resources map[string]interface{} `json:"resources,omitempty"`
}

type DistroContainer struct {
	// Image is the distro image
	Image Image `json:"image,omitempty"`
	// ImagePullPolicy is the pull policy for the distro image
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`
	// Command is the command to start the distro binary. This will override the existing command.
	Command []string `json:"command,omitempty"`
	// ExtraArgs are additional arguments to pass to the distro binary.
	ExtraArgs []string `json:"extraArgs,omitempty"`
}

type DistroContainerDisabled struct {
	// Disabled signals this container should be disabled.
	Disabled bool `json:"disabled,omitempty"`
	// Image is the distro image
	Image Image `json:"image,omitempty"`
	// ImagePullPolicy is the pull policy for the distro image
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`
	// Command is the command to start the distro binary. This will override the existing command.
	Command []string `json:"command,omitempty"`
	// ExtraArgs are additional arguments to pass to the distro binary.
	ExtraArgs []string `json:"extraArgs,omitempty"`
}

type Image struct {
	// Repository is the registry and repository of the container image, e.g. my-registry.com/my-repo/my-image
	Repository string `json:"repository,omitempty"`
	// Tag is the tag of the container image, e.g. latest
	Tag string `json:"tag,omitempty"`
}

// LocalObjectReference contains enough information to let you locate the
// referenced object inside the same namespace.
type LocalObjectReference struct {
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
}

type VirtualClusterKubeConfig struct {
	// KubeConfig is the virtual cluster kube config path.
	KubeConfig string `json:"kubeConfig,omitempty"`
	// ServerCAKey is the server ca key path.
	ServerCAKey string `json:"serverCAKey,omitempty"`
	// ServerCAKey is the server ca cert path.
	ServerCACert string `json:"serverCACert,omitempty"`
	// ServerCAKey is the client ca cert path.
	ClientCACert string `json:"clientCACert,omitempty"`
	// RequestHeaderCACert is the request header ca cert path.
	RequestHeaderCACert string `json:"requestHeaderCACert,omitempty"`
}

type BackingStore struct {
	// EmbeddedEtcd defines to use embedded etcd as a storage backend for the vCluster
	EmbeddedEtcd EmbeddedEtcd `json:"embeddedEtcd,omitempty" product:"pro"`
	// ExternalEtcd defines to use an external etcd deployed by the helm chart as a storage backend for the vCluster
	ExternalEtcd ExternalEtcd `json:"externalEtcd,omitempty"`
}

type EmbeddedEtcd struct {
	// Enabled defines if the embedded etcd should be used.
	Enabled bool `json:"enabled,omitempty"`
	// MigrateFromExternalEtcd signals that vCluster should migrate from the external etcd.
	MigrateFromExternalEtcd bool `json:"migrateFromExternalEtcd,omitempty"`
}

type ExternalEtcd struct {
	// Enabled defines if the external etcd should be used.
	Enabled bool `json:"enabled,omitempty"`

	// StatefulSet holds options for the external etcd statefulSet.
	StatefulSet ExternalEtcdStatefulSet `json:"statefulSet,omitempty"`
	// Service holds options for the external etcd service.
	Service ExternalEtcdService `json:"service,omitempty"`
	// HeadlessService holds options for the external etcd headless service.
	HeadlessService ExternalEtcdHeadlessService `json:"headlessService,omitempty"`
}

type ExternalEtcdService struct {
	// Enabled defines if the etcd service should be deployed
	Enabled bool `json:"enabled,omitempty"`
	// Annotations are extra annotations for the external etcd service
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ExternalEtcdHeadlessService struct {
	// Enabled defines if the etcd headless service should be deployed
	Enabled bool `json:"enabled,omitempty"`
	// Annotations are extra annotations for the external etcd headless service
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ExternalEtcdStatefulSet struct {
	// Enabled defines if the statefulSet should be deployed
	Enabled bool `json:"enabled,omitempty"`
	// Image is the image to use for the external etcd statefulSet
	Image Image `json:"image,omitempty"`
	// ImagePullPolicy is the pull policy for the external etcd image
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`
	// Env are extra environment variables
	Env []map[string]interface{} `json:"env,omitempty"`
	// ExtraArgs are appended to the etcd command.
	ExtraArgs []string `json:"extraArgs,omitempty"`

	// Resources the etcd can consume
	Resources Resources `json:"resources,omitempty"`
	// Pods defines extra metadata for the etcd pods.
	Pods LabelsAndAnnotations `json:"pods,omitempty"`
	// HighAvailability are high availability options
	HighAvailability ExternalEtcdHighAvailability `json:"highAvailability,omitempty"`
	// Scheduling options for the etcd pods.
	Scheduling ControlPlaneScheduling `json:"scheduling,omitempty"`
	// Security options for the etcd pods.
	Security ControlPlaneSecurity `json:"security,omitempty"`
	// Persistence options for the etcd pods.
	Persistence ControlPlanePersistence `json:"persistence,omitempty"`

	LabelsAndAnnotations `json:",inline"`
}

type Resources struct {
	// Limits are resource limits for the container
	Limits map[string]interface{} `json:"limits,omitempty"`
	// Requests are minimal resources that will be consumed by the container
	Requests map[string]interface{} `json:"requests,omitempty"`
}

type ExternalEtcdHighAvailability struct {
	// Replicas are the amount of pods to use.
	Replicas int `json:"replicas,omitempty"`
}

type HostPathMapper struct {
	// Enabled specifies if the host path mapper will be used
	Enabled bool `json:"enabled,omitempty"`
	// Central specifies if the central host path mapper will be used
	Central bool `json:"central,omitempty" product:"pro"`
}

type CoreDNS struct {
	// Enabled defines if coredns is enabled
	Enabled bool `json:"enabled,omitempty"`
	// Embedded defines if vCluster will start the embedded coredns service
	Embedded bool `json:"embedded,omitempty" product:"pro"`
	// Service holds extra options for the coredns service deployed within the vCluster
	Service CoreDNSService `json:"service,omitempty"`
	// Deployment holds extra options for the coredns deployment deployed within the vCluster
	Deployment CoreDNSDeployment `json:"deployment,omitempty"`

	// OverwriteConfig can be used to overwrite the coredns config
	OverwriteConfig string `json:"overwriteConfig,omitempty"`
	// OverwriteManifests can be used to overwrite the coredns manifests used to deploy coredns
	OverwriteManifests string `json:"overwriteManifests,omitempty"`
}

type CoreDNSService struct {
	// Spec holds extra options for the coredns service
	Spec map[string]interface{} `json:"spec,omitempty"`

	LabelsAndAnnotations `json:",inline"`
}

type CoreDNSDeployment struct {
	// Image is the coredns image to use
	Image string `json:"image,omitempty"`
	// Replicas is the amount of coredns pods to run.
	Replicas int `json:"replicas,omitempty"`
	// NodeSelector is the node selector to use for coredns.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// Resources are the desired resources for coredns.
	Resources Resources `json:"resources,omitempty"`
	// Pods is additional metadata for the coredns pods.
	Pods LabelsAndAnnotations `json:"pods,omitempty"`

	LabelsAndAnnotations `json:",inline"`
}

type ControlPlaneProxy struct {
	// BindAddress under which the vCluster will expose the proxy.
	BindAddress string `json:"bindAddress,omitempty"`
	// Port under which the vCluster will expose the proxy.
	Port int `json:"port,omitempty"`
	// ExtraSANs are extra hostnames to sign the vCluster proxy certificate for.
	ExtraSANs []string `json:"extraSANs,omitempty"`
}

type ControlPlaneService struct {
	// Enabled defines if the control plane service should be enabled
	Enabled bool `json:"enabled,omitempty"`
	// KubeletNodePort is the node port where the fake kubelet is exposed. Defaults to 0.
	KubeletNodePort int `json:"kubeletNodePort,omitempty"`
	// HTTPSNodePort is the node port where https is exposed. Defaults to 0.
	HTTPSNodePort int `json:"httpsNodePort,omitempty"`
	// Spec allows you to configure extra service options.
	Spec map[string]interface{} `json:"spec,omitempty"`

	LabelsAndAnnotations `json:",inline"`
}

type ControlPlaneIngress struct {
	// Enabled defines if the control plane ingress should be enabled
	Enabled bool `json:"enabled,omitempty"`

	// Host is the host where vCluster will be reachable
	Host string `json:"host,omitempty"`
	// PathType is the path type of the ingress
	PathType string `json:"pathType,omitempty"`
	// Spec allows you to configure extra ingress options.
	Spec map[string]interface{} `json:"spec,omitempty"`

	LabelsAndAnnotations `json:",inline"`
}

type ControlPlaneHighAvailability struct {
	// Replicas is the amount of replicas to use for the statefulSet.
	Replicas int32 `json:"replicas,omitempty"`

	// LeaseDuration is the time to lease for the leader.
	LeaseDuration int `json:"leaseDuration,omitempty"`

	// RenewDeadline is the deadline to renew a lease for the leader.
	RenewDeadline int `json:"renewDeadline,omitempty"`

	// RetryPeriod is the time until a replica will retry to get a lease.
	RetryPeriod int `json:"retryPeriod,omitempty"`
}

type ControlPlaneAdvanced struct {
	// DefaultImageRegistry will be used as a prefix for all internal images deployed by vCluster or helm. This makes it easy to
	// upload all required vCluster images to a single private repository and set this value. Workload images are not affected by this.
	DefaultImageRegistry string `json:"defaultImageRegistry,omitempty"`

	// VirtualScheduler defines if a scheduler should be used within the vCluster or the scheduling decision for workloads will be made by the host cluster.
	VirtualScheduler EnableSwitch `json:"virtualScheduler,omitempty"`

	// ServiceAccount specifies options for the vCluster control-plane service account.
	ServiceAccount ControlPlaneServiceAccount `json:"serviceAccount,omitempty"`

	// WorkloadServiceAccount specifies options for the service account that will be used for the workloads that run within the vCluster.
	WorkloadServiceAccount ControlPlaneWorkloadServiceAccount `json:"workloadServiceAccount,omitempty"`

	// HeadlessService specifies options for the headless service used for the vCluster statefulSet.
	HeadlessService ControlPlaneHeadlessService `json:"headlessService,omitempty"`

	// GlobalMetadata is metadata that will be added to all resources deployed by helm.
	GlobalMetadata ControlPlaneGlobalMetadata `json:"globalMetadata,omitempty"`
}

type ControlPlaneHeadlessService struct {
	// Annotations are extra annotations for this resource.
	Annotations map[string]string `json:"annotations,omitempty"`
	// Labels are extra labels for this resource.
	Labels map[string]string `json:"labels,omitempty"`
}

type ControlPlanePersistence struct {
	// VolumeClaim can be used to configure the persistent volume claim.
	VolumeClaim VolumeClaim `json:"volumeClaim,omitempty"`
	// VolumeClaimTemplates defines the volumeClaimTemplates for the statefulSet
	VolumeClaimTemplates []map[string]interface{} `json:"volumeClaimTemplates,omitempty"`
	// AddVolumes defines extra volumes for the pod
	AddVolumes []map[string]interface{} `json:"addVolumes,omitempty"`
	// AddVolumeMounts defines extra volume mounts for the container
	AddVolumeMounts []VolumeMount `json:"addVolumeMounts,omitempty"`
}

type VolumeClaim struct {
	// Disabled signals to disable deploying a persistent volume claim. If false, vCluster will automatically determine
	// based on the chosen distro and other options if this is required.
	Disabled bool `json:"disabled,omitempty"`
	// AccessModes are the persistent volume claim access modes.
	AccessModes []string `json:"accessModes,omitempty"`
	// RetentionPolicy is the persistent volume claim retention policy.
	RetentionPolicy string `json:"retentionPolicy,omitempty"`
	// Size is the persistent volume claim storage size.
	Size string `json:"size,omitempty"`
	// StorageClass is the persistent volume claim storage class.
	StorageClass string `json:"storageClass,omitempty"`
}

// VolumeMount describes a mounting of a Volume within a container.
type VolumeMount struct {
	// This must match the Name of a Volume.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	// Defaults to false.
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"`
	// Path within the container at which the volume should be mounted.  Must
	// not contain ':'.
	MountPath string `json:"mountPath" protobuf:"bytes,3,opt,name=mountPath"`
	// Path within the volume from which the container's volume should be mounted.
	// Defaults to "" (volume's root).
	SubPath string `json:"subPath,omitempty" protobuf:"bytes,4,opt,name=subPath"`
	// mountPropagation determines how mounts are propagated from the host
	// to container and the other way around.
	// When not set, MountPropagationNone is used.
	// This field is beta in 1.10.
	MountPropagation *string `json:"mountPropagation,omitempty" protobuf:"bytes,5,opt,name=mountPropagation,casttype=MountPropagationMode"`
	// Expanded path within the volume from which the container's volume should be mounted.
	// Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.
	// Defaults to "" (volume's root).
	// SubPathExpr and SubPath are mutually exclusive.
	SubPathExpr string `json:"subPathExpr,omitempty" protobuf:"bytes,6,opt,name=subPathExpr"`
}

type ControlPlaneScheduling struct {
	// NodeSelector is the node selector to apply to the pod.
	NodeSelector map[string]interface{} `json:"nodeSelector,omitempty"`
	// Affinity is the affinity to apply to the pod.
	Affinity map[string]interface{} `json:"affinity,omitempty"`
	// Tolerations are the tolerations to apply to the pod.
	Tolerations []interface{} `json:"tolerations,omitempty"`
	// PriorityClassName is the priority class name for the the pod.
	PriorityClassName string `json:"priorityClassName,omitempty"`
	// PodManagementPolicy is the statefulSet pod management policy.
	PodManagementPolicy string `json:"podManagementPolicy,omitempty"`
	// TopologySpreadConstraints are the topology spread constraints for the pod.
	TopologySpreadConstraints []interface{} `json:"topologySpreadConstraints,omitempty"`
}

type ControlPlaneServiceAccount struct {
	// Enabled specifies if the service account should get deployed.
	Enabled bool `json:"enabled,omitempty"`
	// Name specifies what name to use for the service account.
	Name string `json:"name,omitempty"`
	// ImagePullSecrets defines extra image pull secrets for the service account.
	ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets,omitempty"`
	// Annotations are extra annotations for this resource.
	Annotations map[string]string `json:"annotations,omitempty"`
	// Labels are extra labels for this resource.
	Labels map[string]string `json:"labels,omitempty"`
}

type ControlPlaneWorkloadServiceAccount struct {
	// Enabled specifies if the service account for the workloads should get deployed.
	Enabled bool `json:"enabled,omitempty"`
	// Name specifies what name to use for the service account for the vCluster workloads.
	Name string `json:"name,omitempty"`
	// ImagePullSecrets defines extra image pull secrets for the workload service account.
	ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets,omitempty"`
	// Annotations are extra annotations for this resource.
	Annotations map[string]string `json:"annotations,omitempty"`
	// Labels are extra labels for this resource.
	Labels map[string]string `json:"labels,omitempty"`
}

type ControlPlaneProbes struct {
	// LivenessProbe specifies if the liveness probe for the container should be enabled
	LivenessProbe EnableSwitch `json:"livenessProbe,omitempty"`
	// ReadinessProbe specifies if the readiness probe for the container should be enabled
	ReadinessProbe EnableSwitch `json:"readinessProbe,omitempty"`
	// StartupProbe specifies if the startup probe for the container should be enabled
	StartupProbe EnableSwitch `json:"startupProbe,omitempty"`
}

type ControlPlaneSecurity struct {
	// PodSecurityContext specifies security context options on the pod level.
	PodSecurityContext map[string]interface{} `json:"podSecurityContext,omitempty"`
	// ContainerSecurityContext specifies security context options on the container level.
	ContainerSecurityContext map[string]interface{} `json:"containerSecurityContext,omitempty"`
}

type ControlPlaneGlobalMetadata struct {
	// Annotations are extra annotations for this resource.
	Annotations map[string]string `json:"annotations,omitempty"`
}

type LabelsAndAnnotations struct {
	// Annotations are extra annotations for this resource.
	Annotations map[string]string `json:"annotations,omitempty"`
	// Labels are extra labels for this resource.
	Labels map[string]string `json:"labels,omitempty"`
}

type Policies struct {
	// PodSecurityStandard that can be enforced can be one of: empty (""), baseline, restricted or privileged
	PodSecurityStandard string `json:"podSecurityStandard,omitempty"`
	// ResourceQuota specifies resource quota options.
	ResourceQuota ResourceQuota `json:"resourceQuota,omitempty"`
	// LimitRange specifies limit range options.
	LimitRange LimitRange `json:"limitRange,omitempty"`
	// NetworkPolicy specifies network policy options.
	NetworkPolicy NetworkPolicy `json:"networkPolicy,omitempty"`
	// CentralAdmission defines what validating or mutating webhooks should be enforced within the vCluster.
	CentralAdmission CentralAdmission `json:"centralAdmission,omitempty" product:"pro"`
}

type ResourceQuota struct {
	// Enabled defines if the resource quota should be enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Quota are the quota options
	Quota map[string]interface{} `json:"quota,omitempty"`
	// ScopeSelector is the resource quota scope selector
	ScopeSelector ScopeSelector `json:"scopeSelector,omitempty"`
	// Scopes are the resource quota scopes
	Scopes []string `json:"scopes,omitempty"`

	LabelsAndAnnotations `json:",inline"`
}

type ScopeSelector struct {
	MatchExpressions []LabelSelectorRequirement `json:"matchExpressions,omitempty"`
}

type LabelSelectorRequirement struct {
	// key is the label key that the selector applies to.
	Key string `json:"key"`
	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator string `json:"operator"`
	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. This array is replaced during a strategic
	// merge patch.
	Values []string `json:"values,omitempty"`
}

type LimitRange struct {
	// Enabled defines if the limit range should be deployed by vCluster.
	Enabled bool `json:"enabled,omitempty"`

	// Default are the default limits for the limit range
	Default map[string]interface{} `json:"default,omitempty"`
	// DefaultRequest are the default request options for the limit range
	DefaultRequest map[string]interface{} `json:"defaultRequest,omitempty"`

	LabelsAndAnnotations `json:",inline"`
}

type NetworkPolicy struct {
	// Enabled defines if the network policy should be deployed by vCluster.
	Enabled bool `json:"enabled,omitempty"`

	FallbackDNS         string              `json:"fallbackDns,omitempty"`
	OutgoingConnections OutgoingConnections `json:"outgoingConnections,omitempty"`

	LabelsAndAnnotations `json:",inline"`
}

type OutgoingConnections struct {
	IPBlock IPBlock `json:"ipBlock,omitempty"`
}

// IPBlock describes a particular CIDR (Ex. "192.168.1.0/24","2001:db8::/64") that is allowed
// to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs
// that should not be included within this rule.
type IPBlock struct {
	// cidr is a string representing the IPBlock
	// Valid examples are "192.168.1.0/24" or "2001:db8::/64"
	CIDR string `json:"cidr,omitempty"`

	// except is a slice of CIDRs that should not be included within an IPBlock
	// Valid examples are "192.168.1.0/24" or "2001:db8::/64"
	// Except values will be rejected if they are outside the cidr range
	// +optional
	Except []string `json:"except,omitempty"`
}

type CentralAdmission struct {
	// ValidatingWebhooks are validating webhooks that should be enforced in the vCluster
	ValidatingWebhooks []interface{} `json:"validatingWebhooks,omitempty"`
	// MutatingWebhooks are mutating webhooks that should be enforced in the vCluster
	MutatingWebhooks []interface{} `json:"mutatingWebhooks,omitempty"`
}

type RBAC struct {
	// Role holds vCluster role configuration
	Role RBACRole `json:"role,omitempty"`
	// ClusterRole holds vCluster cluster role configuration
	ClusterRole RBACClusterRole `json:"clusterRole,omitempty"`
}

type RBACClusterRole struct {
	// Disabled defines if the cluster role should be disabled. Otherwise, its automatically determined if vCluster requires a cluster role.
	Disabled bool `json:"disabled,omitempty"`
	// OverwriteRules will overwrite the cluster role rules completely.
	OverwriteRules []map[string]interface{} `json:"overwriteRules,omitempty"`
	// ExtraRules will add rules to the cluster role.
	ExtraRules []map[string]interface{} `json:"extraRules,omitempty"`
}

type RBACRole struct {
	// Enabled
	Enabled bool `json:"enabled,omitempty"`
	// OverwriteRules will overwrite the role rules completely.
	OverwriteRules []map[string]interface{} `json:"overwriteRules,omitempty"`
	// ExtraRules will add rules to the role.
	ExtraRules []map[string]interface{} `json:"extraRules,omitempty"`
}

type Telemetry struct {
	// Disabled specifies that the telemetry for vCluster control plane should be disabled.
	Disabled           bool   `json:"disabled,omitempty"`
	InstanceCreator    string `json:"instanceCreator,omitempty"`
	PlatformUserID     string `json:"platformUserID,omitempty"`
	PlatformInstanceID string `json:"platformInstanceID,omitempty"`
	MachineID          string `json:"machineID,omitempty"`
}

type Experimental struct {
	// IsolatedControlPlane is a feature to run the vCluster control plane in a different Kubernetes cluster than the workloads themselves.
	IsolatedControlPlane ExperimentalIsolatedControlPlane `json:"isolatedControlPlane,omitempty"`
	// SyncSettings are advanced settings for the syncer controller.
	SyncSettings ExperimentalSyncSettings `json:"syncSettings,omitempty"`
	// GenericSync holds options to generically sync resources from vCluster to host.
	GenericSync ExperimentalGenericSync `json:"genericSync,omitempty"`
	// Deploy allows you to configure manifests and helm charts to deploy within the vCluster.
	Deploy ExperimentalDeploy `json:"deploy,omitempty"`
	// MultiNamespaceMode tells vCluster to sync to multiple namespaces instead of a single one. This will map each vCluster namespace to a single namespace in the host cluster.
	MultiNamespaceMode ExperimentalMultiNamespaceMode `json:"multiNamespaceMode,omitempty"`

	// VirtualClusterKubeConfig allows you to override distro specifics and specify where vCluster will find the required certificates and vCluster config.
	VirtualClusterKubeConfig VirtualClusterKubeConfig `json:"virtualClusterKubeConfig,omitempty"`
}

type ExperimentalMultiNamespaceMode struct {
	// Enabled specifies if multi namespace mode should get enabled
	Enabled bool `json:"enabled,omitempty"`

	// NamespaceLabels are extra labels that will be added by vCluster to each created namespace.
	NamespaceLabels map[string]string `json:"namespaceLabels,omitempty"`
}

type ExperimentalIsolatedControlPlane struct {
	// Enabled specifies if the isolated control plane feature should be enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Headless states that helm should deploy the vCluster in headless mode for the isolated control plane.
	Headless bool `json:"headless,omitempty"`
	// KubeConfig is the path where to find the remote workload cluster kube config.
	KubeConfig string `json:"kubeConfig,omitempty"`
	// Namespace is the namespace where to sync the workloads into.
	Namespace string `json:"namespace,omitempty"`
	// Service is the vCluster service in the remote cluster.
	Service string `json:"service,omitempty"`
}

type ExperimentalSyncSettings struct {
	// DisableSync will not sync any resources and disable most control plane functionality.
	DisableSync bool `json:"disableSync,omitempty"`
	// RewriteKubernetesService will rewrite the kubernetes service to point to the vCluster if disableSync is enabled
	RewriteKubernetesService bool `json:"rewriteKubernetesService,omitempty"`

	// TargetNamespace is the namespace where the workloads should get synced to.
	TargetNamespace string `json:"targetNamespace,omitempty"`
	// SetOwner specifies if vCluster should set an owner reference on the synced objects to the vCluster service. This allows for easy garbage collection.
	SetOwner bool `json:"setOwner,omitempty"`
	// SyncLabels are labels that should get not rewritten when syncing from vCluster.
	SyncLabels []string `json:"syncLabels,omitempty"`
}

type ExperimentalDeploy struct {
	// Manifests are raw kubernetes manifests that should get applied within the vCluster.
	Manifests string `json:"manifests,omitempty"`
	// ManifestsTemplate is a kubernetes manifest template that will be rendered with vCluster values before applying it within the vCluster.
	ManifestsTemplate string `json:"manifestsTemplate,omitempty"`
	// Helm are helm charts that should get deployed into the vCluster
	Helm []ExperimentalDeployHelm `json:"helm,omitempty"`
}

type ExperimentalDeployHelm struct {
	// Chart defines what chart should get deployed.
	Chart ExperimentalDeployHelmChart `json:"chart,omitempty"`
	// Release defines what release should get deployed.
	Release ExperimentalDeployHelmRelease `json:"release,omitempty"`
	// Values defines what values should get used.
	Values string `json:"values,omitempty"`
	// Timeout defines the timeout for helm
	Timeout string `json:"timeout,omitempty"`
	// Bundle allows to compress the helm chart and specify this instead of an online chart
	Bundle string `json:"bundle,omitempty"`
}

type ExperimentalDeployHelmRelease struct {
	// Name of the release
	Name string `json:"name,omitempty"`
	// Namespace of the release
	Namespace string `json:"namespace,omitempty"`
}

type ExperimentalDeployHelmChart struct {
	Name     string `json:"name,omitempty"`
	Repo     string `json:"repo,omitempty"`
	Insecure bool   `json:"insecure,omitempty"`
	Version  string `json:"version,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Platform struct {
	// Name is the name of the vCluster instance in the vCluster platform
	Name string `json:"name,omitempty"`

	// Owner is the desired owner of the vCluster within the vCluster platform. If empty will take the current user.
	Owner PlatformOwner `json:"owner,omitempty"`

	// Project is the project within the platform where the vCluster should connect to.
	Project string `json:"project,omitempty"`

	// APIKey defines how vCluster can find the api key used for the platform.
	APIKey PlatformAPIKey `json:"apiKey,omitempty"`
}

type PlatformOwner struct {
	// User is the user id within the platform. This is mutually exclusive with team.
	User string `json:"user,omitempty"`

	// Team is the team id within the platform. This is mutually exclusive with user.
	Team string `json:"team,omitempty"`
}

type PlatformAPIKey struct {
	// Value specifies the api key as a regular text value.
	Value string `json:"value,omitempty"`

	// SecretRef defines where to find the platform api key. By default vCluster will search in the following locations in this precedence:
	// * platform.apiKey.value
	// * environment variable called LICENSE
	// * secret specified under platform.secret.name
	// * secret called "vcluster-platform-api-key" in the vCluster namespace
	SecretRef PlatformAPIKeySecretReference `json:"secretRef,omitempty"`
}

// PlatformAPIKeySecretReference defines where to find the platform api key. The secret key name doesn't matter as long as the secret only contains a single key.
type PlatformAPIKeySecretReference struct {
	// Name is the name of the secret where the platform api key is stored. This defaults to vcluster-platform-api-key if undefined.
	Name string `json:"name,omitempty"`

	// Namespace defines the namespace where the api key secret should be retrieved from. If this is not equal to the namespace
	// where the vCluster is deployed, you need to make sure vCluster has access to this other namespace.
	Namespace string `json:"namespace,omitempty"`
}

type ExperimentalGenericSync struct {
	// Version is the config version
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Exports syncs a resource from the virtual cluster to the host
	Exports []*Export `json:"export,omitempty" yaml:"export,omitempty"`

	// Imports syncs a resource from the host cluster to virtual cluster
	Imports []*Import `json:"import,omitempty" yaml:"import,omitempty"`

	// Hooks are hooks that can be used to inject custom patches before syncing
	Hooks *Hooks `json:"hooks,omitempty" yaml:"hooks,omitempty"`

	ClusterRole ExperimentalGenericSyncExtraRules `json:"clusterRole,omitempty"`
	Role        ExperimentalGenericSyncExtraRules `json:"role,omitempty"`
}

type ExperimentalGenericSyncExtraRules struct {
	ExtraRules []interface{} `json:"extraRules,omitempty"`
}

type Hooks struct {
	// HostToVirtual is a hook that is executed before syncing from the host to the virtual cluster
	HostToVirtual []*Hook `json:"hostToVirtual,omitempty" yaml:"hostToVirtual,omitempty"`

	// VirtualToHost is a hook that is executed before syncing from the virtual to the host cluster
	VirtualToHost []*Hook `json:"virtualToHost,omitempty" yaml:"virtualToHost,omitempty"`
}

type Hook struct {
	TypeInformation

	// Verbs are the verbs that the hook should mutate
	Verbs []string `json:"verbs,omitempty" yaml:"verbs,omitempty"`

	// Patches are the patches to apply on the object to be synced
	Patches []*Patch `json:"patches,omitempty" yaml:"patches,omitempty"`
}

type Import struct {
	SyncBase `json:",inline" yaml:",inline"`
}

type SyncBase struct {
	TypeInformation `json:",inline" yaml:",inline"`

	Optional bool `json:"optional,omitempty" yaml:"optional,omitempty"`

	// ReplaceWhenInvalid determines if the controller should try to recreate the object
	// if there is a problem applying
	ReplaceWhenInvalid bool `json:"replaceOnConflict,omitempty" yaml:"replaceOnConflict,omitempty"`

	// Patches are the patches to apply on the virtual cluster objects
	// when syncing them from the host cluster
	Patches []*Patch `json:"patches,omitempty" yaml:"patches,omitempty"`

	// ReversePatches are the patches to apply to host cluster objects
	// after it has been synced to the virtual cluster
	ReversePatches []*Patch `json:"reversePatches,omitempty" yaml:"reversePatches,omitempty"`
}

type Export struct {
	SyncBase `json:",inline" yaml:",inline"`

	// Selector is a label selector to select the synced objects in the virtual cluster.
	// If empty, all objects will be synced.
	Selector *Selector `json:"selector,omitempty" yaml:"selector,omitempty"`
}

type TypeInformation struct {
	// APIVersion of the object to sync
	APIVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`

	// Kind of the object to sync
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
}

type Selector struct {
	// LabelSelector are the labels to select the object from
	LabelSelector map[string]string `json:"labelSelector,omitempty" yaml:"labelSelector,omitempty"`
}

type Patch struct {
	// Operation is the type of the patch
	Operation PatchType `json:"op,omitempty" yaml:"op,omitempty"`

	// FromPath is the path from the other object
	FromPath string `json:"fromPath,omitempty" yaml:"fromPath,omitempty"`

	// Path is the path of the patch
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// NamePath is the path to the name of a child resource within Path
	NamePath string `json:"namePath,omitempty" yaml:"namePath,omitempty"`

	// NamespacePath is path to the namespace of a child resource within Path
	NamespacePath string `json:"namespacePath,omitempty" yaml:"namespacePath,omitempty"`

	// Value is the new value to be set to the path
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`

	// Regex - is regular expresion used to identify the Name,
	// and optionally Namespace, parts of the field value that
	// will be replaced with the rewritten Name and/or Namespace
	Regex       string         `json:"regex,omitempty" yaml:"regex,omitempty"`
	ParsedRegex *regexp.Regexp `json:"-"               yaml:"-"`

	// Conditions are conditions that must be true for
	// the patch to get executed
	Conditions []*PatchCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	// Ignore determines if the path should be ignored if handled as a reverse patch
	Ignore *bool `json:"ignore,omitempty" yaml:"ignore,omitempty"`

	// Sync defines if a specialized syncer should be initialized using values
	// from the rewriteName operation as Secret/Configmap names to be synced
	Sync *PatchSync `json:"sync,omitempty" yaml:"sync,omitempty"`
}

type PatchType string

const (
	PatchTypeRewriteName                     PatchType = "rewriteName"
	PatchTypeRewriteLabelKey                 PatchType = "rewriteLabelKey"
	PatchTypeRewriteLabelSelector            PatchType = "rewriteLabelSelector"
	PatchTypeRewriteLabelExpressionsSelector PatchType = "rewriteLabelExpressionsSelector"

	PatchTypeCopyFromObject PatchType = "copyFromObject"
	PatchTypeAdd            PatchType = "add"
	PatchTypeReplace        PatchType = "replace"
	PatchTypeRemove         PatchType = "remove"
)

type PatchCondition struct {
	// Path is the path within the object to select
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// SubPath is the path below the selected object to select
	SubPath string `json:"subPath,omitempty" yaml:"subPath,omitempty"`

	// Equal is the value the path should be equal to
	Equal interface{} `json:"equal,omitempty" yaml:"equal,omitempty"`

	// NotEqual is the value the path should not be equal to
	NotEqual interface{} `json:"notEqual,omitempty" yaml:"notEqual,omitempty"`

	// Empty means that the path value should be empty or unset
	Empty *bool `json:"empty,omitempty" yaml:"empty,omitempty"`
}

type PatchSync struct {
	Secret    *bool `json:"secret,omitempty"    yaml:"secret,omitempty"`
	ConfigMap *bool `json:"configmap,omitempty" yaml:"configmap,omitempty"`
}

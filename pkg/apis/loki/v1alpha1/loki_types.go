package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LokiSpec defines the desired state of Loki
type LokiSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	ServiceMonitorSelector          *metav1.LabelSelector    `json:"serviceMonitorSelector,omitempty"`
	ServiceMonitorNamespaceSelector *metav1.LabelSelector    `json:"serviceMonitorNamespaceSelector,omitempty"`
	ServiceAccount                  *LokiServiceAccount      `json:"serviceAccount,omitempty"`
	Containers                      []v1.Container           `json:"containers,omitempty"`
	InitContainers                  []v1.Container           `json:"initContainers,omitempty"`
	InitResources                   *v1.ResourceRequirements `json:"initResources,omitempty"`
	Deployment                      *LokiDeployment          `json:"deployment,omitempty"`
	Secrets                         []string                 `json:"secrets,omitempty"`
	ConfigMaps                      []string                 `json:"configMaps,omitempty"`
	Service                         *LokiService             `json:"service,omitempty"`
	Ingress                         *LokiIngress             `json:"ingress,omitempty"`
}

// LokiStatus defines the observed state of Loki
type LokiStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Loki is the Schema for the lokis API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=lokis,scope=Namespaced
type Loki struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LokiSpec   `json:"spec,omitempty"`
	Status LokiStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LokiList contains a list of Loki
type LokiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Loki `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Loki{}, &LokiList{})
}

// LokiDeployment ...
type LokiDeployment struct {
	Annotations                   map[string]string      `json:"annotations,omitempty"`
	Labels                        map[string]string      `json:"labels,omitempty"`
	Replicas                      int32                  `json:"replicas"`
	NodeSelector                  map[string]string      `json:"nodeSelector,omitempty"`
	Tolerations                   []v1.Toleration        `json:"tolerations,omitempty"`
	Affinity                      *v1.Affinity           `json:"affinity,omitempty"`
	SecurityContext               *v1.PodSecurityContext `json:"securityContext,omitempty"`
	ContainerSecurityContext      *v1.SecurityContext    `json:"containerSecurityContext,omitempty"`
	TerminationGracePeriodSeconds int64                  `json:"terminationGracePeriodSeconds"`
}

// LokiIngress provides a means to configure the ingress
type LokiIngress struct {
	Annotations   map[string]string `json:"annotations,omitempty"`
	Hostname      string            `json:"hostname,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Path          string            `json:"path,omitempty"`
	Enabled       bool              `json:"enabled,omitempty"`
	TLSEnabled    bool              `json:"tlsEnabled,omitempty"`
	TLSSecretName string            `json:"tlsSecretName,omitempty"`
	TargetPort    string            `json:"targetPort,omitempty"`
}

// LokiService provides a means to configure the service
type LokiService struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Type        v1.ServiceType    `json:"type,omitempty"`
	Ports       []v1.ServicePort  `json:"ports,omitempty"`
}

// LokiServiceAccount ...
type LokiServiceAccount struct {
	Annotations      map[string]string         `json:"annotations,omitempty"`
	Labels           map[string]string         `json:"labels,omitempty"`
	ImagePullSecrets []v1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
}

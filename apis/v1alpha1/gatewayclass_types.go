/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatewayClass describes a class of Gateways available to the user
// for creating Gateway resources.
//
// GatewayClass is a Cluster level resource.
//
// Support: Core.
type GatewayClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec for this GatewayClass.
	Spec GatewayClassSpec `json:"spec,omitempty"`
	// Status of the GatewayClass.
	// +kubebuilder:default={conditions: {{type: "InvalidParameters", status: "Unknown", message: "Waiting for controller", reason: "Waiting", lastTransitionTime: "1970-01-01T00:00:00Z"}}}
	Status GatewayClassStatus `json:"status,omitempty"`
}

// GatewayClassSpec reflects the configuration of a class of Gateways.
type GatewayClassSpec struct {
	// Controller is a domain/path string that indicates the
	// controller that is managing Gateways of this class.
	//
	// Example: "acme.io/gateway-controller".
	//
	// This field is not mutable and cannot be empty.
	//
	// The format of this field is DOMAIN "/" PATH, where DOMAIN
	// and PATH are valid Kubernetes names
	// (https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Support: Core
	//
	// +required
	Controller string `json:"controller"`

	// GatewayNamespaceSelector is a selector of namespaces. Gateways that
	// run in matched namespaces can use this GatewayClass. This is a standard
	// Kubernetes LabelSelector. Controllers must not support Gateways in
	// namespaces outside this selector.
	//
	// An empty selector (default) indicates that this GatewayClass is available
	// to use by Gateways in any namespace.
	//
	// When a Gateway attempts to use this class from a namespace that is not
	// allowed by this selector, the controller implementing the GatewayClass
	// may add a new "ForbiddenNamespaceForClass" condition to the Gateway
	// status. Adding this condition is considered optional since not all
	// controllers will have access to all namespaces.
	//
	// Support: Core
	//
	// +optional
	GatewayNamespaceSelector metav1.LabelSelector `json:"gatewayNamespaceSelector,omitempty"`

	// AllowedRouteNamespaces indicates in which namespaces Routes can be
	// selected for Gateways of this class. This is restricted to the namespace
	// of the Gateway by default.
	//
	// When any Routes are selected by a Gateway in a namespace that is not
	// allowed by this selector, the controller implementing the GatewayClass
	// may add a new "ForbiddenRoutesForClass" condition to the Gateway status.
	// Adding this condition is considered optional since not all controllers
	// will have access to all namespaces.
	//
	// Support: Core
	//
	// +optional
	// +kubebuilder:default={onlySameNamespace:true}
	AllowedRouteNamespaces RouteNamespaces `json:"allowedRouteNamespaces,omitempty"`

	// ParametersRef is a controller-specific resource containing the
	// configuration parameters corresponding to this class. This is optional if
	// the controller does not require any additional configuration.
	//
	// Parameters resources are implementation specific custom resources. These
	// resources must be cluster-scoped.
	//
	// If the referent cannot be found, the GatewayClass's "InvalidParameters"
	// status condition will be true.
	//
	// Support: Custom
	//
	// +optional
	ParametersRef *GatewayClassParametersObjectReference `json:"parametersRef,omitempty"`
}

// RouteNamespaces is used by Gateway and GatewayClass to indicate which
// namespaces Routes should be selected from.
type RouteNamespaces struct {
	// NamespaceSelector is a selector of namespaces that Routes should be
	// selected from. This is a standard Kubernetes LabelSelector, a label query
	// over a set of resources. The result of matchLabels and matchExpressions
	// are ANDed. Controllers must not support Routes in namespaces outside this
	// selector.
	//
	// An empty selector (default) indicates that Routes in any namespace can be
	// selected.
	//
	// The OnlySameNamespace field takes precedence over this field. This
	// selector will only take effect when OnlySameNamespace is false.
	//
	// Support: Core
	//
	// +optional
	NamespaceSelector metav1.LabelSelector `json:"namespaceSelector"`

	// OnlySameNamespace is a boolean used to indicate if Route references are
	// limited to the same Namespace as the Gateway. When true, only Routes
	// within the same Namespace as the Gateway should be selected.
	//
	// This field takes precedence over the NamespaceSelector field. That
	// selector should only take effect when this field is set to false.
	//
	// Support: Core
	//
	// +optional
	// +kubebuilder:default=true
	OnlySameNamespace bool `json:"onlySameNamespace"`
}

// GatewayClassParametersObjectReference identifies a cluster-scoped parameters
// resource for a GatewayClass.
//
// +k8s:deepcopy-gen=false
type GatewayClassParametersObjectReference = LocalObjectReference

// GatewayClassConditionType is the type of status conditions. This
// type should be used with the GatewayClassStatus.Conditions field.
type GatewayClassConditionType string

const (
	// GatewayClassConditionStatusInvalidParameters indicates the
	// validity of the Parameters set for a given controller. This
	// will initially start off as "Unknown".
	GatewayClassConditionStatusInvalidParameters GatewayClassConditionType = "InvalidParameters"
)

// GatewayClassStatus is the current status for the GatewayClass.
type GatewayClassStatus struct {
	// Conditions is the current status from the controller for
	// this GatewayClass.
	// +optional
	// +kubebuilder:default={{type: "InvalidParameters", status: "Unknown", message: "Waiting for controller", reason: "Waiting", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayClassList contains a list of GatewayClass
type GatewayClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayClass `json:"items"`
}

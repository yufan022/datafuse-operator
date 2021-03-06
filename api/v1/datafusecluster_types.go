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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DatafuseClusterSpec defines the desired state of DatafuseCluster
type DatafuseClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Define a set of compute groups belongs to the cluster
	ComputeGroups []*DatafuseComputeGroup `json:"computegroup,omitempty"`
	// Fuse Query and Fuse Store will share the same version
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Default=latest
	// +kubebuilder:validation:Optional
	Version       *string                  `json:"version,omitempty"`
}

// DatafuseClusterStatus defines the observed state of DatafuseCluster
type DatafuseClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// state of each compute group. map from namespace to compute group
	ComputeGroupStates map[string]ComputeGroupState	`json:"computegroupstates",omitempty`
}

// +kubebuilder:object:root=true

// DatafuseCluster is the Schema for the datafuseclusters API
type DatafuseCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatafuseClusterSpec   `json:"spec,omitempty"`
	Status DatafuseClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatafuseClusterList contains a list of DatafuseCluster
type DatafuseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatafuseCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatafuseCluster{}, &DatafuseClusterList{})
}

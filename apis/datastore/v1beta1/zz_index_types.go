/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IndexObservation struct {

	// Policy for including ancestors in the index.
	// Default value is NONE.
	// Possible values are: NONE, ALL_ANCESTORS.
	Ancestor *string `json:"ancestor,omitempty" tf:"ancestor,omitempty"`

	// an identifier for the resource with format projects/{{project}}/indexes/{{index_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The index id.
	IndexID *string `json:"indexId,omitempty" tf:"index_id,omitempty"`

	// The entity kind which the index applies to.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// An ordered list of properties to index on.
	// Structure is documented below.
	Properties []PropertiesObservation `json:"properties,omitempty" tf:"properties,omitempty"`
}

type IndexParameters struct {

	// Policy for including ancestors in the index.
	// Default value is NONE.
	// Possible values are: NONE, ALL_ANCESTORS.
	// +kubebuilder:validation:Optional
	Ancestor *string `json:"ancestor,omitempty" tf:"ancestor,omitempty"`

	// The entity kind which the index applies to.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// An ordered list of properties to index on.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Properties []PropertiesParameters `json:"properties,omitempty" tf:"properties,omitempty"`
}

type PropertiesObservation struct {

	// The direction the index should optimize for sorting.
	// Possible values are: ASCENDING, DESCENDING.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The property name to index.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PropertiesParameters struct {

	// The direction the index should optimize for sorting.
	// Possible values are: ASCENDING, DESCENDING.
	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// The property name to index.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// IndexSpec defines the desired state of Index
type IndexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IndexParameters `json:"forProvider"`
}

// IndexStatus defines the observed state of Index.
type IndexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IndexObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Index is the Schema for the Indexs API. Describes a composite index for Cloud Datastore.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Index struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.kind)",message="kind is a required parameter"
	Spec   IndexSpec   `json:"spec"`
	Status IndexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IndexList contains a list of Indexs
type IndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Index `json:"items"`
}

// Repository type metadata.
var (
	Index_Kind             = "Index"
	Index_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Index_Kind}.String()
	Index_KindAPIVersion   = Index_Kind + "." + CRDGroupVersion.String()
	Index_GroupVersionKind = CRDGroupVersion.WithKind(Index_Kind)
)

func init() {
	SchemeBuilder.Register(&Index{}, &IndexList{})
}

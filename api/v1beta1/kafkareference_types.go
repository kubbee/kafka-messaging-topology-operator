/*
Copyright 2022.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KafkaReferenceSpec defines the desired state of KafkaReference
type KafkaReferenceSpec struct {
	Tenant          string                   `json:"tenant,omitempty"`
	BootstrapServer KafkaReferenceBootstrap  `json:"bootstrapServer,omitempty"`
	ZookeeperServer KafkaReferenceZookeeper  `json:"zookeeperServer,omitempty"`
	KafkaConnection KafkaReferenceConnection `json:"kafkaConnection,omitempty"`
}

// KafkaReferenceStatus defines the observed state of KafkaReference
type KafkaReferenceStatus struct {
	ObservedGeneration int64       `json:"observedGeneration,omitempty"`
	Conditions         []Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KafkaReference is the Schema for the kafkareferences API
type KafkaReference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaReferenceSpec   `json:"spec,omitempty"`
	Status KafkaReferenceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KafkaReferenceList contains a list of KafkaReference
type KafkaReferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaReference `json:"items"`
}

type KafkaReferenceConnection struct {
	ConnectionSecret string `json:"connectionSecret,omitempty"`
}

type KafkaReferenceBootstrap struct {
	Uri  string `json:"uri,omitempty"`
	Port string `json:"port,omitempty"`
}

type KafkaReferenceZookeeper struct {
	Uri  string `json:"uri,omitempty"`
	Port string `json:"port,omitempty"`
}

func init() {
	SchemeBuilder.Register(&KafkaReference{}, &KafkaReferenceList{})
}

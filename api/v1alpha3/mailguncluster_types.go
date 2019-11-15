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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Priority string

const (
	// PriorityUrgent means do this right away
	PriorityUrgent = Priority("Urgent")

	// PriorityUrgent means do this immediately
	PriorityExtremelyUrgent = Priority("ExtremelyUrgent")

	// PriorityBusinessCritical means you absolutely need to do this now
	PriorityBusinessCritical = Priority("BusinessCritical")
)

// MailgunClusterSpec defines the desired state of MailgunCluster
type MailgunClusterSpec struct {
	// Priority is how quickly you need this cluster
	Priority Priority `json:"priority,omitempty"`
	// Request is where you ask extra nicely
	Request string `json:"request,omitempty"`
	// Requester is the email of the person sending the request
	Requester string `json:"requester,omitempty"`
}

// MailgunClusterStatus defines the observed state of MailgunCluster
type MailgunClusterStatus struct {
	// MessageID is set to the message ID from Mailgun when our message has been sent
	MessageID *string `json:"response"`
}

// +kubebuilder:object:root=true

// MailgunCluster is the Schema for the mailgunclusters API
type MailgunCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MailgunClusterSpec   `json:"spec,omitempty"`
	Status MailgunClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MailgunClusterList contains a list of MailgunCluster
type MailgunClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MailgunCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MailgunCluster{}, &MailgunClusterList{})
}

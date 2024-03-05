/*
MIT License

Copyright (c) 2023 mercari

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ScheduledScalingSpec defines the desired state of ScheduledScaling
type ScheduledScalingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	/// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Schedule Schedule `json:"schedule" protobuf:"bytes,1,name=schedule"`

	// TargetRef is the targets that need to be scheduled
	TargetRefs TargetRefs `json:"targetRefs" protobuf:"bytes,2,name=targetRefs"`

	// Currently only static, might implement others(?)
	Strategy Strategy `json:"strategy" protobuf:"bytes,3,name=strategy"`
}

// ScheduledScalingStatus defines the observed state of ScheduledScaling
type ScheduledScalingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ScheduledScalingPhase ScheduledScalingPhase `json:"scheduledScalingPhase" protobuf:"bytes,1,name=scheduledScalingPhase"`
}

type ScheduledScalingPhase string

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ScheduledScaling is the Schema for the scheduledscalings API
type ScheduledScaling struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScheduledScalingSpec   `json:"spec,omitempty"`
	Status ScheduledScalingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ScheduledScalingList contains a list of ScheduledScaling
type ScheduledScalingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScheduledScaling `json:"items"`
}

type TargetRefs struct {
	// ScaleTargetRef is the target of scaling.
	// It should be the same as the target of HPA.
	ScaleTargetRef CrossVersionObjectReference `json:"scaleTargetRef" protobuf:"bytes,1,name=scaleTargetRef"`
	//Tortoise to be targeted for scheduled scaling
	TortoiseName *string `json:"tortoiseName,omitempty" protobuf:"bytes,2,name=tortoiseName"`
}

// CrossVersionObjectReference contains enough information toet identify the referred resource.
type CrossVersionObjectReference struct {
	// kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`

	// name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`

	// apiVersion is the API version of the referent
	// +optional
	APIVersion *string `json:"apiVersion,omitempty" protobuf:"bytes,3,opt,name=apiVersion"`
}

type Schedule struct {
	/// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	// start of schedule
	StartAt *string `json:"startAt,omitempty" protobuf:"bytes,1,opt,name=startAt"`
	// end of schedule
	FinishAt *string `json:"finishAt,omitempty" protobuf:"bytes,2,name=finishAt"`
}

type Strategy struct {
	// Resource scaling requirements
	Static *Static `json:"static,omitempty" protobuf:"bytes,1,opt,name=static"`
}

type Static struct {
	// Min replicas to be deployed on schedule
	MinimumMinReplicas *int `json:"minimumMinReplica,omitempty" protobuf:"bytes,1,opt,name=minimumMinReplica"`
	// Resources requested per container
	MinAllocatedResources []ContainerResourceRequests `json:"minAllocatedResourcesomitempty,omitempty" protobuf:"bytes,2,opt,name=minAllocatedResources"`
}

// Not sure if I need this here, duplicated code from tortoise types
type ContainerResourceRequests struct {
	// ContainerName is the name of target container.
	ContainerName string          `json:"containerName" protobuf:"bytes,1,name=containerName"`
	Resource      v1.ResourceList `json:"resource" protobuf:"bytes,2,name=resource"`
}

func init() {
	SchemeBuilder.Register(&ScheduledScaling{}, &ScheduledScalingList{})
}
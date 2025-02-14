// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package types

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8srt "k8s.io/apimachinery/pkg/runtime"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
)

// RuntimeMetaObject contains both the Kubernetes apimachinery/runtime.Object
// and apimachinery/apis/meta/v1.Object interfaces
//
// NOTE(jaypipes): This really belongs as an upstream apimachinery type
type RuntimeMetaObject interface {
	metav1.Object
	k8srt.Object
}

// AWSResource represents a custom resource object in the Kubernetes API that
// corresponds to a resource in an AWS service API.
type AWSResource interface {
	ConditionManager
	// Identifiers returns an AWSResourceIdentifiers object containing various
	// identifying information, including the AWS account ID that owns the
	// resource, the resource's AWS Resource Name (ARN)
	Identifiers() AWSResourceIdentifiers
	// IsBeingDeleted returns true if the Kubernetes resource has a non-zero
	// deletion timestamp
	IsBeingDeleted() bool
	// RuntimeObject returns the Kubernetes apimachinery/runtime representation
	// of the AWSResource
	RuntimeObject() k8srt.Object
	// MetaObject returns the Kubernetes apimachinery/apis/meta/v1.Object
	// representation of the AWSResource
	MetaObject() metav1.Object
	// RuntimeMetaObject returns an object that implements both the Kubernetes
	// apimachinery/runtime.Object and the Kubernetes
	// apimachinery/apis/meta/v1.Object interfaces
	RuntimeMetaObject() RuntimeMetaObject
	// SetObjectMeta sets the ObjectMeta field for the resource
	SetObjectMeta(meta metav1.ObjectMeta)
	// SetIdentifiers will set the the Spec or Status field that represents the
	// identifier for the resource.
	SetIdentifiers(*ackv1alpha1.AWSIdentifiers) error
	// SetStatus will set the Status field for the resource
	SetStatus(AWSResource)
	// DeepCopy will return a copy of the resource
	DeepCopy() AWSResource
}

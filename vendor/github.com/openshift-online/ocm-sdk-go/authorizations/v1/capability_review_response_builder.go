/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

// CapabilityReviewResponseBuilder contains the data and logic needed to build 'capability_review_response' objects.
//
// Representation of a capability review response.
type CapabilityReviewResponseBuilder struct {
	result *string
}

// NewCapabilityReviewResponse creates a new builder of 'capability_review_response' objects.
func NewCapabilityReviewResponse() *CapabilityReviewResponseBuilder {
	return new(CapabilityReviewResponseBuilder)
}

// Result sets the value of the 'result' attribute to the given value.
//
//
func (b *CapabilityReviewResponseBuilder) Result(value string) *CapabilityReviewResponseBuilder {
	b.result = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *CapabilityReviewResponseBuilder) Copy(object *CapabilityReviewResponse) *CapabilityReviewResponseBuilder {
	if object == nil {
		return b
	}
	b.result = object.result
	return b
}

// Build creates a 'capability_review_response' object using the configuration stored in the builder.
func (b *CapabilityReviewResponseBuilder) Build() (object *CapabilityReviewResponse, err error) {
	object = new(CapabilityReviewResponse)
	object.result = b.result
	return
}
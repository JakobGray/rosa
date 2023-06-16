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

// SelfAccessReviewResponse represents the values of the 'self_access_review_response' type.
//
// Representation of an access review response, performed against oneself
type SelfAccessReviewResponse struct {
	bitmap_        uint32
	action         string
	clusterID      string
	clusterUUID    string
	organizationID string
	reason         string
	resourceType   string
	subscriptionID string
	allowed        bool
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *SelfAccessReviewResponse) Empty() bool {
	return o == nil || o.bitmap_ == 0
}

// Action returns the value of the 'action' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates the action, one of: [get,list,create,delete,update]
func (o *SelfAccessReviewResponse) Action() string {
	if o != nil && o.bitmap_&1 != 0 {
		return o.action
	}
	return ""
}

// GetAction returns the value of the 'action' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates the action, one of: [get,list,create,delete,update]
func (o *SelfAccessReviewResponse) GetAction() (value string, ok bool) {
	ok = o != nil && o.bitmap_&1 != 0
	if ok {
		value = o.action
	}
	return
}

// Allowed returns the value of the 'allowed' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Defines whether the action on the specified resource type is allowed
func (o *SelfAccessReviewResponse) Allowed() bool {
	if o != nil && o.bitmap_&2 != 0 {
		return o.allowed
	}
	return false
}

// GetAllowed returns the value of the 'allowed' attribute and
// a flag indicating if the attribute has a value.
//
// Defines whether the action on the specified resource type is allowed
func (o *SelfAccessReviewResponse) GetAllowed() (value bool, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.allowed
	}
	return
}

// ClusterID returns the value of the 'cluster_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Cluster (internal id) the resource type belongs to
func (o *SelfAccessReviewResponse) ClusterID() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.clusterID
	}
	return ""
}

// GetClusterID returns the value of the 'cluster_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Cluster (internal id) the resource type belongs to
func (o *SelfAccessReviewResponse) GetClusterID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.clusterID
	}
	return
}

// ClusterUUID returns the value of the 'cluster_UUID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Cluster (external id) the resource type belongs to
func (o *SelfAccessReviewResponse) ClusterUUID() string {
	if o != nil && o.bitmap_&8 != 0 {
		return o.clusterUUID
	}
	return ""
}

// GetClusterUUID returns the value of the 'cluster_UUID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Cluster (external id) the resource type belongs to
func (o *SelfAccessReviewResponse) GetClusterUUID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.clusterUUID
	}
	return
}

// OrganizationID returns the value of the 'organization_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Organization the resource type belongs to
func (o *SelfAccessReviewResponse) OrganizationID() string {
	if o != nil && o.bitmap_&16 != 0 {
		return o.organizationID
	}
	return ""
}

// GetOrganizationID returns the value of the 'organization_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Organization the resource type belongs to
func (o *SelfAccessReviewResponse) GetOrganizationID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&16 != 0
	if ok {
		value = o.organizationID
	}
	return
}

// Reason returns the value of the 'reason' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Defines further context for the value in allowed (if applicable)
func (o *SelfAccessReviewResponse) Reason() string {
	if o != nil && o.bitmap_&32 != 0 {
		return o.reason
	}
	return ""
}

// GetReason returns the value of the 'reason' attribute and
// a flag indicating if the attribute has a value.
//
// Defines further context for the value in allowed (if applicable)
func (o *SelfAccessReviewResponse) GetReason() (value string, ok bool) {
	ok = o != nil && o.bitmap_&32 != 0
	if ok {
		value = o.reason
	}
	return
}

// ResourceType returns the value of the 'resource_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates the type of the resource an action would be taken on.
// See uhc-account-manager/openapi/openapi.yaml for a list of possible values
func (o *SelfAccessReviewResponse) ResourceType() string {
	if o != nil && o.bitmap_&64 != 0 {
		return o.resourceType
	}
	return ""
}

// GetResourceType returns the value of the 'resource_type' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates the type of the resource an action would be taken on.
// See uhc-account-manager/openapi/openapi.yaml for a list of possible values
func (o *SelfAccessReviewResponse) GetResourceType() (value string, ok bool) {
	ok = o != nil && o.bitmap_&64 != 0
	if ok {
		value = o.resourceType
	}
	return
}

// SubscriptionID returns the value of the 'subscription_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Subscription the resource type belongs to
func (o *SelfAccessReviewResponse) SubscriptionID() string {
	if o != nil && o.bitmap_&128 != 0 {
		return o.subscriptionID
	}
	return ""
}

// GetSubscriptionID returns the value of the 'subscription_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Subscription the resource type belongs to
func (o *SelfAccessReviewResponse) GetSubscriptionID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&128 != 0
	if ok {
		value = o.subscriptionID
	}
	return
}

// SelfAccessReviewResponseListKind is the name of the type used to represent list of objects of
// type 'self_access_review_response'.
const SelfAccessReviewResponseListKind = "SelfAccessReviewResponseList"

// SelfAccessReviewResponseListLinkKind is the name of the type used to represent links to list
// of objects of type 'self_access_review_response'.
const SelfAccessReviewResponseListLinkKind = "SelfAccessReviewResponseListLink"

// SelfAccessReviewResponseNilKind is the name of the type used to nil lists of objects of
// type 'self_access_review_response'.
const SelfAccessReviewResponseListNilKind = "SelfAccessReviewResponseListNil"

// SelfAccessReviewResponseList is a list of values of the 'self_access_review_response' type.
type SelfAccessReviewResponseList struct {
	href  string
	link  bool
	items []*SelfAccessReviewResponse
}

// Len returns the length of the list.
func (l *SelfAccessReviewResponseList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *SelfAccessReviewResponseList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *SelfAccessReviewResponseList) Get(i int) *SelfAccessReviewResponse {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *SelfAccessReviewResponseList) Slice() []*SelfAccessReviewResponse {
	var slice []*SelfAccessReviewResponse
	if l == nil {
		slice = make([]*SelfAccessReviewResponse, 0)
	} else {
		slice = make([]*SelfAccessReviewResponse, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *SelfAccessReviewResponseList) Each(f func(item *SelfAccessReviewResponse) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *SelfAccessReviewResponseList) Range(f func(index int, item *SelfAccessReviewResponse) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}

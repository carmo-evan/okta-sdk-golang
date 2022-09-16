/*
Okta Management APIs

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 3.0.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// EmailDefaultContentAllOf struct for EmailDefaultContentAllOf
type EmailDefaultContentAllOf struct {
	Links *EmailDefaultContentAllOfLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailDefaultContentAllOf EmailDefaultContentAllOf

// NewEmailDefaultContentAllOf instantiates a new EmailDefaultContentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDefaultContentAllOf() *EmailDefaultContentAllOf {
	this := EmailDefaultContentAllOf{}
	return &this
}

// NewEmailDefaultContentAllOfWithDefaults instantiates a new EmailDefaultContentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDefaultContentAllOfWithDefaults() *EmailDefaultContentAllOf {
	this := EmailDefaultContentAllOf{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailDefaultContentAllOf) GetLinks() EmailDefaultContentAllOfLinks {
	if o == nil || o.Links == nil {
		var ret EmailDefaultContentAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDefaultContentAllOf) GetLinksOk() (*EmailDefaultContentAllOfLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailDefaultContentAllOf) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EmailDefaultContentAllOfLinks and assigns it to the Links field.
func (o *EmailDefaultContentAllOf) SetLinks(v EmailDefaultContentAllOfLinks) {
	o.Links = &v
}

func (o EmailDefaultContentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailDefaultContentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEmailDefaultContentAllOf := _EmailDefaultContentAllOf{}

	err = json.Unmarshal(bytes, &varEmailDefaultContentAllOf)
	if err == nil {
		*o = EmailDefaultContentAllOf(varEmailDefaultContentAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailDefaultContentAllOf struct {
	value *EmailDefaultContentAllOf
	isSet bool
}

func (v NullableEmailDefaultContentAllOf) Get() *EmailDefaultContentAllOf {
	return v.value
}

func (v *NullableEmailDefaultContentAllOf) Set(val *EmailDefaultContentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDefaultContentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDefaultContentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDefaultContentAllOf(val *EmailDefaultContentAllOf) *NullableEmailDefaultContentAllOf {
	return &NullableEmailDefaultContentAllOf{value: val, isSet: true}
}

func (v NullableEmailDefaultContentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDefaultContentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


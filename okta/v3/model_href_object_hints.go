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

// HrefObjectHints struct for HrefObjectHints
type HrefObjectHints struct {
	Allow []string `json:"allow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HrefObjectHints HrefObjectHints

// NewHrefObjectHints instantiates a new HrefObjectHints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHrefObjectHints() *HrefObjectHints {
	this := HrefObjectHints{}
	return &this
}

// NewHrefObjectHintsWithDefaults instantiates a new HrefObjectHints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrefObjectHintsWithDefaults() *HrefObjectHints {
	this := HrefObjectHints{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *HrefObjectHints) GetAllow() []string {
	if o == nil || o.Allow == nil {
		var ret []string
		return ret
	}
	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefObjectHints) GetAllowOk() ([]string, bool) {
	if o == nil || o.Allow == nil {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *HrefObjectHints) HasAllow() bool {
	if o != nil && o.Allow != nil {
		return true
	}

	return false
}

// SetAllow gets a reference to the given []string and assigns it to the Allow field.
func (o *HrefObjectHints) SetAllow(v []string) {
	o.Allow = v
}

func (o HrefObjectHints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allow != nil {
		toSerialize["allow"] = o.Allow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HrefObjectHints) UnmarshalJSON(bytes []byte) (err error) {
	varHrefObjectHints := _HrefObjectHints{}

	err = json.Unmarshal(bytes, &varHrefObjectHints)
	if err == nil {
		*o = HrefObjectHints(varHrefObjectHints)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allow")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableHrefObjectHints struct {
	value *HrefObjectHints
	isSet bool
}

func (v NullableHrefObjectHints) Get() *HrefObjectHints {
	return v.value
}

func (v *NullableHrefObjectHints) Set(val *HrefObjectHints) {
	v.value = val
	v.isSet = true
}

func (v NullableHrefObjectHints) IsSet() bool {
	return v.isSet
}

func (v *NullableHrefObjectHints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHrefObjectHints(val *HrefObjectHints) *NullableHrefObjectHints {
	return &NullableHrefObjectHints{value: val, isSet: true}
}

func (v NullableHrefObjectHints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHrefObjectHints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


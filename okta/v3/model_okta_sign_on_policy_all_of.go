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

// OktaSignOnPolicyAllOf struct for OktaSignOnPolicyAllOf
type OktaSignOnPolicyAllOf struct {
	Conditions *OktaSignOnPolicyConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSignOnPolicyAllOf OktaSignOnPolicyAllOf

// NewOktaSignOnPolicyAllOf instantiates a new OktaSignOnPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSignOnPolicyAllOf() *OktaSignOnPolicyAllOf {
	this := OktaSignOnPolicyAllOf{}
	return &this
}

// NewOktaSignOnPolicyAllOfWithDefaults instantiates a new OktaSignOnPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSignOnPolicyAllOfWithDefaults() *OktaSignOnPolicyAllOf {
	this := OktaSignOnPolicyAllOf{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *OktaSignOnPolicyAllOf) GetConditions() OktaSignOnPolicyConditions {
	if o == nil || o.Conditions == nil {
		var ret OktaSignOnPolicyConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyAllOf) GetConditionsOk() (*OktaSignOnPolicyConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *OktaSignOnPolicyAllOf) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given OktaSignOnPolicyConditions and assigns it to the Conditions field.
func (o *OktaSignOnPolicyAllOf) SetConditions(v OktaSignOnPolicyConditions) {
	o.Conditions = &v
}

func (o OktaSignOnPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaSignOnPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOktaSignOnPolicyAllOf := _OktaSignOnPolicyAllOf{}

	err = json.Unmarshal(bytes, &varOktaSignOnPolicyAllOf)
	if err == nil {
		*o = OktaSignOnPolicyAllOf(varOktaSignOnPolicyAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conditions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaSignOnPolicyAllOf struct {
	value *OktaSignOnPolicyAllOf
	isSet bool
}

func (v NullableOktaSignOnPolicyAllOf) Get() *OktaSignOnPolicyAllOf {
	return v.value
}

func (v *NullableOktaSignOnPolicyAllOf) Set(val *OktaSignOnPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicyAllOf(val *OktaSignOnPolicyAllOf) *NullableOktaSignOnPolicyAllOf {
	return &NullableOktaSignOnPolicyAllOf{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


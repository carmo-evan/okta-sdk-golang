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

// UserStatusPolicyRuleCondition struct for UserStatusPolicyRuleCondition
type UserStatusPolicyRuleCondition struct {
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserStatusPolicyRuleCondition UserStatusPolicyRuleCondition

// NewUserStatusPolicyRuleCondition instantiates a new UserStatusPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStatusPolicyRuleCondition() *UserStatusPolicyRuleCondition {
	this := UserStatusPolicyRuleCondition{}
	return &this
}

// NewUserStatusPolicyRuleConditionWithDefaults instantiates a new UserStatusPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatusPolicyRuleConditionWithDefaults() *UserStatusPolicyRuleCondition {
	this := UserStatusPolicyRuleCondition{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UserStatusPolicyRuleCondition) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserStatusPolicyRuleCondition) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UserStatusPolicyRuleCondition) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UserStatusPolicyRuleCondition) SetValue(v string) {
	o.Value = &v
}

func (o UserStatusPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserStatusPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varUserStatusPolicyRuleCondition := _UserStatusPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varUserStatusPolicyRuleCondition)
	if err == nil {
		*o = UserStatusPolicyRuleCondition(varUserStatusPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserStatusPolicyRuleCondition struct {
	value *UserStatusPolicyRuleCondition
	isSet bool
}

func (v NullableUserStatusPolicyRuleCondition) Get() *UserStatusPolicyRuleCondition {
	return v.value
}

func (v *NullableUserStatusPolicyRuleCondition) Set(val *UserStatusPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatusPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatusPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatusPolicyRuleCondition(val *UserStatusPolicyRuleCondition) *NullableUserStatusPolicyRuleCondition {
	return &NullableUserStatusPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableUserStatusPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatusPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


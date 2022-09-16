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

// AccessPolicyRuleConditionsAllOf struct for AccessPolicyRuleConditionsAllOf
type AccessPolicyRuleConditionsAllOf struct {
	Device *DeviceAccessPolicyRuleCondition `json:"device,omitempty"`
	ElCondition *AccessPolicyRuleCustomCondition `json:"elCondition,omitempty"`
	UserType *UserTypeCondition `json:"userType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyRuleConditionsAllOf AccessPolicyRuleConditionsAllOf

// NewAccessPolicyRuleConditionsAllOf instantiates a new AccessPolicyRuleConditionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyRuleConditionsAllOf() *AccessPolicyRuleConditionsAllOf {
	this := AccessPolicyRuleConditionsAllOf{}
	return &this
}

// NewAccessPolicyRuleConditionsAllOfWithDefaults instantiates a new AccessPolicyRuleConditionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyRuleConditionsAllOfWithDefaults() *AccessPolicyRuleConditionsAllOf {
	this := AccessPolicyRuleConditionsAllOf{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditionsAllOf) GetDevice() DeviceAccessPolicyRuleCondition {
	if o == nil || o.Device == nil {
		var ret DeviceAccessPolicyRuleCondition
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditionsAllOf) GetDeviceOk() (*DeviceAccessPolicyRuleCondition, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditionsAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given DeviceAccessPolicyRuleCondition and assigns it to the Device field.
func (o *AccessPolicyRuleConditionsAllOf) SetDevice(v DeviceAccessPolicyRuleCondition) {
	o.Device = &v
}

// GetElCondition returns the ElCondition field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditionsAllOf) GetElCondition() AccessPolicyRuleCustomCondition {
	if o == nil || o.ElCondition == nil {
		var ret AccessPolicyRuleCustomCondition
		return ret
	}
	return *o.ElCondition
}

// GetElConditionOk returns a tuple with the ElCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditionsAllOf) GetElConditionOk() (*AccessPolicyRuleCustomCondition, bool) {
	if o == nil || o.ElCondition == nil {
		return nil, false
	}
	return o.ElCondition, true
}

// HasElCondition returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditionsAllOf) HasElCondition() bool {
	if o != nil && o.ElCondition != nil {
		return true
	}

	return false
}

// SetElCondition gets a reference to the given AccessPolicyRuleCustomCondition and assigns it to the ElCondition field.
func (o *AccessPolicyRuleConditionsAllOf) SetElCondition(v AccessPolicyRuleCustomCondition) {
	o.ElCondition = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditionsAllOf) GetUserType() UserTypeCondition {
	if o == nil || o.UserType == nil {
		var ret UserTypeCondition
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditionsAllOf) GetUserTypeOk() (*UserTypeCondition, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditionsAllOf) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given UserTypeCondition and assigns it to the UserType field.
func (o *AccessPolicyRuleConditionsAllOf) SetUserType(v UserTypeCondition) {
	o.UserType = &v
}

func (o AccessPolicyRuleConditionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.ElCondition != nil {
		toSerialize["elCondition"] = o.ElCondition
	}
	if o.UserType != nil {
		toSerialize["userType"] = o.UserType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyRuleConditionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyRuleConditionsAllOf := _AccessPolicyRuleConditionsAllOf{}

	err = json.Unmarshal(bytes, &varAccessPolicyRuleConditionsAllOf)
	if err == nil {
		*o = AccessPolicyRuleConditionsAllOf(varAccessPolicyRuleConditionsAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "elCondition")
		delete(additionalProperties, "userType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAccessPolicyRuleConditionsAllOf struct {
	value *AccessPolicyRuleConditionsAllOf
	isSet bool
}

func (v NullableAccessPolicyRuleConditionsAllOf) Get() *AccessPolicyRuleConditionsAllOf {
	return v.value
}

func (v *NullableAccessPolicyRuleConditionsAllOf) Set(val *AccessPolicyRuleConditionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyRuleConditionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyRuleConditionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyRuleConditionsAllOf(val *AccessPolicyRuleConditionsAllOf) *NullableAccessPolicyRuleConditionsAllOf {
	return &NullableAccessPolicyRuleConditionsAllOf{value: val, isSet: true}
}

func (v NullableAccessPolicyRuleConditionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyRuleConditionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


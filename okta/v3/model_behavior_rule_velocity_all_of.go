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

// BehaviorRuleVelocityAllOf struct for BehaviorRuleVelocityAllOf
type BehaviorRuleVelocityAllOf struct {
	Settings *BehaviorRuleSettingsVelocity `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleVelocityAllOf BehaviorRuleVelocityAllOf

// NewBehaviorRuleVelocityAllOf instantiates a new BehaviorRuleVelocityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleVelocityAllOf() *BehaviorRuleVelocityAllOf {
	this := BehaviorRuleVelocityAllOf{}
	return &this
}

// NewBehaviorRuleVelocityAllOfWithDefaults instantiates a new BehaviorRuleVelocityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleVelocityAllOfWithDefaults() *BehaviorRuleVelocityAllOf {
	this := BehaviorRuleVelocityAllOf{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *BehaviorRuleVelocityAllOf) GetSettings() BehaviorRuleSettingsVelocity {
	if o == nil || o.Settings == nil {
		var ret BehaviorRuleSettingsVelocity
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleVelocityAllOf) GetSettingsOk() (*BehaviorRuleSettingsVelocity, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *BehaviorRuleVelocityAllOf) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given BehaviorRuleSettingsVelocity and assigns it to the Settings field.
func (o *BehaviorRuleVelocityAllOf) SetSettings(v BehaviorRuleSettingsVelocity) {
	o.Settings = &v
}

func (o BehaviorRuleVelocityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BehaviorRuleVelocityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBehaviorRuleVelocityAllOf := _BehaviorRuleVelocityAllOf{}

	err = json.Unmarshal(bytes, &varBehaviorRuleVelocityAllOf)
	if err == nil {
		*o = BehaviorRuleVelocityAllOf(varBehaviorRuleVelocityAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBehaviorRuleVelocityAllOf struct {
	value *BehaviorRuleVelocityAllOf
	isSet bool
}

func (v NullableBehaviorRuleVelocityAllOf) Get() *BehaviorRuleVelocityAllOf {
	return v.value
}

func (v *NullableBehaviorRuleVelocityAllOf) Set(val *BehaviorRuleVelocityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleVelocityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleVelocityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleVelocityAllOf(val *BehaviorRuleVelocityAllOf) *NullableBehaviorRuleVelocityAllOf {
	return &NullableBehaviorRuleVelocityAllOf{value: val, isSet: true}
}

func (v NullableBehaviorRuleVelocityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleVelocityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


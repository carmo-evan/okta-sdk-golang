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

// FCMPushProviderAllOf struct for FCMPushProviderAllOf
type FCMPushProviderAllOf struct {
	Configuration *FCMConfiguration `json:"configuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FCMPushProviderAllOf FCMPushProviderAllOf

// NewFCMPushProviderAllOf instantiates a new FCMPushProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFCMPushProviderAllOf() *FCMPushProviderAllOf {
	this := FCMPushProviderAllOf{}
	return &this
}

// NewFCMPushProviderAllOfWithDefaults instantiates a new FCMPushProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFCMPushProviderAllOfWithDefaults() *FCMPushProviderAllOf {
	this := FCMPushProviderAllOf{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *FCMPushProviderAllOf) GetConfiguration() FCMConfiguration {
	if o == nil || o.Configuration == nil {
		var ret FCMConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCMPushProviderAllOf) GetConfigurationOk() (*FCMConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *FCMPushProviderAllOf) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given FCMConfiguration and assigns it to the Configuration field.
func (o *FCMPushProviderAllOf) SetConfiguration(v FCMConfiguration) {
	o.Configuration = &v
}

func (o FCMPushProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FCMPushProviderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFCMPushProviderAllOf := _FCMPushProviderAllOf{}

	err = json.Unmarshal(bytes, &varFCMPushProviderAllOf)
	if err == nil {
		*o = FCMPushProviderAllOf(varFCMPushProviderAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "configuration")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFCMPushProviderAllOf struct {
	value *FCMPushProviderAllOf
	isSet bool
}

func (v NullableFCMPushProviderAllOf) Get() *FCMPushProviderAllOf {
	return v.value
}

func (v *NullableFCMPushProviderAllOf) Set(val *FCMPushProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFCMPushProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFCMPushProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFCMPushProviderAllOf(val *FCMPushProviderAllOf) *NullableFCMPushProviderAllOf {
	return &NullableFCMPushProviderAllOf{value: val, isSet: true}
}

func (v NullableFCMPushProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFCMPushProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


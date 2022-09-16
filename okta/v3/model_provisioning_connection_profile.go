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

// ProvisioningConnectionProfile struct for ProvisioningConnectionProfile
type ProvisioningConnectionProfile struct {
	AuthScheme *string `json:"authScheme,omitempty"`
	Token *string `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionProfile ProvisioningConnectionProfile

// NewProvisioningConnectionProfile instantiates a new ProvisioningConnectionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionProfile() *ProvisioningConnectionProfile {
	this := ProvisioningConnectionProfile{}
	return &this
}

// NewProvisioningConnectionProfileWithDefaults instantiates a new ProvisioningConnectionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionProfileWithDefaults() *ProvisioningConnectionProfile {
	this := ProvisioningConnectionProfile{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise.
func (o *ProvisioningConnectionProfile) GetAuthScheme() string {
	if o == nil || o.AuthScheme == nil {
		var ret string
		return ret
	}
	return *o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionProfile) GetAuthSchemeOk() (*string, bool) {
	if o == nil || o.AuthScheme == nil {
		return nil, false
	}
	return o.AuthScheme, true
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *ProvisioningConnectionProfile) HasAuthScheme() bool {
	if o != nil && o.AuthScheme != nil {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given string and assigns it to the AuthScheme field.
func (o *ProvisioningConnectionProfile) SetAuthScheme(v string) {
	o.AuthScheme = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProvisioningConnectionProfile) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionProfile) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProvisioningConnectionProfile) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProvisioningConnectionProfile) SetToken(v string) {
	o.Token = &v
}

func (o ProvisioningConnectionProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthScheme != nil {
		toSerialize["authScheme"] = o.AuthScheme
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnectionProfile) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnectionProfile := _ProvisioningConnectionProfile{}

	err = json.Unmarshal(bytes, &varProvisioningConnectionProfile)
	if err == nil {
		*o = ProvisioningConnectionProfile(varProvisioningConnectionProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProvisioningConnectionProfile struct {
	value *ProvisioningConnectionProfile
	isSet bool
}

func (v NullableProvisioningConnectionProfile) Get() *ProvisioningConnectionProfile {
	return v.value
}

func (v *NullableProvisioningConnectionProfile) Set(val *ProvisioningConnectionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionProfile(val *ProvisioningConnectionProfile) *NullableProvisioningConnectionProfile {
	return &NullableProvisioningConnectionProfile{value: val, isSet: true}
}

func (v NullableProvisioningConnectionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


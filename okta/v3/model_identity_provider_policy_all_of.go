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

// IdentityProviderPolicyAllOf struct for IdentityProviderPolicyAllOf
type IdentityProviderPolicyAllOf struct {
	AccountLink *PolicyAccountLink `json:"accountLink,omitempty"`
	Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	MaxClockSkew *int32 `json:"maxClockSkew,omitempty"`
	Provisioning *Provisioning `json:"provisioning,omitempty"`
	Subject *PolicySubject `json:"subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderPolicyAllOf IdentityProviderPolicyAllOf

// NewIdentityProviderPolicyAllOf instantiates a new IdentityProviderPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderPolicyAllOf() *IdentityProviderPolicyAllOf {
	this := IdentityProviderPolicyAllOf{}
	return &this
}

// NewIdentityProviderPolicyAllOfWithDefaults instantiates a new IdentityProviderPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderPolicyAllOfWithDefaults() *IdentityProviderPolicyAllOf {
	this := IdentityProviderPolicyAllOf{}
	return &this
}

// GetAccountLink returns the AccountLink field value if set, zero value otherwise.
func (o *IdentityProviderPolicyAllOf) GetAccountLink() PolicyAccountLink {
	if o == nil || o.AccountLink == nil {
		var ret PolicyAccountLink
		return ret
	}
	return *o.AccountLink
}

// GetAccountLinkOk returns a tuple with the AccountLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicyAllOf) GetAccountLinkOk() (*PolicyAccountLink, bool) {
	if o == nil || o.AccountLink == nil {
		return nil, false
	}
	return o.AccountLink, true
}

// HasAccountLink returns a boolean if a field has been set.
func (o *IdentityProviderPolicyAllOf) HasAccountLink() bool {
	if o != nil && o.AccountLink != nil {
		return true
	}

	return false
}

// SetAccountLink gets a reference to the given PolicyAccountLink and assigns it to the AccountLink field.
func (o *IdentityProviderPolicyAllOf) SetAccountLink(v PolicyAccountLink) {
	o.AccountLink = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *IdentityProviderPolicyAllOf) GetConditions() PolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret PolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicyAllOf) GetConditionsOk() (*PolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *IdentityProviderPolicyAllOf) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PolicyRuleConditions and assigns it to the Conditions field.
func (o *IdentityProviderPolicyAllOf) SetConditions(v PolicyRuleConditions) {
	o.Conditions = &v
}

// GetMaxClockSkew returns the MaxClockSkew field value if set, zero value otherwise.
func (o *IdentityProviderPolicyAllOf) GetMaxClockSkew() int32 {
	if o == nil || o.MaxClockSkew == nil {
		var ret int32
		return ret
	}
	return *o.MaxClockSkew
}

// GetMaxClockSkewOk returns a tuple with the MaxClockSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicyAllOf) GetMaxClockSkewOk() (*int32, bool) {
	if o == nil || o.MaxClockSkew == nil {
		return nil, false
	}
	return o.MaxClockSkew, true
}

// HasMaxClockSkew returns a boolean if a field has been set.
func (o *IdentityProviderPolicyAllOf) HasMaxClockSkew() bool {
	if o != nil && o.MaxClockSkew != nil {
		return true
	}

	return false
}

// SetMaxClockSkew gets a reference to the given int32 and assigns it to the MaxClockSkew field.
func (o *IdentityProviderPolicyAllOf) SetMaxClockSkew(v int32) {
	o.MaxClockSkew = &v
}

// GetProvisioning returns the Provisioning field value if set, zero value otherwise.
func (o *IdentityProviderPolicyAllOf) GetProvisioning() Provisioning {
	if o == nil || o.Provisioning == nil {
		var ret Provisioning
		return ret
	}
	return *o.Provisioning
}

// GetProvisioningOk returns a tuple with the Provisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicyAllOf) GetProvisioningOk() (*Provisioning, bool) {
	if o == nil || o.Provisioning == nil {
		return nil, false
	}
	return o.Provisioning, true
}

// HasProvisioning returns a boolean if a field has been set.
func (o *IdentityProviderPolicyAllOf) HasProvisioning() bool {
	if o != nil && o.Provisioning != nil {
		return true
	}

	return false
}

// SetProvisioning gets a reference to the given Provisioning and assigns it to the Provisioning field.
func (o *IdentityProviderPolicyAllOf) SetProvisioning(v Provisioning) {
	o.Provisioning = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *IdentityProviderPolicyAllOf) GetSubject() PolicySubject {
	if o == nil || o.Subject == nil {
		var ret PolicySubject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicyAllOf) GetSubjectOk() (*PolicySubject, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *IdentityProviderPolicyAllOf) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given PolicySubject and assigns it to the Subject field.
func (o *IdentityProviderPolicyAllOf) SetSubject(v PolicySubject) {
	o.Subject = &v
}

func (o IdentityProviderPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountLink != nil {
		toSerialize["accountLink"] = o.AccountLink
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.MaxClockSkew != nil {
		toSerialize["maxClockSkew"] = o.MaxClockSkew
	}
	if o.Provisioning != nil {
		toSerialize["provisioning"] = o.Provisioning
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderPolicyAllOf := _IdentityProviderPolicyAllOf{}

	err = json.Unmarshal(bytes, &varIdentityProviderPolicyAllOf)
	if err == nil {
		*o = IdentityProviderPolicyAllOf(varIdentityProviderPolicyAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accountLink")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "maxClockSkew")
		delete(additionalProperties, "provisioning")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderPolicyAllOf struct {
	value *IdentityProviderPolicyAllOf
	isSet bool
}

func (v NullableIdentityProviderPolicyAllOf) Get() *IdentityProviderPolicyAllOf {
	return v.value
}

func (v *NullableIdentityProviderPolicyAllOf) Set(val *IdentityProviderPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderPolicyAllOf(val *IdentityProviderPolicyAllOf) *NullableIdentityProviderPolicyAllOf {
	return &NullableIdentityProviderPolicyAllOf{value: val, isSet: true}
}

func (v NullableIdentityProviderPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


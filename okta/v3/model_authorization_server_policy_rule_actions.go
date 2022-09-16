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

// AuthorizationServerPolicyRuleActions struct for AuthorizationServerPolicyRuleActions
type AuthorizationServerPolicyRuleActions struct {
	Enroll *PolicyRuleActionsEnroll `json:"enroll,omitempty"`
	Idp *IdpPolicyRuleAction `json:"idp,omitempty"`
	PasswordChange *PasswordPolicyRuleAction `json:"passwordChange,omitempty"`
	SelfServicePasswordReset *PasswordPolicyRuleAction `json:"selfServicePasswordReset,omitempty"`
	SelfServiceUnlock *PasswordPolicyRuleAction `json:"selfServiceUnlock,omitempty"`
	Signon *OktaSignOnPolicyRuleSignonActions `json:"signon,omitempty"`
	Token *TokenAuthorizationServerPolicyRuleAction `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyRuleActions AuthorizationServerPolicyRuleActions

// NewAuthorizationServerPolicyRuleActions instantiates a new AuthorizationServerPolicyRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyRuleActions() *AuthorizationServerPolicyRuleActions {
	this := AuthorizationServerPolicyRuleActions{}
	return &this
}

// NewAuthorizationServerPolicyRuleActionsWithDefaults instantiates a new AuthorizationServerPolicyRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyRuleActionsWithDefaults() *AuthorizationServerPolicyRuleActions {
	this := AuthorizationServerPolicyRuleActions{}
	return &this
}

// GetEnroll returns the Enroll field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleActions) GetEnroll() PolicyRuleActionsEnroll {
	if o == nil || o.Enroll == nil {
		var ret PolicyRuleActionsEnroll
		return ret
	}
	return *o.Enroll
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleActions) GetEnrollOk() (*PolicyRuleActionsEnroll, bool) {
	if o == nil || o.Enroll == nil {
		return nil, false
	}
	return o.Enroll, true
}

// HasEnroll returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleActions) HasEnroll() bool {
	if o != nil && o.Enroll != nil {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given PolicyRuleActionsEnroll and assigns it to the Enroll field.
func (o *AuthorizationServerPolicyRuleActions) SetEnroll(v PolicyRuleActionsEnroll) {
	o.Enroll = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleActions) GetIdp() IdpPolicyRuleAction {
	if o == nil || o.Idp == nil {
		var ret IdpPolicyRuleAction
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleActions) GetIdpOk() (*IdpPolicyRuleAction, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleActions) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IdpPolicyRuleAction and assigns it to the Idp field.
func (o *AuthorizationServerPolicyRuleActions) SetIdp(v IdpPolicyRuleAction) {
	o.Idp = &v
}

// GetPasswordChange returns the PasswordChange field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleActions) GetPasswordChange() PasswordPolicyRuleAction {
	if o == nil || o.PasswordChange == nil {
		var ret PasswordPolicyRuleAction
		return ret
	}
	return *o.PasswordChange
}

// GetPasswordChangeOk returns a tuple with the PasswordChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleActions) GetPasswordChangeOk() (*PasswordPolicyRuleAction, bool) {
	if o == nil || o.PasswordChange == nil {
		return nil, false
	}
	return o.PasswordChange, true
}

// HasPasswordChange returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleActions) HasPasswordChange() bool {
	if o != nil && o.PasswordChange != nil {
		return true
	}

	return false
}

// SetPasswordChange gets a reference to the given PasswordPolicyRuleAction and assigns it to the PasswordChange field.
func (o *AuthorizationServerPolicyRuleActions) SetPasswordChange(v PasswordPolicyRuleAction) {
	o.PasswordChange = &v
}

// GetSelfServicePasswordReset returns the SelfServicePasswordReset field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleActions) GetSelfServicePasswordReset() PasswordPolicyRuleAction {
	if o == nil || o.SelfServicePasswordReset == nil {
		var ret PasswordPolicyRuleAction
		return ret
	}
	return *o.SelfServicePasswordReset
}

// GetSelfServicePasswordResetOk returns a tuple with the SelfServicePasswordReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleActions) GetSelfServicePasswordResetOk() (*PasswordPolicyRuleAction, bool) {
	if o == nil || o.SelfServicePasswordReset == nil {
		return nil, false
	}
	return o.SelfServicePasswordReset, true
}

// HasSelfServicePasswordReset returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleActions) HasSelfServicePasswordReset() bool {
	if o != nil && o.SelfServicePasswordReset != nil {
		return true
	}

	return false
}

// SetSelfServicePasswordReset gets a reference to the given PasswordPolicyRuleAction and assigns it to the SelfServicePasswordReset field.
func (o *AuthorizationServerPolicyRuleActions) SetSelfServicePasswordReset(v PasswordPolicyRuleAction) {
	o.SelfServicePasswordReset = &v
}

// GetSelfServiceUnlock returns the SelfServiceUnlock field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleActions) GetSelfServiceUnlock() PasswordPolicyRuleAction {
	if o == nil || o.SelfServiceUnlock == nil {
		var ret PasswordPolicyRuleAction
		return ret
	}
	return *o.SelfServiceUnlock
}

// GetSelfServiceUnlockOk returns a tuple with the SelfServiceUnlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleActions) GetSelfServiceUnlockOk() (*PasswordPolicyRuleAction, bool) {
	if o == nil || o.SelfServiceUnlock == nil {
		return nil, false
	}
	return o.SelfServiceUnlock, true
}

// HasSelfServiceUnlock returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleActions) HasSelfServiceUnlock() bool {
	if o != nil && o.SelfServiceUnlock != nil {
		return true
	}

	return false
}

// SetSelfServiceUnlock gets a reference to the given PasswordPolicyRuleAction and assigns it to the SelfServiceUnlock field.
func (o *AuthorizationServerPolicyRuleActions) SetSelfServiceUnlock(v PasswordPolicyRuleAction) {
	o.SelfServiceUnlock = &v
}

// GetSignon returns the Signon field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleActions) GetSignon() OktaSignOnPolicyRuleSignonActions {
	if o == nil || o.Signon == nil {
		var ret OktaSignOnPolicyRuleSignonActions
		return ret
	}
	return *o.Signon
}

// GetSignonOk returns a tuple with the Signon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleActions) GetSignonOk() (*OktaSignOnPolicyRuleSignonActions, bool) {
	if o == nil || o.Signon == nil {
		return nil, false
	}
	return o.Signon, true
}

// HasSignon returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleActions) HasSignon() bool {
	if o != nil && o.Signon != nil {
		return true
	}

	return false
}

// SetSignon gets a reference to the given OktaSignOnPolicyRuleSignonActions and assigns it to the Signon field.
func (o *AuthorizationServerPolicyRuleActions) SetSignon(v OktaSignOnPolicyRuleSignonActions) {
	o.Signon = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleActions) GetToken() TokenAuthorizationServerPolicyRuleAction {
	if o == nil || o.Token == nil {
		var ret TokenAuthorizationServerPolicyRuleAction
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleActions) GetTokenOk() (*TokenAuthorizationServerPolicyRuleAction, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleActions) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given TokenAuthorizationServerPolicyRuleAction and assigns it to the Token field.
func (o *AuthorizationServerPolicyRuleActions) SetToken(v TokenAuthorizationServerPolicyRuleAction) {
	o.Token = &v
}

func (o AuthorizationServerPolicyRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enroll != nil {
		toSerialize["enroll"] = o.Enroll
	}
	if o.Idp != nil {
		toSerialize["idp"] = o.Idp
	}
	if o.PasswordChange != nil {
		toSerialize["passwordChange"] = o.PasswordChange
	}
	if o.SelfServicePasswordReset != nil {
		toSerialize["selfServicePasswordReset"] = o.SelfServicePasswordReset
	}
	if o.SelfServiceUnlock != nil {
		toSerialize["selfServiceUnlock"] = o.SelfServiceUnlock
	}
	if o.Signon != nil {
		toSerialize["signon"] = o.Signon
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicyRuleActions) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicyRuleActions := _AuthorizationServerPolicyRuleActions{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyRuleActions)
	if err == nil {
		*o = AuthorizationServerPolicyRuleActions(varAuthorizationServerPolicyRuleActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "idp")
		delete(additionalProperties, "passwordChange")
		delete(additionalProperties, "selfServicePasswordReset")
		delete(additionalProperties, "selfServiceUnlock")
		delete(additionalProperties, "signon")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicyRuleActions struct {
	value *AuthorizationServerPolicyRuleActions
	isSet bool
}

func (v NullableAuthorizationServerPolicyRuleActions) Get() *AuthorizationServerPolicyRuleActions {
	return v.value
}

func (v *NullableAuthorizationServerPolicyRuleActions) Set(val *AuthorizationServerPolicyRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyRuleActions(val *AuthorizationServerPolicyRuleActions) *NullableAuthorizationServerPolicyRuleActions {
	return &NullableAuthorizationServerPolicyRuleActions{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


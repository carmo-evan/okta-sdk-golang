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

// ErrorErrorCausesInner struct for ErrorErrorCausesInner
type ErrorErrorCausesInner struct {
	ErrorSummary *string `json:"errorSummary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorErrorCausesInner ErrorErrorCausesInner

// NewErrorErrorCausesInner instantiates a new ErrorErrorCausesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorErrorCausesInner() *ErrorErrorCausesInner {
	this := ErrorErrorCausesInner{}
	return &this
}

// NewErrorErrorCausesInnerWithDefaults instantiates a new ErrorErrorCausesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorErrorCausesInnerWithDefaults() *ErrorErrorCausesInner {
	this := ErrorErrorCausesInner{}
	return &this
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *ErrorErrorCausesInner) GetErrorSummary() string {
	if o == nil || o.ErrorSummary == nil {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorErrorCausesInner) GetErrorSummaryOk() (*string, bool) {
	if o == nil || o.ErrorSummary == nil {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *ErrorErrorCausesInner) HasErrorSummary() bool {
	if o != nil && o.ErrorSummary != nil {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *ErrorErrorCausesInner) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

func (o ErrorErrorCausesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorSummary != nil {
		toSerialize["errorSummary"] = o.ErrorSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ErrorErrorCausesInner) UnmarshalJSON(bytes []byte) (err error) {
	varErrorErrorCausesInner := _ErrorErrorCausesInner{}

	err = json.Unmarshal(bytes, &varErrorErrorCausesInner)
	if err == nil {
		*o = ErrorErrorCausesInner(varErrorErrorCausesInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "errorSummary")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableErrorErrorCausesInner struct {
	value *ErrorErrorCausesInner
	isSet bool
}

func (v NullableErrorErrorCausesInner) Get() *ErrorErrorCausesInner {
	return v.value
}

func (v *NullableErrorErrorCausesInner) Set(val *ErrorErrorCausesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorErrorCausesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorErrorCausesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorErrorCausesInner(val *ErrorErrorCausesInner) *NullableErrorErrorCausesInner {
	return &NullableErrorErrorCausesInner{value: val, isSet: true}
}

func (v NullableErrorErrorCausesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorErrorCausesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


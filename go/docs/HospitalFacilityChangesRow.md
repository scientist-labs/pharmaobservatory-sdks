# HospitalFacilityChangesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccn** | Pointer to **NullableString** | CMS Certification Number of the hospital the event is about. | [optional] 
**AsOf** | Pointer to **NullableString** | The quarter-end date the change was detected at. | [optional] 
**ChangeType** | Pointer to **NullableString** | rename (same CCN, new provider name), appeared (CCN enters the series after its start), disappeared (CCN leaves it before its end), merge_candidate (a retired CCN with one successor) or split_candidate (a retired CCN with several). | [optional] 
**FromValue** | Pointer to **NullableString** | The provider name before the change (rename), or the last name of the retired CCN (disappeared, merge and split candidates). Null on appeared. | [optional] 
**ToValue** | Pointer to **NullableString** | The provider name after the change (rename), the first name of the new CCN (appeared, merge and split candidates). Null on disappeared. | [optional] 
**RelatedCcn** | Pointer to **NullableString** | The nearest later CCN that shares the retired CCN&#39;s normalised name, city and state. Set on merge_candidate and split_candidate rows. | [optional] 
**SuccessorCcn** | Pointer to **NullableString** | The CCN that absorbed the retired one. Set on merge_candidate rows only; null on split_candidate, where more than one CCN matches. | [optional] 

## Methods

### NewHospitalFacilityChangesRow

`func NewHospitalFacilityChangesRow() *HospitalFacilityChangesRow`

NewHospitalFacilityChangesRow instantiates a new HospitalFacilityChangesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHospitalFacilityChangesRowWithDefaults

`func NewHospitalFacilityChangesRowWithDefaults() *HospitalFacilityChangesRow`

NewHospitalFacilityChangesRowWithDefaults instantiates a new HospitalFacilityChangesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcn

`func (o *HospitalFacilityChangesRow) GetCcn() string`

GetCcn returns the Ccn field if non-nil, zero value otherwise.

### GetCcnOk

`func (o *HospitalFacilityChangesRow) GetCcnOk() (*string, bool)`

GetCcnOk returns a tuple with the Ccn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcn

`func (o *HospitalFacilityChangesRow) SetCcn(v string)`

SetCcn sets Ccn field to given value.

### HasCcn

`func (o *HospitalFacilityChangesRow) HasCcn() bool`

HasCcn returns a boolean if a field has been set.

### SetCcnNil

`func (o *HospitalFacilityChangesRow) SetCcnNil(b bool)`

 SetCcnNil sets the value for Ccn to be an explicit nil

### UnsetCcn
`func (o *HospitalFacilityChangesRow) UnsetCcn()`

UnsetCcn ensures that no value is present for Ccn, not even an explicit nil
### GetAsOf

`func (o *HospitalFacilityChangesRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *HospitalFacilityChangesRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *HospitalFacilityChangesRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *HospitalFacilityChangesRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *HospitalFacilityChangesRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *HospitalFacilityChangesRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetChangeType

`func (o *HospitalFacilityChangesRow) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *HospitalFacilityChangesRow) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *HospitalFacilityChangesRow) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *HospitalFacilityChangesRow) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### SetChangeTypeNil

`func (o *HospitalFacilityChangesRow) SetChangeTypeNil(b bool)`

 SetChangeTypeNil sets the value for ChangeType to be an explicit nil

### UnsetChangeType
`func (o *HospitalFacilityChangesRow) UnsetChangeType()`

UnsetChangeType ensures that no value is present for ChangeType, not even an explicit nil
### GetFromValue

`func (o *HospitalFacilityChangesRow) GetFromValue() string`

GetFromValue returns the FromValue field if non-nil, zero value otherwise.

### GetFromValueOk

`func (o *HospitalFacilityChangesRow) GetFromValueOk() (*string, bool)`

GetFromValueOk returns a tuple with the FromValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromValue

`func (o *HospitalFacilityChangesRow) SetFromValue(v string)`

SetFromValue sets FromValue field to given value.

### HasFromValue

`func (o *HospitalFacilityChangesRow) HasFromValue() bool`

HasFromValue returns a boolean if a field has been set.

### SetFromValueNil

`func (o *HospitalFacilityChangesRow) SetFromValueNil(b bool)`

 SetFromValueNil sets the value for FromValue to be an explicit nil

### UnsetFromValue
`func (o *HospitalFacilityChangesRow) UnsetFromValue()`

UnsetFromValue ensures that no value is present for FromValue, not even an explicit nil
### GetToValue

`func (o *HospitalFacilityChangesRow) GetToValue() string`

GetToValue returns the ToValue field if non-nil, zero value otherwise.

### GetToValueOk

`func (o *HospitalFacilityChangesRow) GetToValueOk() (*string, bool)`

GetToValueOk returns a tuple with the ToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToValue

`func (o *HospitalFacilityChangesRow) SetToValue(v string)`

SetToValue sets ToValue field to given value.

### HasToValue

`func (o *HospitalFacilityChangesRow) HasToValue() bool`

HasToValue returns a boolean if a field has been set.

### SetToValueNil

`func (o *HospitalFacilityChangesRow) SetToValueNil(b bool)`

 SetToValueNil sets the value for ToValue to be an explicit nil

### UnsetToValue
`func (o *HospitalFacilityChangesRow) UnsetToValue()`

UnsetToValue ensures that no value is present for ToValue, not even an explicit nil
### GetRelatedCcn

`func (o *HospitalFacilityChangesRow) GetRelatedCcn() string`

GetRelatedCcn returns the RelatedCcn field if non-nil, zero value otherwise.

### GetRelatedCcnOk

`func (o *HospitalFacilityChangesRow) GetRelatedCcnOk() (*string, bool)`

GetRelatedCcnOk returns a tuple with the RelatedCcn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedCcn

`func (o *HospitalFacilityChangesRow) SetRelatedCcn(v string)`

SetRelatedCcn sets RelatedCcn field to given value.

### HasRelatedCcn

`func (o *HospitalFacilityChangesRow) HasRelatedCcn() bool`

HasRelatedCcn returns a boolean if a field has been set.

### SetRelatedCcnNil

`func (o *HospitalFacilityChangesRow) SetRelatedCcnNil(b bool)`

 SetRelatedCcnNil sets the value for RelatedCcn to be an explicit nil

### UnsetRelatedCcn
`func (o *HospitalFacilityChangesRow) UnsetRelatedCcn()`

UnsetRelatedCcn ensures that no value is present for RelatedCcn, not even an explicit nil
### GetSuccessorCcn

`func (o *HospitalFacilityChangesRow) GetSuccessorCcn() string`

GetSuccessorCcn returns the SuccessorCcn field if non-nil, zero value otherwise.

### GetSuccessorCcnOk

`func (o *HospitalFacilityChangesRow) GetSuccessorCcnOk() (*string, bool)`

GetSuccessorCcnOk returns a tuple with the SuccessorCcn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessorCcn

`func (o *HospitalFacilityChangesRow) SetSuccessorCcn(v string)`

SetSuccessorCcn sets SuccessorCcn field to given value.

### HasSuccessorCcn

`func (o *HospitalFacilityChangesRow) HasSuccessorCcn() bool`

HasSuccessorCcn returns a boolean if a field has been set.

### SetSuccessorCcnNil

`func (o *HospitalFacilityChangesRow) SetSuccessorCcnNil(b bool)`

 SetSuccessorCcnNil sets the value for SuccessorCcn to be an explicit nil

### UnsetSuccessorCcn
`func (o *HospitalFacilityChangesRow) UnsetSuccessorCcn()`

UnsetSuccessorCcn ensures that no value is present for SuccessorCcn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



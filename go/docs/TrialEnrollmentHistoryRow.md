# TrialEnrollmentHistoryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**AsOf** | Pointer to **NullableString** | The registry snapshot date this row describes. | [optional] 
**Enrollment** | Pointer to **NullableInt32** | Enrollment registered at the snapshot (a planned target or an achieved count, see enrollment_type). | [optional] 
**EnrollmentType** | Pointer to **NullableString** | ACTUAL or ESTIMATED at the snapshot. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The trial&#39;s overall status code at the snapshot. | [optional] 
**StartDate** | Pointer to **NullableString** | Start date at the snapshot. | [optional] 
**PrimaryCompletionDate** | Pointer to **NullableString** | Primary completion date at the snapshot. | [optional] 
**CompletionDate** | Pointer to **NullableString** | Completion date at the snapshot. | [optional] 
**EnrollmentIsPlausible** | Pointer to **NullableBool** | False when the registered enrollment is a placeholder value rather than a count. | [optional] 
**EnrollmentFlagReason** | Pointer to **NullableString** | Why the enrollment was marked a placeholder (null when it was not). | [optional] 

## Methods

### NewTrialEnrollmentHistoryRow

`func NewTrialEnrollmentHistoryRow() *TrialEnrollmentHistoryRow`

NewTrialEnrollmentHistoryRow instantiates a new TrialEnrollmentHistoryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialEnrollmentHistoryRowWithDefaults

`func NewTrialEnrollmentHistoryRowWithDefaults() *TrialEnrollmentHistoryRow`

NewTrialEnrollmentHistoryRowWithDefaults instantiates a new TrialEnrollmentHistoryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialEnrollmentHistoryRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialEnrollmentHistoryRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialEnrollmentHistoryRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialEnrollmentHistoryRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialEnrollmentHistoryRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialEnrollmentHistoryRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetAsOf

`func (o *TrialEnrollmentHistoryRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *TrialEnrollmentHistoryRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *TrialEnrollmentHistoryRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *TrialEnrollmentHistoryRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *TrialEnrollmentHistoryRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *TrialEnrollmentHistoryRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetEnrollment

`func (o *TrialEnrollmentHistoryRow) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *TrialEnrollmentHistoryRow) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *TrialEnrollmentHistoryRow) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *TrialEnrollmentHistoryRow) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### SetEnrollmentNil

`func (o *TrialEnrollmentHistoryRow) SetEnrollmentNil(b bool)`

 SetEnrollmentNil sets the value for Enrollment to be an explicit nil

### UnsetEnrollment
`func (o *TrialEnrollmentHistoryRow) UnsetEnrollment()`

UnsetEnrollment ensures that no value is present for Enrollment, not even an explicit nil
### GetEnrollmentType

`func (o *TrialEnrollmentHistoryRow) GetEnrollmentType() string`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *TrialEnrollmentHistoryRow) GetEnrollmentTypeOk() (*string, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentType

`func (o *TrialEnrollmentHistoryRow) SetEnrollmentType(v string)`

SetEnrollmentType sets EnrollmentType field to given value.

### HasEnrollmentType

`func (o *TrialEnrollmentHistoryRow) HasEnrollmentType() bool`

HasEnrollmentType returns a boolean if a field has been set.

### SetEnrollmentTypeNil

`func (o *TrialEnrollmentHistoryRow) SetEnrollmentTypeNil(b bool)`

 SetEnrollmentTypeNil sets the value for EnrollmentType to be an explicit nil

### UnsetEnrollmentType
`func (o *TrialEnrollmentHistoryRow) UnsetEnrollmentType()`

UnsetEnrollmentType ensures that no value is present for EnrollmentType, not even an explicit nil
### GetOverallStatus

`func (o *TrialEnrollmentHistoryRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *TrialEnrollmentHistoryRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *TrialEnrollmentHistoryRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *TrialEnrollmentHistoryRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *TrialEnrollmentHistoryRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *TrialEnrollmentHistoryRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetStartDate

`func (o *TrialEnrollmentHistoryRow) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TrialEnrollmentHistoryRow) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TrialEnrollmentHistoryRow) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TrialEnrollmentHistoryRow) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TrialEnrollmentHistoryRow) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TrialEnrollmentHistoryRow) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetPrimaryCompletionDate

`func (o *TrialEnrollmentHistoryRow) GetPrimaryCompletionDate() string`

GetPrimaryCompletionDate returns the PrimaryCompletionDate field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateOk

`func (o *TrialEnrollmentHistoryRow) GetPrimaryCompletionDateOk() (*string, bool)`

GetPrimaryCompletionDateOk returns a tuple with the PrimaryCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDate

`func (o *TrialEnrollmentHistoryRow) SetPrimaryCompletionDate(v string)`

SetPrimaryCompletionDate sets PrimaryCompletionDate field to given value.

### HasPrimaryCompletionDate

`func (o *TrialEnrollmentHistoryRow) HasPrimaryCompletionDate() bool`

HasPrimaryCompletionDate returns a boolean if a field has been set.

### SetPrimaryCompletionDateNil

`func (o *TrialEnrollmentHistoryRow) SetPrimaryCompletionDateNil(b bool)`

 SetPrimaryCompletionDateNil sets the value for PrimaryCompletionDate to be an explicit nil

### UnsetPrimaryCompletionDate
`func (o *TrialEnrollmentHistoryRow) UnsetPrimaryCompletionDate()`

UnsetPrimaryCompletionDate ensures that no value is present for PrimaryCompletionDate, not even an explicit nil
### GetCompletionDate

`func (o *TrialEnrollmentHistoryRow) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *TrialEnrollmentHistoryRow) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *TrialEnrollmentHistoryRow) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *TrialEnrollmentHistoryRow) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDateNil

`func (o *TrialEnrollmentHistoryRow) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *TrialEnrollmentHistoryRow) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetEnrollmentIsPlausible

`func (o *TrialEnrollmentHistoryRow) GetEnrollmentIsPlausible() bool`

GetEnrollmentIsPlausible returns the EnrollmentIsPlausible field if non-nil, zero value otherwise.

### GetEnrollmentIsPlausibleOk

`func (o *TrialEnrollmentHistoryRow) GetEnrollmentIsPlausibleOk() (*bool, bool)`

GetEnrollmentIsPlausibleOk returns a tuple with the EnrollmentIsPlausible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentIsPlausible

`func (o *TrialEnrollmentHistoryRow) SetEnrollmentIsPlausible(v bool)`

SetEnrollmentIsPlausible sets EnrollmentIsPlausible field to given value.

### HasEnrollmentIsPlausible

`func (o *TrialEnrollmentHistoryRow) HasEnrollmentIsPlausible() bool`

HasEnrollmentIsPlausible returns a boolean if a field has been set.

### SetEnrollmentIsPlausibleNil

`func (o *TrialEnrollmentHistoryRow) SetEnrollmentIsPlausibleNil(b bool)`

 SetEnrollmentIsPlausibleNil sets the value for EnrollmentIsPlausible to be an explicit nil

### UnsetEnrollmentIsPlausible
`func (o *TrialEnrollmentHistoryRow) UnsetEnrollmentIsPlausible()`

UnsetEnrollmentIsPlausible ensures that no value is present for EnrollmentIsPlausible, not even an explicit nil
### GetEnrollmentFlagReason

`func (o *TrialEnrollmentHistoryRow) GetEnrollmentFlagReason() string`

GetEnrollmentFlagReason returns the EnrollmentFlagReason field if non-nil, zero value otherwise.

### GetEnrollmentFlagReasonOk

`func (o *TrialEnrollmentHistoryRow) GetEnrollmentFlagReasonOk() (*string, bool)`

GetEnrollmentFlagReasonOk returns a tuple with the EnrollmentFlagReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlagReason

`func (o *TrialEnrollmentHistoryRow) SetEnrollmentFlagReason(v string)`

SetEnrollmentFlagReason sets EnrollmentFlagReason field to given value.

### HasEnrollmentFlagReason

`func (o *TrialEnrollmentHistoryRow) HasEnrollmentFlagReason() bool`

HasEnrollmentFlagReason returns a boolean if a field has been set.

### SetEnrollmentFlagReasonNil

`func (o *TrialEnrollmentHistoryRow) SetEnrollmentFlagReasonNil(b bool)`

 SetEnrollmentFlagReasonNil sets the value for EnrollmentFlagReason to be an explicit nil

### UnsetEnrollmentFlagReason
`func (o *TrialEnrollmentHistoryRow) UnsetEnrollmentFlagReason()`

UnsetEnrollmentFlagReason ensures that no value is present for EnrollmentFlagReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



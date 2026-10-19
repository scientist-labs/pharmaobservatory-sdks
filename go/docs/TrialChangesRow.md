# TrialChangesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**ObservedAt** | Pointer to **NullableTime** | When the platform observed the change. | [optional] 
**EventType** | Pointer to **NullableString** | The kind of event (first_seen, status_change and so on). | [optional] 
**Field** | Pointer to **NullableString** | The record field that changed (null for a first_seen event). | [optional] 
**OldValue** | Pointer to **NullableString** | The value before the change. | [optional] 
**NewValue** | Pointer to **NullableString** | The value after the change. | [optional] 
**TrialTitle** | Pointer to **NullableString** | The trial&#39;s brief title. | [optional] 
**ConditionIds** | Pointer to **[]string** | MeSH (Medical Subject Headings) condition ids the trial is linked to, direct and through the MeSH tree. Filter on one id to scope changes to a condition&#39;s trials. | [optional] 
**DrugIds** | Pointer to **[]string** | Drug ids the trial&#39;s interventions resolved to. Filter on one id to scope changes to a drug&#39;s trials. | [optional] 
**SponsorOrgIds** | Pointer to **[]string** | Organization record ids of the sponsors on the trial, lead and collaborator. Filter on one id to scope changes to a sponsor&#39;s trials. | [optional] 

## Methods

### NewTrialChangesRow

`func NewTrialChangesRow() *TrialChangesRow`

NewTrialChangesRow instantiates a new TrialChangesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialChangesRowWithDefaults

`func NewTrialChangesRowWithDefaults() *TrialChangesRow`

NewTrialChangesRowWithDefaults instantiates a new TrialChangesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialChangesRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialChangesRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialChangesRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialChangesRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialChangesRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialChangesRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetObservedAt

`func (o *TrialChangesRow) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *TrialChangesRow) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *TrialChangesRow) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *TrialChangesRow) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### SetObservedAtNil

`func (o *TrialChangesRow) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *TrialChangesRow) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil
### GetEventType

`func (o *TrialChangesRow) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TrialChangesRow) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TrialChangesRow) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TrialChangesRow) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *TrialChangesRow) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *TrialChangesRow) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetField

`func (o *TrialChangesRow) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TrialChangesRow) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TrialChangesRow) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *TrialChangesRow) HasField() bool`

HasField returns a boolean if a field has been set.

### SetFieldNil

`func (o *TrialChangesRow) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *TrialChangesRow) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil
### GetOldValue

`func (o *TrialChangesRow) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *TrialChangesRow) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *TrialChangesRow) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *TrialChangesRow) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### SetOldValueNil

`func (o *TrialChangesRow) SetOldValueNil(b bool)`

 SetOldValueNil sets the value for OldValue to be an explicit nil

### UnsetOldValue
`func (o *TrialChangesRow) UnsetOldValue()`

UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
### GetNewValue

`func (o *TrialChangesRow) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *TrialChangesRow) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *TrialChangesRow) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *TrialChangesRow) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### SetNewValueNil

`func (o *TrialChangesRow) SetNewValueNil(b bool)`

 SetNewValueNil sets the value for NewValue to be an explicit nil

### UnsetNewValue
`func (o *TrialChangesRow) UnsetNewValue()`

UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
### GetTrialTitle

`func (o *TrialChangesRow) GetTrialTitle() string`

GetTrialTitle returns the TrialTitle field if non-nil, zero value otherwise.

### GetTrialTitleOk

`func (o *TrialChangesRow) GetTrialTitleOk() (*string, bool)`

GetTrialTitleOk returns a tuple with the TrialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialTitle

`func (o *TrialChangesRow) SetTrialTitle(v string)`

SetTrialTitle sets TrialTitle field to given value.

### HasTrialTitle

`func (o *TrialChangesRow) HasTrialTitle() bool`

HasTrialTitle returns a boolean if a field has been set.

### SetTrialTitleNil

`func (o *TrialChangesRow) SetTrialTitleNil(b bool)`

 SetTrialTitleNil sets the value for TrialTitle to be an explicit nil

### UnsetTrialTitle
`func (o *TrialChangesRow) UnsetTrialTitle()`

UnsetTrialTitle ensures that no value is present for TrialTitle, not even an explicit nil
### GetConditionIds

`func (o *TrialChangesRow) GetConditionIds() []string`

GetConditionIds returns the ConditionIds field if non-nil, zero value otherwise.

### GetConditionIdsOk

`func (o *TrialChangesRow) GetConditionIdsOk() (*[]string, bool)`

GetConditionIdsOk returns a tuple with the ConditionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIds

`func (o *TrialChangesRow) SetConditionIds(v []string)`

SetConditionIds sets ConditionIds field to given value.

### HasConditionIds

`func (o *TrialChangesRow) HasConditionIds() bool`

HasConditionIds returns a boolean if a field has been set.

### SetConditionIdsNil

`func (o *TrialChangesRow) SetConditionIdsNil(b bool)`

 SetConditionIdsNil sets the value for ConditionIds to be an explicit nil

### UnsetConditionIds
`func (o *TrialChangesRow) UnsetConditionIds()`

UnsetConditionIds ensures that no value is present for ConditionIds, not even an explicit nil
### GetDrugIds

`func (o *TrialChangesRow) GetDrugIds() []string`

GetDrugIds returns the DrugIds field if non-nil, zero value otherwise.

### GetDrugIdsOk

`func (o *TrialChangesRow) GetDrugIdsOk() (*[]string, bool)`

GetDrugIdsOk returns a tuple with the DrugIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugIds

`func (o *TrialChangesRow) SetDrugIds(v []string)`

SetDrugIds sets DrugIds field to given value.

### HasDrugIds

`func (o *TrialChangesRow) HasDrugIds() bool`

HasDrugIds returns a boolean if a field has been set.

### SetDrugIdsNil

`func (o *TrialChangesRow) SetDrugIdsNil(b bool)`

 SetDrugIdsNil sets the value for DrugIds to be an explicit nil

### UnsetDrugIds
`func (o *TrialChangesRow) UnsetDrugIds()`

UnsetDrugIds ensures that no value is present for DrugIds, not even an explicit nil
### GetSponsorOrgIds

`func (o *TrialChangesRow) GetSponsorOrgIds() []string`

GetSponsorOrgIds returns the SponsorOrgIds field if non-nil, zero value otherwise.

### GetSponsorOrgIdsOk

`func (o *TrialChangesRow) GetSponsorOrgIdsOk() (*[]string, bool)`

GetSponsorOrgIdsOk returns a tuple with the SponsorOrgIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgIds

`func (o *TrialChangesRow) SetSponsorOrgIds(v []string)`

SetSponsorOrgIds sets SponsorOrgIds field to given value.

### HasSponsorOrgIds

`func (o *TrialChangesRow) HasSponsorOrgIds() bool`

HasSponsorOrgIds returns a boolean if a field has been set.

### SetSponsorOrgIdsNil

`func (o *TrialChangesRow) SetSponsorOrgIdsNil(b bool)`

 SetSponsorOrgIdsNil sets the value for SponsorOrgIds to be an explicit nil

### UnsetSponsorOrgIds
`func (o *TrialChangesRow) UnsetSponsorOrgIds()`

UnsetSponsorOrgIds ensures that no value is present for SponsorOrgIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SiteTrialConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key (a MeSH, Medical Subject Headings, descriptor id). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s name. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall status of the trial. | [optional] 
**StartYear** | Pointer to **NullableInt32** | Year of the trial start. | [optional] 

## Methods

### NewSiteTrialConditionsRow

`func NewSiteTrialConditionsRow() *SiteTrialConditionsRow`

NewSiteTrialConditionsRow instantiates a new SiteTrialConditionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteTrialConditionsRowWithDefaults

`func NewSiteTrialConditionsRowWithDefaults() *SiteTrialConditionsRow`

NewSiteTrialConditionsRowWithDefaults instantiates a new SiteTrialConditionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *SiteTrialConditionsRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteTrialConditionsRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteTrialConditionsRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteTrialConditionsRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *SiteTrialConditionsRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *SiteTrialConditionsRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetTrialId

`func (o *SiteTrialConditionsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *SiteTrialConditionsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *SiteTrialConditionsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *SiteTrialConditionsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *SiteTrialConditionsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *SiteTrialConditionsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetConditionId

`func (o *SiteTrialConditionsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *SiteTrialConditionsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *SiteTrialConditionsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *SiteTrialConditionsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *SiteTrialConditionsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *SiteTrialConditionsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *SiteTrialConditionsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *SiteTrialConditionsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *SiteTrialConditionsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *SiteTrialConditionsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *SiteTrialConditionsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *SiteTrialConditionsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetOverallStatus

`func (o *SiteTrialConditionsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *SiteTrialConditionsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *SiteTrialConditionsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *SiteTrialConditionsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *SiteTrialConditionsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *SiteTrialConditionsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetStartYear

`func (o *SiteTrialConditionsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *SiteTrialConditionsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *SiteTrialConditionsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *SiteTrialConditionsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *SiteTrialConditionsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *SiteTrialConditionsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



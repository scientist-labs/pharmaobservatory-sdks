# SiteCroTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**CroOrgId** | Pointer to **NullableString** | Legal-entity level: the CRO&#39;s organization record. | [optional] 
**CroName** | Pointer to **NullableString** | The CRO organization&#39;s name. | [optional] 
**CroCompanyId** | Pointer to **NullableString** | Company level: the company the CRO belongs to (null when unresolved). | [optional] 
**CroCompanyName** | Pointer to **NullableString** | Display name of the CRO&#39;s company. | [optional] 
**CroGroupId** | Pointer to **NullableString** | Group level: the CRO&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**CroGroupName** | Pointer to **NullableString** | Display name of the CRO&#39;s group. | [optional] 
**CroAttributionMethod** | Pointer to **NullableString** | How the CRO was inferred (sponsor_match, collaborator_match or llm_extraction). | [optional] 
**CroConfidence** | Pointer to **NullableFloat32** | Confidence of the inference, from 0 to 1. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall status of the trial. | [optional] 
**StartYear** | Pointer to **NullableInt32** | Year of the trial start. | [optional] 

## Methods

### NewSiteCroTrialsRow

`func NewSiteCroTrialsRow() *SiteCroTrialsRow`

NewSiteCroTrialsRow instantiates a new SiteCroTrialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteCroTrialsRowWithDefaults

`func NewSiteCroTrialsRowWithDefaults() *SiteCroTrialsRow`

NewSiteCroTrialsRowWithDefaults instantiates a new SiteCroTrialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *SiteCroTrialsRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteCroTrialsRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteCroTrialsRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteCroTrialsRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *SiteCroTrialsRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *SiteCroTrialsRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetTrialId

`func (o *SiteCroTrialsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *SiteCroTrialsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *SiteCroTrialsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *SiteCroTrialsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *SiteCroTrialsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *SiteCroTrialsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetCroOrgId

`func (o *SiteCroTrialsRow) GetCroOrgId() string`

GetCroOrgId returns the CroOrgId field if non-nil, zero value otherwise.

### GetCroOrgIdOk

`func (o *SiteCroTrialsRow) GetCroOrgIdOk() (*string, bool)`

GetCroOrgIdOk returns a tuple with the CroOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroOrgId

`func (o *SiteCroTrialsRow) SetCroOrgId(v string)`

SetCroOrgId sets CroOrgId field to given value.

### HasCroOrgId

`func (o *SiteCroTrialsRow) HasCroOrgId() bool`

HasCroOrgId returns a boolean if a field has been set.

### SetCroOrgIdNil

`func (o *SiteCroTrialsRow) SetCroOrgIdNil(b bool)`

 SetCroOrgIdNil sets the value for CroOrgId to be an explicit nil

### UnsetCroOrgId
`func (o *SiteCroTrialsRow) UnsetCroOrgId()`

UnsetCroOrgId ensures that no value is present for CroOrgId, not even an explicit nil
### GetCroName

`func (o *SiteCroTrialsRow) GetCroName() string`

GetCroName returns the CroName field if non-nil, zero value otherwise.

### GetCroNameOk

`func (o *SiteCroTrialsRow) GetCroNameOk() (*string, bool)`

GetCroNameOk returns a tuple with the CroName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroName

`func (o *SiteCroTrialsRow) SetCroName(v string)`

SetCroName sets CroName field to given value.

### HasCroName

`func (o *SiteCroTrialsRow) HasCroName() bool`

HasCroName returns a boolean if a field has been set.

### SetCroNameNil

`func (o *SiteCroTrialsRow) SetCroNameNil(b bool)`

 SetCroNameNil sets the value for CroName to be an explicit nil

### UnsetCroName
`func (o *SiteCroTrialsRow) UnsetCroName()`

UnsetCroName ensures that no value is present for CroName, not even an explicit nil
### GetCroCompanyId

`func (o *SiteCroTrialsRow) GetCroCompanyId() string`

GetCroCompanyId returns the CroCompanyId field if non-nil, zero value otherwise.

### GetCroCompanyIdOk

`func (o *SiteCroTrialsRow) GetCroCompanyIdOk() (*string, bool)`

GetCroCompanyIdOk returns a tuple with the CroCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroCompanyId

`func (o *SiteCroTrialsRow) SetCroCompanyId(v string)`

SetCroCompanyId sets CroCompanyId field to given value.

### HasCroCompanyId

`func (o *SiteCroTrialsRow) HasCroCompanyId() bool`

HasCroCompanyId returns a boolean if a field has been set.

### SetCroCompanyIdNil

`func (o *SiteCroTrialsRow) SetCroCompanyIdNil(b bool)`

 SetCroCompanyIdNil sets the value for CroCompanyId to be an explicit nil

### UnsetCroCompanyId
`func (o *SiteCroTrialsRow) UnsetCroCompanyId()`

UnsetCroCompanyId ensures that no value is present for CroCompanyId, not even an explicit nil
### GetCroCompanyName

`func (o *SiteCroTrialsRow) GetCroCompanyName() string`

GetCroCompanyName returns the CroCompanyName field if non-nil, zero value otherwise.

### GetCroCompanyNameOk

`func (o *SiteCroTrialsRow) GetCroCompanyNameOk() (*string, bool)`

GetCroCompanyNameOk returns a tuple with the CroCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroCompanyName

`func (o *SiteCroTrialsRow) SetCroCompanyName(v string)`

SetCroCompanyName sets CroCompanyName field to given value.

### HasCroCompanyName

`func (o *SiteCroTrialsRow) HasCroCompanyName() bool`

HasCroCompanyName returns a boolean if a field has been set.

### SetCroCompanyNameNil

`func (o *SiteCroTrialsRow) SetCroCompanyNameNil(b bool)`

 SetCroCompanyNameNil sets the value for CroCompanyName to be an explicit nil

### UnsetCroCompanyName
`func (o *SiteCroTrialsRow) UnsetCroCompanyName()`

UnsetCroCompanyName ensures that no value is present for CroCompanyName, not even an explicit nil
### GetCroGroupId

`func (o *SiteCroTrialsRow) GetCroGroupId() string`

GetCroGroupId returns the CroGroupId field if non-nil, zero value otherwise.

### GetCroGroupIdOk

`func (o *SiteCroTrialsRow) GetCroGroupIdOk() (*string, bool)`

GetCroGroupIdOk returns a tuple with the CroGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroGroupId

`func (o *SiteCroTrialsRow) SetCroGroupId(v string)`

SetCroGroupId sets CroGroupId field to given value.

### HasCroGroupId

`func (o *SiteCroTrialsRow) HasCroGroupId() bool`

HasCroGroupId returns a boolean if a field has been set.

### SetCroGroupIdNil

`func (o *SiteCroTrialsRow) SetCroGroupIdNil(b bool)`

 SetCroGroupIdNil sets the value for CroGroupId to be an explicit nil

### UnsetCroGroupId
`func (o *SiteCroTrialsRow) UnsetCroGroupId()`

UnsetCroGroupId ensures that no value is present for CroGroupId, not even an explicit nil
### GetCroGroupName

`func (o *SiteCroTrialsRow) GetCroGroupName() string`

GetCroGroupName returns the CroGroupName field if non-nil, zero value otherwise.

### GetCroGroupNameOk

`func (o *SiteCroTrialsRow) GetCroGroupNameOk() (*string, bool)`

GetCroGroupNameOk returns a tuple with the CroGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroGroupName

`func (o *SiteCroTrialsRow) SetCroGroupName(v string)`

SetCroGroupName sets CroGroupName field to given value.

### HasCroGroupName

`func (o *SiteCroTrialsRow) HasCroGroupName() bool`

HasCroGroupName returns a boolean if a field has been set.

### SetCroGroupNameNil

`func (o *SiteCroTrialsRow) SetCroGroupNameNil(b bool)`

 SetCroGroupNameNil sets the value for CroGroupName to be an explicit nil

### UnsetCroGroupName
`func (o *SiteCroTrialsRow) UnsetCroGroupName()`

UnsetCroGroupName ensures that no value is present for CroGroupName, not even an explicit nil
### GetCroAttributionMethod

`func (o *SiteCroTrialsRow) GetCroAttributionMethod() string`

GetCroAttributionMethod returns the CroAttributionMethod field if non-nil, zero value otherwise.

### GetCroAttributionMethodOk

`func (o *SiteCroTrialsRow) GetCroAttributionMethodOk() (*string, bool)`

GetCroAttributionMethodOk returns a tuple with the CroAttributionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroAttributionMethod

`func (o *SiteCroTrialsRow) SetCroAttributionMethod(v string)`

SetCroAttributionMethod sets CroAttributionMethod field to given value.

### HasCroAttributionMethod

`func (o *SiteCroTrialsRow) HasCroAttributionMethod() bool`

HasCroAttributionMethod returns a boolean if a field has been set.

### SetCroAttributionMethodNil

`func (o *SiteCroTrialsRow) SetCroAttributionMethodNil(b bool)`

 SetCroAttributionMethodNil sets the value for CroAttributionMethod to be an explicit nil

### UnsetCroAttributionMethod
`func (o *SiteCroTrialsRow) UnsetCroAttributionMethod()`

UnsetCroAttributionMethod ensures that no value is present for CroAttributionMethod, not even an explicit nil
### GetCroConfidence

`func (o *SiteCroTrialsRow) GetCroConfidence() float32`

GetCroConfidence returns the CroConfidence field if non-nil, zero value otherwise.

### GetCroConfidenceOk

`func (o *SiteCroTrialsRow) GetCroConfidenceOk() (*float32, bool)`

GetCroConfidenceOk returns a tuple with the CroConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCroConfidence

`func (o *SiteCroTrialsRow) SetCroConfidence(v float32)`

SetCroConfidence sets CroConfidence field to given value.

### HasCroConfidence

`func (o *SiteCroTrialsRow) HasCroConfidence() bool`

HasCroConfidence returns a boolean if a field has been set.

### SetCroConfidenceNil

`func (o *SiteCroTrialsRow) SetCroConfidenceNil(b bool)`

 SetCroConfidenceNil sets the value for CroConfidence to be an explicit nil

### UnsetCroConfidence
`func (o *SiteCroTrialsRow) UnsetCroConfidence()`

UnsetCroConfidence ensures that no value is present for CroConfidence, not even an explicit nil
### GetOverallStatus

`func (o *SiteCroTrialsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *SiteCroTrialsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *SiteCroTrialsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *SiteCroTrialsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *SiteCroTrialsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *SiteCroTrialsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetStartYear

`func (o *SiteCroTrialsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *SiteCroTrialsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *SiteCroTrialsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *SiteCroTrialsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *SiteCroTrialsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *SiteCroTrialsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



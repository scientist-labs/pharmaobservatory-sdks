# CroTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**VendorOrgId** | Pointer to **NullableString** | Legal-entity level: the contract research organization&#39;s organization record, as filed. | [optional] 
**VendorName** | Pointer to **NullableString** | The organization&#39;s name as filed. | [optional] 
**VendorCompanyId** | Pointer to **NullableString** | Company level: the contract research organization&#39;s company (null when unresolved). Group by this column for the company leaderboard. | [optional] 
**VendorCompanyName** | Pointer to **NullableString** | Display name of the contract research organization&#39;s company. | [optional] 
**VendorGroupId** | Pointer to **NullableString** | Group level: the contract research organization&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**VendorGroupName** | Pointer to **NullableString** | Display name of the contract research organization&#39;s group. | [optional] 
**AttributionMethod** | Pointer to **NullableString** | How the vendor was inferred for the trial: sponsor_match or collaborator_match (the organization is named on the registry record and is on the curated list) or llm_extraction. | [optional] 
**Confidence** | Pointer to **NullableFloat32** | Confidence of the inference, 0 to 1. | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code of the trial (PHASE1, PHASE2, PHASE3 and so on). Null when the registry gives none. | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall status code of the trial (RECRUITING, COMPLETED, TERMINATED and so on). | [optional] 
**StartYear** | Pointer to **NullableInt32** | Start year of the trial. | [optional] 
**LeadSponsorName** | Pointer to **NullableString** | The trial&#39;s lead sponsor as filed. | [optional] 
**FacilityCountries** | Pointer to **[]string** | Countries of the trial&#39;s facilities. Filter on one country (for example facility_countries&#x3D;United States) to match trials that list it. | [optional] 
**ConditionIds** | Pointer to **[]string** | The MeSH (Medical Subject Headings) condition ids of the trial, direct and ancestor (D followed by six to nine digits). Filter on one id (for example condition_ids&#x3D;D003316) to scope the rows to an indication. | [optional] 

## Methods

### NewCroTrialsRow

`func NewCroTrialsRow() *CroTrialsRow`

NewCroTrialsRow instantiates a new CroTrialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCroTrialsRowWithDefaults

`func NewCroTrialsRowWithDefaults() *CroTrialsRow`

NewCroTrialsRowWithDefaults instantiates a new CroTrialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *CroTrialsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *CroTrialsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *CroTrialsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *CroTrialsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *CroTrialsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *CroTrialsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetVendorOrgId

`func (o *CroTrialsRow) GetVendorOrgId() string`

GetVendorOrgId returns the VendorOrgId field if non-nil, zero value otherwise.

### GetVendorOrgIdOk

`func (o *CroTrialsRow) GetVendorOrgIdOk() (*string, bool)`

GetVendorOrgIdOk returns a tuple with the VendorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorOrgId

`func (o *CroTrialsRow) SetVendorOrgId(v string)`

SetVendorOrgId sets VendorOrgId field to given value.

### HasVendorOrgId

`func (o *CroTrialsRow) HasVendorOrgId() bool`

HasVendorOrgId returns a boolean if a field has been set.

### SetVendorOrgIdNil

`func (o *CroTrialsRow) SetVendorOrgIdNil(b bool)`

 SetVendorOrgIdNil sets the value for VendorOrgId to be an explicit nil

### UnsetVendorOrgId
`func (o *CroTrialsRow) UnsetVendorOrgId()`

UnsetVendorOrgId ensures that no value is present for VendorOrgId, not even an explicit nil
### GetVendorName

`func (o *CroTrialsRow) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *CroTrialsRow) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *CroTrialsRow) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *CroTrialsRow) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### SetVendorNameNil

`func (o *CroTrialsRow) SetVendorNameNil(b bool)`

 SetVendorNameNil sets the value for VendorName to be an explicit nil

### UnsetVendorName
`func (o *CroTrialsRow) UnsetVendorName()`

UnsetVendorName ensures that no value is present for VendorName, not even an explicit nil
### GetVendorCompanyId

`func (o *CroTrialsRow) GetVendorCompanyId() string`

GetVendorCompanyId returns the VendorCompanyId field if non-nil, zero value otherwise.

### GetVendorCompanyIdOk

`func (o *CroTrialsRow) GetVendorCompanyIdOk() (*string, bool)`

GetVendorCompanyIdOk returns a tuple with the VendorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCompanyId

`func (o *CroTrialsRow) SetVendorCompanyId(v string)`

SetVendorCompanyId sets VendorCompanyId field to given value.

### HasVendorCompanyId

`func (o *CroTrialsRow) HasVendorCompanyId() bool`

HasVendorCompanyId returns a boolean if a field has been set.

### SetVendorCompanyIdNil

`func (o *CroTrialsRow) SetVendorCompanyIdNil(b bool)`

 SetVendorCompanyIdNil sets the value for VendorCompanyId to be an explicit nil

### UnsetVendorCompanyId
`func (o *CroTrialsRow) UnsetVendorCompanyId()`

UnsetVendorCompanyId ensures that no value is present for VendorCompanyId, not even an explicit nil
### GetVendorCompanyName

`func (o *CroTrialsRow) GetVendorCompanyName() string`

GetVendorCompanyName returns the VendorCompanyName field if non-nil, zero value otherwise.

### GetVendorCompanyNameOk

`func (o *CroTrialsRow) GetVendorCompanyNameOk() (*string, bool)`

GetVendorCompanyNameOk returns a tuple with the VendorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCompanyName

`func (o *CroTrialsRow) SetVendorCompanyName(v string)`

SetVendorCompanyName sets VendorCompanyName field to given value.

### HasVendorCompanyName

`func (o *CroTrialsRow) HasVendorCompanyName() bool`

HasVendorCompanyName returns a boolean if a field has been set.

### SetVendorCompanyNameNil

`func (o *CroTrialsRow) SetVendorCompanyNameNil(b bool)`

 SetVendorCompanyNameNil sets the value for VendorCompanyName to be an explicit nil

### UnsetVendorCompanyName
`func (o *CroTrialsRow) UnsetVendorCompanyName()`

UnsetVendorCompanyName ensures that no value is present for VendorCompanyName, not even an explicit nil
### GetVendorGroupId

`func (o *CroTrialsRow) GetVendorGroupId() string`

GetVendorGroupId returns the VendorGroupId field if non-nil, zero value otherwise.

### GetVendorGroupIdOk

`func (o *CroTrialsRow) GetVendorGroupIdOk() (*string, bool)`

GetVendorGroupIdOk returns a tuple with the VendorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorGroupId

`func (o *CroTrialsRow) SetVendorGroupId(v string)`

SetVendorGroupId sets VendorGroupId field to given value.

### HasVendorGroupId

`func (o *CroTrialsRow) HasVendorGroupId() bool`

HasVendorGroupId returns a boolean if a field has been set.

### SetVendorGroupIdNil

`func (o *CroTrialsRow) SetVendorGroupIdNil(b bool)`

 SetVendorGroupIdNil sets the value for VendorGroupId to be an explicit nil

### UnsetVendorGroupId
`func (o *CroTrialsRow) UnsetVendorGroupId()`

UnsetVendorGroupId ensures that no value is present for VendorGroupId, not even an explicit nil
### GetVendorGroupName

`func (o *CroTrialsRow) GetVendorGroupName() string`

GetVendorGroupName returns the VendorGroupName field if non-nil, zero value otherwise.

### GetVendorGroupNameOk

`func (o *CroTrialsRow) GetVendorGroupNameOk() (*string, bool)`

GetVendorGroupNameOk returns a tuple with the VendorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorGroupName

`func (o *CroTrialsRow) SetVendorGroupName(v string)`

SetVendorGroupName sets VendorGroupName field to given value.

### HasVendorGroupName

`func (o *CroTrialsRow) HasVendorGroupName() bool`

HasVendorGroupName returns a boolean if a field has been set.

### SetVendorGroupNameNil

`func (o *CroTrialsRow) SetVendorGroupNameNil(b bool)`

 SetVendorGroupNameNil sets the value for VendorGroupName to be an explicit nil

### UnsetVendorGroupName
`func (o *CroTrialsRow) UnsetVendorGroupName()`

UnsetVendorGroupName ensures that no value is present for VendorGroupName, not even an explicit nil
### GetAttributionMethod

`func (o *CroTrialsRow) GetAttributionMethod() string`

GetAttributionMethod returns the AttributionMethod field if non-nil, zero value otherwise.

### GetAttributionMethodOk

`func (o *CroTrialsRow) GetAttributionMethodOk() (*string, bool)`

GetAttributionMethodOk returns a tuple with the AttributionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionMethod

`func (o *CroTrialsRow) SetAttributionMethod(v string)`

SetAttributionMethod sets AttributionMethod field to given value.

### HasAttributionMethod

`func (o *CroTrialsRow) HasAttributionMethod() bool`

HasAttributionMethod returns a boolean if a field has been set.

### SetAttributionMethodNil

`func (o *CroTrialsRow) SetAttributionMethodNil(b bool)`

 SetAttributionMethodNil sets the value for AttributionMethod to be an explicit nil

### UnsetAttributionMethod
`func (o *CroTrialsRow) UnsetAttributionMethod()`

UnsetAttributionMethod ensures that no value is present for AttributionMethod, not even an explicit nil
### GetConfidence

`func (o *CroTrialsRow) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CroTrialsRow) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CroTrialsRow) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *CroTrialsRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *CroTrialsRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *CroTrialsRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetPhase

`func (o *CroTrialsRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CroTrialsRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CroTrialsRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CroTrialsRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *CroTrialsRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *CroTrialsRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetOverallStatus

`func (o *CroTrialsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *CroTrialsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *CroTrialsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *CroTrialsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *CroTrialsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *CroTrialsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetStartYear

`func (o *CroTrialsRow) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *CroTrialsRow) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *CroTrialsRow) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.

### HasStartYear

`func (o *CroTrialsRow) HasStartYear() bool`

HasStartYear returns a boolean if a field has been set.

### SetStartYearNil

`func (o *CroTrialsRow) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *CroTrialsRow) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetLeadSponsorName

`func (o *CroTrialsRow) GetLeadSponsorName() string`

GetLeadSponsorName returns the LeadSponsorName field if non-nil, zero value otherwise.

### GetLeadSponsorNameOk

`func (o *CroTrialsRow) GetLeadSponsorNameOk() (*string, bool)`

GetLeadSponsorNameOk returns a tuple with the LeadSponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorName

`func (o *CroTrialsRow) SetLeadSponsorName(v string)`

SetLeadSponsorName sets LeadSponsorName field to given value.

### HasLeadSponsorName

`func (o *CroTrialsRow) HasLeadSponsorName() bool`

HasLeadSponsorName returns a boolean if a field has been set.

### SetLeadSponsorNameNil

`func (o *CroTrialsRow) SetLeadSponsorNameNil(b bool)`

 SetLeadSponsorNameNil sets the value for LeadSponsorName to be an explicit nil

### UnsetLeadSponsorName
`func (o *CroTrialsRow) UnsetLeadSponsorName()`

UnsetLeadSponsorName ensures that no value is present for LeadSponsorName, not even an explicit nil
### GetFacilityCountries

`func (o *CroTrialsRow) GetFacilityCountries() []string`

GetFacilityCountries returns the FacilityCountries field if non-nil, zero value otherwise.

### GetFacilityCountriesOk

`func (o *CroTrialsRow) GetFacilityCountriesOk() (*[]string, bool)`

GetFacilityCountriesOk returns a tuple with the FacilityCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountries

`func (o *CroTrialsRow) SetFacilityCountries(v []string)`

SetFacilityCountries sets FacilityCountries field to given value.

### HasFacilityCountries

`func (o *CroTrialsRow) HasFacilityCountries() bool`

HasFacilityCountries returns a boolean if a field has been set.

### SetFacilityCountriesNil

`func (o *CroTrialsRow) SetFacilityCountriesNil(b bool)`

 SetFacilityCountriesNil sets the value for FacilityCountries to be an explicit nil

### UnsetFacilityCountries
`func (o *CroTrialsRow) UnsetFacilityCountries()`

UnsetFacilityCountries ensures that no value is present for FacilityCountries, not even an explicit nil
### GetConditionIds

`func (o *CroTrialsRow) GetConditionIds() []string`

GetConditionIds returns the ConditionIds field if non-nil, zero value otherwise.

### GetConditionIdsOk

`func (o *CroTrialsRow) GetConditionIdsOk() (*[]string, bool)`

GetConditionIdsOk returns a tuple with the ConditionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIds

`func (o *CroTrialsRow) SetConditionIds(v []string)`

SetConditionIds sets ConditionIds field to given value.

### HasConditionIds

`func (o *CroTrialsRow) HasConditionIds() bool`

HasConditionIds returns a boolean if a field has been set.

### SetConditionIdsNil

`func (o *CroTrialsRow) SetConditionIdsNil(b bool)`

 SetConditionIdsNil sets the value for ConditionIds to be an explicit nil

### UnsetConditionIds
`func (o *CroTrialsRow) UnsetConditionIds()`

UnsetConditionIds ensures that no value is present for ConditionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



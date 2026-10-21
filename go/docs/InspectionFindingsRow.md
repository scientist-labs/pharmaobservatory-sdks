# InspectionFindingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitationId** | Pointer to **NullableString** | The citation id. | [optional] 
**InspectionId** | Pointer to **NullableString** | The inspection the citation belongs to; join to the inspections dataset. | [optional] 
**Fei** | Pointer to **NullableString** | The FDA establishment identifier (FEI) of the inspected site. | [optional] 
**FirmName** | Pointer to **NullableString** | The firm&#39;s name as filed with the regulator. | [optional] 
**FirmOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] 
**FirmCompanyId** | Pointer to **NullableString** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] 
**FirmCompanyName** | Pointer to **NullableString** | Display name of the firm&#39;s company. | [optional] 
**FirmGroupId** | Pointer to **NullableString** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FirmGroupName** | Pointer to **NullableString** | Display name of the firm&#39;s group. | [optional] 
**IsOrgResolved** | Pointer to **NullableBool** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] 
**MatchTier** | Pointer to **NullableString** | How well the firm was attached to an organization: high, medium, low, registry, ambiguous or unresolved. | [optional] 
**InspectionDate** | Pointer to **NullableString** | The date of the inspection. | [optional] 
**Classification** | Pointer to **NullableString** | The outcome of the inspection: NAI, VAI or OAI. | [optional] 
**ActCfrNumber** | Pointer to **NullableString** | The cited regulation (a Code of Federal Regulations section or Act citation). | [optional] 
**ShortDescription** | Pointer to **NullableString** | The FDA&#39;s short description of the observation. | [optional] 
**ProgramArea** | Pointer to **NullableString** | The FDA program area of the citation. | [optional] 

## Methods

### NewInspectionFindingsRow

`func NewInspectionFindingsRow() *InspectionFindingsRow`

NewInspectionFindingsRow instantiates a new InspectionFindingsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInspectionFindingsRowWithDefaults

`func NewInspectionFindingsRowWithDefaults() *InspectionFindingsRow`

NewInspectionFindingsRowWithDefaults instantiates a new InspectionFindingsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitationId

`func (o *InspectionFindingsRow) GetCitationId() string`

GetCitationId returns the CitationId field if non-nil, zero value otherwise.

### GetCitationIdOk

`func (o *InspectionFindingsRow) GetCitationIdOk() (*string, bool)`

GetCitationIdOk returns a tuple with the CitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationId

`func (o *InspectionFindingsRow) SetCitationId(v string)`

SetCitationId sets CitationId field to given value.

### HasCitationId

`func (o *InspectionFindingsRow) HasCitationId() bool`

HasCitationId returns a boolean if a field has been set.

### SetCitationIdNil

`func (o *InspectionFindingsRow) SetCitationIdNil(b bool)`

 SetCitationIdNil sets the value for CitationId to be an explicit nil

### UnsetCitationId
`func (o *InspectionFindingsRow) UnsetCitationId()`

UnsetCitationId ensures that no value is present for CitationId, not even an explicit nil
### GetInspectionId

`func (o *InspectionFindingsRow) GetInspectionId() string`

GetInspectionId returns the InspectionId field if non-nil, zero value otherwise.

### GetInspectionIdOk

`func (o *InspectionFindingsRow) GetInspectionIdOk() (*string, bool)`

GetInspectionIdOk returns a tuple with the InspectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionId

`func (o *InspectionFindingsRow) SetInspectionId(v string)`

SetInspectionId sets InspectionId field to given value.

### HasInspectionId

`func (o *InspectionFindingsRow) HasInspectionId() bool`

HasInspectionId returns a boolean if a field has been set.

### SetInspectionIdNil

`func (o *InspectionFindingsRow) SetInspectionIdNil(b bool)`

 SetInspectionIdNil sets the value for InspectionId to be an explicit nil

### UnsetInspectionId
`func (o *InspectionFindingsRow) UnsetInspectionId()`

UnsetInspectionId ensures that no value is present for InspectionId, not even an explicit nil
### GetFei

`func (o *InspectionFindingsRow) GetFei() string`

GetFei returns the Fei field if non-nil, zero value otherwise.

### GetFeiOk

`func (o *InspectionFindingsRow) GetFeiOk() (*string, bool)`

GetFeiOk returns a tuple with the Fei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFei

`func (o *InspectionFindingsRow) SetFei(v string)`

SetFei sets Fei field to given value.

### HasFei

`func (o *InspectionFindingsRow) HasFei() bool`

HasFei returns a boolean if a field has been set.

### SetFeiNil

`func (o *InspectionFindingsRow) SetFeiNil(b bool)`

 SetFeiNil sets the value for Fei to be an explicit nil

### UnsetFei
`func (o *InspectionFindingsRow) UnsetFei()`

UnsetFei ensures that no value is present for Fei, not even an explicit nil
### GetFirmName

`func (o *InspectionFindingsRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *InspectionFindingsRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *InspectionFindingsRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *InspectionFindingsRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *InspectionFindingsRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *InspectionFindingsRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *InspectionFindingsRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *InspectionFindingsRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *InspectionFindingsRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *InspectionFindingsRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *InspectionFindingsRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *InspectionFindingsRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *InspectionFindingsRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *InspectionFindingsRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *InspectionFindingsRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *InspectionFindingsRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *InspectionFindingsRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *InspectionFindingsRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *InspectionFindingsRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *InspectionFindingsRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *InspectionFindingsRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *InspectionFindingsRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *InspectionFindingsRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *InspectionFindingsRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *InspectionFindingsRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *InspectionFindingsRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *InspectionFindingsRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *InspectionFindingsRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *InspectionFindingsRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *InspectionFindingsRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *InspectionFindingsRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *InspectionFindingsRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *InspectionFindingsRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *InspectionFindingsRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *InspectionFindingsRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *InspectionFindingsRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetIsOrgResolved

`func (o *InspectionFindingsRow) GetIsOrgResolved() bool`

GetIsOrgResolved returns the IsOrgResolved field if non-nil, zero value otherwise.

### GetIsOrgResolvedOk

`func (o *InspectionFindingsRow) GetIsOrgResolvedOk() (*bool, bool)`

GetIsOrgResolvedOk returns a tuple with the IsOrgResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgResolved

`func (o *InspectionFindingsRow) SetIsOrgResolved(v bool)`

SetIsOrgResolved sets IsOrgResolved field to given value.

### HasIsOrgResolved

`func (o *InspectionFindingsRow) HasIsOrgResolved() bool`

HasIsOrgResolved returns a boolean if a field has been set.

### SetIsOrgResolvedNil

`func (o *InspectionFindingsRow) SetIsOrgResolvedNil(b bool)`

 SetIsOrgResolvedNil sets the value for IsOrgResolved to be an explicit nil

### UnsetIsOrgResolved
`func (o *InspectionFindingsRow) UnsetIsOrgResolved()`

UnsetIsOrgResolved ensures that no value is present for IsOrgResolved, not even an explicit nil
### GetMatchTier

`func (o *InspectionFindingsRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *InspectionFindingsRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *InspectionFindingsRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *InspectionFindingsRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *InspectionFindingsRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *InspectionFindingsRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetInspectionDate

`func (o *InspectionFindingsRow) GetInspectionDate() string`

GetInspectionDate returns the InspectionDate field if non-nil, zero value otherwise.

### GetInspectionDateOk

`func (o *InspectionFindingsRow) GetInspectionDateOk() (*string, bool)`

GetInspectionDateOk returns a tuple with the InspectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionDate

`func (o *InspectionFindingsRow) SetInspectionDate(v string)`

SetInspectionDate sets InspectionDate field to given value.

### HasInspectionDate

`func (o *InspectionFindingsRow) HasInspectionDate() bool`

HasInspectionDate returns a boolean if a field has been set.

### SetInspectionDateNil

`func (o *InspectionFindingsRow) SetInspectionDateNil(b bool)`

 SetInspectionDateNil sets the value for InspectionDate to be an explicit nil

### UnsetInspectionDate
`func (o *InspectionFindingsRow) UnsetInspectionDate()`

UnsetInspectionDate ensures that no value is present for InspectionDate, not even an explicit nil
### GetClassification

`func (o *InspectionFindingsRow) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *InspectionFindingsRow) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *InspectionFindingsRow) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *InspectionFindingsRow) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *InspectionFindingsRow) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *InspectionFindingsRow) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetActCfrNumber

`func (o *InspectionFindingsRow) GetActCfrNumber() string`

GetActCfrNumber returns the ActCfrNumber field if non-nil, zero value otherwise.

### GetActCfrNumberOk

`func (o *InspectionFindingsRow) GetActCfrNumberOk() (*string, bool)`

GetActCfrNumberOk returns a tuple with the ActCfrNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActCfrNumber

`func (o *InspectionFindingsRow) SetActCfrNumber(v string)`

SetActCfrNumber sets ActCfrNumber field to given value.

### HasActCfrNumber

`func (o *InspectionFindingsRow) HasActCfrNumber() bool`

HasActCfrNumber returns a boolean if a field has been set.

### SetActCfrNumberNil

`func (o *InspectionFindingsRow) SetActCfrNumberNil(b bool)`

 SetActCfrNumberNil sets the value for ActCfrNumber to be an explicit nil

### UnsetActCfrNumber
`func (o *InspectionFindingsRow) UnsetActCfrNumber()`

UnsetActCfrNumber ensures that no value is present for ActCfrNumber, not even an explicit nil
### GetShortDescription

`func (o *InspectionFindingsRow) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *InspectionFindingsRow) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *InspectionFindingsRow) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *InspectionFindingsRow) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### SetShortDescriptionNil

`func (o *InspectionFindingsRow) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *InspectionFindingsRow) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
### GetProgramArea

`func (o *InspectionFindingsRow) GetProgramArea() string`

GetProgramArea returns the ProgramArea field if non-nil, zero value otherwise.

### GetProgramAreaOk

`func (o *InspectionFindingsRow) GetProgramAreaOk() (*string, bool)`

GetProgramAreaOk returns a tuple with the ProgramArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramArea

`func (o *InspectionFindingsRow) SetProgramArea(v string)`

SetProgramArea sets ProgramArea field to given value.

### HasProgramArea

`func (o *InspectionFindingsRow) HasProgramArea() bool`

HasProgramArea returns a boolean if a field has been set.

### SetProgramAreaNil

`func (o *InspectionFindingsRow) SetProgramAreaNil(b bool)`

 SetProgramAreaNil sets the value for ProgramArea to be an explicit nil

### UnsetProgramArea
`func (o *InspectionFindingsRow) UnsetProgramArea()`

UnsetProgramArea ensures that no value is present for ProgramArea, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



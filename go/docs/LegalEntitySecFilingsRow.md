# LegalEntitySecFilingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **NullableString** | Legal-entity level: the SEC filer organization record. | [optional] 
**OrgName** | Pointer to **NullableString** | Display name of the filer organization record. | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: the company the filer belongs to (null when unresolved). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the filer&#39;s company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the filer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the filer&#39;s group. | [optional] 
**EdgarCik** | Pointer to **NullableString** | The filer&#39;s 10-digit SEC Central Index Key (CIK). | [optional] 
**JurisdictionOfInc** | Pointer to **NullableString** | The jurisdiction of incorporation as the latest filing states it (a US state or a country, free text); null when blank. | [optional] 
**EntityType** | Pointer to **NullableString** | The entity type the latest filing states (Corporation, Limited Liability Company, Limited Partnership, General Partnership, Business Trust or Other); null when blank. | [optional] 
**EntityTypeOtherDesc** | Pointer to **NullableString** | The filer&#39;s free-text description when it ticked entity type Other. | [optional] 
**YearOfIncBand** | Pointer to **NullableString** | The filing&#39;s year-of-incorporation choice (Over Five Years Ago, Within Last Five Years or Yet to Be Formed); null when blank. | [optional] 
**YearOfInc** | Pointer to **NullableInt32** | The year of incorporation when the filing gives a four-digit year. | [optional] 
**IndustryGroupTypes** | Pointer to **[]string** | The industry groups the filer&#39;s in-scope filings name. | [optional] 
**PreviousNames** | Pointer to **[]string** | Earlier names the filer states, other than its current name. | [optional] 

## Methods

### NewLegalEntitySecFilingsRow

`func NewLegalEntitySecFilingsRow() *LegalEntitySecFilingsRow`

NewLegalEntitySecFilingsRow instantiates a new LegalEntitySecFilingsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalEntitySecFilingsRowWithDefaults

`func NewLegalEntitySecFilingsRowWithDefaults() *LegalEntitySecFilingsRow`

NewLegalEntitySecFilingsRowWithDefaults instantiates a new LegalEntitySecFilingsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *LegalEntitySecFilingsRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *LegalEntitySecFilingsRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *LegalEntitySecFilingsRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *LegalEntitySecFilingsRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *LegalEntitySecFilingsRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *LegalEntitySecFilingsRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *LegalEntitySecFilingsRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *LegalEntitySecFilingsRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *LegalEntitySecFilingsRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *LegalEntitySecFilingsRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *LegalEntitySecFilingsRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *LegalEntitySecFilingsRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetCompanyId

`func (o *LegalEntitySecFilingsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *LegalEntitySecFilingsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *LegalEntitySecFilingsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *LegalEntitySecFilingsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *LegalEntitySecFilingsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *LegalEntitySecFilingsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *LegalEntitySecFilingsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *LegalEntitySecFilingsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *LegalEntitySecFilingsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *LegalEntitySecFilingsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *LegalEntitySecFilingsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *LegalEntitySecFilingsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *LegalEntitySecFilingsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *LegalEntitySecFilingsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *LegalEntitySecFilingsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *LegalEntitySecFilingsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *LegalEntitySecFilingsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *LegalEntitySecFilingsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *LegalEntitySecFilingsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *LegalEntitySecFilingsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *LegalEntitySecFilingsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *LegalEntitySecFilingsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *LegalEntitySecFilingsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *LegalEntitySecFilingsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetEdgarCik

`func (o *LegalEntitySecFilingsRow) GetEdgarCik() string`

GetEdgarCik returns the EdgarCik field if non-nil, zero value otherwise.

### GetEdgarCikOk

`func (o *LegalEntitySecFilingsRow) GetEdgarCikOk() (*string, bool)`

GetEdgarCikOk returns a tuple with the EdgarCik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgarCik

`func (o *LegalEntitySecFilingsRow) SetEdgarCik(v string)`

SetEdgarCik sets EdgarCik field to given value.

### HasEdgarCik

`func (o *LegalEntitySecFilingsRow) HasEdgarCik() bool`

HasEdgarCik returns a boolean if a field has been set.

### SetEdgarCikNil

`func (o *LegalEntitySecFilingsRow) SetEdgarCikNil(b bool)`

 SetEdgarCikNil sets the value for EdgarCik to be an explicit nil

### UnsetEdgarCik
`func (o *LegalEntitySecFilingsRow) UnsetEdgarCik()`

UnsetEdgarCik ensures that no value is present for EdgarCik, not even an explicit nil
### GetJurisdictionOfInc

`func (o *LegalEntitySecFilingsRow) GetJurisdictionOfInc() string`

GetJurisdictionOfInc returns the JurisdictionOfInc field if non-nil, zero value otherwise.

### GetJurisdictionOfIncOk

`func (o *LegalEntitySecFilingsRow) GetJurisdictionOfIncOk() (*string, bool)`

GetJurisdictionOfIncOk returns a tuple with the JurisdictionOfInc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionOfInc

`func (o *LegalEntitySecFilingsRow) SetJurisdictionOfInc(v string)`

SetJurisdictionOfInc sets JurisdictionOfInc field to given value.

### HasJurisdictionOfInc

`func (o *LegalEntitySecFilingsRow) HasJurisdictionOfInc() bool`

HasJurisdictionOfInc returns a boolean if a field has been set.

### SetJurisdictionOfIncNil

`func (o *LegalEntitySecFilingsRow) SetJurisdictionOfIncNil(b bool)`

 SetJurisdictionOfIncNil sets the value for JurisdictionOfInc to be an explicit nil

### UnsetJurisdictionOfInc
`func (o *LegalEntitySecFilingsRow) UnsetJurisdictionOfInc()`

UnsetJurisdictionOfInc ensures that no value is present for JurisdictionOfInc, not even an explicit nil
### GetEntityType

`func (o *LegalEntitySecFilingsRow) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *LegalEntitySecFilingsRow) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *LegalEntitySecFilingsRow) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *LegalEntitySecFilingsRow) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *LegalEntitySecFilingsRow) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *LegalEntitySecFilingsRow) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetEntityTypeOtherDesc

`func (o *LegalEntitySecFilingsRow) GetEntityTypeOtherDesc() string`

GetEntityTypeOtherDesc returns the EntityTypeOtherDesc field if non-nil, zero value otherwise.

### GetEntityTypeOtherDescOk

`func (o *LegalEntitySecFilingsRow) GetEntityTypeOtherDescOk() (*string, bool)`

GetEntityTypeOtherDescOk returns a tuple with the EntityTypeOtherDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeOtherDesc

`func (o *LegalEntitySecFilingsRow) SetEntityTypeOtherDesc(v string)`

SetEntityTypeOtherDesc sets EntityTypeOtherDesc field to given value.

### HasEntityTypeOtherDesc

`func (o *LegalEntitySecFilingsRow) HasEntityTypeOtherDesc() bool`

HasEntityTypeOtherDesc returns a boolean if a field has been set.

### SetEntityTypeOtherDescNil

`func (o *LegalEntitySecFilingsRow) SetEntityTypeOtherDescNil(b bool)`

 SetEntityTypeOtherDescNil sets the value for EntityTypeOtherDesc to be an explicit nil

### UnsetEntityTypeOtherDesc
`func (o *LegalEntitySecFilingsRow) UnsetEntityTypeOtherDesc()`

UnsetEntityTypeOtherDesc ensures that no value is present for EntityTypeOtherDesc, not even an explicit nil
### GetYearOfIncBand

`func (o *LegalEntitySecFilingsRow) GetYearOfIncBand() string`

GetYearOfIncBand returns the YearOfIncBand field if non-nil, zero value otherwise.

### GetYearOfIncBandOk

`func (o *LegalEntitySecFilingsRow) GetYearOfIncBandOk() (*string, bool)`

GetYearOfIncBandOk returns a tuple with the YearOfIncBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOfIncBand

`func (o *LegalEntitySecFilingsRow) SetYearOfIncBand(v string)`

SetYearOfIncBand sets YearOfIncBand field to given value.

### HasYearOfIncBand

`func (o *LegalEntitySecFilingsRow) HasYearOfIncBand() bool`

HasYearOfIncBand returns a boolean if a field has been set.

### SetYearOfIncBandNil

`func (o *LegalEntitySecFilingsRow) SetYearOfIncBandNil(b bool)`

 SetYearOfIncBandNil sets the value for YearOfIncBand to be an explicit nil

### UnsetYearOfIncBand
`func (o *LegalEntitySecFilingsRow) UnsetYearOfIncBand()`

UnsetYearOfIncBand ensures that no value is present for YearOfIncBand, not even an explicit nil
### GetYearOfInc

`func (o *LegalEntitySecFilingsRow) GetYearOfInc() int32`

GetYearOfInc returns the YearOfInc field if non-nil, zero value otherwise.

### GetYearOfIncOk

`func (o *LegalEntitySecFilingsRow) GetYearOfIncOk() (*int32, bool)`

GetYearOfIncOk returns a tuple with the YearOfInc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOfInc

`func (o *LegalEntitySecFilingsRow) SetYearOfInc(v int32)`

SetYearOfInc sets YearOfInc field to given value.

### HasYearOfInc

`func (o *LegalEntitySecFilingsRow) HasYearOfInc() bool`

HasYearOfInc returns a boolean if a field has been set.

### SetYearOfIncNil

`func (o *LegalEntitySecFilingsRow) SetYearOfIncNil(b bool)`

 SetYearOfIncNil sets the value for YearOfInc to be an explicit nil

### UnsetYearOfInc
`func (o *LegalEntitySecFilingsRow) UnsetYearOfInc()`

UnsetYearOfInc ensures that no value is present for YearOfInc, not even an explicit nil
### GetIndustryGroupTypes

`func (o *LegalEntitySecFilingsRow) GetIndustryGroupTypes() []string`

GetIndustryGroupTypes returns the IndustryGroupTypes field if non-nil, zero value otherwise.

### GetIndustryGroupTypesOk

`func (o *LegalEntitySecFilingsRow) GetIndustryGroupTypesOk() (*[]string, bool)`

GetIndustryGroupTypesOk returns a tuple with the IndustryGroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryGroupTypes

`func (o *LegalEntitySecFilingsRow) SetIndustryGroupTypes(v []string)`

SetIndustryGroupTypes sets IndustryGroupTypes field to given value.

### HasIndustryGroupTypes

`func (o *LegalEntitySecFilingsRow) HasIndustryGroupTypes() bool`

HasIndustryGroupTypes returns a boolean if a field has been set.

### SetIndustryGroupTypesNil

`func (o *LegalEntitySecFilingsRow) SetIndustryGroupTypesNil(b bool)`

 SetIndustryGroupTypesNil sets the value for IndustryGroupTypes to be an explicit nil

### UnsetIndustryGroupTypes
`func (o *LegalEntitySecFilingsRow) UnsetIndustryGroupTypes()`

UnsetIndustryGroupTypes ensures that no value is present for IndustryGroupTypes, not even an explicit nil
### GetPreviousNames

`func (o *LegalEntitySecFilingsRow) GetPreviousNames() []string`

GetPreviousNames returns the PreviousNames field if non-nil, zero value otherwise.

### GetPreviousNamesOk

`func (o *LegalEntitySecFilingsRow) GetPreviousNamesOk() (*[]string, bool)`

GetPreviousNamesOk returns a tuple with the PreviousNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousNames

`func (o *LegalEntitySecFilingsRow) SetPreviousNames(v []string)`

SetPreviousNames sets PreviousNames field to given value.

### HasPreviousNames

`func (o *LegalEntitySecFilingsRow) HasPreviousNames() bool`

HasPreviousNames returns a boolean if a field has been set.

### SetPreviousNamesNil

`func (o *LegalEntitySecFilingsRow) SetPreviousNamesNil(b bool)`

 SetPreviousNamesNil sets the value for PreviousNames to be an explicit nil

### UnsetPreviousNames
`func (o *LegalEntitySecFilingsRow) UnsetPreviousNames()`

UnsetPreviousNames ensures that no value is present for PreviousNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



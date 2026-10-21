# InspectionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InspectionId** | Pointer to **NullableString** | The inspection id. | [optional] 
**Fei** | Pointer to **NullableString** | The FDA establishment identifier (FEI) of the inspected site. | [optional] 
**FirmName** | Pointer to **NullableString** | The firm&#39;s name as filed with the regulator (the inspected establishment&#39;s firm). | [optional] 
**FirmOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] 
**FirmCompanyId** | Pointer to **NullableString** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] 
**FirmCompanyName** | Pointer to **NullableString** | Display name of the firm&#39;s company. | [optional] 
**FirmGroupId** | Pointer to **NullableString** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FirmGroupName** | Pointer to **NullableString** | Display name of the firm&#39;s group. | [optional] 
**IsOrgResolved** | Pointer to **NullableBool** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] 
**MatchTier** | Pointer to **NullableString** | How well the firm was attached to an organization: high, medium, low, registry, ambiguous or unresolved. | [optional] 
**OrgResolutionMethod** | Pointer to **NullableString** | The route that attached the firm: fei_bridge (the FDA establishment identifier, FEI, bridge), fei_firm_mint or fei_firm_fold (the operating firm&#39;s registry record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional] 
**InspectionDate** | Pointer to **NullableString** | The date the inspection ended. | [optional] 
**FiscalYear** | Pointer to **NullableString** | The federal fiscal year of the inspection. | [optional] 
**Classification** | Pointer to **NullableString** | The outcome: NAI (no action indicated), VAI (voluntary action indicated) or OAI (official action indicated). | [optional] 
**Had483** | Pointer to **NullableBool** | True when the inspection ended with a Form 483 of observations. | [optional] 
**ProductType** | Pointer to **NullableString** | The product type inspected (Drugs, Devices, Biologics, Veterinary, Tobacco or Food/Cosmetics). | [optional] 
**ProjectArea** | Pointer to **NullableString** | The FDA project area of the inspection. | [optional] 
**City** | Pointer to **NullableString** | The establishment&#39;s city. | [optional] 
**State** | Pointer to **NullableString** | The establishment&#39;s state or province. | [optional] 
**PostalCode** | Pointer to **NullableString** | The establishment&#39;s postal code. | [optional] 
**CountryName** | Pointer to **NullableString** | The establishment&#39;s country name. | [optional] 
**CountryIso2** | Pointer to **NullableString** | The establishment&#39;s two-letter country code. | [optional] 

## Methods

### NewInspectionsRow

`func NewInspectionsRow() *InspectionsRow`

NewInspectionsRow instantiates a new InspectionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInspectionsRowWithDefaults

`func NewInspectionsRowWithDefaults() *InspectionsRow`

NewInspectionsRowWithDefaults instantiates a new InspectionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInspectionId

`func (o *InspectionsRow) GetInspectionId() string`

GetInspectionId returns the InspectionId field if non-nil, zero value otherwise.

### GetInspectionIdOk

`func (o *InspectionsRow) GetInspectionIdOk() (*string, bool)`

GetInspectionIdOk returns a tuple with the InspectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionId

`func (o *InspectionsRow) SetInspectionId(v string)`

SetInspectionId sets InspectionId field to given value.

### HasInspectionId

`func (o *InspectionsRow) HasInspectionId() bool`

HasInspectionId returns a boolean if a field has been set.

### SetInspectionIdNil

`func (o *InspectionsRow) SetInspectionIdNil(b bool)`

 SetInspectionIdNil sets the value for InspectionId to be an explicit nil

### UnsetInspectionId
`func (o *InspectionsRow) UnsetInspectionId()`

UnsetInspectionId ensures that no value is present for InspectionId, not even an explicit nil
### GetFei

`func (o *InspectionsRow) GetFei() string`

GetFei returns the Fei field if non-nil, zero value otherwise.

### GetFeiOk

`func (o *InspectionsRow) GetFeiOk() (*string, bool)`

GetFeiOk returns a tuple with the Fei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFei

`func (o *InspectionsRow) SetFei(v string)`

SetFei sets Fei field to given value.

### HasFei

`func (o *InspectionsRow) HasFei() bool`

HasFei returns a boolean if a field has been set.

### SetFeiNil

`func (o *InspectionsRow) SetFeiNil(b bool)`

 SetFeiNil sets the value for Fei to be an explicit nil

### UnsetFei
`func (o *InspectionsRow) UnsetFei()`

UnsetFei ensures that no value is present for Fei, not even an explicit nil
### GetFirmName

`func (o *InspectionsRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *InspectionsRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *InspectionsRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *InspectionsRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *InspectionsRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *InspectionsRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *InspectionsRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *InspectionsRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *InspectionsRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *InspectionsRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *InspectionsRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *InspectionsRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *InspectionsRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *InspectionsRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *InspectionsRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *InspectionsRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *InspectionsRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *InspectionsRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *InspectionsRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *InspectionsRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *InspectionsRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *InspectionsRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *InspectionsRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *InspectionsRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *InspectionsRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *InspectionsRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *InspectionsRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *InspectionsRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *InspectionsRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *InspectionsRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *InspectionsRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *InspectionsRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *InspectionsRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *InspectionsRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *InspectionsRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *InspectionsRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetIsOrgResolved

`func (o *InspectionsRow) GetIsOrgResolved() bool`

GetIsOrgResolved returns the IsOrgResolved field if non-nil, zero value otherwise.

### GetIsOrgResolvedOk

`func (o *InspectionsRow) GetIsOrgResolvedOk() (*bool, bool)`

GetIsOrgResolvedOk returns a tuple with the IsOrgResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgResolved

`func (o *InspectionsRow) SetIsOrgResolved(v bool)`

SetIsOrgResolved sets IsOrgResolved field to given value.

### HasIsOrgResolved

`func (o *InspectionsRow) HasIsOrgResolved() bool`

HasIsOrgResolved returns a boolean if a field has been set.

### SetIsOrgResolvedNil

`func (o *InspectionsRow) SetIsOrgResolvedNil(b bool)`

 SetIsOrgResolvedNil sets the value for IsOrgResolved to be an explicit nil

### UnsetIsOrgResolved
`func (o *InspectionsRow) UnsetIsOrgResolved()`

UnsetIsOrgResolved ensures that no value is present for IsOrgResolved, not even an explicit nil
### GetMatchTier

`func (o *InspectionsRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *InspectionsRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *InspectionsRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *InspectionsRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *InspectionsRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *InspectionsRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetOrgResolutionMethod

`func (o *InspectionsRow) GetOrgResolutionMethod() string`

GetOrgResolutionMethod returns the OrgResolutionMethod field if non-nil, zero value otherwise.

### GetOrgResolutionMethodOk

`func (o *InspectionsRow) GetOrgResolutionMethodOk() (*string, bool)`

GetOrgResolutionMethodOk returns a tuple with the OrgResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgResolutionMethod

`func (o *InspectionsRow) SetOrgResolutionMethod(v string)`

SetOrgResolutionMethod sets OrgResolutionMethod field to given value.

### HasOrgResolutionMethod

`func (o *InspectionsRow) HasOrgResolutionMethod() bool`

HasOrgResolutionMethod returns a boolean if a field has been set.

### SetOrgResolutionMethodNil

`func (o *InspectionsRow) SetOrgResolutionMethodNil(b bool)`

 SetOrgResolutionMethodNil sets the value for OrgResolutionMethod to be an explicit nil

### UnsetOrgResolutionMethod
`func (o *InspectionsRow) UnsetOrgResolutionMethod()`

UnsetOrgResolutionMethod ensures that no value is present for OrgResolutionMethod, not even an explicit nil
### GetInspectionDate

`func (o *InspectionsRow) GetInspectionDate() string`

GetInspectionDate returns the InspectionDate field if non-nil, zero value otherwise.

### GetInspectionDateOk

`func (o *InspectionsRow) GetInspectionDateOk() (*string, bool)`

GetInspectionDateOk returns a tuple with the InspectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionDate

`func (o *InspectionsRow) SetInspectionDate(v string)`

SetInspectionDate sets InspectionDate field to given value.

### HasInspectionDate

`func (o *InspectionsRow) HasInspectionDate() bool`

HasInspectionDate returns a boolean if a field has been set.

### SetInspectionDateNil

`func (o *InspectionsRow) SetInspectionDateNil(b bool)`

 SetInspectionDateNil sets the value for InspectionDate to be an explicit nil

### UnsetInspectionDate
`func (o *InspectionsRow) UnsetInspectionDate()`

UnsetInspectionDate ensures that no value is present for InspectionDate, not even an explicit nil
### GetFiscalYear

`func (o *InspectionsRow) GetFiscalYear() string`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *InspectionsRow) GetFiscalYearOk() (*string, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *InspectionsRow) SetFiscalYear(v string)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *InspectionsRow) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### SetFiscalYearNil

`func (o *InspectionsRow) SetFiscalYearNil(b bool)`

 SetFiscalYearNil sets the value for FiscalYear to be an explicit nil

### UnsetFiscalYear
`func (o *InspectionsRow) UnsetFiscalYear()`

UnsetFiscalYear ensures that no value is present for FiscalYear, not even an explicit nil
### GetClassification

`func (o *InspectionsRow) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *InspectionsRow) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *InspectionsRow) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *InspectionsRow) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *InspectionsRow) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *InspectionsRow) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetHad483

`func (o *InspectionsRow) GetHad483() bool`

GetHad483 returns the Had483 field if non-nil, zero value otherwise.

### GetHad483Ok

`func (o *InspectionsRow) GetHad483Ok() (*bool, bool)`

GetHad483Ok returns a tuple with the Had483 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHad483

`func (o *InspectionsRow) SetHad483(v bool)`

SetHad483 sets Had483 field to given value.

### HasHad483

`func (o *InspectionsRow) HasHad483() bool`

HasHad483 returns a boolean if a field has been set.

### SetHad483Nil

`func (o *InspectionsRow) SetHad483Nil(b bool)`

 SetHad483Nil sets the value for Had483 to be an explicit nil

### UnsetHad483
`func (o *InspectionsRow) UnsetHad483()`

UnsetHad483 ensures that no value is present for Had483, not even an explicit nil
### GetProductType

`func (o *InspectionsRow) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InspectionsRow) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InspectionsRow) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InspectionsRow) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### SetProductTypeNil

`func (o *InspectionsRow) SetProductTypeNil(b bool)`

 SetProductTypeNil sets the value for ProductType to be an explicit nil

### UnsetProductType
`func (o *InspectionsRow) UnsetProductType()`

UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
### GetProjectArea

`func (o *InspectionsRow) GetProjectArea() string`

GetProjectArea returns the ProjectArea field if non-nil, zero value otherwise.

### GetProjectAreaOk

`func (o *InspectionsRow) GetProjectAreaOk() (*string, bool)`

GetProjectAreaOk returns a tuple with the ProjectArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectArea

`func (o *InspectionsRow) SetProjectArea(v string)`

SetProjectArea sets ProjectArea field to given value.

### HasProjectArea

`func (o *InspectionsRow) HasProjectArea() bool`

HasProjectArea returns a boolean if a field has been set.

### SetProjectAreaNil

`func (o *InspectionsRow) SetProjectAreaNil(b bool)`

 SetProjectAreaNil sets the value for ProjectArea to be an explicit nil

### UnsetProjectArea
`func (o *InspectionsRow) UnsetProjectArea()`

UnsetProjectArea ensures that no value is present for ProjectArea, not even an explicit nil
### GetCity

`func (o *InspectionsRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *InspectionsRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *InspectionsRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *InspectionsRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *InspectionsRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *InspectionsRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *InspectionsRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InspectionsRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InspectionsRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InspectionsRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *InspectionsRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *InspectionsRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *InspectionsRow) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *InspectionsRow) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *InspectionsRow) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *InspectionsRow) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *InspectionsRow) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *InspectionsRow) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryName

`func (o *InspectionsRow) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *InspectionsRow) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *InspectionsRow) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *InspectionsRow) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### SetCountryNameNil

`func (o *InspectionsRow) SetCountryNameNil(b bool)`

 SetCountryNameNil sets the value for CountryName to be an explicit nil

### UnsetCountryName
`func (o *InspectionsRow) UnsetCountryName()`

UnsetCountryName ensures that no value is present for CountryName, not even an explicit nil
### GetCountryIso2

`func (o *InspectionsRow) GetCountryIso2() string`

GetCountryIso2 returns the CountryIso2 field if non-nil, zero value otherwise.

### GetCountryIso2Ok

`func (o *InspectionsRow) GetCountryIso2Ok() (*string, bool)`

GetCountryIso2Ok returns a tuple with the CountryIso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso2

`func (o *InspectionsRow) SetCountryIso2(v string)`

SetCountryIso2 sets CountryIso2 field to given value.

### HasCountryIso2

`func (o *InspectionsRow) HasCountryIso2() bool`

HasCountryIso2 returns a boolean if a field has been set.

### SetCountryIso2Nil

`func (o *InspectionsRow) SetCountryIso2Nil(b bool)`

 SetCountryIso2Nil sets the value for CountryIso2 to be an explicit nil

### UnsetCountryIso2
`func (o *InspectionsRow) UnsetCountryIso2()`

UnsetCountryIso2 ensures that no value is present for CountryIso2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



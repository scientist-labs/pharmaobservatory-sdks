# ImportRefusalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefusalId** | Pointer to **NullableString** | The refusal line id. | [optional] 
**Fei** | Pointer to **NullableString** | The FDA establishment identifier (FEI) of the refused firm. | [optional] 
**FirmName** | Pointer to **NullableString** | The firm&#39;s name as filed with the regulator (the firm named on the refusal). | [optional] 
**FirmOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] 
**FirmCompanyId** | Pointer to **NullableString** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] 
**FirmCompanyName** | Pointer to **NullableString** | Display name of the firm&#39;s company. | [optional] 
**FirmGroupId** | Pointer to **NullableString** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FirmGroupName** | Pointer to **NullableString** | Display name of the firm&#39;s group. | [optional] 
**IsOrgResolved** | Pointer to **NullableBool** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] 
**MatchTier** | Pointer to **NullableString** | How well the firm was attached to an organization: high, medium, low, registry, ambiguous or unresolved. | [optional] 
**OrgResolutionMethod** | Pointer to **NullableString** | The route that attached the firm: fei_bridge (the FDA establishment identifier, FEI, bridge), fei_firm_mint or fei_firm_fold (the operating firm&#39;s registry record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional] 
**RefusalDate** | Pointer to **NullableString** | The date the import was refused. | [optional] 
**ProductCode** | Pointer to **NullableString** | The FDA product code of the refused product. | [optional] 
**ProductDescription** | Pointer to **NullableString** | The refused product as described in the report. | [optional] 
**IndustryCode** | Pointer to **NullableString** | The FDA industry code of the product. | [optional] 
**RefusalCharges** | Pointer to **NullableString** | The charges (violations) the refusal cites. | [optional] 
**City** | Pointer to **NullableString** | The firm&#39;s city. | [optional] 
**State** | Pointer to **NullableString** | The firm&#39;s state or province. | [optional] 
**PostalCode** | Pointer to **NullableString** | The firm&#39;s postal code. | [optional] 
**CountryName** | Pointer to **NullableString** | The country the shipment came from. | [optional] 
**CountryCode** | Pointer to **NullableString** | The FDA country code of the shipment. | [optional] 
**CountryIso2** | Pointer to **NullableString** | The two-letter country code of the shipment. | [optional] 

## Methods

### NewImportRefusalsRow

`func NewImportRefusalsRow() *ImportRefusalsRow`

NewImportRefusalsRow instantiates a new ImportRefusalsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRefusalsRowWithDefaults

`func NewImportRefusalsRowWithDefaults() *ImportRefusalsRow`

NewImportRefusalsRowWithDefaults instantiates a new ImportRefusalsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefusalId

`func (o *ImportRefusalsRow) GetRefusalId() string`

GetRefusalId returns the RefusalId field if non-nil, zero value otherwise.

### GetRefusalIdOk

`func (o *ImportRefusalsRow) GetRefusalIdOk() (*string, bool)`

GetRefusalIdOk returns a tuple with the RefusalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalId

`func (o *ImportRefusalsRow) SetRefusalId(v string)`

SetRefusalId sets RefusalId field to given value.

### HasRefusalId

`func (o *ImportRefusalsRow) HasRefusalId() bool`

HasRefusalId returns a boolean if a field has been set.

### SetRefusalIdNil

`func (o *ImportRefusalsRow) SetRefusalIdNil(b bool)`

 SetRefusalIdNil sets the value for RefusalId to be an explicit nil

### UnsetRefusalId
`func (o *ImportRefusalsRow) UnsetRefusalId()`

UnsetRefusalId ensures that no value is present for RefusalId, not even an explicit nil
### GetFei

`func (o *ImportRefusalsRow) GetFei() string`

GetFei returns the Fei field if non-nil, zero value otherwise.

### GetFeiOk

`func (o *ImportRefusalsRow) GetFeiOk() (*string, bool)`

GetFeiOk returns a tuple with the Fei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFei

`func (o *ImportRefusalsRow) SetFei(v string)`

SetFei sets Fei field to given value.

### HasFei

`func (o *ImportRefusalsRow) HasFei() bool`

HasFei returns a boolean if a field has been set.

### SetFeiNil

`func (o *ImportRefusalsRow) SetFeiNil(b bool)`

 SetFeiNil sets the value for Fei to be an explicit nil

### UnsetFei
`func (o *ImportRefusalsRow) UnsetFei()`

UnsetFei ensures that no value is present for Fei, not even an explicit nil
### GetFirmName

`func (o *ImportRefusalsRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *ImportRefusalsRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *ImportRefusalsRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *ImportRefusalsRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *ImportRefusalsRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *ImportRefusalsRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *ImportRefusalsRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *ImportRefusalsRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *ImportRefusalsRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *ImportRefusalsRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *ImportRefusalsRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *ImportRefusalsRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *ImportRefusalsRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *ImportRefusalsRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *ImportRefusalsRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *ImportRefusalsRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *ImportRefusalsRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *ImportRefusalsRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *ImportRefusalsRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *ImportRefusalsRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *ImportRefusalsRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *ImportRefusalsRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *ImportRefusalsRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *ImportRefusalsRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *ImportRefusalsRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *ImportRefusalsRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *ImportRefusalsRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *ImportRefusalsRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *ImportRefusalsRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *ImportRefusalsRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *ImportRefusalsRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *ImportRefusalsRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *ImportRefusalsRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *ImportRefusalsRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *ImportRefusalsRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *ImportRefusalsRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetIsOrgResolved

`func (o *ImportRefusalsRow) GetIsOrgResolved() bool`

GetIsOrgResolved returns the IsOrgResolved field if non-nil, zero value otherwise.

### GetIsOrgResolvedOk

`func (o *ImportRefusalsRow) GetIsOrgResolvedOk() (*bool, bool)`

GetIsOrgResolvedOk returns a tuple with the IsOrgResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgResolved

`func (o *ImportRefusalsRow) SetIsOrgResolved(v bool)`

SetIsOrgResolved sets IsOrgResolved field to given value.

### HasIsOrgResolved

`func (o *ImportRefusalsRow) HasIsOrgResolved() bool`

HasIsOrgResolved returns a boolean if a field has been set.

### SetIsOrgResolvedNil

`func (o *ImportRefusalsRow) SetIsOrgResolvedNil(b bool)`

 SetIsOrgResolvedNil sets the value for IsOrgResolved to be an explicit nil

### UnsetIsOrgResolved
`func (o *ImportRefusalsRow) UnsetIsOrgResolved()`

UnsetIsOrgResolved ensures that no value is present for IsOrgResolved, not even an explicit nil
### GetMatchTier

`func (o *ImportRefusalsRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *ImportRefusalsRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *ImportRefusalsRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *ImportRefusalsRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *ImportRefusalsRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *ImportRefusalsRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetOrgResolutionMethod

`func (o *ImportRefusalsRow) GetOrgResolutionMethod() string`

GetOrgResolutionMethod returns the OrgResolutionMethod field if non-nil, zero value otherwise.

### GetOrgResolutionMethodOk

`func (o *ImportRefusalsRow) GetOrgResolutionMethodOk() (*string, bool)`

GetOrgResolutionMethodOk returns a tuple with the OrgResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgResolutionMethod

`func (o *ImportRefusalsRow) SetOrgResolutionMethod(v string)`

SetOrgResolutionMethod sets OrgResolutionMethod field to given value.

### HasOrgResolutionMethod

`func (o *ImportRefusalsRow) HasOrgResolutionMethod() bool`

HasOrgResolutionMethod returns a boolean if a field has been set.

### SetOrgResolutionMethodNil

`func (o *ImportRefusalsRow) SetOrgResolutionMethodNil(b bool)`

 SetOrgResolutionMethodNil sets the value for OrgResolutionMethod to be an explicit nil

### UnsetOrgResolutionMethod
`func (o *ImportRefusalsRow) UnsetOrgResolutionMethod()`

UnsetOrgResolutionMethod ensures that no value is present for OrgResolutionMethod, not even an explicit nil
### GetRefusalDate

`func (o *ImportRefusalsRow) GetRefusalDate() string`

GetRefusalDate returns the RefusalDate field if non-nil, zero value otherwise.

### GetRefusalDateOk

`func (o *ImportRefusalsRow) GetRefusalDateOk() (*string, bool)`

GetRefusalDateOk returns a tuple with the RefusalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalDate

`func (o *ImportRefusalsRow) SetRefusalDate(v string)`

SetRefusalDate sets RefusalDate field to given value.

### HasRefusalDate

`func (o *ImportRefusalsRow) HasRefusalDate() bool`

HasRefusalDate returns a boolean if a field has been set.

### SetRefusalDateNil

`func (o *ImportRefusalsRow) SetRefusalDateNil(b bool)`

 SetRefusalDateNil sets the value for RefusalDate to be an explicit nil

### UnsetRefusalDate
`func (o *ImportRefusalsRow) UnsetRefusalDate()`

UnsetRefusalDate ensures that no value is present for RefusalDate, not even an explicit nil
### GetProductCode

`func (o *ImportRefusalsRow) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ImportRefusalsRow) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ImportRefusalsRow) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *ImportRefusalsRow) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCodeNil

`func (o *ImportRefusalsRow) SetProductCodeNil(b bool)`

 SetProductCodeNil sets the value for ProductCode to be an explicit nil

### UnsetProductCode
`func (o *ImportRefusalsRow) UnsetProductCode()`

UnsetProductCode ensures that no value is present for ProductCode, not even an explicit nil
### GetProductDescription

`func (o *ImportRefusalsRow) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *ImportRefusalsRow) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *ImportRefusalsRow) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *ImportRefusalsRow) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### SetProductDescriptionNil

`func (o *ImportRefusalsRow) SetProductDescriptionNil(b bool)`

 SetProductDescriptionNil sets the value for ProductDescription to be an explicit nil

### UnsetProductDescription
`func (o *ImportRefusalsRow) UnsetProductDescription()`

UnsetProductDescription ensures that no value is present for ProductDescription, not even an explicit nil
### GetIndustryCode

`func (o *ImportRefusalsRow) GetIndustryCode() string`

GetIndustryCode returns the IndustryCode field if non-nil, zero value otherwise.

### GetIndustryCodeOk

`func (o *ImportRefusalsRow) GetIndustryCodeOk() (*string, bool)`

GetIndustryCodeOk returns a tuple with the IndustryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryCode

`func (o *ImportRefusalsRow) SetIndustryCode(v string)`

SetIndustryCode sets IndustryCode field to given value.

### HasIndustryCode

`func (o *ImportRefusalsRow) HasIndustryCode() bool`

HasIndustryCode returns a boolean if a field has been set.

### SetIndustryCodeNil

`func (o *ImportRefusalsRow) SetIndustryCodeNil(b bool)`

 SetIndustryCodeNil sets the value for IndustryCode to be an explicit nil

### UnsetIndustryCode
`func (o *ImportRefusalsRow) UnsetIndustryCode()`

UnsetIndustryCode ensures that no value is present for IndustryCode, not even an explicit nil
### GetRefusalCharges

`func (o *ImportRefusalsRow) GetRefusalCharges() string`

GetRefusalCharges returns the RefusalCharges field if non-nil, zero value otherwise.

### GetRefusalChargesOk

`func (o *ImportRefusalsRow) GetRefusalChargesOk() (*string, bool)`

GetRefusalChargesOk returns a tuple with the RefusalCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalCharges

`func (o *ImportRefusalsRow) SetRefusalCharges(v string)`

SetRefusalCharges sets RefusalCharges field to given value.

### HasRefusalCharges

`func (o *ImportRefusalsRow) HasRefusalCharges() bool`

HasRefusalCharges returns a boolean if a field has been set.

### SetRefusalChargesNil

`func (o *ImportRefusalsRow) SetRefusalChargesNil(b bool)`

 SetRefusalChargesNil sets the value for RefusalCharges to be an explicit nil

### UnsetRefusalCharges
`func (o *ImportRefusalsRow) UnsetRefusalCharges()`

UnsetRefusalCharges ensures that no value is present for RefusalCharges, not even an explicit nil
### GetCity

`func (o *ImportRefusalsRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ImportRefusalsRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ImportRefusalsRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ImportRefusalsRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ImportRefusalsRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ImportRefusalsRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *ImportRefusalsRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ImportRefusalsRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ImportRefusalsRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ImportRefusalsRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ImportRefusalsRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ImportRefusalsRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *ImportRefusalsRow) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ImportRefusalsRow) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ImportRefusalsRow) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ImportRefusalsRow) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ImportRefusalsRow) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ImportRefusalsRow) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryName

`func (o *ImportRefusalsRow) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *ImportRefusalsRow) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *ImportRefusalsRow) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *ImportRefusalsRow) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### SetCountryNameNil

`func (o *ImportRefusalsRow) SetCountryNameNil(b bool)`

 SetCountryNameNil sets the value for CountryName to be an explicit nil

### UnsetCountryName
`func (o *ImportRefusalsRow) UnsetCountryName()`

UnsetCountryName ensures that no value is present for CountryName, not even an explicit nil
### GetCountryCode

`func (o *ImportRefusalsRow) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ImportRefusalsRow) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ImportRefusalsRow) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ImportRefusalsRow) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *ImportRefusalsRow) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *ImportRefusalsRow) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetCountryIso2

`func (o *ImportRefusalsRow) GetCountryIso2() string`

GetCountryIso2 returns the CountryIso2 field if non-nil, zero value otherwise.

### GetCountryIso2Ok

`func (o *ImportRefusalsRow) GetCountryIso2Ok() (*string, bool)`

GetCountryIso2Ok returns a tuple with the CountryIso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso2

`func (o *ImportRefusalsRow) SetCountryIso2(v string)`

SetCountryIso2 sets CountryIso2 field to given value.

### HasCountryIso2

`func (o *ImportRefusalsRow) HasCountryIso2() bool`

HasCountryIso2 returns a boolean if a field has been set.

### SetCountryIso2Nil

`func (o *ImportRefusalsRow) SetCountryIso2Nil(b bool)`

 SetCountryIso2Nil sets the value for CountryIso2 to be an explicit nil

### UnsetCountryIso2
`func (o *ImportRefusalsRow) UnsetCountryIso2()`

UnsetCountryIso2 ensures that no value is present for CountryIso2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DrugSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry. | [optional] 
**SiteName** | Pointer to **NullableString** | The facility name as registered. | [optional] 
**City** | Pointer to **NullableString** | The facility&#39;s city. | [optional] 
**Country** | Pointer to **NullableString** | The facility&#39;s country. | [optional] 
**NTrials** | Pointer to **NullableInt32** | The drug&#39;s trials registered at the site. | [optional] 

## Methods

### NewDrugSitesRow

`func NewDrugSitesRow() *DrugSitesRow`

NewDrugSitesRow instantiates a new DrugSitesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugSitesRowWithDefaults

`func NewDrugSitesRowWithDefaults() *DrugSitesRow`

NewDrugSitesRowWithDefaults instantiates a new DrugSitesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugSitesRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugSitesRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugSitesRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugSitesRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugSitesRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugSitesRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetSiteId

`func (o *DrugSitesRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DrugSitesRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DrugSitesRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DrugSitesRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *DrugSitesRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *DrugSitesRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetSiteName

`func (o *DrugSitesRow) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *DrugSitesRow) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *DrugSitesRow) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *DrugSitesRow) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *DrugSitesRow) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *DrugSitesRow) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetCity

`func (o *DrugSitesRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DrugSitesRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DrugSitesRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *DrugSitesRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *DrugSitesRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *DrugSitesRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *DrugSitesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DrugSitesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DrugSitesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DrugSitesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *DrugSitesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *DrugSitesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetNTrials

`func (o *DrugSitesRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *DrugSitesRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *DrugSitesRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *DrugSitesRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *DrugSitesRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *DrugSitesRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



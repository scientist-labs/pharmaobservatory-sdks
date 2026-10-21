# SupplierSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **NullableString** | The supplier site id. | [optional] 
**ProviderId** | Pointer to **NullableString** | The supplier id: the Scientist.com provider id as text. | [optional] 
**SupplierName** | Pointer to **NullableString** | The supplier&#39;s name as listed on the marketplace. | [optional] 
**Category** | Pointer to **NullableString** | The kind of site as the supplier recorded it (Headquarters, Laboratory, Warehouse, Other). Examples, not a closed set. | [optional] 
**City** | Pointer to **NullableString** | City as recorded. | [optional] 
**State** | Pointer to **NullableString** | State or province as recorded. | [optional] 
**Country** | Pointer to **NullableString** | Country as recorded. | [optional] 
**Latitude** | Pointer to **NullableFloat32** | Latitude of the site. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | Longitude of the site. | [optional] 
**CompanyType** | Pointer to **NullableString** | The supplier&#39;s company type, repeated on each of its sites. | [optional] 
**ServiceAreas** | Pointer to **[]string** | The supplier&#39;s service areas, repeated on each of its sites. Filter with &#39;Biology&#39; &#x3D; ANY(service_areas). | [optional] 
**HqCountry** | Pointer to **NullableString** | ISO 3166-1 alpha-2 country code of the supplier&#39;s headquarters point, repeated on each of its sites. | [optional] 
**HqRegion** | Pointer to **NullableString** | The broad world region of the supplier&#39;s headquarters country, repeated on each of its sites; other when unknown. | [optional] 

## Methods

### NewSupplierSitesRow

`func NewSupplierSitesRow() *SupplierSitesRow`

NewSupplierSitesRow instantiates a new SupplierSitesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierSitesRowWithDefaults

`func NewSupplierSitesRowWithDefaults() *SupplierSitesRow`

NewSupplierSitesRowWithDefaults instantiates a new SupplierSitesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *SupplierSitesRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SupplierSitesRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SupplierSitesRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SupplierSitesRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *SupplierSitesRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *SupplierSitesRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetProviderId

`func (o *SupplierSitesRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *SupplierSitesRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *SupplierSitesRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *SupplierSitesRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *SupplierSitesRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *SupplierSitesRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetSupplierName

`func (o *SupplierSitesRow) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *SupplierSitesRow) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *SupplierSitesRow) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.

### HasSupplierName

`func (o *SupplierSitesRow) HasSupplierName() bool`

HasSupplierName returns a boolean if a field has been set.

### SetSupplierNameNil

`func (o *SupplierSitesRow) SetSupplierNameNil(b bool)`

 SetSupplierNameNil sets the value for SupplierName to be an explicit nil

### UnsetSupplierName
`func (o *SupplierSitesRow) UnsetSupplierName()`

UnsetSupplierName ensures that no value is present for SupplierName, not even an explicit nil
### GetCategory

`func (o *SupplierSitesRow) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SupplierSitesRow) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SupplierSitesRow) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SupplierSitesRow) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *SupplierSitesRow) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *SupplierSitesRow) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCity

`func (o *SupplierSitesRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SupplierSitesRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SupplierSitesRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SupplierSitesRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *SupplierSitesRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *SupplierSitesRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *SupplierSitesRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SupplierSitesRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SupplierSitesRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SupplierSitesRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *SupplierSitesRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *SupplierSitesRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *SupplierSitesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SupplierSitesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SupplierSitesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SupplierSitesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *SupplierSitesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *SupplierSitesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLatitude

`func (o *SupplierSitesRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SupplierSitesRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SupplierSitesRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *SupplierSitesRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *SupplierSitesRow) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *SupplierSitesRow) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *SupplierSitesRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SupplierSitesRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SupplierSitesRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *SupplierSitesRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *SupplierSitesRow) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *SupplierSitesRow) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetCompanyType

`func (o *SupplierSitesRow) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *SupplierSitesRow) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *SupplierSitesRow) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *SupplierSitesRow) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### SetCompanyTypeNil

`func (o *SupplierSitesRow) SetCompanyTypeNil(b bool)`

 SetCompanyTypeNil sets the value for CompanyType to be an explicit nil

### UnsetCompanyType
`func (o *SupplierSitesRow) UnsetCompanyType()`

UnsetCompanyType ensures that no value is present for CompanyType, not even an explicit nil
### GetServiceAreas

`func (o *SupplierSitesRow) GetServiceAreas() []string`

GetServiceAreas returns the ServiceAreas field if non-nil, zero value otherwise.

### GetServiceAreasOk

`func (o *SupplierSitesRow) GetServiceAreasOk() (*[]string, bool)`

GetServiceAreasOk returns a tuple with the ServiceAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAreas

`func (o *SupplierSitesRow) SetServiceAreas(v []string)`

SetServiceAreas sets ServiceAreas field to given value.

### HasServiceAreas

`func (o *SupplierSitesRow) HasServiceAreas() bool`

HasServiceAreas returns a boolean if a field has been set.

### SetServiceAreasNil

`func (o *SupplierSitesRow) SetServiceAreasNil(b bool)`

 SetServiceAreasNil sets the value for ServiceAreas to be an explicit nil

### UnsetServiceAreas
`func (o *SupplierSitesRow) UnsetServiceAreas()`

UnsetServiceAreas ensures that no value is present for ServiceAreas, not even an explicit nil
### GetHqCountry

`func (o *SupplierSitesRow) GetHqCountry() string`

GetHqCountry returns the HqCountry field if non-nil, zero value otherwise.

### GetHqCountryOk

`func (o *SupplierSitesRow) GetHqCountryOk() (*string, bool)`

GetHqCountryOk returns a tuple with the HqCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqCountry

`func (o *SupplierSitesRow) SetHqCountry(v string)`

SetHqCountry sets HqCountry field to given value.

### HasHqCountry

`func (o *SupplierSitesRow) HasHqCountry() bool`

HasHqCountry returns a boolean if a field has been set.

### SetHqCountryNil

`func (o *SupplierSitesRow) SetHqCountryNil(b bool)`

 SetHqCountryNil sets the value for HqCountry to be an explicit nil

### UnsetHqCountry
`func (o *SupplierSitesRow) UnsetHqCountry()`

UnsetHqCountry ensures that no value is present for HqCountry, not even an explicit nil
### GetHqRegion

`func (o *SupplierSitesRow) GetHqRegion() string`

GetHqRegion returns the HqRegion field if non-nil, zero value otherwise.

### GetHqRegionOk

`func (o *SupplierSitesRow) GetHqRegionOk() (*string, bool)`

GetHqRegionOk returns a tuple with the HqRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqRegion

`func (o *SupplierSitesRow) SetHqRegion(v string)`

SetHqRegion sets HqRegion field to given value.

### HasHqRegion

`func (o *SupplierSitesRow) HasHqRegion() bool`

HasHqRegion returns a boolean if a field has been set.

### SetHqRegionNil

`func (o *SupplierSitesRow) SetHqRegionNil(b bool)`

 SetHqRegionNil sets the value for HqRegion to be an explicit nil

### UnsetHqRegion
`func (o *SupplierSitesRow) UnsetHqRegion()`

UnsetHqRegion ensures that no value is present for HqRegion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



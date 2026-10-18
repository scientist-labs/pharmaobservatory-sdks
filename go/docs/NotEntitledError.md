# NotEntitledError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Visible** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Tier** | **string** |  | 
**Attribution** | **string** |  | 
**Page** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Section** | Pointer to **string** |  | [optional] 
**Heading** | Pointer to **string** |  | [optional] 
**Lens** | Pointer to **NullableString** |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 
**BuildState** | Pointer to [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**Restriction** | Pointer to [**OrgsBenzingaCatalysts200ResponseRestriction**](OrgsBenzingaCatalysts200ResponseRestriction.md) |  | [optional] 

## Methods

### NewNotEntitledError

`func NewNotEntitledError(error_ string, tier string, attribution string, ) *NotEntitledError`

NewNotEntitledError instantiates a new NotEntitledError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotEntitledErrorWithDefaults

`func NewNotEntitledErrorWithDefaults() *NotEntitledError`

NewNotEntitledErrorWithDefaults instantiates a new NotEntitledError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *NotEntitledError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NotEntitledError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NotEntitledError) SetError(v string)`

SetError sets Error field to given value.


### GetVisible

`func (o *NotEntitledError) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *NotEntitledError) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *NotEntitledError) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *NotEntitledError) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetReason

`func (o *NotEntitledError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NotEntitledError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NotEntitledError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *NotEntitledError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTier

`func (o *NotEntitledError) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *NotEntitledError) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *NotEntitledError) SetTier(v string)`

SetTier sets Tier field to given value.


### GetAttribution

`func (o *NotEntitledError) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *NotEntitledError) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *NotEntitledError) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.


### GetPage

`func (o *NotEntitledError) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *NotEntitledError) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *NotEntitledError) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *NotEntitledError) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetId

`func (o *NotEntitledError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotEntitledError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotEntitledError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotEntitledError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *NotEntitledError) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *NotEntitledError) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *NotEntitledError) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *NotEntitledError) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetHeading

`func (o *NotEntitledError) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *NotEntitledError) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *NotEntitledError) SetHeading(v string)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *NotEntitledError) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### GetLens

`func (o *NotEntitledError) GetLens() string`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *NotEntitledError) GetLensOk() (*string, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *NotEntitledError) SetLens(v string)`

SetLens sets Lens field to given value.

### HasLens

`func (o *NotEntitledError) HasLens() bool`

HasLens returns a boolean if a field has been set.

### SetLensNil

`func (o *NotEntitledError) SetLensNil(b bool)`

 SetLensNil sets the value for Lens to be an explicit nil

### UnsetLens
`func (o *NotEntitledError) UnsetLens()`

UnsetLens ensures that no value is present for Lens, not even an explicit nil
### GetWebUrl

`func (o *NotEntitledError) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *NotEntitledError) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *NotEntitledError) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *NotEntitledError) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetBuildState

`func (o *NotEntitledError) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *NotEntitledError) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *NotEntitledError) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *NotEntitledError) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetLicense

`func (o *NotEntitledError) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *NotEntitledError) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *NotEntitledError) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *NotEntitledError) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetRestriction

`func (o *NotEntitledError) GetRestriction() OrgsBenzingaCatalysts200ResponseRestriction`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *NotEntitledError) GetRestrictionOk() (*OrgsBenzingaCatalysts200ResponseRestriction, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *NotEntitledError) SetRestriction(v OrgsBenzingaCatalysts200ResponseRestriction)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *NotEntitledError) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



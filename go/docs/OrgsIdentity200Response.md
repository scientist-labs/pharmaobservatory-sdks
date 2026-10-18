# OrgsIdentity200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **string** |  | 
**Id** | **string** |  | 
**Section** | **string** |  | 
**Heading** | **string** |  | 
**Lens** | **interface{}** |  | 
**WebUrl** | **string** |  | 
**Shape** | [**BiomarkersList200ResponseShape**](BiomarkersList200ResponseShape.md) |  | 
**Columns** | [**[]BiomarkersList200ResponseColumnsInner**](BiomarkersList200ResponseColumnsInner.md) |  | 
**Rows** | [**[]OrgsGet200ResponseIdentityRowsInner**](OrgsGet200ResponseIdentityRowsInner.md) |  | 
**Attribution** | [**[]BiomarkersList200ResponseAttributionInner**](BiomarkersList200ResponseAttributionInner.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 
**ResolutionConfidence** | [**BiomarkersGet200ResponseResolutionConfidence**](BiomarkersGet200ResponseResolutionConfidence.md) |  | 

## Methods

### NewOrgsIdentity200Response

`func NewOrgsIdentity200Response(page string, id string, section string, heading string, lens interface{}, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []OrgsGet200ResponseIdentityRowsInner, attribution []BiomarkersList200ResponseAttributionInner, buildState BiomarkersList200ResponseBuildState, resolutionConfidence BiomarkersGet200ResponseResolutionConfidence, ) *OrgsIdentity200Response`

NewOrgsIdentity200Response instantiates a new OrgsIdentity200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsIdentity200ResponseWithDefaults

`func NewOrgsIdentity200ResponseWithDefaults() *OrgsIdentity200Response`

NewOrgsIdentity200ResponseWithDefaults instantiates a new OrgsIdentity200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrgsIdentity200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgsIdentity200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgsIdentity200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *OrgsIdentity200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgsIdentity200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgsIdentity200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSection

`func (o *OrgsIdentity200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *OrgsIdentity200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *OrgsIdentity200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *OrgsIdentity200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *OrgsIdentity200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *OrgsIdentity200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *OrgsIdentity200Response) GetLens() interface{}`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *OrgsIdentity200Response) GetLensOk() (*interface{}, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *OrgsIdentity200Response) SetLens(v interface{})`

SetLens sets Lens field to given value.


### SetLensNil

`func (o *OrgsIdentity200Response) SetLensNil(b bool)`

 SetLensNil sets the value for Lens to be an explicit nil

### UnsetLens
`func (o *OrgsIdentity200Response) UnsetLens()`

UnsetLens ensures that no value is present for Lens, not even an explicit nil
### GetWebUrl

`func (o *OrgsIdentity200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OrgsIdentity200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OrgsIdentity200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *OrgsIdentity200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *OrgsIdentity200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *OrgsIdentity200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *OrgsIdentity200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *OrgsIdentity200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *OrgsIdentity200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *OrgsIdentity200Response) GetRows() []OrgsGet200ResponseIdentityRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *OrgsIdentity200Response) GetRowsOk() (*[]OrgsGet200ResponseIdentityRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *OrgsIdentity200Response) SetRows(v []OrgsGet200ResponseIdentityRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *OrgsIdentity200Response) GetAttribution() []BiomarkersList200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsIdentity200Response) GetAttributionOk() (*[]BiomarkersList200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsIdentity200Response) SetAttribution(v []BiomarkersList200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetBuildState

`func (o *OrgsIdentity200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *OrgsIdentity200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *OrgsIdentity200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetResolutionConfidence

`func (o *OrgsIdentity200Response) GetResolutionConfidence() BiomarkersGet200ResponseResolutionConfidence`

GetResolutionConfidence returns the ResolutionConfidence field if non-nil, zero value otherwise.

### GetResolutionConfidenceOk

`func (o *OrgsIdentity200Response) GetResolutionConfidenceOk() (*BiomarkersGet200ResponseResolutionConfidence, bool)`

GetResolutionConfidenceOk returns a tuple with the ResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionConfidence

`func (o *OrgsIdentity200Response) SetResolutionConfidence(v BiomarkersGet200ResponseResolutionConfidence)`

SetResolutionConfidence sets ResolutionConfidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



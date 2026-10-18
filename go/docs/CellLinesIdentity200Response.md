# CellLinesIdentity200Response

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
**Rows** | [**[]CellLinesGet200ResponseIdentityRowsInner**](CellLinesGet200ResponseIdentityRowsInner.md) |  | 
**Attribution** | [**[]CellLinesList200ResponseAttributionInner**](CellLinesList200ResponseAttributionInner.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 
**ResolutionConfidence** | [**BiomarkersGet200ResponseResolutionConfidence**](BiomarkersGet200ResponseResolutionConfidence.md) |  | 

## Methods

### NewCellLinesIdentity200Response

`func NewCellLinesIdentity200Response(page string, id string, section string, heading string, lens interface{}, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []CellLinesGet200ResponseIdentityRowsInner, attribution []CellLinesList200ResponseAttributionInner, buildState BiomarkersList200ResponseBuildState, resolutionConfidence BiomarkersGet200ResponseResolutionConfidence, ) *CellLinesIdentity200Response`

NewCellLinesIdentity200Response instantiates a new CellLinesIdentity200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellLinesIdentity200ResponseWithDefaults

`func NewCellLinesIdentity200ResponseWithDefaults() *CellLinesIdentity200Response`

NewCellLinesIdentity200ResponseWithDefaults instantiates a new CellLinesIdentity200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CellLinesIdentity200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CellLinesIdentity200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CellLinesIdentity200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *CellLinesIdentity200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CellLinesIdentity200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CellLinesIdentity200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSection

`func (o *CellLinesIdentity200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *CellLinesIdentity200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *CellLinesIdentity200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *CellLinesIdentity200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *CellLinesIdentity200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *CellLinesIdentity200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *CellLinesIdentity200Response) GetLens() interface{}`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *CellLinesIdentity200Response) GetLensOk() (*interface{}, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *CellLinesIdentity200Response) SetLens(v interface{})`

SetLens sets Lens field to given value.


### SetLensNil

`func (o *CellLinesIdentity200Response) SetLensNil(b bool)`

 SetLensNil sets the value for Lens to be an explicit nil

### UnsetLens
`func (o *CellLinesIdentity200Response) UnsetLens()`

UnsetLens ensures that no value is present for Lens, not even an explicit nil
### GetWebUrl

`func (o *CellLinesIdentity200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *CellLinesIdentity200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *CellLinesIdentity200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *CellLinesIdentity200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *CellLinesIdentity200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *CellLinesIdentity200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *CellLinesIdentity200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *CellLinesIdentity200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *CellLinesIdentity200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *CellLinesIdentity200Response) GetRows() []CellLinesGet200ResponseIdentityRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *CellLinesIdentity200Response) GetRowsOk() (*[]CellLinesGet200ResponseIdentityRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *CellLinesIdentity200Response) SetRows(v []CellLinesGet200ResponseIdentityRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *CellLinesIdentity200Response) GetAttribution() []CellLinesList200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *CellLinesIdentity200Response) GetAttributionOk() (*[]CellLinesList200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *CellLinesIdentity200Response) SetAttribution(v []CellLinesList200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetBuildState

`func (o *CellLinesIdentity200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *CellLinesIdentity200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *CellLinesIdentity200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetResolutionConfidence

`func (o *CellLinesIdentity200Response) GetResolutionConfidence() BiomarkersGet200ResponseResolutionConfidence`

GetResolutionConfidence returns the ResolutionConfidence field if non-nil, zero value otherwise.

### GetResolutionConfidenceOk

`func (o *CellLinesIdentity200Response) GetResolutionConfidenceOk() (*BiomarkersGet200ResponseResolutionConfidence, bool)`

GetResolutionConfidenceOk returns a tuple with the ResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionConfidence

`func (o *CellLinesIdentity200Response) SetResolutionConfidence(v BiomarkersGet200ResponseResolutionConfidence)`

SetResolutionConfidence sets ResolutionConfidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



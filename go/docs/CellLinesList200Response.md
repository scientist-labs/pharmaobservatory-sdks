# CellLinesList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **string** |  | 
**Id** | **interface{}** |  | 
**Section** | **string** |  | 
**Heading** | **string** |  | 
**Lens** | **interface{}** |  | 
**WebUrl** | **string** |  | 
**Shape** | [**BiomarkersList200ResponseShape**](BiomarkersList200ResponseShape.md) |  | 
**Columns** | [**[]BiomarkersList200ResponseColumnsInner**](BiomarkersList200ResponseColumnsInner.md) |  | 
**Rows** | [**[]CellLinesList200ResponseRowsInner**](CellLinesList200ResponseRowsInner.md) |  | 
**Attribution** | [**[]CellLinesList200ResponseAttributionInner**](CellLinesList200ResponseAttributionInner.md) |  | 
**Pagination** | [**BiomarkersList200ResponsePagination**](BiomarkersList200ResponsePagination.md) |  | 
**Query** | [**CellLinesList200ResponseQuery**](CellLinesList200ResponseQuery.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 
**Facets** | Pointer to [**[]BiomarkersList200ResponseFacetsInner**](BiomarkersList200ResponseFacetsInner.md) |  | [optional] 

## Methods

### NewCellLinesList200Response

`func NewCellLinesList200Response(page string, id interface{}, section string, heading string, lens interface{}, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []CellLinesList200ResponseRowsInner, attribution []CellLinesList200ResponseAttributionInner, pagination BiomarkersList200ResponsePagination, query CellLinesList200ResponseQuery, buildState BiomarkersList200ResponseBuildState, ) *CellLinesList200Response`

NewCellLinesList200Response instantiates a new CellLinesList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellLinesList200ResponseWithDefaults

`func NewCellLinesList200ResponseWithDefaults() *CellLinesList200Response`

NewCellLinesList200ResponseWithDefaults instantiates a new CellLinesList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CellLinesList200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CellLinesList200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CellLinesList200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *CellLinesList200Response) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CellLinesList200Response) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CellLinesList200Response) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *CellLinesList200Response) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CellLinesList200Response) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSection

`func (o *CellLinesList200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *CellLinesList200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *CellLinesList200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *CellLinesList200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *CellLinesList200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *CellLinesList200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *CellLinesList200Response) GetLens() interface{}`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *CellLinesList200Response) GetLensOk() (*interface{}, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *CellLinesList200Response) SetLens(v interface{})`

SetLens sets Lens field to given value.


### SetLensNil

`func (o *CellLinesList200Response) SetLensNil(b bool)`

 SetLensNil sets the value for Lens to be an explicit nil

### UnsetLens
`func (o *CellLinesList200Response) UnsetLens()`

UnsetLens ensures that no value is present for Lens, not even an explicit nil
### GetWebUrl

`func (o *CellLinesList200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *CellLinesList200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *CellLinesList200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *CellLinesList200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *CellLinesList200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *CellLinesList200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *CellLinesList200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *CellLinesList200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *CellLinesList200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *CellLinesList200Response) GetRows() []CellLinesList200ResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *CellLinesList200Response) GetRowsOk() (*[]CellLinesList200ResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *CellLinesList200Response) SetRows(v []CellLinesList200ResponseRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *CellLinesList200Response) GetAttribution() []CellLinesList200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *CellLinesList200Response) GetAttributionOk() (*[]CellLinesList200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *CellLinesList200Response) SetAttribution(v []CellLinesList200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetPagination

`func (o *CellLinesList200Response) GetPagination() BiomarkersList200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CellLinesList200Response) GetPaginationOk() (*BiomarkersList200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CellLinesList200Response) SetPagination(v BiomarkersList200ResponsePagination)`

SetPagination sets Pagination field to given value.


### GetQuery

`func (o *CellLinesList200Response) GetQuery() CellLinesList200ResponseQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CellLinesList200Response) GetQueryOk() (*CellLinesList200ResponseQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CellLinesList200Response) SetQuery(v CellLinesList200ResponseQuery)`

SetQuery sets Query field to given value.


### GetBuildState

`func (o *CellLinesList200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *CellLinesList200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *CellLinesList200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetFacets

`func (o *CellLinesList200Response) GetFacets() []BiomarkersList200ResponseFacetsInner`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *CellLinesList200Response) GetFacetsOk() (*[]BiomarkersList200ResponseFacetsInner, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *CellLinesList200Response) SetFacets(v []BiomarkersList200ResponseFacetsInner)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *CellLinesList200Response) HasFacets() bool`

HasFacets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



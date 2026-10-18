# OrgsOwnershipHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **string** |  | 
**Id** | **string** |  | 
**Section** | **string** |  | 
**Heading** | **string** |  | 
**Lens** | **string** |  | 
**WebUrl** | **string** |  | 
**Shape** | [**BiomarkersList200ResponseShape**](BiomarkersList200ResponseShape.md) |  | 
**Columns** | [**[]BiomarkersList200ResponseColumnsInner**](BiomarkersList200ResponseColumnsInner.md) |  | 
**Rows** | [**[]OrgsOwnershipHistory200ResponseRowsInner**](OrgsOwnershipHistory200ResponseRowsInner.md) |  | 
**Attribution** | [**[]OrgsCurrentOwnership200ResponseAttributionInner**](OrgsCurrentOwnership200ResponseAttributionInner.md) |  | 
**Pagination** | [**BiomarkersAssayProviderLedger200ResponsePagination**](BiomarkersAssayProviderLedger200ResponsePagination.md) |  | 
**Context** | [**OrgsOwnershipHistory200ResponseContext**](OrgsOwnershipHistory200ResponseContext.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 

## Methods

### NewOrgsOwnershipHistory200Response

`func NewOrgsOwnershipHistory200Response(page string, id string, section string, heading string, lens string, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []OrgsOwnershipHistory200ResponseRowsInner, attribution []OrgsCurrentOwnership200ResponseAttributionInner, pagination BiomarkersAssayProviderLedger200ResponsePagination, context OrgsOwnershipHistory200ResponseContext, buildState BiomarkersList200ResponseBuildState, ) *OrgsOwnershipHistory200Response`

NewOrgsOwnershipHistory200Response instantiates a new OrgsOwnershipHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsOwnershipHistory200ResponseWithDefaults

`func NewOrgsOwnershipHistory200ResponseWithDefaults() *OrgsOwnershipHistory200Response`

NewOrgsOwnershipHistory200ResponseWithDefaults instantiates a new OrgsOwnershipHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrgsOwnershipHistory200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgsOwnershipHistory200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgsOwnershipHistory200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *OrgsOwnershipHistory200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgsOwnershipHistory200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgsOwnershipHistory200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSection

`func (o *OrgsOwnershipHistory200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *OrgsOwnershipHistory200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *OrgsOwnershipHistory200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *OrgsOwnershipHistory200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *OrgsOwnershipHistory200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *OrgsOwnershipHistory200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *OrgsOwnershipHistory200Response) GetLens() string`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *OrgsOwnershipHistory200Response) GetLensOk() (*string, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *OrgsOwnershipHistory200Response) SetLens(v string)`

SetLens sets Lens field to given value.


### GetWebUrl

`func (o *OrgsOwnershipHistory200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OrgsOwnershipHistory200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OrgsOwnershipHistory200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *OrgsOwnershipHistory200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *OrgsOwnershipHistory200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *OrgsOwnershipHistory200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *OrgsOwnershipHistory200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *OrgsOwnershipHistory200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *OrgsOwnershipHistory200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *OrgsOwnershipHistory200Response) GetRows() []OrgsOwnershipHistory200ResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *OrgsOwnershipHistory200Response) GetRowsOk() (*[]OrgsOwnershipHistory200ResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *OrgsOwnershipHistory200Response) SetRows(v []OrgsOwnershipHistory200ResponseRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *OrgsOwnershipHistory200Response) GetAttribution() []OrgsCurrentOwnership200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsOwnershipHistory200Response) GetAttributionOk() (*[]OrgsCurrentOwnership200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsOwnershipHistory200Response) SetAttribution(v []OrgsCurrentOwnership200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetPagination

`func (o *OrgsOwnershipHistory200Response) GetPagination() BiomarkersAssayProviderLedger200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *OrgsOwnershipHistory200Response) GetPaginationOk() (*BiomarkersAssayProviderLedger200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *OrgsOwnershipHistory200Response) SetPagination(v BiomarkersAssayProviderLedger200ResponsePagination)`

SetPagination sets Pagination field to given value.


### GetContext

`func (o *OrgsOwnershipHistory200Response) GetContext() OrgsOwnershipHistory200ResponseContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OrgsOwnershipHistory200Response) GetContextOk() (*OrgsOwnershipHistory200ResponseContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OrgsOwnershipHistory200Response) SetContext(v OrgsOwnershipHistory200ResponseContext)`

SetContext sets Context field to given value.


### GetBuildState

`func (o *OrgsOwnershipHistory200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *OrgsOwnershipHistory200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *OrgsOwnershipHistory200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



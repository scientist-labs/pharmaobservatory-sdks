# OrgsOrgOwnership200Response

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
**Rows** | [**[]OrgsCurrentOwnership200ResponseRowsInner**](OrgsCurrentOwnership200ResponseRowsInner.md) |  | 
**Attribution** | [**[]OrgsCurrentOwnership200ResponseAttributionInner**](OrgsCurrentOwnership200ResponseAttributionInner.md) |  | 
**Context** | [**OrgsOrgOwnership200ResponseContext**](OrgsOrgOwnership200ResponseContext.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 

## Methods

### NewOrgsOrgOwnership200Response

`func NewOrgsOrgOwnership200Response(page string, id string, section string, heading string, lens string, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []OrgsCurrentOwnership200ResponseRowsInner, attribution []OrgsCurrentOwnership200ResponseAttributionInner, context OrgsOrgOwnership200ResponseContext, buildState BiomarkersList200ResponseBuildState, ) *OrgsOrgOwnership200Response`

NewOrgsOrgOwnership200Response instantiates a new OrgsOrgOwnership200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsOrgOwnership200ResponseWithDefaults

`func NewOrgsOrgOwnership200ResponseWithDefaults() *OrgsOrgOwnership200Response`

NewOrgsOrgOwnership200ResponseWithDefaults instantiates a new OrgsOrgOwnership200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrgsOrgOwnership200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgsOrgOwnership200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgsOrgOwnership200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *OrgsOrgOwnership200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgsOrgOwnership200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgsOrgOwnership200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSection

`func (o *OrgsOrgOwnership200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *OrgsOrgOwnership200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *OrgsOrgOwnership200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *OrgsOrgOwnership200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *OrgsOrgOwnership200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *OrgsOrgOwnership200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *OrgsOrgOwnership200Response) GetLens() string`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *OrgsOrgOwnership200Response) GetLensOk() (*string, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *OrgsOrgOwnership200Response) SetLens(v string)`

SetLens sets Lens field to given value.


### GetWebUrl

`func (o *OrgsOrgOwnership200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OrgsOrgOwnership200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OrgsOrgOwnership200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *OrgsOrgOwnership200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *OrgsOrgOwnership200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *OrgsOrgOwnership200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *OrgsOrgOwnership200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *OrgsOrgOwnership200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *OrgsOrgOwnership200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *OrgsOrgOwnership200Response) GetRows() []OrgsCurrentOwnership200ResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *OrgsOrgOwnership200Response) GetRowsOk() (*[]OrgsCurrentOwnership200ResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *OrgsOrgOwnership200Response) SetRows(v []OrgsCurrentOwnership200ResponseRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *OrgsOrgOwnership200Response) GetAttribution() []OrgsCurrentOwnership200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsOrgOwnership200Response) GetAttributionOk() (*[]OrgsCurrentOwnership200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsOrgOwnership200Response) SetAttribution(v []OrgsCurrentOwnership200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetContext

`func (o *OrgsOrgOwnership200Response) GetContext() OrgsOrgOwnership200ResponseContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OrgsOrgOwnership200Response) GetContextOk() (*OrgsOrgOwnership200ResponseContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OrgsOrgOwnership200Response) SetContext(v OrgsOrgOwnership200ResponseContext)`

SetContext sets Context field to given value.


### GetBuildState

`func (o *OrgsOrgOwnership200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *OrgsOrgOwnership200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *OrgsOrgOwnership200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



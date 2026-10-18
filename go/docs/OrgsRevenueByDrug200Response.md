# OrgsRevenueByDrug200Response

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
**Rows** | [**[]OrgsRevenueByDrug200ResponseRowsInner**](OrgsRevenueByDrug200ResponseRowsInner.md) |  | 
**Attribution** | [**[]BiomarkersAssayProviderLedger200ResponseAttributionInner**](BiomarkersAssayProviderLedger200ResponseAttributionInner.md) |  | 
**Context** | [**OrgsRevenueByDrug200ResponseContext**](OrgsRevenueByDrug200ResponseContext.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 

## Methods

### NewOrgsRevenueByDrug200Response

`func NewOrgsRevenueByDrug200Response(page string, id string, section string, heading string, lens string, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []OrgsRevenueByDrug200ResponseRowsInner, attribution []BiomarkersAssayProviderLedger200ResponseAttributionInner, context OrgsRevenueByDrug200ResponseContext, buildState BiomarkersList200ResponseBuildState, ) *OrgsRevenueByDrug200Response`

NewOrgsRevenueByDrug200Response instantiates a new OrgsRevenueByDrug200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsRevenueByDrug200ResponseWithDefaults

`func NewOrgsRevenueByDrug200ResponseWithDefaults() *OrgsRevenueByDrug200Response`

NewOrgsRevenueByDrug200ResponseWithDefaults instantiates a new OrgsRevenueByDrug200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrgsRevenueByDrug200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgsRevenueByDrug200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgsRevenueByDrug200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *OrgsRevenueByDrug200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgsRevenueByDrug200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgsRevenueByDrug200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSection

`func (o *OrgsRevenueByDrug200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *OrgsRevenueByDrug200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *OrgsRevenueByDrug200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *OrgsRevenueByDrug200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *OrgsRevenueByDrug200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *OrgsRevenueByDrug200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *OrgsRevenueByDrug200Response) GetLens() string`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *OrgsRevenueByDrug200Response) GetLensOk() (*string, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *OrgsRevenueByDrug200Response) SetLens(v string)`

SetLens sets Lens field to given value.


### GetWebUrl

`func (o *OrgsRevenueByDrug200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OrgsRevenueByDrug200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OrgsRevenueByDrug200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *OrgsRevenueByDrug200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *OrgsRevenueByDrug200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *OrgsRevenueByDrug200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *OrgsRevenueByDrug200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *OrgsRevenueByDrug200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *OrgsRevenueByDrug200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *OrgsRevenueByDrug200Response) GetRows() []OrgsRevenueByDrug200ResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *OrgsRevenueByDrug200Response) GetRowsOk() (*[]OrgsRevenueByDrug200ResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *OrgsRevenueByDrug200Response) SetRows(v []OrgsRevenueByDrug200ResponseRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *OrgsRevenueByDrug200Response) GetAttribution() []BiomarkersAssayProviderLedger200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsRevenueByDrug200Response) GetAttributionOk() (*[]BiomarkersAssayProviderLedger200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsRevenueByDrug200Response) SetAttribution(v []BiomarkersAssayProviderLedger200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetContext

`func (o *OrgsRevenueByDrug200Response) GetContext() OrgsRevenueByDrug200ResponseContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OrgsRevenueByDrug200Response) GetContextOk() (*OrgsRevenueByDrug200ResponseContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OrgsRevenueByDrug200Response) SetContext(v OrgsRevenueByDrug200ResponseContext)`

SetContext sets Context field to given value.


### GetBuildState

`func (o *OrgsRevenueByDrug200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *OrgsRevenueByDrug200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *OrgsRevenueByDrug200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



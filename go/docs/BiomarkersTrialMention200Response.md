# BiomarkersTrialMention200Response

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
**Rows** | [**[]BiomarkersTrialMention200ResponseRowsInner**](BiomarkersTrialMention200ResponseRowsInner.md) |  | 
**Attribution** | [**[]BiomarkersCgi200ResponseAttributionInner**](BiomarkersCgi200ResponseAttributionInner.md) |  | 
**Pagination** | [**BiomarkersAssayProviderLedger200ResponsePagination**](BiomarkersAssayProviderLedger200ResponsePagination.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 

## Methods

### NewBiomarkersTrialMention200Response

`func NewBiomarkersTrialMention200Response(page string, id string, section string, heading string, lens interface{}, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []BiomarkersTrialMention200ResponseRowsInner, attribution []BiomarkersCgi200ResponseAttributionInner, pagination BiomarkersAssayProviderLedger200ResponsePagination, buildState BiomarkersList200ResponseBuildState, ) *BiomarkersTrialMention200Response`

NewBiomarkersTrialMention200Response instantiates a new BiomarkersTrialMention200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkersTrialMention200ResponseWithDefaults

`func NewBiomarkersTrialMention200ResponseWithDefaults() *BiomarkersTrialMention200Response`

NewBiomarkersTrialMention200ResponseWithDefaults instantiates a new BiomarkersTrialMention200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *BiomarkersTrialMention200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BiomarkersTrialMention200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BiomarkersTrialMention200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *BiomarkersTrialMention200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BiomarkersTrialMention200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BiomarkersTrialMention200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSection

`func (o *BiomarkersTrialMention200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *BiomarkersTrialMention200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *BiomarkersTrialMention200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *BiomarkersTrialMention200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *BiomarkersTrialMention200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *BiomarkersTrialMention200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *BiomarkersTrialMention200Response) GetLens() interface{}`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *BiomarkersTrialMention200Response) GetLensOk() (*interface{}, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *BiomarkersTrialMention200Response) SetLens(v interface{})`

SetLens sets Lens field to given value.


### SetLensNil

`func (o *BiomarkersTrialMention200Response) SetLensNil(b bool)`

 SetLensNil sets the value for Lens to be an explicit nil

### UnsetLens
`func (o *BiomarkersTrialMention200Response) UnsetLens()`

UnsetLens ensures that no value is present for Lens, not even an explicit nil
### GetWebUrl

`func (o *BiomarkersTrialMention200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *BiomarkersTrialMention200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *BiomarkersTrialMention200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *BiomarkersTrialMention200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *BiomarkersTrialMention200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *BiomarkersTrialMention200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *BiomarkersTrialMention200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *BiomarkersTrialMention200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *BiomarkersTrialMention200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *BiomarkersTrialMention200Response) GetRows() []BiomarkersTrialMention200ResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *BiomarkersTrialMention200Response) GetRowsOk() (*[]BiomarkersTrialMention200ResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *BiomarkersTrialMention200Response) SetRows(v []BiomarkersTrialMention200ResponseRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *BiomarkersTrialMention200Response) GetAttribution() []BiomarkersCgi200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *BiomarkersTrialMention200Response) GetAttributionOk() (*[]BiomarkersCgi200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *BiomarkersTrialMention200Response) SetAttribution(v []BiomarkersCgi200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetPagination

`func (o *BiomarkersTrialMention200Response) GetPagination() BiomarkersAssayProviderLedger200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *BiomarkersTrialMention200Response) GetPaginationOk() (*BiomarkersAssayProviderLedger200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *BiomarkersTrialMention200Response) SetPagination(v BiomarkersAssayProviderLedger200ResponsePagination)`

SetPagination sets Pagination field to given value.


### GetBuildState

`func (o *BiomarkersTrialMention200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *BiomarkersTrialMention200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *BiomarkersTrialMention200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



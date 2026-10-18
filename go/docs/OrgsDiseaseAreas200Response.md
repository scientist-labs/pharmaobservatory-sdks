# OrgsDiseaseAreas200Response

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
**Rows** | [**[]OrgsDiseaseAreas200ResponseRowsInner**](OrgsDiseaseAreas200ResponseRowsInner.md) |  | 
**Attribution** | [**[]BiomarkersCgi200ResponseAttributionInner**](BiomarkersCgi200ResponseAttributionInner.md) |  | 
**Pagination** | [**BiomarkersAssayProviderLedger200ResponsePagination**](BiomarkersAssayProviderLedger200ResponsePagination.md) |  | 
**Context** | [**OrgsApprovedAssets200ResponseContext**](OrgsApprovedAssets200ResponseContext.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 

## Methods

### NewOrgsDiseaseAreas200Response

`func NewOrgsDiseaseAreas200Response(page string, id string, section string, heading string, lens interface{}, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []OrgsDiseaseAreas200ResponseRowsInner, attribution []BiomarkersCgi200ResponseAttributionInner, pagination BiomarkersAssayProviderLedger200ResponsePagination, context OrgsApprovedAssets200ResponseContext, buildState BiomarkersList200ResponseBuildState, ) *OrgsDiseaseAreas200Response`

NewOrgsDiseaseAreas200Response instantiates a new OrgsDiseaseAreas200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsDiseaseAreas200ResponseWithDefaults

`func NewOrgsDiseaseAreas200ResponseWithDefaults() *OrgsDiseaseAreas200Response`

NewOrgsDiseaseAreas200ResponseWithDefaults instantiates a new OrgsDiseaseAreas200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrgsDiseaseAreas200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgsDiseaseAreas200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgsDiseaseAreas200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *OrgsDiseaseAreas200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgsDiseaseAreas200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgsDiseaseAreas200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSection

`func (o *OrgsDiseaseAreas200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *OrgsDiseaseAreas200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *OrgsDiseaseAreas200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *OrgsDiseaseAreas200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *OrgsDiseaseAreas200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *OrgsDiseaseAreas200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *OrgsDiseaseAreas200Response) GetLens() interface{}`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *OrgsDiseaseAreas200Response) GetLensOk() (*interface{}, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *OrgsDiseaseAreas200Response) SetLens(v interface{})`

SetLens sets Lens field to given value.


### SetLensNil

`func (o *OrgsDiseaseAreas200Response) SetLensNil(b bool)`

 SetLensNil sets the value for Lens to be an explicit nil

### UnsetLens
`func (o *OrgsDiseaseAreas200Response) UnsetLens()`

UnsetLens ensures that no value is present for Lens, not even an explicit nil
### GetWebUrl

`func (o *OrgsDiseaseAreas200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OrgsDiseaseAreas200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OrgsDiseaseAreas200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *OrgsDiseaseAreas200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *OrgsDiseaseAreas200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *OrgsDiseaseAreas200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *OrgsDiseaseAreas200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *OrgsDiseaseAreas200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *OrgsDiseaseAreas200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *OrgsDiseaseAreas200Response) GetRows() []OrgsDiseaseAreas200ResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *OrgsDiseaseAreas200Response) GetRowsOk() (*[]OrgsDiseaseAreas200ResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *OrgsDiseaseAreas200Response) SetRows(v []OrgsDiseaseAreas200ResponseRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *OrgsDiseaseAreas200Response) GetAttribution() []BiomarkersCgi200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsDiseaseAreas200Response) GetAttributionOk() (*[]BiomarkersCgi200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsDiseaseAreas200Response) SetAttribution(v []BiomarkersCgi200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetPagination

`func (o *OrgsDiseaseAreas200Response) GetPagination() BiomarkersAssayProviderLedger200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *OrgsDiseaseAreas200Response) GetPaginationOk() (*BiomarkersAssayProviderLedger200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *OrgsDiseaseAreas200Response) SetPagination(v BiomarkersAssayProviderLedger200ResponsePagination)`

SetPagination sets Pagination field to given value.


### GetContext

`func (o *OrgsDiseaseAreas200Response) GetContext() OrgsApprovedAssets200ResponseContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OrgsDiseaseAreas200Response) GetContextOk() (*OrgsApprovedAssets200ResponseContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OrgsDiseaseAreas200Response) SetContext(v OrgsApprovedAssets200ResponseContext)`

SetContext sets Context field to given value.


### GetBuildState

`func (o *OrgsDiseaseAreas200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *OrgsDiseaseAreas200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *OrgsDiseaseAreas200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



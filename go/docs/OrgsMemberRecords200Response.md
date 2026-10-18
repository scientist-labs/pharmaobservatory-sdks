# OrgsMemberRecords200Response

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
**Rows** | [**[]OrgsMemberRecords200ResponseRowsInner**](OrgsMemberRecords200ResponseRowsInner.md) |  | 
**Attribution** | **[]interface{}** |  | 
**Context** | [**OrgsMemberRecords200ResponseContext**](OrgsMemberRecords200ResponseContext.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 

## Methods

### NewOrgsMemberRecords200Response

`func NewOrgsMemberRecords200Response(page string, id string, section string, heading string, lens interface{}, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []OrgsMemberRecords200ResponseRowsInner, attribution []interface{}, context OrgsMemberRecords200ResponseContext, buildState BiomarkersList200ResponseBuildState, ) *OrgsMemberRecords200Response`

NewOrgsMemberRecords200Response instantiates a new OrgsMemberRecords200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsMemberRecords200ResponseWithDefaults

`func NewOrgsMemberRecords200ResponseWithDefaults() *OrgsMemberRecords200Response`

NewOrgsMemberRecords200ResponseWithDefaults instantiates a new OrgsMemberRecords200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrgsMemberRecords200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgsMemberRecords200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgsMemberRecords200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *OrgsMemberRecords200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgsMemberRecords200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgsMemberRecords200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSection

`func (o *OrgsMemberRecords200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *OrgsMemberRecords200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *OrgsMemberRecords200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *OrgsMemberRecords200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *OrgsMemberRecords200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *OrgsMemberRecords200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *OrgsMemberRecords200Response) GetLens() interface{}`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *OrgsMemberRecords200Response) GetLensOk() (*interface{}, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *OrgsMemberRecords200Response) SetLens(v interface{})`

SetLens sets Lens field to given value.


### SetLensNil

`func (o *OrgsMemberRecords200Response) SetLensNil(b bool)`

 SetLensNil sets the value for Lens to be an explicit nil

### UnsetLens
`func (o *OrgsMemberRecords200Response) UnsetLens()`

UnsetLens ensures that no value is present for Lens, not even an explicit nil
### GetWebUrl

`func (o *OrgsMemberRecords200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OrgsMemberRecords200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OrgsMemberRecords200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *OrgsMemberRecords200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *OrgsMemberRecords200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *OrgsMemberRecords200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *OrgsMemberRecords200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *OrgsMemberRecords200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *OrgsMemberRecords200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *OrgsMemberRecords200Response) GetRows() []OrgsMemberRecords200ResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *OrgsMemberRecords200Response) GetRowsOk() (*[]OrgsMemberRecords200ResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *OrgsMemberRecords200Response) SetRows(v []OrgsMemberRecords200ResponseRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *OrgsMemberRecords200Response) GetAttribution() []interface{}`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsMemberRecords200Response) GetAttributionOk() (*[]interface{}, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsMemberRecords200Response) SetAttribution(v []interface{})`

SetAttribution sets Attribution field to given value.


### GetContext

`func (o *OrgsMemberRecords200Response) GetContext() OrgsMemberRecords200ResponseContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OrgsMemberRecords200Response) GetContextOk() (*OrgsMemberRecords200ResponseContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OrgsMemberRecords200Response) SetContext(v OrgsMemberRecords200ResponseContext)`

SetContext sets Context field to given value.


### GetBuildState

`func (o *OrgsMemberRecords200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *OrgsMemberRecords200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *OrgsMemberRecords200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



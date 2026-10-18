# OrgsPdufaReconciliation200Response

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
**Rows** | [**[]OrgsPdufaReconciliation200ResponseRowsInner**](OrgsPdufaReconciliation200ResponseRowsInner.md) |  | 
**Attribution** | [**[]OrgsBenzingaCatalysts200ResponseAttributionInner**](OrgsBenzingaCatalysts200ResponseAttributionInner.md) |  | 
**Context** | [**OrgsPdufaReconciliation200ResponseContext**](OrgsPdufaReconciliation200ResponseContext.md) |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 
**Restriction** | [**OrgsBenzingaCatalysts200ResponseRestriction**](OrgsBenzingaCatalysts200ResponseRestriction.md) |  | 

## Methods

### NewOrgsPdufaReconciliation200Response

`func NewOrgsPdufaReconciliation200Response(page string, id string, section string, heading string, lens string, webUrl string, shape BiomarkersList200ResponseShape, columns []BiomarkersList200ResponseColumnsInner, rows []OrgsPdufaReconciliation200ResponseRowsInner, attribution []OrgsBenzingaCatalysts200ResponseAttributionInner, context OrgsPdufaReconciliation200ResponseContext, buildState BiomarkersList200ResponseBuildState, restriction OrgsBenzingaCatalysts200ResponseRestriction, ) *OrgsPdufaReconciliation200Response`

NewOrgsPdufaReconciliation200Response instantiates a new OrgsPdufaReconciliation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsPdufaReconciliation200ResponseWithDefaults

`func NewOrgsPdufaReconciliation200ResponseWithDefaults() *OrgsPdufaReconciliation200Response`

NewOrgsPdufaReconciliation200ResponseWithDefaults instantiates a new OrgsPdufaReconciliation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrgsPdufaReconciliation200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgsPdufaReconciliation200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgsPdufaReconciliation200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *OrgsPdufaReconciliation200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgsPdufaReconciliation200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgsPdufaReconciliation200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSection

`func (o *OrgsPdufaReconciliation200Response) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *OrgsPdufaReconciliation200Response) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *OrgsPdufaReconciliation200Response) SetSection(v string)`

SetSection sets Section field to given value.


### GetHeading

`func (o *OrgsPdufaReconciliation200Response) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *OrgsPdufaReconciliation200Response) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *OrgsPdufaReconciliation200Response) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetLens

`func (o *OrgsPdufaReconciliation200Response) GetLens() string`

GetLens returns the Lens field if non-nil, zero value otherwise.

### GetLensOk

`func (o *OrgsPdufaReconciliation200Response) GetLensOk() (*string, bool)`

GetLensOk returns a tuple with the Lens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLens

`func (o *OrgsPdufaReconciliation200Response) SetLens(v string)`

SetLens sets Lens field to given value.


### GetWebUrl

`func (o *OrgsPdufaReconciliation200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OrgsPdufaReconciliation200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OrgsPdufaReconciliation200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetShape

`func (o *OrgsPdufaReconciliation200Response) GetShape() BiomarkersList200ResponseShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *OrgsPdufaReconciliation200Response) GetShapeOk() (*BiomarkersList200ResponseShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *OrgsPdufaReconciliation200Response) SetShape(v BiomarkersList200ResponseShape)`

SetShape sets Shape field to given value.


### GetColumns

`func (o *OrgsPdufaReconciliation200Response) GetColumns() []BiomarkersList200ResponseColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *OrgsPdufaReconciliation200Response) GetColumnsOk() (*[]BiomarkersList200ResponseColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *OrgsPdufaReconciliation200Response) SetColumns(v []BiomarkersList200ResponseColumnsInner)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *OrgsPdufaReconciliation200Response) GetRows() []OrgsPdufaReconciliation200ResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *OrgsPdufaReconciliation200Response) GetRowsOk() (*[]OrgsPdufaReconciliation200ResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *OrgsPdufaReconciliation200Response) SetRows(v []OrgsPdufaReconciliation200ResponseRowsInner)`

SetRows sets Rows field to given value.


### GetAttribution

`func (o *OrgsPdufaReconciliation200Response) GetAttribution() []OrgsBenzingaCatalysts200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsPdufaReconciliation200Response) GetAttributionOk() (*[]OrgsBenzingaCatalysts200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsPdufaReconciliation200Response) SetAttribution(v []OrgsBenzingaCatalysts200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetContext

`func (o *OrgsPdufaReconciliation200Response) GetContext() OrgsPdufaReconciliation200ResponseContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OrgsPdufaReconciliation200Response) GetContextOk() (*OrgsPdufaReconciliation200ResponseContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OrgsPdufaReconciliation200Response) SetContext(v OrgsPdufaReconciliation200ResponseContext)`

SetContext sets Context field to given value.


### GetBuildState

`func (o *OrgsPdufaReconciliation200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *OrgsPdufaReconciliation200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *OrgsPdufaReconciliation200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetRestriction

`func (o *OrgsPdufaReconciliation200Response) GetRestriction() OrgsBenzingaCatalysts200ResponseRestriction`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *OrgsPdufaReconciliation200Response) GetRestrictionOk() (*OrgsBenzingaCatalysts200ResponseRestriction, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *OrgsPdufaReconciliation200Response) SetRestriction(v OrgsBenzingaCatalysts200ResponseRestriction)`

SetRestriction sets Restriction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



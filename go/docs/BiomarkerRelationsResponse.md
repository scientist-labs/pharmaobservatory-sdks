# BiomarkerRelationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**Grain** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to [**NullableBiomarkerCdxAssertionsResponseScope**](BiomarkerCdxAssertionsResponseScope.md) |  | [optional] 
**Columns** | [**[]DatasetColumn**](DatasetColumn.md) |  | 
**Rows** | [**[]BiomarkerRelationsRow**](BiomarkerRelationsRow.md) |  | 
**Total** | Pointer to **NullableInt32** | The exact count of the filtered rows on an ungrouped read, carried by the cursor; null when grouped. | [optional] 
**NextCursor** | Pointer to **NullableString** |  | [optional] 
**Sort** | Pointer to **[]string** |  | [optional] 
**Filters** | Pointer to **[]map[string]interface{}** |  | [optional] 
**GroupBy** | Pointer to **[]string** |  | [optional] 
**Measures** | Pointer to **[]string** |  | [optional] 
**Attribution** | Pointer to [**[]DatasetAttribution**](DatasetAttribution.md) |  | [optional] 
**SourceCredit** | Pointer to **string** |  | [optional] 
**WithheldTiers** | Pointer to **[]string** |  | [optional] 
**WithheldColumns** | Pointer to **[]string** |  | [optional] 
**Freshness** | Pointer to [**BiomarkerCdxAssertionsResponseFreshness**](BiomarkerCdxAssertionsResponseFreshness.md) |  | [optional] 
**Build** | Pointer to [**BiomarkerCdxAssertionsResponseBuild**](BiomarkerCdxAssertionsResponseBuild.md) |  | [optional] 
**BuildState** | Pointer to [**DatasetBuildState**](DatasetBuildState.md) |  | [optional] 
**Restriction** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBiomarkerRelationsResponse

`func NewBiomarkerRelationsResponse(dataset string, columns []DatasetColumn, rows []BiomarkerRelationsRow, ) *BiomarkerRelationsResponse`

NewBiomarkerRelationsResponse instantiates a new BiomarkerRelationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerRelationsResponseWithDefaults

`func NewBiomarkerRelationsResponseWithDefaults() *BiomarkerRelationsResponse`

NewBiomarkerRelationsResponseWithDefaults instantiates a new BiomarkerRelationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *BiomarkerRelationsResponse) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *BiomarkerRelationsResponse) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *BiomarkerRelationsResponse) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetGrain

`func (o *BiomarkerRelationsResponse) GetGrain() string`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *BiomarkerRelationsResponse) GetGrainOk() (*string, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *BiomarkerRelationsResponse) SetGrain(v string)`

SetGrain sets Grain field to given value.

### HasGrain

`func (o *BiomarkerRelationsResponse) HasGrain() bool`

HasGrain returns a boolean if a field has been set.

### GetObjectType

`func (o *BiomarkerRelationsResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiomarkerRelationsResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiomarkerRelationsResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *BiomarkerRelationsResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *BiomarkerRelationsResponse) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *BiomarkerRelationsResponse) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetId

`func (o *BiomarkerRelationsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BiomarkerRelationsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BiomarkerRelationsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BiomarkerRelationsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BiomarkerRelationsResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BiomarkerRelationsResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetScope

`func (o *BiomarkerRelationsResponse) GetScope() BiomarkerCdxAssertionsResponseScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *BiomarkerRelationsResponse) GetScopeOk() (*BiomarkerCdxAssertionsResponseScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *BiomarkerRelationsResponse) SetScope(v BiomarkerCdxAssertionsResponseScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *BiomarkerRelationsResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *BiomarkerRelationsResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *BiomarkerRelationsResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetColumns

`func (o *BiomarkerRelationsResponse) GetColumns() []DatasetColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *BiomarkerRelationsResponse) GetColumnsOk() (*[]DatasetColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *BiomarkerRelationsResponse) SetColumns(v []DatasetColumn)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *BiomarkerRelationsResponse) GetRows() []BiomarkerRelationsRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *BiomarkerRelationsResponse) GetRowsOk() (*[]BiomarkerRelationsRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *BiomarkerRelationsResponse) SetRows(v []BiomarkerRelationsRow)`

SetRows sets Rows field to given value.


### GetTotal

`func (o *BiomarkerRelationsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BiomarkerRelationsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BiomarkerRelationsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BiomarkerRelationsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *BiomarkerRelationsResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *BiomarkerRelationsResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetNextCursor

`func (o *BiomarkerRelationsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *BiomarkerRelationsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *BiomarkerRelationsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *BiomarkerRelationsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *BiomarkerRelationsResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *BiomarkerRelationsResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetSort

`func (o *BiomarkerRelationsResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *BiomarkerRelationsResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *BiomarkerRelationsResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *BiomarkerRelationsResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFilters

`func (o *BiomarkerRelationsResponse) GetFilters() []map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BiomarkerRelationsResponse) GetFiltersOk() (*[]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BiomarkerRelationsResponse) SetFilters(v []map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *BiomarkerRelationsResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *BiomarkerRelationsResponse) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *BiomarkerRelationsResponse) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *BiomarkerRelationsResponse) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *BiomarkerRelationsResponse) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetMeasures

`func (o *BiomarkerRelationsResponse) GetMeasures() []string`

GetMeasures returns the Measures field if non-nil, zero value otherwise.

### GetMeasuresOk

`func (o *BiomarkerRelationsResponse) GetMeasuresOk() (*[]string, bool)`

GetMeasuresOk returns a tuple with the Measures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasures

`func (o *BiomarkerRelationsResponse) SetMeasures(v []string)`

SetMeasures sets Measures field to given value.

### HasMeasures

`func (o *BiomarkerRelationsResponse) HasMeasures() bool`

HasMeasures returns a boolean if a field has been set.

### GetAttribution

`func (o *BiomarkerRelationsResponse) GetAttribution() []DatasetAttribution`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *BiomarkerRelationsResponse) GetAttributionOk() (*[]DatasetAttribution, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *BiomarkerRelationsResponse) SetAttribution(v []DatasetAttribution)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *BiomarkerRelationsResponse) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetSourceCredit

`func (o *BiomarkerRelationsResponse) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *BiomarkerRelationsResponse) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *BiomarkerRelationsResponse) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.

### HasSourceCredit

`func (o *BiomarkerRelationsResponse) HasSourceCredit() bool`

HasSourceCredit returns a boolean if a field has been set.

### GetWithheldTiers

`func (o *BiomarkerRelationsResponse) GetWithheldTiers() []string`

GetWithheldTiers returns the WithheldTiers field if non-nil, zero value otherwise.

### GetWithheldTiersOk

`func (o *BiomarkerRelationsResponse) GetWithheldTiersOk() (*[]string, bool)`

GetWithheldTiersOk returns a tuple with the WithheldTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldTiers

`func (o *BiomarkerRelationsResponse) SetWithheldTiers(v []string)`

SetWithheldTiers sets WithheldTiers field to given value.

### HasWithheldTiers

`func (o *BiomarkerRelationsResponse) HasWithheldTiers() bool`

HasWithheldTiers returns a boolean if a field has been set.

### GetWithheldColumns

`func (o *BiomarkerRelationsResponse) GetWithheldColumns() []string`

GetWithheldColumns returns the WithheldColumns field if non-nil, zero value otherwise.

### GetWithheldColumnsOk

`func (o *BiomarkerRelationsResponse) GetWithheldColumnsOk() (*[]string, bool)`

GetWithheldColumnsOk returns a tuple with the WithheldColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldColumns

`func (o *BiomarkerRelationsResponse) SetWithheldColumns(v []string)`

SetWithheldColumns sets WithheldColumns field to given value.

### HasWithheldColumns

`func (o *BiomarkerRelationsResponse) HasWithheldColumns() bool`

HasWithheldColumns returns a boolean if a field has been set.

### GetFreshness

`func (o *BiomarkerRelationsResponse) GetFreshness() BiomarkerCdxAssertionsResponseFreshness`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *BiomarkerRelationsResponse) GetFreshnessOk() (*BiomarkerCdxAssertionsResponseFreshness, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *BiomarkerRelationsResponse) SetFreshness(v BiomarkerCdxAssertionsResponseFreshness)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *BiomarkerRelationsResponse) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetBuild

`func (o *BiomarkerRelationsResponse) GetBuild() BiomarkerCdxAssertionsResponseBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *BiomarkerRelationsResponse) GetBuildOk() (*BiomarkerCdxAssertionsResponseBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *BiomarkerRelationsResponse) SetBuild(v BiomarkerCdxAssertionsResponseBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *BiomarkerRelationsResponse) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetBuildState

`func (o *BiomarkerRelationsResponse) GetBuildState() DatasetBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *BiomarkerRelationsResponse) GetBuildStateOk() (*DatasetBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *BiomarkerRelationsResponse) SetBuildState(v DatasetBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *BiomarkerRelationsResponse) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetRestriction

`func (o *BiomarkerRelationsResponse) GetRestriction() map[string]interface{}`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *BiomarkerRelationsResponse) GetRestrictionOk() (*map[string]interface{}, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *BiomarkerRelationsResponse) SetRestriction(v map[string]interface{})`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *BiomarkerRelationsResponse) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### SetRestrictionNil

`func (o *BiomarkerRelationsResponse) SetRestrictionNil(b bool)`

 SetRestrictionNil sets the value for Restriction to be an explicit nil

### UnsetRestriction
`func (o *BiomarkerRelationsResponse) UnsetRestriction()`

UnsetRestriction ensures that no value is present for Restriction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



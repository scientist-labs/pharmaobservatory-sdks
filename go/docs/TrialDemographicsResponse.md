# TrialDemographicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**Grain** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to [**NullableBiomarkerCdxAssertionsResponseScope**](BiomarkerCdxAssertionsResponseScope.md) |  | [optional] 
**Columns** | [**[]DatasetColumn**](DatasetColumn.md) |  | 
**Rows** | [**[]TrialDemographicsRow**](TrialDemographicsRow.md) |  | 
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

### NewTrialDemographicsResponse

`func NewTrialDemographicsResponse(dataset string, columns []DatasetColumn, rows []TrialDemographicsRow, ) *TrialDemographicsResponse`

NewTrialDemographicsResponse instantiates a new TrialDemographicsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialDemographicsResponseWithDefaults

`func NewTrialDemographicsResponseWithDefaults() *TrialDemographicsResponse`

NewTrialDemographicsResponseWithDefaults instantiates a new TrialDemographicsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *TrialDemographicsResponse) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *TrialDemographicsResponse) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *TrialDemographicsResponse) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetGrain

`func (o *TrialDemographicsResponse) GetGrain() string`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *TrialDemographicsResponse) GetGrainOk() (*string, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *TrialDemographicsResponse) SetGrain(v string)`

SetGrain sets Grain field to given value.

### HasGrain

`func (o *TrialDemographicsResponse) HasGrain() bool`

HasGrain returns a boolean if a field has been set.

### GetObjectType

`func (o *TrialDemographicsResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TrialDemographicsResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TrialDemographicsResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *TrialDemographicsResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *TrialDemographicsResponse) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *TrialDemographicsResponse) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetId

`func (o *TrialDemographicsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrialDemographicsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrialDemographicsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrialDemographicsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TrialDemographicsResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TrialDemographicsResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetScope

`func (o *TrialDemographicsResponse) GetScope() BiomarkerCdxAssertionsResponseScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TrialDemographicsResponse) GetScopeOk() (*BiomarkerCdxAssertionsResponseScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TrialDemographicsResponse) SetScope(v BiomarkerCdxAssertionsResponseScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TrialDemographicsResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *TrialDemographicsResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *TrialDemographicsResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetColumns

`func (o *TrialDemographicsResponse) GetColumns() []DatasetColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TrialDemographicsResponse) GetColumnsOk() (*[]DatasetColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TrialDemographicsResponse) SetColumns(v []DatasetColumn)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *TrialDemographicsResponse) GetRows() []TrialDemographicsRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *TrialDemographicsResponse) GetRowsOk() (*[]TrialDemographicsRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *TrialDemographicsResponse) SetRows(v []TrialDemographicsRow)`

SetRows sets Rows field to given value.


### GetTotal

`func (o *TrialDemographicsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrialDemographicsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrialDemographicsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TrialDemographicsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *TrialDemographicsResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *TrialDemographicsResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetNextCursor

`func (o *TrialDemographicsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *TrialDemographicsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *TrialDemographicsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *TrialDemographicsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *TrialDemographicsResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *TrialDemographicsResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetSort

`func (o *TrialDemographicsResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TrialDemographicsResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TrialDemographicsResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TrialDemographicsResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFilters

`func (o *TrialDemographicsResponse) GetFilters() []map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TrialDemographicsResponse) GetFiltersOk() (*[]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TrialDemographicsResponse) SetFilters(v []map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TrialDemographicsResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *TrialDemographicsResponse) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *TrialDemographicsResponse) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *TrialDemographicsResponse) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *TrialDemographicsResponse) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetMeasures

`func (o *TrialDemographicsResponse) GetMeasures() []string`

GetMeasures returns the Measures field if non-nil, zero value otherwise.

### GetMeasuresOk

`func (o *TrialDemographicsResponse) GetMeasuresOk() (*[]string, bool)`

GetMeasuresOk returns a tuple with the Measures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasures

`func (o *TrialDemographicsResponse) SetMeasures(v []string)`

SetMeasures sets Measures field to given value.

### HasMeasures

`func (o *TrialDemographicsResponse) HasMeasures() bool`

HasMeasures returns a boolean if a field has been set.

### GetAttribution

`func (o *TrialDemographicsResponse) GetAttribution() []DatasetAttribution`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *TrialDemographicsResponse) GetAttributionOk() (*[]DatasetAttribution, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *TrialDemographicsResponse) SetAttribution(v []DatasetAttribution)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *TrialDemographicsResponse) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetSourceCredit

`func (o *TrialDemographicsResponse) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *TrialDemographicsResponse) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *TrialDemographicsResponse) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.

### HasSourceCredit

`func (o *TrialDemographicsResponse) HasSourceCredit() bool`

HasSourceCredit returns a boolean if a field has been set.

### GetWithheldTiers

`func (o *TrialDemographicsResponse) GetWithheldTiers() []string`

GetWithheldTiers returns the WithheldTiers field if non-nil, zero value otherwise.

### GetWithheldTiersOk

`func (o *TrialDemographicsResponse) GetWithheldTiersOk() (*[]string, bool)`

GetWithheldTiersOk returns a tuple with the WithheldTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldTiers

`func (o *TrialDemographicsResponse) SetWithheldTiers(v []string)`

SetWithheldTiers sets WithheldTiers field to given value.

### HasWithheldTiers

`func (o *TrialDemographicsResponse) HasWithheldTiers() bool`

HasWithheldTiers returns a boolean if a field has been set.

### GetWithheldColumns

`func (o *TrialDemographicsResponse) GetWithheldColumns() []string`

GetWithheldColumns returns the WithheldColumns field if non-nil, zero value otherwise.

### GetWithheldColumnsOk

`func (o *TrialDemographicsResponse) GetWithheldColumnsOk() (*[]string, bool)`

GetWithheldColumnsOk returns a tuple with the WithheldColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldColumns

`func (o *TrialDemographicsResponse) SetWithheldColumns(v []string)`

SetWithheldColumns sets WithheldColumns field to given value.

### HasWithheldColumns

`func (o *TrialDemographicsResponse) HasWithheldColumns() bool`

HasWithheldColumns returns a boolean if a field has been set.

### GetFreshness

`func (o *TrialDemographicsResponse) GetFreshness() BiomarkerCdxAssertionsResponseFreshness`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *TrialDemographicsResponse) GetFreshnessOk() (*BiomarkerCdxAssertionsResponseFreshness, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *TrialDemographicsResponse) SetFreshness(v BiomarkerCdxAssertionsResponseFreshness)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *TrialDemographicsResponse) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetBuild

`func (o *TrialDemographicsResponse) GetBuild() BiomarkerCdxAssertionsResponseBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TrialDemographicsResponse) GetBuildOk() (*BiomarkerCdxAssertionsResponseBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TrialDemographicsResponse) SetBuild(v BiomarkerCdxAssertionsResponseBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TrialDemographicsResponse) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetBuildState

`func (o *TrialDemographicsResponse) GetBuildState() DatasetBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *TrialDemographicsResponse) GetBuildStateOk() (*DatasetBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *TrialDemographicsResponse) SetBuildState(v DatasetBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *TrialDemographicsResponse) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetRestriction

`func (o *TrialDemographicsResponse) GetRestriction() map[string]interface{}`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *TrialDemographicsResponse) GetRestrictionOk() (*map[string]interface{}, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *TrialDemographicsResponse) SetRestriction(v map[string]interface{})`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *TrialDemographicsResponse) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### SetRestrictionNil

`func (o *TrialDemographicsResponse) SetRestrictionNil(b bool)`

 SetRestrictionNil sets the value for Restriction to be an explicit nil

### UnsetRestriction
`func (o *TrialDemographicsResponse) UnsetRestriction()`

UnsetRestriction ensures that no value is present for Restriction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



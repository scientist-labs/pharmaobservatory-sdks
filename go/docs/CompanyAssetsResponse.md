# CompanyAssetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**Grain** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to [**NullableBiomarkerCdxAssertionsResponseScope**](BiomarkerCdxAssertionsResponseScope.md) |  | [optional] 
**Columns** | [**[]DatasetColumn**](DatasetColumn.md) |  | 
**Rows** | [**[]CompanyAssetsRow**](CompanyAssetsRow.md) |  | 
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

### NewCompanyAssetsResponse

`func NewCompanyAssetsResponse(dataset string, columns []DatasetColumn, rows []CompanyAssetsRow, ) *CompanyAssetsResponse`

NewCompanyAssetsResponse instantiates a new CompanyAssetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyAssetsResponseWithDefaults

`func NewCompanyAssetsResponseWithDefaults() *CompanyAssetsResponse`

NewCompanyAssetsResponseWithDefaults instantiates a new CompanyAssetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *CompanyAssetsResponse) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *CompanyAssetsResponse) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *CompanyAssetsResponse) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetGrain

`func (o *CompanyAssetsResponse) GetGrain() string`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *CompanyAssetsResponse) GetGrainOk() (*string, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *CompanyAssetsResponse) SetGrain(v string)`

SetGrain sets Grain field to given value.

### HasGrain

`func (o *CompanyAssetsResponse) HasGrain() bool`

HasGrain returns a boolean if a field has been set.

### GetObjectType

`func (o *CompanyAssetsResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CompanyAssetsResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CompanyAssetsResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CompanyAssetsResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *CompanyAssetsResponse) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *CompanyAssetsResponse) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetId

`func (o *CompanyAssetsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyAssetsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyAssetsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyAssetsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CompanyAssetsResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CompanyAssetsResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetScope

`func (o *CompanyAssetsResponse) GetScope() BiomarkerCdxAssertionsResponseScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CompanyAssetsResponse) GetScopeOk() (*BiomarkerCdxAssertionsResponseScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CompanyAssetsResponse) SetScope(v BiomarkerCdxAssertionsResponseScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CompanyAssetsResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *CompanyAssetsResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *CompanyAssetsResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetColumns

`func (o *CompanyAssetsResponse) GetColumns() []DatasetColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *CompanyAssetsResponse) GetColumnsOk() (*[]DatasetColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *CompanyAssetsResponse) SetColumns(v []DatasetColumn)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *CompanyAssetsResponse) GetRows() []CompanyAssetsRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *CompanyAssetsResponse) GetRowsOk() (*[]CompanyAssetsRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *CompanyAssetsResponse) SetRows(v []CompanyAssetsRow)`

SetRows sets Rows field to given value.


### GetTotal

`func (o *CompanyAssetsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CompanyAssetsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CompanyAssetsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CompanyAssetsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *CompanyAssetsResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *CompanyAssetsResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetNextCursor

`func (o *CompanyAssetsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CompanyAssetsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CompanyAssetsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *CompanyAssetsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *CompanyAssetsResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *CompanyAssetsResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetSort

`func (o *CompanyAssetsResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CompanyAssetsResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CompanyAssetsResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CompanyAssetsResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFilters

`func (o *CompanyAssetsResponse) GetFilters() []map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CompanyAssetsResponse) GetFiltersOk() (*[]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CompanyAssetsResponse) SetFilters(v []map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CompanyAssetsResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *CompanyAssetsResponse) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *CompanyAssetsResponse) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *CompanyAssetsResponse) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *CompanyAssetsResponse) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetMeasures

`func (o *CompanyAssetsResponse) GetMeasures() []string`

GetMeasures returns the Measures field if non-nil, zero value otherwise.

### GetMeasuresOk

`func (o *CompanyAssetsResponse) GetMeasuresOk() (*[]string, bool)`

GetMeasuresOk returns a tuple with the Measures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasures

`func (o *CompanyAssetsResponse) SetMeasures(v []string)`

SetMeasures sets Measures field to given value.

### HasMeasures

`func (o *CompanyAssetsResponse) HasMeasures() bool`

HasMeasures returns a boolean if a field has been set.

### GetAttribution

`func (o *CompanyAssetsResponse) GetAttribution() []DatasetAttribution`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *CompanyAssetsResponse) GetAttributionOk() (*[]DatasetAttribution, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *CompanyAssetsResponse) SetAttribution(v []DatasetAttribution)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *CompanyAssetsResponse) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetSourceCredit

`func (o *CompanyAssetsResponse) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *CompanyAssetsResponse) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *CompanyAssetsResponse) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.

### HasSourceCredit

`func (o *CompanyAssetsResponse) HasSourceCredit() bool`

HasSourceCredit returns a boolean if a field has been set.

### GetWithheldTiers

`func (o *CompanyAssetsResponse) GetWithheldTiers() []string`

GetWithheldTiers returns the WithheldTiers field if non-nil, zero value otherwise.

### GetWithheldTiersOk

`func (o *CompanyAssetsResponse) GetWithheldTiersOk() (*[]string, bool)`

GetWithheldTiersOk returns a tuple with the WithheldTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldTiers

`func (o *CompanyAssetsResponse) SetWithheldTiers(v []string)`

SetWithheldTiers sets WithheldTiers field to given value.

### HasWithheldTiers

`func (o *CompanyAssetsResponse) HasWithheldTiers() bool`

HasWithheldTiers returns a boolean if a field has been set.

### GetWithheldColumns

`func (o *CompanyAssetsResponse) GetWithheldColumns() []string`

GetWithheldColumns returns the WithheldColumns field if non-nil, zero value otherwise.

### GetWithheldColumnsOk

`func (o *CompanyAssetsResponse) GetWithheldColumnsOk() (*[]string, bool)`

GetWithheldColumnsOk returns a tuple with the WithheldColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldColumns

`func (o *CompanyAssetsResponse) SetWithheldColumns(v []string)`

SetWithheldColumns sets WithheldColumns field to given value.

### HasWithheldColumns

`func (o *CompanyAssetsResponse) HasWithheldColumns() bool`

HasWithheldColumns returns a boolean if a field has been set.

### GetFreshness

`func (o *CompanyAssetsResponse) GetFreshness() BiomarkerCdxAssertionsResponseFreshness`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *CompanyAssetsResponse) GetFreshnessOk() (*BiomarkerCdxAssertionsResponseFreshness, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *CompanyAssetsResponse) SetFreshness(v BiomarkerCdxAssertionsResponseFreshness)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *CompanyAssetsResponse) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetBuild

`func (o *CompanyAssetsResponse) GetBuild() BiomarkerCdxAssertionsResponseBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *CompanyAssetsResponse) GetBuildOk() (*BiomarkerCdxAssertionsResponseBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *CompanyAssetsResponse) SetBuild(v BiomarkerCdxAssertionsResponseBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *CompanyAssetsResponse) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetBuildState

`func (o *CompanyAssetsResponse) GetBuildState() DatasetBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *CompanyAssetsResponse) GetBuildStateOk() (*DatasetBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *CompanyAssetsResponse) SetBuildState(v DatasetBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *CompanyAssetsResponse) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetRestriction

`func (o *CompanyAssetsResponse) GetRestriction() map[string]interface{}`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *CompanyAssetsResponse) GetRestrictionOk() (*map[string]interface{}, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *CompanyAssetsResponse) SetRestriction(v map[string]interface{})`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *CompanyAssetsResponse) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### SetRestrictionNil

`func (o *CompanyAssetsResponse) SetRestrictionNil(b bool)`

 SetRestrictionNil sets the value for Restriction to be an explicit nil

### UnsetRestriction
`func (o *CompanyAssetsResponse) UnsetRestriction()`

UnsetRestriction ensures that no value is present for Restriction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



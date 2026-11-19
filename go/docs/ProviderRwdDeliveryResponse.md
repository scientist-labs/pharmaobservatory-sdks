# ProviderRwdDeliveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**Grain** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to [**NullableAnimalWelfareAssurancesResponseScope**](AnimalWelfareAssurancesResponseScope.md) |  | [optional] 
**Columns** | [**[]DatasetColumn**](DatasetColumn.md) |  | 
**Rows** | [**[]ProviderRwdDeliveryRow**](ProviderRwdDeliveryRow.md) |  | 
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
**Freshness** | Pointer to [**AnimalWelfareAssurancesResponseFreshness**](AnimalWelfareAssurancesResponseFreshness.md) |  | [optional] 
**Build** | Pointer to [**AnimalWelfareAssurancesResponseBuild**](AnimalWelfareAssurancesResponseBuild.md) |  | [optional] 
**BuildState** | Pointer to [**DatasetBuildState**](DatasetBuildState.md) |  | [optional] 
**Restriction** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProviderRwdDeliveryResponse

`func NewProviderRwdDeliveryResponse(dataset string, columns []DatasetColumn, rows []ProviderRwdDeliveryRow, ) *ProviderRwdDeliveryResponse`

NewProviderRwdDeliveryResponse instantiates a new ProviderRwdDeliveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRwdDeliveryResponseWithDefaults

`func NewProviderRwdDeliveryResponseWithDefaults() *ProviderRwdDeliveryResponse`

NewProviderRwdDeliveryResponseWithDefaults instantiates a new ProviderRwdDeliveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *ProviderRwdDeliveryResponse) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *ProviderRwdDeliveryResponse) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *ProviderRwdDeliveryResponse) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetGrain

`func (o *ProviderRwdDeliveryResponse) GetGrain() string`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *ProviderRwdDeliveryResponse) GetGrainOk() (*string, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *ProviderRwdDeliveryResponse) SetGrain(v string)`

SetGrain sets Grain field to given value.

### HasGrain

`func (o *ProviderRwdDeliveryResponse) HasGrain() bool`

HasGrain returns a boolean if a field has been set.

### GetObjectType

`func (o *ProviderRwdDeliveryResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProviderRwdDeliveryResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProviderRwdDeliveryResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ProviderRwdDeliveryResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *ProviderRwdDeliveryResponse) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *ProviderRwdDeliveryResponse) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetId

`func (o *ProviderRwdDeliveryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderRwdDeliveryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderRwdDeliveryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderRwdDeliveryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProviderRwdDeliveryResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProviderRwdDeliveryResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetScope

`func (o *ProviderRwdDeliveryResponse) GetScope() AnimalWelfareAssurancesResponseScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ProviderRwdDeliveryResponse) GetScopeOk() (*AnimalWelfareAssurancesResponseScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ProviderRwdDeliveryResponse) SetScope(v AnimalWelfareAssurancesResponseScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ProviderRwdDeliveryResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *ProviderRwdDeliveryResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *ProviderRwdDeliveryResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetColumns

`func (o *ProviderRwdDeliveryResponse) GetColumns() []DatasetColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ProviderRwdDeliveryResponse) GetColumnsOk() (*[]DatasetColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ProviderRwdDeliveryResponse) SetColumns(v []DatasetColumn)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *ProviderRwdDeliveryResponse) GetRows() []ProviderRwdDeliveryRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ProviderRwdDeliveryResponse) GetRowsOk() (*[]ProviderRwdDeliveryRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ProviderRwdDeliveryResponse) SetRows(v []ProviderRwdDeliveryRow)`

SetRows sets Rows field to given value.


### GetTotal

`func (o *ProviderRwdDeliveryResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProviderRwdDeliveryResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProviderRwdDeliveryResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProviderRwdDeliveryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ProviderRwdDeliveryResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ProviderRwdDeliveryResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetNextCursor

`func (o *ProviderRwdDeliveryResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ProviderRwdDeliveryResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ProviderRwdDeliveryResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ProviderRwdDeliveryResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ProviderRwdDeliveryResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ProviderRwdDeliveryResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetSort

`func (o *ProviderRwdDeliveryResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ProviderRwdDeliveryResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ProviderRwdDeliveryResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ProviderRwdDeliveryResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFilters

`func (o *ProviderRwdDeliveryResponse) GetFilters() []map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ProviderRwdDeliveryResponse) GetFiltersOk() (*[]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ProviderRwdDeliveryResponse) SetFilters(v []map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ProviderRwdDeliveryResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *ProviderRwdDeliveryResponse) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ProviderRwdDeliveryResponse) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ProviderRwdDeliveryResponse) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ProviderRwdDeliveryResponse) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetMeasures

`func (o *ProviderRwdDeliveryResponse) GetMeasures() []string`

GetMeasures returns the Measures field if non-nil, zero value otherwise.

### GetMeasuresOk

`func (o *ProviderRwdDeliveryResponse) GetMeasuresOk() (*[]string, bool)`

GetMeasuresOk returns a tuple with the Measures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasures

`func (o *ProviderRwdDeliveryResponse) SetMeasures(v []string)`

SetMeasures sets Measures field to given value.

### HasMeasures

`func (o *ProviderRwdDeliveryResponse) HasMeasures() bool`

HasMeasures returns a boolean if a field has been set.

### GetAttribution

`func (o *ProviderRwdDeliveryResponse) GetAttribution() []DatasetAttribution`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ProviderRwdDeliveryResponse) GetAttributionOk() (*[]DatasetAttribution, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ProviderRwdDeliveryResponse) SetAttribution(v []DatasetAttribution)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ProviderRwdDeliveryResponse) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetSourceCredit

`func (o *ProviderRwdDeliveryResponse) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *ProviderRwdDeliveryResponse) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *ProviderRwdDeliveryResponse) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.

### HasSourceCredit

`func (o *ProviderRwdDeliveryResponse) HasSourceCredit() bool`

HasSourceCredit returns a boolean if a field has been set.

### GetWithheldTiers

`func (o *ProviderRwdDeliveryResponse) GetWithheldTiers() []string`

GetWithheldTiers returns the WithheldTiers field if non-nil, zero value otherwise.

### GetWithheldTiersOk

`func (o *ProviderRwdDeliveryResponse) GetWithheldTiersOk() (*[]string, bool)`

GetWithheldTiersOk returns a tuple with the WithheldTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldTiers

`func (o *ProviderRwdDeliveryResponse) SetWithheldTiers(v []string)`

SetWithheldTiers sets WithheldTiers field to given value.

### HasWithheldTiers

`func (o *ProviderRwdDeliveryResponse) HasWithheldTiers() bool`

HasWithheldTiers returns a boolean if a field has been set.

### GetWithheldColumns

`func (o *ProviderRwdDeliveryResponse) GetWithheldColumns() []string`

GetWithheldColumns returns the WithheldColumns field if non-nil, zero value otherwise.

### GetWithheldColumnsOk

`func (o *ProviderRwdDeliveryResponse) GetWithheldColumnsOk() (*[]string, bool)`

GetWithheldColumnsOk returns a tuple with the WithheldColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldColumns

`func (o *ProviderRwdDeliveryResponse) SetWithheldColumns(v []string)`

SetWithheldColumns sets WithheldColumns field to given value.

### HasWithheldColumns

`func (o *ProviderRwdDeliveryResponse) HasWithheldColumns() bool`

HasWithheldColumns returns a boolean if a field has been set.

### GetFreshness

`func (o *ProviderRwdDeliveryResponse) GetFreshness() AnimalWelfareAssurancesResponseFreshness`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *ProviderRwdDeliveryResponse) GetFreshnessOk() (*AnimalWelfareAssurancesResponseFreshness, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *ProviderRwdDeliveryResponse) SetFreshness(v AnimalWelfareAssurancesResponseFreshness)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *ProviderRwdDeliveryResponse) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetBuild

`func (o *ProviderRwdDeliveryResponse) GetBuild() AnimalWelfareAssurancesResponseBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ProviderRwdDeliveryResponse) GetBuildOk() (*AnimalWelfareAssurancesResponseBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ProviderRwdDeliveryResponse) SetBuild(v AnimalWelfareAssurancesResponseBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ProviderRwdDeliveryResponse) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetBuildState

`func (o *ProviderRwdDeliveryResponse) GetBuildState() DatasetBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ProviderRwdDeliveryResponse) GetBuildStateOk() (*DatasetBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ProviderRwdDeliveryResponse) SetBuildState(v DatasetBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *ProviderRwdDeliveryResponse) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetRestriction

`func (o *ProviderRwdDeliveryResponse) GetRestriction() map[string]interface{}`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ProviderRwdDeliveryResponse) GetRestrictionOk() (*map[string]interface{}, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ProviderRwdDeliveryResponse) SetRestriction(v map[string]interface{})`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *ProviderRwdDeliveryResponse) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### SetRestrictionNil

`func (o *ProviderRwdDeliveryResponse) SetRestrictionNil(b bool)`

 SetRestrictionNil sets the value for Restriction to be an explicit nil

### UnsetRestriction
`func (o *ProviderRwdDeliveryResponse) UnsetRestriction()`

UnsetRestriction ensures that no value is present for Restriction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



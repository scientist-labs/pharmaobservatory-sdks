# ProviderBiomarkerCapabilitySummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**Grain** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to [**NullableAnimalWelfareAssurancesResponseScope**](AnimalWelfareAssurancesResponseScope.md) |  | [optional] 
**Columns** | [**[]DatasetColumn**](DatasetColumn.md) |  | 
**Rows** | [**[]ProviderBiomarkerCapabilitySummaryRow**](ProviderBiomarkerCapabilitySummaryRow.md) |  | 
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

### NewProviderBiomarkerCapabilitySummaryResponse

`func NewProviderBiomarkerCapabilitySummaryResponse(dataset string, columns []DatasetColumn, rows []ProviderBiomarkerCapabilitySummaryRow, ) *ProviderBiomarkerCapabilitySummaryResponse`

NewProviderBiomarkerCapabilitySummaryResponse instantiates a new ProviderBiomarkerCapabilitySummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderBiomarkerCapabilitySummaryResponseWithDefaults

`func NewProviderBiomarkerCapabilitySummaryResponseWithDefaults() *ProviderBiomarkerCapabilitySummaryResponse`

NewProviderBiomarkerCapabilitySummaryResponseWithDefaults instantiates a new ProviderBiomarkerCapabilitySummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetGrain

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetGrain() string`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetGrainOk() (*string, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetGrain(v string)`

SetGrain sets Grain field to given value.

### HasGrain

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasGrain() bool`

HasGrain returns a boolean if a field has been set.

### GetObjectType

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *ProviderBiomarkerCapabilitySummaryResponse) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetId

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProviderBiomarkerCapabilitySummaryResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetScope

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetScope() AnimalWelfareAssurancesResponseScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetScopeOk() (*AnimalWelfareAssurancesResponseScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetScope(v AnimalWelfareAssurancesResponseScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *ProviderBiomarkerCapabilitySummaryResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetColumns

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetColumns() []DatasetColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetColumnsOk() (*[]DatasetColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetColumns(v []DatasetColumn)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetRows() []ProviderBiomarkerCapabilitySummaryRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetRowsOk() (*[]ProviderBiomarkerCapabilitySummaryRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetRows(v []ProviderBiomarkerCapabilitySummaryRow)`

SetRows sets Rows field to given value.


### GetTotal

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ProviderBiomarkerCapabilitySummaryResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetNextCursor

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ProviderBiomarkerCapabilitySummaryResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetSort

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFilters

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetFilters() []map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetFiltersOk() (*[]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetFilters(v []map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetMeasures

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetMeasures() []string`

GetMeasures returns the Measures field if non-nil, zero value otherwise.

### GetMeasuresOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetMeasuresOk() (*[]string, bool)`

GetMeasuresOk returns a tuple with the Measures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasures

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetMeasures(v []string)`

SetMeasures sets Measures field to given value.

### HasMeasures

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasMeasures() bool`

HasMeasures returns a boolean if a field has been set.

### GetAttribution

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetAttribution() []DatasetAttribution`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetAttributionOk() (*[]DatasetAttribution, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetAttribution(v []DatasetAttribution)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetSourceCredit

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.

### HasSourceCredit

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasSourceCredit() bool`

HasSourceCredit returns a boolean if a field has been set.

### GetWithheldTiers

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetWithheldTiers() []string`

GetWithheldTiers returns the WithheldTiers field if non-nil, zero value otherwise.

### GetWithheldTiersOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetWithheldTiersOk() (*[]string, bool)`

GetWithheldTiersOk returns a tuple with the WithheldTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldTiers

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetWithheldTiers(v []string)`

SetWithheldTiers sets WithheldTiers field to given value.

### HasWithheldTiers

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasWithheldTiers() bool`

HasWithheldTiers returns a boolean if a field has been set.

### GetWithheldColumns

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetWithheldColumns() []string`

GetWithheldColumns returns the WithheldColumns field if non-nil, zero value otherwise.

### GetWithheldColumnsOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetWithheldColumnsOk() (*[]string, bool)`

GetWithheldColumnsOk returns a tuple with the WithheldColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldColumns

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetWithheldColumns(v []string)`

SetWithheldColumns sets WithheldColumns field to given value.

### HasWithheldColumns

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasWithheldColumns() bool`

HasWithheldColumns returns a boolean if a field has been set.

### GetFreshness

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetFreshness() AnimalWelfareAssurancesResponseFreshness`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetFreshnessOk() (*AnimalWelfareAssurancesResponseFreshness, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetFreshness(v AnimalWelfareAssurancesResponseFreshness)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetBuild

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetBuild() AnimalWelfareAssurancesResponseBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetBuildOk() (*AnimalWelfareAssurancesResponseBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetBuild(v AnimalWelfareAssurancesResponseBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetBuildState

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetBuildState() DatasetBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetBuildStateOk() (*DatasetBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetBuildState(v DatasetBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### GetRestriction

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetRestriction() map[string]interface{}`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ProviderBiomarkerCapabilitySummaryResponse) GetRestrictionOk() (*map[string]interface{}, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetRestriction(v map[string]interface{})`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *ProviderBiomarkerCapabilitySummaryResponse) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### SetRestrictionNil

`func (o *ProviderBiomarkerCapabilitySummaryResponse) SetRestrictionNil(b bool)`

 SetRestrictionNil sets the value for Restriction to be an explicit nil

### UnsetRestriction
`func (o *ProviderBiomarkerCapabilitySummaryResponse) UnsetRestriction()`

UnsetRestriction ensures that no value is present for Restriction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



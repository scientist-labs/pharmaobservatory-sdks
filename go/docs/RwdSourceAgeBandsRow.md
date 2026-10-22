# RwdSourceAgeBandsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | Pointer to **NullableString** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] 
**BandOrdinal** | Pointer to **NullableInt32** | The position of the age band in the source&#39;s declared list, from 1. | [optional] 
**AgeBand** | Pointer to **NullableString** | The age band label as the holder declared it. Bands can overlap (see is_nested_scheme); never sum them. | [optional] 
**PopulationSize** | Pointer to **NullableInt64** | The population size the holder declared for the source (all registered persons). | [optional] 
**ActivePopulationSize** | Pointer to **NullableInt64** | The active population size the holder declared for the source. | [optional] 
**HasPopulationBreakdown** | Pointer to **NullableBool** | Whether the holder published a population size breakdown by age band at all. | [optional] 
**HasActivePopulationBreakdown** | Pointer to **NullableBool** | Whether the holder published an active population breakdown by age band at all. | [optional] 
**IsNestedScheme** | Pointer to **NullableBool** | True when the source&#39;s bands overlap, so a sum over bands double counts. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdSourceAgeBandsRow

`func NewRwdSourceAgeBandsRow() *RwdSourceAgeBandsRow`

NewRwdSourceAgeBandsRow instantiates a new RwdSourceAgeBandsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdSourceAgeBandsRowWithDefaults

`func NewRwdSourceAgeBandsRowWithDefaults() *RwdSourceAgeBandsRow`

NewRwdSourceAgeBandsRowWithDefaults instantiates a new RwdSourceAgeBandsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *RwdSourceAgeBandsRow) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *RwdSourceAgeBandsRow) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *RwdSourceAgeBandsRow) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *RwdSourceAgeBandsRow) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *RwdSourceAgeBandsRow) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *RwdSourceAgeBandsRow) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetBandOrdinal

`func (o *RwdSourceAgeBandsRow) GetBandOrdinal() int32`

GetBandOrdinal returns the BandOrdinal field if non-nil, zero value otherwise.

### GetBandOrdinalOk

`func (o *RwdSourceAgeBandsRow) GetBandOrdinalOk() (*int32, bool)`

GetBandOrdinalOk returns a tuple with the BandOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOrdinal

`func (o *RwdSourceAgeBandsRow) SetBandOrdinal(v int32)`

SetBandOrdinal sets BandOrdinal field to given value.

### HasBandOrdinal

`func (o *RwdSourceAgeBandsRow) HasBandOrdinal() bool`

HasBandOrdinal returns a boolean if a field has been set.

### SetBandOrdinalNil

`func (o *RwdSourceAgeBandsRow) SetBandOrdinalNil(b bool)`

 SetBandOrdinalNil sets the value for BandOrdinal to be an explicit nil

### UnsetBandOrdinal
`func (o *RwdSourceAgeBandsRow) UnsetBandOrdinal()`

UnsetBandOrdinal ensures that no value is present for BandOrdinal, not even an explicit nil
### GetAgeBand

`func (o *RwdSourceAgeBandsRow) GetAgeBand() string`

GetAgeBand returns the AgeBand field if non-nil, zero value otherwise.

### GetAgeBandOk

`func (o *RwdSourceAgeBandsRow) GetAgeBandOk() (*string, bool)`

GetAgeBandOk returns a tuple with the AgeBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeBand

`func (o *RwdSourceAgeBandsRow) SetAgeBand(v string)`

SetAgeBand sets AgeBand field to given value.

### HasAgeBand

`func (o *RwdSourceAgeBandsRow) HasAgeBand() bool`

HasAgeBand returns a boolean if a field has been set.

### SetAgeBandNil

`func (o *RwdSourceAgeBandsRow) SetAgeBandNil(b bool)`

 SetAgeBandNil sets the value for AgeBand to be an explicit nil

### UnsetAgeBand
`func (o *RwdSourceAgeBandsRow) UnsetAgeBand()`

UnsetAgeBand ensures that no value is present for AgeBand, not even an explicit nil
### GetPopulationSize

`func (o *RwdSourceAgeBandsRow) GetPopulationSize() int64`

GetPopulationSize returns the PopulationSize field if non-nil, zero value otherwise.

### GetPopulationSizeOk

`func (o *RwdSourceAgeBandsRow) GetPopulationSizeOk() (*int64, bool)`

GetPopulationSizeOk returns a tuple with the PopulationSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationSize

`func (o *RwdSourceAgeBandsRow) SetPopulationSize(v int64)`

SetPopulationSize sets PopulationSize field to given value.

### HasPopulationSize

`func (o *RwdSourceAgeBandsRow) HasPopulationSize() bool`

HasPopulationSize returns a boolean if a field has been set.

### SetPopulationSizeNil

`func (o *RwdSourceAgeBandsRow) SetPopulationSizeNil(b bool)`

 SetPopulationSizeNil sets the value for PopulationSize to be an explicit nil

### UnsetPopulationSize
`func (o *RwdSourceAgeBandsRow) UnsetPopulationSize()`

UnsetPopulationSize ensures that no value is present for PopulationSize, not even an explicit nil
### GetActivePopulationSize

`func (o *RwdSourceAgeBandsRow) GetActivePopulationSize() int64`

GetActivePopulationSize returns the ActivePopulationSize field if non-nil, zero value otherwise.

### GetActivePopulationSizeOk

`func (o *RwdSourceAgeBandsRow) GetActivePopulationSizeOk() (*int64, bool)`

GetActivePopulationSizeOk returns a tuple with the ActivePopulationSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePopulationSize

`func (o *RwdSourceAgeBandsRow) SetActivePopulationSize(v int64)`

SetActivePopulationSize sets ActivePopulationSize field to given value.

### HasActivePopulationSize

`func (o *RwdSourceAgeBandsRow) HasActivePopulationSize() bool`

HasActivePopulationSize returns a boolean if a field has been set.

### SetActivePopulationSizeNil

`func (o *RwdSourceAgeBandsRow) SetActivePopulationSizeNil(b bool)`

 SetActivePopulationSizeNil sets the value for ActivePopulationSize to be an explicit nil

### UnsetActivePopulationSize
`func (o *RwdSourceAgeBandsRow) UnsetActivePopulationSize()`

UnsetActivePopulationSize ensures that no value is present for ActivePopulationSize, not even an explicit nil
### GetHasPopulationBreakdown

`func (o *RwdSourceAgeBandsRow) GetHasPopulationBreakdown() bool`

GetHasPopulationBreakdown returns the HasPopulationBreakdown field if non-nil, zero value otherwise.

### GetHasPopulationBreakdownOk

`func (o *RwdSourceAgeBandsRow) GetHasPopulationBreakdownOk() (*bool, bool)`

GetHasPopulationBreakdownOk returns a tuple with the HasPopulationBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPopulationBreakdown

`func (o *RwdSourceAgeBandsRow) SetHasPopulationBreakdown(v bool)`

SetHasPopulationBreakdown sets HasPopulationBreakdown field to given value.

### HasHasPopulationBreakdown

`func (o *RwdSourceAgeBandsRow) HasHasPopulationBreakdown() bool`

HasHasPopulationBreakdown returns a boolean if a field has been set.

### SetHasPopulationBreakdownNil

`func (o *RwdSourceAgeBandsRow) SetHasPopulationBreakdownNil(b bool)`

 SetHasPopulationBreakdownNil sets the value for HasPopulationBreakdown to be an explicit nil

### UnsetHasPopulationBreakdown
`func (o *RwdSourceAgeBandsRow) UnsetHasPopulationBreakdown()`

UnsetHasPopulationBreakdown ensures that no value is present for HasPopulationBreakdown, not even an explicit nil
### GetHasActivePopulationBreakdown

`func (o *RwdSourceAgeBandsRow) GetHasActivePopulationBreakdown() bool`

GetHasActivePopulationBreakdown returns the HasActivePopulationBreakdown field if non-nil, zero value otherwise.

### GetHasActivePopulationBreakdownOk

`func (o *RwdSourceAgeBandsRow) GetHasActivePopulationBreakdownOk() (*bool, bool)`

GetHasActivePopulationBreakdownOk returns a tuple with the HasActivePopulationBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActivePopulationBreakdown

`func (o *RwdSourceAgeBandsRow) SetHasActivePopulationBreakdown(v bool)`

SetHasActivePopulationBreakdown sets HasActivePopulationBreakdown field to given value.

### HasHasActivePopulationBreakdown

`func (o *RwdSourceAgeBandsRow) HasHasActivePopulationBreakdown() bool`

HasHasActivePopulationBreakdown returns a boolean if a field has been set.

### SetHasActivePopulationBreakdownNil

`func (o *RwdSourceAgeBandsRow) SetHasActivePopulationBreakdownNil(b bool)`

 SetHasActivePopulationBreakdownNil sets the value for HasActivePopulationBreakdown to be an explicit nil

### UnsetHasActivePopulationBreakdown
`func (o *RwdSourceAgeBandsRow) UnsetHasActivePopulationBreakdown()`

UnsetHasActivePopulationBreakdown ensures that no value is present for HasActivePopulationBreakdown, not even an explicit nil
### GetIsNestedScheme

`func (o *RwdSourceAgeBandsRow) GetIsNestedScheme() bool`

GetIsNestedScheme returns the IsNestedScheme field if non-nil, zero value otherwise.

### GetIsNestedSchemeOk

`func (o *RwdSourceAgeBandsRow) GetIsNestedSchemeOk() (*bool, bool)`

GetIsNestedSchemeOk returns a tuple with the IsNestedScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedScheme

`func (o *RwdSourceAgeBandsRow) SetIsNestedScheme(v bool)`

SetIsNestedScheme sets IsNestedScheme field to given value.

### HasIsNestedScheme

`func (o *RwdSourceAgeBandsRow) HasIsNestedScheme() bool`

HasIsNestedScheme returns a boolean if a field has been set.

### SetIsNestedSchemeNil

`func (o *RwdSourceAgeBandsRow) SetIsNestedSchemeNil(b bool)`

 SetIsNestedSchemeNil sets the value for IsNestedScheme to be an explicit nil

### UnsetIsNestedScheme
`func (o *RwdSourceAgeBandsRow) UnsetIsNestedScheme()`

UnsetIsNestedScheme ensures that no value is present for IsNestedScheme, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdSourceAgeBandsRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdSourceAgeBandsRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdSourceAgeBandsRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdSourceAgeBandsRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdSourceAgeBandsRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdSourceAgeBandsRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdSourceAgeBandsRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdSourceAgeBandsRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdSourceAgeBandsRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdSourceAgeBandsRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdSourceAgeBandsRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdSourceAgeBandsRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdSourceAgeBandsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdSourceAgeBandsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdSourceAgeBandsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdSourceAgeBandsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdSourceAgeBandsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdSourceAgeBandsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdSourceAgeBandsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdSourceAgeBandsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdSourceAgeBandsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdSourceAgeBandsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdSourceAgeBandsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdSourceAgeBandsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



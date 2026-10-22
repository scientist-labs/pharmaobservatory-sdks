# RwdConductorFootprintRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConductorNameKey** | Pointer to **NullableString** | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. | [optional] 
**ConductorId** | Pointer to **NullableString** | The conductor slug: the register&#39;s normalised institution name key with spaces as hyphens, the id in the web page address (/rwd-conductors/&lt;id&gt;). The entry key of the rwd_conductors object type; injective because no name key carries a hyphen. | [optional] 
**ConductorName** | Pointer to **NullableString** | The conductor&#39;s display name: one register spelling of the key. | [optional] 
**Dimension** | Pointer to **NullableString** | What the footprint value is: delivered_country (a country of a conducting study) or delivered_data_source (a registered data source a conducting study names). | [optional] 
**Value** | Pointer to **NullableString** | The country or data source name the row counts. | [optional] 
**StudyCount** | Pointer to **NullableInt32** | Distinct conducting studies of the conductor for this country or data source. | [optional] 
**ConductorStudyCount** | Pointer to **NullableInt32** | Distinct conducting studies of the conductor, the denominator beside study_count. | [optional] 
**DataSourceIds** | Pointer to **[]string** | The register Data source IDs of the registered sources with this name (delivered_data_source rows only; null for a country). | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdConductorFootprintRow

`func NewRwdConductorFootprintRow() *RwdConductorFootprintRow`

NewRwdConductorFootprintRow instantiates a new RwdConductorFootprintRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdConductorFootprintRowWithDefaults

`func NewRwdConductorFootprintRowWithDefaults() *RwdConductorFootprintRow`

NewRwdConductorFootprintRowWithDefaults instantiates a new RwdConductorFootprintRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConductorNameKey

`func (o *RwdConductorFootprintRow) GetConductorNameKey() string`

GetConductorNameKey returns the ConductorNameKey field if non-nil, zero value otherwise.

### GetConductorNameKeyOk

`func (o *RwdConductorFootprintRow) GetConductorNameKeyOk() (*string, bool)`

GetConductorNameKeyOk returns a tuple with the ConductorNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorNameKey

`func (o *RwdConductorFootprintRow) SetConductorNameKey(v string)`

SetConductorNameKey sets ConductorNameKey field to given value.

### HasConductorNameKey

`func (o *RwdConductorFootprintRow) HasConductorNameKey() bool`

HasConductorNameKey returns a boolean if a field has been set.

### SetConductorNameKeyNil

`func (o *RwdConductorFootprintRow) SetConductorNameKeyNil(b bool)`

 SetConductorNameKeyNil sets the value for ConductorNameKey to be an explicit nil

### UnsetConductorNameKey
`func (o *RwdConductorFootprintRow) UnsetConductorNameKey()`

UnsetConductorNameKey ensures that no value is present for ConductorNameKey, not even an explicit nil
### GetConductorId

`func (o *RwdConductorFootprintRow) GetConductorId() string`

GetConductorId returns the ConductorId field if non-nil, zero value otherwise.

### GetConductorIdOk

`func (o *RwdConductorFootprintRow) GetConductorIdOk() (*string, bool)`

GetConductorIdOk returns a tuple with the ConductorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorId

`func (o *RwdConductorFootprintRow) SetConductorId(v string)`

SetConductorId sets ConductorId field to given value.

### HasConductorId

`func (o *RwdConductorFootprintRow) HasConductorId() bool`

HasConductorId returns a boolean if a field has been set.

### SetConductorIdNil

`func (o *RwdConductorFootprintRow) SetConductorIdNil(b bool)`

 SetConductorIdNil sets the value for ConductorId to be an explicit nil

### UnsetConductorId
`func (o *RwdConductorFootprintRow) UnsetConductorId()`

UnsetConductorId ensures that no value is present for ConductorId, not even an explicit nil
### GetConductorName

`func (o *RwdConductorFootprintRow) GetConductorName() string`

GetConductorName returns the ConductorName field if non-nil, zero value otherwise.

### GetConductorNameOk

`func (o *RwdConductorFootprintRow) GetConductorNameOk() (*string, bool)`

GetConductorNameOk returns a tuple with the ConductorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorName

`func (o *RwdConductorFootprintRow) SetConductorName(v string)`

SetConductorName sets ConductorName field to given value.

### HasConductorName

`func (o *RwdConductorFootprintRow) HasConductorName() bool`

HasConductorName returns a boolean if a field has been set.

### SetConductorNameNil

`func (o *RwdConductorFootprintRow) SetConductorNameNil(b bool)`

 SetConductorNameNil sets the value for ConductorName to be an explicit nil

### UnsetConductorName
`func (o *RwdConductorFootprintRow) UnsetConductorName()`

UnsetConductorName ensures that no value is present for ConductorName, not even an explicit nil
### GetDimension

`func (o *RwdConductorFootprintRow) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *RwdConductorFootprintRow) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *RwdConductorFootprintRow) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *RwdConductorFootprintRow) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### SetDimensionNil

`func (o *RwdConductorFootprintRow) SetDimensionNil(b bool)`

 SetDimensionNil sets the value for Dimension to be an explicit nil

### UnsetDimension
`func (o *RwdConductorFootprintRow) UnsetDimension()`

UnsetDimension ensures that no value is present for Dimension, not even an explicit nil
### GetValue

`func (o *RwdConductorFootprintRow) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RwdConductorFootprintRow) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RwdConductorFootprintRow) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RwdConductorFootprintRow) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *RwdConductorFootprintRow) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *RwdConductorFootprintRow) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetStudyCount

`func (o *RwdConductorFootprintRow) GetStudyCount() int32`

GetStudyCount returns the StudyCount field if non-nil, zero value otherwise.

### GetStudyCountOk

`func (o *RwdConductorFootprintRow) GetStudyCountOk() (*int32, bool)`

GetStudyCountOk returns a tuple with the StudyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyCount

`func (o *RwdConductorFootprintRow) SetStudyCount(v int32)`

SetStudyCount sets StudyCount field to given value.

### HasStudyCount

`func (o *RwdConductorFootprintRow) HasStudyCount() bool`

HasStudyCount returns a boolean if a field has been set.

### SetStudyCountNil

`func (o *RwdConductorFootprintRow) SetStudyCountNil(b bool)`

 SetStudyCountNil sets the value for StudyCount to be an explicit nil

### UnsetStudyCount
`func (o *RwdConductorFootprintRow) UnsetStudyCount()`

UnsetStudyCount ensures that no value is present for StudyCount, not even an explicit nil
### GetConductorStudyCount

`func (o *RwdConductorFootprintRow) GetConductorStudyCount() int32`

GetConductorStudyCount returns the ConductorStudyCount field if non-nil, zero value otherwise.

### GetConductorStudyCountOk

`func (o *RwdConductorFootprintRow) GetConductorStudyCountOk() (*int32, bool)`

GetConductorStudyCountOk returns a tuple with the ConductorStudyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductorStudyCount

`func (o *RwdConductorFootprintRow) SetConductorStudyCount(v int32)`

SetConductorStudyCount sets ConductorStudyCount field to given value.

### HasConductorStudyCount

`func (o *RwdConductorFootprintRow) HasConductorStudyCount() bool`

HasConductorStudyCount returns a boolean if a field has been set.

### SetConductorStudyCountNil

`func (o *RwdConductorFootprintRow) SetConductorStudyCountNil(b bool)`

 SetConductorStudyCountNil sets the value for ConductorStudyCount to be an explicit nil

### UnsetConductorStudyCount
`func (o *RwdConductorFootprintRow) UnsetConductorStudyCount()`

UnsetConductorStudyCount ensures that no value is present for ConductorStudyCount, not even an explicit nil
### GetDataSourceIds

`func (o *RwdConductorFootprintRow) GetDataSourceIds() []string`

GetDataSourceIds returns the DataSourceIds field if non-nil, zero value otherwise.

### GetDataSourceIdsOk

`func (o *RwdConductorFootprintRow) GetDataSourceIdsOk() (*[]string, bool)`

GetDataSourceIdsOk returns a tuple with the DataSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceIds

`func (o *RwdConductorFootprintRow) SetDataSourceIds(v []string)`

SetDataSourceIds sets DataSourceIds field to given value.

### HasDataSourceIds

`func (o *RwdConductorFootprintRow) HasDataSourceIds() bool`

HasDataSourceIds returns a boolean if a field has been set.

### SetDataSourceIdsNil

`func (o *RwdConductorFootprintRow) SetDataSourceIdsNil(b bool)`

 SetDataSourceIdsNil sets the value for DataSourceIds to be an explicit nil

### UnsetDataSourceIds
`func (o *RwdConductorFootprintRow) UnsetDataSourceIds()`

UnsetDataSourceIds ensures that no value is present for DataSourceIds, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdConductorFootprintRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdConductorFootprintRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdConductorFootprintRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdConductorFootprintRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdConductorFootprintRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdConductorFootprintRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdConductorFootprintRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdConductorFootprintRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdConductorFootprintRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdConductorFootprintRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdConductorFootprintRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdConductorFootprintRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdConductorFootprintRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdConductorFootprintRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdConductorFootprintRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdConductorFootprintRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdConductorFootprintRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdConductorFootprintRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdConductorFootprintRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdConductorFootprintRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdConductorFootprintRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdConductorFootprintRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdConductorFootprintRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdConductorFootprintRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



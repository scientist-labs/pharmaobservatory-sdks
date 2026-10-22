# RwdSourceTypesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | Pointer to **NullableString** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] 
**ItemOrdinal** | Pointer to **NullableInt32** | The position of the item in the register&#39;s own list, from 1. | [optional] 
**TypeValue** | Pointer to **NullableString** | The declared data source type, as typed after the register&#39;s list is split. | [optional] 
**TypeKey** | Pointer to **NullableString** | A stable lowercase key for the data source type; the filter value. | [optional] 
**TypeLabel** | Pointer to **NullableString** | The register&#39;s most common spelling of the type. | [optional] 
**IsOtherSentinel** | Pointer to **NullableBool** | True when the item is the register&#39;s Other sentinel; the holder&#39;s free text for it is on the source row. | [optional] 
**SourceValue** | Pointer to **NullableString** | The whole register answer the item was split from. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdSourceTypesRow

`func NewRwdSourceTypesRow() *RwdSourceTypesRow`

NewRwdSourceTypesRow instantiates a new RwdSourceTypesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdSourceTypesRowWithDefaults

`func NewRwdSourceTypesRowWithDefaults() *RwdSourceTypesRow`

NewRwdSourceTypesRowWithDefaults instantiates a new RwdSourceTypesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *RwdSourceTypesRow) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *RwdSourceTypesRow) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *RwdSourceTypesRow) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *RwdSourceTypesRow) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *RwdSourceTypesRow) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *RwdSourceTypesRow) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetItemOrdinal

`func (o *RwdSourceTypesRow) GetItemOrdinal() int32`

GetItemOrdinal returns the ItemOrdinal field if non-nil, zero value otherwise.

### GetItemOrdinalOk

`func (o *RwdSourceTypesRow) GetItemOrdinalOk() (*int32, bool)`

GetItemOrdinalOk returns a tuple with the ItemOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemOrdinal

`func (o *RwdSourceTypesRow) SetItemOrdinal(v int32)`

SetItemOrdinal sets ItemOrdinal field to given value.

### HasItemOrdinal

`func (o *RwdSourceTypesRow) HasItemOrdinal() bool`

HasItemOrdinal returns a boolean if a field has been set.

### SetItemOrdinalNil

`func (o *RwdSourceTypesRow) SetItemOrdinalNil(b bool)`

 SetItemOrdinalNil sets the value for ItemOrdinal to be an explicit nil

### UnsetItemOrdinal
`func (o *RwdSourceTypesRow) UnsetItemOrdinal()`

UnsetItemOrdinal ensures that no value is present for ItemOrdinal, not even an explicit nil
### GetTypeValue

`func (o *RwdSourceTypesRow) GetTypeValue() string`

GetTypeValue returns the TypeValue field if non-nil, zero value otherwise.

### GetTypeValueOk

`func (o *RwdSourceTypesRow) GetTypeValueOk() (*string, bool)`

GetTypeValueOk returns a tuple with the TypeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeValue

`func (o *RwdSourceTypesRow) SetTypeValue(v string)`

SetTypeValue sets TypeValue field to given value.

### HasTypeValue

`func (o *RwdSourceTypesRow) HasTypeValue() bool`

HasTypeValue returns a boolean if a field has been set.

### SetTypeValueNil

`func (o *RwdSourceTypesRow) SetTypeValueNil(b bool)`

 SetTypeValueNil sets the value for TypeValue to be an explicit nil

### UnsetTypeValue
`func (o *RwdSourceTypesRow) UnsetTypeValue()`

UnsetTypeValue ensures that no value is present for TypeValue, not even an explicit nil
### GetTypeKey

`func (o *RwdSourceTypesRow) GetTypeKey() string`

GetTypeKey returns the TypeKey field if non-nil, zero value otherwise.

### GetTypeKeyOk

`func (o *RwdSourceTypesRow) GetTypeKeyOk() (*string, bool)`

GetTypeKeyOk returns a tuple with the TypeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeKey

`func (o *RwdSourceTypesRow) SetTypeKey(v string)`

SetTypeKey sets TypeKey field to given value.

### HasTypeKey

`func (o *RwdSourceTypesRow) HasTypeKey() bool`

HasTypeKey returns a boolean if a field has been set.

### SetTypeKeyNil

`func (o *RwdSourceTypesRow) SetTypeKeyNil(b bool)`

 SetTypeKeyNil sets the value for TypeKey to be an explicit nil

### UnsetTypeKey
`func (o *RwdSourceTypesRow) UnsetTypeKey()`

UnsetTypeKey ensures that no value is present for TypeKey, not even an explicit nil
### GetTypeLabel

`func (o *RwdSourceTypesRow) GetTypeLabel() string`

GetTypeLabel returns the TypeLabel field if non-nil, zero value otherwise.

### GetTypeLabelOk

`func (o *RwdSourceTypesRow) GetTypeLabelOk() (*string, bool)`

GetTypeLabelOk returns a tuple with the TypeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeLabel

`func (o *RwdSourceTypesRow) SetTypeLabel(v string)`

SetTypeLabel sets TypeLabel field to given value.

### HasTypeLabel

`func (o *RwdSourceTypesRow) HasTypeLabel() bool`

HasTypeLabel returns a boolean if a field has been set.

### SetTypeLabelNil

`func (o *RwdSourceTypesRow) SetTypeLabelNil(b bool)`

 SetTypeLabelNil sets the value for TypeLabel to be an explicit nil

### UnsetTypeLabel
`func (o *RwdSourceTypesRow) UnsetTypeLabel()`

UnsetTypeLabel ensures that no value is present for TypeLabel, not even an explicit nil
### GetIsOtherSentinel

`func (o *RwdSourceTypesRow) GetIsOtherSentinel() bool`

GetIsOtherSentinel returns the IsOtherSentinel field if non-nil, zero value otherwise.

### GetIsOtherSentinelOk

`func (o *RwdSourceTypesRow) GetIsOtherSentinelOk() (*bool, bool)`

GetIsOtherSentinelOk returns a tuple with the IsOtherSentinel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOtherSentinel

`func (o *RwdSourceTypesRow) SetIsOtherSentinel(v bool)`

SetIsOtherSentinel sets IsOtherSentinel field to given value.

### HasIsOtherSentinel

`func (o *RwdSourceTypesRow) HasIsOtherSentinel() bool`

HasIsOtherSentinel returns a boolean if a field has been set.

### SetIsOtherSentinelNil

`func (o *RwdSourceTypesRow) SetIsOtherSentinelNil(b bool)`

 SetIsOtherSentinelNil sets the value for IsOtherSentinel to be an explicit nil

### UnsetIsOtherSentinel
`func (o *RwdSourceTypesRow) UnsetIsOtherSentinel()`

UnsetIsOtherSentinel ensures that no value is present for IsOtherSentinel, not even an explicit nil
### GetSourceValue

`func (o *RwdSourceTypesRow) GetSourceValue() string`

GetSourceValue returns the SourceValue field if non-nil, zero value otherwise.

### GetSourceValueOk

`func (o *RwdSourceTypesRow) GetSourceValueOk() (*string, bool)`

GetSourceValueOk returns a tuple with the SourceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValue

`func (o *RwdSourceTypesRow) SetSourceValue(v string)`

SetSourceValue sets SourceValue field to given value.

### HasSourceValue

`func (o *RwdSourceTypesRow) HasSourceValue() bool`

HasSourceValue returns a boolean if a field has been set.

### SetSourceValueNil

`func (o *RwdSourceTypesRow) SetSourceValueNil(b bool)`

 SetSourceValueNil sets the value for SourceValue to be an explicit nil

### UnsetSourceValue
`func (o *RwdSourceTypesRow) UnsetSourceValue()`

UnsetSourceValue ensures that no value is present for SourceValue, not even an explicit nil
### GetRestrictionTier

`func (o *RwdSourceTypesRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdSourceTypesRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdSourceTypesRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdSourceTypesRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdSourceTypesRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdSourceTypesRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdSourceTypesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdSourceTypesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdSourceTypesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdSourceTypesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdSourceTypesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdSourceTypesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



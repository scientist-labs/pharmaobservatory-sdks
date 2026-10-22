# RwdSourceVocabulariesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | Pointer to **NullableString** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] 
**CapabilityKey** | Pointer to **NullableString** | Which capability the coding scheme answers: exposure (prescriptions, dispensing, medicinal product, indication) or outcomes (diagnosis, cause of death, procedures, and so on). | [optional] 
**CapabilityLabel** | Pointer to **NullableString** | The capability&#39;s label as the register words it. | [optional] 
**SourceColumn** | Pointer to **NullableString** | The register column the vocabulary item came from. | [optional] 
**ItemOrdinal** | Pointer to **NullableInt32** | The position of the item in the register&#39;s own list, from 1. | [optional] 
**VocabularyValue** | Pointer to **NullableString** | The declared coding scheme as typed. | [optional] 
**VocabularyKey** | Pointer to **NullableString** | A stable lowercase key for the coding scheme; the filter value. | [optional] 
**VocabularyLabel** | Pointer to **NullableString** | The register&#39;s most common spelling of the coding scheme. | [optional] 
**IsFreeText** | Pointer to **NullableBool** | True when the item is a free-text answer to the register&#39;s other-vocabulary question, not a coding scheme from its list. | [optional] 
**IsOtherSentinel** | Pointer to **NullableBool** | True when the item is the register&#39;s Other sentinel; the holder&#39;s free text for it is on the source row. | [optional] 
**SourceValue** | Pointer to **NullableString** | The whole register answer the item was split from. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdSourceVocabulariesRow

`func NewRwdSourceVocabulariesRow() *RwdSourceVocabulariesRow`

NewRwdSourceVocabulariesRow instantiates a new RwdSourceVocabulariesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdSourceVocabulariesRowWithDefaults

`func NewRwdSourceVocabulariesRowWithDefaults() *RwdSourceVocabulariesRow`

NewRwdSourceVocabulariesRowWithDefaults instantiates a new RwdSourceVocabulariesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *RwdSourceVocabulariesRow) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *RwdSourceVocabulariesRow) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *RwdSourceVocabulariesRow) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *RwdSourceVocabulariesRow) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *RwdSourceVocabulariesRow) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *RwdSourceVocabulariesRow) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetCapabilityKey

`func (o *RwdSourceVocabulariesRow) GetCapabilityKey() string`

GetCapabilityKey returns the CapabilityKey field if non-nil, zero value otherwise.

### GetCapabilityKeyOk

`func (o *RwdSourceVocabulariesRow) GetCapabilityKeyOk() (*string, bool)`

GetCapabilityKeyOk returns a tuple with the CapabilityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityKey

`func (o *RwdSourceVocabulariesRow) SetCapabilityKey(v string)`

SetCapabilityKey sets CapabilityKey field to given value.

### HasCapabilityKey

`func (o *RwdSourceVocabulariesRow) HasCapabilityKey() bool`

HasCapabilityKey returns a boolean if a field has been set.

### SetCapabilityKeyNil

`func (o *RwdSourceVocabulariesRow) SetCapabilityKeyNil(b bool)`

 SetCapabilityKeyNil sets the value for CapabilityKey to be an explicit nil

### UnsetCapabilityKey
`func (o *RwdSourceVocabulariesRow) UnsetCapabilityKey()`

UnsetCapabilityKey ensures that no value is present for CapabilityKey, not even an explicit nil
### GetCapabilityLabel

`func (o *RwdSourceVocabulariesRow) GetCapabilityLabel() string`

GetCapabilityLabel returns the CapabilityLabel field if non-nil, zero value otherwise.

### GetCapabilityLabelOk

`func (o *RwdSourceVocabulariesRow) GetCapabilityLabelOk() (*string, bool)`

GetCapabilityLabelOk returns a tuple with the CapabilityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityLabel

`func (o *RwdSourceVocabulariesRow) SetCapabilityLabel(v string)`

SetCapabilityLabel sets CapabilityLabel field to given value.

### HasCapabilityLabel

`func (o *RwdSourceVocabulariesRow) HasCapabilityLabel() bool`

HasCapabilityLabel returns a boolean if a field has been set.

### SetCapabilityLabelNil

`func (o *RwdSourceVocabulariesRow) SetCapabilityLabelNil(b bool)`

 SetCapabilityLabelNil sets the value for CapabilityLabel to be an explicit nil

### UnsetCapabilityLabel
`func (o *RwdSourceVocabulariesRow) UnsetCapabilityLabel()`

UnsetCapabilityLabel ensures that no value is present for CapabilityLabel, not even an explicit nil
### GetSourceColumn

`func (o *RwdSourceVocabulariesRow) GetSourceColumn() string`

GetSourceColumn returns the SourceColumn field if non-nil, zero value otherwise.

### GetSourceColumnOk

`func (o *RwdSourceVocabulariesRow) GetSourceColumnOk() (*string, bool)`

GetSourceColumnOk returns a tuple with the SourceColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceColumn

`func (o *RwdSourceVocabulariesRow) SetSourceColumn(v string)`

SetSourceColumn sets SourceColumn field to given value.

### HasSourceColumn

`func (o *RwdSourceVocabulariesRow) HasSourceColumn() bool`

HasSourceColumn returns a boolean if a field has been set.

### SetSourceColumnNil

`func (o *RwdSourceVocabulariesRow) SetSourceColumnNil(b bool)`

 SetSourceColumnNil sets the value for SourceColumn to be an explicit nil

### UnsetSourceColumn
`func (o *RwdSourceVocabulariesRow) UnsetSourceColumn()`

UnsetSourceColumn ensures that no value is present for SourceColumn, not even an explicit nil
### GetItemOrdinal

`func (o *RwdSourceVocabulariesRow) GetItemOrdinal() int32`

GetItemOrdinal returns the ItemOrdinal field if non-nil, zero value otherwise.

### GetItemOrdinalOk

`func (o *RwdSourceVocabulariesRow) GetItemOrdinalOk() (*int32, bool)`

GetItemOrdinalOk returns a tuple with the ItemOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemOrdinal

`func (o *RwdSourceVocabulariesRow) SetItemOrdinal(v int32)`

SetItemOrdinal sets ItemOrdinal field to given value.

### HasItemOrdinal

`func (o *RwdSourceVocabulariesRow) HasItemOrdinal() bool`

HasItemOrdinal returns a boolean if a field has been set.

### SetItemOrdinalNil

`func (o *RwdSourceVocabulariesRow) SetItemOrdinalNil(b bool)`

 SetItemOrdinalNil sets the value for ItemOrdinal to be an explicit nil

### UnsetItemOrdinal
`func (o *RwdSourceVocabulariesRow) UnsetItemOrdinal()`

UnsetItemOrdinal ensures that no value is present for ItemOrdinal, not even an explicit nil
### GetVocabularyValue

`func (o *RwdSourceVocabulariesRow) GetVocabularyValue() string`

GetVocabularyValue returns the VocabularyValue field if non-nil, zero value otherwise.

### GetVocabularyValueOk

`func (o *RwdSourceVocabulariesRow) GetVocabularyValueOk() (*string, bool)`

GetVocabularyValueOk returns a tuple with the VocabularyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVocabularyValue

`func (o *RwdSourceVocabulariesRow) SetVocabularyValue(v string)`

SetVocabularyValue sets VocabularyValue field to given value.

### HasVocabularyValue

`func (o *RwdSourceVocabulariesRow) HasVocabularyValue() bool`

HasVocabularyValue returns a boolean if a field has been set.

### SetVocabularyValueNil

`func (o *RwdSourceVocabulariesRow) SetVocabularyValueNil(b bool)`

 SetVocabularyValueNil sets the value for VocabularyValue to be an explicit nil

### UnsetVocabularyValue
`func (o *RwdSourceVocabulariesRow) UnsetVocabularyValue()`

UnsetVocabularyValue ensures that no value is present for VocabularyValue, not even an explicit nil
### GetVocabularyKey

`func (o *RwdSourceVocabulariesRow) GetVocabularyKey() string`

GetVocabularyKey returns the VocabularyKey field if non-nil, zero value otherwise.

### GetVocabularyKeyOk

`func (o *RwdSourceVocabulariesRow) GetVocabularyKeyOk() (*string, bool)`

GetVocabularyKeyOk returns a tuple with the VocabularyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVocabularyKey

`func (o *RwdSourceVocabulariesRow) SetVocabularyKey(v string)`

SetVocabularyKey sets VocabularyKey field to given value.

### HasVocabularyKey

`func (o *RwdSourceVocabulariesRow) HasVocabularyKey() bool`

HasVocabularyKey returns a boolean if a field has been set.

### SetVocabularyKeyNil

`func (o *RwdSourceVocabulariesRow) SetVocabularyKeyNil(b bool)`

 SetVocabularyKeyNil sets the value for VocabularyKey to be an explicit nil

### UnsetVocabularyKey
`func (o *RwdSourceVocabulariesRow) UnsetVocabularyKey()`

UnsetVocabularyKey ensures that no value is present for VocabularyKey, not even an explicit nil
### GetVocabularyLabel

`func (o *RwdSourceVocabulariesRow) GetVocabularyLabel() string`

GetVocabularyLabel returns the VocabularyLabel field if non-nil, zero value otherwise.

### GetVocabularyLabelOk

`func (o *RwdSourceVocabulariesRow) GetVocabularyLabelOk() (*string, bool)`

GetVocabularyLabelOk returns a tuple with the VocabularyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVocabularyLabel

`func (o *RwdSourceVocabulariesRow) SetVocabularyLabel(v string)`

SetVocabularyLabel sets VocabularyLabel field to given value.

### HasVocabularyLabel

`func (o *RwdSourceVocabulariesRow) HasVocabularyLabel() bool`

HasVocabularyLabel returns a boolean if a field has been set.

### SetVocabularyLabelNil

`func (o *RwdSourceVocabulariesRow) SetVocabularyLabelNil(b bool)`

 SetVocabularyLabelNil sets the value for VocabularyLabel to be an explicit nil

### UnsetVocabularyLabel
`func (o *RwdSourceVocabulariesRow) UnsetVocabularyLabel()`

UnsetVocabularyLabel ensures that no value is present for VocabularyLabel, not even an explicit nil
### GetIsFreeText

`func (o *RwdSourceVocabulariesRow) GetIsFreeText() bool`

GetIsFreeText returns the IsFreeText field if non-nil, zero value otherwise.

### GetIsFreeTextOk

`func (o *RwdSourceVocabulariesRow) GetIsFreeTextOk() (*bool, bool)`

GetIsFreeTextOk returns a tuple with the IsFreeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeText

`func (o *RwdSourceVocabulariesRow) SetIsFreeText(v bool)`

SetIsFreeText sets IsFreeText field to given value.

### HasIsFreeText

`func (o *RwdSourceVocabulariesRow) HasIsFreeText() bool`

HasIsFreeText returns a boolean if a field has been set.

### SetIsFreeTextNil

`func (o *RwdSourceVocabulariesRow) SetIsFreeTextNil(b bool)`

 SetIsFreeTextNil sets the value for IsFreeText to be an explicit nil

### UnsetIsFreeText
`func (o *RwdSourceVocabulariesRow) UnsetIsFreeText()`

UnsetIsFreeText ensures that no value is present for IsFreeText, not even an explicit nil
### GetIsOtherSentinel

`func (o *RwdSourceVocabulariesRow) GetIsOtherSentinel() bool`

GetIsOtherSentinel returns the IsOtherSentinel field if non-nil, zero value otherwise.

### GetIsOtherSentinelOk

`func (o *RwdSourceVocabulariesRow) GetIsOtherSentinelOk() (*bool, bool)`

GetIsOtherSentinelOk returns a tuple with the IsOtherSentinel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOtherSentinel

`func (o *RwdSourceVocabulariesRow) SetIsOtherSentinel(v bool)`

SetIsOtherSentinel sets IsOtherSentinel field to given value.

### HasIsOtherSentinel

`func (o *RwdSourceVocabulariesRow) HasIsOtherSentinel() bool`

HasIsOtherSentinel returns a boolean if a field has been set.

### SetIsOtherSentinelNil

`func (o *RwdSourceVocabulariesRow) SetIsOtherSentinelNil(b bool)`

 SetIsOtherSentinelNil sets the value for IsOtherSentinel to be an explicit nil

### UnsetIsOtherSentinel
`func (o *RwdSourceVocabulariesRow) UnsetIsOtherSentinel()`

UnsetIsOtherSentinel ensures that no value is present for IsOtherSentinel, not even an explicit nil
### GetSourceValue

`func (o *RwdSourceVocabulariesRow) GetSourceValue() string`

GetSourceValue returns the SourceValue field if non-nil, zero value otherwise.

### GetSourceValueOk

`func (o *RwdSourceVocabulariesRow) GetSourceValueOk() (*string, bool)`

GetSourceValueOk returns a tuple with the SourceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValue

`func (o *RwdSourceVocabulariesRow) SetSourceValue(v string)`

SetSourceValue sets SourceValue field to given value.

### HasSourceValue

`func (o *RwdSourceVocabulariesRow) HasSourceValue() bool`

HasSourceValue returns a boolean if a field has been set.

### SetSourceValueNil

`func (o *RwdSourceVocabulariesRow) SetSourceValueNil(b bool)`

 SetSourceValueNil sets the value for SourceValue to be an explicit nil

### UnsetSourceValue
`func (o *RwdSourceVocabulariesRow) UnsetSourceValue()`

UnsetSourceValue ensures that no value is present for SourceValue, not even an explicit nil
### GetRestrictionTier

`func (o *RwdSourceVocabulariesRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdSourceVocabulariesRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdSourceVocabulariesRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdSourceVocabulariesRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdSourceVocabulariesRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdSourceVocabulariesRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdSourceVocabulariesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdSourceVocabulariesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdSourceVocabulariesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdSourceVocabulariesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdSourceVocabulariesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdSourceVocabulariesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



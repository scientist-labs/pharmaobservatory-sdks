# RwdSourceEtlMappingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | Pointer to **NullableString** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] 
**MappingOrdinal** | Pointer to **NullableInt32** | The position of the mapping among the source&#39;s declared CDM mappings, from 1. | [optional] 
**DeclaredMappingCount** | Pointer to **NullableInt32** | How many CDM mappings the holder declared, after register join artefacts are removed; 0 when the question was never answered. | [optional] 
**EtlStatus** | Pointer to **NullableString** | The extract-transform-load (ETL) status of the CDM mapping as typed; one comma-joined answer per declared mapping (rwd_source_etl_mappings has one row per mapping). | [optional] 
**EtlStatusAligned** | Pointer to **NullableBool** | False when the register list for this column splits into a different number of pieces than the number of mappings declared; the row value is then null and the whole answer is in etl_status_source. | [optional] 
**EtlStatusSource** | Pointer to **NullableString** | The whole answer the ETL status was split from. | [optional] 
**EtlCdmVersion** | Pointer to **NullableString** | The CDM version(s) declared for the ETL, comma-joined as typed. | [optional] 
**EtlCdmVersionAligned** | Pointer to **NullableBool** | False when the register list for this column splits into a different number of pieces than the number of mappings declared; the row value is then null and the whole answer is in etl_cdm_version_source. | [optional] 
**EtlCdmVersionSource** | Pointer to **NullableString** | The whole answer the CDM version was split from. | [optional] 
**EtlFrequency** | Pointer to **NullableString** | The ETL frequency declared, comma-joined as typed. | [optional] 
**EtlFrequencyAligned** | Pointer to **NullableBool** | False when the register list for this column splits into a different number of pieces than the number of mappings declared; the row value is then null and the whole answer is in etl_frequency_source. | [optional] 
**EtlFrequencySource** | Pointer to **NullableString** | The whole answer the ETL frequency was split from. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdSourceEtlMappingsRow

`func NewRwdSourceEtlMappingsRow() *RwdSourceEtlMappingsRow`

NewRwdSourceEtlMappingsRow instantiates a new RwdSourceEtlMappingsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdSourceEtlMappingsRowWithDefaults

`func NewRwdSourceEtlMappingsRowWithDefaults() *RwdSourceEtlMappingsRow`

NewRwdSourceEtlMappingsRowWithDefaults instantiates a new RwdSourceEtlMappingsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *RwdSourceEtlMappingsRow) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *RwdSourceEtlMappingsRow) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *RwdSourceEtlMappingsRow) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *RwdSourceEtlMappingsRow) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *RwdSourceEtlMappingsRow) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *RwdSourceEtlMappingsRow) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetMappingOrdinal

`func (o *RwdSourceEtlMappingsRow) GetMappingOrdinal() int32`

GetMappingOrdinal returns the MappingOrdinal field if non-nil, zero value otherwise.

### GetMappingOrdinalOk

`func (o *RwdSourceEtlMappingsRow) GetMappingOrdinalOk() (*int32, bool)`

GetMappingOrdinalOk returns a tuple with the MappingOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingOrdinal

`func (o *RwdSourceEtlMappingsRow) SetMappingOrdinal(v int32)`

SetMappingOrdinal sets MappingOrdinal field to given value.

### HasMappingOrdinal

`func (o *RwdSourceEtlMappingsRow) HasMappingOrdinal() bool`

HasMappingOrdinal returns a boolean if a field has been set.

### SetMappingOrdinalNil

`func (o *RwdSourceEtlMappingsRow) SetMappingOrdinalNil(b bool)`

 SetMappingOrdinalNil sets the value for MappingOrdinal to be an explicit nil

### UnsetMappingOrdinal
`func (o *RwdSourceEtlMappingsRow) UnsetMappingOrdinal()`

UnsetMappingOrdinal ensures that no value is present for MappingOrdinal, not even an explicit nil
### GetDeclaredMappingCount

`func (o *RwdSourceEtlMappingsRow) GetDeclaredMappingCount() int32`

GetDeclaredMappingCount returns the DeclaredMappingCount field if non-nil, zero value otherwise.

### GetDeclaredMappingCountOk

`func (o *RwdSourceEtlMappingsRow) GetDeclaredMappingCountOk() (*int32, bool)`

GetDeclaredMappingCountOk returns a tuple with the DeclaredMappingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredMappingCount

`func (o *RwdSourceEtlMappingsRow) SetDeclaredMappingCount(v int32)`

SetDeclaredMappingCount sets DeclaredMappingCount field to given value.

### HasDeclaredMappingCount

`func (o *RwdSourceEtlMappingsRow) HasDeclaredMappingCount() bool`

HasDeclaredMappingCount returns a boolean if a field has been set.

### SetDeclaredMappingCountNil

`func (o *RwdSourceEtlMappingsRow) SetDeclaredMappingCountNil(b bool)`

 SetDeclaredMappingCountNil sets the value for DeclaredMappingCount to be an explicit nil

### UnsetDeclaredMappingCount
`func (o *RwdSourceEtlMappingsRow) UnsetDeclaredMappingCount()`

UnsetDeclaredMappingCount ensures that no value is present for DeclaredMappingCount, not even an explicit nil
### GetEtlStatus

`func (o *RwdSourceEtlMappingsRow) GetEtlStatus() string`

GetEtlStatus returns the EtlStatus field if non-nil, zero value otherwise.

### GetEtlStatusOk

`func (o *RwdSourceEtlMappingsRow) GetEtlStatusOk() (*string, bool)`

GetEtlStatusOk returns a tuple with the EtlStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlStatus

`func (o *RwdSourceEtlMappingsRow) SetEtlStatus(v string)`

SetEtlStatus sets EtlStatus field to given value.

### HasEtlStatus

`func (o *RwdSourceEtlMappingsRow) HasEtlStatus() bool`

HasEtlStatus returns a boolean if a field has been set.

### SetEtlStatusNil

`func (o *RwdSourceEtlMappingsRow) SetEtlStatusNil(b bool)`

 SetEtlStatusNil sets the value for EtlStatus to be an explicit nil

### UnsetEtlStatus
`func (o *RwdSourceEtlMappingsRow) UnsetEtlStatus()`

UnsetEtlStatus ensures that no value is present for EtlStatus, not even an explicit nil
### GetEtlStatusAligned

`func (o *RwdSourceEtlMappingsRow) GetEtlStatusAligned() bool`

GetEtlStatusAligned returns the EtlStatusAligned field if non-nil, zero value otherwise.

### GetEtlStatusAlignedOk

`func (o *RwdSourceEtlMappingsRow) GetEtlStatusAlignedOk() (*bool, bool)`

GetEtlStatusAlignedOk returns a tuple with the EtlStatusAligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlStatusAligned

`func (o *RwdSourceEtlMappingsRow) SetEtlStatusAligned(v bool)`

SetEtlStatusAligned sets EtlStatusAligned field to given value.

### HasEtlStatusAligned

`func (o *RwdSourceEtlMappingsRow) HasEtlStatusAligned() bool`

HasEtlStatusAligned returns a boolean if a field has been set.

### SetEtlStatusAlignedNil

`func (o *RwdSourceEtlMappingsRow) SetEtlStatusAlignedNil(b bool)`

 SetEtlStatusAlignedNil sets the value for EtlStatusAligned to be an explicit nil

### UnsetEtlStatusAligned
`func (o *RwdSourceEtlMappingsRow) UnsetEtlStatusAligned()`

UnsetEtlStatusAligned ensures that no value is present for EtlStatusAligned, not even an explicit nil
### GetEtlStatusSource

`func (o *RwdSourceEtlMappingsRow) GetEtlStatusSource() string`

GetEtlStatusSource returns the EtlStatusSource field if non-nil, zero value otherwise.

### GetEtlStatusSourceOk

`func (o *RwdSourceEtlMappingsRow) GetEtlStatusSourceOk() (*string, bool)`

GetEtlStatusSourceOk returns a tuple with the EtlStatusSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlStatusSource

`func (o *RwdSourceEtlMappingsRow) SetEtlStatusSource(v string)`

SetEtlStatusSource sets EtlStatusSource field to given value.

### HasEtlStatusSource

`func (o *RwdSourceEtlMappingsRow) HasEtlStatusSource() bool`

HasEtlStatusSource returns a boolean if a field has been set.

### SetEtlStatusSourceNil

`func (o *RwdSourceEtlMappingsRow) SetEtlStatusSourceNil(b bool)`

 SetEtlStatusSourceNil sets the value for EtlStatusSource to be an explicit nil

### UnsetEtlStatusSource
`func (o *RwdSourceEtlMappingsRow) UnsetEtlStatusSource()`

UnsetEtlStatusSource ensures that no value is present for EtlStatusSource, not even an explicit nil
### GetEtlCdmVersion

`func (o *RwdSourceEtlMappingsRow) GetEtlCdmVersion() string`

GetEtlCdmVersion returns the EtlCdmVersion field if non-nil, zero value otherwise.

### GetEtlCdmVersionOk

`func (o *RwdSourceEtlMappingsRow) GetEtlCdmVersionOk() (*string, bool)`

GetEtlCdmVersionOk returns a tuple with the EtlCdmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlCdmVersion

`func (o *RwdSourceEtlMappingsRow) SetEtlCdmVersion(v string)`

SetEtlCdmVersion sets EtlCdmVersion field to given value.

### HasEtlCdmVersion

`func (o *RwdSourceEtlMappingsRow) HasEtlCdmVersion() bool`

HasEtlCdmVersion returns a boolean if a field has been set.

### SetEtlCdmVersionNil

`func (o *RwdSourceEtlMappingsRow) SetEtlCdmVersionNil(b bool)`

 SetEtlCdmVersionNil sets the value for EtlCdmVersion to be an explicit nil

### UnsetEtlCdmVersion
`func (o *RwdSourceEtlMappingsRow) UnsetEtlCdmVersion()`

UnsetEtlCdmVersion ensures that no value is present for EtlCdmVersion, not even an explicit nil
### GetEtlCdmVersionAligned

`func (o *RwdSourceEtlMappingsRow) GetEtlCdmVersionAligned() bool`

GetEtlCdmVersionAligned returns the EtlCdmVersionAligned field if non-nil, zero value otherwise.

### GetEtlCdmVersionAlignedOk

`func (o *RwdSourceEtlMappingsRow) GetEtlCdmVersionAlignedOk() (*bool, bool)`

GetEtlCdmVersionAlignedOk returns a tuple with the EtlCdmVersionAligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlCdmVersionAligned

`func (o *RwdSourceEtlMappingsRow) SetEtlCdmVersionAligned(v bool)`

SetEtlCdmVersionAligned sets EtlCdmVersionAligned field to given value.

### HasEtlCdmVersionAligned

`func (o *RwdSourceEtlMappingsRow) HasEtlCdmVersionAligned() bool`

HasEtlCdmVersionAligned returns a boolean if a field has been set.

### SetEtlCdmVersionAlignedNil

`func (o *RwdSourceEtlMappingsRow) SetEtlCdmVersionAlignedNil(b bool)`

 SetEtlCdmVersionAlignedNil sets the value for EtlCdmVersionAligned to be an explicit nil

### UnsetEtlCdmVersionAligned
`func (o *RwdSourceEtlMappingsRow) UnsetEtlCdmVersionAligned()`

UnsetEtlCdmVersionAligned ensures that no value is present for EtlCdmVersionAligned, not even an explicit nil
### GetEtlCdmVersionSource

`func (o *RwdSourceEtlMappingsRow) GetEtlCdmVersionSource() string`

GetEtlCdmVersionSource returns the EtlCdmVersionSource field if non-nil, zero value otherwise.

### GetEtlCdmVersionSourceOk

`func (o *RwdSourceEtlMappingsRow) GetEtlCdmVersionSourceOk() (*string, bool)`

GetEtlCdmVersionSourceOk returns a tuple with the EtlCdmVersionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlCdmVersionSource

`func (o *RwdSourceEtlMappingsRow) SetEtlCdmVersionSource(v string)`

SetEtlCdmVersionSource sets EtlCdmVersionSource field to given value.

### HasEtlCdmVersionSource

`func (o *RwdSourceEtlMappingsRow) HasEtlCdmVersionSource() bool`

HasEtlCdmVersionSource returns a boolean if a field has been set.

### SetEtlCdmVersionSourceNil

`func (o *RwdSourceEtlMappingsRow) SetEtlCdmVersionSourceNil(b bool)`

 SetEtlCdmVersionSourceNil sets the value for EtlCdmVersionSource to be an explicit nil

### UnsetEtlCdmVersionSource
`func (o *RwdSourceEtlMappingsRow) UnsetEtlCdmVersionSource()`

UnsetEtlCdmVersionSource ensures that no value is present for EtlCdmVersionSource, not even an explicit nil
### GetEtlFrequency

`func (o *RwdSourceEtlMappingsRow) GetEtlFrequency() string`

GetEtlFrequency returns the EtlFrequency field if non-nil, zero value otherwise.

### GetEtlFrequencyOk

`func (o *RwdSourceEtlMappingsRow) GetEtlFrequencyOk() (*string, bool)`

GetEtlFrequencyOk returns a tuple with the EtlFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlFrequency

`func (o *RwdSourceEtlMappingsRow) SetEtlFrequency(v string)`

SetEtlFrequency sets EtlFrequency field to given value.

### HasEtlFrequency

`func (o *RwdSourceEtlMappingsRow) HasEtlFrequency() bool`

HasEtlFrequency returns a boolean if a field has been set.

### SetEtlFrequencyNil

`func (o *RwdSourceEtlMappingsRow) SetEtlFrequencyNil(b bool)`

 SetEtlFrequencyNil sets the value for EtlFrequency to be an explicit nil

### UnsetEtlFrequency
`func (o *RwdSourceEtlMappingsRow) UnsetEtlFrequency()`

UnsetEtlFrequency ensures that no value is present for EtlFrequency, not even an explicit nil
### GetEtlFrequencyAligned

`func (o *RwdSourceEtlMappingsRow) GetEtlFrequencyAligned() bool`

GetEtlFrequencyAligned returns the EtlFrequencyAligned field if non-nil, zero value otherwise.

### GetEtlFrequencyAlignedOk

`func (o *RwdSourceEtlMappingsRow) GetEtlFrequencyAlignedOk() (*bool, bool)`

GetEtlFrequencyAlignedOk returns a tuple with the EtlFrequencyAligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlFrequencyAligned

`func (o *RwdSourceEtlMappingsRow) SetEtlFrequencyAligned(v bool)`

SetEtlFrequencyAligned sets EtlFrequencyAligned field to given value.

### HasEtlFrequencyAligned

`func (o *RwdSourceEtlMappingsRow) HasEtlFrequencyAligned() bool`

HasEtlFrequencyAligned returns a boolean if a field has been set.

### SetEtlFrequencyAlignedNil

`func (o *RwdSourceEtlMappingsRow) SetEtlFrequencyAlignedNil(b bool)`

 SetEtlFrequencyAlignedNil sets the value for EtlFrequencyAligned to be an explicit nil

### UnsetEtlFrequencyAligned
`func (o *RwdSourceEtlMappingsRow) UnsetEtlFrequencyAligned()`

UnsetEtlFrequencyAligned ensures that no value is present for EtlFrequencyAligned, not even an explicit nil
### GetEtlFrequencySource

`func (o *RwdSourceEtlMappingsRow) GetEtlFrequencySource() string`

GetEtlFrequencySource returns the EtlFrequencySource field if non-nil, zero value otherwise.

### GetEtlFrequencySourceOk

`func (o *RwdSourceEtlMappingsRow) GetEtlFrequencySourceOk() (*string, bool)`

GetEtlFrequencySourceOk returns a tuple with the EtlFrequencySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlFrequencySource

`func (o *RwdSourceEtlMappingsRow) SetEtlFrequencySource(v string)`

SetEtlFrequencySource sets EtlFrequencySource field to given value.

### HasEtlFrequencySource

`func (o *RwdSourceEtlMappingsRow) HasEtlFrequencySource() bool`

HasEtlFrequencySource returns a boolean if a field has been set.

### SetEtlFrequencySourceNil

`func (o *RwdSourceEtlMappingsRow) SetEtlFrequencySourceNil(b bool)`

 SetEtlFrequencySourceNil sets the value for EtlFrequencySource to be an explicit nil

### UnsetEtlFrequencySource
`func (o *RwdSourceEtlMappingsRow) UnsetEtlFrequencySource()`

UnsetEtlFrequencySource ensures that no value is present for EtlFrequencySource, not even an explicit nil
### GetRestrictionTier

`func (o *RwdSourceEtlMappingsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdSourceEtlMappingsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdSourceEtlMappingsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdSourceEtlMappingsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdSourceEtlMappingsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdSourceEtlMappingsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdSourceEtlMappingsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdSourceEtlMappingsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdSourceEtlMappingsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdSourceEtlMappingsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdSourceEtlMappingsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdSourceEtlMappingsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



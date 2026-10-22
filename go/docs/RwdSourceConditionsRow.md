# RwdSourceConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | Pointer to **NullableString** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] 
**DataSourceNameKey** | Pointer to **NullableString** | The normalised name key the study register is joined to the data source on (the study register carries no data source id). | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition&#39;s MeSH descriptor id (D followed by six to nine digits), the id of the conditions object type. | [optional] 
**ConditionLabel** | Pointer to **NullableString** | The condition&#39;s MeSH heading. | [optional] 
**MeshTreeCategory** | Pointer to **NullableString** | The MeSH tree category the condition sits under. | [optional] 
**MeshTreePrefix** | Pointer to **NullableString** | The MeSH tree prefix of that category. | [optional] 
**IsDiseaseDescriptor** | Pointer to **NullableBool** | Whether the MeSH descriptor is a disease descriptor. | [optional] 
**EmaConditionTerms** | Pointer to **[]string** | The register&#39;s own condition terms that resolved to this MeSH descriptor. | [optional] 
**EmaConditionTermCount** | Pointer to **NullableInt32** | The number of register condition terms behind the row. | [optional] 
**StudyCount** | Pointer to **NullableInt32** | How many registered studies named this source. A registered study NAMED the source; it is never evidence that the source holds data on a condition. Read it with study_register_total and studies_naming_a_registered_source, because a zero is not a finding that the source is unused. | [optional] 
**LinkedStudyIds** | Pointer to **[]string** | The ids of the registered studies behind the row (studies of this condition that named this source). | [optional] 
**SourceStudyCount** | Pointer to **NullableInt32** | How many registered studies named this source, whatever their condition: the denominator under study_count on this row. | [optional] 
**SourceConditionStudyCount** | Pointer to **NullableInt32** | How many distinct registered studies of this source carry any MeSH-resolved condition, whatever the condition. | [optional] 
**SourceConditionCount** | Pointer to **NullableInt32** | How many distinct conditions have a row for this source. | [optional] 
**RegisterSourceCount** | Pointer to **NullableInt32** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] 
**SourcesWithAnyStudyLink** | Pointer to **NullableInt32** | How many registered sources are named by at least one registered study. | [optional] 
**SourcesWithAnyConditionLink** | Pointer to **NullableInt32** | How many registered sources have at least one condition row (named by a study with a resolved condition). | [optional] 
**StudyRegisterTotal** | Pointer to **NullableInt32** | The number of studies in the whole study register at the snapshot. | [optional] 
**StudiesNamingARegisteredSource** | Pointer to **NullableInt32** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] 
**ConditionTermTotal** | Pointer to **NullableInt32** | How many distinct condition terms the study register holds. | [optional] 
**ConditionTermResolved** | Pointer to **NullableInt32** | How many of those terms resolved to a MeSH descriptor. | [optional] 
**LinkedConditionCount** | Pointer to **NullableInt32** | How many distinct MeSH conditions appear on any row of the dataset. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdSourceConditionsRow

`func NewRwdSourceConditionsRow() *RwdSourceConditionsRow`

NewRwdSourceConditionsRow instantiates a new RwdSourceConditionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdSourceConditionsRowWithDefaults

`func NewRwdSourceConditionsRowWithDefaults() *RwdSourceConditionsRow`

NewRwdSourceConditionsRowWithDefaults instantiates a new RwdSourceConditionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *RwdSourceConditionsRow) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *RwdSourceConditionsRow) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *RwdSourceConditionsRow) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *RwdSourceConditionsRow) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *RwdSourceConditionsRow) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *RwdSourceConditionsRow) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetDataSourceNameKey

`func (o *RwdSourceConditionsRow) GetDataSourceNameKey() string`

GetDataSourceNameKey returns the DataSourceNameKey field if non-nil, zero value otherwise.

### GetDataSourceNameKeyOk

`func (o *RwdSourceConditionsRow) GetDataSourceNameKeyOk() (*string, bool)`

GetDataSourceNameKeyOk returns a tuple with the DataSourceNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceNameKey

`func (o *RwdSourceConditionsRow) SetDataSourceNameKey(v string)`

SetDataSourceNameKey sets DataSourceNameKey field to given value.

### HasDataSourceNameKey

`func (o *RwdSourceConditionsRow) HasDataSourceNameKey() bool`

HasDataSourceNameKey returns a boolean if a field has been set.

### SetDataSourceNameKeyNil

`func (o *RwdSourceConditionsRow) SetDataSourceNameKeyNil(b bool)`

 SetDataSourceNameKeyNil sets the value for DataSourceNameKey to be an explicit nil

### UnsetDataSourceNameKey
`func (o *RwdSourceConditionsRow) UnsetDataSourceNameKey()`

UnsetDataSourceNameKey ensures that no value is present for DataSourceNameKey, not even an explicit nil
### GetConditionId

`func (o *RwdSourceConditionsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *RwdSourceConditionsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *RwdSourceConditionsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *RwdSourceConditionsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *RwdSourceConditionsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *RwdSourceConditionsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionLabel

`func (o *RwdSourceConditionsRow) GetConditionLabel() string`

GetConditionLabel returns the ConditionLabel field if non-nil, zero value otherwise.

### GetConditionLabelOk

`func (o *RwdSourceConditionsRow) GetConditionLabelOk() (*string, bool)`

GetConditionLabelOk returns a tuple with the ConditionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionLabel

`func (o *RwdSourceConditionsRow) SetConditionLabel(v string)`

SetConditionLabel sets ConditionLabel field to given value.

### HasConditionLabel

`func (o *RwdSourceConditionsRow) HasConditionLabel() bool`

HasConditionLabel returns a boolean if a field has been set.

### SetConditionLabelNil

`func (o *RwdSourceConditionsRow) SetConditionLabelNil(b bool)`

 SetConditionLabelNil sets the value for ConditionLabel to be an explicit nil

### UnsetConditionLabel
`func (o *RwdSourceConditionsRow) UnsetConditionLabel()`

UnsetConditionLabel ensures that no value is present for ConditionLabel, not even an explicit nil
### GetMeshTreeCategory

`func (o *RwdSourceConditionsRow) GetMeshTreeCategory() string`

GetMeshTreeCategory returns the MeshTreeCategory field if non-nil, zero value otherwise.

### GetMeshTreeCategoryOk

`func (o *RwdSourceConditionsRow) GetMeshTreeCategoryOk() (*string, bool)`

GetMeshTreeCategoryOk returns a tuple with the MeshTreeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshTreeCategory

`func (o *RwdSourceConditionsRow) SetMeshTreeCategory(v string)`

SetMeshTreeCategory sets MeshTreeCategory field to given value.

### HasMeshTreeCategory

`func (o *RwdSourceConditionsRow) HasMeshTreeCategory() bool`

HasMeshTreeCategory returns a boolean if a field has been set.

### SetMeshTreeCategoryNil

`func (o *RwdSourceConditionsRow) SetMeshTreeCategoryNil(b bool)`

 SetMeshTreeCategoryNil sets the value for MeshTreeCategory to be an explicit nil

### UnsetMeshTreeCategory
`func (o *RwdSourceConditionsRow) UnsetMeshTreeCategory()`

UnsetMeshTreeCategory ensures that no value is present for MeshTreeCategory, not even an explicit nil
### GetMeshTreePrefix

`func (o *RwdSourceConditionsRow) GetMeshTreePrefix() string`

GetMeshTreePrefix returns the MeshTreePrefix field if non-nil, zero value otherwise.

### GetMeshTreePrefixOk

`func (o *RwdSourceConditionsRow) GetMeshTreePrefixOk() (*string, bool)`

GetMeshTreePrefixOk returns a tuple with the MeshTreePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshTreePrefix

`func (o *RwdSourceConditionsRow) SetMeshTreePrefix(v string)`

SetMeshTreePrefix sets MeshTreePrefix field to given value.

### HasMeshTreePrefix

`func (o *RwdSourceConditionsRow) HasMeshTreePrefix() bool`

HasMeshTreePrefix returns a boolean if a field has been set.

### SetMeshTreePrefixNil

`func (o *RwdSourceConditionsRow) SetMeshTreePrefixNil(b bool)`

 SetMeshTreePrefixNil sets the value for MeshTreePrefix to be an explicit nil

### UnsetMeshTreePrefix
`func (o *RwdSourceConditionsRow) UnsetMeshTreePrefix()`

UnsetMeshTreePrefix ensures that no value is present for MeshTreePrefix, not even an explicit nil
### GetIsDiseaseDescriptor

`func (o *RwdSourceConditionsRow) GetIsDiseaseDescriptor() bool`

GetIsDiseaseDescriptor returns the IsDiseaseDescriptor field if non-nil, zero value otherwise.

### GetIsDiseaseDescriptorOk

`func (o *RwdSourceConditionsRow) GetIsDiseaseDescriptorOk() (*bool, bool)`

GetIsDiseaseDescriptorOk returns a tuple with the IsDiseaseDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiseaseDescriptor

`func (o *RwdSourceConditionsRow) SetIsDiseaseDescriptor(v bool)`

SetIsDiseaseDescriptor sets IsDiseaseDescriptor field to given value.

### HasIsDiseaseDescriptor

`func (o *RwdSourceConditionsRow) HasIsDiseaseDescriptor() bool`

HasIsDiseaseDescriptor returns a boolean if a field has been set.

### SetIsDiseaseDescriptorNil

`func (o *RwdSourceConditionsRow) SetIsDiseaseDescriptorNil(b bool)`

 SetIsDiseaseDescriptorNil sets the value for IsDiseaseDescriptor to be an explicit nil

### UnsetIsDiseaseDescriptor
`func (o *RwdSourceConditionsRow) UnsetIsDiseaseDescriptor()`

UnsetIsDiseaseDescriptor ensures that no value is present for IsDiseaseDescriptor, not even an explicit nil
### GetEmaConditionTerms

`func (o *RwdSourceConditionsRow) GetEmaConditionTerms() []string`

GetEmaConditionTerms returns the EmaConditionTerms field if non-nil, zero value otherwise.

### GetEmaConditionTermsOk

`func (o *RwdSourceConditionsRow) GetEmaConditionTermsOk() (*[]string, bool)`

GetEmaConditionTermsOk returns a tuple with the EmaConditionTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmaConditionTerms

`func (o *RwdSourceConditionsRow) SetEmaConditionTerms(v []string)`

SetEmaConditionTerms sets EmaConditionTerms field to given value.

### HasEmaConditionTerms

`func (o *RwdSourceConditionsRow) HasEmaConditionTerms() bool`

HasEmaConditionTerms returns a boolean if a field has been set.

### SetEmaConditionTermsNil

`func (o *RwdSourceConditionsRow) SetEmaConditionTermsNil(b bool)`

 SetEmaConditionTermsNil sets the value for EmaConditionTerms to be an explicit nil

### UnsetEmaConditionTerms
`func (o *RwdSourceConditionsRow) UnsetEmaConditionTerms()`

UnsetEmaConditionTerms ensures that no value is present for EmaConditionTerms, not even an explicit nil
### GetEmaConditionTermCount

`func (o *RwdSourceConditionsRow) GetEmaConditionTermCount() int32`

GetEmaConditionTermCount returns the EmaConditionTermCount field if non-nil, zero value otherwise.

### GetEmaConditionTermCountOk

`func (o *RwdSourceConditionsRow) GetEmaConditionTermCountOk() (*int32, bool)`

GetEmaConditionTermCountOk returns a tuple with the EmaConditionTermCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmaConditionTermCount

`func (o *RwdSourceConditionsRow) SetEmaConditionTermCount(v int32)`

SetEmaConditionTermCount sets EmaConditionTermCount field to given value.

### HasEmaConditionTermCount

`func (o *RwdSourceConditionsRow) HasEmaConditionTermCount() bool`

HasEmaConditionTermCount returns a boolean if a field has been set.

### SetEmaConditionTermCountNil

`func (o *RwdSourceConditionsRow) SetEmaConditionTermCountNil(b bool)`

 SetEmaConditionTermCountNil sets the value for EmaConditionTermCount to be an explicit nil

### UnsetEmaConditionTermCount
`func (o *RwdSourceConditionsRow) UnsetEmaConditionTermCount()`

UnsetEmaConditionTermCount ensures that no value is present for EmaConditionTermCount, not even an explicit nil
### GetStudyCount

`func (o *RwdSourceConditionsRow) GetStudyCount() int32`

GetStudyCount returns the StudyCount field if non-nil, zero value otherwise.

### GetStudyCountOk

`func (o *RwdSourceConditionsRow) GetStudyCountOk() (*int32, bool)`

GetStudyCountOk returns a tuple with the StudyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyCount

`func (o *RwdSourceConditionsRow) SetStudyCount(v int32)`

SetStudyCount sets StudyCount field to given value.

### HasStudyCount

`func (o *RwdSourceConditionsRow) HasStudyCount() bool`

HasStudyCount returns a boolean if a field has been set.

### SetStudyCountNil

`func (o *RwdSourceConditionsRow) SetStudyCountNil(b bool)`

 SetStudyCountNil sets the value for StudyCount to be an explicit nil

### UnsetStudyCount
`func (o *RwdSourceConditionsRow) UnsetStudyCount()`

UnsetStudyCount ensures that no value is present for StudyCount, not even an explicit nil
### GetLinkedStudyIds

`func (o *RwdSourceConditionsRow) GetLinkedStudyIds() []string`

GetLinkedStudyIds returns the LinkedStudyIds field if non-nil, zero value otherwise.

### GetLinkedStudyIdsOk

`func (o *RwdSourceConditionsRow) GetLinkedStudyIdsOk() (*[]string, bool)`

GetLinkedStudyIdsOk returns a tuple with the LinkedStudyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedStudyIds

`func (o *RwdSourceConditionsRow) SetLinkedStudyIds(v []string)`

SetLinkedStudyIds sets LinkedStudyIds field to given value.

### HasLinkedStudyIds

`func (o *RwdSourceConditionsRow) HasLinkedStudyIds() bool`

HasLinkedStudyIds returns a boolean if a field has been set.

### SetLinkedStudyIdsNil

`func (o *RwdSourceConditionsRow) SetLinkedStudyIdsNil(b bool)`

 SetLinkedStudyIdsNil sets the value for LinkedStudyIds to be an explicit nil

### UnsetLinkedStudyIds
`func (o *RwdSourceConditionsRow) UnsetLinkedStudyIds()`

UnsetLinkedStudyIds ensures that no value is present for LinkedStudyIds, not even an explicit nil
### GetSourceStudyCount

`func (o *RwdSourceConditionsRow) GetSourceStudyCount() int32`

GetSourceStudyCount returns the SourceStudyCount field if non-nil, zero value otherwise.

### GetSourceStudyCountOk

`func (o *RwdSourceConditionsRow) GetSourceStudyCountOk() (*int32, bool)`

GetSourceStudyCountOk returns a tuple with the SourceStudyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStudyCount

`func (o *RwdSourceConditionsRow) SetSourceStudyCount(v int32)`

SetSourceStudyCount sets SourceStudyCount field to given value.

### HasSourceStudyCount

`func (o *RwdSourceConditionsRow) HasSourceStudyCount() bool`

HasSourceStudyCount returns a boolean if a field has been set.

### SetSourceStudyCountNil

`func (o *RwdSourceConditionsRow) SetSourceStudyCountNil(b bool)`

 SetSourceStudyCountNil sets the value for SourceStudyCount to be an explicit nil

### UnsetSourceStudyCount
`func (o *RwdSourceConditionsRow) UnsetSourceStudyCount()`

UnsetSourceStudyCount ensures that no value is present for SourceStudyCount, not even an explicit nil
### GetSourceConditionStudyCount

`func (o *RwdSourceConditionsRow) GetSourceConditionStudyCount() int32`

GetSourceConditionStudyCount returns the SourceConditionStudyCount field if non-nil, zero value otherwise.

### GetSourceConditionStudyCountOk

`func (o *RwdSourceConditionsRow) GetSourceConditionStudyCountOk() (*int32, bool)`

GetSourceConditionStudyCountOk returns a tuple with the SourceConditionStudyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConditionStudyCount

`func (o *RwdSourceConditionsRow) SetSourceConditionStudyCount(v int32)`

SetSourceConditionStudyCount sets SourceConditionStudyCount field to given value.

### HasSourceConditionStudyCount

`func (o *RwdSourceConditionsRow) HasSourceConditionStudyCount() bool`

HasSourceConditionStudyCount returns a boolean if a field has been set.

### SetSourceConditionStudyCountNil

`func (o *RwdSourceConditionsRow) SetSourceConditionStudyCountNil(b bool)`

 SetSourceConditionStudyCountNil sets the value for SourceConditionStudyCount to be an explicit nil

### UnsetSourceConditionStudyCount
`func (o *RwdSourceConditionsRow) UnsetSourceConditionStudyCount()`

UnsetSourceConditionStudyCount ensures that no value is present for SourceConditionStudyCount, not even an explicit nil
### GetSourceConditionCount

`func (o *RwdSourceConditionsRow) GetSourceConditionCount() int32`

GetSourceConditionCount returns the SourceConditionCount field if non-nil, zero value otherwise.

### GetSourceConditionCountOk

`func (o *RwdSourceConditionsRow) GetSourceConditionCountOk() (*int32, bool)`

GetSourceConditionCountOk returns a tuple with the SourceConditionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConditionCount

`func (o *RwdSourceConditionsRow) SetSourceConditionCount(v int32)`

SetSourceConditionCount sets SourceConditionCount field to given value.

### HasSourceConditionCount

`func (o *RwdSourceConditionsRow) HasSourceConditionCount() bool`

HasSourceConditionCount returns a boolean if a field has been set.

### SetSourceConditionCountNil

`func (o *RwdSourceConditionsRow) SetSourceConditionCountNil(b bool)`

 SetSourceConditionCountNil sets the value for SourceConditionCount to be an explicit nil

### UnsetSourceConditionCount
`func (o *RwdSourceConditionsRow) UnsetSourceConditionCount()`

UnsetSourceConditionCount ensures that no value is present for SourceConditionCount, not even an explicit nil
### GetRegisterSourceCount

`func (o *RwdSourceConditionsRow) GetRegisterSourceCount() int32`

GetRegisterSourceCount returns the RegisterSourceCount field if non-nil, zero value otherwise.

### GetRegisterSourceCountOk

`func (o *RwdSourceConditionsRow) GetRegisterSourceCountOk() (*int32, bool)`

GetRegisterSourceCountOk returns a tuple with the RegisterSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterSourceCount

`func (o *RwdSourceConditionsRow) SetRegisterSourceCount(v int32)`

SetRegisterSourceCount sets RegisterSourceCount field to given value.

### HasRegisterSourceCount

`func (o *RwdSourceConditionsRow) HasRegisterSourceCount() bool`

HasRegisterSourceCount returns a boolean if a field has been set.

### SetRegisterSourceCountNil

`func (o *RwdSourceConditionsRow) SetRegisterSourceCountNil(b bool)`

 SetRegisterSourceCountNil sets the value for RegisterSourceCount to be an explicit nil

### UnsetRegisterSourceCount
`func (o *RwdSourceConditionsRow) UnsetRegisterSourceCount()`

UnsetRegisterSourceCount ensures that no value is present for RegisterSourceCount, not even an explicit nil
### GetSourcesWithAnyStudyLink

`func (o *RwdSourceConditionsRow) GetSourcesWithAnyStudyLink() int32`

GetSourcesWithAnyStudyLink returns the SourcesWithAnyStudyLink field if non-nil, zero value otherwise.

### GetSourcesWithAnyStudyLinkOk

`func (o *RwdSourceConditionsRow) GetSourcesWithAnyStudyLinkOk() (*int32, bool)`

GetSourcesWithAnyStudyLinkOk returns a tuple with the SourcesWithAnyStudyLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesWithAnyStudyLink

`func (o *RwdSourceConditionsRow) SetSourcesWithAnyStudyLink(v int32)`

SetSourcesWithAnyStudyLink sets SourcesWithAnyStudyLink field to given value.

### HasSourcesWithAnyStudyLink

`func (o *RwdSourceConditionsRow) HasSourcesWithAnyStudyLink() bool`

HasSourcesWithAnyStudyLink returns a boolean if a field has been set.

### SetSourcesWithAnyStudyLinkNil

`func (o *RwdSourceConditionsRow) SetSourcesWithAnyStudyLinkNil(b bool)`

 SetSourcesWithAnyStudyLinkNil sets the value for SourcesWithAnyStudyLink to be an explicit nil

### UnsetSourcesWithAnyStudyLink
`func (o *RwdSourceConditionsRow) UnsetSourcesWithAnyStudyLink()`

UnsetSourcesWithAnyStudyLink ensures that no value is present for SourcesWithAnyStudyLink, not even an explicit nil
### GetSourcesWithAnyConditionLink

`func (o *RwdSourceConditionsRow) GetSourcesWithAnyConditionLink() int32`

GetSourcesWithAnyConditionLink returns the SourcesWithAnyConditionLink field if non-nil, zero value otherwise.

### GetSourcesWithAnyConditionLinkOk

`func (o *RwdSourceConditionsRow) GetSourcesWithAnyConditionLinkOk() (*int32, bool)`

GetSourcesWithAnyConditionLinkOk returns a tuple with the SourcesWithAnyConditionLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesWithAnyConditionLink

`func (o *RwdSourceConditionsRow) SetSourcesWithAnyConditionLink(v int32)`

SetSourcesWithAnyConditionLink sets SourcesWithAnyConditionLink field to given value.

### HasSourcesWithAnyConditionLink

`func (o *RwdSourceConditionsRow) HasSourcesWithAnyConditionLink() bool`

HasSourcesWithAnyConditionLink returns a boolean if a field has been set.

### SetSourcesWithAnyConditionLinkNil

`func (o *RwdSourceConditionsRow) SetSourcesWithAnyConditionLinkNil(b bool)`

 SetSourcesWithAnyConditionLinkNil sets the value for SourcesWithAnyConditionLink to be an explicit nil

### UnsetSourcesWithAnyConditionLink
`func (o *RwdSourceConditionsRow) UnsetSourcesWithAnyConditionLink()`

UnsetSourcesWithAnyConditionLink ensures that no value is present for SourcesWithAnyConditionLink, not even an explicit nil
### GetStudyRegisterTotal

`func (o *RwdSourceConditionsRow) GetStudyRegisterTotal() int32`

GetStudyRegisterTotal returns the StudyRegisterTotal field if non-nil, zero value otherwise.

### GetStudyRegisterTotalOk

`func (o *RwdSourceConditionsRow) GetStudyRegisterTotalOk() (*int32, bool)`

GetStudyRegisterTotalOk returns a tuple with the StudyRegisterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyRegisterTotal

`func (o *RwdSourceConditionsRow) SetStudyRegisterTotal(v int32)`

SetStudyRegisterTotal sets StudyRegisterTotal field to given value.

### HasStudyRegisterTotal

`func (o *RwdSourceConditionsRow) HasStudyRegisterTotal() bool`

HasStudyRegisterTotal returns a boolean if a field has been set.

### SetStudyRegisterTotalNil

`func (o *RwdSourceConditionsRow) SetStudyRegisterTotalNil(b bool)`

 SetStudyRegisterTotalNil sets the value for StudyRegisterTotal to be an explicit nil

### UnsetStudyRegisterTotal
`func (o *RwdSourceConditionsRow) UnsetStudyRegisterTotal()`

UnsetStudyRegisterTotal ensures that no value is present for StudyRegisterTotal, not even an explicit nil
### GetStudiesNamingARegisteredSource

`func (o *RwdSourceConditionsRow) GetStudiesNamingARegisteredSource() int32`

GetStudiesNamingARegisteredSource returns the StudiesNamingARegisteredSource field if non-nil, zero value otherwise.

### GetStudiesNamingARegisteredSourceOk

`func (o *RwdSourceConditionsRow) GetStudiesNamingARegisteredSourceOk() (*int32, bool)`

GetStudiesNamingARegisteredSourceOk returns a tuple with the StudiesNamingARegisteredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudiesNamingARegisteredSource

`func (o *RwdSourceConditionsRow) SetStudiesNamingARegisteredSource(v int32)`

SetStudiesNamingARegisteredSource sets StudiesNamingARegisteredSource field to given value.

### HasStudiesNamingARegisteredSource

`func (o *RwdSourceConditionsRow) HasStudiesNamingARegisteredSource() bool`

HasStudiesNamingARegisteredSource returns a boolean if a field has been set.

### SetStudiesNamingARegisteredSourceNil

`func (o *RwdSourceConditionsRow) SetStudiesNamingARegisteredSourceNil(b bool)`

 SetStudiesNamingARegisteredSourceNil sets the value for StudiesNamingARegisteredSource to be an explicit nil

### UnsetStudiesNamingARegisteredSource
`func (o *RwdSourceConditionsRow) UnsetStudiesNamingARegisteredSource()`

UnsetStudiesNamingARegisteredSource ensures that no value is present for StudiesNamingARegisteredSource, not even an explicit nil
### GetConditionTermTotal

`func (o *RwdSourceConditionsRow) GetConditionTermTotal() int32`

GetConditionTermTotal returns the ConditionTermTotal field if non-nil, zero value otherwise.

### GetConditionTermTotalOk

`func (o *RwdSourceConditionsRow) GetConditionTermTotalOk() (*int32, bool)`

GetConditionTermTotalOk returns a tuple with the ConditionTermTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTermTotal

`func (o *RwdSourceConditionsRow) SetConditionTermTotal(v int32)`

SetConditionTermTotal sets ConditionTermTotal field to given value.

### HasConditionTermTotal

`func (o *RwdSourceConditionsRow) HasConditionTermTotal() bool`

HasConditionTermTotal returns a boolean if a field has been set.

### SetConditionTermTotalNil

`func (o *RwdSourceConditionsRow) SetConditionTermTotalNil(b bool)`

 SetConditionTermTotalNil sets the value for ConditionTermTotal to be an explicit nil

### UnsetConditionTermTotal
`func (o *RwdSourceConditionsRow) UnsetConditionTermTotal()`

UnsetConditionTermTotal ensures that no value is present for ConditionTermTotal, not even an explicit nil
### GetConditionTermResolved

`func (o *RwdSourceConditionsRow) GetConditionTermResolved() int32`

GetConditionTermResolved returns the ConditionTermResolved field if non-nil, zero value otherwise.

### GetConditionTermResolvedOk

`func (o *RwdSourceConditionsRow) GetConditionTermResolvedOk() (*int32, bool)`

GetConditionTermResolvedOk returns a tuple with the ConditionTermResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTermResolved

`func (o *RwdSourceConditionsRow) SetConditionTermResolved(v int32)`

SetConditionTermResolved sets ConditionTermResolved field to given value.

### HasConditionTermResolved

`func (o *RwdSourceConditionsRow) HasConditionTermResolved() bool`

HasConditionTermResolved returns a boolean if a field has been set.

### SetConditionTermResolvedNil

`func (o *RwdSourceConditionsRow) SetConditionTermResolvedNil(b bool)`

 SetConditionTermResolvedNil sets the value for ConditionTermResolved to be an explicit nil

### UnsetConditionTermResolved
`func (o *RwdSourceConditionsRow) UnsetConditionTermResolved()`

UnsetConditionTermResolved ensures that no value is present for ConditionTermResolved, not even an explicit nil
### GetLinkedConditionCount

`func (o *RwdSourceConditionsRow) GetLinkedConditionCount() int32`

GetLinkedConditionCount returns the LinkedConditionCount field if non-nil, zero value otherwise.

### GetLinkedConditionCountOk

`func (o *RwdSourceConditionsRow) GetLinkedConditionCountOk() (*int32, bool)`

GetLinkedConditionCountOk returns a tuple with the LinkedConditionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedConditionCount

`func (o *RwdSourceConditionsRow) SetLinkedConditionCount(v int32)`

SetLinkedConditionCount sets LinkedConditionCount field to given value.

### HasLinkedConditionCount

`func (o *RwdSourceConditionsRow) HasLinkedConditionCount() bool`

HasLinkedConditionCount returns a boolean if a field has been set.

### SetLinkedConditionCountNil

`func (o *RwdSourceConditionsRow) SetLinkedConditionCountNil(b bool)`

 SetLinkedConditionCountNil sets the value for LinkedConditionCount to be an explicit nil

### UnsetLinkedConditionCount
`func (o *RwdSourceConditionsRow) UnsetLinkedConditionCount()`

UnsetLinkedConditionCount ensures that no value is present for LinkedConditionCount, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdSourceConditionsRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdSourceConditionsRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdSourceConditionsRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdSourceConditionsRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdSourceConditionsRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdSourceConditionsRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdSourceConditionsRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdSourceConditionsRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdSourceConditionsRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdSourceConditionsRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdSourceConditionsRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdSourceConditionsRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdSourceConditionsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdSourceConditionsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdSourceConditionsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdSourceConditionsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdSourceConditionsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdSourceConditionsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdSourceConditionsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdSourceConditionsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdSourceConditionsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdSourceConditionsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdSourceConditionsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdSourceConditionsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



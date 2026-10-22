# RwdSourceLinkagesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | Pointer to **NullableString** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] 
**LinkageOrdinal** | Pointer to **NullableInt32** | The position of the linkage among the source&#39;s declared linkages, from 1. | [optional] 
**DeclaredLinkageCount** | Pointer to **NullableInt32** | How many linkages the holder declared; 0 when none. | [optional] 
**IsPreLinked** | Pointer to **NullableBool** | Whether the holder answered Yes to Pre linked for this linkage. | [optional] 
**PreLinkedAnswer** | Pointer to **NullableString** | The holder&#39;s Yes or No answer to Pre linked for this linkage. | [optional] 
**LinkedDataSourceName** | Pointer to **NullableString** | The linked data source named for this linkage, from the register&#39;s list. | [optional] 
**LinkedDataSourceNameAligned** | Pointer to **NullableBool** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linked_data_source_name_source. | [optional] 
**LinkedDataSourceNameSource** | Pointer to **NullableString** | The whole comma-joined answer the linked data source name was split from, kept so an unaligned row can still be read. | [optional] 
**NamesOtherSentinel** | Pointer to **NullableBool** | True when the linked data source name is the register&#39;s Other data source sentinel. | [optional] 
**NamesRegisteredSource** | Pointer to **NullableBool** | True when the linked data source name is itself a registered data source. | [optional] 
**LinkedDataSourceOther** | Pointer to **NullableString** | The linked data source named in free text when the holder chose Other data source. | [optional] 
**LinkedDataSourceOtherAligned** | Pointer to **NullableBool** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linked_data_source_other_source. | [optional] 
**LinkedDataSourceOtherSource** | Pointer to **NullableString** | The whole answer the other-data-source text was split from. | [optional] 
**LinkageCompleteness** | Pointer to **NullableString** | The linkage completeness declared for this linkage, as typed. | [optional] 
**LinkageCompletenessAligned** | Pointer to **NullableBool** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linkage_completeness_source. | [optional] 
**LinkageCompletenessSource** | Pointer to **NullableString** | The whole answer linkage completeness was split from. | [optional] 
**LinkageVariable** | Pointer to **NullableString** | The linkage variable declared for this linkage, as typed. | [optional] 
**LinkageVariableAligned** | Pointer to **NullableBool** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linkage_variable_source. | [optional] 
**LinkageVariableSource** | Pointer to **NullableString** | The whole answer the linkage variable was split from. | [optional] 
**LinkageStrategy** | Pointer to **NullableString** | The linkage strategy declared for this linkage, as typed. | [optional] 
**LinkageStrategyAligned** | Pointer to **NullableBool** | False when the register list for this column splits into a different number of pieces than the number of linkages declared; the row value is then null and the whole answer is in linkage_strategy_source. | [optional] 
**LinkageStrategySource** | Pointer to **NullableString** | The whole answer the linkage strategy was split from. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdSourceLinkagesRow

`func NewRwdSourceLinkagesRow() *RwdSourceLinkagesRow`

NewRwdSourceLinkagesRow instantiates a new RwdSourceLinkagesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdSourceLinkagesRowWithDefaults

`func NewRwdSourceLinkagesRowWithDefaults() *RwdSourceLinkagesRow`

NewRwdSourceLinkagesRowWithDefaults instantiates a new RwdSourceLinkagesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *RwdSourceLinkagesRow) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *RwdSourceLinkagesRow) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *RwdSourceLinkagesRow) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *RwdSourceLinkagesRow) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *RwdSourceLinkagesRow) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *RwdSourceLinkagesRow) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetLinkageOrdinal

`func (o *RwdSourceLinkagesRow) GetLinkageOrdinal() int32`

GetLinkageOrdinal returns the LinkageOrdinal field if non-nil, zero value otherwise.

### GetLinkageOrdinalOk

`func (o *RwdSourceLinkagesRow) GetLinkageOrdinalOk() (*int32, bool)`

GetLinkageOrdinalOk returns a tuple with the LinkageOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageOrdinal

`func (o *RwdSourceLinkagesRow) SetLinkageOrdinal(v int32)`

SetLinkageOrdinal sets LinkageOrdinal field to given value.

### HasLinkageOrdinal

`func (o *RwdSourceLinkagesRow) HasLinkageOrdinal() bool`

HasLinkageOrdinal returns a boolean if a field has been set.

### SetLinkageOrdinalNil

`func (o *RwdSourceLinkagesRow) SetLinkageOrdinalNil(b bool)`

 SetLinkageOrdinalNil sets the value for LinkageOrdinal to be an explicit nil

### UnsetLinkageOrdinal
`func (o *RwdSourceLinkagesRow) UnsetLinkageOrdinal()`

UnsetLinkageOrdinal ensures that no value is present for LinkageOrdinal, not even an explicit nil
### GetDeclaredLinkageCount

`func (o *RwdSourceLinkagesRow) GetDeclaredLinkageCount() int32`

GetDeclaredLinkageCount returns the DeclaredLinkageCount field if non-nil, zero value otherwise.

### GetDeclaredLinkageCountOk

`func (o *RwdSourceLinkagesRow) GetDeclaredLinkageCountOk() (*int32, bool)`

GetDeclaredLinkageCountOk returns a tuple with the DeclaredLinkageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredLinkageCount

`func (o *RwdSourceLinkagesRow) SetDeclaredLinkageCount(v int32)`

SetDeclaredLinkageCount sets DeclaredLinkageCount field to given value.

### HasDeclaredLinkageCount

`func (o *RwdSourceLinkagesRow) HasDeclaredLinkageCount() bool`

HasDeclaredLinkageCount returns a boolean if a field has been set.

### SetDeclaredLinkageCountNil

`func (o *RwdSourceLinkagesRow) SetDeclaredLinkageCountNil(b bool)`

 SetDeclaredLinkageCountNil sets the value for DeclaredLinkageCount to be an explicit nil

### UnsetDeclaredLinkageCount
`func (o *RwdSourceLinkagesRow) UnsetDeclaredLinkageCount()`

UnsetDeclaredLinkageCount ensures that no value is present for DeclaredLinkageCount, not even an explicit nil
### GetIsPreLinked

`func (o *RwdSourceLinkagesRow) GetIsPreLinked() bool`

GetIsPreLinked returns the IsPreLinked field if non-nil, zero value otherwise.

### GetIsPreLinkedOk

`func (o *RwdSourceLinkagesRow) GetIsPreLinkedOk() (*bool, bool)`

GetIsPreLinkedOk returns a tuple with the IsPreLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreLinked

`func (o *RwdSourceLinkagesRow) SetIsPreLinked(v bool)`

SetIsPreLinked sets IsPreLinked field to given value.

### HasIsPreLinked

`func (o *RwdSourceLinkagesRow) HasIsPreLinked() bool`

HasIsPreLinked returns a boolean if a field has been set.

### SetIsPreLinkedNil

`func (o *RwdSourceLinkagesRow) SetIsPreLinkedNil(b bool)`

 SetIsPreLinkedNil sets the value for IsPreLinked to be an explicit nil

### UnsetIsPreLinked
`func (o *RwdSourceLinkagesRow) UnsetIsPreLinked()`

UnsetIsPreLinked ensures that no value is present for IsPreLinked, not even an explicit nil
### GetPreLinkedAnswer

`func (o *RwdSourceLinkagesRow) GetPreLinkedAnswer() string`

GetPreLinkedAnswer returns the PreLinkedAnswer field if non-nil, zero value otherwise.

### GetPreLinkedAnswerOk

`func (o *RwdSourceLinkagesRow) GetPreLinkedAnswerOk() (*string, bool)`

GetPreLinkedAnswerOk returns a tuple with the PreLinkedAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLinkedAnswer

`func (o *RwdSourceLinkagesRow) SetPreLinkedAnswer(v string)`

SetPreLinkedAnswer sets PreLinkedAnswer field to given value.

### HasPreLinkedAnswer

`func (o *RwdSourceLinkagesRow) HasPreLinkedAnswer() bool`

HasPreLinkedAnswer returns a boolean if a field has been set.

### SetPreLinkedAnswerNil

`func (o *RwdSourceLinkagesRow) SetPreLinkedAnswerNil(b bool)`

 SetPreLinkedAnswerNil sets the value for PreLinkedAnswer to be an explicit nil

### UnsetPreLinkedAnswer
`func (o *RwdSourceLinkagesRow) UnsetPreLinkedAnswer()`

UnsetPreLinkedAnswer ensures that no value is present for PreLinkedAnswer, not even an explicit nil
### GetLinkedDataSourceName

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceName() string`

GetLinkedDataSourceName returns the LinkedDataSourceName field if non-nil, zero value otherwise.

### GetLinkedDataSourceNameOk

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceNameOk() (*string, bool)`

GetLinkedDataSourceNameOk returns a tuple with the LinkedDataSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDataSourceName

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceName(v string)`

SetLinkedDataSourceName sets LinkedDataSourceName field to given value.

### HasLinkedDataSourceName

`func (o *RwdSourceLinkagesRow) HasLinkedDataSourceName() bool`

HasLinkedDataSourceName returns a boolean if a field has been set.

### SetLinkedDataSourceNameNil

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceNameNil(b bool)`

 SetLinkedDataSourceNameNil sets the value for LinkedDataSourceName to be an explicit nil

### UnsetLinkedDataSourceName
`func (o *RwdSourceLinkagesRow) UnsetLinkedDataSourceName()`

UnsetLinkedDataSourceName ensures that no value is present for LinkedDataSourceName, not even an explicit nil
### GetLinkedDataSourceNameAligned

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceNameAligned() bool`

GetLinkedDataSourceNameAligned returns the LinkedDataSourceNameAligned field if non-nil, zero value otherwise.

### GetLinkedDataSourceNameAlignedOk

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceNameAlignedOk() (*bool, bool)`

GetLinkedDataSourceNameAlignedOk returns a tuple with the LinkedDataSourceNameAligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDataSourceNameAligned

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceNameAligned(v bool)`

SetLinkedDataSourceNameAligned sets LinkedDataSourceNameAligned field to given value.

### HasLinkedDataSourceNameAligned

`func (o *RwdSourceLinkagesRow) HasLinkedDataSourceNameAligned() bool`

HasLinkedDataSourceNameAligned returns a boolean if a field has been set.

### SetLinkedDataSourceNameAlignedNil

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceNameAlignedNil(b bool)`

 SetLinkedDataSourceNameAlignedNil sets the value for LinkedDataSourceNameAligned to be an explicit nil

### UnsetLinkedDataSourceNameAligned
`func (o *RwdSourceLinkagesRow) UnsetLinkedDataSourceNameAligned()`

UnsetLinkedDataSourceNameAligned ensures that no value is present for LinkedDataSourceNameAligned, not even an explicit nil
### GetLinkedDataSourceNameSource

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceNameSource() string`

GetLinkedDataSourceNameSource returns the LinkedDataSourceNameSource field if non-nil, zero value otherwise.

### GetLinkedDataSourceNameSourceOk

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceNameSourceOk() (*string, bool)`

GetLinkedDataSourceNameSourceOk returns a tuple with the LinkedDataSourceNameSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDataSourceNameSource

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceNameSource(v string)`

SetLinkedDataSourceNameSource sets LinkedDataSourceNameSource field to given value.

### HasLinkedDataSourceNameSource

`func (o *RwdSourceLinkagesRow) HasLinkedDataSourceNameSource() bool`

HasLinkedDataSourceNameSource returns a boolean if a field has been set.

### SetLinkedDataSourceNameSourceNil

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceNameSourceNil(b bool)`

 SetLinkedDataSourceNameSourceNil sets the value for LinkedDataSourceNameSource to be an explicit nil

### UnsetLinkedDataSourceNameSource
`func (o *RwdSourceLinkagesRow) UnsetLinkedDataSourceNameSource()`

UnsetLinkedDataSourceNameSource ensures that no value is present for LinkedDataSourceNameSource, not even an explicit nil
### GetNamesOtherSentinel

`func (o *RwdSourceLinkagesRow) GetNamesOtherSentinel() bool`

GetNamesOtherSentinel returns the NamesOtherSentinel field if non-nil, zero value otherwise.

### GetNamesOtherSentinelOk

`func (o *RwdSourceLinkagesRow) GetNamesOtherSentinelOk() (*bool, bool)`

GetNamesOtherSentinelOk returns a tuple with the NamesOtherSentinel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamesOtherSentinel

`func (o *RwdSourceLinkagesRow) SetNamesOtherSentinel(v bool)`

SetNamesOtherSentinel sets NamesOtherSentinel field to given value.

### HasNamesOtherSentinel

`func (o *RwdSourceLinkagesRow) HasNamesOtherSentinel() bool`

HasNamesOtherSentinel returns a boolean if a field has been set.

### SetNamesOtherSentinelNil

`func (o *RwdSourceLinkagesRow) SetNamesOtherSentinelNil(b bool)`

 SetNamesOtherSentinelNil sets the value for NamesOtherSentinel to be an explicit nil

### UnsetNamesOtherSentinel
`func (o *RwdSourceLinkagesRow) UnsetNamesOtherSentinel()`

UnsetNamesOtherSentinel ensures that no value is present for NamesOtherSentinel, not even an explicit nil
### GetNamesRegisteredSource

`func (o *RwdSourceLinkagesRow) GetNamesRegisteredSource() bool`

GetNamesRegisteredSource returns the NamesRegisteredSource field if non-nil, zero value otherwise.

### GetNamesRegisteredSourceOk

`func (o *RwdSourceLinkagesRow) GetNamesRegisteredSourceOk() (*bool, bool)`

GetNamesRegisteredSourceOk returns a tuple with the NamesRegisteredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamesRegisteredSource

`func (o *RwdSourceLinkagesRow) SetNamesRegisteredSource(v bool)`

SetNamesRegisteredSource sets NamesRegisteredSource field to given value.

### HasNamesRegisteredSource

`func (o *RwdSourceLinkagesRow) HasNamesRegisteredSource() bool`

HasNamesRegisteredSource returns a boolean if a field has been set.

### SetNamesRegisteredSourceNil

`func (o *RwdSourceLinkagesRow) SetNamesRegisteredSourceNil(b bool)`

 SetNamesRegisteredSourceNil sets the value for NamesRegisteredSource to be an explicit nil

### UnsetNamesRegisteredSource
`func (o *RwdSourceLinkagesRow) UnsetNamesRegisteredSource()`

UnsetNamesRegisteredSource ensures that no value is present for NamesRegisteredSource, not even an explicit nil
### GetLinkedDataSourceOther

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceOther() string`

GetLinkedDataSourceOther returns the LinkedDataSourceOther field if non-nil, zero value otherwise.

### GetLinkedDataSourceOtherOk

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceOtherOk() (*string, bool)`

GetLinkedDataSourceOtherOk returns a tuple with the LinkedDataSourceOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDataSourceOther

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceOther(v string)`

SetLinkedDataSourceOther sets LinkedDataSourceOther field to given value.

### HasLinkedDataSourceOther

`func (o *RwdSourceLinkagesRow) HasLinkedDataSourceOther() bool`

HasLinkedDataSourceOther returns a boolean if a field has been set.

### SetLinkedDataSourceOtherNil

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceOtherNil(b bool)`

 SetLinkedDataSourceOtherNil sets the value for LinkedDataSourceOther to be an explicit nil

### UnsetLinkedDataSourceOther
`func (o *RwdSourceLinkagesRow) UnsetLinkedDataSourceOther()`

UnsetLinkedDataSourceOther ensures that no value is present for LinkedDataSourceOther, not even an explicit nil
### GetLinkedDataSourceOtherAligned

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceOtherAligned() bool`

GetLinkedDataSourceOtherAligned returns the LinkedDataSourceOtherAligned field if non-nil, zero value otherwise.

### GetLinkedDataSourceOtherAlignedOk

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceOtherAlignedOk() (*bool, bool)`

GetLinkedDataSourceOtherAlignedOk returns a tuple with the LinkedDataSourceOtherAligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDataSourceOtherAligned

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceOtherAligned(v bool)`

SetLinkedDataSourceOtherAligned sets LinkedDataSourceOtherAligned field to given value.

### HasLinkedDataSourceOtherAligned

`func (o *RwdSourceLinkagesRow) HasLinkedDataSourceOtherAligned() bool`

HasLinkedDataSourceOtherAligned returns a boolean if a field has been set.

### SetLinkedDataSourceOtherAlignedNil

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceOtherAlignedNil(b bool)`

 SetLinkedDataSourceOtherAlignedNil sets the value for LinkedDataSourceOtherAligned to be an explicit nil

### UnsetLinkedDataSourceOtherAligned
`func (o *RwdSourceLinkagesRow) UnsetLinkedDataSourceOtherAligned()`

UnsetLinkedDataSourceOtherAligned ensures that no value is present for LinkedDataSourceOtherAligned, not even an explicit nil
### GetLinkedDataSourceOtherSource

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceOtherSource() string`

GetLinkedDataSourceOtherSource returns the LinkedDataSourceOtherSource field if non-nil, zero value otherwise.

### GetLinkedDataSourceOtherSourceOk

`func (o *RwdSourceLinkagesRow) GetLinkedDataSourceOtherSourceOk() (*string, bool)`

GetLinkedDataSourceOtherSourceOk returns a tuple with the LinkedDataSourceOtherSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDataSourceOtherSource

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceOtherSource(v string)`

SetLinkedDataSourceOtherSource sets LinkedDataSourceOtherSource field to given value.

### HasLinkedDataSourceOtherSource

`func (o *RwdSourceLinkagesRow) HasLinkedDataSourceOtherSource() bool`

HasLinkedDataSourceOtherSource returns a boolean if a field has been set.

### SetLinkedDataSourceOtherSourceNil

`func (o *RwdSourceLinkagesRow) SetLinkedDataSourceOtherSourceNil(b bool)`

 SetLinkedDataSourceOtherSourceNil sets the value for LinkedDataSourceOtherSource to be an explicit nil

### UnsetLinkedDataSourceOtherSource
`func (o *RwdSourceLinkagesRow) UnsetLinkedDataSourceOtherSource()`

UnsetLinkedDataSourceOtherSource ensures that no value is present for LinkedDataSourceOtherSource, not even an explicit nil
### GetLinkageCompleteness

`func (o *RwdSourceLinkagesRow) GetLinkageCompleteness() string`

GetLinkageCompleteness returns the LinkageCompleteness field if non-nil, zero value otherwise.

### GetLinkageCompletenessOk

`func (o *RwdSourceLinkagesRow) GetLinkageCompletenessOk() (*string, bool)`

GetLinkageCompletenessOk returns a tuple with the LinkageCompleteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageCompleteness

`func (o *RwdSourceLinkagesRow) SetLinkageCompleteness(v string)`

SetLinkageCompleteness sets LinkageCompleteness field to given value.

### HasLinkageCompleteness

`func (o *RwdSourceLinkagesRow) HasLinkageCompleteness() bool`

HasLinkageCompleteness returns a boolean if a field has been set.

### SetLinkageCompletenessNil

`func (o *RwdSourceLinkagesRow) SetLinkageCompletenessNil(b bool)`

 SetLinkageCompletenessNil sets the value for LinkageCompleteness to be an explicit nil

### UnsetLinkageCompleteness
`func (o *RwdSourceLinkagesRow) UnsetLinkageCompleteness()`

UnsetLinkageCompleteness ensures that no value is present for LinkageCompleteness, not even an explicit nil
### GetLinkageCompletenessAligned

`func (o *RwdSourceLinkagesRow) GetLinkageCompletenessAligned() bool`

GetLinkageCompletenessAligned returns the LinkageCompletenessAligned field if non-nil, zero value otherwise.

### GetLinkageCompletenessAlignedOk

`func (o *RwdSourceLinkagesRow) GetLinkageCompletenessAlignedOk() (*bool, bool)`

GetLinkageCompletenessAlignedOk returns a tuple with the LinkageCompletenessAligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageCompletenessAligned

`func (o *RwdSourceLinkagesRow) SetLinkageCompletenessAligned(v bool)`

SetLinkageCompletenessAligned sets LinkageCompletenessAligned field to given value.

### HasLinkageCompletenessAligned

`func (o *RwdSourceLinkagesRow) HasLinkageCompletenessAligned() bool`

HasLinkageCompletenessAligned returns a boolean if a field has been set.

### SetLinkageCompletenessAlignedNil

`func (o *RwdSourceLinkagesRow) SetLinkageCompletenessAlignedNil(b bool)`

 SetLinkageCompletenessAlignedNil sets the value for LinkageCompletenessAligned to be an explicit nil

### UnsetLinkageCompletenessAligned
`func (o *RwdSourceLinkagesRow) UnsetLinkageCompletenessAligned()`

UnsetLinkageCompletenessAligned ensures that no value is present for LinkageCompletenessAligned, not even an explicit nil
### GetLinkageCompletenessSource

`func (o *RwdSourceLinkagesRow) GetLinkageCompletenessSource() string`

GetLinkageCompletenessSource returns the LinkageCompletenessSource field if non-nil, zero value otherwise.

### GetLinkageCompletenessSourceOk

`func (o *RwdSourceLinkagesRow) GetLinkageCompletenessSourceOk() (*string, bool)`

GetLinkageCompletenessSourceOk returns a tuple with the LinkageCompletenessSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageCompletenessSource

`func (o *RwdSourceLinkagesRow) SetLinkageCompletenessSource(v string)`

SetLinkageCompletenessSource sets LinkageCompletenessSource field to given value.

### HasLinkageCompletenessSource

`func (o *RwdSourceLinkagesRow) HasLinkageCompletenessSource() bool`

HasLinkageCompletenessSource returns a boolean if a field has been set.

### SetLinkageCompletenessSourceNil

`func (o *RwdSourceLinkagesRow) SetLinkageCompletenessSourceNil(b bool)`

 SetLinkageCompletenessSourceNil sets the value for LinkageCompletenessSource to be an explicit nil

### UnsetLinkageCompletenessSource
`func (o *RwdSourceLinkagesRow) UnsetLinkageCompletenessSource()`

UnsetLinkageCompletenessSource ensures that no value is present for LinkageCompletenessSource, not even an explicit nil
### GetLinkageVariable

`func (o *RwdSourceLinkagesRow) GetLinkageVariable() string`

GetLinkageVariable returns the LinkageVariable field if non-nil, zero value otherwise.

### GetLinkageVariableOk

`func (o *RwdSourceLinkagesRow) GetLinkageVariableOk() (*string, bool)`

GetLinkageVariableOk returns a tuple with the LinkageVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageVariable

`func (o *RwdSourceLinkagesRow) SetLinkageVariable(v string)`

SetLinkageVariable sets LinkageVariable field to given value.

### HasLinkageVariable

`func (o *RwdSourceLinkagesRow) HasLinkageVariable() bool`

HasLinkageVariable returns a boolean if a field has been set.

### SetLinkageVariableNil

`func (o *RwdSourceLinkagesRow) SetLinkageVariableNil(b bool)`

 SetLinkageVariableNil sets the value for LinkageVariable to be an explicit nil

### UnsetLinkageVariable
`func (o *RwdSourceLinkagesRow) UnsetLinkageVariable()`

UnsetLinkageVariable ensures that no value is present for LinkageVariable, not even an explicit nil
### GetLinkageVariableAligned

`func (o *RwdSourceLinkagesRow) GetLinkageVariableAligned() bool`

GetLinkageVariableAligned returns the LinkageVariableAligned field if non-nil, zero value otherwise.

### GetLinkageVariableAlignedOk

`func (o *RwdSourceLinkagesRow) GetLinkageVariableAlignedOk() (*bool, bool)`

GetLinkageVariableAlignedOk returns a tuple with the LinkageVariableAligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageVariableAligned

`func (o *RwdSourceLinkagesRow) SetLinkageVariableAligned(v bool)`

SetLinkageVariableAligned sets LinkageVariableAligned field to given value.

### HasLinkageVariableAligned

`func (o *RwdSourceLinkagesRow) HasLinkageVariableAligned() bool`

HasLinkageVariableAligned returns a boolean if a field has been set.

### SetLinkageVariableAlignedNil

`func (o *RwdSourceLinkagesRow) SetLinkageVariableAlignedNil(b bool)`

 SetLinkageVariableAlignedNil sets the value for LinkageVariableAligned to be an explicit nil

### UnsetLinkageVariableAligned
`func (o *RwdSourceLinkagesRow) UnsetLinkageVariableAligned()`

UnsetLinkageVariableAligned ensures that no value is present for LinkageVariableAligned, not even an explicit nil
### GetLinkageVariableSource

`func (o *RwdSourceLinkagesRow) GetLinkageVariableSource() string`

GetLinkageVariableSource returns the LinkageVariableSource field if non-nil, zero value otherwise.

### GetLinkageVariableSourceOk

`func (o *RwdSourceLinkagesRow) GetLinkageVariableSourceOk() (*string, bool)`

GetLinkageVariableSourceOk returns a tuple with the LinkageVariableSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageVariableSource

`func (o *RwdSourceLinkagesRow) SetLinkageVariableSource(v string)`

SetLinkageVariableSource sets LinkageVariableSource field to given value.

### HasLinkageVariableSource

`func (o *RwdSourceLinkagesRow) HasLinkageVariableSource() bool`

HasLinkageVariableSource returns a boolean if a field has been set.

### SetLinkageVariableSourceNil

`func (o *RwdSourceLinkagesRow) SetLinkageVariableSourceNil(b bool)`

 SetLinkageVariableSourceNil sets the value for LinkageVariableSource to be an explicit nil

### UnsetLinkageVariableSource
`func (o *RwdSourceLinkagesRow) UnsetLinkageVariableSource()`

UnsetLinkageVariableSource ensures that no value is present for LinkageVariableSource, not even an explicit nil
### GetLinkageStrategy

`func (o *RwdSourceLinkagesRow) GetLinkageStrategy() string`

GetLinkageStrategy returns the LinkageStrategy field if non-nil, zero value otherwise.

### GetLinkageStrategyOk

`func (o *RwdSourceLinkagesRow) GetLinkageStrategyOk() (*string, bool)`

GetLinkageStrategyOk returns a tuple with the LinkageStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageStrategy

`func (o *RwdSourceLinkagesRow) SetLinkageStrategy(v string)`

SetLinkageStrategy sets LinkageStrategy field to given value.

### HasLinkageStrategy

`func (o *RwdSourceLinkagesRow) HasLinkageStrategy() bool`

HasLinkageStrategy returns a boolean if a field has been set.

### SetLinkageStrategyNil

`func (o *RwdSourceLinkagesRow) SetLinkageStrategyNil(b bool)`

 SetLinkageStrategyNil sets the value for LinkageStrategy to be an explicit nil

### UnsetLinkageStrategy
`func (o *RwdSourceLinkagesRow) UnsetLinkageStrategy()`

UnsetLinkageStrategy ensures that no value is present for LinkageStrategy, not even an explicit nil
### GetLinkageStrategyAligned

`func (o *RwdSourceLinkagesRow) GetLinkageStrategyAligned() bool`

GetLinkageStrategyAligned returns the LinkageStrategyAligned field if non-nil, zero value otherwise.

### GetLinkageStrategyAlignedOk

`func (o *RwdSourceLinkagesRow) GetLinkageStrategyAlignedOk() (*bool, bool)`

GetLinkageStrategyAlignedOk returns a tuple with the LinkageStrategyAligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageStrategyAligned

`func (o *RwdSourceLinkagesRow) SetLinkageStrategyAligned(v bool)`

SetLinkageStrategyAligned sets LinkageStrategyAligned field to given value.

### HasLinkageStrategyAligned

`func (o *RwdSourceLinkagesRow) HasLinkageStrategyAligned() bool`

HasLinkageStrategyAligned returns a boolean if a field has been set.

### SetLinkageStrategyAlignedNil

`func (o *RwdSourceLinkagesRow) SetLinkageStrategyAlignedNil(b bool)`

 SetLinkageStrategyAlignedNil sets the value for LinkageStrategyAligned to be an explicit nil

### UnsetLinkageStrategyAligned
`func (o *RwdSourceLinkagesRow) UnsetLinkageStrategyAligned()`

UnsetLinkageStrategyAligned ensures that no value is present for LinkageStrategyAligned, not even an explicit nil
### GetLinkageStrategySource

`func (o *RwdSourceLinkagesRow) GetLinkageStrategySource() string`

GetLinkageStrategySource returns the LinkageStrategySource field if non-nil, zero value otherwise.

### GetLinkageStrategySourceOk

`func (o *RwdSourceLinkagesRow) GetLinkageStrategySourceOk() (*string, bool)`

GetLinkageStrategySourceOk returns a tuple with the LinkageStrategySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkageStrategySource

`func (o *RwdSourceLinkagesRow) SetLinkageStrategySource(v string)`

SetLinkageStrategySource sets LinkageStrategySource field to given value.

### HasLinkageStrategySource

`func (o *RwdSourceLinkagesRow) HasLinkageStrategySource() bool`

HasLinkageStrategySource returns a boolean if a field has been set.

### SetLinkageStrategySourceNil

`func (o *RwdSourceLinkagesRow) SetLinkageStrategySourceNil(b bool)`

 SetLinkageStrategySourceNil sets the value for LinkageStrategySource to be an explicit nil

### UnsetLinkageStrategySource
`func (o *RwdSourceLinkagesRow) UnsetLinkageStrategySource()`

UnsetLinkageStrategySource ensures that no value is present for LinkageStrategySource, not even an explicit nil
### GetRestrictionTier

`func (o *RwdSourceLinkagesRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdSourceLinkagesRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdSourceLinkagesRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdSourceLinkagesRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdSourceLinkagesRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdSourceLinkagesRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdSourceLinkagesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdSourceLinkagesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdSourceLinkagesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdSourceLinkagesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdSourceLinkagesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdSourceLinkagesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



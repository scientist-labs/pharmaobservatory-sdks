# RwdWithheldFieldsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisterKey** | Pointer to **NullableString** | Which of the five exports the row concerns: study, data_source_human, data_source_veterinary, institution or network. The entry key of the rwd_registers object type. | [optional] 
**RegisterLabel** | Pointer to **NullableString** | The register&#39;s plain-words label. | [optional] 
**SourceOrdinal** | Pointer to **NullableInt32** | The withheld column&#39;s position in the export. | [optional] 
**SourceHeader** | Pointer to **NullableString** | The export&#39;s header for the withheld column. | [optional] 
**ColumnName** | Pointer to **NullableString** | The name of the withheld column in the lane&#39;s landed schema. | [optional] 
**WithheldReason** | Pointer to **NullableString** | Why the column is not published. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdWithheldFieldsRow

`func NewRwdWithheldFieldsRow() *RwdWithheldFieldsRow`

NewRwdWithheldFieldsRow instantiates a new RwdWithheldFieldsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdWithheldFieldsRowWithDefaults

`func NewRwdWithheldFieldsRowWithDefaults() *RwdWithheldFieldsRow`

NewRwdWithheldFieldsRowWithDefaults instantiates a new RwdWithheldFieldsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisterKey

`func (o *RwdWithheldFieldsRow) GetRegisterKey() string`

GetRegisterKey returns the RegisterKey field if non-nil, zero value otherwise.

### GetRegisterKeyOk

`func (o *RwdWithheldFieldsRow) GetRegisterKeyOk() (*string, bool)`

GetRegisterKeyOk returns a tuple with the RegisterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterKey

`func (o *RwdWithheldFieldsRow) SetRegisterKey(v string)`

SetRegisterKey sets RegisterKey field to given value.

### HasRegisterKey

`func (o *RwdWithheldFieldsRow) HasRegisterKey() bool`

HasRegisterKey returns a boolean if a field has been set.

### SetRegisterKeyNil

`func (o *RwdWithheldFieldsRow) SetRegisterKeyNil(b bool)`

 SetRegisterKeyNil sets the value for RegisterKey to be an explicit nil

### UnsetRegisterKey
`func (o *RwdWithheldFieldsRow) UnsetRegisterKey()`

UnsetRegisterKey ensures that no value is present for RegisterKey, not even an explicit nil
### GetRegisterLabel

`func (o *RwdWithheldFieldsRow) GetRegisterLabel() string`

GetRegisterLabel returns the RegisterLabel field if non-nil, zero value otherwise.

### GetRegisterLabelOk

`func (o *RwdWithheldFieldsRow) GetRegisterLabelOk() (*string, bool)`

GetRegisterLabelOk returns a tuple with the RegisterLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterLabel

`func (o *RwdWithheldFieldsRow) SetRegisterLabel(v string)`

SetRegisterLabel sets RegisterLabel field to given value.

### HasRegisterLabel

`func (o *RwdWithheldFieldsRow) HasRegisterLabel() bool`

HasRegisterLabel returns a boolean if a field has been set.

### SetRegisterLabelNil

`func (o *RwdWithheldFieldsRow) SetRegisterLabelNil(b bool)`

 SetRegisterLabelNil sets the value for RegisterLabel to be an explicit nil

### UnsetRegisterLabel
`func (o *RwdWithheldFieldsRow) UnsetRegisterLabel()`

UnsetRegisterLabel ensures that no value is present for RegisterLabel, not even an explicit nil
### GetSourceOrdinal

`func (o *RwdWithheldFieldsRow) GetSourceOrdinal() int32`

GetSourceOrdinal returns the SourceOrdinal field if non-nil, zero value otherwise.

### GetSourceOrdinalOk

`func (o *RwdWithheldFieldsRow) GetSourceOrdinalOk() (*int32, bool)`

GetSourceOrdinalOk returns a tuple with the SourceOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrdinal

`func (o *RwdWithheldFieldsRow) SetSourceOrdinal(v int32)`

SetSourceOrdinal sets SourceOrdinal field to given value.

### HasSourceOrdinal

`func (o *RwdWithheldFieldsRow) HasSourceOrdinal() bool`

HasSourceOrdinal returns a boolean if a field has been set.

### SetSourceOrdinalNil

`func (o *RwdWithheldFieldsRow) SetSourceOrdinalNil(b bool)`

 SetSourceOrdinalNil sets the value for SourceOrdinal to be an explicit nil

### UnsetSourceOrdinal
`func (o *RwdWithheldFieldsRow) UnsetSourceOrdinal()`

UnsetSourceOrdinal ensures that no value is present for SourceOrdinal, not even an explicit nil
### GetSourceHeader

`func (o *RwdWithheldFieldsRow) GetSourceHeader() string`

GetSourceHeader returns the SourceHeader field if non-nil, zero value otherwise.

### GetSourceHeaderOk

`func (o *RwdWithheldFieldsRow) GetSourceHeaderOk() (*string, bool)`

GetSourceHeaderOk returns a tuple with the SourceHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHeader

`func (o *RwdWithheldFieldsRow) SetSourceHeader(v string)`

SetSourceHeader sets SourceHeader field to given value.

### HasSourceHeader

`func (o *RwdWithheldFieldsRow) HasSourceHeader() bool`

HasSourceHeader returns a boolean if a field has been set.

### SetSourceHeaderNil

`func (o *RwdWithheldFieldsRow) SetSourceHeaderNil(b bool)`

 SetSourceHeaderNil sets the value for SourceHeader to be an explicit nil

### UnsetSourceHeader
`func (o *RwdWithheldFieldsRow) UnsetSourceHeader()`

UnsetSourceHeader ensures that no value is present for SourceHeader, not even an explicit nil
### GetColumnName

`func (o *RwdWithheldFieldsRow) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *RwdWithheldFieldsRow) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *RwdWithheldFieldsRow) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *RwdWithheldFieldsRow) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### SetColumnNameNil

`func (o *RwdWithheldFieldsRow) SetColumnNameNil(b bool)`

 SetColumnNameNil sets the value for ColumnName to be an explicit nil

### UnsetColumnName
`func (o *RwdWithheldFieldsRow) UnsetColumnName()`

UnsetColumnName ensures that no value is present for ColumnName, not even an explicit nil
### GetWithheldReason

`func (o *RwdWithheldFieldsRow) GetWithheldReason() string`

GetWithheldReason returns the WithheldReason field if non-nil, zero value otherwise.

### GetWithheldReasonOk

`func (o *RwdWithheldFieldsRow) GetWithheldReasonOk() (*string, bool)`

GetWithheldReasonOk returns a tuple with the WithheldReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldReason

`func (o *RwdWithheldFieldsRow) SetWithheldReason(v string)`

SetWithheldReason sets WithheldReason field to given value.

### HasWithheldReason

`func (o *RwdWithheldFieldsRow) HasWithheldReason() bool`

HasWithheldReason returns a boolean if a field has been set.

### SetWithheldReasonNil

`func (o *RwdWithheldFieldsRow) SetWithheldReasonNil(b bool)`

 SetWithheldReasonNil sets the value for WithheldReason to be an explicit nil

### UnsetWithheldReason
`func (o *RwdWithheldFieldsRow) UnsetWithheldReason()`

UnsetWithheldReason ensures that no value is present for WithheldReason, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdWithheldFieldsRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdWithheldFieldsRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdWithheldFieldsRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdWithheldFieldsRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdWithheldFieldsRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdWithheldFieldsRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdWithheldFieldsRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdWithheldFieldsRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdWithheldFieldsRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdWithheldFieldsRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdWithheldFieldsRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdWithheldFieldsRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdWithheldFieldsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdWithheldFieldsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdWithheldFieldsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdWithheldFieldsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdWithheldFieldsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdWithheldFieldsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdWithheldFieldsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdWithheldFieldsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdWithheldFieldsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdWithheldFieldsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdWithheldFieldsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdWithheldFieldsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



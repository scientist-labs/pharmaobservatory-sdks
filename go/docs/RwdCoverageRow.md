# RwdCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainKey** | Pointer to **NullableString** | Which denominator chain the row belongs to. | [optional] 
**ChainLabel** | Pointer to **NullableString** | The chain&#39;s plain-words label. | [optional] 
**ChainIsAPartition** | Pointer to **NullableBool** | True when the chain&#39;s steps partition the denominator (they sum to it); false when steps overlap. | [optional] 
**StepOrdinal** | Pointer to **NullableInt32** | The step&#39;s position in the chain; 0 is the denominator itself, published as a step so the chain reads top to bottom. | [optional] 
**StepKey** | Pointer to **NullableString** | The step&#39;s key. | [optional] 
**StepLabel** | Pointer to **NullableString** | The step&#39;s plain-words label. | [optional] 
**RegisterKey** | Pointer to **NullableString** | Which of the five exports the row concerns: study, data_source_human, data_source_veterinary, institution or network. The entry key of the rwd_registers object type. | [optional] 
**MeasuredValue** | Pointer to **NullableInt64** | The measured count at this step. | [optional] 
**DenominatorKey** | Pointer to **NullableString** | Which denominator the step is read against. | [optional] 
**DenominatorLabel** | Pointer to **NullableString** | The denominator&#39;s plain-words label. | [optional] 
**DenominatorValue** | Pointer to **NullableInt64** | The denominator&#39;s value. | [optional] 
**ShareOfDenominator** | Pointer to **NullableFloat32** | measured_value over denominator_value, as a fraction. | [optional] 
**DerivationNote** | Pointer to **NullableString** | How the step is derived. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**SnapshotDate** | Pointer to **NullableString** | The date of the catalogue snapshot. | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdCoverageRow

`func NewRwdCoverageRow() *RwdCoverageRow`

NewRwdCoverageRow instantiates a new RwdCoverageRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdCoverageRowWithDefaults

`func NewRwdCoverageRowWithDefaults() *RwdCoverageRow`

NewRwdCoverageRowWithDefaults instantiates a new RwdCoverageRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainKey

`func (o *RwdCoverageRow) GetChainKey() string`

GetChainKey returns the ChainKey field if non-nil, zero value otherwise.

### GetChainKeyOk

`func (o *RwdCoverageRow) GetChainKeyOk() (*string, bool)`

GetChainKeyOk returns a tuple with the ChainKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainKey

`func (o *RwdCoverageRow) SetChainKey(v string)`

SetChainKey sets ChainKey field to given value.

### HasChainKey

`func (o *RwdCoverageRow) HasChainKey() bool`

HasChainKey returns a boolean if a field has been set.

### SetChainKeyNil

`func (o *RwdCoverageRow) SetChainKeyNil(b bool)`

 SetChainKeyNil sets the value for ChainKey to be an explicit nil

### UnsetChainKey
`func (o *RwdCoverageRow) UnsetChainKey()`

UnsetChainKey ensures that no value is present for ChainKey, not even an explicit nil
### GetChainLabel

`func (o *RwdCoverageRow) GetChainLabel() string`

GetChainLabel returns the ChainLabel field if non-nil, zero value otherwise.

### GetChainLabelOk

`func (o *RwdCoverageRow) GetChainLabelOk() (*string, bool)`

GetChainLabelOk returns a tuple with the ChainLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainLabel

`func (o *RwdCoverageRow) SetChainLabel(v string)`

SetChainLabel sets ChainLabel field to given value.

### HasChainLabel

`func (o *RwdCoverageRow) HasChainLabel() bool`

HasChainLabel returns a boolean if a field has been set.

### SetChainLabelNil

`func (o *RwdCoverageRow) SetChainLabelNil(b bool)`

 SetChainLabelNil sets the value for ChainLabel to be an explicit nil

### UnsetChainLabel
`func (o *RwdCoverageRow) UnsetChainLabel()`

UnsetChainLabel ensures that no value is present for ChainLabel, not even an explicit nil
### GetChainIsAPartition

`func (o *RwdCoverageRow) GetChainIsAPartition() bool`

GetChainIsAPartition returns the ChainIsAPartition field if non-nil, zero value otherwise.

### GetChainIsAPartitionOk

`func (o *RwdCoverageRow) GetChainIsAPartitionOk() (*bool, bool)`

GetChainIsAPartitionOk returns a tuple with the ChainIsAPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainIsAPartition

`func (o *RwdCoverageRow) SetChainIsAPartition(v bool)`

SetChainIsAPartition sets ChainIsAPartition field to given value.

### HasChainIsAPartition

`func (o *RwdCoverageRow) HasChainIsAPartition() bool`

HasChainIsAPartition returns a boolean if a field has been set.

### SetChainIsAPartitionNil

`func (o *RwdCoverageRow) SetChainIsAPartitionNil(b bool)`

 SetChainIsAPartitionNil sets the value for ChainIsAPartition to be an explicit nil

### UnsetChainIsAPartition
`func (o *RwdCoverageRow) UnsetChainIsAPartition()`

UnsetChainIsAPartition ensures that no value is present for ChainIsAPartition, not even an explicit nil
### GetStepOrdinal

`func (o *RwdCoverageRow) GetStepOrdinal() int32`

GetStepOrdinal returns the StepOrdinal field if non-nil, zero value otherwise.

### GetStepOrdinalOk

`func (o *RwdCoverageRow) GetStepOrdinalOk() (*int32, bool)`

GetStepOrdinalOk returns a tuple with the StepOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepOrdinal

`func (o *RwdCoverageRow) SetStepOrdinal(v int32)`

SetStepOrdinal sets StepOrdinal field to given value.

### HasStepOrdinal

`func (o *RwdCoverageRow) HasStepOrdinal() bool`

HasStepOrdinal returns a boolean if a field has been set.

### SetStepOrdinalNil

`func (o *RwdCoverageRow) SetStepOrdinalNil(b bool)`

 SetStepOrdinalNil sets the value for StepOrdinal to be an explicit nil

### UnsetStepOrdinal
`func (o *RwdCoverageRow) UnsetStepOrdinal()`

UnsetStepOrdinal ensures that no value is present for StepOrdinal, not even an explicit nil
### GetStepKey

`func (o *RwdCoverageRow) GetStepKey() string`

GetStepKey returns the StepKey field if non-nil, zero value otherwise.

### GetStepKeyOk

`func (o *RwdCoverageRow) GetStepKeyOk() (*string, bool)`

GetStepKeyOk returns a tuple with the StepKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepKey

`func (o *RwdCoverageRow) SetStepKey(v string)`

SetStepKey sets StepKey field to given value.

### HasStepKey

`func (o *RwdCoverageRow) HasStepKey() bool`

HasStepKey returns a boolean if a field has been set.

### SetStepKeyNil

`func (o *RwdCoverageRow) SetStepKeyNil(b bool)`

 SetStepKeyNil sets the value for StepKey to be an explicit nil

### UnsetStepKey
`func (o *RwdCoverageRow) UnsetStepKey()`

UnsetStepKey ensures that no value is present for StepKey, not even an explicit nil
### GetStepLabel

`func (o *RwdCoverageRow) GetStepLabel() string`

GetStepLabel returns the StepLabel field if non-nil, zero value otherwise.

### GetStepLabelOk

`func (o *RwdCoverageRow) GetStepLabelOk() (*string, bool)`

GetStepLabelOk returns a tuple with the StepLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepLabel

`func (o *RwdCoverageRow) SetStepLabel(v string)`

SetStepLabel sets StepLabel field to given value.

### HasStepLabel

`func (o *RwdCoverageRow) HasStepLabel() bool`

HasStepLabel returns a boolean if a field has been set.

### SetStepLabelNil

`func (o *RwdCoverageRow) SetStepLabelNil(b bool)`

 SetStepLabelNil sets the value for StepLabel to be an explicit nil

### UnsetStepLabel
`func (o *RwdCoverageRow) UnsetStepLabel()`

UnsetStepLabel ensures that no value is present for StepLabel, not even an explicit nil
### GetRegisterKey

`func (o *RwdCoverageRow) GetRegisterKey() string`

GetRegisterKey returns the RegisterKey field if non-nil, zero value otherwise.

### GetRegisterKeyOk

`func (o *RwdCoverageRow) GetRegisterKeyOk() (*string, bool)`

GetRegisterKeyOk returns a tuple with the RegisterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterKey

`func (o *RwdCoverageRow) SetRegisterKey(v string)`

SetRegisterKey sets RegisterKey field to given value.

### HasRegisterKey

`func (o *RwdCoverageRow) HasRegisterKey() bool`

HasRegisterKey returns a boolean if a field has been set.

### SetRegisterKeyNil

`func (o *RwdCoverageRow) SetRegisterKeyNil(b bool)`

 SetRegisterKeyNil sets the value for RegisterKey to be an explicit nil

### UnsetRegisterKey
`func (o *RwdCoverageRow) UnsetRegisterKey()`

UnsetRegisterKey ensures that no value is present for RegisterKey, not even an explicit nil
### GetMeasuredValue

`func (o *RwdCoverageRow) GetMeasuredValue() int64`

GetMeasuredValue returns the MeasuredValue field if non-nil, zero value otherwise.

### GetMeasuredValueOk

`func (o *RwdCoverageRow) GetMeasuredValueOk() (*int64, bool)`

GetMeasuredValueOk returns a tuple with the MeasuredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredValue

`func (o *RwdCoverageRow) SetMeasuredValue(v int64)`

SetMeasuredValue sets MeasuredValue field to given value.

### HasMeasuredValue

`func (o *RwdCoverageRow) HasMeasuredValue() bool`

HasMeasuredValue returns a boolean if a field has been set.

### SetMeasuredValueNil

`func (o *RwdCoverageRow) SetMeasuredValueNil(b bool)`

 SetMeasuredValueNil sets the value for MeasuredValue to be an explicit nil

### UnsetMeasuredValue
`func (o *RwdCoverageRow) UnsetMeasuredValue()`

UnsetMeasuredValue ensures that no value is present for MeasuredValue, not even an explicit nil
### GetDenominatorKey

`func (o *RwdCoverageRow) GetDenominatorKey() string`

GetDenominatorKey returns the DenominatorKey field if non-nil, zero value otherwise.

### GetDenominatorKeyOk

`func (o *RwdCoverageRow) GetDenominatorKeyOk() (*string, bool)`

GetDenominatorKeyOk returns a tuple with the DenominatorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorKey

`func (o *RwdCoverageRow) SetDenominatorKey(v string)`

SetDenominatorKey sets DenominatorKey field to given value.

### HasDenominatorKey

`func (o *RwdCoverageRow) HasDenominatorKey() bool`

HasDenominatorKey returns a boolean if a field has been set.

### SetDenominatorKeyNil

`func (o *RwdCoverageRow) SetDenominatorKeyNil(b bool)`

 SetDenominatorKeyNil sets the value for DenominatorKey to be an explicit nil

### UnsetDenominatorKey
`func (o *RwdCoverageRow) UnsetDenominatorKey()`

UnsetDenominatorKey ensures that no value is present for DenominatorKey, not even an explicit nil
### GetDenominatorLabel

`func (o *RwdCoverageRow) GetDenominatorLabel() string`

GetDenominatorLabel returns the DenominatorLabel field if non-nil, zero value otherwise.

### GetDenominatorLabelOk

`func (o *RwdCoverageRow) GetDenominatorLabelOk() (*string, bool)`

GetDenominatorLabelOk returns a tuple with the DenominatorLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorLabel

`func (o *RwdCoverageRow) SetDenominatorLabel(v string)`

SetDenominatorLabel sets DenominatorLabel field to given value.

### HasDenominatorLabel

`func (o *RwdCoverageRow) HasDenominatorLabel() bool`

HasDenominatorLabel returns a boolean if a field has been set.

### SetDenominatorLabelNil

`func (o *RwdCoverageRow) SetDenominatorLabelNil(b bool)`

 SetDenominatorLabelNil sets the value for DenominatorLabel to be an explicit nil

### UnsetDenominatorLabel
`func (o *RwdCoverageRow) UnsetDenominatorLabel()`

UnsetDenominatorLabel ensures that no value is present for DenominatorLabel, not even an explicit nil
### GetDenominatorValue

`func (o *RwdCoverageRow) GetDenominatorValue() int64`

GetDenominatorValue returns the DenominatorValue field if non-nil, zero value otherwise.

### GetDenominatorValueOk

`func (o *RwdCoverageRow) GetDenominatorValueOk() (*int64, bool)`

GetDenominatorValueOk returns a tuple with the DenominatorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorValue

`func (o *RwdCoverageRow) SetDenominatorValue(v int64)`

SetDenominatorValue sets DenominatorValue field to given value.

### HasDenominatorValue

`func (o *RwdCoverageRow) HasDenominatorValue() bool`

HasDenominatorValue returns a boolean if a field has been set.

### SetDenominatorValueNil

`func (o *RwdCoverageRow) SetDenominatorValueNil(b bool)`

 SetDenominatorValueNil sets the value for DenominatorValue to be an explicit nil

### UnsetDenominatorValue
`func (o *RwdCoverageRow) UnsetDenominatorValue()`

UnsetDenominatorValue ensures that no value is present for DenominatorValue, not even an explicit nil
### GetShareOfDenominator

`func (o *RwdCoverageRow) GetShareOfDenominator() float32`

GetShareOfDenominator returns the ShareOfDenominator field if non-nil, zero value otherwise.

### GetShareOfDenominatorOk

`func (o *RwdCoverageRow) GetShareOfDenominatorOk() (*float32, bool)`

GetShareOfDenominatorOk returns a tuple with the ShareOfDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareOfDenominator

`func (o *RwdCoverageRow) SetShareOfDenominator(v float32)`

SetShareOfDenominator sets ShareOfDenominator field to given value.

### HasShareOfDenominator

`func (o *RwdCoverageRow) HasShareOfDenominator() bool`

HasShareOfDenominator returns a boolean if a field has been set.

### SetShareOfDenominatorNil

`func (o *RwdCoverageRow) SetShareOfDenominatorNil(b bool)`

 SetShareOfDenominatorNil sets the value for ShareOfDenominator to be an explicit nil

### UnsetShareOfDenominator
`func (o *RwdCoverageRow) UnsetShareOfDenominator()`

UnsetShareOfDenominator ensures that no value is present for ShareOfDenominator, not even an explicit nil
### GetDerivationNote

`func (o *RwdCoverageRow) GetDerivationNote() string`

GetDerivationNote returns the DerivationNote field if non-nil, zero value otherwise.

### GetDerivationNoteOk

`func (o *RwdCoverageRow) GetDerivationNoteOk() (*string, bool)`

GetDerivationNoteOk returns a tuple with the DerivationNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationNote

`func (o *RwdCoverageRow) SetDerivationNote(v string)`

SetDerivationNote sets DerivationNote field to given value.

### HasDerivationNote

`func (o *RwdCoverageRow) HasDerivationNote() bool`

HasDerivationNote returns a boolean if a field has been set.

### SetDerivationNoteNil

`func (o *RwdCoverageRow) SetDerivationNoteNil(b bool)`

 SetDerivationNoteNil sets the value for DerivationNote to be an explicit nil

### UnsetDerivationNote
`func (o *RwdCoverageRow) UnsetDerivationNote()`

UnsetDerivationNote ensures that no value is present for DerivationNote, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdCoverageRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdCoverageRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdCoverageRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdCoverageRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdCoverageRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdCoverageRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetSnapshotDate

`func (o *RwdCoverageRow) GetSnapshotDate() string`

GetSnapshotDate returns the SnapshotDate field if non-nil, zero value otherwise.

### GetSnapshotDateOk

`func (o *RwdCoverageRow) GetSnapshotDateOk() (*string, bool)`

GetSnapshotDateOk returns a tuple with the SnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDate

`func (o *RwdCoverageRow) SetSnapshotDate(v string)`

SetSnapshotDate sets SnapshotDate field to given value.

### HasSnapshotDate

`func (o *RwdCoverageRow) HasSnapshotDate() bool`

HasSnapshotDate returns a boolean if a field has been set.

### SetSnapshotDateNil

`func (o *RwdCoverageRow) SetSnapshotDateNil(b bool)`

 SetSnapshotDateNil sets the value for SnapshotDate to be an explicit nil

### UnsetSnapshotDate
`func (o *RwdCoverageRow) UnsetSnapshotDate()`

UnsetSnapshotDate ensures that no value is present for SnapshotDate, not even an explicit nil
### GetAccessedMonth

`func (o *RwdCoverageRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdCoverageRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdCoverageRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdCoverageRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdCoverageRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdCoverageRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdCoverageRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdCoverageRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdCoverageRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdCoverageRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdCoverageRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdCoverageRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdCoverageRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdCoverageRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdCoverageRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdCoverageRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdCoverageRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdCoverageRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TargetCrowdingRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] 
**MaxPhase** | Pointer to **NullableInt32** | The band: the furthest phase a program reached, 0 to 4 (0 is phase not reported). | [optional] 
**NPrograms** | Pointer to **NullableInt32** | Distinct drug and indication programs whose furthest phase is the band. | [optional] 
**NActivePrograms** | Pointer to **NullableInt32** | Programs in the band with no terminated, withdrawn or suspended trial. | [optional] 
**NDiscontinued** | Pointer to **NullableInt32** | Programs in the band with at least one terminated, withdrawn or suspended trial. | [optional] 
**NDrugs** | Pointer to **NullableInt32** | Distinct drugs among the band&#39;s programs. | [optional] 

## Methods

### NewTargetCrowdingRow

`func NewTargetCrowdingRow() *TargetCrowdingRow`

NewTargetCrowdingRow instantiates a new TargetCrowdingRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCrowdingRowWithDefaults

`func NewTargetCrowdingRowWithDefaults() *TargetCrowdingRow`

NewTargetCrowdingRowWithDefaults instantiates a new TargetCrowdingRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetCrowdingRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetCrowdingRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetCrowdingRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetCrowdingRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetCrowdingRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetCrowdingRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *TargetCrowdingRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *TargetCrowdingRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *TargetCrowdingRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *TargetCrowdingRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *TargetCrowdingRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *TargetCrowdingRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetMaxPhase

`func (o *TargetCrowdingRow) GetMaxPhase() int32`

GetMaxPhase returns the MaxPhase field if non-nil, zero value otherwise.

### GetMaxPhaseOk

`func (o *TargetCrowdingRow) GetMaxPhaseOk() (*int32, bool)`

GetMaxPhaseOk returns a tuple with the MaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhase

`func (o *TargetCrowdingRow) SetMaxPhase(v int32)`

SetMaxPhase sets MaxPhase field to given value.

### HasMaxPhase

`func (o *TargetCrowdingRow) HasMaxPhase() bool`

HasMaxPhase returns a boolean if a field has been set.

### SetMaxPhaseNil

`func (o *TargetCrowdingRow) SetMaxPhaseNil(b bool)`

 SetMaxPhaseNil sets the value for MaxPhase to be an explicit nil

### UnsetMaxPhase
`func (o *TargetCrowdingRow) UnsetMaxPhase()`

UnsetMaxPhase ensures that no value is present for MaxPhase, not even an explicit nil
### GetNPrograms

`func (o *TargetCrowdingRow) GetNPrograms() int32`

GetNPrograms returns the NPrograms field if non-nil, zero value otherwise.

### GetNProgramsOk

`func (o *TargetCrowdingRow) GetNProgramsOk() (*int32, bool)`

GetNProgramsOk returns a tuple with the NPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPrograms

`func (o *TargetCrowdingRow) SetNPrograms(v int32)`

SetNPrograms sets NPrograms field to given value.

### HasNPrograms

`func (o *TargetCrowdingRow) HasNPrograms() bool`

HasNPrograms returns a boolean if a field has been set.

### SetNProgramsNil

`func (o *TargetCrowdingRow) SetNProgramsNil(b bool)`

 SetNProgramsNil sets the value for NPrograms to be an explicit nil

### UnsetNPrograms
`func (o *TargetCrowdingRow) UnsetNPrograms()`

UnsetNPrograms ensures that no value is present for NPrograms, not even an explicit nil
### GetNActivePrograms

`func (o *TargetCrowdingRow) GetNActivePrograms() int32`

GetNActivePrograms returns the NActivePrograms field if non-nil, zero value otherwise.

### GetNActiveProgramsOk

`func (o *TargetCrowdingRow) GetNActiveProgramsOk() (*int32, bool)`

GetNActiveProgramsOk returns a tuple with the NActivePrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActivePrograms

`func (o *TargetCrowdingRow) SetNActivePrograms(v int32)`

SetNActivePrograms sets NActivePrograms field to given value.

### HasNActivePrograms

`func (o *TargetCrowdingRow) HasNActivePrograms() bool`

HasNActivePrograms returns a boolean if a field has been set.

### SetNActiveProgramsNil

`func (o *TargetCrowdingRow) SetNActiveProgramsNil(b bool)`

 SetNActiveProgramsNil sets the value for NActivePrograms to be an explicit nil

### UnsetNActivePrograms
`func (o *TargetCrowdingRow) UnsetNActivePrograms()`

UnsetNActivePrograms ensures that no value is present for NActivePrograms, not even an explicit nil
### GetNDiscontinued

`func (o *TargetCrowdingRow) GetNDiscontinued() int32`

GetNDiscontinued returns the NDiscontinued field if non-nil, zero value otherwise.

### GetNDiscontinuedOk

`func (o *TargetCrowdingRow) GetNDiscontinuedOk() (*int32, bool)`

GetNDiscontinuedOk returns a tuple with the NDiscontinued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDiscontinued

`func (o *TargetCrowdingRow) SetNDiscontinued(v int32)`

SetNDiscontinued sets NDiscontinued field to given value.

### HasNDiscontinued

`func (o *TargetCrowdingRow) HasNDiscontinued() bool`

HasNDiscontinued returns a boolean if a field has been set.

### SetNDiscontinuedNil

`func (o *TargetCrowdingRow) SetNDiscontinuedNil(b bool)`

 SetNDiscontinuedNil sets the value for NDiscontinued to be an explicit nil

### UnsetNDiscontinued
`func (o *TargetCrowdingRow) UnsetNDiscontinued()`

UnsetNDiscontinued ensures that no value is present for NDiscontinued, not even an explicit nil
### GetNDrugs

`func (o *TargetCrowdingRow) GetNDrugs() int32`

GetNDrugs returns the NDrugs field if non-nil, zero value otherwise.

### GetNDrugsOk

`func (o *TargetCrowdingRow) GetNDrugsOk() (*int32, bool)`

GetNDrugsOk returns a tuple with the NDrugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDrugs

`func (o *TargetCrowdingRow) SetNDrugs(v int32)`

SetNDrugs sets NDrugs field to given value.

### HasNDrugs

`func (o *TargetCrowdingRow) HasNDrugs() bool`

HasNDrugs returns a boolean if a field has been set.

### SetNDrugsNil

`func (o *TargetCrowdingRow) SetNDrugsNil(b bool)`

 SetNDrugsNil sets the value for NDrugs to be an explicit nil

### UnsetNDrugs
`func (o *TargetCrowdingRow) UnsetNDrugs()`

UnsetNDrugs ensures that no value is present for NDrugs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



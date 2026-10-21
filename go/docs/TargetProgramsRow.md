# TargetProgramsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] 
**DrugId** | Pointer to **NullableString** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition as a MeSH descriptor id (D followed by digits), the id the conditions datasets use. | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s name. | [optional] 
**MaxPhase** | Pointer to **NullableInt32** | The furthest phase the program reached, 0 to 4 (0 is phase not reported). | [optional] 
**NTrials** | Pointer to **NullableInt32** | Distinct trials of the drug in the indication. | [optional] 
**NCompleted** | Pointer to **NullableInt32** | Of n_trials, those with status COMPLETED. | [optional] 
**NTerminatedTrials** | Pointer to **NullableInt32** | Of n_trials, those with status TERMINATED. | [optional] 
**FirstYear** | Pointer to **NullableInt32** | The earliest plausible trial start year. | [optional] 
**LastYear** | Pointer to **NullableInt32** | The latest plausible trial completion or start year. | [optional] 
**IsDiscontinued** | Pointer to **NullableBool** | True when the program has a terminated, withdrawn or suspended trial. | [optional] 
**NDiscontinued** | Pointer to **NullableInt32** | Distinct stopped trials (terminated, withdrawn or suspended); 0 for a live program. | [optional] 
**NTerminated** | Pointer to **NullableInt32** | Of the stopped trials, those TERMINATED. | [optional] 
**NWithdrawn** | Pointer to **NullableInt32** | Of the stopped trials, those WITHDRAWN. | [optional] 
**NSuspended** | Pointer to **NullableInt32** | Of the stopped trials, those SUSPENDED. | [optional] 
**MaxPhaseAtStop** | Pointer to **NullableInt32** | The furthest phase among the stopped trials, 0 to 4; null for a live program. | [optional] 
**LastStopYear** | Pointer to **NullableInt32** | The latest year a stopped trial ended; null for a live program. | [optional] 
**WhyStoppedReasons** | Pointer to **[]string** | The stop reasons registrants disclosed, verbatim; empty when none. | [optional] 
**WhyStoppedCategories** | Pointer to **[]string** | The normalized stop-reason categories (efficacy_futility, safety, funding and so on); empty when none. | [optional] 

## Methods

### NewTargetProgramsRow

`func NewTargetProgramsRow() *TargetProgramsRow`

NewTargetProgramsRow instantiates a new TargetProgramsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetProgramsRowWithDefaults

`func NewTargetProgramsRowWithDefaults() *TargetProgramsRow`

NewTargetProgramsRowWithDefaults instantiates a new TargetProgramsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetProgramsRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetProgramsRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetProgramsRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetProgramsRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetProgramsRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetProgramsRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *TargetProgramsRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *TargetProgramsRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *TargetProgramsRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *TargetProgramsRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *TargetProgramsRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *TargetProgramsRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetDrugId

`func (o *TargetProgramsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *TargetProgramsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *TargetProgramsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *TargetProgramsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *TargetProgramsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *TargetProgramsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *TargetProgramsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *TargetProgramsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *TargetProgramsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *TargetProgramsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *TargetProgramsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *TargetProgramsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetConditionId

`func (o *TargetProgramsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *TargetProgramsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *TargetProgramsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *TargetProgramsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *TargetProgramsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *TargetProgramsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *TargetProgramsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *TargetProgramsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *TargetProgramsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *TargetProgramsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *TargetProgramsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *TargetProgramsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetMaxPhase

`func (o *TargetProgramsRow) GetMaxPhase() int32`

GetMaxPhase returns the MaxPhase field if non-nil, zero value otherwise.

### GetMaxPhaseOk

`func (o *TargetProgramsRow) GetMaxPhaseOk() (*int32, bool)`

GetMaxPhaseOk returns a tuple with the MaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhase

`func (o *TargetProgramsRow) SetMaxPhase(v int32)`

SetMaxPhase sets MaxPhase field to given value.

### HasMaxPhase

`func (o *TargetProgramsRow) HasMaxPhase() bool`

HasMaxPhase returns a boolean if a field has been set.

### SetMaxPhaseNil

`func (o *TargetProgramsRow) SetMaxPhaseNil(b bool)`

 SetMaxPhaseNil sets the value for MaxPhase to be an explicit nil

### UnsetMaxPhase
`func (o *TargetProgramsRow) UnsetMaxPhase()`

UnsetMaxPhase ensures that no value is present for MaxPhase, not even an explicit nil
### GetNTrials

`func (o *TargetProgramsRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *TargetProgramsRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *TargetProgramsRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *TargetProgramsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *TargetProgramsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *TargetProgramsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNCompleted

`func (o *TargetProgramsRow) GetNCompleted() int32`

GetNCompleted returns the NCompleted field if non-nil, zero value otherwise.

### GetNCompletedOk

`func (o *TargetProgramsRow) GetNCompletedOk() (*int32, bool)`

GetNCompletedOk returns a tuple with the NCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompleted

`func (o *TargetProgramsRow) SetNCompleted(v int32)`

SetNCompleted sets NCompleted field to given value.

### HasNCompleted

`func (o *TargetProgramsRow) HasNCompleted() bool`

HasNCompleted returns a boolean if a field has been set.

### SetNCompletedNil

`func (o *TargetProgramsRow) SetNCompletedNil(b bool)`

 SetNCompletedNil sets the value for NCompleted to be an explicit nil

### UnsetNCompleted
`func (o *TargetProgramsRow) UnsetNCompleted()`

UnsetNCompleted ensures that no value is present for NCompleted, not even an explicit nil
### GetNTerminatedTrials

`func (o *TargetProgramsRow) GetNTerminatedTrials() int32`

GetNTerminatedTrials returns the NTerminatedTrials field if non-nil, zero value otherwise.

### GetNTerminatedTrialsOk

`func (o *TargetProgramsRow) GetNTerminatedTrialsOk() (*int32, bool)`

GetNTerminatedTrialsOk returns a tuple with the NTerminatedTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTerminatedTrials

`func (o *TargetProgramsRow) SetNTerminatedTrials(v int32)`

SetNTerminatedTrials sets NTerminatedTrials field to given value.

### HasNTerminatedTrials

`func (o *TargetProgramsRow) HasNTerminatedTrials() bool`

HasNTerminatedTrials returns a boolean if a field has been set.

### SetNTerminatedTrialsNil

`func (o *TargetProgramsRow) SetNTerminatedTrialsNil(b bool)`

 SetNTerminatedTrialsNil sets the value for NTerminatedTrials to be an explicit nil

### UnsetNTerminatedTrials
`func (o *TargetProgramsRow) UnsetNTerminatedTrials()`

UnsetNTerminatedTrials ensures that no value is present for NTerminatedTrials, not even an explicit nil
### GetFirstYear

`func (o *TargetProgramsRow) GetFirstYear() int32`

GetFirstYear returns the FirstYear field if non-nil, zero value otherwise.

### GetFirstYearOk

`func (o *TargetProgramsRow) GetFirstYearOk() (*int32, bool)`

GetFirstYearOk returns a tuple with the FirstYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstYear

`func (o *TargetProgramsRow) SetFirstYear(v int32)`

SetFirstYear sets FirstYear field to given value.

### HasFirstYear

`func (o *TargetProgramsRow) HasFirstYear() bool`

HasFirstYear returns a boolean if a field has been set.

### SetFirstYearNil

`func (o *TargetProgramsRow) SetFirstYearNil(b bool)`

 SetFirstYearNil sets the value for FirstYear to be an explicit nil

### UnsetFirstYear
`func (o *TargetProgramsRow) UnsetFirstYear()`

UnsetFirstYear ensures that no value is present for FirstYear, not even an explicit nil
### GetLastYear

`func (o *TargetProgramsRow) GetLastYear() int32`

GetLastYear returns the LastYear field if non-nil, zero value otherwise.

### GetLastYearOk

`func (o *TargetProgramsRow) GetLastYearOk() (*int32, bool)`

GetLastYearOk returns a tuple with the LastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYear

`func (o *TargetProgramsRow) SetLastYear(v int32)`

SetLastYear sets LastYear field to given value.

### HasLastYear

`func (o *TargetProgramsRow) HasLastYear() bool`

HasLastYear returns a boolean if a field has been set.

### SetLastYearNil

`func (o *TargetProgramsRow) SetLastYearNil(b bool)`

 SetLastYearNil sets the value for LastYear to be an explicit nil

### UnsetLastYear
`func (o *TargetProgramsRow) UnsetLastYear()`

UnsetLastYear ensures that no value is present for LastYear, not even an explicit nil
### GetIsDiscontinued

`func (o *TargetProgramsRow) GetIsDiscontinued() bool`

GetIsDiscontinued returns the IsDiscontinued field if non-nil, zero value otherwise.

### GetIsDiscontinuedOk

`func (o *TargetProgramsRow) GetIsDiscontinuedOk() (*bool, bool)`

GetIsDiscontinuedOk returns a tuple with the IsDiscontinued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiscontinued

`func (o *TargetProgramsRow) SetIsDiscontinued(v bool)`

SetIsDiscontinued sets IsDiscontinued field to given value.

### HasIsDiscontinued

`func (o *TargetProgramsRow) HasIsDiscontinued() bool`

HasIsDiscontinued returns a boolean if a field has been set.

### SetIsDiscontinuedNil

`func (o *TargetProgramsRow) SetIsDiscontinuedNil(b bool)`

 SetIsDiscontinuedNil sets the value for IsDiscontinued to be an explicit nil

### UnsetIsDiscontinued
`func (o *TargetProgramsRow) UnsetIsDiscontinued()`

UnsetIsDiscontinued ensures that no value is present for IsDiscontinued, not even an explicit nil
### GetNDiscontinued

`func (o *TargetProgramsRow) GetNDiscontinued() int32`

GetNDiscontinued returns the NDiscontinued field if non-nil, zero value otherwise.

### GetNDiscontinuedOk

`func (o *TargetProgramsRow) GetNDiscontinuedOk() (*int32, bool)`

GetNDiscontinuedOk returns a tuple with the NDiscontinued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDiscontinued

`func (o *TargetProgramsRow) SetNDiscontinued(v int32)`

SetNDiscontinued sets NDiscontinued field to given value.

### HasNDiscontinued

`func (o *TargetProgramsRow) HasNDiscontinued() bool`

HasNDiscontinued returns a boolean if a field has been set.

### SetNDiscontinuedNil

`func (o *TargetProgramsRow) SetNDiscontinuedNil(b bool)`

 SetNDiscontinuedNil sets the value for NDiscontinued to be an explicit nil

### UnsetNDiscontinued
`func (o *TargetProgramsRow) UnsetNDiscontinued()`

UnsetNDiscontinued ensures that no value is present for NDiscontinued, not even an explicit nil
### GetNTerminated

`func (o *TargetProgramsRow) GetNTerminated() int32`

GetNTerminated returns the NTerminated field if non-nil, zero value otherwise.

### GetNTerminatedOk

`func (o *TargetProgramsRow) GetNTerminatedOk() (*int32, bool)`

GetNTerminatedOk returns a tuple with the NTerminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTerminated

`func (o *TargetProgramsRow) SetNTerminated(v int32)`

SetNTerminated sets NTerminated field to given value.

### HasNTerminated

`func (o *TargetProgramsRow) HasNTerminated() bool`

HasNTerminated returns a boolean if a field has been set.

### SetNTerminatedNil

`func (o *TargetProgramsRow) SetNTerminatedNil(b bool)`

 SetNTerminatedNil sets the value for NTerminated to be an explicit nil

### UnsetNTerminated
`func (o *TargetProgramsRow) UnsetNTerminated()`

UnsetNTerminated ensures that no value is present for NTerminated, not even an explicit nil
### GetNWithdrawn

`func (o *TargetProgramsRow) GetNWithdrawn() int32`

GetNWithdrawn returns the NWithdrawn field if non-nil, zero value otherwise.

### GetNWithdrawnOk

`func (o *TargetProgramsRow) GetNWithdrawnOk() (*int32, bool)`

GetNWithdrawnOk returns a tuple with the NWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNWithdrawn

`func (o *TargetProgramsRow) SetNWithdrawn(v int32)`

SetNWithdrawn sets NWithdrawn field to given value.

### HasNWithdrawn

`func (o *TargetProgramsRow) HasNWithdrawn() bool`

HasNWithdrawn returns a boolean if a field has been set.

### SetNWithdrawnNil

`func (o *TargetProgramsRow) SetNWithdrawnNil(b bool)`

 SetNWithdrawnNil sets the value for NWithdrawn to be an explicit nil

### UnsetNWithdrawn
`func (o *TargetProgramsRow) UnsetNWithdrawn()`

UnsetNWithdrawn ensures that no value is present for NWithdrawn, not even an explicit nil
### GetNSuspended

`func (o *TargetProgramsRow) GetNSuspended() int32`

GetNSuspended returns the NSuspended field if non-nil, zero value otherwise.

### GetNSuspendedOk

`func (o *TargetProgramsRow) GetNSuspendedOk() (*int32, bool)`

GetNSuspendedOk returns a tuple with the NSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSuspended

`func (o *TargetProgramsRow) SetNSuspended(v int32)`

SetNSuspended sets NSuspended field to given value.

### HasNSuspended

`func (o *TargetProgramsRow) HasNSuspended() bool`

HasNSuspended returns a boolean if a field has been set.

### SetNSuspendedNil

`func (o *TargetProgramsRow) SetNSuspendedNil(b bool)`

 SetNSuspendedNil sets the value for NSuspended to be an explicit nil

### UnsetNSuspended
`func (o *TargetProgramsRow) UnsetNSuspended()`

UnsetNSuspended ensures that no value is present for NSuspended, not even an explicit nil
### GetMaxPhaseAtStop

`func (o *TargetProgramsRow) GetMaxPhaseAtStop() int32`

GetMaxPhaseAtStop returns the MaxPhaseAtStop field if non-nil, zero value otherwise.

### GetMaxPhaseAtStopOk

`func (o *TargetProgramsRow) GetMaxPhaseAtStopOk() (*int32, bool)`

GetMaxPhaseAtStopOk returns a tuple with the MaxPhaseAtStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhaseAtStop

`func (o *TargetProgramsRow) SetMaxPhaseAtStop(v int32)`

SetMaxPhaseAtStop sets MaxPhaseAtStop field to given value.

### HasMaxPhaseAtStop

`func (o *TargetProgramsRow) HasMaxPhaseAtStop() bool`

HasMaxPhaseAtStop returns a boolean if a field has been set.

### SetMaxPhaseAtStopNil

`func (o *TargetProgramsRow) SetMaxPhaseAtStopNil(b bool)`

 SetMaxPhaseAtStopNil sets the value for MaxPhaseAtStop to be an explicit nil

### UnsetMaxPhaseAtStop
`func (o *TargetProgramsRow) UnsetMaxPhaseAtStop()`

UnsetMaxPhaseAtStop ensures that no value is present for MaxPhaseAtStop, not even an explicit nil
### GetLastStopYear

`func (o *TargetProgramsRow) GetLastStopYear() int32`

GetLastStopYear returns the LastStopYear field if non-nil, zero value otherwise.

### GetLastStopYearOk

`func (o *TargetProgramsRow) GetLastStopYearOk() (*int32, bool)`

GetLastStopYearOk returns a tuple with the LastStopYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStopYear

`func (o *TargetProgramsRow) SetLastStopYear(v int32)`

SetLastStopYear sets LastStopYear field to given value.

### HasLastStopYear

`func (o *TargetProgramsRow) HasLastStopYear() bool`

HasLastStopYear returns a boolean if a field has been set.

### SetLastStopYearNil

`func (o *TargetProgramsRow) SetLastStopYearNil(b bool)`

 SetLastStopYearNil sets the value for LastStopYear to be an explicit nil

### UnsetLastStopYear
`func (o *TargetProgramsRow) UnsetLastStopYear()`

UnsetLastStopYear ensures that no value is present for LastStopYear, not even an explicit nil
### GetWhyStoppedReasons

`func (o *TargetProgramsRow) GetWhyStoppedReasons() []string`

GetWhyStoppedReasons returns the WhyStoppedReasons field if non-nil, zero value otherwise.

### GetWhyStoppedReasonsOk

`func (o *TargetProgramsRow) GetWhyStoppedReasonsOk() (*[]string, bool)`

GetWhyStoppedReasonsOk returns a tuple with the WhyStoppedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhyStoppedReasons

`func (o *TargetProgramsRow) SetWhyStoppedReasons(v []string)`

SetWhyStoppedReasons sets WhyStoppedReasons field to given value.

### HasWhyStoppedReasons

`func (o *TargetProgramsRow) HasWhyStoppedReasons() bool`

HasWhyStoppedReasons returns a boolean if a field has been set.

### SetWhyStoppedReasonsNil

`func (o *TargetProgramsRow) SetWhyStoppedReasonsNil(b bool)`

 SetWhyStoppedReasonsNil sets the value for WhyStoppedReasons to be an explicit nil

### UnsetWhyStoppedReasons
`func (o *TargetProgramsRow) UnsetWhyStoppedReasons()`

UnsetWhyStoppedReasons ensures that no value is present for WhyStoppedReasons, not even an explicit nil
### GetWhyStoppedCategories

`func (o *TargetProgramsRow) GetWhyStoppedCategories() []string`

GetWhyStoppedCategories returns the WhyStoppedCategories field if non-nil, zero value otherwise.

### GetWhyStoppedCategoriesOk

`func (o *TargetProgramsRow) GetWhyStoppedCategoriesOk() (*[]string, bool)`

GetWhyStoppedCategoriesOk returns a tuple with the WhyStoppedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhyStoppedCategories

`func (o *TargetProgramsRow) SetWhyStoppedCategories(v []string)`

SetWhyStoppedCategories sets WhyStoppedCategories field to given value.

### HasWhyStoppedCategories

`func (o *TargetProgramsRow) HasWhyStoppedCategories() bool`

HasWhyStoppedCategories returns a boolean if a field has been set.

### SetWhyStoppedCategoriesNil

`func (o *TargetProgramsRow) SetWhyStoppedCategoriesNil(b bool)`

 SetWhyStoppedCategoriesNil sets the value for WhyStoppedCategories to be an explicit nil

### UnsetWhyStoppedCategories
`func (o *TargetProgramsRow) UnsetWhyStoppedCategories()`

UnsetWhyStoppedCategories ensures that no value is present for WhyStoppedCategories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



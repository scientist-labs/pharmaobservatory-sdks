# ConditionDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s MeSH heading. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Trials in the roll-up, counted under each condition they are tagged with and each condition above it in the MeSH tree. | [optional] 
**NTrialsWithObservedPlan** | Pointer to **NullableInt64** | Trials whose first recorded plan is known. | [optional] 
**NTrialsWithEnrollmentDelta** | Pointer to **NullableInt64** | Trials graded on enrollment. | [optional] 
**NTrialsWithCompletionDelta** | Pointer to **NullableInt64** | Trials graded on completion date. | [optional] 
**NEnrollmentMet** | Pointer to **NullableInt64** | Graded trials that met their planned enrollment. | [optional] 
**NEnrollmentMissed** | Pointer to **NullableInt64** | Graded trials that missed their planned enrollment. | [optional] 
**EnrollmentHitRate** | Pointer to **NullableFloat32** | Share of graded trials that met planned enrollment (0 to 1). | [optional] 
**MedianEnrollmentAttainmentPct** | Pointer to **NullableFloat64** | Median enrollment attained as a percent of the plan. | [optional] 
**NOnTime** | Pointer to **NullableInt64** | Graded trials that completed on or before the planned date. | [optional] 
**NSlipped** | Pointer to **NullableInt64** | Graded trials that completed after the planned date. | [optional] 
**OnTimeRate** | Pointer to **NullableFloat32** | Share of graded trials that completed on time (0 to 1). | [optional] 
**MedianCompletionSlipDays** | Pointer to **NullableFloat64** | Median days between planned and actual primary completion, among graded trials; negative when earlier than planned. | [optional] 

## Methods

### NewConditionDeliveryRow

`func NewConditionDeliveryRow() *ConditionDeliveryRow`

NewConditionDeliveryRow instantiates a new ConditionDeliveryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionDeliveryRowWithDefaults

`func NewConditionDeliveryRowWithDefaults() *ConditionDeliveryRow`

NewConditionDeliveryRowWithDefaults instantiates a new ConditionDeliveryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionDeliveryRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionDeliveryRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionDeliveryRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionDeliveryRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionDeliveryRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionDeliveryRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ConditionDeliveryRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ConditionDeliveryRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ConditionDeliveryRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ConditionDeliveryRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ConditionDeliveryRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ConditionDeliveryRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetNTrials

`func (o *ConditionDeliveryRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ConditionDeliveryRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ConditionDeliveryRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ConditionDeliveryRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ConditionDeliveryRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ConditionDeliveryRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNTrialsWithObservedPlan

`func (o *ConditionDeliveryRow) GetNTrialsWithObservedPlan() int64`

GetNTrialsWithObservedPlan returns the NTrialsWithObservedPlan field if non-nil, zero value otherwise.

### GetNTrialsWithObservedPlanOk

`func (o *ConditionDeliveryRow) GetNTrialsWithObservedPlanOk() (*int64, bool)`

GetNTrialsWithObservedPlanOk returns a tuple with the NTrialsWithObservedPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsWithObservedPlan

`func (o *ConditionDeliveryRow) SetNTrialsWithObservedPlan(v int64)`

SetNTrialsWithObservedPlan sets NTrialsWithObservedPlan field to given value.

### HasNTrialsWithObservedPlan

`func (o *ConditionDeliveryRow) HasNTrialsWithObservedPlan() bool`

HasNTrialsWithObservedPlan returns a boolean if a field has been set.

### SetNTrialsWithObservedPlanNil

`func (o *ConditionDeliveryRow) SetNTrialsWithObservedPlanNil(b bool)`

 SetNTrialsWithObservedPlanNil sets the value for NTrialsWithObservedPlan to be an explicit nil

### UnsetNTrialsWithObservedPlan
`func (o *ConditionDeliveryRow) UnsetNTrialsWithObservedPlan()`

UnsetNTrialsWithObservedPlan ensures that no value is present for NTrialsWithObservedPlan, not even an explicit nil
### GetNTrialsWithEnrollmentDelta

`func (o *ConditionDeliveryRow) GetNTrialsWithEnrollmentDelta() int64`

GetNTrialsWithEnrollmentDelta returns the NTrialsWithEnrollmentDelta field if non-nil, zero value otherwise.

### GetNTrialsWithEnrollmentDeltaOk

`func (o *ConditionDeliveryRow) GetNTrialsWithEnrollmentDeltaOk() (*int64, bool)`

GetNTrialsWithEnrollmentDeltaOk returns a tuple with the NTrialsWithEnrollmentDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsWithEnrollmentDelta

`func (o *ConditionDeliveryRow) SetNTrialsWithEnrollmentDelta(v int64)`

SetNTrialsWithEnrollmentDelta sets NTrialsWithEnrollmentDelta field to given value.

### HasNTrialsWithEnrollmentDelta

`func (o *ConditionDeliveryRow) HasNTrialsWithEnrollmentDelta() bool`

HasNTrialsWithEnrollmentDelta returns a boolean if a field has been set.

### SetNTrialsWithEnrollmentDeltaNil

`func (o *ConditionDeliveryRow) SetNTrialsWithEnrollmentDeltaNil(b bool)`

 SetNTrialsWithEnrollmentDeltaNil sets the value for NTrialsWithEnrollmentDelta to be an explicit nil

### UnsetNTrialsWithEnrollmentDelta
`func (o *ConditionDeliveryRow) UnsetNTrialsWithEnrollmentDelta()`

UnsetNTrialsWithEnrollmentDelta ensures that no value is present for NTrialsWithEnrollmentDelta, not even an explicit nil
### GetNTrialsWithCompletionDelta

`func (o *ConditionDeliveryRow) GetNTrialsWithCompletionDelta() int64`

GetNTrialsWithCompletionDelta returns the NTrialsWithCompletionDelta field if non-nil, zero value otherwise.

### GetNTrialsWithCompletionDeltaOk

`func (o *ConditionDeliveryRow) GetNTrialsWithCompletionDeltaOk() (*int64, bool)`

GetNTrialsWithCompletionDeltaOk returns a tuple with the NTrialsWithCompletionDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsWithCompletionDelta

`func (o *ConditionDeliveryRow) SetNTrialsWithCompletionDelta(v int64)`

SetNTrialsWithCompletionDelta sets NTrialsWithCompletionDelta field to given value.

### HasNTrialsWithCompletionDelta

`func (o *ConditionDeliveryRow) HasNTrialsWithCompletionDelta() bool`

HasNTrialsWithCompletionDelta returns a boolean if a field has been set.

### SetNTrialsWithCompletionDeltaNil

`func (o *ConditionDeliveryRow) SetNTrialsWithCompletionDeltaNil(b bool)`

 SetNTrialsWithCompletionDeltaNil sets the value for NTrialsWithCompletionDelta to be an explicit nil

### UnsetNTrialsWithCompletionDelta
`func (o *ConditionDeliveryRow) UnsetNTrialsWithCompletionDelta()`

UnsetNTrialsWithCompletionDelta ensures that no value is present for NTrialsWithCompletionDelta, not even an explicit nil
### GetNEnrollmentMet

`func (o *ConditionDeliveryRow) GetNEnrollmentMet() int64`

GetNEnrollmentMet returns the NEnrollmentMet field if non-nil, zero value otherwise.

### GetNEnrollmentMetOk

`func (o *ConditionDeliveryRow) GetNEnrollmentMetOk() (*int64, bool)`

GetNEnrollmentMetOk returns a tuple with the NEnrollmentMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEnrollmentMet

`func (o *ConditionDeliveryRow) SetNEnrollmentMet(v int64)`

SetNEnrollmentMet sets NEnrollmentMet field to given value.

### HasNEnrollmentMet

`func (o *ConditionDeliveryRow) HasNEnrollmentMet() bool`

HasNEnrollmentMet returns a boolean if a field has been set.

### SetNEnrollmentMetNil

`func (o *ConditionDeliveryRow) SetNEnrollmentMetNil(b bool)`

 SetNEnrollmentMetNil sets the value for NEnrollmentMet to be an explicit nil

### UnsetNEnrollmentMet
`func (o *ConditionDeliveryRow) UnsetNEnrollmentMet()`

UnsetNEnrollmentMet ensures that no value is present for NEnrollmentMet, not even an explicit nil
### GetNEnrollmentMissed

`func (o *ConditionDeliveryRow) GetNEnrollmentMissed() int64`

GetNEnrollmentMissed returns the NEnrollmentMissed field if non-nil, zero value otherwise.

### GetNEnrollmentMissedOk

`func (o *ConditionDeliveryRow) GetNEnrollmentMissedOk() (*int64, bool)`

GetNEnrollmentMissedOk returns a tuple with the NEnrollmentMissed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEnrollmentMissed

`func (o *ConditionDeliveryRow) SetNEnrollmentMissed(v int64)`

SetNEnrollmentMissed sets NEnrollmentMissed field to given value.

### HasNEnrollmentMissed

`func (o *ConditionDeliveryRow) HasNEnrollmentMissed() bool`

HasNEnrollmentMissed returns a boolean if a field has been set.

### SetNEnrollmentMissedNil

`func (o *ConditionDeliveryRow) SetNEnrollmentMissedNil(b bool)`

 SetNEnrollmentMissedNil sets the value for NEnrollmentMissed to be an explicit nil

### UnsetNEnrollmentMissed
`func (o *ConditionDeliveryRow) UnsetNEnrollmentMissed()`

UnsetNEnrollmentMissed ensures that no value is present for NEnrollmentMissed, not even an explicit nil
### GetEnrollmentHitRate

`func (o *ConditionDeliveryRow) GetEnrollmentHitRate() float32`

GetEnrollmentHitRate returns the EnrollmentHitRate field if non-nil, zero value otherwise.

### GetEnrollmentHitRateOk

`func (o *ConditionDeliveryRow) GetEnrollmentHitRateOk() (*float32, bool)`

GetEnrollmentHitRateOk returns a tuple with the EnrollmentHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentHitRate

`func (o *ConditionDeliveryRow) SetEnrollmentHitRate(v float32)`

SetEnrollmentHitRate sets EnrollmentHitRate field to given value.

### HasEnrollmentHitRate

`func (o *ConditionDeliveryRow) HasEnrollmentHitRate() bool`

HasEnrollmentHitRate returns a boolean if a field has been set.

### SetEnrollmentHitRateNil

`func (o *ConditionDeliveryRow) SetEnrollmentHitRateNil(b bool)`

 SetEnrollmentHitRateNil sets the value for EnrollmentHitRate to be an explicit nil

### UnsetEnrollmentHitRate
`func (o *ConditionDeliveryRow) UnsetEnrollmentHitRate()`

UnsetEnrollmentHitRate ensures that no value is present for EnrollmentHitRate, not even an explicit nil
### GetMedianEnrollmentAttainmentPct

`func (o *ConditionDeliveryRow) GetMedianEnrollmentAttainmentPct() float64`

GetMedianEnrollmentAttainmentPct returns the MedianEnrollmentAttainmentPct field if non-nil, zero value otherwise.

### GetMedianEnrollmentAttainmentPctOk

`func (o *ConditionDeliveryRow) GetMedianEnrollmentAttainmentPctOk() (*float64, bool)`

GetMedianEnrollmentAttainmentPctOk returns a tuple with the MedianEnrollmentAttainmentPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianEnrollmentAttainmentPct

`func (o *ConditionDeliveryRow) SetMedianEnrollmentAttainmentPct(v float64)`

SetMedianEnrollmentAttainmentPct sets MedianEnrollmentAttainmentPct field to given value.

### HasMedianEnrollmentAttainmentPct

`func (o *ConditionDeliveryRow) HasMedianEnrollmentAttainmentPct() bool`

HasMedianEnrollmentAttainmentPct returns a boolean if a field has been set.

### SetMedianEnrollmentAttainmentPctNil

`func (o *ConditionDeliveryRow) SetMedianEnrollmentAttainmentPctNil(b bool)`

 SetMedianEnrollmentAttainmentPctNil sets the value for MedianEnrollmentAttainmentPct to be an explicit nil

### UnsetMedianEnrollmentAttainmentPct
`func (o *ConditionDeliveryRow) UnsetMedianEnrollmentAttainmentPct()`

UnsetMedianEnrollmentAttainmentPct ensures that no value is present for MedianEnrollmentAttainmentPct, not even an explicit nil
### GetNOnTime

`func (o *ConditionDeliveryRow) GetNOnTime() int64`

GetNOnTime returns the NOnTime field if non-nil, zero value otherwise.

### GetNOnTimeOk

`func (o *ConditionDeliveryRow) GetNOnTimeOk() (*int64, bool)`

GetNOnTimeOk returns a tuple with the NOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOnTime

`func (o *ConditionDeliveryRow) SetNOnTime(v int64)`

SetNOnTime sets NOnTime field to given value.

### HasNOnTime

`func (o *ConditionDeliveryRow) HasNOnTime() bool`

HasNOnTime returns a boolean if a field has been set.

### SetNOnTimeNil

`func (o *ConditionDeliveryRow) SetNOnTimeNil(b bool)`

 SetNOnTimeNil sets the value for NOnTime to be an explicit nil

### UnsetNOnTime
`func (o *ConditionDeliveryRow) UnsetNOnTime()`

UnsetNOnTime ensures that no value is present for NOnTime, not even an explicit nil
### GetNSlipped

`func (o *ConditionDeliveryRow) GetNSlipped() int64`

GetNSlipped returns the NSlipped field if non-nil, zero value otherwise.

### GetNSlippedOk

`func (o *ConditionDeliveryRow) GetNSlippedOk() (*int64, bool)`

GetNSlippedOk returns a tuple with the NSlipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSlipped

`func (o *ConditionDeliveryRow) SetNSlipped(v int64)`

SetNSlipped sets NSlipped field to given value.

### HasNSlipped

`func (o *ConditionDeliveryRow) HasNSlipped() bool`

HasNSlipped returns a boolean if a field has been set.

### SetNSlippedNil

`func (o *ConditionDeliveryRow) SetNSlippedNil(b bool)`

 SetNSlippedNil sets the value for NSlipped to be an explicit nil

### UnsetNSlipped
`func (o *ConditionDeliveryRow) UnsetNSlipped()`

UnsetNSlipped ensures that no value is present for NSlipped, not even an explicit nil
### GetOnTimeRate

`func (o *ConditionDeliveryRow) GetOnTimeRate() float32`

GetOnTimeRate returns the OnTimeRate field if non-nil, zero value otherwise.

### GetOnTimeRateOk

`func (o *ConditionDeliveryRow) GetOnTimeRateOk() (*float32, bool)`

GetOnTimeRateOk returns a tuple with the OnTimeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnTimeRate

`func (o *ConditionDeliveryRow) SetOnTimeRate(v float32)`

SetOnTimeRate sets OnTimeRate field to given value.

### HasOnTimeRate

`func (o *ConditionDeliveryRow) HasOnTimeRate() bool`

HasOnTimeRate returns a boolean if a field has been set.

### SetOnTimeRateNil

`func (o *ConditionDeliveryRow) SetOnTimeRateNil(b bool)`

 SetOnTimeRateNil sets the value for OnTimeRate to be an explicit nil

### UnsetOnTimeRate
`func (o *ConditionDeliveryRow) UnsetOnTimeRate()`

UnsetOnTimeRate ensures that no value is present for OnTimeRate, not even an explicit nil
### GetMedianCompletionSlipDays

`func (o *ConditionDeliveryRow) GetMedianCompletionSlipDays() float64`

GetMedianCompletionSlipDays returns the MedianCompletionSlipDays field if non-nil, zero value otherwise.

### GetMedianCompletionSlipDaysOk

`func (o *ConditionDeliveryRow) GetMedianCompletionSlipDaysOk() (*float64, bool)`

GetMedianCompletionSlipDaysOk returns a tuple with the MedianCompletionSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianCompletionSlipDays

`func (o *ConditionDeliveryRow) SetMedianCompletionSlipDays(v float64)`

SetMedianCompletionSlipDays sets MedianCompletionSlipDays field to given value.

### HasMedianCompletionSlipDays

`func (o *ConditionDeliveryRow) HasMedianCompletionSlipDays() bool`

HasMedianCompletionSlipDays returns a boolean if a field has been set.

### SetMedianCompletionSlipDaysNil

`func (o *ConditionDeliveryRow) SetMedianCompletionSlipDaysNil(b bool)`

 SetMedianCompletionSlipDaysNil sets the value for MedianCompletionSlipDays to be an explicit nil

### UnsetMedianCompletionSlipDays
`func (o *ConditionDeliveryRow) UnsetMedianCompletionSlipDays()`

UnsetMedianCompletionSlipDays ensures that no value is present for MedianCompletionSlipDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



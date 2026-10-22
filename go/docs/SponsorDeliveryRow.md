# SponsorDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the lead sponsor organization record. | [optional] 
**SponsorName** | Pointer to **NullableString** | Display name of the sponsor organization record. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the sponsor belongs to (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**NTrialsTotal** | Pointer to **NullableInt64** | Lead-sponsored trials of the organization. | [optional] 
**NTrialsWithObservedPlan** | Pointer to **NullableInt64** | Trials whose planned dates come from the earliest estimate in the registry&#39;s own history. | [optional] 
**NTrialsWithEnrollmentDelta** | Pointer to **NullableInt64** | Trials with a measured difference between planned and actual enrollment; the base of the enrollment rate. | [optional] 
**NTrialsWithCompletionDelta** | Pointer to **NullableInt64** | Trials with a measured difference between planned and actual primary completion; the base of the on-time rate. | [optional] 
**NEnrollmentMet** | Pointer to **NullableInt64** | Trials with an enrollment delta that reached their planned enrollment. | [optional] 
**NEnrollmentMissed** | Pointer to **NullableInt64** | Trials with an enrollment delta that did not reach their planned enrollment. | [optional] 
**EnrollmentHitRate** | Pointer to **NullableFloat32** | n_enrollment_met divided by n_trials_with_enrollment_delta; null when that base is zero. | [optional] 
**MedianEnrollmentAttainmentPct** | Pointer to **NullableFloat64** | Median of actual enrollment as a percentage of planned, over trials with an enrollment delta. | [optional] 
**NOntime** | Pointer to **NullableInt64** | Trials with a completion delta that were not delayed past their planned primary completion. | [optional] 
**NSlipped** | Pointer to **NullableInt64** | Trials with a completion delta that completed or are projected after their planned primary completion. | [optional] 
**OnTimeRate** | Pointer to **NullableFloat32** | n_ontime divided by n_trials_with_completion_delta; null when that base is zero. | [optional] 
**MedianCompletionSlipDays** | Pointer to **NullableFloat64** | Median days between planned and actual primary completion, over trials with a completion delta. | [optional] 

## Methods

### NewSponsorDeliveryRow

`func NewSponsorDeliveryRow() *SponsorDeliveryRow`

NewSponsorDeliveryRow instantiates a new SponsorDeliveryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSponsorDeliveryRowWithDefaults

`func NewSponsorDeliveryRowWithDefaults() *SponsorDeliveryRow`

NewSponsorDeliveryRowWithDefaults instantiates a new SponsorDeliveryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponsorOrgId

`func (o *SponsorDeliveryRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *SponsorDeliveryRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *SponsorDeliveryRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *SponsorDeliveryRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *SponsorDeliveryRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *SponsorDeliveryRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorName

`func (o *SponsorDeliveryRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *SponsorDeliveryRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *SponsorDeliveryRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *SponsorDeliveryRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *SponsorDeliveryRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *SponsorDeliveryRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorCompanyId

`func (o *SponsorDeliveryRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *SponsorDeliveryRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *SponsorDeliveryRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *SponsorDeliveryRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *SponsorDeliveryRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *SponsorDeliveryRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *SponsorDeliveryRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *SponsorDeliveryRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *SponsorDeliveryRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *SponsorDeliveryRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *SponsorDeliveryRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *SponsorDeliveryRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *SponsorDeliveryRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *SponsorDeliveryRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *SponsorDeliveryRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *SponsorDeliveryRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *SponsorDeliveryRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *SponsorDeliveryRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *SponsorDeliveryRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *SponsorDeliveryRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *SponsorDeliveryRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *SponsorDeliveryRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *SponsorDeliveryRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *SponsorDeliveryRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetNTrialsTotal

`func (o *SponsorDeliveryRow) GetNTrialsTotal() int64`

GetNTrialsTotal returns the NTrialsTotal field if non-nil, zero value otherwise.

### GetNTrialsTotalOk

`func (o *SponsorDeliveryRow) GetNTrialsTotalOk() (*int64, bool)`

GetNTrialsTotalOk returns a tuple with the NTrialsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsTotal

`func (o *SponsorDeliveryRow) SetNTrialsTotal(v int64)`

SetNTrialsTotal sets NTrialsTotal field to given value.

### HasNTrialsTotal

`func (o *SponsorDeliveryRow) HasNTrialsTotal() bool`

HasNTrialsTotal returns a boolean if a field has been set.

### SetNTrialsTotalNil

`func (o *SponsorDeliveryRow) SetNTrialsTotalNil(b bool)`

 SetNTrialsTotalNil sets the value for NTrialsTotal to be an explicit nil

### UnsetNTrialsTotal
`func (o *SponsorDeliveryRow) UnsetNTrialsTotal()`

UnsetNTrialsTotal ensures that no value is present for NTrialsTotal, not even an explicit nil
### GetNTrialsWithObservedPlan

`func (o *SponsorDeliveryRow) GetNTrialsWithObservedPlan() int64`

GetNTrialsWithObservedPlan returns the NTrialsWithObservedPlan field if non-nil, zero value otherwise.

### GetNTrialsWithObservedPlanOk

`func (o *SponsorDeliveryRow) GetNTrialsWithObservedPlanOk() (*int64, bool)`

GetNTrialsWithObservedPlanOk returns a tuple with the NTrialsWithObservedPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsWithObservedPlan

`func (o *SponsorDeliveryRow) SetNTrialsWithObservedPlan(v int64)`

SetNTrialsWithObservedPlan sets NTrialsWithObservedPlan field to given value.

### HasNTrialsWithObservedPlan

`func (o *SponsorDeliveryRow) HasNTrialsWithObservedPlan() bool`

HasNTrialsWithObservedPlan returns a boolean if a field has been set.

### SetNTrialsWithObservedPlanNil

`func (o *SponsorDeliveryRow) SetNTrialsWithObservedPlanNil(b bool)`

 SetNTrialsWithObservedPlanNil sets the value for NTrialsWithObservedPlan to be an explicit nil

### UnsetNTrialsWithObservedPlan
`func (o *SponsorDeliveryRow) UnsetNTrialsWithObservedPlan()`

UnsetNTrialsWithObservedPlan ensures that no value is present for NTrialsWithObservedPlan, not even an explicit nil
### GetNTrialsWithEnrollmentDelta

`func (o *SponsorDeliveryRow) GetNTrialsWithEnrollmentDelta() int64`

GetNTrialsWithEnrollmentDelta returns the NTrialsWithEnrollmentDelta field if non-nil, zero value otherwise.

### GetNTrialsWithEnrollmentDeltaOk

`func (o *SponsorDeliveryRow) GetNTrialsWithEnrollmentDeltaOk() (*int64, bool)`

GetNTrialsWithEnrollmentDeltaOk returns a tuple with the NTrialsWithEnrollmentDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsWithEnrollmentDelta

`func (o *SponsorDeliveryRow) SetNTrialsWithEnrollmentDelta(v int64)`

SetNTrialsWithEnrollmentDelta sets NTrialsWithEnrollmentDelta field to given value.

### HasNTrialsWithEnrollmentDelta

`func (o *SponsorDeliveryRow) HasNTrialsWithEnrollmentDelta() bool`

HasNTrialsWithEnrollmentDelta returns a boolean if a field has been set.

### SetNTrialsWithEnrollmentDeltaNil

`func (o *SponsorDeliveryRow) SetNTrialsWithEnrollmentDeltaNil(b bool)`

 SetNTrialsWithEnrollmentDeltaNil sets the value for NTrialsWithEnrollmentDelta to be an explicit nil

### UnsetNTrialsWithEnrollmentDelta
`func (o *SponsorDeliveryRow) UnsetNTrialsWithEnrollmentDelta()`

UnsetNTrialsWithEnrollmentDelta ensures that no value is present for NTrialsWithEnrollmentDelta, not even an explicit nil
### GetNTrialsWithCompletionDelta

`func (o *SponsorDeliveryRow) GetNTrialsWithCompletionDelta() int64`

GetNTrialsWithCompletionDelta returns the NTrialsWithCompletionDelta field if non-nil, zero value otherwise.

### GetNTrialsWithCompletionDeltaOk

`func (o *SponsorDeliveryRow) GetNTrialsWithCompletionDeltaOk() (*int64, bool)`

GetNTrialsWithCompletionDeltaOk returns a tuple with the NTrialsWithCompletionDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsWithCompletionDelta

`func (o *SponsorDeliveryRow) SetNTrialsWithCompletionDelta(v int64)`

SetNTrialsWithCompletionDelta sets NTrialsWithCompletionDelta field to given value.

### HasNTrialsWithCompletionDelta

`func (o *SponsorDeliveryRow) HasNTrialsWithCompletionDelta() bool`

HasNTrialsWithCompletionDelta returns a boolean if a field has been set.

### SetNTrialsWithCompletionDeltaNil

`func (o *SponsorDeliveryRow) SetNTrialsWithCompletionDeltaNil(b bool)`

 SetNTrialsWithCompletionDeltaNil sets the value for NTrialsWithCompletionDelta to be an explicit nil

### UnsetNTrialsWithCompletionDelta
`func (o *SponsorDeliveryRow) UnsetNTrialsWithCompletionDelta()`

UnsetNTrialsWithCompletionDelta ensures that no value is present for NTrialsWithCompletionDelta, not even an explicit nil
### GetNEnrollmentMet

`func (o *SponsorDeliveryRow) GetNEnrollmentMet() int64`

GetNEnrollmentMet returns the NEnrollmentMet field if non-nil, zero value otherwise.

### GetNEnrollmentMetOk

`func (o *SponsorDeliveryRow) GetNEnrollmentMetOk() (*int64, bool)`

GetNEnrollmentMetOk returns a tuple with the NEnrollmentMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEnrollmentMet

`func (o *SponsorDeliveryRow) SetNEnrollmentMet(v int64)`

SetNEnrollmentMet sets NEnrollmentMet field to given value.

### HasNEnrollmentMet

`func (o *SponsorDeliveryRow) HasNEnrollmentMet() bool`

HasNEnrollmentMet returns a boolean if a field has been set.

### SetNEnrollmentMetNil

`func (o *SponsorDeliveryRow) SetNEnrollmentMetNil(b bool)`

 SetNEnrollmentMetNil sets the value for NEnrollmentMet to be an explicit nil

### UnsetNEnrollmentMet
`func (o *SponsorDeliveryRow) UnsetNEnrollmentMet()`

UnsetNEnrollmentMet ensures that no value is present for NEnrollmentMet, not even an explicit nil
### GetNEnrollmentMissed

`func (o *SponsorDeliveryRow) GetNEnrollmentMissed() int64`

GetNEnrollmentMissed returns the NEnrollmentMissed field if non-nil, zero value otherwise.

### GetNEnrollmentMissedOk

`func (o *SponsorDeliveryRow) GetNEnrollmentMissedOk() (*int64, bool)`

GetNEnrollmentMissedOk returns a tuple with the NEnrollmentMissed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEnrollmentMissed

`func (o *SponsorDeliveryRow) SetNEnrollmentMissed(v int64)`

SetNEnrollmentMissed sets NEnrollmentMissed field to given value.

### HasNEnrollmentMissed

`func (o *SponsorDeliveryRow) HasNEnrollmentMissed() bool`

HasNEnrollmentMissed returns a boolean if a field has been set.

### SetNEnrollmentMissedNil

`func (o *SponsorDeliveryRow) SetNEnrollmentMissedNil(b bool)`

 SetNEnrollmentMissedNil sets the value for NEnrollmentMissed to be an explicit nil

### UnsetNEnrollmentMissed
`func (o *SponsorDeliveryRow) UnsetNEnrollmentMissed()`

UnsetNEnrollmentMissed ensures that no value is present for NEnrollmentMissed, not even an explicit nil
### GetEnrollmentHitRate

`func (o *SponsorDeliveryRow) GetEnrollmentHitRate() float32`

GetEnrollmentHitRate returns the EnrollmentHitRate field if non-nil, zero value otherwise.

### GetEnrollmentHitRateOk

`func (o *SponsorDeliveryRow) GetEnrollmentHitRateOk() (*float32, bool)`

GetEnrollmentHitRateOk returns a tuple with the EnrollmentHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentHitRate

`func (o *SponsorDeliveryRow) SetEnrollmentHitRate(v float32)`

SetEnrollmentHitRate sets EnrollmentHitRate field to given value.

### HasEnrollmentHitRate

`func (o *SponsorDeliveryRow) HasEnrollmentHitRate() bool`

HasEnrollmentHitRate returns a boolean if a field has been set.

### SetEnrollmentHitRateNil

`func (o *SponsorDeliveryRow) SetEnrollmentHitRateNil(b bool)`

 SetEnrollmentHitRateNil sets the value for EnrollmentHitRate to be an explicit nil

### UnsetEnrollmentHitRate
`func (o *SponsorDeliveryRow) UnsetEnrollmentHitRate()`

UnsetEnrollmentHitRate ensures that no value is present for EnrollmentHitRate, not even an explicit nil
### GetMedianEnrollmentAttainmentPct

`func (o *SponsorDeliveryRow) GetMedianEnrollmentAttainmentPct() float64`

GetMedianEnrollmentAttainmentPct returns the MedianEnrollmentAttainmentPct field if non-nil, zero value otherwise.

### GetMedianEnrollmentAttainmentPctOk

`func (o *SponsorDeliveryRow) GetMedianEnrollmentAttainmentPctOk() (*float64, bool)`

GetMedianEnrollmentAttainmentPctOk returns a tuple with the MedianEnrollmentAttainmentPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianEnrollmentAttainmentPct

`func (o *SponsorDeliveryRow) SetMedianEnrollmentAttainmentPct(v float64)`

SetMedianEnrollmentAttainmentPct sets MedianEnrollmentAttainmentPct field to given value.

### HasMedianEnrollmentAttainmentPct

`func (o *SponsorDeliveryRow) HasMedianEnrollmentAttainmentPct() bool`

HasMedianEnrollmentAttainmentPct returns a boolean if a field has been set.

### SetMedianEnrollmentAttainmentPctNil

`func (o *SponsorDeliveryRow) SetMedianEnrollmentAttainmentPctNil(b bool)`

 SetMedianEnrollmentAttainmentPctNil sets the value for MedianEnrollmentAttainmentPct to be an explicit nil

### UnsetMedianEnrollmentAttainmentPct
`func (o *SponsorDeliveryRow) UnsetMedianEnrollmentAttainmentPct()`

UnsetMedianEnrollmentAttainmentPct ensures that no value is present for MedianEnrollmentAttainmentPct, not even an explicit nil
### GetNOntime

`func (o *SponsorDeliveryRow) GetNOntime() int64`

GetNOntime returns the NOntime field if non-nil, zero value otherwise.

### GetNOntimeOk

`func (o *SponsorDeliveryRow) GetNOntimeOk() (*int64, bool)`

GetNOntimeOk returns a tuple with the NOntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOntime

`func (o *SponsorDeliveryRow) SetNOntime(v int64)`

SetNOntime sets NOntime field to given value.

### HasNOntime

`func (o *SponsorDeliveryRow) HasNOntime() bool`

HasNOntime returns a boolean if a field has been set.

### SetNOntimeNil

`func (o *SponsorDeliveryRow) SetNOntimeNil(b bool)`

 SetNOntimeNil sets the value for NOntime to be an explicit nil

### UnsetNOntime
`func (o *SponsorDeliveryRow) UnsetNOntime()`

UnsetNOntime ensures that no value is present for NOntime, not even an explicit nil
### GetNSlipped

`func (o *SponsorDeliveryRow) GetNSlipped() int64`

GetNSlipped returns the NSlipped field if non-nil, zero value otherwise.

### GetNSlippedOk

`func (o *SponsorDeliveryRow) GetNSlippedOk() (*int64, bool)`

GetNSlippedOk returns a tuple with the NSlipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSlipped

`func (o *SponsorDeliveryRow) SetNSlipped(v int64)`

SetNSlipped sets NSlipped field to given value.

### HasNSlipped

`func (o *SponsorDeliveryRow) HasNSlipped() bool`

HasNSlipped returns a boolean if a field has been set.

### SetNSlippedNil

`func (o *SponsorDeliveryRow) SetNSlippedNil(b bool)`

 SetNSlippedNil sets the value for NSlipped to be an explicit nil

### UnsetNSlipped
`func (o *SponsorDeliveryRow) UnsetNSlipped()`

UnsetNSlipped ensures that no value is present for NSlipped, not even an explicit nil
### GetOnTimeRate

`func (o *SponsorDeliveryRow) GetOnTimeRate() float32`

GetOnTimeRate returns the OnTimeRate field if non-nil, zero value otherwise.

### GetOnTimeRateOk

`func (o *SponsorDeliveryRow) GetOnTimeRateOk() (*float32, bool)`

GetOnTimeRateOk returns a tuple with the OnTimeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnTimeRate

`func (o *SponsorDeliveryRow) SetOnTimeRate(v float32)`

SetOnTimeRate sets OnTimeRate field to given value.

### HasOnTimeRate

`func (o *SponsorDeliveryRow) HasOnTimeRate() bool`

HasOnTimeRate returns a boolean if a field has been set.

### SetOnTimeRateNil

`func (o *SponsorDeliveryRow) SetOnTimeRateNil(b bool)`

 SetOnTimeRateNil sets the value for OnTimeRate to be an explicit nil

### UnsetOnTimeRate
`func (o *SponsorDeliveryRow) UnsetOnTimeRate()`

UnsetOnTimeRate ensures that no value is present for OnTimeRate, not even an explicit nil
### GetMedianCompletionSlipDays

`func (o *SponsorDeliveryRow) GetMedianCompletionSlipDays() float64`

GetMedianCompletionSlipDays returns the MedianCompletionSlipDays field if non-nil, zero value otherwise.

### GetMedianCompletionSlipDaysOk

`func (o *SponsorDeliveryRow) GetMedianCompletionSlipDaysOk() (*float64, bool)`

GetMedianCompletionSlipDaysOk returns a tuple with the MedianCompletionSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianCompletionSlipDays

`func (o *SponsorDeliveryRow) SetMedianCompletionSlipDays(v float64)`

SetMedianCompletionSlipDays sets MedianCompletionSlipDays field to given value.

### HasMedianCompletionSlipDays

`func (o *SponsorDeliveryRow) HasMedianCompletionSlipDays() bool`

HasMedianCompletionSlipDays returns a boolean if a field has been set.

### SetMedianCompletionSlipDaysNil

`func (o *SponsorDeliveryRow) SetMedianCompletionSlipDaysNil(b bool)`

 SetMedianCompletionSlipDaysNil sets the value for MedianCompletionSlipDays to be an explicit nil

### UnsetMedianCompletionSlipDays
`func (o *SponsorDeliveryRow) UnsetMedianCompletionSlipDays()`

UnsetMedianCompletionSlipDays ensures that no value is present for MedianCompletionSlipDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



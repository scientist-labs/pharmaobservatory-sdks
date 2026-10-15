# ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Note** | **string** |  | 
**LifetimeStudies** | **int32** |  | 
**RecentStudies** | **int32** |  | 
**OngoingStudies** | **int32** |  | 
**MandatedStudies** | **int32** |  | 
**MandatedShare** | **string** |  | 
**FinalReportPairs** | **int32** |  | 
**BackfilledPairs** | **int32** |  | 
**ComparablePairs** | **int32** |  | 
**BackfilledShare** | **string** |  | 
**MeetsNFloor** | **bool** |  | 
**NFloor** | **int32** |  | 
**MedianSlipDays** | **float32** |  | 
**P25SlipDays** | **float32** |  | 
**P75SlipDays** | **float32** |  | 
**LateOver30d** | **int32** |  | 
**SuppressionNote** | **interface{}** |  | 

## Methods

### NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary

`func NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary(label string, note string, lifetimeStudies int32, recentStudies int32, ongoingStudies int32, mandatedStudies int32, mandatedShare string, finalReportPairs int32, backfilledPairs int32, comparablePairs int32, backfilledShare string, meetsNFloor bool, nFloor int32, medianSlipDays float32, p25SlipDays float32, p75SlipDays float32, lateOver30d int32, suppressionNote interface{}, ) *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary`

NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary instantiates a new ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntaryWithDefaults

`func NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntaryWithDefaults() *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary`

NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntaryWithDefaults instantiates a new ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNote

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetNote(v string)`

SetNote sets Note field to given value.


### GetLifetimeStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetLifetimeStudies() int32`

GetLifetimeStudies returns the LifetimeStudies field if non-nil, zero value otherwise.

### GetLifetimeStudiesOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetLifetimeStudiesOk() (*int32, bool)`

GetLifetimeStudiesOk returns a tuple with the LifetimeStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetLifetimeStudies(v int32)`

SetLifetimeStudies sets LifetimeStudies field to given value.


### GetRecentStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetRecentStudies() int32`

GetRecentStudies returns the RecentStudies field if non-nil, zero value otherwise.

### GetRecentStudiesOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetRecentStudiesOk() (*int32, bool)`

GetRecentStudiesOk returns a tuple with the RecentStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetRecentStudies(v int32)`

SetRecentStudies sets RecentStudies field to given value.


### GetOngoingStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetOngoingStudies() int32`

GetOngoingStudies returns the OngoingStudies field if non-nil, zero value otherwise.

### GetOngoingStudiesOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetOngoingStudiesOk() (*int32, bool)`

GetOngoingStudiesOk returns a tuple with the OngoingStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoingStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetOngoingStudies(v int32)`

SetOngoingStudies sets OngoingStudies field to given value.


### GetMandatedStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetMandatedStudies() int32`

GetMandatedStudies returns the MandatedStudies field if non-nil, zero value otherwise.

### GetMandatedStudiesOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetMandatedStudiesOk() (*int32, bool)`

GetMandatedStudiesOk returns a tuple with the MandatedStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatedStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetMandatedStudies(v int32)`

SetMandatedStudies sets MandatedStudies field to given value.


### GetMandatedShare

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetMandatedShare() string`

GetMandatedShare returns the MandatedShare field if non-nil, zero value otherwise.

### GetMandatedShareOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetMandatedShareOk() (*string, bool)`

GetMandatedShareOk returns a tuple with the MandatedShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatedShare

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetMandatedShare(v string)`

SetMandatedShare sets MandatedShare field to given value.


### GetFinalReportPairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetFinalReportPairs() int32`

GetFinalReportPairs returns the FinalReportPairs field if non-nil, zero value otherwise.

### GetFinalReportPairsOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetFinalReportPairsOk() (*int32, bool)`

GetFinalReportPairsOk returns a tuple with the FinalReportPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportPairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetFinalReportPairs(v int32)`

SetFinalReportPairs sets FinalReportPairs field to given value.


### GetBackfilledPairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetBackfilledPairs() int32`

GetBackfilledPairs returns the BackfilledPairs field if non-nil, zero value otherwise.

### GetBackfilledPairsOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetBackfilledPairsOk() (*int32, bool)`

GetBackfilledPairsOk returns a tuple with the BackfilledPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilledPairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetBackfilledPairs(v int32)`

SetBackfilledPairs sets BackfilledPairs field to given value.


### GetComparablePairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetComparablePairs() int32`

GetComparablePairs returns the ComparablePairs field if non-nil, zero value otherwise.

### GetComparablePairsOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetComparablePairsOk() (*int32, bool)`

GetComparablePairsOk returns a tuple with the ComparablePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparablePairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetComparablePairs(v int32)`

SetComparablePairs sets ComparablePairs field to given value.


### GetBackfilledShare

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetBackfilledShare() string`

GetBackfilledShare returns the BackfilledShare field if non-nil, zero value otherwise.

### GetBackfilledShareOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetBackfilledShareOk() (*string, bool)`

GetBackfilledShareOk returns a tuple with the BackfilledShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilledShare

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetBackfilledShare(v string)`

SetBackfilledShare sets BackfilledShare field to given value.


### GetMeetsNFloor

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetMeetsNFloor() bool`

GetMeetsNFloor returns the MeetsNFloor field if non-nil, zero value otherwise.

### GetMeetsNFloorOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetMeetsNFloorOk() (*bool, bool)`

GetMeetsNFloorOk returns a tuple with the MeetsNFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetsNFloor

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetMeetsNFloor(v bool)`

SetMeetsNFloor sets MeetsNFloor field to given value.


### GetNFloor

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetNFloor() int32`

GetNFloor returns the NFloor field if non-nil, zero value otherwise.

### GetNFloorOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetNFloorOk() (*int32, bool)`

GetNFloorOk returns a tuple with the NFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFloor

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetNFloor(v int32)`

SetNFloor sets NFloor field to given value.


### GetMedianSlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetMedianSlipDays() float32`

GetMedianSlipDays returns the MedianSlipDays field if non-nil, zero value otherwise.

### GetMedianSlipDaysOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetMedianSlipDaysOk() (*float32, bool)`

GetMedianSlipDaysOk returns a tuple with the MedianSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianSlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetMedianSlipDays(v float32)`

SetMedianSlipDays sets MedianSlipDays field to given value.


### GetP25SlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetP25SlipDays() float32`

GetP25SlipDays returns the P25SlipDays field if non-nil, zero value otherwise.

### GetP25SlipDaysOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetP25SlipDaysOk() (*float32, bool)`

GetP25SlipDaysOk returns a tuple with the P25SlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP25SlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetP25SlipDays(v float32)`

SetP25SlipDays sets P25SlipDays field to given value.


### GetP75SlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetP75SlipDays() float32`

GetP75SlipDays returns the P75SlipDays field if non-nil, zero value otherwise.

### GetP75SlipDaysOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetP75SlipDaysOk() (*float32, bool)`

GetP75SlipDaysOk returns a tuple with the P75SlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP75SlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetP75SlipDays(v float32)`

SetP75SlipDays sets P75SlipDays field to given value.


### GetLateOver30d

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetLateOver30d() int32`

GetLateOver30d returns the LateOver30d field if non-nil, zero value otherwise.

### GetLateOver30dOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetLateOver30dOk() (*int32, bool)`

GetLateOver30dOk returns a tuple with the LateOver30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateOver30d

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetLateOver30d(v int32)`

SetLateOver30d sets LateOver30d field to given value.


### GetSuppressionNote

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetSuppressionNote() interface{}`

GetSuppressionNote returns the SuppressionNote field if non-nil, zero value otherwise.

### GetSuppressionNoteOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) GetSuppressionNoteOk() (*interface{}, bool)`

GetSuppressionNoteOk returns a tuple with the SuppressionNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionNote

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetSuppressionNote(v interface{})`

SetSuppressionNote sets SuppressionNote field to given value.


### SetSuppressionNoteNil

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) SetSuppressionNoteNil(b bool)`

 SetSuppressionNoteNil sets the value for SuppressionNote to be an explicit nil

### UnsetSuppressionNote
`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateVoluntary) UnsetSuppressionNote()`

UnsetSuppressionNote ensures that no value is present for SuppressionNote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



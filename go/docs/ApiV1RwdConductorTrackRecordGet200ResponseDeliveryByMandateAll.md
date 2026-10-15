# ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll

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
**BackfilledShare** | **NullableString** |  | 
**MeetsNFloor** | **bool** |  | 
**NFloor** | **int32** |  | 
**MedianSlipDays** | **NullableFloat32** |  | 
**P25SlipDays** | **NullableFloat32** |  | 
**P75SlipDays** | **NullableFloat32** |  | 
**LateOver30d** | **NullableInt32** |  | 
**SuppressionNote** | **NullableString** |  | 

## Methods

### NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll

`func NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll(label string, note string, lifetimeStudies int32, recentStudies int32, ongoingStudies int32, mandatedStudies int32, mandatedShare string, finalReportPairs int32, backfilledPairs int32, comparablePairs int32, backfilledShare NullableString, meetsNFloor bool, nFloor int32, medianSlipDays NullableFloat32, p25SlipDays NullableFloat32, p75SlipDays NullableFloat32, lateOver30d NullableInt32, suppressionNote NullableString, ) *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll`

NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll instantiates a new ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAllWithDefaults

`func NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAllWithDefaults() *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll`

NewApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAllWithDefaults instantiates a new ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNote

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetNote(v string)`

SetNote sets Note field to given value.


### GetLifetimeStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetLifetimeStudies() int32`

GetLifetimeStudies returns the LifetimeStudies field if non-nil, zero value otherwise.

### GetLifetimeStudiesOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetLifetimeStudiesOk() (*int32, bool)`

GetLifetimeStudiesOk returns a tuple with the LifetimeStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetLifetimeStudies(v int32)`

SetLifetimeStudies sets LifetimeStudies field to given value.


### GetRecentStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetRecentStudies() int32`

GetRecentStudies returns the RecentStudies field if non-nil, zero value otherwise.

### GetRecentStudiesOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetRecentStudiesOk() (*int32, bool)`

GetRecentStudiesOk returns a tuple with the RecentStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetRecentStudies(v int32)`

SetRecentStudies sets RecentStudies field to given value.


### GetOngoingStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetOngoingStudies() int32`

GetOngoingStudies returns the OngoingStudies field if non-nil, zero value otherwise.

### GetOngoingStudiesOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetOngoingStudiesOk() (*int32, bool)`

GetOngoingStudiesOk returns a tuple with the OngoingStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoingStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetOngoingStudies(v int32)`

SetOngoingStudies sets OngoingStudies field to given value.


### GetMandatedStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetMandatedStudies() int32`

GetMandatedStudies returns the MandatedStudies field if non-nil, zero value otherwise.

### GetMandatedStudiesOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetMandatedStudiesOk() (*int32, bool)`

GetMandatedStudiesOk returns a tuple with the MandatedStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatedStudies

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetMandatedStudies(v int32)`

SetMandatedStudies sets MandatedStudies field to given value.


### GetMandatedShare

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetMandatedShare() string`

GetMandatedShare returns the MandatedShare field if non-nil, zero value otherwise.

### GetMandatedShareOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetMandatedShareOk() (*string, bool)`

GetMandatedShareOk returns a tuple with the MandatedShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatedShare

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetMandatedShare(v string)`

SetMandatedShare sets MandatedShare field to given value.


### GetFinalReportPairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetFinalReportPairs() int32`

GetFinalReportPairs returns the FinalReportPairs field if non-nil, zero value otherwise.

### GetFinalReportPairsOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetFinalReportPairsOk() (*int32, bool)`

GetFinalReportPairsOk returns a tuple with the FinalReportPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportPairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetFinalReportPairs(v int32)`

SetFinalReportPairs sets FinalReportPairs field to given value.


### GetBackfilledPairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetBackfilledPairs() int32`

GetBackfilledPairs returns the BackfilledPairs field if non-nil, zero value otherwise.

### GetBackfilledPairsOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetBackfilledPairsOk() (*int32, bool)`

GetBackfilledPairsOk returns a tuple with the BackfilledPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilledPairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetBackfilledPairs(v int32)`

SetBackfilledPairs sets BackfilledPairs field to given value.


### GetComparablePairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetComparablePairs() int32`

GetComparablePairs returns the ComparablePairs field if non-nil, zero value otherwise.

### GetComparablePairsOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetComparablePairsOk() (*int32, bool)`

GetComparablePairsOk returns a tuple with the ComparablePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparablePairs

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetComparablePairs(v int32)`

SetComparablePairs sets ComparablePairs field to given value.


### GetBackfilledShare

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetBackfilledShare() string`

GetBackfilledShare returns the BackfilledShare field if non-nil, zero value otherwise.

### GetBackfilledShareOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetBackfilledShareOk() (*string, bool)`

GetBackfilledShareOk returns a tuple with the BackfilledShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilledShare

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetBackfilledShare(v string)`

SetBackfilledShare sets BackfilledShare field to given value.


### SetBackfilledShareNil

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetBackfilledShareNil(b bool)`

 SetBackfilledShareNil sets the value for BackfilledShare to be an explicit nil

### UnsetBackfilledShare
`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) UnsetBackfilledShare()`

UnsetBackfilledShare ensures that no value is present for BackfilledShare, not even an explicit nil
### GetMeetsNFloor

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetMeetsNFloor() bool`

GetMeetsNFloor returns the MeetsNFloor field if non-nil, zero value otherwise.

### GetMeetsNFloorOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetMeetsNFloorOk() (*bool, bool)`

GetMeetsNFloorOk returns a tuple with the MeetsNFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetsNFloor

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetMeetsNFloor(v bool)`

SetMeetsNFloor sets MeetsNFloor field to given value.


### GetNFloor

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetNFloor() int32`

GetNFloor returns the NFloor field if non-nil, zero value otherwise.

### GetNFloorOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetNFloorOk() (*int32, bool)`

GetNFloorOk returns a tuple with the NFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFloor

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetNFloor(v int32)`

SetNFloor sets NFloor field to given value.


### GetMedianSlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetMedianSlipDays() float32`

GetMedianSlipDays returns the MedianSlipDays field if non-nil, zero value otherwise.

### GetMedianSlipDaysOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetMedianSlipDaysOk() (*float32, bool)`

GetMedianSlipDaysOk returns a tuple with the MedianSlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianSlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetMedianSlipDays(v float32)`

SetMedianSlipDays sets MedianSlipDays field to given value.


### SetMedianSlipDaysNil

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetMedianSlipDaysNil(b bool)`

 SetMedianSlipDaysNil sets the value for MedianSlipDays to be an explicit nil

### UnsetMedianSlipDays
`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) UnsetMedianSlipDays()`

UnsetMedianSlipDays ensures that no value is present for MedianSlipDays, not even an explicit nil
### GetP25SlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetP25SlipDays() float32`

GetP25SlipDays returns the P25SlipDays field if non-nil, zero value otherwise.

### GetP25SlipDaysOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetP25SlipDaysOk() (*float32, bool)`

GetP25SlipDaysOk returns a tuple with the P25SlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP25SlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetP25SlipDays(v float32)`

SetP25SlipDays sets P25SlipDays field to given value.


### SetP25SlipDaysNil

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetP25SlipDaysNil(b bool)`

 SetP25SlipDaysNil sets the value for P25SlipDays to be an explicit nil

### UnsetP25SlipDays
`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) UnsetP25SlipDays()`

UnsetP25SlipDays ensures that no value is present for P25SlipDays, not even an explicit nil
### GetP75SlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetP75SlipDays() float32`

GetP75SlipDays returns the P75SlipDays field if non-nil, zero value otherwise.

### GetP75SlipDaysOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetP75SlipDaysOk() (*float32, bool)`

GetP75SlipDaysOk returns a tuple with the P75SlipDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP75SlipDays

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetP75SlipDays(v float32)`

SetP75SlipDays sets P75SlipDays field to given value.


### SetP75SlipDaysNil

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetP75SlipDaysNil(b bool)`

 SetP75SlipDaysNil sets the value for P75SlipDays to be an explicit nil

### UnsetP75SlipDays
`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) UnsetP75SlipDays()`

UnsetP75SlipDays ensures that no value is present for P75SlipDays, not even an explicit nil
### GetLateOver30d

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetLateOver30d() int32`

GetLateOver30d returns the LateOver30d field if non-nil, zero value otherwise.

### GetLateOver30dOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetLateOver30dOk() (*int32, bool)`

GetLateOver30dOk returns a tuple with the LateOver30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateOver30d

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetLateOver30d(v int32)`

SetLateOver30d sets LateOver30d field to given value.


### SetLateOver30dNil

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetLateOver30dNil(b bool)`

 SetLateOver30dNil sets the value for LateOver30d to be an explicit nil

### UnsetLateOver30d
`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) UnsetLateOver30d()`

UnsetLateOver30d ensures that no value is present for LateOver30d, not even an explicit nil
### GetSuppressionNote

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetSuppressionNote() string`

GetSuppressionNote returns the SuppressionNote field if non-nil, zero value otherwise.

### GetSuppressionNoteOk

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) GetSuppressionNoteOk() (*string, bool)`

GetSuppressionNoteOk returns a tuple with the SuppressionNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionNote

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetSuppressionNote(v string)`

SetSuppressionNote sets SuppressionNote field to given value.


### SetSuppressionNoteNil

`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) SetSuppressionNoteNil(b bool)`

 SetSuppressionNoteNil sets the value for SuppressionNote to be an explicit nil

### UnsetSuppressionNote
`func (o *ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandateAll) UnsetSuppressionNote()`

UnsetSuppressionNote ensures that no value is present for SuppressionNote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pairs** | **int32** |  | 
**Backfilled** | **int32** |  | 
**Comparable** | **int32** |  | 
**Median** | **NullableFloat32** |  | 
**P25** | **NullableFloat32** |  | 
**P75** | **NullableFloat32** |  | 
**Late30** | **int32** |  | 
**Floor** | **int32** |  | 
**Suppressed** | **bool** |  | 

## Methods

### NewApiV1RwdStudyScreenGet200ResponseDeliveryDistribution

`func NewApiV1RwdStudyScreenGet200ResponseDeliveryDistribution(pairs int32, backfilled int32, comparable int32, median NullableFloat32, p25 NullableFloat32, p75 NullableFloat32, late30 int32, floor int32, suppressed bool, ) *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution`

NewApiV1RwdStudyScreenGet200ResponseDeliveryDistribution instantiates a new ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RwdStudyScreenGet200ResponseDeliveryDistributionWithDefaults

`func NewApiV1RwdStudyScreenGet200ResponseDeliveryDistributionWithDefaults() *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution`

NewApiV1RwdStudyScreenGet200ResponseDeliveryDistributionWithDefaults instantiates a new ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPairs

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetPairs() int32`

GetPairs returns the Pairs field if non-nil, zero value otherwise.

### GetPairsOk

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetPairsOk() (*int32, bool)`

GetPairsOk returns a tuple with the Pairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairs

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetPairs(v int32)`

SetPairs sets Pairs field to given value.


### GetBackfilled

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetBackfilled() int32`

GetBackfilled returns the Backfilled field if non-nil, zero value otherwise.

### GetBackfilledOk

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetBackfilledOk() (*int32, bool)`

GetBackfilledOk returns a tuple with the Backfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilled

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetBackfilled(v int32)`

SetBackfilled sets Backfilled field to given value.


### GetComparable

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetComparable() int32`

GetComparable returns the Comparable field if non-nil, zero value otherwise.

### GetComparableOk

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetComparableOk() (*int32, bool)`

GetComparableOk returns a tuple with the Comparable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparable

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetComparable(v int32)`

SetComparable sets Comparable field to given value.


### GetMedian

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetMedian() float32`

GetMedian returns the Median field if non-nil, zero value otherwise.

### GetMedianOk

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetMedianOk() (*float32, bool)`

GetMedianOk returns a tuple with the Median field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedian

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetMedian(v float32)`

SetMedian sets Median field to given value.


### SetMedianNil

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetMedianNil(b bool)`

 SetMedianNil sets the value for Median to be an explicit nil

### UnsetMedian
`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) UnsetMedian()`

UnsetMedian ensures that no value is present for Median, not even an explicit nil
### GetP25

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetP25() float32`

GetP25 returns the P25 field if non-nil, zero value otherwise.

### GetP25Ok

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetP25Ok() (*float32, bool)`

GetP25Ok returns a tuple with the P25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP25

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetP25(v float32)`

SetP25 sets P25 field to given value.


### SetP25Nil

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetP25Nil(b bool)`

 SetP25Nil sets the value for P25 to be an explicit nil

### UnsetP25
`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) UnsetP25()`

UnsetP25 ensures that no value is present for P25, not even an explicit nil
### GetP75

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetP75() float32`

GetP75 returns the P75 field if non-nil, zero value otherwise.

### GetP75Ok

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetP75Ok() (*float32, bool)`

GetP75Ok returns a tuple with the P75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP75

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetP75(v float32)`

SetP75 sets P75 field to given value.


### SetP75Nil

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetP75Nil(b bool)`

 SetP75Nil sets the value for P75 to be an explicit nil

### UnsetP75
`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) UnsetP75()`

UnsetP75 ensures that no value is present for P75, not even an explicit nil
### GetLate30

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetLate30() int32`

GetLate30 returns the Late30 field if non-nil, zero value otherwise.

### GetLate30Ok

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetLate30Ok() (*int32, bool)`

GetLate30Ok returns a tuple with the Late30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLate30

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetLate30(v int32)`

SetLate30 sets Late30 field to given value.


### GetFloor

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetFloor() int32`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetFloorOk() (*int32, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetFloor(v int32)`

SetFloor sets Floor field to given value.


### GetSuppressed

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *ApiV1RwdStudyScreenGet200ResponseDeliveryDistribution) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



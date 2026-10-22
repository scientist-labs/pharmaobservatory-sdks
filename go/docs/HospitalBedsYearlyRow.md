# HospitalBedsYearlyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccn** | Pointer to **NullableString** | CMS Certification Number of the hospital. | [optional] 
**Year** | Pointer to **NullableInt32** | The calendar year. | [optional] 
**BedCount** | Pointer to **NullableInt32** | Certified bed count that year. | [optional] 
**IsTeaching** | Pointer to **NullableBool** | Whether the hospital was a teaching hospital that year. | [optional] 
**PosVintage** | Pointer to **NullableString** | The Provider of Services file vintage the row comes from. | [optional] 
**AsOf** | Pointer to **NullableString** | The as-of date of that file. | [optional] 

## Methods

### NewHospitalBedsYearlyRow

`func NewHospitalBedsYearlyRow() *HospitalBedsYearlyRow`

NewHospitalBedsYearlyRow instantiates a new HospitalBedsYearlyRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHospitalBedsYearlyRowWithDefaults

`func NewHospitalBedsYearlyRowWithDefaults() *HospitalBedsYearlyRow`

NewHospitalBedsYearlyRowWithDefaults instantiates a new HospitalBedsYearlyRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcn

`func (o *HospitalBedsYearlyRow) GetCcn() string`

GetCcn returns the Ccn field if non-nil, zero value otherwise.

### GetCcnOk

`func (o *HospitalBedsYearlyRow) GetCcnOk() (*string, bool)`

GetCcnOk returns a tuple with the Ccn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcn

`func (o *HospitalBedsYearlyRow) SetCcn(v string)`

SetCcn sets Ccn field to given value.

### HasCcn

`func (o *HospitalBedsYearlyRow) HasCcn() bool`

HasCcn returns a boolean if a field has been set.

### SetCcnNil

`func (o *HospitalBedsYearlyRow) SetCcnNil(b bool)`

 SetCcnNil sets the value for Ccn to be an explicit nil

### UnsetCcn
`func (o *HospitalBedsYearlyRow) UnsetCcn()`

UnsetCcn ensures that no value is present for Ccn, not even an explicit nil
### GetYear

`func (o *HospitalBedsYearlyRow) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *HospitalBedsYearlyRow) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *HospitalBedsYearlyRow) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *HospitalBedsYearlyRow) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *HospitalBedsYearlyRow) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *HospitalBedsYearlyRow) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetBedCount

`func (o *HospitalBedsYearlyRow) GetBedCount() int32`

GetBedCount returns the BedCount field if non-nil, zero value otherwise.

### GetBedCountOk

`func (o *HospitalBedsYearlyRow) GetBedCountOk() (*int32, bool)`

GetBedCountOk returns a tuple with the BedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBedCount

`func (o *HospitalBedsYearlyRow) SetBedCount(v int32)`

SetBedCount sets BedCount field to given value.

### HasBedCount

`func (o *HospitalBedsYearlyRow) HasBedCount() bool`

HasBedCount returns a boolean if a field has been set.

### SetBedCountNil

`func (o *HospitalBedsYearlyRow) SetBedCountNil(b bool)`

 SetBedCountNil sets the value for BedCount to be an explicit nil

### UnsetBedCount
`func (o *HospitalBedsYearlyRow) UnsetBedCount()`

UnsetBedCount ensures that no value is present for BedCount, not even an explicit nil
### GetIsTeaching

`func (o *HospitalBedsYearlyRow) GetIsTeaching() bool`

GetIsTeaching returns the IsTeaching field if non-nil, zero value otherwise.

### GetIsTeachingOk

`func (o *HospitalBedsYearlyRow) GetIsTeachingOk() (*bool, bool)`

GetIsTeachingOk returns a tuple with the IsTeaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTeaching

`func (o *HospitalBedsYearlyRow) SetIsTeaching(v bool)`

SetIsTeaching sets IsTeaching field to given value.

### HasIsTeaching

`func (o *HospitalBedsYearlyRow) HasIsTeaching() bool`

HasIsTeaching returns a boolean if a field has been set.

### SetIsTeachingNil

`func (o *HospitalBedsYearlyRow) SetIsTeachingNil(b bool)`

 SetIsTeachingNil sets the value for IsTeaching to be an explicit nil

### UnsetIsTeaching
`func (o *HospitalBedsYearlyRow) UnsetIsTeaching()`

UnsetIsTeaching ensures that no value is present for IsTeaching, not even an explicit nil
### GetPosVintage

`func (o *HospitalBedsYearlyRow) GetPosVintage() string`

GetPosVintage returns the PosVintage field if non-nil, zero value otherwise.

### GetPosVintageOk

`func (o *HospitalBedsYearlyRow) GetPosVintageOk() (*string, bool)`

GetPosVintageOk returns a tuple with the PosVintage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosVintage

`func (o *HospitalBedsYearlyRow) SetPosVintage(v string)`

SetPosVintage sets PosVintage field to given value.

### HasPosVintage

`func (o *HospitalBedsYearlyRow) HasPosVintage() bool`

HasPosVintage returns a boolean if a field has been set.

### SetPosVintageNil

`func (o *HospitalBedsYearlyRow) SetPosVintageNil(b bool)`

 SetPosVintageNil sets the value for PosVintage to be an explicit nil

### UnsetPosVintage
`func (o *HospitalBedsYearlyRow) UnsetPosVintage()`

UnsetPosVintage ensures that no value is present for PosVintage, not even an explicit nil
### GetAsOf

`func (o *HospitalBedsYearlyRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *HospitalBedsYearlyRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *HospitalBedsYearlyRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *HospitalBedsYearlyRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *HospitalBedsYearlyRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *HospitalBedsYearlyRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



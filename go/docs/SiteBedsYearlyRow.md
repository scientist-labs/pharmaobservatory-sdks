# SiteBedsYearlyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **NullableString** | The site key: one facility spelling from the registry. | [optional] 
**Year** | Pointer to **NullableInt32** | The registry year. | [optional] 
**Ccn** | Pointer to **NullableString** | CMS Certification Number of the matched hospital. | [optional] 
**BedCount** | Pointer to **NullableInt32** | Certified bed count that year. | [optional] 
**IsTeaching** | Pointer to **NullableBool** | Whether the hospital was a teaching hospital that year. | [optional] 
**PosVintage** | Pointer to **NullableString** | The Provider of Services file vintage the row comes from. | [optional] 
**AsOf** | Pointer to **NullableString** | The as-of date of that file. | [optional] 

## Methods

### NewSiteBedsYearlyRow

`func NewSiteBedsYearlyRow() *SiteBedsYearlyRow`

NewSiteBedsYearlyRow instantiates a new SiteBedsYearlyRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteBedsYearlyRowWithDefaults

`func NewSiteBedsYearlyRowWithDefaults() *SiteBedsYearlyRow`

NewSiteBedsYearlyRowWithDefaults instantiates a new SiteBedsYearlyRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *SiteBedsYearlyRow) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteBedsYearlyRow) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteBedsYearlyRow) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteBedsYearlyRow) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *SiteBedsYearlyRow) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *SiteBedsYearlyRow) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetYear

`func (o *SiteBedsYearlyRow) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SiteBedsYearlyRow) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SiteBedsYearlyRow) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *SiteBedsYearlyRow) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *SiteBedsYearlyRow) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *SiteBedsYearlyRow) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetCcn

`func (o *SiteBedsYearlyRow) GetCcn() string`

GetCcn returns the Ccn field if non-nil, zero value otherwise.

### GetCcnOk

`func (o *SiteBedsYearlyRow) GetCcnOk() (*string, bool)`

GetCcnOk returns a tuple with the Ccn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcn

`func (o *SiteBedsYearlyRow) SetCcn(v string)`

SetCcn sets Ccn field to given value.

### HasCcn

`func (o *SiteBedsYearlyRow) HasCcn() bool`

HasCcn returns a boolean if a field has been set.

### SetCcnNil

`func (o *SiteBedsYearlyRow) SetCcnNil(b bool)`

 SetCcnNil sets the value for Ccn to be an explicit nil

### UnsetCcn
`func (o *SiteBedsYearlyRow) UnsetCcn()`

UnsetCcn ensures that no value is present for Ccn, not even an explicit nil
### GetBedCount

`func (o *SiteBedsYearlyRow) GetBedCount() int32`

GetBedCount returns the BedCount field if non-nil, zero value otherwise.

### GetBedCountOk

`func (o *SiteBedsYearlyRow) GetBedCountOk() (*int32, bool)`

GetBedCountOk returns a tuple with the BedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBedCount

`func (o *SiteBedsYearlyRow) SetBedCount(v int32)`

SetBedCount sets BedCount field to given value.

### HasBedCount

`func (o *SiteBedsYearlyRow) HasBedCount() bool`

HasBedCount returns a boolean if a field has been set.

### SetBedCountNil

`func (o *SiteBedsYearlyRow) SetBedCountNil(b bool)`

 SetBedCountNil sets the value for BedCount to be an explicit nil

### UnsetBedCount
`func (o *SiteBedsYearlyRow) UnsetBedCount()`

UnsetBedCount ensures that no value is present for BedCount, not even an explicit nil
### GetIsTeaching

`func (o *SiteBedsYearlyRow) GetIsTeaching() bool`

GetIsTeaching returns the IsTeaching field if non-nil, zero value otherwise.

### GetIsTeachingOk

`func (o *SiteBedsYearlyRow) GetIsTeachingOk() (*bool, bool)`

GetIsTeachingOk returns a tuple with the IsTeaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTeaching

`func (o *SiteBedsYearlyRow) SetIsTeaching(v bool)`

SetIsTeaching sets IsTeaching field to given value.

### HasIsTeaching

`func (o *SiteBedsYearlyRow) HasIsTeaching() bool`

HasIsTeaching returns a boolean if a field has been set.

### SetIsTeachingNil

`func (o *SiteBedsYearlyRow) SetIsTeachingNil(b bool)`

 SetIsTeachingNil sets the value for IsTeaching to be an explicit nil

### UnsetIsTeaching
`func (o *SiteBedsYearlyRow) UnsetIsTeaching()`

UnsetIsTeaching ensures that no value is present for IsTeaching, not even an explicit nil
### GetPosVintage

`func (o *SiteBedsYearlyRow) GetPosVintage() string`

GetPosVintage returns the PosVintage field if non-nil, zero value otherwise.

### GetPosVintageOk

`func (o *SiteBedsYearlyRow) GetPosVintageOk() (*string, bool)`

GetPosVintageOk returns a tuple with the PosVintage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosVintage

`func (o *SiteBedsYearlyRow) SetPosVintage(v string)`

SetPosVintage sets PosVintage field to given value.

### HasPosVintage

`func (o *SiteBedsYearlyRow) HasPosVintage() bool`

HasPosVintage returns a boolean if a field has been set.

### SetPosVintageNil

`func (o *SiteBedsYearlyRow) SetPosVintageNil(b bool)`

 SetPosVintageNil sets the value for PosVintage to be an explicit nil

### UnsetPosVintage
`func (o *SiteBedsYearlyRow) UnsetPosVintage()`

UnsetPosVintage ensures that no value is present for PosVintage, not even an explicit nil
### GetAsOf

`func (o *SiteBedsYearlyRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *SiteBedsYearlyRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *SiteBedsYearlyRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *SiteBedsYearlyRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *SiteBedsYearlyRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *SiteBedsYearlyRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



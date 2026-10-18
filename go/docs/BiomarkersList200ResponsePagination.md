# BiomarkersList200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Total** | **int32** |  | 
**TotalBasis** | **string** |  | 
**Page** | **int32** |  | 
**NextCursor** | **NullableString** |  | 
**PrevCursor** | **NullableString** |  | 

## Methods

### NewBiomarkersList200ResponsePagination

`func NewBiomarkersList200ResponsePagination(limit int32, total int32, totalBasis string, page int32, nextCursor NullableString, prevCursor NullableString, ) *BiomarkersList200ResponsePagination`

NewBiomarkersList200ResponsePagination instantiates a new BiomarkersList200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkersList200ResponsePaginationWithDefaults

`func NewBiomarkersList200ResponsePaginationWithDefaults() *BiomarkersList200ResponsePagination`

NewBiomarkersList200ResponsePaginationWithDefaults instantiates a new BiomarkersList200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *BiomarkersList200ResponsePagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BiomarkersList200ResponsePagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BiomarkersList200ResponsePagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetTotal

`func (o *BiomarkersList200ResponsePagination) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BiomarkersList200ResponsePagination) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BiomarkersList200ResponsePagination) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetTotalBasis

`func (o *BiomarkersList200ResponsePagination) GetTotalBasis() string`

GetTotalBasis returns the TotalBasis field if non-nil, zero value otherwise.

### GetTotalBasisOk

`func (o *BiomarkersList200ResponsePagination) GetTotalBasisOk() (*string, bool)`

GetTotalBasisOk returns a tuple with the TotalBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBasis

`func (o *BiomarkersList200ResponsePagination) SetTotalBasis(v string)`

SetTotalBasis sets TotalBasis field to given value.


### GetPage

`func (o *BiomarkersList200ResponsePagination) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BiomarkersList200ResponsePagination) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BiomarkersList200ResponsePagination) SetPage(v int32)`

SetPage sets Page field to given value.


### GetNextCursor

`func (o *BiomarkersList200ResponsePagination) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *BiomarkersList200ResponsePagination) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *BiomarkersList200ResponsePagination) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *BiomarkersList200ResponsePagination) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *BiomarkersList200ResponsePagination) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPrevCursor

`func (o *BiomarkersList200ResponsePagination) GetPrevCursor() string`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *BiomarkersList200ResponsePagination) GetPrevCursorOk() (*string, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *BiomarkersList200ResponsePagination) SetPrevCursor(v string)`

SetPrevCursor sets PrevCursor field to given value.


### SetPrevCursorNil

`func (o *BiomarkersList200ResponsePagination) SetPrevCursorNil(b bool)`

 SetPrevCursorNil sets the value for PrevCursor to be an explicit nil

### UnsetPrevCursor
`func (o *BiomarkersList200ResponsePagination) UnsetPrevCursor()`

UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



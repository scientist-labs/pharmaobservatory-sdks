# OrgsList200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Total** | **int32** |  | 
**TotalBasis** | **string** |  | 
**Page** | **int32** |  | 
**NextCursor** | **string** |  | 
**PrevCursor** | **interface{}** |  | 

## Methods

### NewOrgsList200ResponsePagination

`func NewOrgsList200ResponsePagination(limit int32, total int32, totalBasis string, page int32, nextCursor string, prevCursor interface{}, ) *OrgsList200ResponsePagination`

NewOrgsList200ResponsePagination instantiates a new OrgsList200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsList200ResponsePaginationWithDefaults

`func NewOrgsList200ResponsePaginationWithDefaults() *OrgsList200ResponsePagination`

NewOrgsList200ResponsePaginationWithDefaults instantiates a new OrgsList200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *OrgsList200ResponsePagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *OrgsList200ResponsePagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *OrgsList200ResponsePagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetTotal

`func (o *OrgsList200ResponsePagination) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrgsList200ResponsePagination) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrgsList200ResponsePagination) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetTotalBasis

`func (o *OrgsList200ResponsePagination) GetTotalBasis() string`

GetTotalBasis returns the TotalBasis field if non-nil, zero value otherwise.

### GetTotalBasisOk

`func (o *OrgsList200ResponsePagination) GetTotalBasisOk() (*string, bool)`

GetTotalBasisOk returns a tuple with the TotalBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBasis

`func (o *OrgsList200ResponsePagination) SetTotalBasis(v string)`

SetTotalBasis sets TotalBasis field to given value.


### GetPage

`func (o *OrgsList200ResponsePagination) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgsList200ResponsePagination) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgsList200ResponsePagination) SetPage(v int32)`

SetPage sets Page field to given value.


### GetNextCursor

`func (o *OrgsList200ResponsePagination) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *OrgsList200ResponsePagination) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *OrgsList200ResponsePagination) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### GetPrevCursor

`func (o *OrgsList200ResponsePagination) GetPrevCursor() interface{}`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *OrgsList200ResponsePagination) GetPrevCursorOk() (*interface{}, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *OrgsList200ResponsePagination) SetPrevCursor(v interface{})`

SetPrevCursor sets PrevCursor field to given value.


### SetPrevCursorNil

`func (o *OrgsList200ResponsePagination) SetPrevCursorNil(b bool)`

 SetPrevCursorNil sets the value for PrevCursor to be an explicit nil

### UnsetPrevCursor
`func (o *OrgsList200ResponsePagination) UnsetPrevCursor()`

UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



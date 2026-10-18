# OrgsOrgOwnership200ResponseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**View** | **string** |  | 
**Limit** | **int32** |  | 
**Totals** | [**OrgsCurrentOwnership200ResponseContextTotals**](OrgsCurrentOwnership200ResponseContextTotals.md) |  | 
**Truncated** | [**OrgsOrgOwnership200ResponseContextTruncated**](OrgsOrgOwnership200ResponseContextTruncated.md) |  | 
**PersonalDataOwners** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrgsOrgOwnership200ResponseContext

`func NewOrgsOrgOwnership200ResponseContext(view string, limit int32, totals OrgsCurrentOwnership200ResponseContextTotals, truncated OrgsOrgOwnership200ResponseContextTruncated, ) *OrgsOrgOwnership200ResponseContext`

NewOrgsOrgOwnership200ResponseContext instantiates a new OrgsOrgOwnership200ResponseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsOrgOwnership200ResponseContextWithDefaults

`func NewOrgsOrgOwnership200ResponseContextWithDefaults() *OrgsOrgOwnership200ResponseContext`

NewOrgsOrgOwnership200ResponseContextWithDefaults instantiates a new OrgsOrgOwnership200ResponseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetView

`func (o *OrgsOrgOwnership200ResponseContext) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *OrgsOrgOwnership200ResponseContext) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *OrgsOrgOwnership200ResponseContext) SetView(v string)`

SetView sets View field to given value.


### GetLimit

`func (o *OrgsOrgOwnership200ResponseContext) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *OrgsOrgOwnership200ResponseContext) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *OrgsOrgOwnership200ResponseContext) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetTotals

`func (o *OrgsOrgOwnership200ResponseContext) GetTotals() OrgsCurrentOwnership200ResponseContextTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *OrgsOrgOwnership200ResponseContext) GetTotalsOk() (*OrgsCurrentOwnership200ResponseContextTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *OrgsOrgOwnership200ResponseContext) SetTotals(v OrgsCurrentOwnership200ResponseContextTotals)`

SetTotals sets Totals field to given value.


### GetTruncated

`func (o *OrgsOrgOwnership200ResponseContext) GetTruncated() OrgsOrgOwnership200ResponseContextTruncated`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *OrgsOrgOwnership200ResponseContext) GetTruncatedOk() (*OrgsOrgOwnership200ResponseContextTruncated, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *OrgsOrgOwnership200ResponseContext) SetTruncated(v OrgsOrgOwnership200ResponseContextTruncated)`

SetTruncated sets Truncated field to given value.


### GetPersonalDataOwners

`func (o *OrgsOrgOwnership200ResponseContext) GetPersonalDataOwners() bool`

GetPersonalDataOwners returns the PersonalDataOwners field if non-nil, zero value otherwise.

### GetPersonalDataOwnersOk

`func (o *OrgsOrgOwnership200ResponseContext) GetPersonalDataOwnersOk() (*bool, bool)`

GetPersonalDataOwnersOk returns a tuple with the PersonalDataOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDataOwners

`func (o *OrgsOrgOwnership200ResponseContext) SetPersonalDataOwners(v bool)`

SetPersonalDataOwners sets PersonalDataOwners field to given value.

### HasPersonalDataOwners

`func (o *OrgsOrgOwnership200ResponseContext) HasPersonalDataOwners() bool`

HasPersonalDataOwners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OrgsCurrentOwnership200ResponseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**View** | **string** |  | 
**Totals** | [**OrgsCurrentOwnership200ResponseContextTotals**](OrgsCurrentOwnership200ResponseContextTotals.md) |  | 
**PersonalDataOwners** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrgsCurrentOwnership200ResponseContext

`func NewOrgsCurrentOwnership200ResponseContext(view string, totals OrgsCurrentOwnership200ResponseContextTotals, ) *OrgsCurrentOwnership200ResponseContext`

NewOrgsCurrentOwnership200ResponseContext instantiates a new OrgsCurrentOwnership200ResponseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsCurrentOwnership200ResponseContextWithDefaults

`func NewOrgsCurrentOwnership200ResponseContextWithDefaults() *OrgsCurrentOwnership200ResponseContext`

NewOrgsCurrentOwnership200ResponseContextWithDefaults instantiates a new OrgsCurrentOwnership200ResponseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetView

`func (o *OrgsCurrentOwnership200ResponseContext) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *OrgsCurrentOwnership200ResponseContext) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *OrgsCurrentOwnership200ResponseContext) SetView(v string)`

SetView sets View field to given value.


### GetTotals

`func (o *OrgsCurrentOwnership200ResponseContext) GetTotals() OrgsCurrentOwnership200ResponseContextTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *OrgsCurrentOwnership200ResponseContext) GetTotalsOk() (*OrgsCurrentOwnership200ResponseContextTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *OrgsCurrentOwnership200ResponseContext) SetTotals(v OrgsCurrentOwnership200ResponseContextTotals)`

SetTotals sets Totals field to given value.


### GetPersonalDataOwners

`func (o *OrgsCurrentOwnership200ResponseContext) GetPersonalDataOwners() bool`

GetPersonalDataOwners returns the PersonalDataOwners field if non-nil, zero value otherwise.

### GetPersonalDataOwnersOk

`func (o *OrgsCurrentOwnership200ResponseContext) GetPersonalDataOwnersOk() (*bool, bool)`

GetPersonalDataOwnersOk returns a tuple with the PersonalDataOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDataOwners

`func (o *OrgsCurrentOwnership200ResponseContext) SetPersonalDataOwners(v bool)`

SetPersonalDataOwners sets PersonalDataOwners field to given value.

### HasPersonalDataOwners

`func (o *OrgsCurrentOwnership200ResponseContext) HasPersonalDataOwners() bool`

HasPersonalDataOwners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



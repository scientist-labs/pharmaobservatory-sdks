# OrgsList200ResponseQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Q** | **interface{}** |  | 
**Grain** | **string** |  | 
**Filters** | **map[string]interface{}** |  | 
**Sort** | **string** |  | 
**Dir** | **string** |  | 

## Methods

### NewOrgsList200ResponseQuery

`func NewOrgsList200ResponseQuery(q interface{}, grain string, filters map[string]interface{}, sort string, dir string, ) *OrgsList200ResponseQuery`

NewOrgsList200ResponseQuery instantiates a new OrgsList200ResponseQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsList200ResponseQueryWithDefaults

`func NewOrgsList200ResponseQueryWithDefaults() *OrgsList200ResponseQuery`

NewOrgsList200ResponseQueryWithDefaults instantiates a new OrgsList200ResponseQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQ

`func (o *OrgsList200ResponseQuery) GetQ() interface{}`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *OrgsList200ResponseQuery) GetQOk() (*interface{}, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *OrgsList200ResponseQuery) SetQ(v interface{})`

SetQ sets Q field to given value.


### SetQNil

`func (o *OrgsList200ResponseQuery) SetQNil(b bool)`

 SetQNil sets the value for Q to be an explicit nil

### UnsetQ
`func (o *OrgsList200ResponseQuery) UnsetQ()`

UnsetQ ensures that no value is present for Q, not even an explicit nil
### GetGrain

`func (o *OrgsList200ResponseQuery) GetGrain() string`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *OrgsList200ResponseQuery) GetGrainOk() (*string, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *OrgsList200ResponseQuery) SetGrain(v string)`

SetGrain sets Grain field to given value.


### GetFilters

`func (o *OrgsList200ResponseQuery) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OrgsList200ResponseQuery) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OrgsList200ResponseQuery) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.


### GetSort

`func (o *OrgsList200ResponseQuery) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *OrgsList200ResponseQuery) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *OrgsList200ResponseQuery) SetSort(v string)`

SetSort sets Sort field to given value.


### GetDir

`func (o *OrgsList200ResponseQuery) GetDir() string`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *OrgsList200ResponseQuery) GetDirOk() (*string, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *OrgsList200ResponseQuery) SetDir(v string)`

SetDir sets Dir field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



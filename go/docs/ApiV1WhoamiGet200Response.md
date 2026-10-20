# ApiV1WhoamiGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | 
**Sub** | **string** |  | 
**Roles** | **[]string** |  | 
**Name** | Pointer to **interface{}** |  | [optional] 
**Email** | Pointer to **interface{}** |  | [optional] 
**Organization** | Pointer to **interface{}** |  | [optional] 
**Plan** | Pointer to **interface{}** |  | [optional] 
**Limits** | Pointer to [**ApiV1WhoamiGet200ResponseLimits**](ApiV1WhoamiGet200ResponseLimits.md) |  | [optional] 
**Quota** | Pointer to [**ApiV1WhoamiGet200ResponseQuota**](ApiV1WhoamiGet200ResponseQuota.md) |  | [optional] 

## Methods

### NewApiV1WhoamiGet200Response

`func NewApiV1WhoamiGet200Response(ok bool, sub string, roles []string, ) *ApiV1WhoamiGet200Response`

NewApiV1WhoamiGet200Response instantiates a new ApiV1WhoamiGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1WhoamiGet200ResponseWithDefaults

`func NewApiV1WhoamiGet200ResponseWithDefaults() *ApiV1WhoamiGet200Response`

NewApiV1WhoamiGet200ResponseWithDefaults instantiates a new ApiV1WhoamiGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *ApiV1WhoamiGet200Response) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *ApiV1WhoamiGet200Response) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *ApiV1WhoamiGet200Response) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetSub

`func (o *ApiV1WhoamiGet200Response) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *ApiV1WhoamiGet200Response) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *ApiV1WhoamiGet200Response) SetSub(v string)`

SetSub sets Sub field to given value.


### GetRoles

`func (o *ApiV1WhoamiGet200Response) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiV1WhoamiGet200Response) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiV1WhoamiGet200Response) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetName

`func (o *ApiV1WhoamiGet200Response) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1WhoamiGet200Response) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1WhoamiGet200Response) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *ApiV1WhoamiGet200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV1WhoamiGet200Response) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV1WhoamiGet200Response) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *ApiV1WhoamiGet200Response) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiV1WhoamiGet200Response) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiV1WhoamiGet200Response) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApiV1WhoamiGet200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ApiV1WhoamiGet200Response) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ApiV1WhoamiGet200Response) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOrganization

`func (o *ApiV1WhoamiGet200Response) GetOrganization() interface{}`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ApiV1WhoamiGet200Response) GetOrganizationOk() (*interface{}, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ApiV1WhoamiGet200Response) SetOrganization(v interface{})`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ApiV1WhoamiGet200Response) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ApiV1WhoamiGet200Response) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ApiV1WhoamiGet200Response) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetPlan

`func (o *ApiV1WhoamiGet200Response) GetPlan() interface{}`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ApiV1WhoamiGet200Response) GetPlanOk() (*interface{}, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ApiV1WhoamiGet200Response) SetPlan(v interface{})`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ApiV1WhoamiGet200Response) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### SetPlanNil

`func (o *ApiV1WhoamiGet200Response) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *ApiV1WhoamiGet200Response) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetLimits

`func (o *ApiV1WhoamiGet200Response) GetLimits() ApiV1WhoamiGet200ResponseLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ApiV1WhoamiGet200Response) GetLimitsOk() (*ApiV1WhoamiGet200ResponseLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ApiV1WhoamiGet200Response) SetLimits(v ApiV1WhoamiGet200ResponseLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *ApiV1WhoamiGet200Response) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetQuota

`func (o *ApiV1WhoamiGet200Response) GetQuota() ApiV1WhoamiGet200ResponseQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ApiV1WhoamiGet200Response) GetQuotaOk() (*ApiV1WhoamiGet200ResponseQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ApiV1WhoamiGet200Response) SetQuota(v ApiV1WhoamiGet200ResponseQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *ApiV1WhoamiGet200Response) HasQuota() bool`

HasQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



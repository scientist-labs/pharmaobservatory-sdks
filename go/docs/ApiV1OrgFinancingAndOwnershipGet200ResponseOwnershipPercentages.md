# ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **NullableString** |  | 
**RequestedOrgId** | **NullableString** |  | 
**Any** | **bool** |  | 
**SharesOutstanding** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesSharesOutstanding**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesSharesOutstanding.md) |  | 
**Insider** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInsider**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInsider.md) |  | 
**Institutional** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInstitutional**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInstitutional.md) |  | 
**Notes** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesNotes**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesNotes.md) |  | 

## Methods

### NewApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages

`func NewApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages(orgId NullableString, requestedOrgId NullableString, any bool, sharesOutstanding ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesSharesOutstanding, insider ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInsider, institutional ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInstitutional, notes ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesNotes, ) *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages`

NewApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages instantiates a new ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesWithDefaults

`func NewApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesWithDefaults() *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages`

NewApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesWithDefaults instantiates a new ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### SetOrgIdNil

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetRequestedOrgId

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetRequestedOrgId() string`

GetRequestedOrgId returns the RequestedOrgId field if non-nil, zero value otherwise.

### GetRequestedOrgIdOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetRequestedOrgIdOk() (*string, bool)`

GetRequestedOrgIdOk returns a tuple with the RequestedOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedOrgId

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) SetRequestedOrgId(v string)`

SetRequestedOrgId sets RequestedOrgId field to given value.


### SetRequestedOrgIdNil

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) SetRequestedOrgIdNil(b bool)`

 SetRequestedOrgIdNil sets the value for RequestedOrgId to be an explicit nil

### UnsetRequestedOrgId
`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) UnsetRequestedOrgId()`

UnsetRequestedOrgId ensures that no value is present for RequestedOrgId, not even an explicit nil
### GetAny

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetAny() bool`

GetAny returns the Any field if non-nil, zero value otherwise.

### GetAnyOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetAnyOk() (*bool, bool)`

GetAnyOk returns a tuple with the Any field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAny

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) SetAny(v bool)`

SetAny sets Any field to given value.


### GetSharesOutstanding

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetSharesOutstanding() ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesSharesOutstanding`

GetSharesOutstanding returns the SharesOutstanding field if non-nil, zero value otherwise.

### GetSharesOutstandingOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetSharesOutstandingOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesSharesOutstanding, bool)`

GetSharesOutstandingOk returns a tuple with the SharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstanding

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) SetSharesOutstanding(v ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesSharesOutstanding)`

SetSharesOutstanding sets SharesOutstanding field to given value.


### GetInsider

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetInsider() ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInsider`

GetInsider returns the Insider field if non-nil, zero value otherwise.

### GetInsiderOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetInsiderOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInsider, bool)`

GetInsiderOk returns a tuple with the Insider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsider

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) SetInsider(v ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInsider)`

SetInsider sets Insider field to given value.


### GetInstitutional

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetInstitutional() ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInstitutional`

GetInstitutional returns the Institutional field if non-nil, zero value otherwise.

### GetInstitutionalOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetInstitutionalOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInstitutional, bool)`

GetInstitutionalOk returns a tuple with the Institutional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutional

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) SetInstitutional(v ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInstitutional)`

SetInstitutional sets Institutional field to given value.


### GetNotes

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetNotes() ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) GetNotesOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages) SetNotes(v ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesNotes)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



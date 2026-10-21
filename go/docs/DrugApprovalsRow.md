# DrugApprovalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**Authority** | Pointer to **NullableString** | The approving authority. | [optional] 
**Jurisdiction** | Pointer to **NullableString** | The jurisdiction the authority covers. | [optional] 
**Status** | Pointer to **NullableString** | The approval status the source states. | [optional] 
**FirstApprovalDate** | Pointer to **NullableString** | The earliest authorisation or first-marketed date the source gives. | [optional] 
**ApprovalId** | Pointer to **NullableString** | The product number the authority uses (EMA product number or Health Canada drug identification number). | [optional] 

## Methods

### NewDrugApprovalsRow

`func NewDrugApprovalsRow() *DrugApprovalsRow`

NewDrugApprovalsRow instantiates a new DrugApprovalsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugApprovalsRowWithDefaults

`func NewDrugApprovalsRowWithDefaults() *DrugApprovalsRow`

NewDrugApprovalsRowWithDefaults instantiates a new DrugApprovalsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugApprovalsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugApprovalsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugApprovalsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugApprovalsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugApprovalsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugApprovalsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetAuthority

`func (o *DrugApprovalsRow) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *DrugApprovalsRow) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *DrugApprovalsRow) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *DrugApprovalsRow) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### SetAuthorityNil

`func (o *DrugApprovalsRow) SetAuthorityNil(b bool)`

 SetAuthorityNil sets the value for Authority to be an explicit nil

### UnsetAuthority
`func (o *DrugApprovalsRow) UnsetAuthority()`

UnsetAuthority ensures that no value is present for Authority, not even an explicit nil
### GetJurisdiction

`func (o *DrugApprovalsRow) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *DrugApprovalsRow) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *DrugApprovalsRow) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *DrugApprovalsRow) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### SetJurisdictionNil

`func (o *DrugApprovalsRow) SetJurisdictionNil(b bool)`

 SetJurisdictionNil sets the value for Jurisdiction to be an explicit nil

### UnsetJurisdiction
`func (o *DrugApprovalsRow) UnsetJurisdiction()`

UnsetJurisdiction ensures that no value is present for Jurisdiction, not even an explicit nil
### GetStatus

`func (o *DrugApprovalsRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DrugApprovalsRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DrugApprovalsRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DrugApprovalsRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DrugApprovalsRow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DrugApprovalsRow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFirstApprovalDate

`func (o *DrugApprovalsRow) GetFirstApprovalDate() string`

GetFirstApprovalDate returns the FirstApprovalDate field if non-nil, zero value otherwise.

### GetFirstApprovalDateOk

`func (o *DrugApprovalsRow) GetFirstApprovalDateOk() (*string, bool)`

GetFirstApprovalDateOk returns a tuple with the FirstApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstApprovalDate

`func (o *DrugApprovalsRow) SetFirstApprovalDate(v string)`

SetFirstApprovalDate sets FirstApprovalDate field to given value.

### HasFirstApprovalDate

`func (o *DrugApprovalsRow) HasFirstApprovalDate() bool`

HasFirstApprovalDate returns a boolean if a field has been set.

### SetFirstApprovalDateNil

`func (o *DrugApprovalsRow) SetFirstApprovalDateNil(b bool)`

 SetFirstApprovalDateNil sets the value for FirstApprovalDate to be an explicit nil

### UnsetFirstApprovalDate
`func (o *DrugApprovalsRow) UnsetFirstApprovalDate()`

UnsetFirstApprovalDate ensures that no value is present for FirstApprovalDate, not even an explicit nil
### GetApprovalId

`func (o *DrugApprovalsRow) GetApprovalId() string`

GetApprovalId returns the ApprovalId field if non-nil, zero value otherwise.

### GetApprovalIdOk

`func (o *DrugApprovalsRow) GetApprovalIdOk() (*string, bool)`

GetApprovalIdOk returns a tuple with the ApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalId

`func (o *DrugApprovalsRow) SetApprovalId(v string)`

SetApprovalId sets ApprovalId field to given value.

### HasApprovalId

`func (o *DrugApprovalsRow) HasApprovalId() bool`

HasApprovalId returns a boolean if a field has been set.

### SetApprovalIdNil

`func (o *DrugApprovalsRow) SetApprovalIdNil(b bool)`

 SetApprovalIdNil sets the value for ApprovalId to be an explicit nil

### UnsetApprovalId
`func (o *DrugApprovalsRow) UnsetApprovalId()`

UnsetApprovalId ensures that no value is present for ApprovalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



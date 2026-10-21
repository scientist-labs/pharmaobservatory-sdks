# DrugAspCodesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**Hcpcs** | Pointer to **NullableString** | The HCPCS billing code. | [optional] 
**HcpcsDesc** | Pointer to **NullableString** | The code&#39;s short description. | [optional] 
**HcpcsDosage** | Pointer to **NullableString** | The billing unit the code is priced per. | [optional] 
**AspPaymentLimit** | Pointer to **NullableFloat32** | The Medicare payment limit per billing unit. | [optional] 
**PricedNdcs** | Pointer to **NullableInt32** | Packages priced under the code. | [optional] 

## Methods

### NewDrugAspCodesRow

`func NewDrugAspCodesRow() *DrugAspCodesRow`

NewDrugAspCodesRow instantiates a new DrugAspCodesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugAspCodesRowWithDefaults

`func NewDrugAspCodesRowWithDefaults() *DrugAspCodesRow`

NewDrugAspCodesRowWithDefaults instantiates a new DrugAspCodesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugAspCodesRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugAspCodesRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugAspCodesRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugAspCodesRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugAspCodesRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugAspCodesRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetHcpcs

`func (o *DrugAspCodesRow) GetHcpcs() string`

GetHcpcs returns the Hcpcs field if non-nil, zero value otherwise.

### GetHcpcsOk

`func (o *DrugAspCodesRow) GetHcpcsOk() (*string, bool)`

GetHcpcsOk returns a tuple with the Hcpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcpcs

`func (o *DrugAspCodesRow) SetHcpcs(v string)`

SetHcpcs sets Hcpcs field to given value.

### HasHcpcs

`func (o *DrugAspCodesRow) HasHcpcs() bool`

HasHcpcs returns a boolean if a field has been set.

### SetHcpcsNil

`func (o *DrugAspCodesRow) SetHcpcsNil(b bool)`

 SetHcpcsNil sets the value for Hcpcs to be an explicit nil

### UnsetHcpcs
`func (o *DrugAspCodesRow) UnsetHcpcs()`

UnsetHcpcs ensures that no value is present for Hcpcs, not even an explicit nil
### GetHcpcsDesc

`func (o *DrugAspCodesRow) GetHcpcsDesc() string`

GetHcpcsDesc returns the HcpcsDesc field if non-nil, zero value otherwise.

### GetHcpcsDescOk

`func (o *DrugAspCodesRow) GetHcpcsDescOk() (*string, bool)`

GetHcpcsDescOk returns a tuple with the HcpcsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcpcsDesc

`func (o *DrugAspCodesRow) SetHcpcsDesc(v string)`

SetHcpcsDesc sets HcpcsDesc field to given value.

### HasHcpcsDesc

`func (o *DrugAspCodesRow) HasHcpcsDesc() bool`

HasHcpcsDesc returns a boolean if a field has been set.

### SetHcpcsDescNil

`func (o *DrugAspCodesRow) SetHcpcsDescNil(b bool)`

 SetHcpcsDescNil sets the value for HcpcsDesc to be an explicit nil

### UnsetHcpcsDesc
`func (o *DrugAspCodesRow) UnsetHcpcsDesc()`

UnsetHcpcsDesc ensures that no value is present for HcpcsDesc, not even an explicit nil
### GetHcpcsDosage

`func (o *DrugAspCodesRow) GetHcpcsDosage() string`

GetHcpcsDosage returns the HcpcsDosage field if non-nil, zero value otherwise.

### GetHcpcsDosageOk

`func (o *DrugAspCodesRow) GetHcpcsDosageOk() (*string, bool)`

GetHcpcsDosageOk returns a tuple with the HcpcsDosage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcpcsDosage

`func (o *DrugAspCodesRow) SetHcpcsDosage(v string)`

SetHcpcsDosage sets HcpcsDosage field to given value.

### HasHcpcsDosage

`func (o *DrugAspCodesRow) HasHcpcsDosage() bool`

HasHcpcsDosage returns a boolean if a field has been set.

### SetHcpcsDosageNil

`func (o *DrugAspCodesRow) SetHcpcsDosageNil(b bool)`

 SetHcpcsDosageNil sets the value for HcpcsDosage to be an explicit nil

### UnsetHcpcsDosage
`func (o *DrugAspCodesRow) UnsetHcpcsDosage()`

UnsetHcpcsDosage ensures that no value is present for HcpcsDosage, not even an explicit nil
### GetAspPaymentLimit

`func (o *DrugAspCodesRow) GetAspPaymentLimit() float32`

GetAspPaymentLimit returns the AspPaymentLimit field if non-nil, zero value otherwise.

### GetAspPaymentLimitOk

`func (o *DrugAspCodesRow) GetAspPaymentLimitOk() (*float32, bool)`

GetAspPaymentLimitOk returns a tuple with the AspPaymentLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspPaymentLimit

`func (o *DrugAspCodesRow) SetAspPaymentLimit(v float32)`

SetAspPaymentLimit sets AspPaymentLimit field to given value.

### HasAspPaymentLimit

`func (o *DrugAspCodesRow) HasAspPaymentLimit() bool`

HasAspPaymentLimit returns a boolean if a field has been set.

### SetAspPaymentLimitNil

`func (o *DrugAspCodesRow) SetAspPaymentLimitNil(b bool)`

 SetAspPaymentLimitNil sets the value for AspPaymentLimit to be an explicit nil

### UnsetAspPaymentLimit
`func (o *DrugAspCodesRow) UnsetAspPaymentLimit()`

UnsetAspPaymentLimit ensures that no value is present for AspPaymentLimit, not even an explicit nil
### GetPricedNdcs

`func (o *DrugAspCodesRow) GetPricedNdcs() int32`

GetPricedNdcs returns the PricedNdcs field if non-nil, zero value otherwise.

### GetPricedNdcsOk

`func (o *DrugAspCodesRow) GetPricedNdcsOk() (*int32, bool)`

GetPricedNdcsOk returns a tuple with the PricedNdcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricedNdcs

`func (o *DrugAspCodesRow) SetPricedNdcs(v int32)`

SetPricedNdcs sets PricedNdcs field to given value.

### HasPricedNdcs

`func (o *DrugAspCodesRow) HasPricedNdcs() bool`

HasPricedNdcs returns a boolean if a field has been set.

### SetPricedNdcsNil

`func (o *DrugAspCodesRow) SetPricedNdcsNil(b bool)`

 SetPricedNdcsNil sets the value for PricedNdcs to be an explicit nil

### UnsetPricedNdcs
`func (o *DrugAspCodesRow) UnsetPricedNdcs()`

UnsetPricedNdcs ensures that no value is present for PricedNdcs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DrugMedicareSpendRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**Year** | Pointer to **NullableInt32** | The calendar year of the spending. | [optional] 
**TotalSpend** | Pointer to **NullableFloat32** | Medicare spending in US dollars. | [optional] 
**Claims** | Pointer to **NullableFloat32** | Claims or services. | [optional] 
**Beneficiaries** | Pointer to **NullableFloat32** | Beneficiaries, summed over payers. | [optional] 
**Payers** | Pointer to **[]string** | The Medicare programs summed (Part B, Part D). | [optional] 

## Methods

### NewDrugMedicareSpendRow

`func NewDrugMedicareSpendRow() *DrugMedicareSpendRow`

NewDrugMedicareSpendRow instantiates a new DrugMedicareSpendRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugMedicareSpendRowWithDefaults

`func NewDrugMedicareSpendRowWithDefaults() *DrugMedicareSpendRow`

NewDrugMedicareSpendRowWithDefaults instantiates a new DrugMedicareSpendRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugMedicareSpendRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugMedicareSpendRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugMedicareSpendRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugMedicareSpendRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugMedicareSpendRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugMedicareSpendRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetYear

`func (o *DrugMedicareSpendRow) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *DrugMedicareSpendRow) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *DrugMedicareSpendRow) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *DrugMedicareSpendRow) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *DrugMedicareSpendRow) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *DrugMedicareSpendRow) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetTotalSpend

`func (o *DrugMedicareSpendRow) GetTotalSpend() float32`

GetTotalSpend returns the TotalSpend field if non-nil, zero value otherwise.

### GetTotalSpendOk

`func (o *DrugMedicareSpendRow) GetTotalSpendOk() (*float32, bool)`

GetTotalSpendOk returns a tuple with the TotalSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpend

`func (o *DrugMedicareSpendRow) SetTotalSpend(v float32)`

SetTotalSpend sets TotalSpend field to given value.

### HasTotalSpend

`func (o *DrugMedicareSpendRow) HasTotalSpend() bool`

HasTotalSpend returns a boolean if a field has been set.

### SetTotalSpendNil

`func (o *DrugMedicareSpendRow) SetTotalSpendNil(b bool)`

 SetTotalSpendNil sets the value for TotalSpend to be an explicit nil

### UnsetTotalSpend
`func (o *DrugMedicareSpendRow) UnsetTotalSpend()`

UnsetTotalSpend ensures that no value is present for TotalSpend, not even an explicit nil
### GetClaims

`func (o *DrugMedicareSpendRow) GetClaims() float32`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *DrugMedicareSpendRow) GetClaimsOk() (*float32, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *DrugMedicareSpendRow) SetClaims(v float32)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *DrugMedicareSpendRow) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *DrugMedicareSpendRow) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *DrugMedicareSpendRow) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil
### GetBeneficiaries

`func (o *DrugMedicareSpendRow) GetBeneficiaries() float32`

GetBeneficiaries returns the Beneficiaries field if non-nil, zero value otherwise.

### GetBeneficiariesOk

`func (o *DrugMedicareSpendRow) GetBeneficiariesOk() (*float32, bool)`

GetBeneficiariesOk returns a tuple with the Beneficiaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaries

`func (o *DrugMedicareSpendRow) SetBeneficiaries(v float32)`

SetBeneficiaries sets Beneficiaries field to given value.

### HasBeneficiaries

`func (o *DrugMedicareSpendRow) HasBeneficiaries() bool`

HasBeneficiaries returns a boolean if a field has been set.

### SetBeneficiariesNil

`func (o *DrugMedicareSpendRow) SetBeneficiariesNil(b bool)`

 SetBeneficiariesNil sets the value for Beneficiaries to be an explicit nil

### UnsetBeneficiaries
`func (o *DrugMedicareSpendRow) UnsetBeneficiaries()`

UnsetBeneficiaries ensures that no value is present for Beneficiaries, not even an explicit nil
### GetPayers

`func (o *DrugMedicareSpendRow) GetPayers() []string`

GetPayers returns the Payers field if non-nil, zero value otherwise.

### GetPayersOk

`func (o *DrugMedicareSpendRow) GetPayersOk() (*[]string, bool)`

GetPayersOk returns a tuple with the Payers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayers

`func (o *DrugMedicareSpendRow) SetPayers(v []string)`

SetPayers sets Payers field to given value.

### HasPayers

`func (o *DrugMedicareSpendRow) HasPayers() bool`

HasPayers returns a boolean if a field has been set.

### SetPayersNil

`func (o *DrugMedicareSpendRow) SetPayersNil(b bool)`

 SetPayersNil sets the value for Payers to be an explicit nil

### UnsetPayers
`func (o *DrugMedicareSpendRow) UnsetPayers()`

UnsetPayers ensures that no value is present for Payers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



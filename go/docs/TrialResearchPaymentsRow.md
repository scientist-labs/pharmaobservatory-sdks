# TrialResearchPaymentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**PaymentKey** | Pointer to **NullableString** | A stable key for the trial, payer and program year row. | [optional] 
**PayerName** | Pointer to **NullableString** | The paying manufacturer&#39;s name as reported. | [optional] 
**PayerOrgId** | Pointer to **NullableString** | Legal-entity level: the payer rolled up to its ultimate parent organization record (null when unresolved). Payments from subsidiaries of one parent share one row per trial and program year. | [optional] 
**PayerCompanyId** | Pointer to **NullableString** | Company level: the company the payer belongs to (null when unresolved). | [optional] 
**PayerCompanyName** | Pointer to **NullableString** | The payer&#39;s company name. | [optional] 
**PayerGroupId** | Pointer to **NullableString** | Group level: the current ultimate parent of the payer&#39;s company. | [optional] 
**PayerGroupName** | Pointer to **NullableString** | The ultimate parent company&#39;s name. | [optional] 
**ProgramYear** | Pointer to **NullableInt32** | The Open Payments program year of the payments. | [optional] 
**ResearchPayments** | Pointer to **NullableInt32** | Payment records behind the row. | [optional] 
**TotalAmountUsd** | Pointer to **NullableFloat64** | Total reported dollars for the row, counted once. | [optional] 
**PiCount** | Pointer to **NullableInt32** | Distinct resolved principal investigators named on the payment records. | [optional] 
**AnyDisputed** | Pointer to **NullableBool** | True when any payment record in the row was disputed by the recipient. | [optional] 
**AnyPreclinical** | Pointer to **NullableBool** | True when any payment record in the row was marked as preclinical research. | [optional] 

## Methods

### NewTrialResearchPaymentsRow

`func NewTrialResearchPaymentsRow() *TrialResearchPaymentsRow`

NewTrialResearchPaymentsRow instantiates a new TrialResearchPaymentsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialResearchPaymentsRowWithDefaults

`func NewTrialResearchPaymentsRowWithDefaults() *TrialResearchPaymentsRow`

NewTrialResearchPaymentsRowWithDefaults instantiates a new TrialResearchPaymentsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialResearchPaymentsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialResearchPaymentsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialResearchPaymentsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialResearchPaymentsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialResearchPaymentsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialResearchPaymentsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetPaymentKey

`func (o *TrialResearchPaymentsRow) GetPaymentKey() string`

GetPaymentKey returns the PaymentKey field if non-nil, zero value otherwise.

### GetPaymentKeyOk

`func (o *TrialResearchPaymentsRow) GetPaymentKeyOk() (*string, bool)`

GetPaymentKeyOk returns a tuple with the PaymentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentKey

`func (o *TrialResearchPaymentsRow) SetPaymentKey(v string)`

SetPaymentKey sets PaymentKey field to given value.

### HasPaymentKey

`func (o *TrialResearchPaymentsRow) HasPaymentKey() bool`

HasPaymentKey returns a boolean if a field has been set.

### SetPaymentKeyNil

`func (o *TrialResearchPaymentsRow) SetPaymentKeyNil(b bool)`

 SetPaymentKeyNil sets the value for PaymentKey to be an explicit nil

### UnsetPaymentKey
`func (o *TrialResearchPaymentsRow) UnsetPaymentKey()`

UnsetPaymentKey ensures that no value is present for PaymentKey, not even an explicit nil
### GetPayerName

`func (o *TrialResearchPaymentsRow) GetPayerName() string`

GetPayerName returns the PayerName field if non-nil, zero value otherwise.

### GetPayerNameOk

`func (o *TrialResearchPaymentsRow) GetPayerNameOk() (*string, bool)`

GetPayerNameOk returns a tuple with the PayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerName

`func (o *TrialResearchPaymentsRow) SetPayerName(v string)`

SetPayerName sets PayerName field to given value.

### HasPayerName

`func (o *TrialResearchPaymentsRow) HasPayerName() bool`

HasPayerName returns a boolean if a field has been set.

### SetPayerNameNil

`func (o *TrialResearchPaymentsRow) SetPayerNameNil(b bool)`

 SetPayerNameNil sets the value for PayerName to be an explicit nil

### UnsetPayerName
`func (o *TrialResearchPaymentsRow) UnsetPayerName()`

UnsetPayerName ensures that no value is present for PayerName, not even an explicit nil
### GetPayerOrgId

`func (o *TrialResearchPaymentsRow) GetPayerOrgId() string`

GetPayerOrgId returns the PayerOrgId field if non-nil, zero value otherwise.

### GetPayerOrgIdOk

`func (o *TrialResearchPaymentsRow) GetPayerOrgIdOk() (*string, bool)`

GetPayerOrgIdOk returns a tuple with the PayerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerOrgId

`func (o *TrialResearchPaymentsRow) SetPayerOrgId(v string)`

SetPayerOrgId sets PayerOrgId field to given value.

### HasPayerOrgId

`func (o *TrialResearchPaymentsRow) HasPayerOrgId() bool`

HasPayerOrgId returns a boolean if a field has been set.

### SetPayerOrgIdNil

`func (o *TrialResearchPaymentsRow) SetPayerOrgIdNil(b bool)`

 SetPayerOrgIdNil sets the value for PayerOrgId to be an explicit nil

### UnsetPayerOrgId
`func (o *TrialResearchPaymentsRow) UnsetPayerOrgId()`

UnsetPayerOrgId ensures that no value is present for PayerOrgId, not even an explicit nil
### GetPayerCompanyId

`func (o *TrialResearchPaymentsRow) GetPayerCompanyId() string`

GetPayerCompanyId returns the PayerCompanyId field if non-nil, zero value otherwise.

### GetPayerCompanyIdOk

`func (o *TrialResearchPaymentsRow) GetPayerCompanyIdOk() (*string, bool)`

GetPayerCompanyIdOk returns a tuple with the PayerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerCompanyId

`func (o *TrialResearchPaymentsRow) SetPayerCompanyId(v string)`

SetPayerCompanyId sets PayerCompanyId field to given value.

### HasPayerCompanyId

`func (o *TrialResearchPaymentsRow) HasPayerCompanyId() bool`

HasPayerCompanyId returns a boolean if a field has been set.

### SetPayerCompanyIdNil

`func (o *TrialResearchPaymentsRow) SetPayerCompanyIdNil(b bool)`

 SetPayerCompanyIdNil sets the value for PayerCompanyId to be an explicit nil

### UnsetPayerCompanyId
`func (o *TrialResearchPaymentsRow) UnsetPayerCompanyId()`

UnsetPayerCompanyId ensures that no value is present for PayerCompanyId, not even an explicit nil
### GetPayerCompanyName

`func (o *TrialResearchPaymentsRow) GetPayerCompanyName() string`

GetPayerCompanyName returns the PayerCompanyName field if non-nil, zero value otherwise.

### GetPayerCompanyNameOk

`func (o *TrialResearchPaymentsRow) GetPayerCompanyNameOk() (*string, bool)`

GetPayerCompanyNameOk returns a tuple with the PayerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerCompanyName

`func (o *TrialResearchPaymentsRow) SetPayerCompanyName(v string)`

SetPayerCompanyName sets PayerCompanyName field to given value.

### HasPayerCompanyName

`func (o *TrialResearchPaymentsRow) HasPayerCompanyName() bool`

HasPayerCompanyName returns a boolean if a field has been set.

### SetPayerCompanyNameNil

`func (o *TrialResearchPaymentsRow) SetPayerCompanyNameNil(b bool)`

 SetPayerCompanyNameNil sets the value for PayerCompanyName to be an explicit nil

### UnsetPayerCompanyName
`func (o *TrialResearchPaymentsRow) UnsetPayerCompanyName()`

UnsetPayerCompanyName ensures that no value is present for PayerCompanyName, not even an explicit nil
### GetPayerGroupId

`func (o *TrialResearchPaymentsRow) GetPayerGroupId() string`

GetPayerGroupId returns the PayerGroupId field if non-nil, zero value otherwise.

### GetPayerGroupIdOk

`func (o *TrialResearchPaymentsRow) GetPayerGroupIdOk() (*string, bool)`

GetPayerGroupIdOk returns a tuple with the PayerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerGroupId

`func (o *TrialResearchPaymentsRow) SetPayerGroupId(v string)`

SetPayerGroupId sets PayerGroupId field to given value.

### HasPayerGroupId

`func (o *TrialResearchPaymentsRow) HasPayerGroupId() bool`

HasPayerGroupId returns a boolean if a field has been set.

### SetPayerGroupIdNil

`func (o *TrialResearchPaymentsRow) SetPayerGroupIdNil(b bool)`

 SetPayerGroupIdNil sets the value for PayerGroupId to be an explicit nil

### UnsetPayerGroupId
`func (o *TrialResearchPaymentsRow) UnsetPayerGroupId()`

UnsetPayerGroupId ensures that no value is present for PayerGroupId, not even an explicit nil
### GetPayerGroupName

`func (o *TrialResearchPaymentsRow) GetPayerGroupName() string`

GetPayerGroupName returns the PayerGroupName field if non-nil, zero value otherwise.

### GetPayerGroupNameOk

`func (o *TrialResearchPaymentsRow) GetPayerGroupNameOk() (*string, bool)`

GetPayerGroupNameOk returns a tuple with the PayerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerGroupName

`func (o *TrialResearchPaymentsRow) SetPayerGroupName(v string)`

SetPayerGroupName sets PayerGroupName field to given value.

### HasPayerGroupName

`func (o *TrialResearchPaymentsRow) HasPayerGroupName() bool`

HasPayerGroupName returns a boolean if a field has been set.

### SetPayerGroupNameNil

`func (o *TrialResearchPaymentsRow) SetPayerGroupNameNil(b bool)`

 SetPayerGroupNameNil sets the value for PayerGroupName to be an explicit nil

### UnsetPayerGroupName
`func (o *TrialResearchPaymentsRow) UnsetPayerGroupName()`

UnsetPayerGroupName ensures that no value is present for PayerGroupName, not even an explicit nil
### GetProgramYear

`func (o *TrialResearchPaymentsRow) GetProgramYear() int32`

GetProgramYear returns the ProgramYear field if non-nil, zero value otherwise.

### GetProgramYearOk

`func (o *TrialResearchPaymentsRow) GetProgramYearOk() (*int32, bool)`

GetProgramYearOk returns a tuple with the ProgramYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramYear

`func (o *TrialResearchPaymentsRow) SetProgramYear(v int32)`

SetProgramYear sets ProgramYear field to given value.

### HasProgramYear

`func (o *TrialResearchPaymentsRow) HasProgramYear() bool`

HasProgramYear returns a boolean if a field has been set.

### SetProgramYearNil

`func (o *TrialResearchPaymentsRow) SetProgramYearNil(b bool)`

 SetProgramYearNil sets the value for ProgramYear to be an explicit nil

### UnsetProgramYear
`func (o *TrialResearchPaymentsRow) UnsetProgramYear()`

UnsetProgramYear ensures that no value is present for ProgramYear, not even an explicit nil
### GetResearchPayments

`func (o *TrialResearchPaymentsRow) GetResearchPayments() int32`

GetResearchPayments returns the ResearchPayments field if non-nil, zero value otherwise.

### GetResearchPaymentsOk

`func (o *TrialResearchPaymentsRow) GetResearchPaymentsOk() (*int32, bool)`

GetResearchPaymentsOk returns a tuple with the ResearchPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearchPayments

`func (o *TrialResearchPaymentsRow) SetResearchPayments(v int32)`

SetResearchPayments sets ResearchPayments field to given value.

### HasResearchPayments

`func (o *TrialResearchPaymentsRow) HasResearchPayments() bool`

HasResearchPayments returns a boolean if a field has been set.

### SetResearchPaymentsNil

`func (o *TrialResearchPaymentsRow) SetResearchPaymentsNil(b bool)`

 SetResearchPaymentsNil sets the value for ResearchPayments to be an explicit nil

### UnsetResearchPayments
`func (o *TrialResearchPaymentsRow) UnsetResearchPayments()`

UnsetResearchPayments ensures that no value is present for ResearchPayments, not even an explicit nil
### GetTotalAmountUsd

`func (o *TrialResearchPaymentsRow) GetTotalAmountUsd() float64`

GetTotalAmountUsd returns the TotalAmountUsd field if non-nil, zero value otherwise.

### GetTotalAmountUsdOk

`func (o *TrialResearchPaymentsRow) GetTotalAmountUsdOk() (*float64, bool)`

GetTotalAmountUsdOk returns a tuple with the TotalAmountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountUsd

`func (o *TrialResearchPaymentsRow) SetTotalAmountUsd(v float64)`

SetTotalAmountUsd sets TotalAmountUsd field to given value.

### HasTotalAmountUsd

`func (o *TrialResearchPaymentsRow) HasTotalAmountUsd() bool`

HasTotalAmountUsd returns a boolean if a field has been set.

### SetTotalAmountUsdNil

`func (o *TrialResearchPaymentsRow) SetTotalAmountUsdNil(b bool)`

 SetTotalAmountUsdNil sets the value for TotalAmountUsd to be an explicit nil

### UnsetTotalAmountUsd
`func (o *TrialResearchPaymentsRow) UnsetTotalAmountUsd()`

UnsetTotalAmountUsd ensures that no value is present for TotalAmountUsd, not even an explicit nil
### GetPiCount

`func (o *TrialResearchPaymentsRow) GetPiCount() int32`

GetPiCount returns the PiCount field if non-nil, zero value otherwise.

### GetPiCountOk

`func (o *TrialResearchPaymentsRow) GetPiCountOk() (*int32, bool)`

GetPiCountOk returns a tuple with the PiCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiCount

`func (o *TrialResearchPaymentsRow) SetPiCount(v int32)`

SetPiCount sets PiCount field to given value.

### HasPiCount

`func (o *TrialResearchPaymentsRow) HasPiCount() bool`

HasPiCount returns a boolean if a field has been set.

### SetPiCountNil

`func (o *TrialResearchPaymentsRow) SetPiCountNil(b bool)`

 SetPiCountNil sets the value for PiCount to be an explicit nil

### UnsetPiCount
`func (o *TrialResearchPaymentsRow) UnsetPiCount()`

UnsetPiCount ensures that no value is present for PiCount, not even an explicit nil
### GetAnyDisputed

`func (o *TrialResearchPaymentsRow) GetAnyDisputed() bool`

GetAnyDisputed returns the AnyDisputed field if non-nil, zero value otherwise.

### GetAnyDisputedOk

`func (o *TrialResearchPaymentsRow) GetAnyDisputedOk() (*bool, bool)`

GetAnyDisputedOk returns a tuple with the AnyDisputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyDisputed

`func (o *TrialResearchPaymentsRow) SetAnyDisputed(v bool)`

SetAnyDisputed sets AnyDisputed field to given value.

### HasAnyDisputed

`func (o *TrialResearchPaymentsRow) HasAnyDisputed() bool`

HasAnyDisputed returns a boolean if a field has been set.

### SetAnyDisputedNil

`func (o *TrialResearchPaymentsRow) SetAnyDisputedNil(b bool)`

 SetAnyDisputedNil sets the value for AnyDisputed to be an explicit nil

### UnsetAnyDisputed
`func (o *TrialResearchPaymentsRow) UnsetAnyDisputed()`

UnsetAnyDisputed ensures that no value is present for AnyDisputed, not even an explicit nil
### GetAnyPreclinical

`func (o *TrialResearchPaymentsRow) GetAnyPreclinical() bool`

GetAnyPreclinical returns the AnyPreclinical field if non-nil, zero value otherwise.

### GetAnyPreclinicalOk

`func (o *TrialResearchPaymentsRow) GetAnyPreclinicalOk() (*bool, bool)`

GetAnyPreclinicalOk returns a tuple with the AnyPreclinical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyPreclinical

`func (o *TrialResearchPaymentsRow) SetAnyPreclinical(v bool)`

SetAnyPreclinical sets AnyPreclinical field to given value.

### HasAnyPreclinical

`func (o *TrialResearchPaymentsRow) HasAnyPreclinical() bool`

HasAnyPreclinical returns a boolean if a field has been set.

### SetAnyPreclinicalNil

`func (o *TrialResearchPaymentsRow) SetAnyPreclinicalNil(b bool)`

 SetAnyPreclinicalNil sets the value for AnyPreclinical to be an explicit nil

### UnsetAnyPreclinical
`func (o *TrialResearchPaymentsRow) UnsetAnyPreclinical()`

UnsetAnyPreclinical ensures that no value is present for AnyPreclinical, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



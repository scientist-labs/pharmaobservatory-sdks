# InvestorPrivateFundsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundId** | Pointer to **NullableString** | The fund key, advfund:&lt;PF-ID&gt;. | [optional] 
**PfId** | Pointer to **NullableString** | The SEC private fund identification number (PF-ID). | [optional] 
**FundName** | Pointer to **NullableString** | The fund&#39;s name as the adviser filed it. | [optional] 
**FundType** | Pointer to **NullableString** | The fund type the adviser reports on Form ADV Schedule D. | [optional] 
**GrossAssetValue** | Pointer to **NullableFloat32** | The fund&#39;s gross asset value in whole US dollars, as filed. | [optional] 
**BeneficialOwnerCount** | Pointer to **NullableInt32** | The number of beneficial owners the fund reports. It counts the fund&#39;s own investors, not investors in this directory. | [optional] 
**IsMaster** | Pointer to **NullableBool** | True when the fund is a master fund. | [optional] 
**IsFeeder** | Pointer to **NullableBool** | True when the fund is a feeder fund. | [optional] 
**MasterFundName** | Pointer to **NullableString** | The master fund&#39;s name when this fund is a feeder. | [optional] 
**MasterPfId** | Pointer to **NullableString** | The master fund&#39;s PF-ID when this fund is a feeder. | [optional] 
**AdviserCount** | Pointer to **NullableInt64** | Advisers that report this PF-ID. | [optional] 
**InvestorId** | Pointer to **NullableString** | The investor key of the adviser the fund is listed under. | [optional] 
**AsOfDate** | Pointer to **NullableString** | The date of the filing the fund row comes from. | [optional] 
**SourceUrl** | Pointer to **NullableString** | A link to the adviser&#39;s Form ADV filing. | [optional] 

## Methods

### NewInvestorPrivateFundsRow

`func NewInvestorPrivateFundsRow() *InvestorPrivateFundsRow`

NewInvestorPrivateFundsRow instantiates a new InvestorPrivateFundsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestorPrivateFundsRowWithDefaults

`func NewInvestorPrivateFundsRowWithDefaults() *InvestorPrivateFundsRow`

NewInvestorPrivateFundsRowWithDefaults instantiates a new InvestorPrivateFundsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundId

`func (o *InvestorPrivateFundsRow) GetFundId() string`

GetFundId returns the FundId field if non-nil, zero value otherwise.

### GetFundIdOk

`func (o *InvestorPrivateFundsRow) GetFundIdOk() (*string, bool)`

GetFundIdOk returns a tuple with the FundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundId

`func (o *InvestorPrivateFundsRow) SetFundId(v string)`

SetFundId sets FundId field to given value.

### HasFundId

`func (o *InvestorPrivateFundsRow) HasFundId() bool`

HasFundId returns a boolean if a field has been set.

### SetFundIdNil

`func (o *InvestorPrivateFundsRow) SetFundIdNil(b bool)`

 SetFundIdNil sets the value for FundId to be an explicit nil

### UnsetFundId
`func (o *InvestorPrivateFundsRow) UnsetFundId()`

UnsetFundId ensures that no value is present for FundId, not even an explicit nil
### GetPfId

`func (o *InvestorPrivateFundsRow) GetPfId() string`

GetPfId returns the PfId field if non-nil, zero value otherwise.

### GetPfIdOk

`func (o *InvestorPrivateFundsRow) GetPfIdOk() (*string, bool)`

GetPfIdOk returns a tuple with the PfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfId

`func (o *InvestorPrivateFundsRow) SetPfId(v string)`

SetPfId sets PfId field to given value.

### HasPfId

`func (o *InvestorPrivateFundsRow) HasPfId() bool`

HasPfId returns a boolean if a field has been set.

### SetPfIdNil

`func (o *InvestorPrivateFundsRow) SetPfIdNil(b bool)`

 SetPfIdNil sets the value for PfId to be an explicit nil

### UnsetPfId
`func (o *InvestorPrivateFundsRow) UnsetPfId()`

UnsetPfId ensures that no value is present for PfId, not even an explicit nil
### GetFundName

`func (o *InvestorPrivateFundsRow) GetFundName() string`

GetFundName returns the FundName field if non-nil, zero value otherwise.

### GetFundNameOk

`func (o *InvestorPrivateFundsRow) GetFundNameOk() (*string, bool)`

GetFundNameOk returns a tuple with the FundName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundName

`func (o *InvestorPrivateFundsRow) SetFundName(v string)`

SetFundName sets FundName field to given value.

### HasFundName

`func (o *InvestorPrivateFundsRow) HasFundName() bool`

HasFundName returns a boolean if a field has been set.

### SetFundNameNil

`func (o *InvestorPrivateFundsRow) SetFundNameNil(b bool)`

 SetFundNameNil sets the value for FundName to be an explicit nil

### UnsetFundName
`func (o *InvestorPrivateFundsRow) UnsetFundName()`

UnsetFundName ensures that no value is present for FundName, not even an explicit nil
### GetFundType

`func (o *InvestorPrivateFundsRow) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *InvestorPrivateFundsRow) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *InvestorPrivateFundsRow) SetFundType(v string)`

SetFundType sets FundType field to given value.

### HasFundType

`func (o *InvestorPrivateFundsRow) HasFundType() bool`

HasFundType returns a boolean if a field has been set.

### SetFundTypeNil

`func (o *InvestorPrivateFundsRow) SetFundTypeNil(b bool)`

 SetFundTypeNil sets the value for FundType to be an explicit nil

### UnsetFundType
`func (o *InvestorPrivateFundsRow) UnsetFundType()`

UnsetFundType ensures that no value is present for FundType, not even an explicit nil
### GetGrossAssetValue

`func (o *InvestorPrivateFundsRow) GetGrossAssetValue() float32`

GetGrossAssetValue returns the GrossAssetValue field if non-nil, zero value otherwise.

### GetGrossAssetValueOk

`func (o *InvestorPrivateFundsRow) GetGrossAssetValueOk() (*float32, bool)`

GetGrossAssetValueOk returns a tuple with the GrossAssetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAssetValue

`func (o *InvestorPrivateFundsRow) SetGrossAssetValue(v float32)`

SetGrossAssetValue sets GrossAssetValue field to given value.

### HasGrossAssetValue

`func (o *InvestorPrivateFundsRow) HasGrossAssetValue() bool`

HasGrossAssetValue returns a boolean if a field has been set.

### SetGrossAssetValueNil

`func (o *InvestorPrivateFundsRow) SetGrossAssetValueNil(b bool)`

 SetGrossAssetValueNil sets the value for GrossAssetValue to be an explicit nil

### UnsetGrossAssetValue
`func (o *InvestorPrivateFundsRow) UnsetGrossAssetValue()`

UnsetGrossAssetValue ensures that no value is present for GrossAssetValue, not even an explicit nil
### GetBeneficialOwnerCount

`func (o *InvestorPrivateFundsRow) GetBeneficialOwnerCount() int32`

GetBeneficialOwnerCount returns the BeneficialOwnerCount field if non-nil, zero value otherwise.

### GetBeneficialOwnerCountOk

`func (o *InvestorPrivateFundsRow) GetBeneficialOwnerCountOk() (*int32, bool)`

GetBeneficialOwnerCountOk returns a tuple with the BeneficialOwnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwnerCount

`func (o *InvestorPrivateFundsRow) SetBeneficialOwnerCount(v int32)`

SetBeneficialOwnerCount sets BeneficialOwnerCount field to given value.

### HasBeneficialOwnerCount

`func (o *InvestorPrivateFundsRow) HasBeneficialOwnerCount() bool`

HasBeneficialOwnerCount returns a boolean if a field has been set.

### SetBeneficialOwnerCountNil

`func (o *InvestorPrivateFundsRow) SetBeneficialOwnerCountNil(b bool)`

 SetBeneficialOwnerCountNil sets the value for BeneficialOwnerCount to be an explicit nil

### UnsetBeneficialOwnerCount
`func (o *InvestorPrivateFundsRow) UnsetBeneficialOwnerCount()`

UnsetBeneficialOwnerCount ensures that no value is present for BeneficialOwnerCount, not even an explicit nil
### GetIsMaster

`func (o *InvestorPrivateFundsRow) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *InvestorPrivateFundsRow) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *InvestorPrivateFundsRow) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *InvestorPrivateFundsRow) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### SetIsMasterNil

`func (o *InvestorPrivateFundsRow) SetIsMasterNil(b bool)`

 SetIsMasterNil sets the value for IsMaster to be an explicit nil

### UnsetIsMaster
`func (o *InvestorPrivateFundsRow) UnsetIsMaster()`

UnsetIsMaster ensures that no value is present for IsMaster, not even an explicit nil
### GetIsFeeder

`func (o *InvestorPrivateFundsRow) GetIsFeeder() bool`

GetIsFeeder returns the IsFeeder field if non-nil, zero value otherwise.

### GetIsFeederOk

`func (o *InvestorPrivateFundsRow) GetIsFeederOk() (*bool, bool)`

GetIsFeederOk returns a tuple with the IsFeeder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeeder

`func (o *InvestorPrivateFundsRow) SetIsFeeder(v bool)`

SetIsFeeder sets IsFeeder field to given value.

### HasIsFeeder

`func (o *InvestorPrivateFundsRow) HasIsFeeder() bool`

HasIsFeeder returns a boolean if a field has been set.

### SetIsFeederNil

`func (o *InvestorPrivateFundsRow) SetIsFeederNil(b bool)`

 SetIsFeederNil sets the value for IsFeeder to be an explicit nil

### UnsetIsFeeder
`func (o *InvestorPrivateFundsRow) UnsetIsFeeder()`

UnsetIsFeeder ensures that no value is present for IsFeeder, not even an explicit nil
### GetMasterFundName

`func (o *InvestorPrivateFundsRow) GetMasterFundName() string`

GetMasterFundName returns the MasterFundName field if non-nil, zero value otherwise.

### GetMasterFundNameOk

`func (o *InvestorPrivateFundsRow) GetMasterFundNameOk() (*string, bool)`

GetMasterFundNameOk returns a tuple with the MasterFundName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterFundName

`func (o *InvestorPrivateFundsRow) SetMasterFundName(v string)`

SetMasterFundName sets MasterFundName field to given value.

### HasMasterFundName

`func (o *InvestorPrivateFundsRow) HasMasterFundName() bool`

HasMasterFundName returns a boolean if a field has been set.

### SetMasterFundNameNil

`func (o *InvestorPrivateFundsRow) SetMasterFundNameNil(b bool)`

 SetMasterFundNameNil sets the value for MasterFundName to be an explicit nil

### UnsetMasterFundName
`func (o *InvestorPrivateFundsRow) UnsetMasterFundName()`

UnsetMasterFundName ensures that no value is present for MasterFundName, not even an explicit nil
### GetMasterPfId

`func (o *InvestorPrivateFundsRow) GetMasterPfId() string`

GetMasterPfId returns the MasterPfId field if non-nil, zero value otherwise.

### GetMasterPfIdOk

`func (o *InvestorPrivateFundsRow) GetMasterPfIdOk() (*string, bool)`

GetMasterPfIdOk returns a tuple with the MasterPfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPfId

`func (o *InvestorPrivateFundsRow) SetMasterPfId(v string)`

SetMasterPfId sets MasterPfId field to given value.

### HasMasterPfId

`func (o *InvestorPrivateFundsRow) HasMasterPfId() bool`

HasMasterPfId returns a boolean if a field has been set.

### SetMasterPfIdNil

`func (o *InvestorPrivateFundsRow) SetMasterPfIdNil(b bool)`

 SetMasterPfIdNil sets the value for MasterPfId to be an explicit nil

### UnsetMasterPfId
`func (o *InvestorPrivateFundsRow) UnsetMasterPfId()`

UnsetMasterPfId ensures that no value is present for MasterPfId, not even an explicit nil
### GetAdviserCount

`func (o *InvestorPrivateFundsRow) GetAdviserCount() int64`

GetAdviserCount returns the AdviserCount field if non-nil, zero value otherwise.

### GetAdviserCountOk

`func (o *InvestorPrivateFundsRow) GetAdviserCountOk() (*int64, bool)`

GetAdviserCountOk returns a tuple with the AdviserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdviserCount

`func (o *InvestorPrivateFundsRow) SetAdviserCount(v int64)`

SetAdviserCount sets AdviserCount field to given value.

### HasAdviserCount

`func (o *InvestorPrivateFundsRow) HasAdviserCount() bool`

HasAdviserCount returns a boolean if a field has been set.

### SetAdviserCountNil

`func (o *InvestorPrivateFundsRow) SetAdviserCountNil(b bool)`

 SetAdviserCountNil sets the value for AdviserCount to be an explicit nil

### UnsetAdviserCount
`func (o *InvestorPrivateFundsRow) UnsetAdviserCount()`

UnsetAdviserCount ensures that no value is present for AdviserCount, not even an explicit nil
### GetInvestorId

`func (o *InvestorPrivateFundsRow) GetInvestorId() string`

GetInvestorId returns the InvestorId field if non-nil, zero value otherwise.

### GetInvestorIdOk

`func (o *InvestorPrivateFundsRow) GetInvestorIdOk() (*string, bool)`

GetInvestorIdOk returns a tuple with the InvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorId

`func (o *InvestorPrivateFundsRow) SetInvestorId(v string)`

SetInvestorId sets InvestorId field to given value.

### HasInvestorId

`func (o *InvestorPrivateFundsRow) HasInvestorId() bool`

HasInvestorId returns a boolean if a field has been set.

### SetInvestorIdNil

`func (o *InvestorPrivateFundsRow) SetInvestorIdNil(b bool)`

 SetInvestorIdNil sets the value for InvestorId to be an explicit nil

### UnsetInvestorId
`func (o *InvestorPrivateFundsRow) UnsetInvestorId()`

UnsetInvestorId ensures that no value is present for InvestorId, not even an explicit nil
### GetAsOfDate

`func (o *InvestorPrivateFundsRow) GetAsOfDate() string`

GetAsOfDate returns the AsOfDate field if non-nil, zero value otherwise.

### GetAsOfDateOk

`func (o *InvestorPrivateFundsRow) GetAsOfDateOk() (*string, bool)`

GetAsOfDateOk returns a tuple with the AsOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfDate

`func (o *InvestorPrivateFundsRow) SetAsOfDate(v string)`

SetAsOfDate sets AsOfDate field to given value.

### HasAsOfDate

`func (o *InvestorPrivateFundsRow) HasAsOfDate() bool`

HasAsOfDate returns a boolean if a field has been set.

### SetAsOfDateNil

`func (o *InvestorPrivateFundsRow) SetAsOfDateNil(b bool)`

 SetAsOfDateNil sets the value for AsOfDate to be an explicit nil

### UnsetAsOfDate
`func (o *InvestorPrivateFundsRow) UnsetAsOfDate()`

UnsetAsOfDate ensures that no value is present for AsOfDate, not even an explicit nil
### GetSourceUrl

`func (o *InvestorPrivateFundsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *InvestorPrivateFundsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *InvestorPrivateFundsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *InvestorPrivateFundsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *InvestorPrivateFundsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *InvestorPrivateFundsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



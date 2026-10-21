# InvestorFundsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestorId** | Pointer to **NullableString** | The investor key. | [optional] 
**FundType** | Pointer to **NullableString** | The Form ADV Item 7.B fund type: hedge, liquidity, pe (private equity), real_estate, securitized, vc (venture capital) or other. | [optional] 
**FundCount** | Pointer to **NullableInt32** | Funds of this type the adviser reports. | [optional] 
**TotalPrivateFundGav** | Pointer to **NullableFloat32** | The adviser&#39;s firm-level total private-fund gross asset value in whole US dollars, repeated on every row of the investor. Do not add it across rows. | [optional] 

## Methods

### NewInvestorFundsRow

`func NewInvestorFundsRow() *InvestorFundsRow`

NewInvestorFundsRow instantiates a new InvestorFundsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestorFundsRowWithDefaults

`func NewInvestorFundsRowWithDefaults() *InvestorFundsRow`

NewInvestorFundsRowWithDefaults instantiates a new InvestorFundsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestorId

`func (o *InvestorFundsRow) GetInvestorId() string`

GetInvestorId returns the InvestorId field if non-nil, zero value otherwise.

### GetInvestorIdOk

`func (o *InvestorFundsRow) GetInvestorIdOk() (*string, bool)`

GetInvestorIdOk returns a tuple with the InvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorId

`func (o *InvestorFundsRow) SetInvestorId(v string)`

SetInvestorId sets InvestorId field to given value.

### HasInvestorId

`func (o *InvestorFundsRow) HasInvestorId() bool`

HasInvestorId returns a boolean if a field has been set.

### SetInvestorIdNil

`func (o *InvestorFundsRow) SetInvestorIdNil(b bool)`

 SetInvestorIdNil sets the value for InvestorId to be an explicit nil

### UnsetInvestorId
`func (o *InvestorFundsRow) UnsetInvestorId()`

UnsetInvestorId ensures that no value is present for InvestorId, not even an explicit nil
### GetFundType

`func (o *InvestorFundsRow) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *InvestorFundsRow) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *InvestorFundsRow) SetFundType(v string)`

SetFundType sets FundType field to given value.

### HasFundType

`func (o *InvestorFundsRow) HasFundType() bool`

HasFundType returns a boolean if a field has been set.

### SetFundTypeNil

`func (o *InvestorFundsRow) SetFundTypeNil(b bool)`

 SetFundTypeNil sets the value for FundType to be an explicit nil

### UnsetFundType
`func (o *InvestorFundsRow) UnsetFundType()`

UnsetFundType ensures that no value is present for FundType, not even an explicit nil
### GetFundCount

`func (o *InvestorFundsRow) GetFundCount() int32`

GetFundCount returns the FundCount field if non-nil, zero value otherwise.

### GetFundCountOk

`func (o *InvestorFundsRow) GetFundCountOk() (*int32, bool)`

GetFundCountOk returns a tuple with the FundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundCount

`func (o *InvestorFundsRow) SetFundCount(v int32)`

SetFundCount sets FundCount field to given value.

### HasFundCount

`func (o *InvestorFundsRow) HasFundCount() bool`

HasFundCount returns a boolean if a field has been set.

### SetFundCountNil

`func (o *InvestorFundsRow) SetFundCountNil(b bool)`

 SetFundCountNil sets the value for FundCount to be an explicit nil

### UnsetFundCount
`func (o *InvestorFundsRow) UnsetFundCount()`

UnsetFundCount ensures that no value is present for FundCount, not even an explicit nil
### GetTotalPrivateFundGav

`func (o *InvestorFundsRow) GetTotalPrivateFundGav() float32`

GetTotalPrivateFundGav returns the TotalPrivateFundGav field if non-nil, zero value otherwise.

### GetTotalPrivateFundGavOk

`func (o *InvestorFundsRow) GetTotalPrivateFundGavOk() (*float32, bool)`

GetTotalPrivateFundGavOk returns a tuple with the TotalPrivateFundGav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrivateFundGav

`func (o *InvestorFundsRow) SetTotalPrivateFundGav(v float32)`

SetTotalPrivateFundGav sets TotalPrivateFundGav field to given value.

### HasTotalPrivateFundGav

`func (o *InvestorFundsRow) HasTotalPrivateFundGav() bool`

HasTotalPrivateFundGav returns a boolean if a field has been set.

### SetTotalPrivateFundGavNil

`func (o *InvestorFundsRow) SetTotalPrivateFundGavNil(b bool)`

 SetTotalPrivateFundGavNil sets the value for TotalPrivateFundGav to be an explicit nil

### UnsetTotalPrivateFundGav
`func (o *InvestorFundsRow) UnsetTotalPrivateFundGav()`

UnsetTotalPrivateFundGav ensures that no value is present for TotalPrivateFundGav, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



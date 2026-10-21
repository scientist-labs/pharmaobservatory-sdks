# InvestorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestorId** | Pointer to **NullableString** | The investor key, the id the web page uses: adv:&lt;CRD&gt; for a Form ADV adviser, lei:&lt;LEI&gt; where one is known, agency:&lt;slug&gt; for an NIH RePORTER funding agency, investor:&lt;hash&gt; for a name-resolved investor. | [optional] 
**LegalName** | Pointer to **NullableString** | The investor&#39;s legal or display name. | [optional] 
**InvestorType** | Pointer to **NullableString** | venture_capital, private_equity, hedge_fund (from the plurality of the adviser&#39;s Form ADV Item 7.B fund types), government_agency or other. | [optional] 
**Source** | Pointer to **NullableString** | The source lane: sec_adv (Form ADV advisers) or reporter (NIH RePORTER funding agencies). | [optional] 
**Country** | Pointer to **NullableString** | The adviser&#39;s main-office country, or United States for an NIH funding agency. | [optional] 
**RaumUsd** | Pointer to **NullableFloat32** | Regulatory assets under management (RAUM) in whole US dollars, as the adviser reports it on Form ADV. Null for an agency or an adviser that reports none. | [optional] 
**RaumBand** | Pointer to **NullableString** | The RAUM band the amount falls in. | [optional] 
**IsEra** | Pointer to **NullableBool** | True for an exempt reporting adviser (ERA), false for a registered adviser, null for an agency. | [optional] 
**IsLifeSciencesFocused** | Pointer to **NullableBool** | True when the investor is life-sciences focused, false when it is not, null when no signal exists. Read it with is_life_sciences_focused_basis. | [optional] 
**IsLifeSciencesFocusedBasis** | Pointer to **NullableString** | Where is_life_sciences_focused came from: human or agent (a recorded classification), override (a curator&#39;s call), edge_share (the share of its portfolio edges in life-sciences issuers) or classifier (the Form 13F biomedical value share). Null exactly when the flag is null. | [optional] 
**PrivateFundCount** | Pointer to **NullableInt32** | Private funds the adviser reports on Form ADV Item 7.B. | [optional] 
**TotalPrivateFundGav** | Pointer to **NullableFloat32** | The adviser&#39;s total private-fund gross asset value (GAV) in whole US dollars, a firm-level figure. | [optional] 
**PortfolioCoCount** | Pointer to **NullableInt64** | Distinct portfolio companies named against the investor in S-1 and 424B offering holder tables. 0 usually means no named-holder link was observed. | [optional] 
**HoldingsPositionCount** | Pointer to **NullableInt64** | Current Form 13F positions reported by the managers that resolve to this investor, summed over its manager keys. 0 when none resolves. | [optional] 
**OrgId** | Pointer to **NullableString** | Legal-entity level: the investor&#39;s organization record where the link resolves (null otherwise). An advfirm:&lt;CRD&gt; shell resolves to no legal-entity record and no company. | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: the investor&#39;s company (null when unresolved). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the investor&#39;s company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the investor&#39;s company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the investor&#39;s group. | [optional] 

## Methods

### NewInvestorsRow

`func NewInvestorsRow() *InvestorsRow`

NewInvestorsRow instantiates a new InvestorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestorsRowWithDefaults

`func NewInvestorsRowWithDefaults() *InvestorsRow`

NewInvestorsRowWithDefaults instantiates a new InvestorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestorId

`func (o *InvestorsRow) GetInvestorId() string`

GetInvestorId returns the InvestorId field if non-nil, zero value otherwise.

### GetInvestorIdOk

`func (o *InvestorsRow) GetInvestorIdOk() (*string, bool)`

GetInvestorIdOk returns a tuple with the InvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorId

`func (o *InvestorsRow) SetInvestorId(v string)`

SetInvestorId sets InvestorId field to given value.

### HasInvestorId

`func (o *InvestorsRow) HasInvestorId() bool`

HasInvestorId returns a boolean if a field has been set.

### SetInvestorIdNil

`func (o *InvestorsRow) SetInvestorIdNil(b bool)`

 SetInvestorIdNil sets the value for InvestorId to be an explicit nil

### UnsetInvestorId
`func (o *InvestorsRow) UnsetInvestorId()`

UnsetInvestorId ensures that no value is present for InvestorId, not even an explicit nil
### GetLegalName

`func (o *InvestorsRow) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *InvestorsRow) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *InvestorsRow) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *InvestorsRow) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *InvestorsRow) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *InvestorsRow) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetInvestorType

`func (o *InvestorsRow) GetInvestorType() string`

GetInvestorType returns the InvestorType field if non-nil, zero value otherwise.

### GetInvestorTypeOk

`func (o *InvestorsRow) GetInvestorTypeOk() (*string, bool)`

GetInvestorTypeOk returns a tuple with the InvestorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorType

`func (o *InvestorsRow) SetInvestorType(v string)`

SetInvestorType sets InvestorType field to given value.

### HasInvestorType

`func (o *InvestorsRow) HasInvestorType() bool`

HasInvestorType returns a boolean if a field has been set.

### SetInvestorTypeNil

`func (o *InvestorsRow) SetInvestorTypeNil(b bool)`

 SetInvestorTypeNil sets the value for InvestorType to be an explicit nil

### UnsetInvestorType
`func (o *InvestorsRow) UnsetInvestorType()`

UnsetInvestorType ensures that no value is present for InvestorType, not even an explicit nil
### GetSource

`func (o *InvestorsRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InvestorsRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InvestorsRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InvestorsRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *InvestorsRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *InvestorsRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCountry

`func (o *InvestorsRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InvestorsRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InvestorsRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InvestorsRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *InvestorsRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *InvestorsRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetRaumUsd

`func (o *InvestorsRow) GetRaumUsd() float32`

GetRaumUsd returns the RaumUsd field if non-nil, zero value otherwise.

### GetRaumUsdOk

`func (o *InvestorsRow) GetRaumUsdOk() (*float32, bool)`

GetRaumUsdOk returns a tuple with the RaumUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaumUsd

`func (o *InvestorsRow) SetRaumUsd(v float32)`

SetRaumUsd sets RaumUsd field to given value.

### HasRaumUsd

`func (o *InvestorsRow) HasRaumUsd() bool`

HasRaumUsd returns a boolean if a field has been set.

### SetRaumUsdNil

`func (o *InvestorsRow) SetRaumUsdNil(b bool)`

 SetRaumUsdNil sets the value for RaumUsd to be an explicit nil

### UnsetRaumUsd
`func (o *InvestorsRow) UnsetRaumUsd()`

UnsetRaumUsd ensures that no value is present for RaumUsd, not even an explicit nil
### GetRaumBand

`func (o *InvestorsRow) GetRaumBand() string`

GetRaumBand returns the RaumBand field if non-nil, zero value otherwise.

### GetRaumBandOk

`func (o *InvestorsRow) GetRaumBandOk() (*string, bool)`

GetRaumBandOk returns a tuple with the RaumBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaumBand

`func (o *InvestorsRow) SetRaumBand(v string)`

SetRaumBand sets RaumBand field to given value.

### HasRaumBand

`func (o *InvestorsRow) HasRaumBand() bool`

HasRaumBand returns a boolean if a field has been set.

### SetRaumBandNil

`func (o *InvestorsRow) SetRaumBandNil(b bool)`

 SetRaumBandNil sets the value for RaumBand to be an explicit nil

### UnsetRaumBand
`func (o *InvestorsRow) UnsetRaumBand()`

UnsetRaumBand ensures that no value is present for RaumBand, not even an explicit nil
### GetIsEra

`func (o *InvestorsRow) GetIsEra() bool`

GetIsEra returns the IsEra field if non-nil, zero value otherwise.

### GetIsEraOk

`func (o *InvestorsRow) GetIsEraOk() (*bool, bool)`

GetIsEraOk returns a tuple with the IsEra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEra

`func (o *InvestorsRow) SetIsEra(v bool)`

SetIsEra sets IsEra field to given value.

### HasIsEra

`func (o *InvestorsRow) HasIsEra() bool`

HasIsEra returns a boolean if a field has been set.

### SetIsEraNil

`func (o *InvestorsRow) SetIsEraNil(b bool)`

 SetIsEraNil sets the value for IsEra to be an explicit nil

### UnsetIsEra
`func (o *InvestorsRow) UnsetIsEra()`

UnsetIsEra ensures that no value is present for IsEra, not even an explicit nil
### GetIsLifeSciencesFocused

`func (o *InvestorsRow) GetIsLifeSciencesFocused() bool`

GetIsLifeSciencesFocused returns the IsLifeSciencesFocused field if non-nil, zero value otherwise.

### GetIsLifeSciencesFocusedOk

`func (o *InvestorsRow) GetIsLifeSciencesFocusedOk() (*bool, bool)`

GetIsLifeSciencesFocusedOk returns a tuple with the IsLifeSciencesFocused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLifeSciencesFocused

`func (o *InvestorsRow) SetIsLifeSciencesFocused(v bool)`

SetIsLifeSciencesFocused sets IsLifeSciencesFocused field to given value.

### HasIsLifeSciencesFocused

`func (o *InvestorsRow) HasIsLifeSciencesFocused() bool`

HasIsLifeSciencesFocused returns a boolean if a field has been set.

### SetIsLifeSciencesFocusedNil

`func (o *InvestorsRow) SetIsLifeSciencesFocusedNil(b bool)`

 SetIsLifeSciencesFocusedNil sets the value for IsLifeSciencesFocused to be an explicit nil

### UnsetIsLifeSciencesFocused
`func (o *InvestorsRow) UnsetIsLifeSciencesFocused()`

UnsetIsLifeSciencesFocused ensures that no value is present for IsLifeSciencesFocused, not even an explicit nil
### GetIsLifeSciencesFocusedBasis

`func (o *InvestorsRow) GetIsLifeSciencesFocusedBasis() string`

GetIsLifeSciencesFocusedBasis returns the IsLifeSciencesFocusedBasis field if non-nil, zero value otherwise.

### GetIsLifeSciencesFocusedBasisOk

`func (o *InvestorsRow) GetIsLifeSciencesFocusedBasisOk() (*string, bool)`

GetIsLifeSciencesFocusedBasisOk returns a tuple with the IsLifeSciencesFocusedBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLifeSciencesFocusedBasis

`func (o *InvestorsRow) SetIsLifeSciencesFocusedBasis(v string)`

SetIsLifeSciencesFocusedBasis sets IsLifeSciencesFocusedBasis field to given value.

### HasIsLifeSciencesFocusedBasis

`func (o *InvestorsRow) HasIsLifeSciencesFocusedBasis() bool`

HasIsLifeSciencesFocusedBasis returns a boolean if a field has been set.

### SetIsLifeSciencesFocusedBasisNil

`func (o *InvestorsRow) SetIsLifeSciencesFocusedBasisNil(b bool)`

 SetIsLifeSciencesFocusedBasisNil sets the value for IsLifeSciencesFocusedBasis to be an explicit nil

### UnsetIsLifeSciencesFocusedBasis
`func (o *InvestorsRow) UnsetIsLifeSciencesFocusedBasis()`

UnsetIsLifeSciencesFocusedBasis ensures that no value is present for IsLifeSciencesFocusedBasis, not even an explicit nil
### GetPrivateFundCount

`func (o *InvestorsRow) GetPrivateFundCount() int32`

GetPrivateFundCount returns the PrivateFundCount field if non-nil, zero value otherwise.

### GetPrivateFundCountOk

`func (o *InvestorsRow) GetPrivateFundCountOk() (*int32, bool)`

GetPrivateFundCountOk returns a tuple with the PrivateFundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateFundCount

`func (o *InvestorsRow) SetPrivateFundCount(v int32)`

SetPrivateFundCount sets PrivateFundCount field to given value.

### HasPrivateFundCount

`func (o *InvestorsRow) HasPrivateFundCount() bool`

HasPrivateFundCount returns a boolean if a field has been set.

### SetPrivateFundCountNil

`func (o *InvestorsRow) SetPrivateFundCountNil(b bool)`

 SetPrivateFundCountNil sets the value for PrivateFundCount to be an explicit nil

### UnsetPrivateFundCount
`func (o *InvestorsRow) UnsetPrivateFundCount()`

UnsetPrivateFundCount ensures that no value is present for PrivateFundCount, not even an explicit nil
### GetTotalPrivateFundGav

`func (o *InvestorsRow) GetTotalPrivateFundGav() float32`

GetTotalPrivateFundGav returns the TotalPrivateFundGav field if non-nil, zero value otherwise.

### GetTotalPrivateFundGavOk

`func (o *InvestorsRow) GetTotalPrivateFundGavOk() (*float32, bool)`

GetTotalPrivateFundGavOk returns a tuple with the TotalPrivateFundGav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrivateFundGav

`func (o *InvestorsRow) SetTotalPrivateFundGav(v float32)`

SetTotalPrivateFundGav sets TotalPrivateFundGav field to given value.

### HasTotalPrivateFundGav

`func (o *InvestorsRow) HasTotalPrivateFundGav() bool`

HasTotalPrivateFundGav returns a boolean if a field has been set.

### SetTotalPrivateFundGavNil

`func (o *InvestorsRow) SetTotalPrivateFundGavNil(b bool)`

 SetTotalPrivateFundGavNil sets the value for TotalPrivateFundGav to be an explicit nil

### UnsetTotalPrivateFundGav
`func (o *InvestorsRow) UnsetTotalPrivateFundGav()`

UnsetTotalPrivateFundGav ensures that no value is present for TotalPrivateFundGav, not even an explicit nil
### GetPortfolioCoCount

`func (o *InvestorsRow) GetPortfolioCoCount() int64`

GetPortfolioCoCount returns the PortfolioCoCount field if non-nil, zero value otherwise.

### GetPortfolioCoCountOk

`func (o *InvestorsRow) GetPortfolioCoCountOk() (*int64, bool)`

GetPortfolioCoCountOk returns a tuple with the PortfolioCoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolioCoCount

`func (o *InvestorsRow) SetPortfolioCoCount(v int64)`

SetPortfolioCoCount sets PortfolioCoCount field to given value.

### HasPortfolioCoCount

`func (o *InvestorsRow) HasPortfolioCoCount() bool`

HasPortfolioCoCount returns a boolean if a field has been set.

### SetPortfolioCoCountNil

`func (o *InvestorsRow) SetPortfolioCoCountNil(b bool)`

 SetPortfolioCoCountNil sets the value for PortfolioCoCount to be an explicit nil

### UnsetPortfolioCoCount
`func (o *InvestorsRow) UnsetPortfolioCoCount()`

UnsetPortfolioCoCount ensures that no value is present for PortfolioCoCount, not even an explicit nil
### GetHoldingsPositionCount

`func (o *InvestorsRow) GetHoldingsPositionCount() int64`

GetHoldingsPositionCount returns the HoldingsPositionCount field if non-nil, zero value otherwise.

### GetHoldingsPositionCountOk

`func (o *InvestorsRow) GetHoldingsPositionCountOk() (*int64, bool)`

GetHoldingsPositionCountOk returns a tuple with the HoldingsPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingsPositionCount

`func (o *InvestorsRow) SetHoldingsPositionCount(v int64)`

SetHoldingsPositionCount sets HoldingsPositionCount field to given value.

### HasHoldingsPositionCount

`func (o *InvestorsRow) HasHoldingsPositionCount() bool`

HasHoldingsPositionCount returns a boolean if a field has been set.

### SetHoldingsPositionCountNil

`func (o *InvestorsRow) SetHoldingsPositionCountNil(b bool)`

 SetHoldingsPositionCountNil sets the value for HoldingsPositionCount to be an explicit nil

### UnsetHoldingsPositionCount
`func (o *InvestorsRow) UnsetHoldingsPositionCount()`

UnsetHoldingsPositionCount ensures that no value is present for HoldingsPositionCount, not even an explicit nil
### GetOrgId

`func (o *InvestorsRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InvestorsRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InvestorsRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *InvestorsRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *InvestorsRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *InvestorsRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetCompanyId

`func (o *InvestorsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InvestorsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InvestorsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InvestorsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *InvestorsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *InvestorsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *InvestorsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvestorsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvestorsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InvestorsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *InvestorsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *InvestorsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *InvestorsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InvestorsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InvestorsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InvestorsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *InvestorsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *InvestorsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *InvestorsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *InvestorsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *InvestorsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *InvestorsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *InvestorsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *InvestorsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CompanyFinancialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**ReportingOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record that filed (the SEC filer). | [optional] 
**FiscalYear** | Pointer to **NullableInt32** | The fiscal year the figures belong to. | [optional] 
**IsPublic** | Pointer to **NullableBool** | True when the record has an SEC registration number and is a listed filer (a Form D issuer that is not listed is false). | [optional] 
**EdgarCik** | Pointer to **NullableString** | The SEC registration number (CIK, Central Index Key) of the filer. | [optional] 
**Ticker** | Pointer to **NullableString** | The stock ticker. | [optional] 
**Exchange** | Pointer to **NullableString** | The primary exchange listing. | [optional] 
**PublicFloatUsd** | Pointer to **NullableFloat32** | The latest reported public float, in USD, repeated on every year row of the filer (a proxy for market value, not a live market capitalization). | [optional] 
**PublicFloatFy** | Pointer to **NullableInt32** | The fiscal year of public_float_usd. | [optional] 
**LatestSharesOutstanding** | Pointer to **NullableFloat32** | The latest reported cover-page share count of the filer, repeated on every year row (the denominator of an ownership percentage). | [optional] 
**LatestSharesOutstandingFy** | Pointer to **NullableInt32** | The fiscal year of latest_shares_outstanding. | [optional] 
**LatestFinancialFy** | Pointer to **NullableInt32** | The latest fiscal year the filer reported any line item. | [optional] 
**Revenue** | Pointer to **NullableFloat32** | Total revenue for the fiscal year, in USD (a non-USD filer is converted at the calendar-year average rate; null when no rate exists). | [optional] 
**NetIncome** | Pointer to **NullableFloat32** | Net income for the fiscal year, in USD. | [optional] 
**Rnd** | Pointer to **NullableFloat32** | Research and development expense for the fiscal year, in USD. | [optional] 
**Assets** | Pointer to **NullableFloat32** | Total assets at the fiscal year end, in USD (converted at the year-end rate). | [optional] 
**Equity** | Pointer to **NullableFloat32** | Stockholders&#39; equity at the fiscal year end, in USD. | [optional] 
**PublicFloat** | Pointer to **NullableFloat32** | The public float on the annual report cover (the value of shares held by non-affiliates at the most recent second fiscal quarter end), in USD, anchored to the fiscal year it was filed with. | [optional] 
**Cash** | Pointer to **NullableFloat32** | Cash and cash equivalents at the fiscal year end, in USD. | [optional] 
**ShortTermInvestments** | Pointer to **NullableFloat32** | Short-term investments at the fiscal year end, in USD. | [optional] 
**CashAndInvestments** | Pointer to **NullableFloat32** | Cash plus short-term investments (a missing investments line counts as zero), in USD; null when cash is null. | [optional] 
**OperatingCashFlow** | Pointer to **NullableFloat32** | Net cash from operating activities for the fiscal year, in USD; negative means burning. | [optional] 
**SharesOutstanding** | Pointer to **NullableFloat32** | Shares outstanding on the annual report cover date, a count (never currency-converted); the cover date falls 0 to 3 months after the fiscal year end. | [optional] 
**StockBasedCompensation** | Pointer to **NullableFloat32** | Stock-based compensation expense, in USD. | [optional] 
**OperatingExpenses** | Pointer to **NullableFloat32** | Operating expenses, in USD. | [optional] 
**DividendsPaid** | Pointer to **NullableFloat32** | Dividends paid, in USD, as filed with no sign test (filers differ on whether an outflow is positive or negative). | [optional] 
**EpsBasic** | Pointer to **NullableFloat32** | Basic earnings per share, in USD. | [optional] 
**EpsDiluted** | Pointer to **NullableFloat32** | Diluted earnings per share, in USD. | [optional] 
**TotalLiabilities** | Pointer to **NullableFloat32** | Total liabilities at the fiscal year end, in USD. | [optional] 
**LongTermDebtNoncurrent** | Pointer to **NullableFloat32** | The non-current portion of long-term debt, in USD (null for a filer that reports only a total). | [optional] 
**LongTermDebtTotal** | Pointer to **NullableFloat32** | Long-term debt, current plus non-current, in USD. | [optional] 
**IntangibleAssets** | Pointer to **NullableFloat32** | Intangible assets, in USD. | [optional] 
**PropertyPlantEquipment** | Pointer to **NullableFloat32** | Property, plant and equipment, in USD. | [optional] 
**Goodwill** | Pointer to **NullableFloat32** | Goodwill, in USD. | [optional] 
**Inventory** | Pointer to **NullableFloat32** | Inventory, in USD. | [optional] 
**AccountsPayable** | Pointer to **NullableFloat32** | Accounts payable, in USD. | [optional] 
**AccountsReceivable** | Pointer to **NullableFloat32** | Accounts receivable, in USD. | [optional] 
**WeightedAverageSharesBasic** | Pointer to **NullableFloat32** | Weighted average basic shares, a count. | [optional] 
**WeightedAverageSharesDiluted** | Pointer to **NullableFloat32** | Weighted average diluted shares, a count. | [optional] 
**BurnBasisUsd** | Pointer to **NullableFloat32** | The average operating cash flow over this fiscal year and the two before it (the years reported), in USD; negative means burning. | [optional] 
**BurnBasisYears** | Pointer to **NullableInt32** | How many reported years burn_basis_usd averages over, 1 to 3. | [optional] 
**RunwayMonths** | Pointer to **NullableFloat32** | Derived, not reported: cash_and_investments divided by the monthly burn (burn_basis_usd over 12), in months, rounded to 0.1. Null when the filer generated cash or has no cash figure. | [optional] 
**SecTotalRevenueUsdMillions** | Pointer to **NullableFloat32** | The filing-level total revenue extracted from the 10-K or 20-F, in USD millions (the line the revenue-by-drug chart draws over the per-product stack); null where no total was extracted. | [optional] 

## Methods

### NewCompanyFinancialsRow

`func NewCompanyFinancialsRow() *CompanyFinancialsRow`

NewCompanyFinancialsRow instantiates a new CompanyFinancialsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyFinancialsRowWithDefaults

`func NewCompanyFinancialsRowWithDefaults() *CompanyFinancialsRow`

NewCompanyFinancialsRowWithDefaults instantiates a new CompanyFinancialsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyFinancialsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyFinancialsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyFinancialsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyFinancialsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyFinancialsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyFinancialsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyFinancialsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyFinancialsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyFinancialsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyFinancialsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyFinancialsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyFinancialsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyFinancialsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyFinancialsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyFinancialsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyFinancialsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyFinancialsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyFinancialsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyFinancialsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyFinancialsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyFinancialsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyFinancialsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyFinancialsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyFinancialsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetReportingOrgId

`func (o *CompanyFinancialsRow) GetReportingOrgId() string`

GetReportingOrgId returns the ReportingOrgId field if non-nil, zero value otherwise.

### GetReportingOrgIdOk

`func (o *CompanyFinancialsRow) GetReportingOrgIdOk() (*string, bool)`

GetReportingOrgIdOk returns a tuple with the ReportingOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOrgId

`func (o *CompanyFinancialsRow) SetReportingOrgId(v string)`

SetReportingOrgId sets ReportingOrgId field to given value.

### HasReportingOrgId

`func (o *CompanyFinancialsRow) HasReportingOrgId() bool`

HasReportingOrgId returns a boolean if a field has been set.

### SetReportingOrgIdNil

`func (o *CompanyFinancialsRow) SetReportingOrgIdNil(b bool)`

 SetReportingOrgIdNil sets the value for ReportingOrgId to be an explicit nil

### UnsetReportingOrgId
`func (o *CompanyFinancialsRow) UnsetReportingOrgId()`

UnsetReportingOrgId ensures that no value is present for ReportingOrgId, not even an explicit nil
### GetFiscalYear

`func (o *CompanyFinancialsRow) GetFiscalYear() int32`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *CompanyFinancialsRow) GetFiscalYearOk() (*int32, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *CompanyFinancialsRow) SetFiscalYear(v int32)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *CompanyFinancialsRow) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### SetFiscalYearNil

`func (o *CompanyFinancialsRow) SetFiscalYearNil(b bool)`

 SetFiscalYearNil sets the value for FiscalYear to be an explicit nil

### UnsetFiscalYear
`func (o *CompanyFinancialsRow) UnsetFiscalYear()`

UnsetFiscalYear ensures that no value is present for FiscalYear, not even an explicit nil
### GetIsPublic

`func (o *CompanyFinancialsRow) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CompanyFinancialsRow) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CompanyFinancialsRow) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *CompanyFinancialsRow) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### SetIsPublicNil

`func (o *CompanyFinancialsRow) SetIsPublicNil(b bool)`

 SetIsPublicNil sets the value for IsPublic to be an explicit nil

### UnsetIsPublic
`func (o *CompanyFinancialsRow) UnsetIsPublic()`

UnsetIsPublic ensures that no value is present for IsPublic, not even an explicit nil
### GetEdgarCik

`func (o *CompanyFinancialsRow) GetEdgarCik() string`

GetEdgarCik returns the EdgarCik field if non-nil, zero value otherwise.

### GetEdgarCikOk

`func (o *CompanyFinancialsRow) GetEdgarCikOk() (*string, bool)`

GetEdgarCikOk returns a tuple with the EdgarCik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgarCik

`func (o *CompanyFinancialsRow) SetEdgarCik(v string)`

SetEdgarCik sets EdgarCik field to given value.

### HasEdgarCik

`func (o *CompanyFinancialsRow) HasEdgarCik() bool`

HasEdgarCik returns a boolean if a field has been set.

### SetEdgarCikNil

`func (o *CompanyFinancialsRow) SetEdgarCikNil(b bool)`

 SetEdgarCikNil sets the value for EdgarCik to be an explicit nil

### UnsetEdgarCik
`func (o *CompanyFinancialsRow) UnsetEdgarCik()`

UnsetEdgarCik ensures that no value is present for EdgarCik, not even an explicit nil
### GetTicker

`func (o *CompanyFinancialsRow) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *CompanyFinancialsRow) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *CompanyFinancialsRow) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *CompanyFinancialsRow) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### SetTickerNil

`func (o *CompanyFinancialsRow) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *CompanyFinancialsRow) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil
### GetExchange

`func (o *CompanyFinancialsRow) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *CompanyFinancialsRow) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *CompanyFinancialsRow) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *CompanyFinancialsRow) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### SetExchangeNil

`func (o *CompanyFinancialsRow) SetExchangeNil(b bool)`

 SetExchangeNil sets the value for Exchange to be an explicit nil

### UnsetExchange
`func (o *CompanyFinancialsRow) UnsetExchange()`

UnsetExchange ensures that no value is present for Exchange, not even an explicit nil
### GetPublicFloatUsd

`func (o *CompanyFinancialsRow) GetPublicFloatUsd() float32`

GetPublicFloatUsd returns the PublicFloatUsd field if non-nil, zero value otherwise.

### GetPublicFloatUsdOk

`func (o *CompanyFinancialsRow) GetPublicFloatUsdOk() (*float32, bool)`

GetPublicFloatUsdOk returns a tuple with the PublicFloatUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFloatUsd

`func (o *CompanyFinancialsRow) SetPublicFloatUsd(v float32)`

SetPublicFloatUsd sets PublicFloatUsd field to given value.

### HasPublicFloatUsd

`func (o *CompanyFinancialsRow) HasPublicFloatUsd() bool`

HasPublicFloatUsd returns a boolean if a field has been set.

### SetPublicFloatUsdNil

`func (o *CompanyFinancialsRow) SetPublicFloatUsdNil(b bool)`

 SetPublicFloatUsdNil sets the value for PublicFloatUsd to be an explicit nil

### UnsetPublicFloatUsd
`func (o *CompanyFinancialsRow) UnsetPublicFloatUsd()`

UnsetPublicFloatUsd ensures that no value is present for PublicFloatUsd, not even an explicit nil
### GetPublicFloatFy

`func (o *CompanyFinancialsRow) GetPublicFloatFy() int32`

GetPublicFloatFy returns the PublicFloatFy field if non-nil, zero value otherwise.

### GetPublicFloatFyOk

`func (o *CompanyFinancialsRow) GetPublicFloatFyOk() (*int32, bool)`

GetPublicFloatFyOk returns a tuple with the PublicFloatFy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFloatFy

`func (o *CompanyFinancialsRow) SetPublicFloatFy(v int32)`

SetPublicFloatFy sets PublicFloatFy field to given value.

### HasPublicFloatFy

`func (o *CompanyFinancialsRow) HasPublicFloatFy() bool`

HasPublicFloatFy returns a boolean if a field has been set.

### SetPublicFloatFyNil

`func (o *CompanyFinancialsRow) SetPublicFloatFyNil(b bool)`

 SetPublicFloatFyNil sets the value for PublicFloatFy to be an explicit nil

### UnsetPublicFloatFy
`func (o *CompanyFinancialsRow) UnsetPublicFloatFy()`

UnsetPublicFloatFy ensures that no value is present for PublicFloatFy, not even an explicit nil
### GetLatestSharesOutstanding

`func (o *CompanyFinancialsRow) GetLatestSharesOutstanding() float32`

GetLatestSharesOutstanding returns the LatestSharesOutstanding field if non-nil, zero value otherwise.

### GetLatestSharesOutstandingOk

`func (o *CompanyFinancialsRow) GetLatestSharesOutstandingOk() (*float32, bool)`

GetLatestSharesOutstandingOk returns a tuple with the LatestSharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSharesOutstanding

`func (o *CompanyFinancialsRow) SetLatestSharesOutstanding(v float32)`

SetLatestSharesOutstanding sets LatestSharesOutstanding field to given value.

### HasLatestSharesOutstanding

`func (o *CompanyFinancialsRow) HasLatestSharesOutstanding() bool`

HasLatestSharesOutstanding returns a boolean if a field has been set.

### SetLatestSharesOutstandingNil

`func (o *CompanyFinancialsRow) SetLatestSharesOutstandingNil(b bool)`

 SetLatestSharesOutstandingNil sets the value for LatestSharesOutstanding to be an explicit nil

### UnsetLatestSharesOutstanding
`func (o *CompanyFinancialsRow) UnsetLatestSharesOutstanding()`

UnsetLatestSharesOutstanding ensures that no value is present for LatestSharesOutstanding, not even an explicit nil
### GetLatestSharesOutstandingFy

`func (o *CompanyFinancialsRow) GetLatestSharesOutstandingFy() int32`

GetLatestSharesOutstandingFy returns the LatestSharesOutstandingFy field if non-nil, zero value otherwise.

### GetLatestSharesOutstandingFyOk

`func (o *CompanyFinancialsRow) GetLatestSharesOutstandingFyOk() (*int32, bool)`

GetLatestSharesOutstandingFyOk returns a tuple with the LatestSharesOutstandingFy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSharesOutstandingFy

`func (o *CompanyFinancialsRow) SetLatestSharesOutstandingFy(v int32)`

SetLatestSharesOutstandingFy sets LatestSharesOutstandingFy field to given value.

### HasLatestSharesOutstandingFy

`func (o *CompanyFinancialsRow) HasLatestSharesOutstandingFy() bool`

HasLatestSharesOutstandingFy returns a boolean if a field has been set.

### SetLatestSharesOutstandingFyNil

`func (o *CompanyFinancialsRow) SetLatestSharesOutstandingFyNil(b bool)`

 SetLatestSharesOutstandingFyNil sets the value for LatestSharesOutstandingFy to be an explicit nil

### UnsetLatestSharesOutstandingFy
`func (o *CompanyFinancialsRow) UnsetLatestSharesOutstandingFy()`

UnsetLatestSharesOutstandingFy ensures that no value is present for LatestSharesOutstandingFy, not even an explicit nil
### GetLatestFinancialFy

`func (o *CompanyFinancialsRow) GetLatestFinancialFy() int32`

GetLatestFinancialFy returns the LatestFinancialFy field if non-nil, zero value otherwise.

### GetLatestFinancialFyOk

`func (o *CompanyFinancialsRow) GetLatestFinancialFyOk() (*int32, bool)`

GetLatestFinancialFyOk returns a tuple with the LatestFinancialFy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFinancialFy

`func (o *CompanyFinancialsRow) SetLatestFinancialFy(v int32)`

SetLatestFinancialFy sets LatestFinancialFy field to given value.

### HasLatestFinancialFy

`func (o *CompanyFinancialsRow) HasLatestFinancialFy() bool`

HasLatestFinancialFy returns a boolean if a field has been set.

### SetLatestFinancialFyNil

`func (o *CompanyFinancialsRow) SetLatestFinancialFyNil(b bool)`

 SetLatestFinancialFyNil sets the value for LatestFinancialFy to be an explicit nil

### UnsetLatestFinancialFy
`func (o *CompanyFinancialsRow) UnsetLatestFinancialFy()`

UnsetLatestFinancialFy ensures that no value is present for LatestFinancialFy, not even an explicit nil
### GetRevenue

`func (o *CompanyFinancialsRow) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CompanyFinancialsRow) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CompanyFinancialsRow) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CompanyFinancialsRow) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *CompanyFinancialsRow) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *CompanyFinancialsRow) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
### GetNetIncome

`func (o *CompanyFinancialsRow) GetNetIncome() float32`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *CompanyFinancialsRow) GetNetIncomeOk() (*float32, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *CompanyFinancialsRow) SetNetIncome(v float32)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *CompanyFinancialsRow) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### SetNetIncomeNil

`func (o *CompanyFinancialsRow) SetNetIncomeNil(b bool)`

 SetNetIncomeNil sets the value for NetIncome to be an explicit nil

### UnsetNetIncome
`func (o *CompanyFinancialsRow) UnsetNetIncome()`

UnsetNetIncome ensures that no value is present for NetIncome, not even an explicit nil
### GetRnd

`func (o *CompanyFinancialsRow) GetRnd() float32`

GetRnd returns the Rnd field if non-nil, zero value otherwise.

### GetRndOk

`func (o *CompanyFinancialsRow) GetRndOk() (*float32, bool)`

GetRndOk returns a tuple with the Rnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnd

`func (o *CompanyFinancialsRow) SetRnd(v float32)`

SetRnd sets Rnd field to given value.

### HasRnd

`func (o *CompanyFinancialsRow) HasRnd() bool`

HasRnd returns a boolean if a field has been set.

### SetRndNil

`func (o *CompanyFinancialsRow) SetRndNil(b bool)`

 SetRndNil sets the value for Rnd to be an explicit nil

### UnsetRnd
`func (o *CompanyFinancialsRow) UnsetRnd()`

UnsetRnd ensures that no value is present for Rnd, not even an explicit nil
### GetAssets

`func (o *CompanyFinancialsRow) GetAssets() float32`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CompanyFinancialsRow) GetAssetsOk() (*float32, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CompanyFinancialsRow) SetAssets(v float32)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CompanyFinancialsRow) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### SetAssetsNil

`func (o *CompanyFinancialsRow) SetAssetsNil(b bool)`

 SetAssetsNil sets the value for Assets to be an explicit nil

### UnsetAssets
`func (o *CompanyFinancialsRow) UnsetAssets()`

UnsetAssets ensures that no value is present for Assets, not even an explicit nil
### GetEquity

`func (o *CompanyFinancialsRow) GetEquity() float32`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *CompanyFinancialsRow) GetEquityOk() (*float32, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *CompanyFinancialsRow) SetEquity(v float32)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *CompanyFinancialsRow) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### SetEquityNil

`func (o *CompanyFinancialsRow) SetEquityNil(b bool)`

 SetEquityNil sets the value for Equity to be an explicit nil

### UnsetEquity
`func (o *CompanyFinancialsRow) UnsetEquity()`

UnsetEquity ensures that no value is present for Equity, not even an explicit nil
### GetPublicFloat

`func (o *CompanyFinancialsRow) GetPublicFloat() float32`

GetPublicFloat returns the PublicFloat field if non-nil, zero value otherwise.

### GetPublicFloatOk

`func (o *CompanyFinancialsRow) GetPublicFloatOk() (*float32, bool)`

GetPublicFloatOk returns a tuple with the PublicFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFloat

`func (o *CompanyFinancialsRow) SetPublicFloat(v float32)`

SetPublicFloat sets PublicFloat field to given value.

### HasPublicFloat

`func (o *CompanyFinancialsRow) HasPublicFloat() bool`

HasPublicFloat returns a boolean if a field has been set.

### SetPublicFloatNil

`func (o *CompanyFinancialsRow) SetPublicFloatNil(b bool)`

 SetPublicFloatNil sets the value for PublicFloat to be an explicit nil

### UnsetPublicFloat
`func (o *CompanyFinancialsRow) UnsetPublicFloat()`

UnsetPublicFloat ensures that no value is present for PublicFloat, not even an explicit nil
### GetCash

`func (o *CompanyFinancialsRow) GetCash() float32`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *CompanyFinancialsRow) GetCashOk() (*float32, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *CompanyFinancialsRow) SetCash(v float32)`

SetCash sets Cash field to given value.

### HasCash

`func (o *CompanyFinancialsRow) HasCash() bool`

HasCash returns a boolean if a field has been set.

### SetCashNil

`func (o *CompanyFinancialsRow) SetCashNil(b bool)`

 SetCashNil sets the value for Cash to be an explicit nil

### UnsetCash
`func (o *CompanyFinancialsRow) UnsetCash()`

UnsetCash ensures that no value is present for Cash, not even an explicit nil
### GetShortTermInvestments

`func (o *CompanyFinancialsRow) GetShortTermInvestments() float32`

GetShortTermInvestments returns the ShortTermInvestments field if non-nil, zero value otherwise.

### GetShortTermInvestmentsOk

`func (o *CompanyFinancialsRow) GetShortTermInvestmentsOk() (*float32, bool)`

GetShortTermInvestmentsOk returns a tuple with the ShortTermInvestments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortTermInvestments

`func (o *CompanyFinancialsRow) SetShortTermInvestments(v float32)`

SetShortTermInvestments sets ShortTermInvestments field to given value.

### HasShortTermInvestments

`func (o *CompanyFinancialsRow) HasShortTermInvestments() bool`

HasShortTermInvestments returns a boolean if a field has been set.

### SetShortTermInvestmentsNil

`func (o *CompanyFinancialsRow) SetShortTermInvestmentsNil(b bool)`

 SetShortTermInvestmentsNil sets the value for ShortTermInvestments to be an explicit nil

### UnsetShortTermInvestments
`func (o *CompanyFinancialsRow) UnsetShortTermInvestments()`

UnsetShortTermInvestments ensures that no value is present for ShortTermInvestments, not even an explicit nil
### GetCashAndInvestments

`func (o *CompanyFinancialsRow) GetCashAndInvestments() float32`

GetCashAndInvestments returns the CashAndInvestments field if non-nil, zero value otherwise.

### GetCashAndInvestmentsOk

`func (o *CompanyFinancialsRow) GetCashAndInvestmentsOk() (*float32, bool)`

GetCashAndInvestmentsOk returns a tuple with the CashAndInvestments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAndInvestments

`func (o *CompanyFinancialsRow) SetCashAndInvestments(v float32)`

SetCashAndInvestments sets CashAndInvestments field to given value.

### HasCashAndInvestments

`func (o *CompanyFinancialsRow) HasCashAndInvestments() bool`

HasCashAndInvestments returns a boolean if a field has been set.

### SetCashAndInvestmentsNil

`func (o *CompanyFinancialsRow) SetCashAndInvestmentsNil(b bool)`

 SetCashAndInvestmentsNil sets the value for CashAndInvestments to be an explicit nil

### UnsetCashAndInvestments
`func (o *CompanyFinancialsRow) UnsetCashAndInvestments()`

UnsetCashAndInvestments ensures that no value is present for CashAndInvestments, not even an explicit nil
### GetOperatingCashFlow

`func (o *CompanyFinancialsRow) GetOperatingCashFlow() float32`

GetOperatingCashFlow returns the OperatingCashFlow field if non-nil, zero value otherwise.

### GetOperatingCashFlowOk

`func (o *CompanyFinancialsRow) GetOperatingCashFlowOk() (*float32, bool)`

GetOperatingCashFlowOk returns a tuple with the OperatingCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingCashFlow

`func (o *CompanyFinancialsRow) SetOperatingCashFlow(v float32)`

SetOperatingCashFlow sets OperatingCashFlow field to given value.

### HasOperatingCashFlow

`func (o *CompanyFinancialsRow) HasOperatingCashFlow() bool`

HasOperatingCashFlow returns a boolean if a field has been set.

### SetOperatingCashFlowNil

`func (o *CompanyFinancialsRow) SetOperatingCashFlowNil(b bool)`

 SetOperatingCashFlowNil sets the value for OperatingCashFlow to be an explicit nil

### UnsetOperatingCashFlow
`func (o *CompanyFinancialsRow) UnsetOperatingCashFlow()`

UnsetOperatingCashFlow ensures that no value is present for OperatingCashFlow, not even an explicit nil
### GetSharesOutstanding

`func (o *CompanyFinancialsRow) GetSharesOutstanding() float32`

GetSharesOutstanding returns the SharesOutstanding field if non-nil, zero value otherwise.

### GetSharesOutstandingOk

`func (o *CompanyFinancialsRow) GetSharesOutstandingOk() (*float32, bool)`

GetSharesOutstandingOk returns a tuple with the SharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstanding

`func (o *CompanyFinancialsRow) SetSharesOutstanding(v float32)`

SetSharesOutstanding sets SharesOutstanding field to given value.

### HasSharesOutstanding

`func (o *CompanyFinancialsRow) HasSharesOutstanding() bool`

HasSharesOutstanding returns a boolean if a field has been set.

### SetSharesOutstandingNil

`func (o *CompanyFinancialsRow) SetSharesOutstandingNil(b bool)`

 SetSharesOutstandingNil sets the value for SharesOutstanding to be an explicit nil

### UnsetSharesOutstanding
`func (o *CompanyFinancialsRow) UnsetSharesOutstanding()`

UnsetSharesOutstanding ensures that no value is present for SharesOutstanding, not even an explicit nil
### GetStockBasedCompensation

`func (o *CompanyFinancialsRow) GetStockBasedCompensation() float32`

GetStockBasedCompensation returns the StockBasedCompensation field if non-nil, zero value otherwise.

### GetStockBasedCompensationOk

`func (o *CompanyFinancialsRow) GetStockBasedCompensationOk() (*float32, bool)`

GetStockBasedCompensationOk returns a tuple with the StockBasedCompensation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockBasedCompensation

`func (o *CompanyFinancialsRow) SetStockBasedCompensation(v float32)`

SetStockBasedCompensation sets StockBasedCompensation field to given value.

### HasStockBasedCompensation

`func (o *CompanyFinancialsRow) HasStockBasedCompensation() bool`

HasStockBasedCompensation returns a boolean if a field has been set.

### SetStockBasedCompensationNil

`func (o *CompanyFinancialsRow) SetStockBasedCompensationNil(b bool)`

 SetStockBasedCompensationNil sets the value for StockBasedCompensation to be an explicit nil

### UnsetStockBasedCompensation
`func (o *CompanyFinancialsRow) UnsetStockBasedCompensation()`

UnsetStockBasedCompensation ensures that no value is present for StockBasedCompensation, not even an explicit nil
### GetOperatingExpenses

`func (o *CompanyFinancialsRow) GetOperatingExpenses() float32`

GetOperatingExpenses returns the OperatingExpenses field if non-nil, zero value otherwise.

### GetOperatingExpensesOk

`func (o *CompanyFinancialsRow) GetOperatingExpensesOk() (*float32, bool)`

GetOperatingExpensesOk returns a tuple with the OperatingExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingExpenses

`func (o *CompanyFinancialsRow) SetOperatingExpenses(v float32)`

SetOperatingExpenses sets OperatingExpenses field to given value.

### HasOperatingExpenses

`func (o *CompanyFinancialsRow) HasOperatingExpenses() bool`

HasOperatingExpenses returns a boolean if a field has been set.

### SetOperatingExpensesNil

`func (o *CompanyFinancialsRow) SetOperatingExpensesNil(b bool)`

 SetOperatingExpensesNil sets the value for OperatingExpenses to be an explicit nil

### UnsetOperatingExpenses
`func (o *CompanyFinancialsRow) UnsetOperatingExpenses()`

UnsetOperatingExpenses ensures that no value is present for OperatingExpenses, not even an explicit nil
### GetDividendsPaid

`func (o *CompanyFinancialsRow) GetDividendsPaid() float32`

GetDividendsPaid returns the DividendsPaid field if non-nil, zero value otherwise.

### GetDividendsPaidOk

`func (o *CompanyFinancialsRow) GetDividendsPaidOk() (*float32, bool)`

GetDividendsPaidOk returns a tuple with the DividendsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendsPaid

`func (o *CompanyFinancialsRow) SetDividendsPaid(v float32)`

SetDividendsPaid sets DividendsPaid field to given value.

### HasDividendsPaid

`func (o *CompanyFinancialsRow) HasDividendsPaid() bool`

HasDividendsPaid returns a boolean if a field has been set.

### SetDividendsPaidNil

`func (o *CompanyFinancialsRow) SetDividendsPaidNil(b bool)`

 SetDividendsPaidNil sets the value for DividendsPaid to be an explicit nil

### UnsetDividendsPaid
`func (o *CompanyFinancialsRow) UnsetDividendsPaid()`

UnsetDividendsPaid ensures that no value is present for DividendsPaid, not even an explicit nil
### GetEpsBasic

`func (o *CompanyFinancialsRow) GetEpsBasic() float32`

GetEpsBasic returns the EpsBasic field if non-nil, zero value otherwise.

### GetEpsBasicOk

`func (o *CompanyFinancialsRow) GetEpsBasicOk() (*float32, bool)`

GetEpsBasicOk returns a tuple with the EpsBasic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBasic

`func (o *CompanyFinancialsRow) SetEpsBasic(v float32)`

SetEpsBasic sets EpsBasic field to given value.

### HasEpsBasic

`func (o *CompanyFinancialsRow) HasEpsBasic() bool`

HasEpsBasic returns a boolean if a field has been set.

### SetEpsBasicNil

`func (o *CompanyFinancialsRow) SetEpsBasicNil(b bool)`

 SetEpsBasicNil sets the value for EpsBasic to be an explicit nil

### UnsetEpsBasic
`func (o *CompanyFinancialsRow) UnsetEpsBasic()`

UnsetEpsBasic ensures that no value is present for EpsBasic, not even an explicit nil
### GetEpsDiluted

`func (o *CompanyFinancialsRow) GetEpsDiluted() float32`

GetEpsDiluted returns the EpsDiluted field if non-nil, zero value otherwise.

### GetEpsDilutedOk

`func (o *CompanyFinancialsRow) GetEpsDilutedOk() (*float32, bool)`

GetEpsDilutedOk returns a tuple with the EpsDiluted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsDiluted

`func (o *CompanyFinancialsRow) SetEpsDiluted(v float32)`

SetEpsDiluted sets EpsDiluted field to given value.

### HasEpsDiluted

`func (o *CompanyFinancialsRow) HasEpsDiluted() bool`

HasEpsDiluted returns a boolean if a field has been set.

### SetEpsDilutedNil

`func (o *CompanyFinancialsRow) SetEpsDilutedNil(b bool)`

 SetEpsDilutedNil sets the value for EpsDiluted to be an explicit nil

### UnsetEpsDiluted
`func (o *CompanyFinancialsRow) UnsetEpsDiluted()`

UnsetEpsDiluted ensures that no value is present for EpsDiluted, not even an explicit nil
### GetTotalLiabilities

`func (o *CompanyFinancialsRow) GetTotalLiabilities() float32`

GetTotalLiabilities returns the TotalLiabilities field if non-nil, zero value otherwise.

### GetTotalLiabilitiesOk

`func (o *CompanyFinancialsRow) GetTotalLiabilitiesOk() (*float32, bool)`

GetTotalLiabilitiesOk returns a tuple with the TotalLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilities

`func (o *CompanyFinancialsRow) SetTotalLiabilities(v float32)`

SetTotalLiabilities sets TotalLiabilities field to given value.

### HasTotalLiabilities

`func (o *CompanyFinancialsRow) HasTotalLiabilities() bool`

HasTotalLiabilities returns a boolean if a field has been set.

### SetTotalLiabilitiesNil

`func (o *CompanyFinancialsRow) SetTotalLiabilitiesNil(b bool)`

 SetTotalLiabilitiesNil sets the value for TotalLiabilities to be an explicit nil

### UnsetTotalLiabilities
`func (o *CompanyFinancialsRow) UnsetTotalLiabilities()`

UnsetTotalLiabilities ensures that no value is present for TotalLiabilities, not even an explicit nil
### GetLongTermDebtNoncurrent

`func (o *CompanyFinancialsRow) GetLongTermDebtNoncurrent() float32`

GetLongTermDebtNoncurrent returns the LongTermDebtNoncurrent field if non-nil, zero value otherwise.

### GetLongTermDebtNoncurrentOk

`func (o *CompanyFinancialsRow) GetLongTermDebtNoncurrentOk() (*float32, bool)`

GetLongTermDebtNoncurrentOk returns a tuple with the LongTermDebtNoncurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermDebtNoncurrent

`func (o *CompanyFinancialsRow) SetLongTermDebtNoncurrent(v float32)`

SetLongTermDebtNoncurrent sets LongTermDebtNoncurrent field to given value.

### HasLongTermDebtNoncurrent

`func (o *CompanyFinancialsRow) HasLongTermDebtNoncurrent() bool`

HasLongTermDebtNoncurrent returns a boolean if a field has been set.

### SetLongTermDebtNoncurrentNil

`func (o *CompanyFinancialsRow) SetLongTermDebtNoncurrentNil(b bool)`

 SetLongTermDebtNoncurrentNil sets the value for LongTermDebtNoncurrent to be an explicit nil

### UnsetLongTermDebtNoncurrent
`func (o *CompanyFinancialsRow) UnsetLongTermDebtNoncurrent()`

UnsetLongTermDebtNoncurrent ensures that no value is present for LongTermDebtNoncurrent, not even an explicit nil
### GetLongTermDebtTotal

`func (o *CompanyFinancialsRow) GetLongTermDebtTotal() float32`

GetLongTermDebtTotal returns the LongTermDebtTotal field if non-nil, zero value otherwise.

### GetLongTermDebtTotalOk

`func (o *CompanyFinancialsRow) GetLongTermDebtTotalOk() (*float32, bool)`

GetLongTermDebtTotalOk returns a tuple with the LongTermDebtTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermDebtTotal

`func (o *CompanyFinancialsRow) SetLongTermDebtTotal(v float32)`

SetLongTermDebtTotal sets LongTermDebtTotal field to given value.

### HasLongTermDebtTotal

`func (o *CompanyFinancialsRow) HasLongTermDebtTotal() bool`

HasLongTermDebtTotal returns a boolean if a field has been set.

### SetLongTermDebtTotalNil

`func (o *CompanyFinancialsRow) SetLongTermDebtTotalNil(b bool)`

 SetLongTermDebtTotalNil sets the value for LongTermDebtTotal to be an explicit nil

### UnsetLongTermDebtTotal
`func (o *CompanyFinancialsRow) UnsetLongTermDebtTotal()`

UnsetLongTermDebtTotal ensures that no value is present for LongTermDebtTotal, not even an explicit nil
### GetIntangibleAssets

`func (o *CompanyFinancialsRow) GetIntangibleAssets() float32`

GetIntangibleAssets returns the IntangibleAssets field if non-nil, zero value otherwise.

### GetIntangibleAssetsOk

`func (o *CompanyFinancialsRow) GetIntangibleAssetsOk() (*float32, bool)`

GetIntangibleAssetsOk returns a tuple with the IntangibleAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntangibleAssets

`func (o *CompanyFinancialsRow) SetIntangibleAssets(v float32)`

SetIntangibleAssets sets IntangibleAssets field to given value.

### HasIntangibleAssets

`func (o *CompanyFinancialsRow) HasIntangibleAssets() bool`

HasIntangibleAssets returns a boolean if a field has been set.

### SetIntangibleAssetsNil

`func (o *CompanyFinancialsRow) SetIntangibleAssetsNil(b bool)`

 SetIntangibleAssetsNil sets the value for IntangibleAssets to be an explicit nil

### UnsetIntangibleAssets
`func (o *CompanyFinancialsRow) UnsetIntangibleAssets()`

UnsetIntangibleAssets ensures that no value is present for IntangibleAssets, not even an explicit nil
### GetPropertyPlantEquipment

`func (o *CompanyFinancialsRow) GetPropertyPlantEquipment() float32`

GetPropertyPlantEquipment returns the PropertyPlantEquipment field if non-nil, zero value otherwise.

### GetPropertyPlantEquipmentOk

`func (o *CompanyFinancialsRow) GetPropertyPlantEquipmentOk() (*float32, bool)`

GetPropertyPlantEquipmentOk returns a tuple with the PropertyPlantEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPlantEquipment

`func (o *CompanyFinancialsRow) SetPropertyPlantEquipment(v float32)`

SetPropertyPlantEquipment sets PropertyPlantEquipment field to given value.

### HasPropertyPlantEquipment

`func (o *CompanyFinancialsRow) HasPropertyPlantEquipment() bool`

HasPropertyPlantEquipment returns a boolean if a field has been set.

### SetPropertyPlantEquipmentNil

`func (o *CompanyFinancialsRow) SetPropertyPlantEquipmentNil(b bool)`

 SetPropertyPlantEquipmentNil sets the value for PropertyPlantEquipment to be an explicit nil

### UnsetPropertyPlantEquipment
`func (o *CompanyFinancialsRow) UnsetPropertyPlantEquipment()`

UnsetPropertyPlantEquipment ensures that no value is present for PropertyPlantEquipment, not even an explicit nil
### GetGoodwill

`func (o *CompanyFinancialsRow) GetGoodwill() float32`

GetGoodwill returns the Goodwill field if non-nil, zero value otherwise.

### GetGoodwillOk

`func (o *CompanyFinancialsRow) GetGoodwillOk() (*float32, bool)`

GetGoodwillOk returns a tuple with the Goodwill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodwill

`func (o *CompanyFinancialsRow) SetGoodwill(v float32)`

SetGoodwill sets Goodwill field to given value.

### HasGoodwill

`func (o *CompanyFinancialsRow) HasGoodwill() bool`

HasGoodwill returns a boolean if a field has been set.

### SetGoodwillNil

`func (o *CompanyFinancialsRow) SetGoodwillNil(b bool)`

 SetGoodwillNil sets the value for Goodwill to be an explicit nil

### UnsetGoodwill
`func (o *CompanyFinancialsRow) UnsetGoodwill()`

UnsetGoodwill ensures that no value is present for Goodwill, not even an explicit nil
### GetInventory

`func (o *CompanyFinancialsRow) GetInventory() float32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *CompanyFinancialsRow) GetInventoryOk() (*float32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *CompanyFinancialsRow) SetInventory(v float32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *CompanyFinancialsRow) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *CompanyFinancialsRow) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *CompanyFinancialsRow) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
### GetAccountsPayable

`func (o *CompanyFinancialsRow) GetAccountsPayable() float32`

GetAccountsPayable returns the AccountsPayable field if non-nil, zero value otherwise.

### GetAccountsPayableOk

`func (o *CompanyFinancialsRow) GetAccountsPayableOk() (*float32, bool)`

GetAccountsPayableOk returns a tuple with the AccountsPayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsPayable

`func (o *CompanyFinancialsRow) SetAccountsPayable(v float32)`

SetAccountsPayable sets AccountsPayable field to given value.

### HasAccountsPayable

`func (o *CompanyFinancialsRow) HasAccountsPayable() bool`

HasAccountsPayable returns a boolean if a field has been set.

### SetAccountsPayableNil

`func (o *CompanyFinancialsRow) SetAccountsPayableNil(b bool)`

 SetAccountsPayableNil sets the value for AccountsPayable to be an explicit nil

### UnsetAccountsPayable
`func (o *CompanyFinancialsRow) UnsetAccountsPayable()`

UnsetAccountsPayable ensures that no value is present for AccountsPayable, not even an explicit nil
### GetAccountsReceivable

`func (o *CompanyFinancialsRow) GetAccountsReceivable() float32`

GetAccountsReceivable returns the AccountsReceivable field if non-nil, zero value otherwise.

### GetAccountsReceivableOk

`func (o *CompanyFinancialsRow) GetAccountsReceivableOk() (*float32, bool)`

GetAccountsReceivableOk returns a tuple with the AccountsReceivable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsReceivable

`func (o *CompanyFinancialsRow) SetAccountsReceivable(v float32)`

SetAccountsReceivable sets AccountsReceivable field to given value.

### HasAccountsReceivable

`func (o *CompanyFinancialsRow) HasAccountsReceivable() bool`

HasAccountsReceivable returns a boolean if a field has been set.

### SetAccountsReceivableNil

`func (o *CompanyFinancialsRow) SetAccountsReceivableNil(b bool)`

 SetAccountsReceivableNil sets the value for AccountsReceivable to be an explicit nil

### UnsetAccountsReceivable
`func (o *CompanyFinancialsRow) UnsetAccountsReceivable()`

UnsetAccountsReceivable ensures that no value is present for AccountsReceivable, not even an explicit nil
### GetWeightedAverageSharesBasic

`func (o *CompanyFinancialsRow) GetWeightedAverageSharesBasic() float32`

GetWeightedAverageSharesBasic returns the WeightedAverageSharesBasic field if non-nil, zero value otherwise.

### GetWeightedAverageSharesBasicOk

`func (o *CompanyFinancialsRow) GetWeightedAverageSharesBasicOk() (*float32, bool)`

GetWeightedAverageSharesBasicOk returns a tuple with the WeightedAverageSharesBasic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAverageSharesBasic

`func (o *CompanyFinancialsRow) SetWeightedAverageSharesBasic(v float32)`

SetWeightedAverageSharesBasic sets WeightedAverageSharesBasic field to given value.

### HasWeightedAverageSharesBasic

`func (o *CompanyFinancialsRow) HasWeightedAverageSharesBasic() bool`

HasWeightedAverageSharesBasic returns a boolean if a field has been set.

### SetWeightedAverageSharesBasicNil

`func (o *CompanyFinancialsRow) SetWeightedAverageSharesBasicNil(b bool)`

 SetWeightedAverageSharesBasicNil sets the value for WeightedAverageSharesBasic to be an explicit nil

### UnsetWeightedAverageSharesBasic
`func (o *CompanyFinancialsRow) UnsetWeightedAverageSharesBasic()`

UnsetWeightedAverageSharesBasic ensures that no value is present for WeightedAverageSharesBasic, not even an explicit nil
### GetWeightedAverageSharesDiluted

`func (o *CompanyFinancialsRow) GetWeightedAverageSharesDiluted() float32`

GetWeightedAverageSharesDiluted returns the WeightedAverageSharesDiluted field if non-nil, zero value otherwise.

### GetWeightedAverageSharesDilutedOk

`func (o *CompanyFinancialsRow) GetWeightedAverageSharesDilutedOk() (*float32, bool)`

GetWeightedAverageSharesDilutedOk returns a tuple with the WeightedAverageSharesDiluted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAverageSharesDiluted

`func (o *CompanyFinancialsRow) SetWeightedAverageSharesDiluted(v float32)`

SetWeightedAverageSharesDiluted sets WeightedAverageSharesDiluted field to given value.

### HasWeightedAverageSharesDiluted

`func (o *CompanyFinancialsRow) HasWeightedAverageSharesDiluted() bool`

HasWeightedAverageSharesDiluted returns a boolean if a field has been set.

### SetWeightedAverageSharesDilutedNil

`func (o *CompanyFinancialsRow) SetWeightedAverageSharesDilutedNil(b bool)`

 SetWeightedAverageSharesDilutedNil sets the value for WeightedAverageSharesDiluted to be an explicit nil

### UnsetWeightedAverageSharesDiluted
`func (o *CompanyFinancialsRow) UnsetWeightedAverageSharesDiluted()`

UnsetWeightedAverageSharesDiluted ensures that no value is present for WeightedAverageSharesDiluted, not even an explicit nil
### GetBurnBasisUsd

`func (o *CompanyFinancialsRow) GetBurnBasisUsd() float32`

GetBurnBasisUsd returns the BurnBasisUsd field if non-nil, zero value otherwise.

### GetBurnBasisUsdOk

`func (o *CompanyFinancialsRow) GetBurnBasisUsdOk() (*float32, bool)`

GetBurnBasisUsdOk returns a tuple with the BurnBasisUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBasisUsd

`func (o *CompanyFinancialsRow) SetBurnBasisUsd(v float32)`

SetBurnBasisUsd sets BurnBasisUsd field to given value.

### HasBurnBasisUsd

`func (o *CompanyFinancialsRow) HasBurnBasisUsd() bool`

HasBurnBasisUsd returns a boolean if a field has been set.

### SetBurnBasisUsdNil

`func (o *CompanyFinancialsRow) SetBurnBasisUsdNil(b bool)`

 SetBurnBasisUsdNil sets the value for BurnBasisUsd to be an explicit nil

### UnsetBurnBasisUsd
`func (o *CompanyFinancialsRow) UnsetBurnBasisUsd()`

UnsetBurnBasisUsd ensures that no value is present for BurnBasisUsd, not even an explicit nil
### GetBurnBasisYears

`func (o *CompanyFinancialsRow) GetBurnBasisYears() int32`

GetBurnBasisYears returns the BurnBasisYears field if non-nil, zero value otherwise.

### GetBurnBasisYearsOk

`func (o *CompanyFinancialsRow) GetBurnBasisYearsOk() (*int32, bool)`

GetBurnBasisYearsOk returns a tuple with the BurnBasisYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBasisYears

`func (o *CompanyFinancialsRow) SetBurnBasisYears(v int32)`

SetBurnBasisYears sets BurnBasisYears field to given value.

### HasBurnBasisYears

`func (o *CompanyFinancialsRow) HasBurnBasisYears() bool`

HasBurnBasisYears returns a boolean if a field has been set.

### SetBurnBasisYearsNil

`func (o *CompanyFinancialsRow) SetBurnBasisYearsNil(b bool)`

 SetBurnBasisYearsNil sets the value for BurnBasisYears to be an explicit nil

### UnsetBurnBasisYears
`func (o *CompanyFinancialsRow) UnsetBurnBasisYears()`

UnsetBurnBasisYears ensures that no value is present for BurnBasisYears, not even an explicit nil
### GetRunwayMonths

`func (o *CompanyFinancialsRow) GetRunwayMonths() float32`

GetRunwayMonths returns the RunwayMonths field if non-nil, zero value otherwise.

### GetRunwayMonthsOk

`func (o *CompanyFinancialsRow) GetRunwayMonthsOk() (*float32, bool)`

GetRunwayMonthsOk returns a tuple with the RunwayMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunwayMonths

`func (o *CompanyFinancialsRow) SetRunwayMonths(v float32)`

SetRunwayMonths sets RunwayMonths field to given value.

### HasRunwayMonths

`func (o *CompanyFinancialsRow) HasRunwayMonths() bool`

HasRunwayMonths returns a boolean if a field has been set.

### SetRunwayMonthsNil

`func (o *CompanyFinancialsRow) SetRunwayMonthsNil(b bool)`

 SetRunwayMonthsNil sets the value for RunwayMonths to be an explicit nil

### UnsetRunwayMonths
`func (o *CompanyFinancialsRow) UnsetRunwayMonths()`

UnsetRunwayMonths ensures that no value is present for RunwayMonths, not even an explicit nil
### GetSecTotalRevenueUsdMillions

`func (o *CompanyFinancialsRow) GetSecTotalRevenueUsdMillions() float32`

GetSecTotalRevenueUsdMillions returns the SecTotalRevenueUsdMillions field if non-nil, zero value otherwise.

### GetSecTotalRevenueUsdMillionsOk

`func (o *CompanyFinancialsRow) GetSecTotalRevenueUsdMillionsOk() (*float32, bool)`

GetSecTotalRevenueUsdMillionsOk returns a tuple with the SecTotalRevenueUsdMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecTotalRevenueUsdMillions

`func (o *CompanyFinancialsRow) SetSecTotalRevenueUsdMillions(v float32)`

SetSecTotalRevenueUsdMillions sets SecTotalRevenueUsdMillions field to given value.

### HasSecTotalRevenueUsdMillions

`func (o *CompanyFinancialsRow) HasSecTotalRevenueUsdMillions() bool`

HasSecTotalRevenueUsdMillions returns a boolean if a field has been set.

### SetSecTotalRevenueUsdMillionsNil

`func (o *CompanyFinancialsRow) SetSecTotalRevenueUsdMillionsNil(b bool)`

 SetSecTotalRevenueUsdMillionsNil sets the value for SecTotalRevenueUsdMillions to be an explicit nil

### UnsetSecTotalRevenueUsdMillions
`func (o *CompanyFinancialsRow) UnsetSecTotalRevenueUsdMillions()`

UnsetSecTotalRevenueUsdMillions ensures that no value is present for SecTotalRevenueUsdMillions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



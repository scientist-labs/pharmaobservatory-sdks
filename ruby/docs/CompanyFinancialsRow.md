# PharmaObservatory::CompanyFinancialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **reporting_org_id** | **String** | Legal-entity level: the organization record that filed (the SEC filer). | [optional] |
| **fiscal_year** | **Integer** | The fiscal year the figures belong to. | [optional] |
| **is_public** | **Boolean** | True when the record has an SEC registration number and is a listed filer (a Form D issuer that is not listed is false). | [optional] |
| **edgar_cik** | **String** | The SEC registration number (CIK, Central Index Key) of the filer. | [optional] |
| **ticker** | **String** | The stock ticker. | [optional] |
| **exchange** | **String** | The primary exchange listing. | [optional] |
| **public_float_usd** | **Float** | The latest reported public float, in USD, repeated on every year row of the filer (a proxy for market value, not a live market capitalization). | [optional] |
| **public_float_fy** | **Integer** | The fiscal year of public_float_usd. | [optional] |
| **latest_shares_outstanding** | **Float** | The latest reported cover-page share count of the filer, repeated on every year row (the denominator of an ownership percentage). | [optional] |
| **latest_shares_outstanding_fy** | **Integer** | The fiscal year of latest_shares_outstanding. | [optional] |
| **latest_financial_fy** | **Integer** | The latest fiscal year the filer reported any line item. | [optional] |
| **revenue** | **Float** | Total revenue for the fiscal year, in USD (a non-USD filer is converted at the calendar-year average rate; null when no rate exists). | [optional] |
| **net_income** | **Float** | Net income for the fiscal year, in USD. | [optional] |
| **rnd** | **Float** | Research and development expense for the fiscal year, in USD. | [optional] |
| **assets** | **Float** | Total assets at the fiscal year end, in USD (converted at the year-end rate). | [optional] |
| **equity** | **Float** | Stockholders&#39; equity at the fiscal year end, in USD. | [optional] |
| **public_float** | **Float** | The public float on the annual report cover (the value of shares held by non-affiliates at the most recent second fiscal quarter end), in USD, anchored to the fiscal year it was filed with. | [optional] |
| **cash** | **Float** | Cash and cash equivalents at the fiscal year end, in USD. | [optional] |
| **short_term_investments** | **Float** | Short-term investments at the fiscal year end, in USD. | [optional] |
| **cash_and_investments** | **Float** | Cash plus short-term investments (a missing investments line counts as zero), in USD; null when cash is null. | [optional] |
| **operating_cash_flow** | **Float** | Net cash from operating activities for the fiscal year, in USD; negative means burning. | [optional] |
| **shares_outstanding** | **Float** | Shares outstanding on the annual report cover date, a count (never currency-converted); the cover date falls 0 to 3 months after the fiscal year end. | [optional] |
| **stock_based_compensation** | **Float** | Stock-based compensation expense, in USD. | [optional] |
| **operating_expenses** | **Float** | Operating expenses, in USD. | [optional] |
| **dividends_paid** | **Float** | Dividends paid, in USD, as filed with no sign test (filers differ on whether an outflow is positive or negative). | [optional] |
| **eps_basic** | **Float** | Basic earnings per share, in USD. | [optional] |
| **eps_diluted** | **Float** | Diluted earnings per share, in USD. | [optional] |
| **total_liabilities** | **Float** | Total liabilities at the fiscal year end, in USD. | [optional] |
| **long_term_debt_noncurrent** | **Float** | The non-current portion of long-term debt, in USD (null for a filer that reports only a total). | [optional] |
| **long_term_debt_total** | **Float** | Long-term debt, current plus non-current, in USD. | [optional] |
| **intangible_assets** | **Float** | Intangible assets, in USD. | [optional] |
| **property_plant_equipment** | **Float** | Property, plant and equipment, in USD. | [optional] |
| **goodwill** | **Float** | Goodwill, in USD. | [optional] |
| **inventory** | **Float** | Inventory, in USD. | [optional] |
| **accounts_payable** | **Float** | Accounts payable, in USD. | [optional] |
| **accounts_receivable** | **Float** | Accounts receivable, in USD. | [optional] |
| **weighted_average_shares_basic** | **Float** | Weighted average basic shares, a count. | [optional] |
| **weighted_average_shares_diluted** | **Float** | Weighted average diluted shares, a count. | [optional] |
| **burn_basis_usd** | **Float** | The average operating cash flow over this fiscal year and the two before it (the years reported), in USD; negative means burning. | [optional] |
| **burn_basis_years** | **Integer** | How many reported years burn_basis_usd averages over, 1 to 3. | [optional] |
| **runway_months** | **Float** | Derived, not reported: cash_and_investments divided by the monthly burn (burn_basis_usd over 12), in months, rounded to 0.1. Null when the filer generated cash or has no cash figure. | [optional] |
| **sec_total_revenue_usd_millions** | **Float** | The filing-level total revenue extracted from the 10-K or 20-F, in USD millions (the line the revenue-by-drug chart draws over the per-product stack); null where no total was extracted. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyFinancialsRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  reporting_org_id: null,
  fiscal_year: null,
  is_public: null,
  edgar_cik: null,
  ticker: null,
  exchange: null,
  public_float_usd: null,
  public_float_fy: null,
  latest_shares_outstanding: null,
  latest_shares_outstanding_fy: null,
  latest_financial_fy: null,
  revenue: null,
  net_income: null,
  rnd: null,
  assets: null,
  equity: null,
  public_float: null,
  cash: null,
  short_term_investments: null,
  cash_and_investments: null,
  operating_cash_flow: null,
  shares_outstanding: null,
  stock_based_compensation: null,
  operating_expenses: null,
  dividends_paid: null,
  eps_basic: null,
  eps_diluted: null,
  total_liabilities: null,
  long_term_debt_noncurrent: null,
  long_term_debt_total: null,
  intangible_assets: null,
  property_plant_equipment: null,
  goodwill: null,
  inventory: null,
  accounts_payable: null,
  accounts_receivable: null,
  weighted_average_shares_basic: null,
  weighted_average_shares_diluted: null,
  burn_basis_usd: null,
  burn_basis_years: null,
  runway_months: null,
  sec_total_revenue_usd_millions: null
)
```


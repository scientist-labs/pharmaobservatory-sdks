# PharmaObservatory::BrandRevenueRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **brand_key** | **String** | The brand key: an upper-case alphanumeric trade-name token (the brand&#39;s leading name tokens), the same key the brands pages use. Not a display name; brand_name is the display form. | [optional] |
| **product_name** | **String** | The product name as filed. | [optional] |
| **reporter_org_id** | **String** | Legal-entity level: the SEC filer that reported the figure, as filed (edgar:&lt;CIK&gt;). | [optional] |
| **reporter_company_id** | **String** | Company level: the reporting company&#39;s company (null when unresolved). | [optional] |
| **reporter_company_name** | **String** | Display name of the reporting company&#39;s company. | [optional] |
| **reporter_group_id** | **String** | Group level: the reporting company&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **reporter_group_name** | **String** | Display name of the reporting company&#39;s group. | [optional] |
| **fiscal_year** | **Integer** | The fiscal year of the figure. | [optional] |
| **revenue_basis** | **String** | The standardised basis: worldwide, us, ex_us, region or country. | [optional] |
| **geography** | **String** | The normalised geography: worldwide, us or ex_us for the spine (equal to the basis), the filer&#39;s own label for a region or country line. | [optional] |
| **revenue_usd_millions** | **Float** | The figure in USD millions (null where no unit or exchange rate is available). | [optional] |
| **revenue_amount** | **Float** | The figure as reported, in reporting_unit and revenue_currency; can be negative for adjustment or return lines. | [optional] |
| **revenue_currency** | **String** | The filing&#39;s currency (ISO code). | [optional] |
| **reporting_unit** | **String** | The filing&#39;s reporting unit (units, thousands, millions or billions); null when the filing did not state a readable one. | [optional] |
| **as_reported** | **Boolean** | True when the figure comes from the filing for its own fiscal year, false when it is a comparative year read from a later filing. | [optional] |
| **n_filings** | **Integer** | Distinct filings that report the figure. | [optional] |
| **restated** | **Boolean** | True when several filings report the figure and differ by more than one percent. | [optional] |
| **confidence** | **Float** | The extraction&#39;s confidence in the line, 0 to 1. | [optional] |
| **source_url** | **String** | The link to the SEC filing the figure was extracted from. | [optional] |
| **evidence_quote** | **String** | The sentence of the filing the figure was extracted from. | [optional] |
| **match_method** | **String** | How the product name matched the brand: alias_prefix, product_prefix or curated_alias. A label of the mechanism, not a confidence. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BrandRevenueRow.new(
  molecule_id: null,
  brand_key: null,
  product_name: null,
  reporter_org_id: null,
  reporter_company_id: null,
  reporter_company_name: null,
  reporter_group_id: null,
  reporter_group_name: null,
  fiscal_year: null,
  revenue_basis: null,
  geography: null,
  revenue_usd_millions: null,
  revenue_amount: null,
  revenue_currency: null,
  reporting_unit: null,
  as_reported: null,
  n_filings: null,
  restated: null,
  confidence: null,
  source_url: null,
  evidence_quote: null,
  match_method: null
)
```


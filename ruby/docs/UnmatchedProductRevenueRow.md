# PharmaObservatory::UnmatchedProductRevenueRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **revenue_line_id** | **String** | The revenue line id: a hash of the document, product name, fiscal year, basis and geography. | [optional] |
| **filer_cik** | **String** | The filer&#39;s SEC Central Index Key, ten digits. | [optional] |
| **filer_name** | **String** | The filer&#39;s name as filed (the reporting company of the document). | [optional] |
| **filer_org_id** | **String** | Legal-entity level: the filer&#39;s organization record, edgar:&lt;cik&gt; when that record exists, else the single record carrying the CIK (null when none; the CIK is kept). | [optional] |
| **filer_company_id** | **String** | Company level: the filer&#39;s company (null when the organization does not resolve to a company). | [optional] |
| **filer_company_name** | **String** | Display name of the filer&#39;s company. | [optional] |
| **filer_group_id** | **String** | Group level: the filer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **filer_group_name** | **String** | Display name of the filer&#39;s group. | [optional] |
| **document_id** | **String** | The extracted filing document id. | [optional] |
| **accession_number** | **String** | The SEC accession number of the filing. | [optional] |
| **form** | **String** | The SEC form the line was extracted from (10-K or 20-F). | [optional] |
| **filing_date** | **Date** | The filing date of the document. | [optional] |
| **fiscal_year_covered** | **Integer** | The fiscal year the document covers, as the extraction states it. | [optional] |
| **product_name** | **String** | The product name as filed, never a drug name: the brand gate did not claim it, and devices, services and apparel that share a word with a drug land here by design. | [optional] |
| **cleaned_key** | **String** | The product name reduced to its brand-alias key, the form the nearest-name pointer compares. | [optional] |
| **fiscal_year** | **Integer** | The fiscal year the line reports revenue for. | [optional] |
| **revenue_basis** | **String** | What the line&#39;s revenue covers, as the extraction states it: worldwide, us, ex_us, region or country. | [optional] |
| **geography** | **String** | The region or country the line covers, as filed; null for a worldwide or us line. | [optional] |
| **revenue_amount** | **Float** | The revenue as filed, in the currency and unit the line states (revenue_currency and reporting_unit). | [optional] |
| **revenue_currency** | **String** | The currency the line states its revenue in. | [optional] |
| **reporting_unit** | **String** | The unit the document states its amounts in (units, thousands, millions or billions). | [optional] |
| **revenue_usd_millions** | **Float** | The line&#39;s revenue in US dollars millions: the amount normalized by reporting_unit, then converted at the annual average exchange rate of the filing&#39;s fiscal year. Null when the unit is unreadable or no rate exists for that currency and year. The extractor&#39;s own USD figure is not published, because it reads 0 for a non-USD line. | [optional] |
| **is_latest_worldwide_line** | **Boolean** | True for the single worldwide line per filer and product with a USD figure from the latest fiscal year (the line the residual mart takes as the product&#39;s latest worldwide figure; ties broken by the headline filing, then the earlier filing date). | [optional] |
| **product_line_count** | **Integer** | How many extracted lines the product has, repeated on every line of the product. | [optional] |
| **best_alias_key** | **String** | The nearest trade-name alias by trigram similarity: a pointer for a person to decide on, not a match. Null when no alias is close enough. | [optional] |
| **best_alias_brand_key** | **String** | The brand key of the nearest alias. | [optional] |
| **best_alias_similarity** | **Float** | The trigram similarity between cleaned_key and best_alias_key, 0 to 1. | [optional] |
| **brand_or_generic** | **String** | The extraction&#39;s own label for the product (brand, generic, franchise and so on). | [optional] |
| **indication** | **String** | The indication the extraction attached to the product, as filed. | [optional] |
| **evidence_quote** | **String** | The filing&#39;s own text the line was extracted from. | [optional] |
| **extraction_confidence** | **Float** | The language model&#39;s self-reported confidence for the line, 0 to 1; not a reliability measure. | [optional] |
| **source_url** | **String** | The link to the filing on SEC EDGAR. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::UnmatchedProductRevenueRow.new(
  revenue_line_id: null,
  filer_cik: null,
  filer_name: null,
  filer_org_id: null,
  filer_company_id: null,
  filer_company_name: null,
  filer_group_id: null,
  filer_group_name: null,
  document_id: null,
  accession_number: null,
  form: null,
  filing_date: null,
  fiscal_year_covered: null,
  product_name: null,
  cleaned_key: null,
  fiscal_year: null,
  revenue_basis: null,
  geography: null,
  revenue_amount: null,
  revenue_currency: null,
  reporting_unit: null,
  revenue_usd_millions: null,
  is_latest_worldwide_line: null,
  product_line_count: null,
  best_alias_key: null,
  best_alias_brand_key: null,
  best_alias_similarity: null,
  brand_or_generic: null,
  indication: null,
  evidence_quote: null,
  extraction_confidence: null,
  source_url: null
)
```


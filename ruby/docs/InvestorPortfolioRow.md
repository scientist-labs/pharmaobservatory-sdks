# PharmaObservatory::InvestorPortfolioRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investor_id** | **String** | The investor key. | [optional] |
| **issuer_cik** | **String** | The issuer&#39;s SEC Central Index Key (CIK), ten digits. | [optional] |
| **issuer_name** | **String** | The issuer&#39;s name: its organization record&#39;s name, else the SEC company name. | [optional] |
| **issuer_org_id** | **String** | Legal-entity level: the issuer&#39;s organization record (null when unresolved). | [optional] |
| **issuer_company_id** | **String** | Company level: the issuer&#39;s company (null when unresolved). | [optional] |
| **issuer_company_name** | **String** | Display name of the issuer&#39;s company. | [optional] |
| **issuer_group_id** | **String** | Group level: the issuer&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **issuer_group_name** | **String** | Display name of the issuer&#39;s group. | [optional] |
| **relationship_types** | **Array&lt;String&gt;** | How the investor is named against the issuer: holder_5pct (a 5 percent holder) or selling_stockholder. | [optional] |
| **edge_count** | **Integer** | Filing mentions behind the row, across every filing and relationship type. | [optional] |
| **first_observed_date** | **Date** | The earliest filing date the investor was named against the issuer. | [optional] |
| **last_observed_date** | **Date** | The latest filing date the investor was named against the issuer. | [optional] |
| **is_life_sciences_issuer** | **Boolean** | True when the issuer&#39;s SIC (Standard Industrial Classification) code is a life-sciences code, false when it is another code, null when the issuer has no SIC evidence. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestorPortfolioRow.new(
  investor_id: null,
  issuer_cik: null,
  issuer_name: null,
  issuer_org_id: null,
  issuer_company_id: null,
  issuer_company_name: null,
  issuer_group_id: null,
  issuer_group_name: null,
  relationship_types: null,
  edge_count: null,
  first_observed_date: null,
  last_observed_date: null,
  is_life_sciences_issuer: null
)
```


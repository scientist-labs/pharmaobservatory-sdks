# PharmaObservatory::ApiV1PatentCliffExposureGet200ResponseExposuresInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **canonical_drug_id** | **String** |  |  |
| **drug_name** | **String** |  |  |
| **canonical_org_id** | **String** |  |  |
| **company** | **String** |  |  |
| **cliff_year** | **Integer** |  |  |
| **cliff_date** | **String** |  |  |
| **cliff_basis** | **String** |  |  |
| **medicare_spend** | **String** |  |  |
| **medicare_spend_year** | **Integer** |  |  |
| **disclosed_revenue_usd_millions** | **String** |  |  |
| **disclosed_revenue_year** | **Integer** |  |  |
| **size_usd** | **String** |  |  |
| **size_basis** | **String** |  |  |
| **generic_available** | **Boolean** |  |  |
| **biosimilar_available** | **Boolean** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1PatentCliffExposureGet200ResponseExposuresInner.new(
  canonical_drug_id: null,
  drug_name: null,
  canonical_org_id: null,
  company: null,
  cliff_year: null,
  cliff_date: null,
  cliff_basis: null,
  medicare_spend: null,
  medicare_spend_year: null,
  disclosed_revenue_usd_millions: null,
  disclosed_revenue_year: null,
  size_usd: null,
  size_basis: null,
  generic_available: null,
  biosimilar_available: null
)
```


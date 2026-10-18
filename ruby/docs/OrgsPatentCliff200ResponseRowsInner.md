# PharmaObservatory::OrgsPatentCliff200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **canonical_drug_id** | **String** |  |  |
| **drug_name** | **String** |  |  |
| **cliff_year** | **Integer** |  |  |
| **cliff_date** | **String** |  |  |
| **cliff_basis** | **String** |  |  |
| **cliff_source** | **Object** |  |  |
| **cliff_confidence** | **Object** |  |  |
| **medicare_spend** | **Object** |  |  |
| **has_generic** | **Boolean** |  |  |
| **biosimilar_available** | **Boolean** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsPatentCliff200ResponseRowsInner.new(
  canonical_drug_id: null,
  drug_name: null,
  cliff_year: null,
  cliff_date: null,
  cliff_basis: null,
  cliff_source: null,
  cliff_confidence: null,
  medicare_spend: null,
  has_generic: null,
  biosimilar_available: null
)
```


# PharmaObservatory::DrugPricingRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **pricing_unit** | **String** | The unit the costs are per (EA each, ML milliliter, GM gram). | [optional] |
| **priced_ndcs** | **Integer** | Packages (National Drug Codes) priced. | [optional] |
| **min_unit_cost** | **Float** | The lowest cost per unit. | [optional] |
| **median_unit_cost** | **Float** | The median cost per unit. | [optional] |
| **max_unit_cost** | **Float** | The highest cost per unit. | [optional] |
| **has_brand_priced** | **Boolean** | True when a brand package is priced. | [optional] |
| **has_generic_priced** | **Boolean** | True when a generic package is priced. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugPricingRow.new(
  molecule_id: null,
  pricing_unit: null,
  priced_ndcs: null,
  min_unit_cost: null,
  median_unit_cost: null,
  max_unit_cost: null,
  has_brand_priced: null,
  has_generic_priced: null
)
```


# PharmaObservatory::DrugMedicareSpendRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **year** | **Integer** | The calendar year of the spending. | [optional] |
| **total_spend** | **Float** | Medicare spending in US dollars. | [optional] |
| **claims** | **Float** | Claims or services. | [optional] |
| **beneficiaries** | **Float** | Beneficiaries, summed over payers. | [optional] |
| **payers** | **Array&lt;String&gt;** | The Medicare programs summed (Part B, Part D). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugMedicareSpendRow.new(
  molecule_id: null,
  year: null,
  total_spend: null,
  claims: null,
  beneficiaries: null,
  payers: null
)
```


# PharmaObservatory::DrugCombinationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **relation** | **String** | has_component when the related drug is a component of this drug; component_of when this drug is a component of the related drug. | [optional] |
| **related_drug_id** | **String** | The molecule key of the related drug. | [optional] |
| **related_drug_name** | **String** | The related drug&#39;s display name. | [optional] |
| **combination_id** | **String** | A stable id for this drug, direction and related drug. | [optional] |
| **n_components** | **Integer** | Components in the combination product. | [optional] |
| **combo_unii_mixture** | **Boolean** | True when the registry itself lists the combination as a mixture substance. | [optional] |
| **has_ingredient_confirmed** | **Boolean** | True when RxNorm&#39;s ingredient list confirms the components. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugCombinationsRow.new(
  molecule_id: null,
  relation: null,
  related_drug_id: null,
  related_drug_name: null,
  combination_id: null,
  n_components: null,
  combo_unii_mixture: null,
  has_ingredient_confirmed: null
)
```


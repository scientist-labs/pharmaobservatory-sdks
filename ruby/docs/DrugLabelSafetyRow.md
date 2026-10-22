# PharmaObservatory::DrugLabelSafetyRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **has_boxed_warning** | **Boolean** | True when the label carries a boxed warning. | [optional] |
| **boxed_warning_text** | **String** | The boxed warning text, null when there is none. | [optional] |
| **warnings** | **String** | The warnings and precautions text. | [optional] |
| **label_effective_date** | **Date** | The effective date of the label. | [optional] |
| **label_count** | **Integer** | Labels matched to the drug. | [optional] |
| **set_id** | **String** | The label set id. The DailyMed page is https://dailymed.nlm.nih.gov/dailymed/drugInfo.cfm?setid&#x3D; followed by this id. | [optional] |
| **match_method** | **String** | How the label was matched to the drug. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugLabelSafetyRow.new(
  molecule_id: null,
  has_boxed_warning: null,
  boxed_warning_text: null,
  warnings: null,
  label_effective_date: null,
  label_count: null,
  set_id: null,
  match_method: null
)
```


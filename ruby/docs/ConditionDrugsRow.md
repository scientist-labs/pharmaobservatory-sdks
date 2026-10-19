# PharmaObservatory::ConditionDrugsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **drug_id** | **String** | The drug key (a canonical drug id; the same value as molecule_id on this dataset). | [optional] |
| **molecule_id** | **String** | The molecule key the drug resolves to. | [optional] |
| **drug_name** | **String** | The molecule&#39;s display name. | [optional] |
| **n_trials** | **Integer** | Trials for this condition that test the molecule. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionDrugsRow.new(
  condition_id: null,
  drug_id: null,
  molecule_id: null,
  drug_name: null,
  n_trials: null
)
```


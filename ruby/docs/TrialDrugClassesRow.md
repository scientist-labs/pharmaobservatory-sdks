# PharmaObservatory::TrialDrugClassesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **drug_id** | **String** | The canonical drug key the trial&#39;s intervention names matched. | [optional] |
| **drug_name** | **String** | The drug&#39;s display name. | [optional] |
| **resolved_drug_id** | **String** | The drug key after same-substance drug records are merged. | [optional] |
| **approved_anywhere** | **Boolean** | True when the drug is marked approved by any regulator (false means investigational or unknown). | [optional] |
| **atc_class** | **String** | The ATC first-level class of the drug (null when it has none). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialDrugClassesRow.new(
  trial_id: null,
  drug_id: null,
  drug_name: null,
  resolved_drug_id: null,
  approved_anywhere: null,
  atc_class: null
)
```


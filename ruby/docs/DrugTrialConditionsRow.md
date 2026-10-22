# PharmaObservatory::DrugTrialConditionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **condition_id** | **String** | The condition key (a MeSH, Medical Subject Headings, descriptor id). | [optional] |
| **condition_name** | **String** | The condition&#39;s name. | [optional] |
| **phase** | **String** | The registry&#39;s phase code of the trial. | [optional] |
| **phase_rank** | **Integer** | The trial&#39;s phase as a number for max and sort (4 is Phase 4, 0 is not reported). | [optional] |
| **overall_status_summary** | **String** | The trial&#39;s overall status grouped as Completed, Ongoing, Terminated or Other. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugTrialConditionsRow.new(
  molecule_id: null,
  trial_id: null,
  condition_id: null,
  condition_name: null,
  phase: null,
  phase_rank: null,
  overall_status_summary: null
)
```


# PharmaObservatory::TrialAdverseEventsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **event_type** | **String** | serious or other. | [optional] |
| **organ_system** | **String** | The organ system class, as filed ((unspecified) where the filing gives none). | [optional] |
| **adverse_event_term** | **String** | The adverse event term, as filed. | [optional] |
| **subjects_affected** | **Integer** | Participants affected, summed over the trial&#39;s groups. | [optional] |
| **subjects_at_risk** | **Integer** | Participants at risk, summed over the trial&#39;s groups. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialAdverseEventsRow.new(
  trial_id: null,
  event_type: null,
  organ_system: null,
  adverse_event_term: null,
  subjects_affected: null,
  subjects_at_risk: null
)
```


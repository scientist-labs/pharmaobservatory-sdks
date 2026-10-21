# PharmaObservatory::TrialKolsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key. | [optional] |
| **investigator_person_name** | **String** | The investigator&#39;s name as filed. | [optional] |
| **investigator_affiliation** | **String** | The affiliation as filed. | [optional] |
| **roles** | **Array&lt;String&gt;** | The roles the person holds on this trial as registered. | [optional] |
| **is_principal_investigator** | **Boolean** | True when the person is named principal investigator on this trial. | [optional] |
| **authored_pubs** | **Float** | Publications the person authored. | [optional] |
| **pi_trials** | **Integer** | Trials on which the person is principal investigator, across the registry. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialKolsRow.new(
  trial_id: null,
  investigator_person_id: null,
  investigator_person_name: null,
  investigator_affiliation: null,
  roles: null,
  is_principal_investigator: null,
  authored_pubs: null,
  pi_trials: null
)
```


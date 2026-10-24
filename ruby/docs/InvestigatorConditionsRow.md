# PharmaObservatory::InvestigatorConditionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **investigator_person_name** | **String** | The investigator&#39;s name as the registry filed it, resolved across the person&#39;s merged name fragments. | [optional] |
| **investigator_affiliation** | **String** | The investigator&#39;s affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional] |
| **condition_id** | **String** | The condition&#39;s MeSH descriptor id, the id in the conditions page address (D001943). | [optional] |
| **condition_name** | **String** | The condition&#39;s MeSH name. | [optional] |
| **n_trials** | **Integer** | Distinct trials the person is named on in the condition. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestigatorConditionsRow.new(
  investigator_person_id: null,
  investigator_person_name: null,
  investigator_affiliation: null,
  condition_id: null,
  condition_name: null,
  n_trials: null
)
```


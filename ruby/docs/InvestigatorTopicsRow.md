# PharmaObservatory::InvestigatorTopicsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **descriptor_id** | **String** | The MeSH descriptor id. | [optional] |
| **descriptor_name** | **String** | The MeSH descriptor name. | [optional] |
| **is_disease** | **Boolean** | True when the descriptor is also a condition. | [optional] |
| **condition_id** | **String** | The descriptor id as a conditions id when the descriptor is a condition (null otherwise). | [optional] |
| **n_publications** | **Integer** | Distinct papers the person authored on the descriptor. | [optional] |
| **n_major_topic_publications** | **Integer** | Those papers where the descriptor is a major topic. | [optional] |
| **topic_weight** | **Integer** | n_publications plus n_major_topic_publications, so a major-topic paper counts twice. | [optional] |
| **topic_rank** | **Integer** | The person&#39;s own topic order by weight (1 is their strongest topic); not a ranking of people. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestigatorTopicsRow.new(
  investigator_person_id: null,
  descriptor_id: null,
  descriptor_name: null,
  is_disease: null,
  condition_id: null,
  n_publications: null,
  n_major_topic_publications: null,
  topic_weight: null,
  topic_rank: null
)
```


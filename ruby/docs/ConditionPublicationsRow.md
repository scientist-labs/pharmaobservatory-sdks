# PharmaObservatory::ConditionPublicationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **pmid** | **String** | The PubMed id. | [optional] |
| **title** | **String** | The article title. | [optional] |
| **pub_year** | **Integer** | Publication year. | [optional] |
| **journal_title** | **String** | The journal title. | [optional] |
| **evidence_type** | **String** | The article&#39;s PubMed publication type, one per article, picked by study design (Randomized Controlled Trial, Journal Article and so on). | [optional] |
| **is_major_topic** | **Boolean** | True when the condition is a major topic of the article. | [optional] |
| **n_linked_trials** | **Integer** | Registry trials linked to the article. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionPublicationsRow.new(
  condition_id: null,
  pmid: null,
  title: null,
  pub_year: null,
  journal_title: null,
  evidence_type: null,
  is_major_topic: null,
  n_linked_trials: null
)
```


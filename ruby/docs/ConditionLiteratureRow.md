# PharmaObservatory::ConditionLiteratureRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **pmid** | **String** | The PubMed id. | [optional] |
| **title** | **String** | The paper&#39;s title. | [optional] |
| **journal_title** | **String** | The journal&#39;s title. | [optional] |
| **pub_year** | **Integer** | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional] |
| **evidence_type** | **String** | The paper&#39;s study-design label, one normalized publication type by study-design priority. | [optional] |
| **is_trial_linked** | **Boolean** | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional] |
| **is_major_topic** | **Boolean** | True when the condition is a major MeSH topic of the paper. | [optional] |
| **citation_count** | **Integer** | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional] |
| **relative_citation_ratio** | **Float** | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionLiteratureRow.new(
  condition_id: null,
  pmid: null,
  title: null,
  journal_title: null,
  pub_year: null,
  evidence_type: null,
  is_trial_linked: null,
  is_major_topic: null,
  citation_count: null,
  relative_citation_ratio: null
)
```


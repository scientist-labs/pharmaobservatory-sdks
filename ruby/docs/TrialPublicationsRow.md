# PharmaObservatory::TrialPublicationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **pmid** | **String** | The PubMed id of the paper. | [optional] |
| **title** | **String** | The paper&#39;s title. | [optional] |
| **journal_title** | **String** | The journal&#39;s title (null when the paper&#39;s journal is not in the journal table). | [optional] |
| **pub_year** | **Integer** | The publication year (null when undated). | [optional] |
| **evidence_type** | **String** | The paper&#39;s study-design label, one normalized publication type by study-design priority; the page&#39;s design column. | [optional] |
| **citation** | **String** | The citation for the paper: one citation per paper, the longest among the trials that cite the PubMed id. | [optional] |
| **reference_types** | **Array&lt;String&gt;** | Every reference kind the registry filed for this paper on this trial, lower-cased (background, result, derived). | [optional] |
| **results** | **Boolean** | True when any reference kind of the row is a results reference; the page&#39;s Results label. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialPublicationsRow.new(
  trial_id: null,
  pmid: null,
  title: null,
  journal_title: null,
  pub_year: null,
  evidence_type: null,
  citation: null,
  reference_types: null,
  results: null
)
```


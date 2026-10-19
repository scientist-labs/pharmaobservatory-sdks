# PharmaObservatory::TrialExtractedEndpointsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **result_reported** | **Boolean** | Whether the extraction found a formal primary-endpoint result. | [optional] |
| **signal** | **String** | The extracted reading, verbatim: met, missed, mixed, reported_indeterminate or not_reported. | [optional] |
| **outcome_category** | **String** | The signal grouped as met, missed, mixed or unclear. | [optional] |
| **has_clean_outcome** | **Boolean** | Whether outcome_category is met, missed or mixed. | [optional] |
| **primary_endpoints** | **Array&lt;String&gt;** | The primary endpoint descriptions the extraction found. | [optional] |
| **primary_endpoint_count** | **Integer** | Number of primary endpoints extracted. | [optional] |
| **confidence** | **Float** | The extraction&#39;s confidence, from 0 to 1. | [optional] |
| **evidence_quote** | **String** | The text quoted as evidence for the reading. | [optional] |
| **source_url** | **String** | The registry page the reading was taken from. | [optional] |
| **model** | **String** | The model that produced the reading. | [optional] |
| **extracted_at** | **Time** | When the extraction ran. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialExtractedEndpointsRow.new(
  trial_id: null,
  result_reported: null,
  signal: null,
  outcome_category: null,
  has_clean_outcome: null,
  primary_endpoints: null,
  primary_endpoint_count: null,
  confidence: null,
  evidence_quote: null,
  source_url: null,
  model: null,
  extracted_at: null
)
```


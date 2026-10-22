# PharmaObservatory::RareDiseaseRarityEvidenceRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **rarity_evidence_id** | **String** | The identifier of the evidence link. | [optional] |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **rarity_assertion_id** | **String** | The assessment the evidence belongs to. | [optional] |
| **prevalence_assertion_id** | **String** | The prevalence observation compared. | [optional] |
| **jurisdiction_code** | **String** | The jurisdiction whose rule was applied (GB or EU). | [optional] |
| **evidence_kind** | **String** | How the observation was compared (prevalence_class_ratio). | [optional] |
| **comparison_result** | **String** | The result of comparing the observation with the threshold: rare, not_rare or unknown. | [optional] |
| **is_comparable** | **Boolean** | True when the observation could be compared with the threshold. | [optional] |
| **comparison_reason** | **String** | The reason the comparison gave its result. | [optional] |
| **assessment_date** | **Date** | The date of the assessment. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseRarityEvidenceRow.new(
  rarity_evidence_id: null,
  orpha_code: null,
  rarity_assertion_id: null,
  prevalence_assertion_id: null,
  jurisdiction_code: null,
  evidence_kind: null,
  comparison_result: null,
  is_comparable: null,
  comparison_reason: null,
  assessment_date: null
)
```


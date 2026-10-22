# PharmaObservatory::RareDiseaseRarityRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **rarity_assertion_id** | **String** | The identifier of the assessment. | [optional] |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **jurisdiction_code** | **String** | The jurisdiction whose definition was applied (GB or EU). | [optional] |
| **jurisdiction_name** | **String** | The jurisdiction&#39;s name. | [optional] |
| **rule_id** | **String** | The identifier of the rule applied. | [optional] |
| **definition_text** | **String** | The rule&#39;s definition as stated by the jurisdiction. | [optional] |
| **comparison_operator** | **String** | The comparison the rule applies to the threshold (&lt; or &lt;&#x3D;). | [optional] |
| **threshold_value** | **Float** | The threshold figure, to be read over population_denominator (for example 5 per 10000). | [optional] |
| **population_denominator** | **Float** | The denominator of the threshold figure. | [optional] |
| **population_unit** | **String** | The unit of the denominator. | [optional] |
| **effective_date** | **Date** | The date the rule took effect. | [optional] |
| **end_date** | **Date** | The date the rule ended, where it has. | [optional] |
| **assessment_date** | **Date** | The date of the assessment. | [optional] |
| **rarity_status** | **String** | rare, not_rare, unknown or conflicting, as a result of applying the rule to the evidence. | [optional] |
| **evidence_basis** | **String** | What the assessment rests on: the bounds of a source prevalence class, or no evidence in the jurisdiction&#39;s geography. | [optional] |
| **evidence_count** | **Integer** | Prevalence observations compared in the assessment. | [optional] |
| **official_source_url** | **String** | The jurisdiction&#39;s official document the rule comes from. | [optional] |
| **license_name** | **String** | The licence of the jurisdiction&#39;s document. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseRarityRow.new(
  rarity_assertion_id: null,
  orpha_code: null,
  jurisdiction_code: null,
  jurisdiction_name: null,
  rule_id: null,
  definition_text: null,
  comparison_operator: null,
  threshold_value: null,
  population_denominator: null,
  population_unit: null,
  effective_date: null,
  end_date: null,
  assessment_date: null,
  rarity_status: null,
  evidence_basis: null,
  evidence_count: null,
  official_source_url: null,
  license_name: null,
  attribution: null
)
```


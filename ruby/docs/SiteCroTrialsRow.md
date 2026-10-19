# PharmaObservatory::SiteCroTrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_id** | **String** | The site key: one facility spelling from the registry. | [optional] |
| **trial_id** | **String** | The trial&#39;s registry id (NCT number). | [optional] |
| **cro_org_id** | **String** | Legal-entity level: the CRO&#39;s organization record. | [optional] |
| **cro_name** | **String** | The CRO organization&#39;s name. | [optional] |
| **cro_company_id** | **String** | Company level: the company the CRO belongs to (null when unresolved). | [optional] |
| **cro_company_name** | **String** | Display name of the CRO&#39;s company. | [optional] |
| **cro_group_id** | **String** | Group level: the CRO&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **cro_group_name** | **String** | Display name of the CRO&#39;s group. | [optional] |
| **cro_attribution_method** | **String** | How the CRO was inferred (sponsor_match, collaborator_match or llm_extraction). | [optional] |
| **cro_confidence** | **Float** | Confidence of the inference, from 0 to 1. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall status of the trial. | [optional] |
| **start_year** | **Integer** | Year of the trial start. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SiteCroTrialsRow.new(
  site_id: null,
  trial_id: null,
  cro_org_id: null,
  cro_name: null,
  cro_company_id: null,
  cro_company_name: null,
  cro_group_id: null,
  cro_group_name: null,
  cro_attribution_method: null,
  cro_confidence: null,
  overall_status: null,
  start_year: null
)
```


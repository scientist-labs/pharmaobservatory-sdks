# PharmaObservatory::CompanyTargetsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sponsor_company_id** | **String** | Company level: the company whose lead-sponsored trials pursue the target. | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **target_id** | **String** | The target&#39;s gene identifier (Ensembl). | [optional] |
| **target_symbol** | **String** | The target&#39;s gene symbol. | [optional] |
| **target_name** | **String** | The target&#39;s name. | [optional] |
| **max_phase** | **Integer** | The furthest phase any qualifying program of the company reached against the target. | [optional] |
| **n_drugs** | **Integer** | Distinct drugs of the company&#39;s qualifying programs against the target. | [optional] |
| **n_conditions** | **Integer** | Distinct conditions of the company&#39;s qualifying programs against the target. | [optional] |
| **n_trials** | **Integer** | Distinct trials the company leads in those programs. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyTargetsRow.new(
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  target_id: null,
  target_symbol: null,
  target_name: null,
  max_phase: null,
  n_drugs: null,
  n_conditions: null,
  n_trials: null
)
```


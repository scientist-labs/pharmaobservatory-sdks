# PharmaObservatory::ApiV1LicensingCandidatesGet200ResponseCandidatesInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **canonical_org_id** | **String** |  |  |
| **org_name** | **String** |  |  |
| **canonical_drug_id** | **String** |  |  |
| **drug_name** | **String** |  |  |
| **condition_ui** | **String** |  |  |
| **condition_name** | **String** |  |  |
| **max_phase** | **String** |  |  |
| **n_trials** | **Integer** |  |  |
| **last_active_year** | **Integer** |  |  |
| **dalys_m** | **Float** |  |  |
| **deaths_m** | **Float** |  |  |
| **druggability** | **Integer** |  |  |
| **druggability_label** | **String** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1LicensingCandidatesGet200ResponseCandidatesInner.new(
  canonical_org_id: null,
  org_name: null,
  canonical_drug_id: null,
  drug_name: null,
  condition_ui: null,
  condition_name: null,
  max_phase: null,
  n_trials: null,
  last_active_year: null,
  dalys_m: null,
  deaths_m: null,
  druggability: null,
  druggability_label: null
)
```


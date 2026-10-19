# PharmaObservatory::SiteTrialDrugsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_id** | **String** | The site key: one facility spelling from the registry. | [optional] |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **drug_id** | **String** | The drug key. | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall status of the trial. | [optional] |
| **start_year** | **Integer** | Year of the trial start. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SiteTrialDrugsRow.new(
  site_id: null,
  trial_id: null,
  drug_id: null,
  drug_name: null,
  overall_status: null,
  start_year: null
)
```


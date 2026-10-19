# PharmaObservatory::SiteTrialConditionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_id** | **String** | The site key: one facility spelling from the registry. | [optional] |
| **trial_id** | **String** | The trial&#39;s registry id (NCT number). | [optional] |
| **condition_id** | **String** | The condition key (a MeSH, Medical Subject Headings, descriptor id). | [optional] |
| **condition_name** | **String** | The condition&#39;s name. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall status of the trial. | [optional] |
| **start_year** | **Integer** | Year of the trial start. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SiteTrialConditionsRow.new(
  site_id: null,
  trial_id: null,
  condition_id: null,
  condition_name: null,
  overall_status: null,
  start_year: null
)
```


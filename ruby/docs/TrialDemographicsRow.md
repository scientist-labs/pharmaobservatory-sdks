# PharmaObservatory::TrialDemographicsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **has_sex_data** | **Boolean** | Whether a sex baseline breakdown was posted. | [optional] |
| **has_race_data** | **Boolean** | Whether a race baseline breakdown was posted. | [optional] |
| **has_ethnicity_data** | **Boolean** | Whether an ethnicity baseline breakdown was posted. | [optional] |
| **has_age_data** | **Boolean** | Whether a categorical age baseline breakdown was posted. | [optional] |
| **pct_female** | **Float** | Percent of baseline participants recorded as female (null when sex was not posted). | [optional] |
| **race_breakdown** | **String** | A JSON object of race category to percent of baseline participants, as posted. | [optional] |
| **n_analyzed** | **Integer** | Baseline participants analyzed. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialDemographicsRow.new(
  trial_id: null,
  has_sex_data: null,
  has_race_data: null,
  has_ethnicity_data: null,
  has_age_data: null,
  pct_female: null,
  race_breakdown: null,
  n_analyzed: null
)
```


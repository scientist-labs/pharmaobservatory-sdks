# PharmaObservatory::RareDiseasesRecordAttributes

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **rare_disease_investigator_coverage** | [**RareDiseaseInvestigatorCoverageRow**](RareDiseaseInvestigatorCoverageRow.md) |  | [optional] |
| **rare_disease_site_coverage** | [**RareDiseaseSiteCoverageRow**](RareDiseaseSiteCoverageRow.md) |  | [optional] |
| **rare_disease_trial_coverage** | [**RareDiseaseTrialCoverageRow**](RareDiseaseTrialCoverageRow.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseasesRecordAttributes.new(
  rare_disease_investigator_coverage: null,
  rare_disease_site_coverage: null,
  rare_disease_trial_coverage: null
)
```


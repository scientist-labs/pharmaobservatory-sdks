# PharmaObservatory::RareDiseaseInvestigatorCoverageRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **coverage_state** | **String** | no_exact_mapping, no_direct_trials, no_publishable_investigator or computed. | [optional] |
| **n_direct_trials** | **Integer** | Exact direct trials of the disease. | [optional] |
| **n_direct_trials_with_named_pi** | **Integer** | Those trials that name a principal investigator. | [optional] |
| **n_pi_trial_edges** | **Integer** | Trial and principal investigator pairs among them. | [optional] |
| **n_pi_people** | **Integer** | Distinct people named as principal investigator. | [optional] |
| **n_investigators** | **Integer** | Rows of rare_disease_investigators for the disease. | [optional] |
| **n_name_ambiguous** | **Integer** | Investigators whose name another investigator of the disease shares. | [optional] |
| **registry_reference_date** | **Date** | The registry&#39;s latest update date the snapshot reflects. | [optional] |
| **license_name** | **String** | The licence of the Orphadata Science release the disease list rests on. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseInvestigatorCoverageRow.new(
  orpha_code: null,
  coverage_state: null,
  n_direct_trials: null,
  n_direct_trials_with_named_pi: null,
  n_pi_trial_edges: null,
  n_pi_people: null,
  n_investigators: null,
  n_name_ambiguous: null,
  registry_reference_date: null,
  license_name: null,
  attribution: null
)
```


# PharmaObservatory::ProgramsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **program_id** | **String** | The program key: a hash of drug_id and condition_id, one value per row. | [optional] |
| **drug_id** | **String** | The drug key of the program, the drug record the trials name (a salt form or a trial arm is its own record). | [optional] |
| **molecule_id** | **String** | The molecule key: the resolved drug the program&#39;s drug record folds to, and the entry key for the drugs object type. | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The condition&#39;s name. | [optional] |
| **is_indication** | **Boolean** | True for a specific disease, false for a broad MeSH umbrella condition. | [optional] |
| **max_phase** | **Integer** | Highest phase reached across the program&#39;s trials: 0 not reported, 1 through 4 for phase 1 through phase 4 (the 0 to 4 scale; see caveats). | [optional] |
| **n_trials** | **Integer** | Distinct trials of the program, across all sponsors. | [optional] |
| **n_completed** | **Integer** | Of those, trials with status completed. | [optional] |
| **n_terminated** | **Integer** | Of those, trials with status terminated. | [optional] |
| **first_year** | **Integer** | Earliest plausible trial start year. | [optional] |
| **last_year** | **Integer** | Latest plausible trial completion or start year. | [optional] |
| **n_sponsor_companies** | **Integer** | Distinct lead-sponsor companies across the program&#39;s trials; an unresolved lead sponsor counts as itself. | [optional] |
| **dalys_millions** | **Float** | Disability-adjusted life years (DALYs, millions) of the condition&#39;s global burden, from the Global Burden of Disease (GBD) 2021 study of the Institute for Health Metrics and Evaluation (IHME). Null for a caller without the licence for this column; see withheld_columns. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ProgramsRow.new(
  program_id: null,
  drug_id: null,
  molecule_id: null,
  drug_name: null,
  condition_id: null,
  condition_name: null,
  is_indication: null,
  max_phase: null,
  n_trials: null,
  n_completed: null,
  n_terminated: null,
  first_year: null,
  last_year: null,
  n_sponsor_companies: null,
  dalys_millions: null
)
```


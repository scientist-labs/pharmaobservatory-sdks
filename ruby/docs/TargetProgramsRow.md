# PharmaObservatory::TargetProgramsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] |
| **target_symbol** | **String** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] |
| **drug_id** | **String** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **condition_id** | **String** | The condition as a MeSH descriptor id (D followed by digits), the id the conditions datasets use. | [optional] |
| **condition_name** | **String** | The condition&#39;s name. | [optional] |
| **max_phase** | **Integer** | The furthest phase the program reached, 0 to 4 (0 is phase not reported). | [optional] |
| **n_trials** | **Integer** | Distinct trials of the drug in the indication. | [optional] |
| **n_completed** | **Integer** | Of n_trials, those with status COMPLETED. | [optional] |
| **n_terminated_trials** | **Integer** | Of n_trials, those with status TERMINATED. | [optional] |
| **first_year** | **Integer** | The earliest plausible trial start year. | [optional] |
| **last_year** | **Integer** | The latest plausible trial completion or start year. | [optional] |
| **is_discontinued** | **Boolean** | True when the program has a terminated, withdrawn or suspended trial. | [optional] |
| **n_discontinued** | **Integer** | Distinct stopped trials (terminated, withdrawn or suspended); 0 for a live program. | [optional] |
| **n_terminated** | **Integer** | Of the stopped trials, those TERMINATED. | [optional] |
| **n_withdrawn** | **Integer** | Of the stopped trials, those WITHDRAWN. | [optional] |
| **n_suspended** | **Integer** | Of the stopped trials, those SUSPENDED. | [optional] |
| **max_phase_at_stop** | **Integer** | The furthest phase among the stopped trials, 0 to 4; null for a live program. | [optional] |
| **last_stop_year** | **Integer** | The latest year a stopped trial ended; null for a live program. | [optional] |
| **why_stopped_reasons** | **Array&lt;String&gt;** | The stop reasons registrants disclosed, verbatim; empty when none. | [optional] |
| **why_stopped_categories** | **Array&lt;String&gt;** | The normalized stop-reason categories (efficacy_futility, safety, funding and so on); empty when none. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetProgramsRow.new(
  target_id: null,
  target_symbol: null,
  drug_id: null,
  drug_name: null,
  condition_id: null,
  condition_name: null,
  max_phase: null,
  n_trials: null,
  n_completed: null,
  n_terminated_trials: null,
  first_year: null,
  last_year: null,
  is_discontinued: null,
  n_discontinued: null,
  n_terminated: null,
  n_withdrawn: null,
  n_suspended: null,
  max_phase_at_stop: null,
  last_stop_year: null,
  why_stopped_reasons: null,
  why_stopped_categories: null
)
```


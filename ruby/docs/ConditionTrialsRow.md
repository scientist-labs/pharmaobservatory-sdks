# PharmaObservatory::ConditionTrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The condition&#39;s MeSH heading. | [optional] |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **is_direct** | **Boolean** | True when the trial is tagged with this condition itself; false when it is tagged with a condition beneath it in the MeSH tree. | [optional] |
| **trial_title** | **String** | The trial&#39;s brief title. | [optional] |
| **study_type** | **String** | The registry&#39;s study type (INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS). | [optional] |
| **phase** | **String** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] |
| **phase_label** | **String** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial&#39;s phase column holds the registry code. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall recruitment status (COMPLETED, RECRUITING, TERMINATED and so on). | [optional] |
| **start_date** | **Date** | Trial start date as registered. | [optional] |
| **start_year** | **Integer** | Calendar year of the start date. | [optional] |
| **primary_completion_date** | **Date** | Primary completion date as registered. | [optional] |
| **last_update_posted_date** | **Date** | The date the registry last posted an update to the trial record. | [optional] |
| **enrollment** | **Integer** | Enrollment as registered. | [optional] |
| **were_results_reported** | **Boolean** | Whether results are posted on the registry. | [optional] |
| **sponsor_name** | **String** | The lead sponsor&#39;s name as filed. | [optional] |
| **sponsor_org_id** | **String** | Legal-entity level: the lead sponsor&#39;s organization record. | [optional] |
| **sponsor_company_id** | **String** | Company level: the lead sponsor&#39;s company (null when the sponsor did not resolve to a company). | [optional] |
| **sponsor_company_name** | **String** | The lead sponsor&#39;s company name. | [optional] |
| **sponsor_group_id** | **String** | Group level: the current ultimate parent of the lead sponsor&#39;s company; dated links are in company_members. | [optional] |
| **sponsor_group_name** | **String** | Display name of the lead sponsor&#39;s group. | [optional] |
| **sponsor_agency_class** | **String** | The funder class the registry declared for the sponsor (INDUSTRY, NIH for the US National Institutes of Health, FED, OTHER_GOV, NETWORK, OTHER, INDIV, AMBIG, UNKNOWN). UNKNOWN is the registry&#39;s own code for a class it does not state. Null when the trial has no lead sponsor row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ConditionTrialsRow.new(
  condition_id: null,
  condition_name: null,
  trial_id: null,
  is_direct: null,
  trial_title: null,
  study_type: null,
  phase: null,
  phase_label: null,
  overall_status: null,
  start_date: null,
  start_year: null,
  primary_completion_date: null,
  last_update_posted_date: null,
  enrollment: null,
  were_results_reported: null,
  sponsor_name: null,
  sponsor_org_id: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  sponsor_agency_class: null
)
```


# PharmaObservatory::DrugTrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **title** | **String** | The trial&#39;s brief title. | [optional] |
| **phase** | **String** | The registry&#39;s phase code. Null when the registry gives no phase; phase_label carries the display label for those. | [optional] |
| **phase_label** | **String** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). | [optional] |
| **phase_rank** | **Integer** | The development phase as a number for max and sort: 4 is Phase 4, 3 is Phase 3 or 2/3, 2 is Phase 2 or 1/2, 1 is Phase 1 or Early Phase 1, 0 is not reported. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall status. | [optional] |
| **overall_status_summary** | **String** | The overall status grouped as Completed, Ongoing, Terminated or Other. | [optional] |
| **study_type** | **String** | INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS. | [optional] |
| **start_date** | **Date** | Trial start date. | [optional] |
| **primary_completion_date** | **Date** | The date primary data collection ends (estimated for an ongoing trial). | [optional] |
| **completion_date** | **Date** | The trial&#39;s completion date (estimated for an ongoing trial). | [optional] |
| **enrollment** | **Integer** | Enrollment as registered for the whole trial. | [optional] |
| **has_results** | **Boolean** | True when the sponsor has posted results to the registry. | [optional] |
| **lead_sponsor_name** | **String** | The lead sponsor&#39;s name as filed. | [optional] |
| **lead_sponsor_org_id** | **String** | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional] |
| **lead_sponsor_company_id** | **String** | Company level: the lead sponsor&#39;s company (null when unresolved). | [optional] |
| **lead_sponsor_company_name** | **String** | Display name of the lead sponsor&#39;s company. | [optional] |
| **lead_sponsor_group_id** | **String** | Group level: the lead sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **lead_sponsor_group_name** | **String** | Display name of the lead sponsor&#39;s group. | [optional] |
| **condition_ids** | **Array&lt;String&gt;** | The condition keys (MeSH, Medical Subject Headings, descriptor ids) the trial studies; drug_trial_conditions has one row per specific indication. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugTrialsRow.new(
  molecule_id: null,
  trial_id: null,
  title: null,
  phase: null,
  phase_label: null,
  phase_rank: null,
  overall_status: null,
  overall_status_summary: null,
  study_type: null,
  start_date: null,
  primary_completion_date: null,
  completion_date: null,
  enrollment: null,
  has_results: null,
  lead_sponsor_name: null,
  lead_sponsor_org_id: null,
  lead_sponsor_company_id: null,
  lead_sponsor_company_name: null,
  lead_sponsor_group_id: null,
  lead_sponsor_group_name: null,
  condition_ids: null
)
```


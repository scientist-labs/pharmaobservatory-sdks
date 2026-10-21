# PharmaObservatory::SponsorTrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **sponsor_role** | **String** | lead or collaborator. | [optional] |
| **sponsor_name** | **String** | The sponsor name as filed. | [optional] |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the sponsor belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **sponsor_agency_class** | **String** | The funder class the registry declared for the sponsor (INDUSTRY, NIH for the US National Institutes of Health, FED, OTHER_GOV, NETWORK, OTHER, INDIV, AMBIG, UNKNOWN). UNKNOWN is the registry&#39;s own code for a class it does not state. | [optional] |
| **sponsor_resolution_method** | **String** | How the sponsor name was resolved to an organization record. | [optional] |
| **trial_title** | **String** | The trial&#39;s brief title. | [optional] |
| **study_type** | **String** | INTERVENTIONAL or OBSERVATIONAL. | [optional] |
| **phase** | **String** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] |
| **phase_label** | **String** | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial&#39;s phase column holds the registry code. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall status. | [optional] |
| **start_date** | **Date** | Trial start date. | [optional] |
| **start_year** | **Integer** | Year of the trial start. | [optional] |
| **primary_completion_date** | **Date** | Primary completion date. | [optional] |
| **enrollment** | **Integer** | Enrollment as registered. | [optional] |
| **were_results_reported** | **Boolean** | Whether results are posted. | [optional] |
| **conditions** | **Array&lt;String&gt;** | Conditions studied, as registered. | [optional] |
| **interventions** | **Array&lt;String&gt;** | Interventions, as registered. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SponsorTrialsRow.new(
  trial_id: null,
  sponsor_role: null,
  sponsor_name: null,
  sponsor_org_id: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  sponsor_agency_class: null,
  sponsor_resolution_method: null,
  trial_title: null,
  study_type: null,
  phase: null,
  phase_label: null,
  overall_status: null,
  start_date: null,
  start_year: null,
  primary_completion_date: null,
  enrollment: null,
  were_results_reported: null,
  conditions: null,
  interventions: null
)
```


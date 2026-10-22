# PharmaObservatory::TrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s id: the ClinicalTrials.gov registry id (NCT, National Clinical Trial, number), or the WHO TrialID of a trial registered only with a WHO ICTRP member registry. | [optional] |
| **nct_id** | **String** | The ClinicalTrials.gov registry id; null for a trial registered only with a WHO ICTRP member registry. | [optional] |
| **source_registry** | **String** | The registry source of the row: clinicaltrials or ictrp (a trial registered only with a WHO ICTRP member registry). | [optional] |
| **registry** | **String** | The registry that holds the record: NCT for ClinicalTrials.gov, or a WHO ICTRP member registry (ChiCTR, ISRCTN, CTIS, ACTRN, NL, KCT, EUCTR and so on). | [optional] |
| **restriction_tier** | **String** | none for a ClinicalTrials.gov trial; licensable for a WHO ICTRP trial, which only a caller holding the licensable tier receives. | [optional] |
| **brief_title** | **String** | The trial&#39;s brief title. | [optional] |
| **official_title** | **String** | The trial&#39;s official title. | [optional] |
| **study_type** | **String** | INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall status code (RECRUITING, COMPLETED, TERMINATED and so on). | [optional] |
| **overall_status_summary** | **String** | A coarse status label derived from the overall status (Recruiting, Completed, Discontinued and so on). | [optional] |
| **phase** | **String** | The registry&#39;s phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional] |
| **phase_label** | **String** | The trial phase as a label (Phase 2, Phase 3, Observational, Not applicable). | [optional] |
| **phase_by_status** | **String** | The phase code and the status summary joined with a colon (PHASE2:Completed). | [optional] |
| **enrollment** | **Integer** | Enrollment as registered: a planned target or an achieved count (see enrollment_type). | [optional] |
| **enrollment_type** | **String** | ACTUAL (a count of enrolled participants) or ESTIMATED (a planned target). | [optional] |
| **enrollment_is_plausible** | **Boolean** | False when the registered enrollment is a placeholder value rather than a count. | [optional] |
| **enrollment_flag_reason** | **String** | Why the registered enrollment was marked a placeholder (null when it was not). | [optional] |
| **start_date** | **Date** | Trial start date. | [optional] |
| **start_year** | **Integer** | Year of the trial start. | [optional] |
| **primary_completion_date** | **Date** | Primary completion date. | [optional] |
| **primary_completion_date_type** | **String** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] |
| **completion_date** | **Date** | Completion date. | [optional] |
| **completion_date_type** | **String** | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry&#39;s older ANTICIPATED is published as ESTIMATED. | [optional] |
| **study_first_posted_date** | **Date** | Date the study was first posted to the registry. | [optional] |
| **last_update_posted_date** | **Date** | Date the registry record was last updated. | [optional] |
| **why_stopped** | **String** | The sponsor&#39;s stated reason a trial was stopped early, as filed. | [optional] |
| **number_of_arms** | **Integer** | Number of study arms. | [optional] |
| **were_results_reported** | **Boolean** | Whether results are posted to the registry. | [optional] |
| **gender** | **String** | Eligible sex: ALL, FEMALE or MALE. | [optional] |
| **minimum_age** | **String** | Minimum eligible age, as filed (for example 18 Years). | [optional] |
| **maximum_age** | **String** | Maximum eligible age, as filed (for example 65 Years). | [optional] |
| **healthy_volunteers** | **Boolean** | Whether healthy volunteers are accepted. | [optional] |
| **eligibility_criteria** | **String** | The inclusion and exclusion criteria text, as filed. | [optional] |
| **lead_sponsor_name** | **String** | The lead sponsor name as filed. | [optional] |
| **lead_sponsor_org_id** | **String** | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional] |
| **lead_sponsor_company_id** | **String** | Company level: the company the lead sponsor belongs to (null when unresolved). | [optional] |
| **lead_sponsor_company_name** | **String** | Display name of the lead sponsor&#39;s company. | [optional] |
| **lead_sponsor_group_id** | **String** | Group level: the lead sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **lead_sponsor_group_name** | **String** | Display name of the lead sponsor&#39;s group. | [optional] |
| **collaborators** | **Array&lt;String&gt;** | Collaborating sponsor names, as filed. | [optional] |
| **participating_sponsors** | **Array&lt;String&gt;** | Names of the lead sponsor and the collaborators, as filed. | [optional] |
| **conditions** | **Array&lt;String&gt;** | Conditions studied, as registered. | [optional] |
| **condition_mesh** | **Array&lt;String&gt;** | Conditions studied, as Medical Subject Headings (MeSH) terms the registry assigns. | [optional] |
| **interventions** | **Array&lt;String&gt;** | Interventions, as registered. | [optional] |
| **intervention_mesh** | **Array&lt;String&gt;** | Interventions as Medical Subject Headings (MeSH) terms the registry assigns. | [optional] |
| **intervention_types** | **Array&lt;String&gt;** | Intervention types (DRUG, BIOLOGICAL, DEVICE and so on). | [optional] |
| **number_of_facilities** | **Integer** | Number of study sites the registry record states. | [optional] |
| **facility_count** | **Integer** | Number of study sites listed on the record. | [optional] |
| **has_us_facility** | **Boolean** | Whether any listed site is in the United States. | [optional] |
| **facility_country_count** | **Integer** | Number of countries with a listed site. | [optional] |
| **facility_countries** | **Array&lt;String&gt;** | Countries with a listed site. | [optional] |
| **condition_ids** | **Array&lt;String&gt;** | MeSH condition ids the trial is linked to, direct and through the MeSH tree. Filter on one id to match the trials that carry it. | [optional] |
| **drug_ids** | **Array&lt;String&gt;** | Drug ids the trial&#39;s interventions resolved to. Filter on one id to match the trials that carry it. | [optional] |
| **sponsor_org_ids** | **Array&lt;String&gt;** | Organization record ids of the sponsors on the trial, lead and collaborator. Filter on one id to match the trials that carry it. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialsRow.new(
  trial_id: null,
  nct_id: null,
  source_registry: null,
  registry: null,
  restriction_tier: null,
  brief_title: null,
  official_title: null,
  study_type: null,
  overall_status: null,
  overall_status_summary: null,
  phase: null,
  phase_label: null,
  phase_by_status: null,
  enrollment: null,
  enrollment_type: null,
  enrollment_is_plausible: null,
  enrollment_flag_reason: null,
  start_date: null,
  start_year: null,
  primary_completion_date: null,
  primary_completion_date_type: null,
  completion_date: null,
  completion_date_type: null,
  study_first_posted_date: null,
  last_update_posted_date: null,
  why_stopped: null,
  number_of_arms: null,
  were_results_reported: null,
  gender: null,
  minimum_age: null,
  maximum_age: null,
  healthy_volunteers: null,
  eligibility_criteria: null,
  lead_sponsor_name: null,
  lead_sponsor_org_id: null,
  lead_sponsor_company_id: null,
  lead_sponsor_company_name: null,
  lead_sponsor_group_id: null,
  lead_sponsor_group_name: null,
  collaborators: null,
  participating_sponsors: null,
  conditions: null,
  condition_mesh: null,
  interventions: null,
  intervention_mesh: null,
  intervention_types: null,
  number_of_facilities: null,
  facility_count: null,
  has_us_facility: null,
  facility_country_count: null,
  facility_countries: null,
  condition_ids: null,
  drug_ids: null,
  sponsor_org_ids: null
)
```


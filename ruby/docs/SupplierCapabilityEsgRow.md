# PharmaObservatory::SupplierCapabilityEsgRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **vendor_org_id** | **String** | Legal-entity level: the contract research organization&#39;s organization record. | [optional] |
| **vendor_name** | **String** | The organization&#39;s name as filed. | [optional] |
| **vendor_company_id** | **String** | Company level: the contract research organization&#39;s company (null when unresolved). | [optional] |
| **vendor_company_name** | **String** | Display name of the contract research organization&#39;s company. | [optional] |
| **vendor_group_id** | **String** | Group level: the contract research organization&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **vendor_group_name** | **String** | Display name of the contract research organization&#39;s group. | [optional] |
| **cro_attribution_method** | **String** | How the vendor was inferred for the trial: sponsor_match or collaborator_match (the organization is named on the registry record and is on the curated list) or llm_extraction. | [optional] |
| **cro_confidence** | **Float** | Confidence of the inference, 0 to 1. | [optional] |
| **phase** | **String** | The registry&#39;s phase code of the trial (PHASE1, PHASE2, PHASE3 and so on). Null when the registry gives none. | [optional] |
| **start_year** | **Integer** | Start year of the trial. | [optional] |
| **lead_sponsor_name** | **String** | The trial&#39;s lead sponsor as filed. | [optional] |
| **facility_countries** | **Array&lt;String&gt;** | Countries of the trial&#39;s facilities. Filter on one country (for example facility_countries&#x3D;United States) to match trials that list it. | [optional] |
| **condition_ids** | **Array&lt;String&gt;** | The MeSH (Medical Subject Headings) condition ids of the trial, direct and ancestor (D followed by six to nine digits). Filter on one id (for example condition_ids&#x3D;D003316) to scope the rows to a condition. | [optional] |
| **enrollment_delta_known** | **Boolean** | True when the trial has both a planned and an actual enrollment, so its attainment is graded. | [optional] |
| **enrollment_met** | **Boolean** | True when actual enrollment reached the plan. Null when not graded. | [optional] |
| **enrollment_attainment_pct** | **Float** | Actual enrollment as a percentage of planned. Null when not graded. | [optional] |
| **climate_target_state** | **String** | The vendor&#39;s climate target state in the open lane (extracted from its SEC filings): disclosed when a net-zero or near-term target is claimed, true_negative when only a no is filed, not_disclosed otherwise, including a vendor the lane does not cover. The same value on every row of the vendor. | [optional] |
| **diversity_target_state** | **String** | The vendor&#39;s diversity target state in the open lane, over the has_dei_targets claim, by the same rule. | [optional] |
| **esg_claim_count** | **Integer** | The vendor&#39;s extracted claims (any type) that carry a value. | [optional] |
| **esg_latest_disclosure_date** | **Date** | The latest disclosure date among the vendor&#39;s extracted claims. | [optional] |
| **sbti_climate_target_state** | **String** | The vendor&#39;s climate target state in the Science Based Targets initiative lane, by the same rule. Null for a caller without the licence for this column; see withheld_columns. | [optional] |
| **sbti_claim_count** | **Integer** | The vendor&#39;s SBTi claims that carry a value. Null for a caller without the licence for this column; see withheld_columns. | [optional] |
| **sbti_latest_disclosure_date** | **Date** | The latest disclosure date among the vendor&#39;s SBTi claims. Null for a caller without the licence for this column; see withheld_columns. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SupplierCapabilityEsgRow.new(
  trial_id: null,
  vendor_org_id: null,
  vendor_name: null,
  vendor_company_id: null,
  vendor_company_name: null,
  vendor_group_id: null,
  vendor_group_name: null,
  cro_attribution_method: null,
  cro_confidence: null,
  phase: null,
  start_year: null,
  lead_sponsor_name: null,
  facility_countries: null,
  condition_ids: null,
  enrollment_delta_known: null,
  enrollment_met: null,
  enrollment_attainment_pct: null,
  climate_target_state: null,
  diversity_target_state: null,
  esg_claim_count: null,
  esg_latest_disclosure_date: null,
  sbti_climate_target_state: null,
  sbti_claim_count: null,
  sbti_latest_disclosure_date: null
)
```


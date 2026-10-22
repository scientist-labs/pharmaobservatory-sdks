# PharmaObservatory::CroTrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **vendor_org_id** | **String** | Legal-entity level: the contract research organization&#39;s organization record, as filed. | [optional] |
| **vendor_name** | **String** | The organization&#39;s name as filed. | [optional] |
| **vendor_company_id** | **String** | Company level: the contract research organization&#39;s company (null when unresolved). Group by this column for the company leaderboard. | [optional] |
| **vendor_company_name** | **String** | Display name of the contract research organization&#39;s company. | [optional] |
| **vendor_group_id** | **String** | Group level: the contract research organization&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **vendor_group_name** | **String** | Display name of the contract research organization&#39;s group. | [optional] |
| **attribution_method** | **String** | How the vendor was inferred for the trial: sponsor_match or collaborator_match (the organization is named on the registry record and is on the curated list) or llm_extraction. | [optional] |
| **confidence** | **Float** | Confidence of the inference, 0 to 1. | [optional] |
| **phase** | **String** | The registry&#39;s phase code of the trial (PHASE1, PHASE2, PHASE3 and so on). Null when the registry gives none. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall status code of the trial (RECRUITING, COMPLETED, TERMINATED and so on). | [optional] |
| **start_year** | **Integer** | Start year of the trial. | [optional] |
| **lead_sponsor_name** | **String** | The trial&#39;s lead sponsor as filed. | [optional] |
| **facility_countries** | **Array&lt;String&gt;** | Countries of the trial&#39;s facilities. Filter on one country (for example facility_countries&#x3D;United States) to match trials that list it. | [optional] |
| **condition_ids** | **Array&lt;String&gt;** | The MeSH (Medical Subject Headings) condition ids of the trial, direct and ancestor (D followed by six to nine digits). Filter on one id (for example condition_ids&#x3D;D003316) to scope the rows to an indication. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CroTrialsRow.new(
  trial_id: null,
  vendor_org_id: null,
  vendor_name: null,
  vendor_company_id: null,
  vendor_company_name: null,
  vendor_group_id: null,
  vendor_group_name: null,
  attribution_method: null,
  confidence: null,
  phase: null,
  overall_status: null,
  start_year: null,
  lead_sponsor_name: null,
  facility_countries: null,
  condition_ids: null
)
```


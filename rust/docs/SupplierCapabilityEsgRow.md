# SupplierCapabilityEsgRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**vendor_org_id** | Option<**String**> | Legal-entity level: the contract research organization's organization record. | [optional]
**vendor_name** | Option<**String**> | The organization's name as filed. | [optional]
**vendor_company_id** | Option<**String**> | Company level: the contract research organization's company (null when unresolved). | [optional]
**vendor_company_name** | Option<**String**> | Display name of the contract research organization's company. | [optional]
**vendor_group_id** | Option<**String**> | Group level: the contract research organization's current ultimate parent; a root's group is itself. | [optional]
**vendor_group_name** | Option<**String**> | Display name of the contract research organization's group. | [optional]
**cro_attribution_method** | Option<**String**> | How the vendor was inferred for the trial: sponsor_match or collaborator_match (the organization is named on the registry record and is on the curated list) or llm_extraction. | [optional]
**cro_confidence** | Option<**f64**> | Confidence of the inference, 0 to 1. | [optional]
**phase** | Option<**String**> | The registry's phase code of the trial (PHASE1, PHASE2, PHASE3 and so on). Null when the registry gives none. | [optional]
**start_year** | Option<**i32**> | Start year of the trial. | [optional]
**lead_sponsor_name** | Option<**String**> | The trial's lead sponsor as filed. | [optional]
**facility_countries** | Option<**Vec<String>**> | Countries of the trial's facilities. Filter on one country (for example facility_countries=United States) to match trials that list it. | [optional]
**condition_ids** | Option<**Vec<String>**> | The MeSH (Medical Subject Headings) condition ids of the trial, direct and ancestor (D followed by six to nine digits). Filter on one id (for example condition_ids=D003316) to scope the rows to a condition. | [optional]
**enrollment_delta_known** | Option<**bool**> | True when the trial has both a planned and an actual enrollment, so its attainment is graded. | [optional]
**enrollment_met** | Option<**bool**> | True when actual enrollment reached the plan. Null when not graded. | [optional]
**enrollment_attainment_pct** | Option<**f64**> | Actual enrollment as a percentage of planned. Null when not graded. | [optional]
**climate_target_state** | Option<**String**> | The vendor's climate target state in the open lane (extracted from its SEC filings): disclosed when a net-zero or near-term target is claimed, true_negative when only a no is filed, not_disclosed otherwise, including a vendor the lane does not cover. The same value on every row of the vendor. | [optional]
**diversity_target_state** | Option<**String**> | The vendor's diversity target state in the open lane, over the has_dei_targets claim, by the same rule. | [optional]
**esg_claim_count** | Option<**i32**> | The vendor's extracted claims (any type) that carry a value. | [optional]
**esg_latest_disclosure_date** | Option<[**String**](string.md)> | The latest disclosure date among the vendor's extracted claims. | [optional]
**sbti_climate_target_state** | Option<**String**> | The vendor's climate target state in the Science Based Targets initiative lane, by the same rule. Null for a caller without the licence for this column; see withheld_columns. | [optional]
**sbti_claim_count** | Option<**i32**> | The vendor's SBTi claims that carry a value. Null for a caller without the licence for this column; see withheld_columns. | [optional]
**sbti_latest_disclosure_date** | Option<[**String**](string.md)> | The latest disclosure date among the vendor's SBTi claims. Null for a caller without the licence for this column; see withheld_columns. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



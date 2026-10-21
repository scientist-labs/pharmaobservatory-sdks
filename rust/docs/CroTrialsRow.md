# CroTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**vendor_org_id** | Option<**String**> | Legal-entity level: the contract research organization's organization record, as filed. | [optional]
**vendor_name** | Option<**String**> | The organization's name as filed. | [optional]
**vendor_company_id** | Option<**String**> | Company level: the contract research organization's company (null when unresolved). Group by this column for the company leaderboard. | [optional]
**vendor_company_name** | Option<**String**> | Display name of the contract research organization's company. | [optional]
**vendor_group_id** | Option<**String**> | Group level: the contract research organization's current ultimate parent; a root's group is itself. | [optional]
**vendor_group_name** | Option<**String**> | Display name of the contract research organization's group. | [optional]
**attribution_method** | Option<**String**> | How the vendor was inferred for the trial: sponsor_match or collaborator_match (the organization is named on the registry record and is on the curated list) or llm_extraction. | [optional]
**confidence** | Option<**f64**> | Confidence of the inference, 0 to 1. | [optional]
**phase** | Option<**String**> | The registry's phase code of the trial (PHASE1, PHASE2, PHASE3 and so on). Null when the registry gives none. | [optional]
**overall_status** | Option<**String**> | The registry's overall status code of the trial (RECRUITING, COMPLETED, TERMINATED and so on). | [optional]
**start_year** | Option<**i32**> | Start year of the trial. | [optional]
**lead_sponsor_name** | Option<**String**> | The trial's lead sponsor as filed. | [optional]
**facility_countries** | Option<**Vec<String>**> | Countries of the trial's facilities. Filter on one country (for example facility_countries=United States) to match trials that list it. | [optional]
**condition_ids** | Option<**Vec<String>**> | The MeSH (Medical Subject Headings) condition ids of the trial, direct and ancestor (D followed by six to nine digits). Filter on one id (for example condition_ids=D003316) to scope the rows to an indication. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



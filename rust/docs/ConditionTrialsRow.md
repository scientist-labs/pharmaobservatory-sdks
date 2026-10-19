# ConditionTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The condition's MeSH heading. | [optional]
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**is_direct** | Option<**bool**> | True when the trial is tagged with this condition itself; false when it is tagged with a condition beneath it in the MeSH tree. | [optional]
**trial_title** | Option<**String**> | The trial's brief title. | [optional]
**study_type** | Option<**String**> | The registry's study type (INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS). | [optional]
**phase** | Option<**String**> | The trial phase label (Phase 1, Phase 2, Phase 3, Phase 4, Not Applicable and so on). | [optional]
**overall_status** | Option<**String**> | The registry's overall recruitment status (COMPLETED, RECRUITING, TERMINATED and so on). | [optional]
**start_date** | Option<[**String**](string.md)> | Trial start date as registered. | [optional]
**start_year** | Option<**i32**> | Calendar year of the start date. | [optional]
**primary_completion_date** | Option<[**String**](string.md)> | Primary completion date as registered. | [optional]
**last_update_posted_date** | Option<[**String**](string.md)> | The date the registry last posted an update to the trial record. | [optional]
**enrollment** | Option<**i32**> | Enrollment as registered. | [optional]
**were_results_reported** | Option<**bool**> | Whether results are posted on the registry. | [optional]
**sponsor_name** | Option<**String**> | The lead sponsor's name as filed. | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the lead sponsor's organization record. | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the lead sponsor's company (null when the sponsor did not resolve to a company). | [optional]
**sponsor_company_name** | Option<**String**> | The lead sponsor's company name. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the current ultimate parent of the lead sponsor's company; dated links are in company_members. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the lead sponsor's group. | [optional]
**sponsor_agency_class** | Option<**String**> | The registry's declared funder class of the lead sponsor (INDUSTRY, NIH for the US National Institutes of Health, OTHER and so on). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



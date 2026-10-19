# SponsorTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's registry id (NCT number). | [optional]
**sponsor_role** | Option<**String**> | lead or collaborator. | [optional]
**sponsor_name** | Option<**String**> | The sponsor name as filed. | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the sponsor belongs to (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**sponsor_agency_class** | Option<**String**> | The registry's agency class for the sponsor (INDUSTRY, NIH, OTHER ...). | [optional]
**sponsor_resolution_method** | Option<**String**> | How the sponsor name was resolved to an organization record. | [optional]
**trial_title** | Option<**String**> | The trial's brief title. | [optional]
**study_type** | Option<**String**> | INTERVENTIONAL or OBSERVATIONAL. | [optional]
**phase** | Option<**String**> | The trial phase label. | [optional]
**overall_status** | Option<**String**> | The registry's overall status. | [optional]
**start_date** | Option<[**String**](string.md)> | Trial start date. | [optional]
**start_year** | Option<**i32**> | Year of the trial start. | [optional]
**primary_completion_date** | Option<[**String**](string.md)> | Primary completion date. | [optional]
**enrollment** | Option<**i32**> | Enrollment as registered. | [optional]
**were_results_reported** | Option<**bool**> | Whether results are posted. | [optional]
**conditions** | Option<**Vec<String>**> | Conditions studied, as registered. | [optional]
**interventions** | Option<**Vec<String>**> | Interventions, as registered. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



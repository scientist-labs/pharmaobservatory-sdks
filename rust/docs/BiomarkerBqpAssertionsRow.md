# BiomarkerBqpAssertionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source_assertion_id** | Option<**String**> | The assertion id. | [optional]
**route_key** | Option<**String**> | The biomarker's route key, the entry key for the biomarkers object type. | [optional]
**biomarker_id** | Option<**String**> | The biomarker key. | [optional]
**ddt_number** | Option<**String**> | The drug development tool number the US Food and Drug Administration (FDA) assigns the project. | [optional]
**ddt_project_id** | Option<**String**> | The FDA project id. | [optional]
**project_name** | Option<**String**> | The project's name. | [optional]
**bqp_signal** | Option<**String**> | The source status of the project. Only qualified means FDA-qualified; every other value is a project status. | [optional]
**is_fda_qualified** | Option<**bool**> | True only when the source status is exactly qualified. | [optional]
**submission_status** | Option<**String**> | The submission status as the source words it. | [optional]
**submission_stage** | Option<**String**> | The submission stage as the source words it. | [optional]
**submitter_name** | Option<**String**> | The submitter as the source names it. | [optional]
**submitter_org_id** | Option<**String**> | Legal-entity level: the organization record the submitter resolved to (null when unresolved). | [optional]
**submitter_company_id** | Option<**String**> | Company level: the company the submitter belongs to (null when unresolved). | [optional]
**submitter_company_name** | Option<**String**> | Display name of the submitter's company. | [optional]
**submitter_group_id** | Option<**String**> | Group level: the submitter's current ultimate parent company; a root company's group is itself. | [optional]
**submitter_group_name** | Option<**String**> | Display name of the submitter's group. | [optional]
**submitter_match_tier** | Option<**String**> | How the submitter name matched an organization record. | [optional]
**submitter_role** | Option<**String**> | The submitter's role on the project. | [optional]
**biomarker_role** | Option<**String**> | The biomarker's role in the project (for example a diagnostic or a surrogate endpoint). | [optional]
**biomarker_role_method** | Option<**String**> | How the biomarker role was assigned. | [optional]
**is_surrogate_endpoint_claim** | Option<**bool**> | True when the project claims a surrogate endpoint. | [optional]
**source_url** | Option<**String**> | Link to the FDA project page. | [optional]
**snapshot_date** | Option<[**String**](string.md)> | The date the source list was captured. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



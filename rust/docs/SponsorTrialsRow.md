# SponsorTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**sponsor_role** | Option<**String**> | lead or collaborator. | [optional]
**sponsor_name** | Option<**String**> | The sponsor name as filed. | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the sponsor belongs to (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**sponsor_agency_class** | Option<**String**> | The funder class the registry declared for the sponsor (INDUSTRY, NIH for the US National Institutes of Health, FED, OTHER_GOV, NETWORK, OTHER, INDIV, AMBIG, UNKNOWN). UNKNOWN is the registry's own code for a class it does not state. | [optional]
**sponsor_resolution_method** | Option<**String**> | How the sponsor name was resolved to an organization record. | [optional]
**trial_title** | Option<**String**> | The trial's brief title. | [optional]
**study_type** | Option<**String**> | INTERVENTIONAL or OBSERVATIONAL. | [optional]
**phase** | Option<**String**> | The registry's phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional]
**phase_label** | Option<**String**> | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial's phase column holds the registry code. | [optional]
**overall_status** | Option<**String**> | The registry's overall status. | [optional]
**start_date** | Option<[**String**](string.md)> | Trial start date. | [optional]
**start_year** | Option<**i32**> | Year of the trial start. | [optional]
**primary_completion_date** | Option<[**String**](string.md)> | Primary completion date. | [optional]
**enrollment** | Option<**i32**> | Enrollment as registered. | [optional]
**were_results_reported** | Option<**bool**> | Whether results are posted. | [optional]
**conditions** | Option<**Vec<String>**> | Conditions studied, as registered. | [optional]
**interventions** | Option<**Vec<String>**> | Interventions, as registered. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



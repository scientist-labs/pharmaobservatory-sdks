# SiteTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**site_id** | Option<**String**> | The site key: one facility spelling from the registry. | [optional]
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**trial_title** | Option<**String**> | The trial's brief title. | [optional]
**study_type** | Option<**String**> | INTERVENTIONAL or OBSERVATIONAL. | [optional]
**phase** | Option<**String**> | The registry's phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional]
**phase_label** | Option<**String**> | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial's phase column holds the registry code. | [optional]
**overall_status** | Option<**String**> | The registry's overall status. | [optional]
**start_date** | Option<[**String**](string.md)> | Trial start date. | [optional]
**start_year** | Option<**i32**> | Year of the trial start. | [optional]
**enrollment** | Option<**i32**> | Enrollment as registered for the whole trial. | [optional]
**lead_sponsor_name** | Option<**String**> | The lead sponsor's name as filed. | [optional]
**lead_sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional]
**lead_sponsor_company_id** | Option<**String**> | Company level: the company the lead sponsor belongs to (null when unresolved). | [optional]
**lead_sponsor_company_name** | Option<**String**> | Display name of the lead sponsor's company. | [optional]
**lead_sponsor_group_id** | Option<**String**> | Group level: the lead sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**lead_sponsor_group_name** | Option<**String**> | Display name of the lead sponsor's group. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DrugTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**title** | Option<**String**> | The trial's brief title. | [optional]
**phase** | Option<**String**> | The registry's phase code. Null when the registry gives no phase; phase_label carries the display label for those. | [optional]
**phase_label** | Option<**String**> | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). | [optional]
**phase_rank** | Option<**i32**> | The development phase as a number for max and sort: 4 is Phase 4, 3 is Phase 3 or 2/3, 2 is Phase 2 or 1/2, 1 is Phase 1 or Early Phase 1, 0 is not reported. | [optional]
**overall_status** | Option<**String**> | The registry's overall status. | [optional]
**overall_status_summary** | Option<**String**> | The overall status grouped as Completed, Ongoing, Terminated or Other. | [optional]
**study_type** | Option<**String**> | INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS. | [optional]
**start_date** | Option<[**String**](string.md)> | Trial start date. | [optional]
**primary_completion_date** | Option<[**String**](string.md)> | The date primary data collection ends (estimated for an ongoing trial). | [optional]
**completion_date** | Option<[**String**](string.md)> | The trial's completion date (estimated for an ongoing trial). | [optional]
**enrollment** | Option<**i32**> | Enrollment as registered for the whole trial. | [optional]
**has_results** | Option<**bool**> | True when the sponsor has posted results to the registry. | [optional]
**lead_sponsor_name** | Option<**String**> | The lead sponsor's name as filed. | [optional]
**lead_sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional]
**lead_sponsor_company_id** | Option<**String**> | Company level: the lead sponsor's company (null when unresolved). | [optional]
**lead_sponsor_company_name** | Option<**String**> | Display name of the lead sponsor's company. | [optional]
**lead_sponsor_group_id** | Option<**String**> | Group level: the lead sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**lead_sponsor_group_name** | Option<**String**> | Display name of the lead sponsor's group. | [optional]
**condition_ids** | Option<**Vec<String>**> | The condition keys (MeSH, Medical Subject Headings, descriptor ids) the trial studies; drug_trial_conditions has one row per specific indication. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



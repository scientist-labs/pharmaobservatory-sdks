# TargetTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**target_symbol** | Option<**String**> | The target's HGNC gene symbol (the id the target page address carries). | [optional]
**trial_id** | Option<**String**> | The ClinicalTrials.gov NCT id, the id the trials datasets use. | [optional]
**evidence_source** | Option<**String**> | The best evidence tier among the trial's drugs that act on the target. | [optional]
**brief_title** | Option<**String**> | The trial's brief title. | [optional]
**phase** | Option<**String**> | The registry's phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional]
**phase_label** | Option<**String**> | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial's phase column holds the registry code. | [optional]
**overall_status** | Option<**String**> | The registry's overall recruitment status. | [optional]
**start_year** | Option<**i32**> | The year the trial started. | [optional]
**lead_sponsor_name** | Option<**String**> | The lead sponsor's name as filed. | [optional]
**lead_sponsor_org_id** | Option<**String**> | Legal-entity level: the lead sponsor organization; null when it does not resolve. | [optional]
**lead_sponsor_company_id** | Option<**String**> | Company level: the lead sponsor's company; null when none. | [optional]
**lead_sponsor_company_name** | Option<**String**> | Display name of the lead sponsor's company. | [optional]
**lead_sponsor_group_id** | Option<**String**> | Group level: the lead sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**lead_sponsor_group_name** | Option<**String**> | Display name of the lead sponsor's group. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



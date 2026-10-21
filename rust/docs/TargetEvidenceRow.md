# TargetEvidenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_evidence_id** | Option<**String**> | The row id: a hash of the target, drug, indication, trial and sponsor organization. | [optional]
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**target_symbol** | Option<**String**> | The target's HGNC gene symbol (the id the target page address carries). | [optional]
**drug_id** | Option<**String**> | The drug's canonical id (the same id the drug datasets use as drug_id). | [optional]
**drug_name** | Option<**String**> | The drug's name. | [optional]
**approved_anywhere** | Option<**bool**> | Whether the drug is approved by any authority we hold. | [optional]
**primary_action_type** | Option<**String**> | The action type of the drug's best-evidence mechanism on the target; null where the source gives none. | [optional]
**evidence_source** | Option<**String**> | The best evidence tier that links the drug to the target: curated_moa (an Open Targets curated mechanism edge), chembl_moa (a ChEMBL mechanism) or intervention_seed (a curated investigational asset). | [optional]
**condition_id** | Option<**String**> | The condition as a MeSH descriptor id (D followed by digits), the id the conditions datasets use. | [optional]
**condition_name** | Option<**String**> | The condition's name. | [optional]
**trial_id** | Option<**String**> | The ClinicalTrials.gov NCT id, the id the trials datasets use. | [optional]
**max_phase** | Option<**i32**> | The furthest phase the program (drug and indication) reached, 0 to 4 (0 is phase not reported). | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the trial's lead sponsor organization as filed; null when it does not resolve. | [optional]
**sponsor_name** | Option<**String**> | The lead sponsor's name as filed. | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the lead sponsor organization belongs to; null when none. | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TargetDrugsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**target_symbol** | Option<**String**> | The target's HGNC gene symbol (the id the target page address carries). | [optional]
**drug_id** | Option<**String**> | The drug's canonical id (the same id the drug datasets use as drug_id). | [optional]
**molecule_id** | Option<**String**> | The molecule the drug resolves to, the id condition_drugs also calls molecule_id; equal to drug_id for a drug that resolves to no other. | [optional]
**drug_name** | Option<**String**> | The drug's name. | [optional]
**primary_action_type** | Option<**String**> | The action type of the best-evidence mechanism (INHIBITOR, AGONIST and so on); null where the source gives none. | [optional]
**evidence_source** | Option<**String**> | The best evidence tier that links the drug to the target: curated_moa (an Open Targets curated mechanism edge), chembl_moa (a ChEMBL mechanism) or intervention_seed (a curated investigational asset). | [optional]
**evidence_rank** | Option<**i32**> | The rank of evidence_source, 1 for curated_moa, 2 for chembl_moa, 3 for intervention_seed. | [optional]
**evidence_sources** | Option<**Vec<String>**> | Every evidence tier that links the drug to the target; more than one is a corroboration signal. | [optional]
**approved_anywhere** | Option<**bool**> | Whether the drug is approved by any authority we hold. | [optional]
**drug_n_trials** | Option<**f64**> | Distinct trials of the drug, on any target. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



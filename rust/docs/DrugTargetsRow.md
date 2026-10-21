# DrugTargetsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**ensembl_id** | Option<**String**> | The target's Ensembl gene id: the entry key of the targets object type. | [optional]
**symbol** | Option<**String**> | The gene symbol: the route key of the biomarker page for the same gene. | [optional]
**target_name** | Option<**String**> | The target's name. | [optional]
**uniprot** | Option<**String**> | The UniProt accession of the target protein. | [optional]
**has_pdb** | Option<**bool**> | True when the Protein Data Bank holds a structure for the protein. | [optional]
**primary_action_type** | Option<**String**> | The mechanism of action (for example INHIBITOR, AGONIST) from the best-ranked evidence that states one. | [optional]
**evidence_source** | Option<**String**> | The best-ranked evidence source for the link. | [optional]
**evidence_rank** | Option<**i32**> | The rank of that source; lower is stronger. | [optional]
**evidence_sources** | Option<**Vec<String>**> | Every evidence source for the link. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



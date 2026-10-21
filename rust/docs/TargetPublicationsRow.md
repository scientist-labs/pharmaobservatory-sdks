# TargetPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**target_symbol** | Option<**String**> | The target's HGNC gene symbol (the id the target page address carries). | [optional]
**pmid** | Option<**String**> | The PubMed id. | [optional]
**citation** | Option<**String**> | The citation string registrants gave. | [optional]
**title** | Option<**String**> | The paper's title from PubMed; null when PubMed has no record. | [optional]
**pub_year** | Option<**i32**> | The publication year. | [optional]
**n_linked_trials** | Option<**i32**> | Distinct trials in the registry that cite the paper, for any target. | [optional]
**n_target_trials** | Option<**i32**> | Distinct trials of this target that cite the paper. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



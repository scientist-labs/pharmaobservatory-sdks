# TargetRegulatoryPrecedentRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**target_symbol** | Option<**String**> | The target's HGNC gene symbol (the id the target page address carries). | [optional]
**drug_id** | Option<**String**> | The drug's canonical id (the same id the drug datasets use as drug_id). | [optional]
**drug_name** | Option<**String**> | The drug's name. | [optional]
**authority** | Option<**String**> | The approving authority. | [optional]
**jurisdiction** | Option<**String**> | The authority's jurisdiction. | [optional]
**status** | Option<**String**> | The approval status the authority records. | [optional]
**first_approval_date** | Option<[**String**](string.md)> | The authority's first approval date of the drug. | [optional]
**approval_id** | Option<**String**> | The authority's own approval identifier. | [optional]
**primary_action_type** | Option<**String**> | The action type of the drug's curated mechanism on the target. | [optional]
**moa_evidence_source** | Option<**String**> | The evidence tier of the mechanism; curated_moa on every row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



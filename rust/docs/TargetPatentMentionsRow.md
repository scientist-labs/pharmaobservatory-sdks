# TargetPatentMentionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**target_symbol** | Option<**String**> | The target's HGNC gene symbol (the id the target page address carries). | [optional]
**drug_id** | Option<**String**> | The drug's canonical id (the same id the drug datasets use as drug_id). | [optional]
**drug_name** | Option<**String**> | The drug's name. | [optional]
**patent_number** | Option<**String**> | The patent number. | [optional]
**country** | Option<**String**> | The patent office's country code. | [optional]
**publication_date** | Option<[**String**](string.md)> | The patent's publication date. | [optional]
**title** | Option<**String**> | The patent's title. | [optional]
**assignee** | Option<**String**> | The patent's assignee as filed; null where none. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



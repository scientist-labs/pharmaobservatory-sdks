# TargetAssociationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**datasource** | Option<**String**> | The Open Targets datasource name (for example genetic_association or literature). | [optional]
**n_disease_associations** | Option<**i64**> | Distinct disease associations the datasource holds for the target. | [optional]
**min_association_score** | Option<**f64**> | The lowest association score in the datasource for the target. | [optional]
**max_association_score** | Option<**f64**> | The highest association score in the datasource for the target. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



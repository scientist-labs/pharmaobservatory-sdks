# TargetCrowdingRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**target_symbol** | Option<**String**> | The target's HGNC gene symbol (the id the target page address carries). | [optional]
**max_phase** | Option<**i32**> | The band: the furthest phase a program reached, 0 to 4 (0 is phase not reported). | [optional]
**n_programs** | Option<**i32**> | Distinct drug and indication programs whose furthest phase is the band. | [optional]
**n_active_programs** | Option<**i32**> | Programs in the band with no terminated, withdrawn or suspended trial. | [optional]
**n_discontinued** | Option<**i32**> | Programs in the band with at least one terminated, withdrawn or suspended trial. | [optional]
**n_drugs** | Option<**i32**> | Distinct drugs among the band's programs. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BiomarkerTraversalStatusRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**route_key** | Option<**String**> | The biomarker's route key, the entry key for the biomarkers object type. | [optional]
**biomarker_id** | Option<**String**> | The biomarker's own key. | [optional]
**biomarker_kind** | Option<**String**> | What kind of biomarker the row describes. | [optional]
**traversal_key** | Option<**String**> | The traversal. | [optional]
**display_order** | Option<**i32**> | The order the web page lists the traversals in, one to nine. | [optional]
**traversal_name** | Option<**String**> | The traversal's display name. | [optional]
**applicability_state** | Option<**String**> | applicable: measured for this kind of biomarker. inapplicable: does not apply to this kind. not_measured: applies in principle but the corpus is not mined for this kind yet. | [optional]
**context_mode** | Option<**String**> | How the traversal reads for this biomarker. direct: its own rows. parent_gene: the variant's gene rows. parent_and_exact: the gene's rows plus the exact variant's. child_variants: the gene's variant children. source_linked: shown when a source assertion resolves to it. bridge_deferred: depends on a bridge that has not landed. none: no traversal. | [optional]
**reason_code** | Option<**String**> | The code behind the applicability state and context mode; state_detail is its sentence. | [optional]
**state_detail** | Option<**String**> | The sentence the web page prints for the reason code. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



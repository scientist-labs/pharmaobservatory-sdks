# CellLinesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cell_line_id** | Option<**String**> | The catalog listing's model id (scientist_model_id). | [optional]
**alternate_cell_line_ids** | Option<**String**> | Other catalog model ids for the same line. | [optional]
**model_lot_id** | Option<**String**> | The lot id. | [optional]
**cell_type** | Option<**String**> | The cell line's name. | [optional]
**provider_id** | Option<**String**> | The supplier's catalog id. | [optional]
**supplier** | Option<**String**> | The supplier as listed. | [optional]
**sample_species** | Option<**String**> | The species. | [optional]
**storefront_enable** | Option<**bool**> | Whether the listing is enabled in the storefront. | [optional]
**mutation_data** | Option<**String**> | Mutation data as listed. | [optional]
**status** | Option<**String**> | The listing status. | [optional]
**safety_status** | Option<**String**> | confirmed_problematic, review_candidate, no_problematic_match or missing_cell_type. | [optional]
**resolution_status** | Option<**String**> | resolved, unresolved_review, no_match or missing_input. | [optional]
**confirmed_problematic** | Option<**bool**> | Whether the line is confirmed problematic in Cellosaurus. | [optional]
**cellosaurus_accession** | Option<**String**> | The displayed Cellosaurus accession. | [optional]
**cellosaurus_url** | Option<**String**> | Link to the Cellosaurus entry. | [optional]
**problem_category** | Option<**String**> | The problem category, for a problematic line. | [optional]
**problematic_comment** | Option<**String**> | The Cellosaurus comment, for a problematic line. | [optional]
**iclac_registration_ids** | Option<**Vec<String>**> | ICLAC registration ids. | [optional]
**release_version** | Option<**String**> | The Cellosaurus release the check ran against. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



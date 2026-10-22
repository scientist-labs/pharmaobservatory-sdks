# RwdWithheldFieldsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**register_key** | Option<**String**> | Which of the five exports the row concerns: study, data_source_human, data_source_veterinary, institution or network. The entry key of the rwd_registers object type. | [optional]
**register_label** | Option<**String**> | The register's plain-words label. | [optional]
**source_ordinal** | Option<**i32**> | The withheld column's position in the export. | [optional]
**source_header** | Option<**String**> | The export's header for the withheld column. | [optional]
**column_name** | Option<**String**> | The name of the withheld column in the lane's landed schema. | [optional]
**withheld_reason** | Option<**String**> | Why the column is not published. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



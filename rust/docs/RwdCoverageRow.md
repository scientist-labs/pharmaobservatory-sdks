# RwdCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**chain_key** | Option<**String**> | Which denominator chain the row belongs to. | [optional]
**chain_label** | Option<**String**> | The chain's plain-words label. | [optional]
**chain_is_a_partition** | Option<**bool**> | True when the chain's steps partition the denominator (they sum to it); false when steps overlap. | [optional]
**step_ordinal** | Option<**i32**> | The step's position in the chain; 0 is the denominator itself, published as a step so the chain reads top to bottom. | [optional]
**step_key** | Option<**String**> | The step's key. | [optional]
**step_label** | Option<**String**> | The step's plain-words label. | [optional]
**register_key** | Option<**String**> | Which of the five exports the row concerns: study, data_source_human, data_source_veterinary, institution or network. The entry key of the rwd_registers object type. | [optional]
**measured_value** | Option<**i64**> | The measured count at this step. | [optional]
**denominator_key** | Option<**String**> | Which denominator the step is read against. | [optional]
**denominator_label** | Option<**String**> | The denominator's plain-words label. | [optional]
**denominator_value** | Option<**i64**> | The denominator's value. | [optional]
**share_of_denominator** | Option<**f64**> | measured_value over denominator_value, as a fraction. | [optional]
**derivation_note** | Option<**String**> | How the step is derived. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**snapshot_date** | Option<[**String**](string.md)> | The date of the catalogue snapshot. | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



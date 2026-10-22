# RareDiseaseRarityRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**rarity_assertion_id** | Option<**String**> | The identifier of the assessment. | [optional]
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**jurisdiction_code** | Option<**String**> | The jurisdiction whose definition was applied (GB or EU). | [optional]
**jurisdiction_name** | Option<**String**> | The jurisdiction's name. | [optional]
**rule_id** | Option<**String**> | The identifier of the rule applied. | [optional]
**definition_text** | Option<**String**> | The rule's definition as stated by the jurisdiction. | [optional]
**comparison_operator** | Option<**String**> | The comparison the rule applies to the threshold (< or <=). | [optional]
**threshold_value** | Option<**f64**> | The threshold figure, to be read over population_denominator (for example 5 per 10000). | [optional]
**population_denominator** | Option<**f64**> | The denominator of the threshold figure. | [optional]
**population_unit** | Option<**String**> | The unit of the denominator. | [optional]
**effective_date** | Option<[**String**](string.md)> | The date the rule took effect. | [optional]
**end_date** | Option<[**String**](string.md)> | The date the rule ended, where it has. | [optional]
**assessment_date** | Option<[**String**](string.md)> | The date of the assessment. | [optional]
**rarity_status** | Option<**String**> | rare, not_rare, unknown or conflicting, as a result of applying the rule to the evidence. | [optional]
**evidence_basis** | Option<**String**> | What the assessment rests on: the bounds of a source prevalence class, or no evidence in the jurisdiction's geography. | [optional]
**evidence_count** | Option<**i64**> | Prevalence observations compared in the assessment. | [optional]
**official_source_url** | Option<**String**> | The jurisdiction's official document the rule comes from. | [optional]
**license_name** | Option<**String**> | The licence of the jurisdiction's document. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



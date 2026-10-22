# RareDiseaseRarityEvidenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**rarity_evidence_id** | Option<**String**> | The identifier of the evidence link. | [optional]
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**rarity_assertion_id** | Option<**String**> | The assessment the evidence belongs to. | [optional]
**prevalence_assertion_id** | Option<**String**> | The prevalence observation compared. | [optional]
**jurisdiction_code** | Option<**String**> | The jurisdiction whose rule was applied (GB or EU). | [optional]
**evidence_kind** | Option<**String**> | How the observation was compared (prevalence_class_ratio). | [optional]
**comparison_result** | Option<**String**> | The result of comparing the observation with the threshold: rare, not_rare or unknown. | [optional]
**is_comparable** | Option<**bool**> | True when the observation could be compared with the threshold. | [optional]
**comparison_reason** | Option<**String**> | The reason the comparison gave its result. | [optional]
**assessment_date** | Option<[**String**](string.md)> | The date of the assessment. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



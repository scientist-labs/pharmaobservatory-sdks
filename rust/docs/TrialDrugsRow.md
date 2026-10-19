# TrialDrugsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**drug_key** | Option<**String**> | The molecule key, or the drug name as registered prefixed unresolved: for a drug that reached no molecule. | [optional]
**molecule_id** | Option<**String**> | The molecule key the drugs resolved to (null for a drug with no molecule). | [optional]
**molecule_name** | Option<**String**> | The molecule's display name (null for a drug with no molecule). | [optional]
**registered_as** | Option<**Vec<String>**> | The drug names as registered on the trial that led to this row. | [optional]
**matched_from_name_part** | Option<**bool**> | Whether every match to the molecule came from one part of a multi-part drug name. | [optional]
**brand_count** | Option<**i32**> | Number of brand names recorded for the molecule. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



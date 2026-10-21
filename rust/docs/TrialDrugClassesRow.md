# TrialDrugClassesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**drug_id** | Option<**String**> | The canonical drug key the trial's intervention names matched. | [optional]
**drug_name** | Option<**String**> | The drug's display name. | [optional]
**resolved_drug_id** | Option<**String**> | The drug key after same-substance drug records are merged. | [optional]
**approved_anywhere** | Option<**bool**> | True when the drug is marked approved by any regulator (false means investigational or unknown). | [optional]
**atc_class** | Option<**String**> | The ATC first-level class of the drug (null when it has none). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



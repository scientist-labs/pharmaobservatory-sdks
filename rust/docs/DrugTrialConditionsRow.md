# DrugTrialConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**condition_id** | Option<**String**> | The condition key (a MeSH, Medical Subject Headings, descriptor id). | [optional]
**condition_name** | Option<**String**> | The condition's name. | [optional]
**phase** | Option<**String**> | The registry's phase code of the trial. | [optional]
**phase_rank** | Option<**i32**> | The trial's phase as a number for max and sort (4 is Phase 4, 0 is not reported). | [optional]
**overall_status_summary** | Option<**String**> | The trial's overall status grouped as Completed, Ongoing, Terminated or Other. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



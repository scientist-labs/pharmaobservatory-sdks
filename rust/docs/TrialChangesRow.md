# TrialChangesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**observed_at** | Option<**String**> | When the platform observed the change. | [optional]
**event_type** | Option<**String**> | The kind of event (first_seen, status_change and so on). | [optional]
**field** | Option<**String**> | The record field that changed (null for a first_seen event). | [optional]
**old_value** | Option<**String**> | The value before the change. | [optional]
**new_value** | Option<**String**> | The value after the change. | [optional]
**trial_title** | Option<**String**> | The trial's brief title. | [optional]
**condition_ids** | Option<**Vec<String>**> | MeSH (Medical Subject Headings) condition ids the trial is linked to, direct and through the MeSH tree. Filter on one id to scope changes to a condition's trials. | [optional]
**drug_ids** | Option<**Vec<String>**> | Drug ids the trial's interventions resolved to. Filter on one id to scope changes to a drug's trials. | [optional]
**sponsor_org_ids** | Option<**Vec<String>**> | Organization record ids of the sponsors on the trial, lead and collaborator. Filter on one id to scope changes to a sponsor's trials. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



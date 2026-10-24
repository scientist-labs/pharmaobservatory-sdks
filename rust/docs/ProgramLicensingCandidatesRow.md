# ProgramLicensingCandidatesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sponsor_org_id** | Option<**String**> | Legal-entity level: the lead sponsor's organization record. | [optional]
**sponsor_name** | Option<**String**> | The sponsor's organization name. | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the sponsor's company (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent; a root's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**drug_id** | Option<**String**> | The drug key. | [optional]
**drug_name** | Option<**String**> | The drug's name. | [optional]
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The condition's name. | [optional]
**condition_ancestor_ids** | Option<**Vec<String>**> | Every MeSH descriptor id at or above the condition in the tree, the condition itself included. Filter a subtree with = ANY. | [optional]
**max_phase** | Option<**i32**> | Highest phase reached: 0 early phase 1 through 6 for phase 4 (marketed); -1 when unstated. | [optional]
**n_trials** | Option<**i64**> | Lead-sponsored trials in the program. | [optional]
**n_active** | Option<**i64**> | Of those, trials in an active status. | [optional]
**last_year** | Option<**i32**> | Latest trial start year. | [optional]
**last_update_year** | Option<**i32**> | Year the program's most recently updated trial record was posted. | [optional]
**withdrawn** | Option<**bool**> | True when the drug has been withdrawn. | [optional]
**cause_name** | Option<**String**> | The Global Burden of Disease cause matched to the condition. | [optional]
**dalys_millions** | Option<**f64**> | Global disability-adjusted life years for the matched cause, in millions. | [optional]
**deaths_millions** | Option<**f64**> | Global deaths for the matched cause, in millions. | [optional]
**druggability** | Option<**i32**> | Druggability class of the matched cause: 1 highly druggable, 2 druggable, 3 challenging, 4 very challenging. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



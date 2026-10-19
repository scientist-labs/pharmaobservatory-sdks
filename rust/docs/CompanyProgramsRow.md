# CompanyProgramsRow

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
**condition_id** | Option<**String**> | The condition key (UMLS CUI). | [optional]
**condition_name** | Option<**String**> | The condition's name. | [optional]
**max_phase** | Option<**i32**> | Highest phase reached: 0 early phase 1 through 6 for phase 4 (marketed); -1 when unstated. | [optional]
**n_trials** | Option<**i64**> | Lead-sponsored trials in the program. | [optional]
**n_active** | Option<**i64**> | Of those, trials in an active status. | [optional]
**last_year** | Option<**i32**> | Latest trial start year. | [optional]
**withdrawn** | Option<**bool**> | True when the drug has been withdrawn. | [optional]
**lifecycle** | Option<**String**> | approved, active or abandoned. | [optional]
**dalys_millions** | Option<**f64**> | Disability-adjusted life years (millions) of the condition's global burden, IHME GBD 2021. Null for a caller without the licence for this column; see withheld_columns. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



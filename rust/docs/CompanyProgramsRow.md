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
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The condition's name. | [optional]
**max_phase** | Option<**i32**> | Highest phase reached on the 0 to 4 scale: 0 not reported, 1 through 4 for phase 1 through phase 4 (4 is marketed); a split phase folds up to its higher phase. | [optional]
**max_phase_label** | Option<**String**> | The furthest phase reached with the split phases kept: Early 1, Phase 1, Phase 1/2, Phase 2, Phase 2/3, Phase 3, Phase 4, or Not reported. max_phase is the same value folded onto 0 to 4 (a split phase counts as its higher phase). | [optional]
**max_phase_fine_rank** | Option<**i32**> | The finest phase rank behind max_phase_label: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early 1, -1 not reported. Sort on it to order programs by furthest phase with the split phases kept; max_phase is the same value folded onto 0 to 4. | [optional]
**n_trials** | Option<**i64**> | Lead-sponsored trials in the program. | [optional]
**n_active** | Option<**i64**> | Of those, trials in an active status. | [optional]
**last_year** | Option<**i32**> | Latest trial start year. | [optional]
**withdrawn** | Option<**bool**> | True when the drug has been withdrawn. | [optional]
**lifecycle** | Option<**String**> | approved, active or abandoned. | [optional]
**dalys_millions** | Option<**f64**> | Disability-adjusted life years (DALYs, millions) of the condition's global burden, from the Global Burden of Disease (GBD) 2021 study of the Institute for Health Metrics and Evaluation (IHME). Null for a caller without the licence for this column; see withheld_columns. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



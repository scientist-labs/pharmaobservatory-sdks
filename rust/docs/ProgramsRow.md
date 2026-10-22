# ProgramsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**program_id** | Option<**String**> | The program key: a hash of drug_id and condition_id, one value per row. | [optional]
**drug_id** | Option<**String**> | The drug key of the program, the drug record the trials name (a salt form or a trial arm is its own record). | [optional]
**molecule_id** | Option<**String**> | The molecule key: the resolved drug the program's drug record folds to, and the entry key for the drugs object type. | [optional]
**drug_name** | Option<**String**> | The drug's name. | [optional]
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The condition's name. | [optional]
**is_indication** | Option<**bool**> | True for a specific disease, false for a broad MeSH umbrella condition. | [optional]
**max_phase** | Option<**i32**> | Highest phase reached across the program's trials: 0 not reported, 1 through 4 for phase 1 through phase 4 (the 0 to 4 scale shared with company_programs; see caveats). | [optional]
**n_trials** | Option<**i64**> | Distinct trials of the program, across all sponsors. | [optional]
**n_completed** | Option<**i64**> | Of those, trials with status completed. | [optional]
**n_terminated** | Option<**i64**> | Of those, trials with status terminated. | [optional]
**first_year** | Option<**i32**> | Earliest plausible trial start year. | [optional]
**last_year** | Option<**i32**> | Latest plausible trial completion or start year. | [optional]
**n_sponsor_companies** | Option<**i32**> | Distinct lead-sponsor companies across the program's trials; an unresolved lead sponsor counts as itself. | [optional]
**dalys_millions** | Option<**f64**> | Disability-adjusted life years (DALYs, millions) of the condition's global burden, from the Global Burden of Disease (GBD) 2021 study of the Institute for Health Metrics and Evaluation (IHME). Null for a caller without the licence for this column; see withheld_columns. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



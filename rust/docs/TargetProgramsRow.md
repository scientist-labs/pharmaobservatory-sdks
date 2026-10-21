# TargetProgramsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional]
**target_symbol** | Option<**String**> | The target's HGNC gene symbol (the id the target page address carries). | [optional]
**drug_id** | Option<**String**> | The drug's canonical id (the same id the drug datasets use as drug_id). | [optional]
**drug_name** | Option<**String**> | The drug's name. | [optional]
**condition_id** | Option<**String**> | The condition as a MeSH descriptor id (D followed by digits), the id the conditions datasets use. | [optional]
**condition_name** | Option<**String**> | The condition's name. | [optional]
**max_phase** | Option<**i32**> | The furthest phase the program reached, 0 to 4 (0 is phase not reported). | [optional]
**n_trials** | Option<**i32**> | Distinct trials of the drug in the indication. | [optional]
**n_completed** | Option<**i32**> | Of n_trials, those with status COMPLETED. | [optional]
**n_terminated_trials** | Option<**i32**> | Of n_trials, those with status TERMINATED. | [optional]
**first_year** | Option<**i32**> | The earliest plausible trial start year. | [optional]
**last_year** | Option<**i32**> | The latest plausible trial completion or start year. | [optional]
**is_discontinued** | Option<**bool**> | True when the program has a terminated, withdrawn or suspended trial. | [optional]
**n_discontinued** | Option<**i32**> | Distinct stopped trials (terminated, withdrawn or suspended); 0 for a live program. | [optional]
**n_terminated** | Option<**i32**> | Of the stopped trials, those TERMINATED. | [optional]
**n_withdrawn** | Option<**i32**> | Of the stopped trials, those WITHDRAWN. | [optional]
**n_suspended** | Option<**i32**> | Of the stopped trials, those SUSPENDED. | [optional]
**max_phase_at_stop** | Option<**i32**> | The furthest phase among the stopped trials, 0 to 4; null for a live program. | [optional]
**last_stop_year** | Option<**i32**> | The latest year a stopped trial ended; null for a live program. | [optional]
**why_stopped_reasons** | Option<**Vec<String>**> | The stop reasons registrants disclosed, verbatim; empty when none. | [optional]
**why_stopped_categories** | Option<**Vec<String>**> | The normalized stop-reason categories (efficacy_futility, safety, funding and so on); empty when none. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TrialResearchPaymentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**payment_key** | Option<**String**> | A stable key for the trial, payer and program year row. | [optional]
**payer_name** | Option<**String**> | The paying manufacturer's name as reported. | [optional]
**payer_org_id** | Option<**String**> | Legal-entity level: the payer rolled up to its ultimate parent organization record (null when unresolved). Payments from subsidiaries of one parent share one row per trial and program year. | [optional]
**payer_company_id** | Option<**String**> | Company level: the company the payer belongs to (null when unresolved). | [optional]
**payer_company_name** | Option<**String**> | The payer's company name. | [optional]
**payer_group_id** | Option<**String**> | Group level: the current ultimate parent of the payer's company. | [optional]
**payer_group_name** | Option<**String**> | The ultimate parent company's name. | [optional]
**program_year** | Option<**i32**> | The Open Payments program year of the payments. | [optional]
**research_payments** | Option<**i32**> | Payment records behind the row. | [optional]
**total_amount_usd** | Option<**f64**> | Total reported dollars for the row, counted once. | [optional]
**pi_count** | Option<**i32**> | Distinct resolved principal investigators named on the payment records. | [optional]
**any_disputed** | Option<**bool**> | True when any payment record in the row was disputed by the recipient. | [optional]
**any_preclinical** | Option<**bool**> | True when any payment record in the row was marked as preclinical research. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CompanyTargetsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sponsor_company_id** | Option<**String**> | Company level: the company whose lead-sponsored trials pursue the target. | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**target_id** | Option<**String**> | The target's gene identifier (Ensembl). | [optional]
**target_symbol** | Option<**String**> | The target's gene symbol. | [optional]
**target_name** | Option<**String**> | The target's name. | [optional]
**max_phase** | Option<**i32**> | The furthest phase any qualifying program of the company reached against the target. | [optional]
**n_drugs** | Option<**i32**> | Distinct drugs of the company's qualifying programs against the target. | [optional]
**n_conditions** | Option<**i32**> | Distinct conditions of the company's qualifying programs against the target. | [optional]
**n_trials** | Option<**i32**> | Distinct trials the company leads in those programs. | [optional]
**n_programs** | Option<**i32**> | Distinct drug and indication programs (a specific indication with at least two trials) the company leads against the target. | [optional]
**last_year** | Option<**i32**> | The latest start year among the company's trials against the target in those programs. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



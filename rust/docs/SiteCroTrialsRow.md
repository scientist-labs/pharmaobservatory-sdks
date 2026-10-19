# SiteCroTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**site_id** | Option<**String**> | The site key: one facility spelling from the registry. | [optional]
**trial_id** | Option<**String**> | The trial's registry id (NCT number). | [optional]
**cro_org_id** | Option<**String**> | Legal-entity level: the CRO's organization record. | [optional]
**cro_name** | Option<**String**> | The CRO organization's name. | [optional]
**cro_company_id** | Option<**String**> | Company level: the company the CRO belongs to (null when unresolved). | [optional]
**cro_company_name** | Option<**String**> | Display name of the CRO's company. | [optional]
**cro_group_id** | Option<**String**> | Group level: the CRO's current ultimate parent company; a root company's group is itself. | [optional]
**cro_group_name** | Option<**String**> | Display name of the CRO's group. | [optional]
**cro_attribution_method** | Option<**String**> | How the CRO was inferred (sponsor_match, collaborator_match or llm_extraction). | [optional]
**cro_confidence** | Option<**f64**> | Confidence of the inference, from 0 to 1. | [optional]
**overall_status** | Option<**String**> | The registry's overall status of the trial. | [optional]
**start_year** | Option<**i32**> | Year of the trial start. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



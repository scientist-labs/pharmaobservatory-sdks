# SbtiCrosscheckRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**org_id** | Option<**String**> | Legal-entity level: the organization the cross-check is for. | [optional]
**org_name** | Option<**String**> | Display name of the organization record. | [optional]
**company_id** | Option<**String**> | Company level: the company the organization belongs to (null when unresolved). | [optional]
**company_name** | Option<**String**> | Display name of the organization's company. | [optional]
**group_id** | Option<**String**> | Group level: the organization's current ultimate parent company; a root company's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the organization's group. | [optional]
**reported_company_name** | Option<**String**> | The company name as the SBTi list or the disclosure gives it. | [optional]
**recon_status** | Option<**String**> | The cross-check result: agree, extracted_only, sbti_only, target_year_conflict or self_claim_conflict. | [optional]
**has_net_zero_target** | Option<**bool**> | Whether a net-zero target is stated (SBTi's value where it lists the organization, else the disclosure's). | [optional]
**net_zero_year** | Option<**i32**> | The net-zero target year. | [optional]
**has_near_term_target** | Option<**bool**> | Whether a near-term target is stated. | [optional]
**near_term_target_year** | Option<**i32**> | The near-term target year. | [optional]
**sbti_validated** | Option<**bool**> | Whether SBTi lists the targets as validated; null when SBTi does not list the organization. | [optional]
**claimed_sbti_status** | Option<**String**> | The SBTi status the organization's own filing claims. | [optional]
**in_sbti** | Option<**bool**> | Whether the SBTi list carries the organization. | [optional]
**in_extracted** | Option<**bool**> | Whether an extraction of the organization's own filing exists. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



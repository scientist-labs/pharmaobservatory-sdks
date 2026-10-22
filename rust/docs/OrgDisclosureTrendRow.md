# OrgDisclosureTrendRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**org_id** | Option<**String**> | Legal-entity level: the organization record of the SEC registrant that filed. | [optional]
**org_name** | Option<**String**> | Display name of the filing organization record. | [optional]
**company_id** | Option<**String**> | Company level: the company the filing organization belongs to (null when unresolved). | [optional]
**company_name** | Option<**String**> | Display name of the filing organization's company. | [optional]
**group_id** | Option<**String**> | Group level: the filing organization's current ultimate parent company; a root company's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the filing organization's group. | [optional]
**reporting_year** | Option<**i32**> | The fiscal year the filing reports. | [optional]
**reporting_company** | Option<**String**> | The company name as the filing states it. | [optional]
**cik** | Option<**String**> | The registrant's SEC Central Index Key (CIK). | [optional]
**ticker** | Option<**String**> | The stock ticker the filing states. | [optional]
**board_pct_women** | Option<**f64**> | Women's share of the board, as the filing states it. | [optional]
**workforce_pct_women** | Option<**f64**> | Women's share of the workforce. | [optional]
**leadership_pct_women** | Option<**f64**> | Women's share of leadership roles. | [optional]
**workforce_pct_underrepresented** | Option<**f64**> | Share of the workforce from underrepresented groups. | [optional]
**gender_pay_gap_pct** | Option<**f64**> | The stated gender pay gap. | [optional]
**has_dei_targets** | Option<**bool**> | Whether the filing states diversity, equity and inclusion targets. | [optional]
**has_net_zero_target** | Option<**bool**> | Whether the filing states a net-zero target. | [optional]
**net_zero_target_year** | Option<**i32**> | The year of the net-zero target. | [optional]
**reduction_pct** | Option<**f64**> | The emissions reduction the filing commits to. | [optional]
**sbti_status_claimed** | Option<**String**> | The Science Based Targets initiative (SBTi) status the filing claims for itself (a claim in the filing, not the SBTi list). | [optional]
**target_temperature_claimed** | Option<**String**> | The warming limit the filing says its targets align with. | [optional]
**climate_document_id** | Option<**String**> | The filing document the climate measures were extracted from. | [optional]
**diversity_document_id** | Option<**String**> | The filing document the diversity measures were extracted from. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



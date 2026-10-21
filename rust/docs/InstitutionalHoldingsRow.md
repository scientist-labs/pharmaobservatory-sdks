# InstitutionalHoldingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**holding_id** | Option<**String**> | The holding line id. | [optional]
**issuer_name** | Option<**String**> | The security issuer as filed. | [optional]
**issuer_org_id** | Option<**String**> | Legal-entity level: the issuer's organization record (null when unresolved). | [optional]
**issuer_company_id** | Option<**String**> | Company level: the issuer's company (null when unresolved). | [optional]
**issuer_company_name** | Option<**String**> | Display name of the issuer's company. | [optional]
**issuer_group_id** | Option<**String**> | Group level: the issuer's current ultimate parent; a root's group is itself. | [optional]
**issuer_group_name** | Option<**String**> | Display name of the issuer's group. | [optional]
**security_ref** | Option<**String**> | The security's CUSIP (Committee on Uniform Securities Identification Procedures) identifier. | [optional]
**holder_name** | Option<**String**> | The institutional manager's name. | [optional]
**holder_identity_key** | Option<**String**> | The institutional manager's stable identity (an SEC registration number), the key that groups all of one manager's holdings. | [optional]
**holder_investor_id** | Option<**String**> | The manager's investor key where the link resolves (null when the manager resolves to no investor or its rows disagree). A manager-level value: the same on every row of one manager. | [optional]
**holder_link_conflict** | Option<**bool**> | True when the manager's own holding rows resolve to different investors, so holder_investor_id is null. | [optional]
**holder_kind** | Option<**String**> | Manager kind. | [optional]
**holder_org_id** | Option<**String**> | Legal-entity level: the manager's organization record (null when unresolved). | [optional]
**holder_company_id** | Option<**String**> | Company level: the manager's company (null when unresolved). | [optional]
**holder_company_name** | Option<**String**> | Display name of the manager's company. | [optional]
**holder_group_id** | Option<**String**> | Group level: the manager's current ultimate parent. | [optional]
**holder_group_name** | Option<**String**> | Display name of the manager's group. | [optional]
**reported_value_amount** | Option<**f64**> | Reported value of the position in whole US dollars at the filing's inferred value scale; equity (SH) lines only. Null where the scale or currency is ambiguous. | [optional]
**reported_value_currency** | Option<**String**> | Currency of the reported value. | [optional]
**shares_basis** | Option<**String**> | How the share figure was reached: sum_latest_period_positions, the sum of the manager's positions in the issuer's latest reporting period. Read a share figure with its basis. | [optional]
**shares** | Option<**f64**> | Shares standing. | [optional]
**percent_basis** | Option<**String**> | The basis the percent is measured on. | [optional]
**reported_percent** | Option<**f64**> | Reported share of the issuer. | [optional]
**edge_state** | Option<**String**> | State of the position (standing, ceased, no longer reported). | [optional]
**first_observed_date** | Option<[**String**](string.md)> | First filing date the position was observed. | [optional]
**last_filing_date** | Option<[**String**](string.md)> | Latest filing date the position was reported. | [optional]
**contains_personal_data** | Option<**bool**> | Whether the holder is a natural person. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



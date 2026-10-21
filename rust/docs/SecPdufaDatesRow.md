# SecPdufaDatesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**pdufa_observation_id** | Option<**String**> | The extracted disclosure id (filing document id and disclosure index). | [optional]
**event_category** | Option<**String**> | pdufa, advisory_committee, priority_review or other. | [optional]
**status** | Option<**String**> | The disclosure's status: pending, approved, crl (complete response letter), or withdrawn. | [optional]
**is_decided** | Option<**bool**> | True when the status is approved, crl or withdrawn. | [optional]
**is_forward** | Option<**bool**> | True when the disclosure is a pending, forward-looking date. | [optional]
**application_type** | Option<**String**> | The application type the filing names (NDA, BLA and so on). | [optional]
**designation** | Option<**String**> | The review designation the filing names (priority review, fast track and so on). | [optional]
**indication** | Option<**String**> | The indication the filing names. | [optional]
**date_phrase** | Option<**String**> | The date as the filing words it. | [optional]
**date_start** | Option<[**String**](string.md)> | Start of the date interval the phrase means. | [optional]
**date_end** | Option<[**String**](string.md)> | End of the date interval the phrase means. | [optional]
**date_precision** | Option<**String**> | How exact the phrase is (day, month, quarter and so on). | [optional]
**is_in_period_now** | Option<**bool**> | True when today falls in the interval and the disclosure is not decided. | [optional]
**drug_name** | Option<**String**> | The drug name as filed. | [optional]
**sponsor_name** | Option<**String**> | The filer's name. | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the SEC registrant's organization record (null when it is not in the spine). | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the filer's company (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the filer's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the filer's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the filer's group. | [optional]
**ticker** | Option<**String**> | The stock ticker the filing names. | [optional]
**source_url** | Option<**String**> | The filing's address on the SEC site. | [optional]
**confidence** | Option<**f64**> | The extraction's own confidence, from 0 to 1. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



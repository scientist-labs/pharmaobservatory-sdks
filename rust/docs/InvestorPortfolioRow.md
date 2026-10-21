# InvestorPortfolioRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investor_id** | Option<**String**> | The investor key. | [optional]
**issuer_cik** | Option<**String**> | The issuer's SEC Central Index Key (CIK), ten digits. | [optional]
**issuer_name** | Option<**String**> | The issuer's name: its organization record's name, else the SEC company name. | [optional]
**issuer_org_id** | Option<**String**> | Legal-entity level: the issuer's organization record (null when unresolved). | [optional]
**issuer_company_id** | Option<**String**> | Company level: the issuer's company (null when unresolved). | [optional]
**issuer_company_name** | Option<**String**> | Display name of the issuer's company. | [optional]
**issuer_group_id** | Option<**String**> | Group level: the issuer's current ultimate parent; a root's group is itself. | [optional]
**issuer_group_name** | Option<**String**> | Display name of the issuer's group. | [optional]
**relationship_types** | Option<**Vec<String>**> | How the investor is named against the issuer: holder_5pct (a 5 percent holder) or selling_stockholder. | [optional]
**edge_count** | Option<**i64**> | Filing mentions behind the row, across every filing and relationship type. | [optional]
**first_observed_date** | Option<[**String**](string.md)> | The earliest filing date the investor was named against the issuer. | [optional]
**last_observed_date** | Option<[**String**](string.md)> | The latest filing date the investor was named against the issuer. | [optional]
**is_life_sciences_issuer** | Option<**bool**> | True when the issuer's SIC (Standard Industrial Classification) code is a life-sciences code, false when it is another code, null when the issuer has no SIC evidence. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



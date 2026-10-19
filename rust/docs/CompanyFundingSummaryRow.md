# CompanyFundingSummaryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | The company key (company:<key>). | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the current ultimate parent company; a root company's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**n_rounds** | Option<**i64**> | Funding events of the company, disclosed or not. | [optional]
**total_non_dilutive_usd** | Option<**f64**> | Disclosed non-dilutive amounts (grants, debt, royalty and similar), in whole US dollars. | [optional]
**total_dilutive_usd** | Option<**f64**> | Disclosed dilutive amounts (equity-issuing rounds), in whole US dollars. | [optional]
**total_unclassified_usd** | Option<**f64**> | Disclosed amounts of events whose instrument is not classified, kept apart from the other two totals. | [optional]
**n_unclassified_rounds** | Option<**i64**> | Events whose instrument is not classified. | [optional]
**total_raised_usd** | Option<**f64**> | All disclosed amounts, the sum of the three totals, in whole US dollars. | [optional]
**first_funded_date** | Option<[**String**](string.md)> | Date of the earliest event. | [optional]
**last_funded_date** | Option<[**String**](string.md)> | Date of the latest event. | [optional]
**latest_round_type** | Option<**String**> | The round type of the latest event. | [optional]
**latest_round_date** | Option<[**String**](string.md)> | The date of the latest event. | [optional]
**stage_label** | Option<**String**> | A coarse reading of the round types: public, growth_private, seed_private, unclassified or non_dilutive_only. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



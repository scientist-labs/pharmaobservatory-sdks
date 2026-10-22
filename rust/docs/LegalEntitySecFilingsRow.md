# LegalEntitySecFilingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**org_id** | Option<**String**> | Legal-entity level: the SEC filer organization record. | [optional]
**org_name** | Option<**String**> | Display name of the filer organization record. | [optional]
**company_id** | Option<**String**> | Company level: the company the filer belongs to (null when unresolved). | [optional]
**company_name** | Option<**String**> | Display name of the filer's company. | [optional]
**group_id** | Option<**String**> | Group level: the filer's current ultimate parent company; a root company's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the filer's group. | [optional]
**edgar_cik** | Option<**String**> | The filer's 10-digit SEC Central Index Key (CIK). | [optional]
**jurisdiction_of_inc** | Option<**String**> | The jurisdiction of incorporation as the latest filing states it (a US state or a country, free text); null when blank. | [optional]
**entity_type** | Option<**String**> | The entity type the latest filing states (Corporation, Limited Liability Company, Limited Partnership, General Partnership, Business Trust or Other); null when blank. | [optional]
**entity_type_other_desc** | Option<**String**> | The filer's free-text description when it ticked entity type Other. | [optional]
**year_of_inc_band** | Option<**String**> | The filing's year-of-incorporation choice (Over Five Years Ago, Within Last Five Years or Yet to Be Formed); null when blank. | [optional]
**year_of_inc** | Option<**i32**> | The year of incorporation when the filing gives a four-digit year. | [optional]
**industry_group_types** | Option<**Vec<String>**> | The industry groups the filer's in-scope filings name. | [optional]
**previous_names** | Option<**Vec<String>**> | Earlier names the filer states, other than its current name. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



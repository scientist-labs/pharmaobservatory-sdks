# CompanyUsPatentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | Company level: the company the row is about; the entry key for the companies object type. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**patent_id** | Option<**String**> | The PatentsView patent id. | [optional]
**patent_number** | Option<**String**> | The US patent number. | [optional]
**patent_title** | Option<**String**> | The patent title. | [optional]
**patent_date** | Option<[**String**](string.md)> | The grant date. | [optional]
**is_biomedical** | Option<**bool**> | True when the patent carries a biomedical CPC class; the split is never a filter on what belongs to the portfolio. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



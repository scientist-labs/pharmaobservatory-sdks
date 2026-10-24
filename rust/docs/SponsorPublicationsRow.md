# SponsorPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | Company level: the company the row is about; the entry key for the companies object type. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**pmid** | Option<**String**> | The PubMed id. | [optional]
**citation** | Option<**String**> | The citation text of the paper as the trial registry cites it. | [optional]
**title** | Option<**String**> | The article title. | [optional]
**pub_year** | Option<**i32**> | The publication year. | [optional]
**journal_nlm_id** | Option<**String**> | The journal's NLM (National Library of Medicine) catalog id; null when the paper's journal is unknown. | [optional]
**journal_title** | Option<**String**> | The journal title; null when the journal is not in the journal directory. | [optional]
**publication_trials** | Option<**i32**> | How many trials of any sponsor cite the paper. | [optional]
**n_trials** | Option<**i32**> | Distinct trials of this company that cite the paper. | [optional]
**n_lead_trials** | Option<**i32**> | Of those, the trials the company leads. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



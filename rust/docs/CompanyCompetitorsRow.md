# CompanyCompetitorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | Company level: the company the row is about; the entry key for the companies object type. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**competitor_company_id** | Option<**String**> | Company level: the competing company. | [optional]
**competitor_company_name** | Option<**String**> | Display name of the competitor's company. | [optional]
**competitor_group_id** | Option<**String**> | Group level: the competitor's current ultimate parent; a root's group is itself. | [optional]
**competitor_group_name** | Option<**String**> | Display name of the competitor's group. | [optional]
**competitor_rank** | Option<**i32**> | The competitor's rank for the company by competitor_score (1 is the highest). | [optional]
**competitor_score** | Option<**f64**> | 100 times the weighted mean of the four similarities below, rounded to four places. | [optional]
**target_sim** | Option<**f64**> | Jaccard similarity, 0 to 1, of the two companies' molecular target sets. | [optional]
**mechanism_sim** | Option<**f64**> | Jaccard similarity, 0 to 1, of their mechanism sets (a target and its action). | [optional]
**indication_sim** | Option<**f64**> | Jaccard similarity, 0 to 1, of their top-25 disease area sets. | [optional]
**phase_sim** | Option<**f64**> | The fraction, 0 to 1, of the shared targets that both companies have taken to phase 2 or later. | [optional]
**n_shared_targets** | Option<**i32**> | Molecular targets both companies pursue. | [optional]
**n_shared_mechanisms** | Option<**i32**> | Mechanisms (target and action) both companies pursue. | [optional]
**n_shared_indications** | Option<**i32**> | Disease areas both companies are in, within each one's top 25. | [optional]
**n_phase_overlap_targets** | Option<**i32**> | Shared targets that both companies have taken to phase 2 or later. | [optional]
**n_company_targets** | Option<**i32**> | Targets the company pursues (the Jaccard denominator). | [optional]
**n_competitor_targets** | Option<**i32**> | Targets the competitor pursues. | [optional]
**n_company_indications** | Option<**i32**> | Disease areas in the company's top 25. | [optional]
**n_competitor_indications** | Option<**i32**> | Disease areas in the competitor's top 25. | [optional]
**top_shared_targets** | Option<**Vec<String>**> | Up to 5 shared target gene symbols, by combined phase then combined trials. An empty list when none. | [optional]
**top_shared_indications** | Option<**Vec<String>**> | Up to 5 shared disease area names, by combined trial count. An empty list when none. | [optional]
**top_shared_mechanisms** | Option<**Vec<String>**> | Up to 3 shared mechanisms written as a gene symbol and an action, by combined phase then combined trials. An empty list when none. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



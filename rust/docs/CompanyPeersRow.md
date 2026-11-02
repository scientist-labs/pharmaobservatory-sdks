# CompanyPeersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | Company level: the company the row is about; the entry key for the companies object type. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**peer_company_id** | Option<**String**> | Company level: the peer company. | [optional]
**peer_company_name** | Option<**String**> | Display name of the peer company. | [optional]
**peer_group_id** | Option<**String**> | Group level: the peer's current ultimate parent; a root's group is itself. | [optional]
**peer_group_name** | Option<**String**> | Display name of the peer's group. | [optional]
**peer_rank** | Option<**i32**> | The peer's rank for the company, 1 is the closest match: match_level first, then the closeness of lead-trial counts, then of biologic share. Dense 1 to N. | [optional]
**match_level** | Option<**i32**> | What the peer matched the company on. 1 stage, modality, area and size; 2 stage, modality and area; 3 stage, modality and size; 4 stage and modality. Stage and modality always match. | [optional]
**matched_area** | Option<**bool**> | True when the peer shares the company's primary therapeutic area. | [optional]
**matched_size** | Option<**bool**> | True when the peer shares the company's size band. | [optional]
**lead_stage** | Option<**String**> | The peer's lead stage, the furthest phase of its live owned programs. Equal to the company's lead stage on every row. | [optional]
**lead_drug_name** | Option<**String**> | The peer's lead drug, the highest-phase live owned drug (ties: most trials). | [optional]
**biologic_share** | Option<**f64**> | Share, 0 to 1, of the peer's live owned drugs that are biologic (GSRS protein or nucleic acid substance type). | [optional]
**modality_band** | Option<**String**> | biologic when biologic_share is at least 0.5, else non_biologic. Equal to the company's band on every row. | [optional]
**primary_area** | Option<**String**> | The peer's primary therapeutic area: the MeSH top-level category with the most live owned program trials. | [optional]
**num_lead_trials** | Option<**i32**> | Distinct lead-sponsored trials of the peer's subtree. | [optional]
**size_band** | Option<**String**> | Band of num_lead_trials: 1-4, 5-19, 20-99 or 100+. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



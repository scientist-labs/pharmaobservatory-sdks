# CompanyMarketCollisionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | Company level: the company the row is about; the entry key for the companies object type. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**molecule_id** | Option<**String**> | Molecule level: the company's asset, the resolved drug (a drug with two canonical ids is one asset). | [optional]
**molecule_name** | Option<**String**> | The asset's display name, from the company's most-trialled record of it. | [optional]
**competitor_company_id** | Option<**String**> | Company level: the company whose asset collides. | [optional]
**competitor_company_name** | Option<**String**> | Display name of the competitor's company. | [optional]
**competitor_group_id** | Option<**String**> | Group level: the competitor's current ultimate parent; a root's group is itself. | [optional]
**competitor_group_name** | Option<**String**> | Display name of the competitor's group. | [optional]
**competitor_molecule_id** | Option<**String**> | Molecule level: the competitor's asset, the resolved drug. | [optional]
**competitor_molecule_name** | Option<**String**> | The competitor asset's display name. | [optional]
**condition_ui** | Option<**String**> | The shared MeSH indication (a MeSH descriptor id): both assets are owned, live and Phase 3 or 4 in it. | [optional]
**condition_name** | Option<**String**> | The MeSH name of the shared indication. | [optional]
**target_symbols** | Option<**Vec<String>**> | Gene symbols of the company drug's recorded targets. Never empty, and none is a target of the competitor drug. | [optional]
**competitor_target_symbols** | Option<**Vec<String>**> | Gene symbols of the competitor drug's recorded targets. | [optional]
**stage** | Option<**i32**> | The less advanced of the two programs' furthest phase: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1, -1 phase not applicable. | [optional]
**last_year** | Option<**i32**> | The later of the two programs' last trial start year. | [optional]
**n_head_to_head_trials** | Option<**i32**> | Trials led by either company in which one drug is experimental and the other an active comparator. | [optional]
**ownership_basis** | Option<**String**> | Why the company owns the drug. experimental_arm and regulatory_holder rest on evidence; inferred_lead is an inference (no arm data says otherwise). | [optional]
**competitor_ownership_basis** | Option<**String**> | Why the competitor owns its drug, same vocabulary as ownership_basis. | [optional]
**indication_idf** | Option<**f64**> | The rarity weight of the indication: the natural log of (companies holding a live, owned Phase 3+ asset with a recorded target) over (companies holding one in this indication). | [optional]
**market_weight** | Option<**f64**> | The weight of the pair, 0.5 times indication_idf times factors for recency and a head-to-head trial. There is no stage factor, both programs are Phase 3 or 4. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



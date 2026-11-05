# CompanyAssetCollisionsRow

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
**target_family** | Option<**String**> | The company drug's target family of the shared gene (the Open Targets target name, so a taxane's tubulin isotypes are one family; the ChEMBL mechanism text for a drug Open Targets lacks; the gene symbol for a curated investigational asset). | [optional]
**shared_symbols** | Option<**Vec<String>**> | Gene symbols of the shared genes within the family. | [optional]
**action_type** | Option<**String**> | The company drug's action on the family (Open Targets or ChEMBL vocabulary); null when no source states one. | [optional]
**competitor_action_type** | Option<**String**> | The competitor drug's action on the shared genes; null when no source states one. | [optional]
**same_action** | Option<**bool**> | True when the two drugs have the same action type on a shared gene of the family. | [optional]
**n_shared_indications** | Option<**i32**> | Disease areas both programs are in, among live programs. | [optional]
**top_shared_indications** | Option<**Vec<String>**> | Up to 3 shared disease area names, by the pair's combined trial count. An empty list when none. | [optional]
**stage** | Option<**i32**> | The less advanced of the two programs' furthest phase: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1, -1 phase not applicable. | [optional]
**last_year** | Option<**i32**> | The later of the two programs' last trial start year. | [optional]
**n_head_to_head_trials** | Option<**i32**> | Trials led by either company in which one drug is experimental and the other an active comparator. | [optional]
**ownership_basis** | Option<**String**> | Why the company owns the drug. experimental_arm and regulatory_holder rest on evidence; inferred_lead is an inference (no arm data says otherwise). | [optional]
**competitor_ownership_basis** | Option<**String**> | Why the competitor owns its drug, same vocabulary as ownership_basis. | [optional]
**family_idf** | Option<**f64**> | The rarity weight of the family: the natural log of (companies holding any live, owned target family) over (companies holding this one). | [optional]
**collision_weight** | Option<**f64**> | The weight of the pair, family_idf times a factor for same action, a shared indication, the stage, recency and a head-to-head trial. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



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
**competitor_class** | Option<**String**> | The competitor's class: industry, or non_industry (not classed as industry by the sponsor-class rule: academic, hospital, government, nonprofit, and any record the rule does not reach). competitor_rank restarts at 1 within each class, so filter or sort on this column first. | [optional]
**n_class_candidates** | Option<**i32**> | How many companies of this competitor_class had at least one asset collision with the company before the top 25 were kept. The same value on every row of one company and class. | [optional]
**competitor_rank** | Option<**i32**> | The competitor's rank within its competitor_class for the company by competitor_score (1 is the highest). | [optional]
**competitor_score** | Option<**f64**> | The weighted share of the company's live target families that the competitor's live assets collide with, 0 to 100, rounded to four places. Each colliding family counts its best collision weight (rarity of the family across companies, times factors for same action, a shared indication, the least advanced stage, recency and a head-to-head trial), capped at the family's rarity weight; the sum is divided by the rarity weights of ALL the company's families. | [optional]
**n_collisions** | Option<**i32**> | Asset pairs (a drug of the company, a drug of the competitor) that act on a common target. | [optional]
**n_colliding_families** | Option<**i32**> | Target families of the company's live assets that the competitor collides with. | [optional]
**n_subject_families** | Option<**i32**> | Target families of the company's own live, owned assets; the denominator of competitor_score is their rarity weights. | [optional]
**n_same_action** | Option<**i32**> | Asset pairs whose drugs act on the shared target with the same action type. | [optional]
**n_shared_indication_pairs** | Option<**i32**> | Asset pairs whose programs share at least one disease area. | [optional]
**n_head_to_head_trials** | Option<**i32**> | Distinct trials led by either company in which one drug of a pair is experimental and the other an active comparator. | [optional]
**top_collisions** | Option<**Vec<String>**> | Up to 3 strongest asset pairs as text: \"drug A / drug B: target family, stage, N shared indications\". The stage is the less advanced of the two programs. An ownership_basis of inferred_lead on the underlying pair (see company_asset_collisions) is an inference. | [optional]
**n_cosponsored_trials** | Option<**i32**> | Distinct trials on which the company and the competitor are both named as sponsors, counted over their subtrees; 0 when none. A consortium trial (more than 25 sponsor companies) is left out. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PharmaObservatory::CompanyCompetitorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **row_level** | **String** | root: the competitor&#39;s ultimate parent company, with its score, counts and evidence recomputed over the union of its members&#39; colliding assets (the best weight per target family and per indication across members, distinct asset pairs and distinct trials, never a sum of member scores). member: a company under that root that itself collides, scored on its own assets. Filter on one level; the two do not add. | [optional] |
| **competitor_company_id** | **String** | Company level: the competing company. At row_level root this is the root company; at member it is the member company itself. | [optional] |
| **competitor_company_name** | **String** | Display name of the competitor&#39;s company. | [optional] |
| **competitor_group_id** | **String** | Group level: the competitor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **competitor_group_name** | **String** | Display name of the competitor&#39;s group. | [optional] |
| **competitor_root_id** | **String** | The ultimate parent company of the competitor, the ladder company&#39;s &#x60;coalesce(ultimate_parent_company_id, company_id)&#x60; (one root per company, never the multi-root group fan-out). Equal to competitor_company_id at row_level root, and the root a member row sits under. | [optional] |
| **competitor_class** | **String** | The competitor&#39;s class: always industry since the not-classed-as-industry group left Competitors; the column is kept so a query written against it keeps working. competitor_rank is 1-based within the class. | [optional] |
| **n_class_candidates** | **Integer** | How many companies of this competitor_class had at least one asset collision with the company before the top 25 were kept. The same value on every row of one company and class. | [optional] |
| **competitor_rank** | **Integer** | The rank of the competitor ROOT within its competitor_class for the company by competitor_score (1 is the highest). The same on a root row and on each of its member rows. | [optional] |
| **member_rank** | **Integer** | 0 on a root row; 1 and up on the member rows of that root, by the member&#39;s own competitor_score. Order by competitor_rank then member_rank to list each root followed by its members. | [optional] |
| **n_members** | **Integer** | Member companies of this root that collide with the company, before the member cut. The same on a root row and its member rows. | [optional] |
| **competitor_score** | **Float** | The weighted share of the company&#39;s live target families that the competitor&#39;s live assets collide with, 0 to 100, rounded to four places. Each colliding family counts its best collision weight (rarity of the family across companies, times factors for same action, a shared indication, the least advanced stage, recency and a head-to-head trial), capped at the family&#39;s rarity weight; the sum is divided by the rarity weights of ALL the company&#39;s families. | [optional] |
| **n_collisions** | **Integer** | Asset pairs (a drug of the company, a drug of the competitor) that act on a common target; 0 when the competitor is listed only for a same-market pair (see n_market_collisions). | [optional] |
| **n_colliding_families** | **Integer** | Target families of the company&#39;s live assets that the competitor collides with. | [optional] |
| **n_subject_families** | **Integer** | Target families of the company&#39;s own live, owned assets; the denominator of competitor_score is their rarity weights. | [optional] |
| **n_same_action** | **Integer** | Asset pairs whose drugs act on the shared target with the same action type. | [optional] |
| **n_shared_indication_pairs** | **Integer** | Asset pairs whose programs share at least one disease area. Every listed pair does, so it equals n_collisions. | [optional] |
| **n_head_to_head_trials** | **Integer** | Distinct trials led by either company in which one drug of a pair is experimental and the other an active comparator. | [optional] |
| **best_stage** | **Integer** | The most advanced pair&#39;s stage, where a pair&#39;s stage is the less advanced of its two programs&#39; furthest phase: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1, -1 phase not applicable. A tie-break key after the score. | [optional] |
| **last_year** | **Integer** | The latest trial start year over the company pair&#39;s programs (the maximum over its asset pairs of the later of the two programs&#39; last start year). A tie-break key after the score; null when no program has a start year. | [optional] |
| **top_collisions** | **Array&lt;String&gt;** | Up to 3 strongest asset pairs as text: \&quot;drug A / drug B: target family, stage, N shared indications\&quot;. The stage is the less advanced of the two programs. An ownership_basis of inferred_lead on the underlying pair (see company_asset_collisions) is an inference. | [optional] |
| **n_market_collisions** | **Integer** | Asset pairs (a Phase 3+ drug of the company, a Phase 3+ drug of the competitor) that share a MeSH indication on drugs with no common target; the pairs are the company_market_collisions dataset. | [optional] |
| **n_market_indications** | **Integer** | Distinct shared Phase 3+ indications of the pair, counted after a shared indication is replaced by a more specific one that the same asset pair also shares. | [optional] |
| **n_subject_markets** | **Integer** | The company&#39;s own Phase 3+ indications; with its target families, the denominator of competitor_score. 0 when it holds no Phase 3+ program. | [optional] |
| **top_market_collisions** | **Array&lt;String&gt;** | Up to 3 strongest same-market asset pairs as text: \&quot;drug A / drug B: indication, same market, different target (SYMBOLS vs SYMBOLS), stage\&quot;. An empty list when none. Kept apart from top_collisions, which is target evidence only. | [optional] |
| **n_cosponsored_trials** | **Integer** | Distinct trials on which the company and the competitor are both named as sponsors, counted over their subtrees; 0 when none. A consortium trial (more than 25 sponsor companies) is left out. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyCompetitorsRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  row_level: null,
  competitor_company_id: null,
  competitor_company_name: null,
  competitor_group_id: null,
  competitor_group_name: null,
  competitor_root_id: null,
  competitor_class: null,
  n_class_candidates: null,
  competitor_rank: null,
  member_rank: null,
  n_members: null,
  competitor_score: null,
  n_collisions: null,
  n_colliding_families: null,
  n_subject_families: null,
  n_same_action: null,
  n_shared_indication_pairs: null,
  n_head_to_head_trials: null,
  best_stage: null,
  last_year: null,
  top_collisions: null,
  n_market_collisions: null,
  n_market_indications: null,
  n_subject_markets: null,
  top_market_collisions: null,
  n_cosponsored_trials: null
)
```


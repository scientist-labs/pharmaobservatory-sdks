# PharmaObservatory::CompanyCompetitorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **competitor_company_id** | **String** | Company level: the competing company. | [optional] |
| **competitor_company_name** | **String** | Display name of the competitor&#39;s company. | [optional] |
| **competitor_group_id** | **String** | Group level: the competitor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **competitor_group_name** | **String** | Display name of the competitor&#39;s group. | [optional] |
| **competitor_class** | **String** | The competitor&#39;s class: always industry since the not-classed-as-industry group left Competitors; the column is kept so a query written against it keeps working. competitor_rank is 1-based within the class. | [optional] |
| **n_class_candidates** | **Integer** | How many companies of this competitor_class had at least one asset collision with the company before the top 25 were kept. The same value on every row of one company and class. | [optional] |
| **competitor_rank** | **Integer** | The competitor&#39;s rank within its competitor_class for the company by competitor_score (1 is the highest). | [optional] |
| **competitor_score** | **Float** | The weighted share of the company&#39;s live target families that the competitor&#39;s live assets collide with, 0 to 100, rounded to four places. Each colliding family counts its best collision weight (rarity of the family across companies, times factors for same action, a shared indication, the least advanced stage, recency and a head-to-head trial), capped at the family&#39;s rarity weight; the sum is divided by the rarity weights of ALL the company&#39;s families. | [optional] |
| **n_collisions** | **Integer** | Asset pairs (a drug of the company, a drug of the competitor) that act on a common target. | [optional] |
| **n_colliding_families** | **Integer** | Target families of the company&#39;s live assets that the competitor collides with. | [optional] |
| **n_subject_families** | **Integer** | Target families of the company&#39;s own live, owned assets; the denominator of competitor_score is their rarity weights. | [optional] |
| **n_same_action** | **Integer** | Asset pairs whose drugs act on the shared target with the same action type. | [optional] |
| **n_shared_indication_pairs** | **Integer** | Asset pairs whose programs share at least one disease area. Every listed pair does, so it equals n_collisions. | [optional] |
| **n_head_to_head_trials** | **Integer** | Distinct trials led by either company in which one drug of a pair is experimental and the other an active comparator. | [optional] |
| **best_stage** | **Integer** | The most advanced pair&#39;s stage, where a pair&#39;s stage is the less advanced of its two programs&#39; furthest phase: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1, -1 phase not applicable. A tie-break key after the score. | [optional] |
| **last_year** | **Integer** | The latest trial start year over the company pair&#39;s programs (the maximum over its asset pairs of the later of the two programs&#39; last start year). A tie-break key after the score; null when no program has a start year. | [optional] |
| **top_collisions** | **Array&lt;String&gt;** | Up to 3 strongest asset pairs as text: \&quot;drug A / drug B: target family, stage, N shared indications\&quot;. The stage is the less advanced of the two programs. An ownership_basis of inferred_lead on the underlying pair (see company_asset_collisions) is an inference. | [optional] |
| **n_cosponsored_trials** | **Integer** | Distinct trials on which the company and the competitor are both named as sponsors, counted over their subtrees; 0 when none. A consortium trial (more than 25 sponsor companies) is left out. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyCompetitorsRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  competitor_company_id: null,
  competitor_company_name: null,
  competitor_group_id: null,
  competitor_group_name: null,
  competitor_class: null,
  n_class_candidates: null,
  competitor_rank: null,
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
  n_cosponsored_trials: null
)
```


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
| **competitor_rank** | **Integer** | The competitor&#39;s rank for the company by competitor_score (1 is the highest). | [optional] |
| **competitor_score** | **Float** | 100 times the weighted mean of the four similarities below, rounded to four places. | [optional] |
| **target_sim** | **Float** | Jaccard similarity, 0 to 1, of the two companies&#39; molecular target sets. | [optional] |
| **mechanism_sim** | **Float** | Jaccard similarity, 0 to 1, of their mechanism sets (a target and its action). | [optional] |
| **indication_sim** | **Float** | Jaccard similarity, 0 to 1, of their top-25 disease area sets. | [optional] |
| **phase_sim** | **Float** | The fraction, 0 to 1, of the shared targets that both companies have taken to phase 2 or later. | [optional] |
| **n_shared_targets** | **Integer** | Molecular targets both companies pursue. | [optional] |
| **n_shared_mechanisms** | **Integer** | Mechanisms (target and action) both companies pursue. | [optional] |
| **n_shared_indications** | **Integer** | Disease areas both companies are in, within each one&#39;s top 25. | [optional] |
| **n_phase_overlap_targets** | **Integer** | Shared targets that both companies have taken to phase 2 or later. | [optional] |
| **n_company_targets** | **Integer** | Targets the company pursues (the Jaccard denominator). | [optional] |
| **n_competitor_targets** | **Integer** | Targets the competitor pursues. | [optional] |
| **n_company_indications** | **Integer** | Disease areas in the company&#39;s top 25. | [optional] |
| **n_competitor_indications** | **Integer** | Disease areas in the competitor&#39;s top 25. | [optional] |
| **top_shared_targets** | **Array&lt;String&gt;** | Up to 5 shared target gene symbols, by combined phase then combined trials. An empty list when none. | [optional] |
| **top_shared_indications** | **Array&lt;String&gt;** | Up to 5 shared disease area names, by combined trial count. An empty list when none. | [optional] |
| **top_shared_mechanisms** | **Array&lt;String&gt;** | Up to 3 shared mechanisms written as a gene symbol and an action, by combined phase then combined trials. An empty list when none. | [optional] |

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
  competitor_rank: null,
  competitor_score: null,
  target_sim: null,
  mechanism_sim: null,
  indication_sim: null,
  phase_sim: null,
  n_shared_targets: null,
  n_shared_mechanisms: null,
  n_shared_indications: null,
  n_phase_overlap_targets: null,
  n_company_targets: null,
  n_competitor_targets: null,
  n_company_indications: null,
  n_competitor_indications: null,
  top_shared_targets: null,
  top_shared_indications: null,
  top_shared_mechanisms: null
)
```


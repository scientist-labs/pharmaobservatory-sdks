# PharmaObservatory::ProviderRwdDeliveryRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **provider_rwd_delivery_id** | **String** | The row id: a hash of provider_uuid, conductor_name_key and mandate_stratum. | [optional] |
| **provider_id** | **String** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). The entry key of the suppliers object type. A supplier is not a company. | [optional] |
| **provider_uuid** | **String** | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] |
| **provider_slug** | **String** | The supplier&#39;s marketplace slug. | [optional] |
| **provider_name** | **String** | The supplier&#39;s name. | [optional] |
| **institution_id** | **String** | The register Institution ID of the winning route to the conductor (strongest method, then lowest id). | [optional] |
| **match_method** | **String** | match_method says how the provider was tied to the register institution: org_member_exact (through the supplier&#39;s resolved organization, match_confidence carried from that match), rx_name_key_exact (an exact match of the register institution&#39;s name key to the supplier&#39;s, match_confidence null, a name match and not a verified identity) or org_unit_of_entry (an organization unit of the supplier&#39;s organization, one-directional, match_confidence the bridge&#39;s own figure). Where a provider reaches one conductor by several routes the strongest wins, in that order, and match_methods lists every route. | [optional] |
| **match_confidence** | **Float** | The bridge&#39;s own confidence for the winning route, 0 to 1; null on an rx_name_key_exact row, whose exact name match defines none. Passed through, never computed here. | [optional] |
| **conductor_name_key** | **String** | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. Not a vendor family. | [optional] |
| **matched_institution_ids** | **Array&lt;String&gt;** | Every register Institution ID that routes the provider to the conductor name key (more than one when the register lists the name twice or two methods reach it). | [optional] |
| **match_methods** | **Array&lt;String&gt;** | Every match method that routes the provider to this conductor name key; match_method is the strongest of them. | [optional] |
| **conductor_name** | **String** | The conductor&#39;s display name: one register spelling of the key. | [optional] |
| **mandate_stratum** | **String** | all (every study of the conductor), mandated (studies a regulator required) or voluntary (studies the registrant says it did not require); studies with an unknown mandate are in all only. The mandated and voluntary books slip differently: choose a stratum. | [optional] |
| **lifetime_studies** | **Integer** | Distinct registered studies the conductor is named as conducting institution on. | [optional] |
| **final_report_pairs** | **Integer** | Conducting studies with both a planned and an actual final report date. | [optional] |
| **comparable_pairs** | **Integer** | Conducting studies with a final report slip (both dates present and different); the N of the figures and the denominator of the floor. | [optional] |
| **backfilled_pairs** | **Integer** | Of the final_report_pairs, the studies whose planned and actual date are equal to the day (a likely backfill, a form artefact). Excluded from every figure, counted here. | [optional] |
| **backfilled_share** | **Float** | backfilled_pairs over final_report_pairs; null when there is no pair. | [optional] |
| **mandated_share** | **Float** | The share of the conductor&#39;s conducting studies a regulator required; read it beside any figure to see how mandated the book is. Null when the conductor has no conducting study. | [optional] |
| **n_floor** | **Integer** | The floor on comparable_pairs below which the four distribution figures are null (suppressed, not zero): 10. | [optional] |
| **meets_n_floor** | **Boolean** | True when comparable_pairs reaches n_floor, so the four figures are shown; false means they are null and the studies are in provider_rwd_delivery_study. | [optional] |
| **median_slip_days** | **Float** | Median final report slip in days against the planned date (positive is late) over the comparable pairs; null below the floor. | [optional] |
| **p25_slip_days** | **Float** | First quartile of the final report slip in days; null below the floor. | [optional] |
| **p75_slip_days** | **Float** | Third quartile of the final report slip in days; null below the floor. | [optional] |
| **late_over_30d** | **Integer** | Comparable studies whose final report slipped more than 30 days; null below the floor. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ProviderRwdDeliveryRow.new(
  provider_rwd_delivery_id: null,
  provider_id: null,
  provider_uuid: null,
  provider_slug: null,
  provider_name: null,
  institution_id: null,
  match_method: null,
  match_confidence: null,
  conductor_name_key: null,
  matched_institution_ids: null,
  match_methods: null,
  conductor_name: null,
  mandate_stratum: null,
  lifetime_studies: null,
  final_report_pairs: null,
  comparable_pairs: null,
  backfilled_pairs: null,
  backfilled_share: null,
  mandated_share: null,
  n_floor: null,
  meets_n_floor: null,
  median_slip_days: null,
  p25_slip_days: null,
  p75_slip_days: null,
  late_over_30d: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```


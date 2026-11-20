# PharmaObservatory::ProviderRwdDeliveryStudyRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **provider_rwd_delivery_study_id** | **String** | The row id: a hash of provider_uuid, conductor_name_key and study_id. | [optional] |
| **provider_id** | **String** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). The entry key of the suppliers object type. A supplier is not a company. | [optional] |
| **provider_uuid** | **String** | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] |
| **provider_slug** | **String** | The supplier&#39;s marketplace slug. | [optional] |
| **provider_name** | **String** | The supplier&#39;s name. | [optional] |
| **institution_id** | **String** | The register Institution ID of the winning route to the conductor (strongest method, then lowest id). | [optional] |
| **match_method** | **String** | match_method says how the provider was tied to the register institution: org_member_exact (through the supplier&#39;s resolved organization, match_confidence carried from that match), rx_name_key_exact (an exact match of the register institution&#39;s name key to the supplier&#39;s, match_confidence null, a name match and not a verified identity) or org_unit_of_entry (an organization unit of the supplier&#39;s organization, one-directional, match_confidence the bridge&#39;s own figure). Where a provider reaches one conductor by several routes the strongest wins, in that order, and match_methods lists every route. | [optional] |
| **match_confidence** | **Float** | The bridge&#39;s own confidence for the winning route, 0 to 1; null on an rx_name_key_exact row, whose exact name match defines none. Passed through, never computed here. | [optional] |
| **conductor_name_key** | **String** | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. Not a vendor family. | [optional] |
| **study_id** | **String** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] |
| **eu_pas_number** | **String** | The EU PAS (post-authorisation study) register number of the study. | [optional] |
| **study_mandate_stratum** | **String** | The study&#39;s own mandate: mandated (a regulator required it), voluntary, or unknown (the register does not say). | [optional] |
| **final_report_planned_on** | **Date** | Planned date of the final study report. | [optional] |
| **final_report_actual_on** | **Date** | Actual date of the final study report. | [optional] |
| **final_report_slip_days** | **Integer** | Actual minus planned final report date in days (positive is late, negative is early). A plain fact about one study, never a statistic. | [optional] |
| **all_meets_n_floor** | **Boolean** | True when the provider and conductor&#39;s all row meets the floor of 10 comparable pairs. | [optional] |
| **stratum_meets_n_floor** | **Boolean** | True when the row of the study&#39;s own mandate stratum meets the floor; null for a study whose mandate is unknown (it belongs to the all row only). | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ProviderRwdDeliveryStudyRow.new(
  provider_rwd_delivery_study_id: null,
  provider_id: null,
  provider_uuid: null,
  provider_slug: null,
  provider_name: null,
  institution_id: null,
  match_method: null,
  match_confidence: null,
  conductor_name_key: null,
  study_id: null,
  eu_pas_number: null,
  study_mandate_stratum: null,
  final_report_planned_on: null,
  final_report_actual_on: null,
  final_report_slip_days: null,
  all_meets_n_floor: null,
  stratum_meets_n_floor: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```


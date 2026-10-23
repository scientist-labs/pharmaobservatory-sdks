# PharmaObservatory::RwdConductorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **conductor_id** | **String** | The conductor slug: the register&#39;s normalised institution name key with spaces as hyphens, the id in the web page address (/rwd-conductors/&lt;id&gt;). The entry key of the rwd_conductors object type; injective because no name key carries a hyphen. | [optional] |
| **conductor_name_key** | **String** | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. | [optional] |
| **conductor_name** | **String** | The conductor&#39;s display name: one register spelling of the key. | [optional] |
| **conductor_names** | **Array&lt;String&gt;** | Every register spelling of the key, as the study lists name it. | [optional] |
| **register_institution_ids** | **Array&lt;String&gt;** | The register Institution IDs of every register row that folds to this conductor (more than one when the register lists the name twice). | [optional] |
| **register_row_count** | **Integer** | How many register institution rows fold to this conductor; above one means the register lists the name more than once and the counts are per name, not per row. | [optional] |
| **declared_office_countries** | **Array&lt;String&gt;** | The office country text the register institutions declare, one element per distinct declaration; the office is not where the conductor delivered studies (see rwd_conductor_footprint). | [optional] |
| **any_encepp_partner** | **Boolean** | True when any register row of the conductor declares itself an ENCePP (European Network of Centres for Pharmacoepidemiology and Pharmacovigilance) partner. | [optional] |
| **lifetime_studies** | **Integer** | Distinct registered studies the conductor is named as conducting institution on. Not additive with additional_role_studies. | [optional] |
| **recent_studies** | **Integer** | Conducting studies that started within five years of the snapshot. | [optional] |
| **ongoing_studies** | **Integer** | Conducting studies whose status is Ongoing. | [optional] |
| **additional_role_studies** | **Integer** | Distinct registered studies the conductor is named on as an additional institution; a different claim from conducting, never added to it. | [optional] |
| **mandated_studies** | **Integer** | Conducting studies a regulator required (mandate_stratum mandated). | [optional] |
| **mandated_share** | **Float** | mandated_studies over lifetime_studies; null when the conductor has no conducting study. | [optional] |
| **final_report_pairs** | **Integer** | Conducting studies with both a planned and an actual final report date. | [optional] |
| **backfilled_pairs** | **Integer** | Of those pairs, the studies whose planned and actual date are equal to the day (a likely backfill). | [optional] |
| **comparable_pairs** | **Integer** | Conducting studies with a final report slip (both dates present and different); the denominator of the delivery figures. | [optional] |
| **backfilled_share** | **Float** | backfilled_pairs over final_report_pairs; null when there is no pair. | [optional] |
| **n_floor** | **Integer** | The floor on comparable_pairs below which the median, quartiles and late count are null (suppressed, not zero). | [optional] |
| **meets_n_floor** | **Boolean** | True when comparable_pairs reaches n_floor, so the slip figures are shown. | [optional] |
| **median_slip_days** | **Float** | Median final report slip in days (positive is late) over the conductor&#39;s comparable pairs; null below the floor. | [optional] |
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

instance = PharmaObservatory::RwdConductorsRow.new(
  conductor_id: null,
  conductor_name_key: null,
  conductor_name: null,
  conductor_names: null,
  register_institution_ids: null,
  register_row_count: null,
  declared_office_countries: null,
  any_encepp_partner: null,
  lifetime_studies: null,
  recent_studies: null,
  ongoing_studies: null,
  additional_role_studies: null,
  mandated_studies: null,
  mandated_share: null,
  final_report_pairs: null,
  backfilled_pairs: null,
  comparable_pairs: null,
  backfilled_share: null,
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


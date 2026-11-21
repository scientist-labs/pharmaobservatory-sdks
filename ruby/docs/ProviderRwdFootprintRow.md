# PharmaObservatory::ProviderRwdFootprintRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **provider_rwd_footprint_id** | **String** | The assertion id: a hash of the provider uuid, role, register entity kind and id, match method and institution. | [optional] |
| **provider_id** | **String** | The supplier id: the Scientist.com provider id as text, the entry key of the suppliers object type. A supplier is not a company. | [optional] |
| **provider_uuid** | **String** | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] |
| **provider_slug** | **String** | The supplier&#39;s marketplace slug. | [optional] |
| **provider_name** | **String** | The supplier&#39;s name. | [optional] |
| **institution_id** | **String** | The register institution the provider was matched THROUGH. The entry key of the rwd_institutions object type. | [optional] |
| **institution_name** | **String** | The register&#39;s own name for that institution. | [optional] |
| **role** | **String** | What the register entity is to the institution: conductor or additional_institution (a study), custodian (a data source&#39;s holder) or network_member (a network&#39;s member institution). | [optional] |
| **evidence_ref_kind** | **String** | Which register entity evidence_ref names: rwd_study_id, rwd_data_source_id or rwd_network_id. | [optional] |
| **evidence_ref** | **String** | The register id of the study, data source or network, per evidence_ref_kind. Ids of different kinds can coincide; read it with evidence_ref_kind. | [optional] |
| **study_id** | **String** | The register Study ID when evidence_ref_kind is rwd_study_id, otherwise null. The entry key of the rwd_studies object type: a study route returns the providers asserted against that study. | [optional] |
| **evidence_label** | **String** | The register&#39;s own entity name (study title, data source name or network name), never prose. | [optional] |
| **catalogue_url** | **String** | The address of the register entity on the HMA-EMA catalogue site. | [optional] |
| **study_start_actual_on** | **Date** | Study rows only: the registered actual start date; null otherwise and where the study has none. | [optional] |
| **study_status** | **String** | Study rows only: the study&#39;s registered status. | [optional] |
| **mandate_stratum** | **String** | Study rows only: mandated, voluntary or unknown. | [optional] |
| **is_darwin_eu_study** | **Boolean** | Study rows only: whether the study is a DARWIN EU (the EMA&#39;s Data Analysis and Real World Interrogation Network) study. | [optional] |
| **data_source_ids** | **Array&lt;String&gt;** | Study rows only: the register data source ids the study names (the names that match a registered source). | [optional] |
| **data_source_names_unmatched** | **Array&lt;String&gt;** | Study rows only: the data source names the study gives that match no registered source. | [optional] |
| **study_country_names** | **Array&lt;String&gt;** | Study rows only: the countries the study names. | [optional] |
| **data_source_type** | **String** | Data source rows only: the data source type(s) the holder ticked, as the register exports them. | [optional] |
| **country_names** | **Array&lt;String&gt;** | Data source rows only: the countries the source covers. | [optional] |
| **match_method** | **String** | How the provider was tied to the institution: org_member_exact (through the supplier&#39;s resolved organization), rx_name_key_exact (an exact name match, not a verified identity) or org_unit_of_entry (the register entry&#39;s own statement of its parent organization). | [optional] |
| **also_matched_by_name** | **Boolean** | Whether the same provider and institution were also tied by an exact name match. | [optional] |
| **match_confidence** | **Float** | The upstream bridge&#39;s confidence figure, passed through unchanged; null on rx_name_key_exact rows. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ProviderRwdFootprintRow.new(
  provider_rwd_footprint_id: null,
  provider_id: null,
  provider_uuid: null,
  provider_slug: null,
  provider_name: null,
  institution_id: null,
  institution_name: null,
  role: null,
  evidence_ref_kind: null,
  evidence_ref: null,
  study_id: null,
  evidence_label: null,
  catalogue_url: null,
  study_start_actual_on: null,
  study_status: null,
  mandate_stratum: null,
  is_darwin_eu_study: null,
  data_source_ids: null,
  data_source_names_unmatched: null,
  study_country_names: null,
  data_source_type: null,
  country_names: null,
  match_method: null,
  also_matched_by_name: null,
  match_confidence: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```


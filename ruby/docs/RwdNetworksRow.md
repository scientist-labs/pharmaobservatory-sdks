# PharmaObservatory::RwdNetworksRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **network_id** | **String** | The register&#39;s own Network ID. The entry key of the rwd_networks object type. | [optional] |
| **network_name** | **String** | The network&#39;s name as the register spells it. | [optional] |
| **network_name_key** | **String** | The normalised name key the study and member edges join on. | [optional] |
| **register_state** | **String** | self_registered_profile when the entry declares at least one of description, countries, therapeutic areas or funding source, catalogue_name_entry otherwise. is_encepp_partner is null on a catalogue name entry. | [optional] |
| **is_self_registered_profile** | **Boolean** | True when registration_state is self_registered_profile. | [optional] |
| **declared_profile_fields** | **Integer** | How many of the declarable profile fields the entry answered. | [optional] |
| **declarable_profile_fields** | **Integer** | How many profile fields the register offers; the denominator of declared_profile_fields. | [optional] |
| **network_countries** | **String** | The country text the network declares, as typed. | [optional] |
| **primary_therapeutic_disease_areas** | **String** | The primary therapeutic and disease areas the network declares, as typed; blank on most networks, so a facet over it would remove the largest networks. | [optional] |
| **has_declared_therapeutic_area** | **Boolean** | Whether the network declares a therapeutic area at all (the blank is the registrant&#39;s silence). | [optional] |
| **sources_of_funding** | **String** | The sources of funding the network declares, as typed. | [optional] |
| **network_website** | **String** | The network&#39;s website as the register holds it. | [optional] |
| **catalogue_url** | **String** | The address of this entry on the HMA-EMA catalogue site (the register&#39;s URL LINK). | [optional] |
| **is_encepp_partner** | **Boolean** | Whether the entry declares an ENCePP partnership; null on a catalogue name entry, where the export writes a default No that is not an answer. | [optional] |
| **declared_member_count** | **Integer** | How many member institutions the network&#39;s own list names; zero says the network named nobody, not that it has no members. | [optional] |
| **declared_member_name_keys** | **Array&lt;String&gt;** | The normalised name keys of the declared members. | [optional] |
| **declared_member_names** | **Array&lt;String&gt;** | The names of the declared members as the register spells them. | [optional] |
| **lists_any_member** | **Boolean** | Whether the network&#39;s own list names at least one member. | [optional] |
| **delivered_studies_lifetime** | **Integer** | Distinct registered studies that named the network. Declared membership and delivered work are different facts and never share a column. | [optional] |
| **delivered_studies_last_5_years** | **Integer** | Delivered studies that started within five years of the snapshot. | [optional] |
| **delivered_studies_ongoing** | **Integer** | Delivered studies whose status is Ongoing. | [optional] |
| **delivered_studies_planned** | **Integer** | Delivered studies whose status is Planned. | [optional] |
| **most_recent_study_start_on** | **Date** | The latest actual start date among the studies it conducted (or delivered, for a network). | [optional] |
| **record_updated_at** | **Time** | When the registrant last edited this catalogue entry. A different clock from last_refreshed_on: it dates the web form, not the database. | [optional] |
| **record_updated_is_mass_edit** | **Boolean** | True when the record&#39;s last update date is a mass-edit date: one shared by an outsized part of the register, so it dates a migration and not the registrant. | [optional] |
| **record_updated_at_attested** | **Time** | record_updated_at with the mass-edit dates removed (null for a mass-edit date or no timestamp): the column to sort a freshness question on. | [optional] |
| **record_updated_date_share** | **Float** | The share of the register whose last update carries the same date, on every row so a mass-edit date is visible. | [optional] |
| **record_first_published_at** | **Time** | When the catalogue entry was first published. | [optional] |
| **is_outdated** | **Boolean** | EMA&#39;s own Outdated flag, carried through untouched; the register flags none of the data sources, so read false as EMA&#39;s answer and not as a currency finding. | [optional] |
| **network_register_total** | **Integer** | The number of networks in the whole network register at the snapshot. | [optional] |
| **networks_listing_any_member** | **Integer** | How many register networks list at least one member. | [optional] |
| **membership_edge_total** | **Integer** | How many network-to-institution membership edges the member lists hold. | [optional] |
| **institutions_in_any_member_list** | **Integer** | How many register institutions at least one network member list names. | [optional] |
| **institution_name_key_total** | **Integer** | How many distinct institution name keys the register holds. | [optional] |
| **networks_declaring_a_therapeutic_area** | **Integer** | How many register networks declare a therapeutic area. | [optional] |
| **networks_with_any_delivered_study** | **Integer** | How many register networks at least one registered study named. | [optional] |
| **self_registered_profile_total** | **Integer** | How many register entries are self-registered profiles. | [optional] |
| **catalogue_name_entry_total** | **Integer** | How many register entries are catalogue name entries. | [optional] |
| **self_registered_profile_encepp_yes_total** | **Integer** | How many self-registered networks declare an ENCePP partnership. | [optional] |
| **catalogue_name_entry_encepp_yes_total** | **Integer** | How many catalogue name entries declare one (the mart nulls the declaration there, so this is the false-No check). | [optional] |
| **register_source_count** | **Integer** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] |
| **study_register_total** | **Integer** | The number of studies in the whole study register at the snapshot. | [optional] |
| **studies_naming_a_registered_source** | **Integer** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdNetworksRow.new(
  network_id: null,
  network_name: null,
  network_name_key: null,
  register_state: null,
  is_self_registered_profile: null,
  declared_profile_fields: null,
  declarable_profile_fields: null,
  network_countries: null,
  primary_therapeutic_disease_areas: null,
  has_declared_therapeutic_area: null,
  sources_of_funding: null,
  network_website: null,
  catalogue_url: null,
  is_encepp_partner: null,
  declared_member_count: null,
  declared_member_name_keys: null,
  declared_member_names: null,
  lists_any_member: null,
  delivered_studies_lifetime: null,
  delivered_studies_last_5_years: null,
  delivered_studies_ongoing: null,
  delivered_studies_planned: null,
  most_recent_study_start_on: null,
  record_updated_at: null,
  record_updated_is_mass_edit: null,
  record_updated_at_attested: null,
  record_updated_date_share: null,
  record_first_published_at: null,
  is_outdated: null,
  network_register_total: null,
  networks_listing_any_member: null,
  membership_edge_total: null,
  institutions_in_any_member_list: null,
  institution_name_key_total: null,
  networks_declaring_a_therapeutic_area: null,
  networks_with_any_delivered_study: null,
  self_registered_profile_total: null,
  catalogue_name_entry_total: null,
  self_registered_profile_encepp_yes_total: null,
  catalogue_name_entry_encepp_yes_total: null,
  register_source_count: null,
  study_register_total: null,
  studies_naming_a_registered_source: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```


# PharmaObservatory::RwdInstitutionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **institution_id** | **String** | The register&#39;s own Institution ID. The entry key of the rwd_institutions object type. The register can list one name under two ids; counts are per id and are not additive across rows. | [optional] |
| **institution_name** | **String** | The institution&#39;s name as the register spells it; an institution here can be registered under a person&#39;s name. | [optional] |
| **institution_org_id** | **String** | The legal entity (organization record id) the institution resolved to; null when it did not resolve. An unresolved institution keeps its row with null ids. | [optional] |
| **institution_company_id** | **String** | The company the institution&#39;s legal entity belongs to; null when unresolved. | [optional] |
| **institution_company_name** | **String** | The company&#39;s name. | [optional] |
| **institution_group_id** | **String** | The current ultimate parent company of the institution&#39;s company. | [optional] |
| **institution_group_name** | **String** | The group&#39;s name. | [optional] |
| **org_match_rung** | **String** | How the institution name resolved to a legal entity (the resolution rung); null when unresolved. | [optional] |
| **org_match_tier** | **String** | The confidence tier of that resolution; null when unresolved. | [optional] |
| **parent_org_id** | **String** | The legal entity the register&#39;s parent institution claim resolved to. Most parent claims are claims between register entries, not between organizations: read it with is_org_to_org. | [optional] |
| **is_org_to_org** | **Boolean** | True when the parent claim is a claim between two resolved legal entities, false when it only links register entries. | [optional] |
| **institution_name_key** | **String** | The normalised name key the study, membership and holder edges join on. | [optional] |
| **parent_name** | **String** | The parent institution&#39;s name as the register spells it. | [optional] |
| **parent_name_key** | **String** | The normalised key of the parent name. | [optional] |
| **registration_state** | **String** | self_registered_profile when the entry carries at least one descriptive field, catalogue_name_entry when it is a name and an id only (a catalogue pick-list entry nobody registered a profile for). The declaration columns are null on a catalogue name entry. | [optional] |
| **is_self_registered_profile** | **Boolean** | True when registration_state is self_registered_profile. | [optional] |
| **declared_profile_fields** | **Integer** | How many of the declarable profile fields the entry answered. | [optional] |
| **declarable_profile_fields** | **Integer** | How many profile fields the register offers; the denominator of declared_profile_fields. | [optional] |
| **type_of_institution** | **String** | The type of institution the entry declares, as typed. | [optional] |
| **institution_role** | **String** | The role the entry declares for the institution, as typed. | [optional] |
| **institution_countries** | **String** | The country text the entry declares, as typed. | [optional] |
| **institution_country_names** | **Array&lt;String&gt;** | The declared countries split by the register&#39;s country list, one element per country. | [optional] |
| **institution_country_count** | **Integer** | How many countries institution_country_names holds. | [optional] |
| **unmatched_country_fragments** | **Array&lt;String&gt;** | Country text items the country vocabulary of this build does not hold. | [optional] |
| **single_declared_country_iso2** | **String** | The ISO 3166-1 alpha-2 code when the entry declares exactly one country, null otherwise. | [optional] |
| **institution_website** | **String** | The institution&#39;s website as the register holds it. | [optional] |
| **catalogue_url** | **String** | The address of this entry on the HMA-EMA catalogue site (the register&#39;s URL LINK). | [optional] |
| **open_to_industry_funded_research** | **Boolean** | Whether the institution declares it is open to industry-funded research; null on a catalogue name entry, where the export writes a default No that is not an answer. | [optional] |
| **collects_data_directly_from_patients** | **Boolean** | Whether the institution declares it collects data directly from patients; null on a catalogue name entry. | [optional] |
| **is_encepp_partner** | **Boolean** | Whether the entry declares an ENCePP partnership; null on a catalogue name entry, where the export writes a default No that is not an answer. | [optional] |
| **record_updated_at** | **Time** | When the registrant last edited this catalogue entry. A different clock from last_refreshed_on: it dates the web form, not the database. | [optional] |
| **record_updated_is_mass_edit** | **Boolean** | True when the record&#39;s last update date is a mass-edit date: one shared by an outsized part of the register, so it dates a migration and not the registrant. | [optional] |
| **record_updated_at_attested** | **Time** | record_updated_at with the mass-edit dates removed (null for a mass-edit date or no timestamp): the column to sort a freshness question on. | [optional] |
| **record_updated_date_share** | **Float** | The share of the register whose last update carries the same date, on every row so a mass-edit date is visible. | [optional] |
| **record_first_published_at** | **Time** | When the catalogue entry was first published. | [optional] |
| **is_outdated** | **Boolean** | EMA&#39;s own Outdated flag, carried through untouched; the register flags none of the data sources, so read false as EMA&#39;s answer and not as a currency finding. | [optional] |
| **conducting_studies_lifetime** | **Integer** | Distinct registered studies the institution&#39;s name conducted. Four separate activity numbers; never added together. | [optional] |
| **conducting_studies_last_5_years** | **Integer** | Conducting studies that started within five years of the snapshot. | [optional] |
| **conducting_studies_ongoing** | **Integer** | Conducting studies whose status is Ongoing. | [optional] |
| **conducting_studies_planned** | **Integer** | Conducting studies whose status is Planned. | [optional] |
| **additional_role_studies** | **Integer** | Distinct registered studies the conductor is named on as an additional institution; a different claim from conducting, never added to it. | [optional] |
| **most_recent_study_start_on** | **Date** | The latest actual start date among the studies it conducted (or delivered, for a network). | [optional] |
| **held_source_count** | **Integer** | How many registered data sources name this institution as their data holder; zero is no source naming it, not a finding about the institution. | [optional] |
| **held_source_ids** | **Array&lt;String&gt;** | The register Data source IDs of those sources. | [optional] |
| **held_active_population** | **Integer** | The sum of the active populations the held sources declare. | [optional] |
| **held_max_active_population** | **Integer** | The largest active population any held source declares. | [optional] |
| **held_sources_declaring_active_population** | **Integer** | How many held sources declare an active population. | [optional] |
| **held_qualified_source_count** | **Integer** | How many held sources declare a data source qualification. | [optional] |
| **held_darwin_partner_source_count** | **Integer** | How many held sources are DARWIN EU data partners. | [optional] |
| **network_membership_count** | **Integer** | How many networks&#39; own member lists name this institution; zero is the network&#39;s silence, not a property of the institution. | [optional] |
| **network_ids** | **Array&lt;String&gt;** | The register Network IDs of those networks. | [optional] |
| **network_names** | **Array&lt;String&gt;** | The names of those networks. | [optional] |
| **institution_register_total** | **Integer** | The number of institutions in the whole institution register at the snapshot. | [optional] |
| **self_registered_profile_total** | **Integer** | How many register entries are self-registered profiles. | [optional] |
| **catalogue_name_entry_total** | **Integer** | How many register entries are catalogue name entries. | [optional] |
| **institutions_declaring_a_country** | **Integer** | How many register institutions declare a country. | [optional] |
| **institutions_in_any_member_list** | **Integer** | How many register institutions at least one network member list names. | [optional] |
| **institution_name_key_total** | **Integer** | How many distinct institution name keys the register holds. | [optional] |
| **institutions_holding_a_source** | **Integer** | How many register institutions hold a registered data source. | [optional] |
| **declared_holder_name_total** | **Integer** | How many distinct data holder names the data source register declares. | [optional] |
| **declared_holder_name_matched** | **Integer** | How many of those holder names match a register institution. | [optional] |
| **register_source_count** | **Integer** | The number of data sources in the whole data source register at the snapshot: the denominator every source count is over. Not a count of real-world data sources in the world. | [optional] |
| **study_register_total** | **Integer** | The number of studies in the whole study register at the snapshot. | [optional] |
| **studies_naming_a_registered_source** | **Integer** | How many registered studies name at least one data source that is itself registered; the denominator under study_count. | [optional] |
| **studies_naming_a_registered_conductor** | **Integer** | How many registered studies name a conducting institution that is itself registered. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdInstitutionsRow.new(
  institution_id: null,
  institution_name: null,
  institution_org_id: null,
  institution_company_id: null,
  institution_company_name: null,
  institution_group_id: null,
  institution_group_name: null,
  org_match_rung: null,
  org_match_tier: null,
  parent_org_id: null,
  is_org_to_org: null,
  institution_name_key: null,
  parent_name: null,
  parent_name_key: null,
  registration_state: null,
  is_self_registered_profile: null,
  declared_profile_fields: null,
  declarable_profile_fields: null,
  type_of_institution: null,
  institution_role: null,
  institution_countries: null,
  institution_country_names: null,
  institution_country_count: null,
  unmatched_country_fragments: null,
  single_declared_country_iso2: null,
  institution_website: null,
  catalogue_url: null,
  open_to_industry_funded_research: null,
  collects_data_directly_from_patients: null,
  is_encepp_partner: null,
  record_updated_at: null,
  record_updated_is_mass_edit: null,
  record_updated_at_attested: null,
  record_updated_date_share: null,
  record_first_published_at: null,
  is_outdated: null,
  conducting_studies_lifetime: null,
  conducting_studies_last_5_years: null,
  conducting_studies_ongoing: null,
  conducting_studies_planned: null,
  additional_role_studies: null,
  most_recent_study_start_on: null,
  held_source_count: null,
  held_source_ids: null,
  held_active_population: null,
  held_max_active_population: null,
  held_sources_declaring_active_population: null,
  held_qualified_source_count: null,
  held_darwin_partner_source_count: null,
  network_membership_count: null,
  network_ids: null,
  network_names: null,
  institution_register_total: null,
  self_registered_profile_total: null,
  catalogue_name_entry_total: null,
  institutions_declaring_a_country: null,
  institutions_in_any_member_list: null,
  institution_name_key_total: null,
  institutions_holding_a_source: null,
  declared_holder_name_total: null,
  declared_holder_name_matched: null,
  register_source_count: null,
  study_register_total: null,
  studies_naming_a_registered_source: null,
  studies_naming_a_registered_conductor: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```


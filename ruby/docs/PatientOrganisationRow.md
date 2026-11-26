# PharmaObservatory::PatientOrganisationRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_kind** | **String** | Which kind of page the row is listed on: rare_disease (a rare disease page, entered on orpha_code) or condition (a condition page, entered on condition_id). | [optional] |
| **target_key** | **String** | The orpha_code of a rare_disease row or the condition_id of a condition row: the key the row is entered on. Always present, so it can stand in a sort or a cursor. | [optional] |
| **orpha_code** | **String** | The ORPHA code of the disease the row is listed under, the digits in the web page address (/rare-diseases/558); null on a condition row. The entry key of the rare_diseases object type. | [optional] |
| **condition_id** | **String** | The MeSH descriptor id of the condition the row is listed under (D followed by six to nine digits); null on a rare disease row. The entry key of the conditions object type. | [optional] |
| **patient_organisation_id** | **String** | The organisation&#39;s id: a spine organisation id (edgar:, ror:, sponsor: and the like) when the organisation folded onto an organisation we already hold, else pag:&lt;registrable domain of its website&gt;. | [optional] |
| **org_id** | **String** | The legal-entity id of the organisation when it folded onto one we already hold, the join to legal_entities; null for a pag: organisation, which has no spine record. | [optional] |
| **name** | **String** | The organisation&#39;s name as its own website states it. | [optional] |
| **website_url** | **String** | The organisation&#39;s website as it states it. | [optional] |
| **country** | **String** | The country the organisation declares (ISO 3166-1 alpha-2), else its primary location&#39;s country. | [optional] |
| **region** | **String** | The region of its primary location. | [optional] |
| **city** | **String** | The city of its primary location. | [optional] |
| **location_count** | **Integer** | How many locations the organisation lists. | [optional] |
| **languages** | **Array&lt;String&gt;** | The languages the organisation says it works in, as it states them. | [optional] |
| **phones** | **Array&lt;String&gt;** | The organisation&#39;s own phone numbers, as stated. Organisation-level channels only: no person&#39;s line. | [optional] |
| **emails** | **Array&lt;String&gt;** | The organisation&#39;s own email addresses, as stated, minus any address whose local part looks like a person&#39;s name. Organisation-level channels only. | [optional] |
| **contact_urls** | **Array&lt;String&gt;** | The organisation&#39;s own contact page addresses. | [optional] |
| **social_urls** | **Array&lt;String&gt;** | The organisation&#39;s own social media addresses. | [optional] |
| **service_codes** | **Array&lt;String&gt;** | The services the organisation&#39;s own pages say it offers, from a closed list: support_groups, helpline, information_and_education, patient_registry, research_funding, advocacy_and_policy, financial_assistance, care_navigation, clinical_trial_matching, other. | [optional] |
| **other_service_labels** | **Array&lt;String&gt;** | For a service coded other, the organisation&#39;s own words for it. Empty when it states none. | [optional] |
| **link_kind** | **String** | How the organisation reaches the disease or condition the row is listed under: declared (its own page names it), mapped (a condition row: it declares a rare disease this condition is mapped to; the mapped_* columns say which), broader (a rare disease row: it declared a broader parent disease) or narrower (a rare disease row: it declared a narrower child disease). A broader or narrower row never says the organisation declares THIS disease. | [optional] |
| **link_rank** | **Integer** | The page&#39;s group order: 1 declared, 2 mapped, 3 broader, 4 narrower. | [optional] |
| **declared_scope** | **String** | How much of its work the organisation says the disease is: sole_focus, one_of_several or umbrella; null when it states none. | [optional] |
| **declared_name** | **String** | The disease name the organisation actually declared. On a broader, narrower or mapped row it is NOT this page&#39;s disease. | [optional] |
| **mapped_orpha_code** | **String** | On a mapped row, the ORPHA code of the declared rare disease the condition is mapped to; null otherwise. | [optional] |
| **mapped_disease_name** | **String** | On a mapped row, the preferred name of that rare disease; null otherwise. | [optional] |
| **mapping_relation** | **String** | On a mapped row, how the rare disease relates to the condition (exact, broader or narrower); null otherwise. | [optional] |
| **source_url** | **String** | The page of the organisation&#39;s own website the declaration came from. | [optional] |
| **retrieved_on** | **Date** | The date that page was retrieved. | [optional] |
| **credit** | **String** | The source credit line for the row: Source: &lt;organisation name&gt;, &lt;page address&gt;, retrieved &lt;date&gt;. Mandatory wherever the row is shown. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::PatientOrganisationRow.new(
  target_kind: null,
  target_key: null,
  orpha_code: null,
  condition_id: null,
  patient_organisation_id: null,
  org_id: null,
  name: null,
  website_url: null,
  country: null,
  region: null,
  city: null,
  location_count: null,
  languages: null,
  phones: null,
  emails: null,
  contact_urls: null,
  social_urls: null,
  service_codes: null,
  other_service_labels: null,
  link_kind: null,
  link_rank: null,
  declared_scope: null,
  declared_name: null,
  mapped_orpha_code: null,
  mapped_disease_name: null,
  mapping_relation: null,
  source_url: null,
  retrieved_on: null,
  credit: null
)
```


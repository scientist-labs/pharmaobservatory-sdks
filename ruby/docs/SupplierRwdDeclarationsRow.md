# PharmaObservatory::SupplierRwdDeclarationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **provider_rwd_declared_source_id** | **String** | The fact id: a hash of the supplier uuid, the fact kind and the fact reference. | [optional] |
| **provider_id** | **String** | The supplier id: the Scientist.com provider id as text, the entry key of the suppliers object type. A supplier is not a company. | [optional] |
| **provider_uuid** | **String** | The Scientist.com public provider uuid, the join key to the RX marketplace. | [optional] |
| **provider_slug** | **String** | The supplier&#39;s marketplace slug. | [optional] |
| **provider_name** | **String** | The supplier&#39;s name. | [optional] |
| **role** | **String** | Always supplier_declared: the supplier said it, on the Scientist.com questionnaire. It is never a register role. | [optional] |
| **evidence_ref_kind** | **String** | Which kind of fact the row is: rx_rfi_owned_database (a database the supplier lists), rx_rfi_data_access (countries and regions it declares data access in, naming no source) or rx_rfi_therapeutic_experience (rare-disease and therapeutic-area experience). | [optional] |
| **evidence_ref** | **String** | The fact reference, per evidence_ref_kind: the questionnaire id with the item position for a database, the questionnaire id otherwise. Read it with evidence_ref_kind. | [optional] |
| **evidence_label** | **String** | The database short name the supplier gave (capped at 120 characters), or a fixed label for the two other kinds. | [optional] |
| **match_method** | **String** | How the fact was tied to the supplier: rx_provider_direct, the questionnaire carries the supplier id and nothing was inferred. | [optional] |
| **country_names** | **Array&lt;String&gt;** | Database rows: the database country. Access rows: the countries the supplier declares access in. Null on therapeutic rows. | [optional] |
| **declared_regions** | **Array&lt;String&gt;** | Access rows: the regions the supplier declares access in, as the questionnaire writes them. Null elsewhere. | [optional] |
| **data_types** | **Array&lt;String&gt;** | The data types, mapped to the current label of each option (a label with no mapping is kept as ticked). Database rows: types the source holds. Access rows: types the supplier declares experience with. | [optional] |
| **data_source_type** | **String** | data_types joined with &#39; | &#39;. | [optional] |
| **data_types_basis** | **String** | What data_types means: source_holds (database rows) or supplier_experience (access rows, which name no source). Null where there are no data types. | [optional] |
| **declared_validation** | **String** | Database rows: the supplier&#39;s own statement of validation, declared_validated, declared_not_validated or declared_validation_in_progress. | [optional] |
| **declared_first_collection_year** | **Integer** | Database rows: the first-collection year the supplier gave, kept where it is between 1950 and the build year, else null. | [optional] |
| **year_implausible** | **Boolean** | Database rows: true where the supplier gave a year outside 1950 to the build year (declared_first_collection_year is then null). | [optional] |
| **declared_linkable** | **Boolean** | Database rows: whether the supplier says the database can be linked; null where it did not answer. | [optional] |
| **therapeutic_areas** | **Array&lt;String&gt;** | Therapeutic rows: the therapeutic areas ticked, current labels. An empty list is a declared negative, not missing data. | [optional] |
| **declared_rare_disease_experience** | **Boolean** | Therapeutic rows: true where the supplier says it has rare-disease experience, false where it says it has none (a declared negative), null where it did not answer. | [optional] |
| **as_of** | **Date** | The date the questionnaire was submitted. | [optional] |
| **is_expired** | **Boolean** | Whether the questionnaire is past its 730-day validity at build time. | [optional] |
| **expired_at** | **Time** | When the questionnaire validity ends. | [optional] |
| **questionnaire_key** | **String** | The questionnaire version the fact came from: rwe_current, or rwe_2019 on a database row. | [optional] |
| **questionnaire_label** | **String** | The questionnaire version as a label. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the source: none, a first-party source with no license restriction. | [optional] |
| **attribution** | **String** | The plain source line to show with a row. | [optional] |
| **authority_class** | **String** | Always declared: the supplier&#39;s own statement, not verified. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SupplierRwdDeclarationsRow.new(
  provider_rwd_declared_source_id: null,
  provider_id: null,
  provider_uuid: null,
  provider_slug: null,
  provider_name: null,
  role: null,
  evidence_ref_kind: null,
  evidence_ref: null,
  evidence_label: null,
  match_method: null,
  country_names: null,
  declared_regions: null,
  data_types: null,
  data_source_type: null,
  data_types_basis: null,
  declared_validation: null,
  declared_first_collection_year: null,
  year_implausible: null,
  declared_linkable: null,
  therapeutic_areas: null,
  declared_rare_disease_experience: null,
  as_of: null,
  is_expired: null,
  expired_at: null,
  questionnaire_key: null,
  questionnaire_label: null,
  restriction_tier: null,
  attribution: null,
  authority_class: null
)
```


# PharmaObservatory::RwdStudyConditionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **study_id** | **String** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] |
| **condition_role** | **String** | primary when the term is the study&#39;s medicinal condition to be studied, additional when it is in the additional conditions list. | [optional] |
| **condition_term** | **String** | The condition term as the register has it. | [optional] |
| **mention_count** | **Integer** | How many times the term is mentioned in the study. | [optional] |
| **condition_id** | **String** | The condition&#39;s MeSH descriptor id (D followed by six to nine digits), the id of the conditions object type. | [optional] |
| **mesh_term** | **String** | The MeSH heading the term resolved to; null when it did not resolve. | [optional] |
| **mesh_tree_category** | **String** | The MeSH tree category the condition sits under. | [optional] |
| **is_disease_descriptor** | **Boolean** | Whether the MeSH descriptor is a disease descriptor. | [optional] |
| **mesh_match_rung** | **String** | How the term resolved to MeSH (curated, entry_term and so on, or unresolved). | [optional] |
| **is_mesh_resolved** | **Boolean** | Whether the term resolved to a MeSH descriptor. | [optional] |
| **study_status** | **String** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] |
| **mandate_stratum** | **String** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] |
| **required_by_regulator** | **String** | The registrant&#39;s answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional] |
| **data_source_count** | **Integer** | How many data source items the study names. | [optional] |
| **conducting_institution_count** | **Integer** | How many conducting institutions the study names. | [optional] |
| **snapshot_label** | **String** | The label of the catalogue snapshot the row was read from (a date). | [optional] |
| **accessed_month** | **String** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdStudyConditionsRow.new(
  study_id: null,
  condition_role: null,
  condition_term: null,
  mention_count: null,
  condition_id: null,
  mesh_term: null,
  mesh_tree_category: null,
  is_disease_descriptor: null,
  mesh_match_rung: null,
  is_mesh_resolved: null,
  study_status: null,
  mandate_stratum: null,
  required_by_regulator: null,
  data_source_count: null,
  conducting_institution_count: null,
  snapshot_label: null,
  accessed_month: null,
  restriction_tier: null,
  attribution: null
)
```


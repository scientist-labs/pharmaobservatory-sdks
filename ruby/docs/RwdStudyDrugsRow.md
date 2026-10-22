# PharmaObservatory::RwdStudyDrugsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **study_id** | **String** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] |
| **mention_source** | **String** | Which study field the drug mention came from: inn or atc_code. | [optional] |
| **item_ordinal** | **Integer** | The position of the item in the register&#39;s own list, from 1. | [optional] |
| **drug_mention** | **String** | The drug mention as typed. | [optional] |
| **source_list** | **String** | The whole list the mention was split from. | [optional] |
| **atc_code** | **String** | The ATC code of the mention. | [optional] |
| **atc_level5_code** | **String** | The full ATC level 5 code (seven characters, a chemical substance), when the mention carries one. | [optional] |
| **atc_level1_code** | **String** | The ATC level 1 code (one letter, the anatomical main group). | [optional] |
| **atc_label** | **String** | The substance name the mention carries beside the code. | [optional] |
| **inn_name** | **String** | The INN of the mention as typed. | [optional] |
| **inn_name_key** | **String** | The normalised INN key. | [optional] |
| **match_rung** | **String** | How the mention resolved to a canonical drug: atc_level5 or inn_name when it reached exactly one, ambiguous when several tied, unresolved when none, not_applicable when the mention has nothing to match. | [optional] |
| **drug_id** | **String** | The canonical drug id (for example unii:53AXN4NNHX) when the mention reached exactly one drug; null when unresolved or ambiguous (candidate_drug_ids holds the tie). | [optional] |
| **candidate_drug_ids** | **Array&lt;String&gt;** | The candidate canonical drug ids when the mention is ambiguous. | [optional] |
| **candidate_drug_count** | **Integer** | How many candidates. | [optional] |
| **restriction_tier** | **String** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] |
| **attribution** | **String** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdStudyDrugsRow.new(
  study_id: null,
  mention_source: null,
  item_ordinal: null,
  drug_mention: null,
  source_list: null,
  atc_code: null,
  atc_level5_code: null,
  atc_level1_code: null,
  atc_label: null,
  inn_name: null,
  inn_name_key: null,
  match_rung: null,
  drug_id: null,
  candidate_drug_ids: null,
  candidate_drug_count: null,
  restriction_tier: null,
  attribution: null
)
```


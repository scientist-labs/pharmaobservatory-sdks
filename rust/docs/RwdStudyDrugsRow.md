# RwdStudyDrugsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**study_id** | Option<**String**> | The register's own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional]
**mention_source** | Option<**String**> | Which study field the drug mention came from: inn or atc_code. | [optional]
**item_ordinal** | Option<**i32**> | The position of the item in the register's own list, from 1. | [optional]
**drug_mention** | Option<**String**> | The drug mention as typed. | [optional]
**source_list** | Option<**String**> | The whole list the mention was split from. | [optional]
**atc_code** | Option<**String**> | The ATC code of the mention. | [optional]
**atc_level5_code** | Option<**String**> | The full ATC level 5 code (seven characters, a chemical substance), when the mention carries one. | [optional]
**atc_level1_code** | Option<**String**> | The ATC level 1 code (one letter, the anatomical main group). | [optional]
**atc_label** | Option<**String**> | The substance name the mention carries beside the code. | [optional]
**inn_name** | Option<**String**> | The INN of the mention as typed. | [optional]
**inn_name_key** | Option<**String**> | The normalised INN key. | [optional]
**match_rung** | Option<**String**> | How the mention resolved to a canonical drug: atc_level5 or inn_name when it reached exactly one, ambiguous when several tied, unresolved when none, not_applicable when the mention has nothing to match. | [optional]
**drug_id** | Option<**String**> | The canonical drug id (for example unii:53AXN4NNHX) when the mention reached exactly one drug; null when unresolved or ambiguous (candidate_drug_ids holds the tie). | [optional]
**candidate_drug_ids** | Option<**Vec<String>**> | The candidate canonical drug ids when the mention is ambiguous. | [optional]
**candidate_drug_count** | Option<**i32**> | How many candidates. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



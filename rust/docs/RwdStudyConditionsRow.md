# RwdStudyConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**study_id** | Option<**String**> | The register's own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional]
**condition_role** | Option<**String**> | primary when the term is the study's medicinal condition to be studied, additional when it is in the additional conditions list. | [optional]
**condition_term** | Option<**String**> | The condition term as the register has it. | [optional]
**mention_count** | Option<**i32**> | How many times the term is mentioned in the study. | [optional]
**condition_id** | Option<**String**> | The condition's MeSH descriptor id (D followed by six to nine digits), the id of the conditions object type. | [optional]
**mesh_term** | Option<**String**> | The MeSH heading the term resolved to; null when it did not resolve. | [optional]
**mesh_tree_category** | Option<**String**> | The MeSH tree category the condition sits under. | [optional]
**is_disease_descriptor** | Option<**bool**> | Whether the MeSH descriptor is a disease descriptor. | [optional]
**mesh_match_rung** | Option<**String**> | How the term resolved to MeSH (curated, entry_term and so on, or unresolved). | [optional]
**is_mesh_resolved** | Option<**bool**> | Whether the term resolved to a MeSH descriptor. | [optional]
**study_status** | Option<**String**> | The study's status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional]
**mandate_stratum** | Option<**String**> | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional]
**required_by_regulator** | Option<**String**> | The registrant's answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional]
**data_source_count** | Option<**i32**> | How many data source items the study names. | [optional]
**conducting_institution_count** | Option<**i32**> | How many conducting institutions the study names. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



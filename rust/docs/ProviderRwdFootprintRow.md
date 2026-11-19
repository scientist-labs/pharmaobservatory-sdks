# ProviderRwdFootprintRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**provider_rwd_footprint_id** | Option<**String**> | The assertion id: a hash of the provider uuid, role, register entity kind and id, match method and institution. | [optional]
**provider_id** | Option<**String**> | The supplier id: the Scientist.com provider id as text, the entry key of the suppliers object type. A supplier is not a company. | [optional]
**provider_uuid** | Option<**String**> | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional]
**provider_slug** | Option<**String**> | The supplier's marketplace slug. | [optional]
**provider_name** | Option<**String**> | The supplier's name. | [optional]
**institution_id** | Option<**String**> | The register institution the provider was matched THROUGH. The entry key of the rwd_institutions object type. | [optional]
**institution_name** | Option<**String**> | The register's own name for that institution. | [optional]
**role** | Option<**String**> | What the register entity is to the institution: conductor or additional_institution (a study), custodian (a data source's holder) or network_member (a network's member institution). | [optional]
**evidence_ref_kind** | Option<**String**> | Which register entity evidence_ref names: rwd_study_id, rwd_data_source_id or rwd_network_id. | [optional]
**evidence_ref** | Option<**String**> | The register id of the study, data source or network, per evidence_ref_kind. Ids of different kinds can coincide; read it with evidence_ref_kind. | [optional]
**evidence_label** | Option<**String**> | The register's own entity name (study title, data source name or network name), never prose. | [optional]
**catalogue_url** | Option<**String**> | The address of the register entity on the HMA-EMA catalogue site. | [optional]
**study_start_actual_on** | Option<[**String**](string.md)> | Study rows only: the registered actual start date; null otherwise and where the study has none. | [optional]
**study_status** | Option<**String**> | Study rows only: the study's registered status. | [optional]
**mandate_stratum** | Option<**String**> | Study rows only: mandated, voluntary or unknown. | [optional]
**is_darwin_eu_study** | Option<**bool**> | Study rows only: whether the study is a DARWIN EU (the EMA's Data Analysis and Real World Interrogation Network) study. | [optional]
**data_source_ids** | Option<**Vec<String>**> | Study rows only: the register data source ids the study names (the names that match a registered source). | [optional]
**data_source_names_unmatched** | Option<**Vec<String>**> | Study rows only: the data source names the study gives that match no registered source. | [optional]
**study_country_names** | Option<**Vec<String>**> | Study rows only: the countries the study names. | [optional]
**data_source_type** | Option<**String**> | Data source rows only: the data source type(s) the holder ticked, as the register exports them. | [optional]
**country_names** | Option<**Vec<String>**> | Data source rows only: the countries the source covers. | [optional]
**match_method** | Option<**String**> | How the provider was tied to the institution: org_member_exact (through the supplier's resolved organization), rx_name_key_exact (an exact name match, not a verified identity) or org_unit_of_entry (the register entry's own statement of its parent organization). | [optional]
**also_matched_by_name** | Option<**bool**> | Whether the same provider and institution were also tied by an exact name match. | [optional]
**match_confidence** | Option<**f64**> | The upstream bridge's confidence figure, passed through unchanged; null on rx_name_key_exact rows. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



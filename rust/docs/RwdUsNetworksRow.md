# RwdUsNetworksRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_key** | Option<**String**> | The network key, the entry key of the rwd_us_networks object type (n3c today). | [optional]
**network_source** | Option<**String**> | The source the network rows come from. | [optional]
**network_name** | Option<**String**> | The network's name. | [optional]
**network_operator** | Option<**String**> | The operating body. | [optional]
**member_total** | Option<**i32**> | Institutions holding at least one published agreement, our count of the register. | [optional]
**members_with_dta** | Option<**i32**> | Institutions holding a Data Transfer Agreement. | [optional]
**members_with_current_dta** | Option<**i32**> | Institutions with a DTA row whose published term end is on or after the retrieval date. | [optional]
**members_with_dua** | Option<**i32**> | Institutions holding a Data Use Agreement. | [optional]
**members_with_current_dua** | Option<**i32**> | Institutions with a DUA row within its published term. | [optional]
**members_with_lhba** | Option<**i32**> | Institutions holding an agreement published with the type code LHBA. | [optional]
**members_resolved_org** | Option<**i32**> | Members resolved to an organization on the spine. | [optional]
**members_matched_rx** | Option<**i32**> | Members matched to at least one RX provider, by an exact method visible in rwd_us_network_memberships. | [optional]
**published_dta_count** | Option<**i32**> | N3C's own published count of DTAs, as N3C labels it. Differs from ours; not reconciled. | [optional]
**published_dua_count** | Option<**i32**> | N3C's own published count of DUAs. Differs from ours; not reconciled. | [optional]
**published_institution_count** | Option<**i32**> | N3C's own published count of institutions. Differs from ours; not reconciled. | [optional]
**published_project_count** | Option<**i32**> | N3C's own published count of projects. Differs from ours; not reconciled. | [optional]
**as_of** | Option<[**String**](string.md)> | The retrieval date of the agreement snapshot. | [optional]
**source_url** | Option<**String**> | The API endpoint the agreements came from. | [optional]
**agreements_url** | Option<**String**> | The public N3C dashboard. | [optional]
**signatories_page_url** | Option<**String**> | N3C's HTML signatories page, kept as a documented reconciliation and not ingested. | [optional]
**retrieved_at** | Option<**String**> | When the snapshot was retrieved, in UTC. | [optional]
**attribution** | Option<**String**> | The credit line every surface carrying this data prints. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



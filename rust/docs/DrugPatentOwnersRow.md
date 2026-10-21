# DrugPatentOwnersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**owner_org_id** | Option<**String**> | Legal-entity level: the owning organization record as filed. | [optional]
**owner_name** | Option<**String**> | The owner's name. | [optional]
**owner_company_id** | Option<**String**> | Company level: the owner's company (null when the owner is not in the spine). | [optional]
**owner_company_name** | Option<**String**> | Display name of the owner's company. | [optional]
**owner_group_id** | Option<**String**> | Group level: the owner's current ultimate parent company; a root company's group is itself. | [optional]
**owner_group_name** | Option<**String**> | Display name of the owner's group. | [optional]
**org_in_spine** | Option<**bool**> | True when the owner is an organization record in the spine. | [optional]
**n_patents** | Option<**i32**> | Distinct declared patents the owner holds for the drug. | [optional]
**earliest_expiry** | Option<[**String**](string.md)> | The earliest declared expiry. | [optional]
**latest_expiry** | Option<[**String**](string.md)> | The latest declared expiry. | [optional]
**best_match_confidence** | Option<**String**> | The best assignee-to-organization match confidence (high, med, low). | [optional]
**link_bases** | Option<**Vec<String>**> | The records that tie the patents to the drug (orange_book, uspto_pte). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



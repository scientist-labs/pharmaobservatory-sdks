# DrugDeclaredPatentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**patent_id** | Option<**String**> | The patent id in the patent dimension. | [optional]
**assignee_sequence** | Option<**i32**> | The assignee's order on the patent; -1 when there is none. | [optional]
**patent_number** | Option<**String**> | The patent number. | [optional]
**patent_title** | Option<**String**> | The patent title. | [optional]
**patent_date** | Option<[**String**](string.md)> | The date the patent was granted. | [optional]
**declared_expiry** | Option<[**String**](string.md)> | The expiry declared for the drug. | [optional]
**appl_nos** | Option<**Vec<String>**> | The drug application numbers the patent is declared against. | [optional]
**link_basis** | Option<**String**> | The record that ties the patent to the drug (orange_book or uspto_pte). | [optional]
**link_bases** | Option<**Vec<String>**> | Every record that ties the patent to the drug. | [optional]
**has_ped_extension** | Option<**bool**> | True when the Orange Book lists a pediatric exclusivity extension. | [optional]
**assignee_organization** | Option<**String**> | The assignee as filed. | [optional]
**assignee_org_id** | Option<**String**> | Legal-entity level: the assignee's organization record (null when unresolved). | [optional]
**assignee_company_id** | Option<**String**> | Company level: the assignee's company (null when unresolved). | [optional]
**assignee_company_name** | Option<**String**> | Display name of the assignee's company. | [optional]
**assignee_group_id** | Option<**String**> | Group level: the assignee's current ultimate parent company; a root company's group is itself. | [optional]
**assignee_group_name** | Option<**String**> | Display name of the assignee's group. | [optional]
**org_in_spine** | Option<**bool**> | True when the assignee is an organization record in the spine. | [optional]
**resolution_method** | Option<**String**> | How the assignee was matched to an organization. | [optional]
**match_confidence** | Option<**String**> | The match confidence (high, med, low). | [optional]
**owner_basis** | Option<**String**> | Why the assignee is, or is not, on record. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



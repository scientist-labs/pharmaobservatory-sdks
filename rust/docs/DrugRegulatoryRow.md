# DrugRegulatoryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**fda_approved** | Option<**bool**> | True when the drug has a record in Drugs@FDA. | [optional]
**in_ndc** | Option<**bool**> | True when the drug has a package in the National Drug Code (NDC) directory. | [optional]
**has_generic** | Option<**bool**> | True when a generic is on record. | [optional]
**biosimilar_available** | Option<**bool**> | True when the Purple Book lists a biosimilar. | [optional]
**interchangeable_available** | Option<**bool**> | True when the Purple Book lists an interchangeable biosimilar. | [optional]
**first_licensure_date** | Option<[**String**](string.md)> | The earliest biologic licensure date in the Purple Book. | [optional]
**pte_expiry** | Option<[**String**](string.md)> | The extended expiry of a granted patent term extension. | [optional]
**patent_protected** | Option<**bool**> | True when an Orange Book patent is still active. | [optional]
**latest_patent_expiry** | Option<[**String**](string.md)> | The latest expiry among the Orange Book patents. | [optional]
**applicant_name** | Option<**String**> | The applicant (company) as filed in the Orange or Purple Book. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



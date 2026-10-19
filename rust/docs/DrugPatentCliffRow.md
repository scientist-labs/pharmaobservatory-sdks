# DrugPatentCliffRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**drug_id** | Option<**String**> | The drug key. | [optional]
**drug_name** | Option<**String**> | The drug's name. | [optional]
**pharm_class** | Option<**String**> | The drug's pharmacologic class. | [optional]
**atc_class** | Option<**String**> | The drug's Anatomical Therapeutic Chemical (ATC) class. | [optional]
**brand** | Option<**String**> | The brand name. | [optional]
**candidate_brands** | Option<**String**> | Other candidate brand names, comma separated. | [optional]
**cliff_date** | Option<[**String**](string.md)> | The date exclusivity is expected to end (loss of exclusivity); always in the future when the dataset is built. | [optional]
**cliff_year** | Option<**i32**> | The year of cliff_date. | [optional]
**cliff_basis** | Option<**String**> | Which regime set the date: curated_loe, orange_book_patent, pte_expiry or bpcia_exclusivity. | [optional]
**cliff_source** | Option<**String**> | Source of a curated date (null on the other bases). | [optional]
**cliff_confidence** | Option<**String**> | Confidence of a curated date (null on the other bases). | [optional]
**has_generic** | Option<**bool**> | True when a generic is already marketed. | [optional]
**biosimilar_available** | Option<**bool**> | True when a biosimilar is already marketed. | [optional]
**applicant_name** | Option<**String**> | The brand or originator applicant as the Orange Book or Purple Book names it. | [optional]
**applicant_org_id** | Option<**String**> | Legal-entity level: the organization record the applicant resolved to (null when unresolved). | [optional]
**applicant_company_id** | Option<**String**> | Company level: the company the applicant belongs to (null when unresolved). | [optional]
**applicant_company_name** | Option<**String**> | Display name of the applicant's company. | [optional]
**applicant_group_id** | Option<**String**> | Group level: the applicant's current ultimate parent company; a root company's group is itself. | [optional]
**applicant_group_name** | Option<**String**> | Display name of the applicant's group. | [optional]
**applicant_match_scope** | Option<**String**> | How the applicant name matched an organization record. | [optional]
**medicare_spend** | Option<**f64**> | Latest-year Medicare Part B and D spend for the drug, in US dollars. | [optional]
**medicare_spend_year** | Option<**i32**> | The year of medicare_spend. | [optional]
**disclosed_revenue_usd_millions** | Option<**f64**> | Latest full-year disclosed worldwide revenue, in US millions. | [optional]
**disclosed_revenue_year** | Option<**i32**> | The fiscal year of the disclosed revenue. | [optional]
**size_usd** | Option<**f64**> | The single size axis in US dollars: disclosed revenue when present, else Medicare spend. | [optional]
**size_basis** | Option<**String**> | Which source size_usd used: disclosed_worldwide_revenue, medicare_spend or none. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



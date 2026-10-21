# BiomarkerPublishingCompaniesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**route_key** | Option<**String**> | The biomarker route key of the gene the papers name (the gene symbol); the entry key for the biomarkers object type. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol. | [optional]
**publisher_company_id** | Option<**String**> | Company level: the publisher's company. | [optional]
**publisher_company_name** | Option<**String**> | Display name of the publisher's company. | [optional]
**publisher_group_id** | Option<**String**> | Group level: the publisher's current ultimate parent company; a root company's group is itself. | [optional]
**publisher_group_name** | Option<**String**> | Display name of the publisher's group. | [optional]
**n_publications** | Option<**i32**> | Papers that name the gene from the company's organizations (a consortium paper is excluded). | [optional]
**n_member_orgs** | Option<**i32**> | The company's organizations the papers reach. | [optional]
**best_match_tier** | Option<**String**> | Strongest affiliation match tier. | [optional]
**first_pub_year** | Option<**i32**> | Earliest publication year. | [optional]
**last_pub_year** | Option<**i32**> | Latest publication year. | [optional]
**supporting_pmids** | Option<**Vec<String>**> | A sample of the supporting PubMed ids. | [optional]
**supporting_pmids_truncated** | Option<**bool**> | True when the sample is shorter than the paper count. | [optional]
**n_papers_excluded_as_consortium** | Option<**i32**> | Papers left out because they name too many organizations. | [optional]
**corpus_articles** | Option<**i32**> | Articles in the trial-cited publication set the publishing counts are drawn from (not all of PubMed). Corpus-wide, the same value on every row, never per biomarker or per company. | [optional]
**n_affiliations** | Option<**i32**> | Distinct author affiliation strings in the corpus. Corpus-wide, the same value on every row. | [optional]
**n_resolved_affiliations** | Option<**i32**> | Affiliation strings that resolved to an organization. Corpus-wide, the same value on every row. | [optional]
**affiliation_resolved_share** | Option<**f64**> | n_resolved_affiliations divided by n_affiliations, rounded to four places. Corpus-wide, the same value on every row: how much of the corpus the company attribution can see, not a figure about this biomarker. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



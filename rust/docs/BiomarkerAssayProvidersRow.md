# BiomarkerAssayProvidersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**route_key** | Option<**String**> | The route key of the gene the biomarker resolved to, the entry key for the biomarkers object type. Null when the extraction did not resolve a gene. | [optional]
**biomarker_key** | Option<**String**> | The grouping key of the biomarker. The biomarker's key when its gene resolved, otherwise raw: followed by the lower-cased name as written, so two spellings of one unresolved biomarker are two keys. | [optional]
**biomarker_id** | Option<**String**> | The biomarker's own key; null for a pair whose gene did not resolve. | [optional]
**biomarker_display_name** | Option<**String**> | The biomarker's name, as resolved or else as extracted. | [optional]
**gene_resolution_status** | Option<**String**> | resolved: the gene resolved. unresolved: a gene symbol was extracted and matched nothing. not_applicable: no gene symbol was extracted. | [optional]
**provider_key** | Option<**String**> | The provider the rows are grouped on. org: followed by the organization key when the company resolves to an organization, otherwise name: followed by the normalised company name. | [optional]
**provider_role** | Option<**String**> | What the papers name the provider as. assay_or_test_maker makes the assay or test; testing_service runs it as a service; both when the pair has papers of each kind. | [optional]
**provider_name** | Option<**String**> | The provider company name as it appears in the papers. | [optional]
**provider_org_id** | Option<**String**> | The provider's legal entity key, where the marketplace match resolves one. | [optional]
**provider_org_match_tier** | Option<**String**> | How the provider was resolved to an organization. rx_provider_link: through its marketplace supplier. unresolved: no organization, the provider stays a name. The remaining values are reserved for resolution rungs that have not shipped. | [optional]
**provider_parent_org_id** | Option<**String**> | The ultimate parent organization of the provider's organization; null when the provider has no organization. | [optional]
**provider_company_id** | Option<**String**> | The provider's company key. | [optional]
**provider_company_name** | Option<**String**> | The provider's company name. | [optional]
**provider_group_id** | Option<**String**> | The provider's current ultimate parent group key. | [optional]
**provider_group_name** | Option<**String**> | The provider's group name. | [optional]
**rx_matched** | Option<**bool**> | Whether the provider name matches a supplier in the Scientist.com marketplace. | [optional]
**rx_provider_id** | Option<**String**> | The matched marketplace supplier's id; join to the suppliers dataset. | [optional]
**rx_provider_name** | Option<**String**> | The matched marketplace supplier's name. | [optional]
**rx_provider_slug** | Option<**String**> | The matched supplier's marketplace profile slug. | [optional]
**has_clinical_validation** | Option<**bool**> | Whether any paper for the pair mentions clinical validation of the assay. | [optional]
**has_analytical_validation** | Option<**bool**> | Whether any paper for the pair mentions analytical validation of the assay. | [optional]
**n_papers** | Option<**i64**> | Papers naming the biomarker and the provider, dated or not. | [optional]
**n_dated_papers** | Option<**i64**> | Papers with a recoverable publication date. | [optional]
**n_undated_papers** | Option<**i64**> | Papers with no recoverable publication date. n_dated_papers plus n_undated_papers is n_papers. | [optional]
**most_recent_published_at** | Option<**String**> | The latest publication date among the dated papers, in UTC. Null when no paper is dated, which is not a statement that none was published recently. | [optional]
**most_recent_pmid** | Option<**String**> | The PubMed id of the most recent dated paper. | [optional]
**most_recent_title** | Option<**String**> | The title of that paper. | [optional]
**meets_curated_threshold** | Option<**bool**> | True when the provider matches a marketplace supplier or at least two papers name the pair. | [optional]
**assay_provider_class** | Option<**String**> | The company-level class. diagnostic_provider makes or supplies the diagnostic assay; reagent_supplier is the vendor whose kit appears in a methods section; context_dependent, pharma_or_other and not_a_company are the other decisions; unclassified means no decision yet. | [optional]
**assay_provider_class_source** | Option<**String**> | Where the class decision came from; null when unclassified. | [optional]
**assay_provider_class_decided_at** | Option<**String**> | When the class was decided, in UTC; null when unclassified. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



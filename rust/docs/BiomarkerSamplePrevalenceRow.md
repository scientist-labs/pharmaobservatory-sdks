# BiomarkerSamplePrevalenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**route_key** | Option<**String**> | The route key of the gene the variants belong to, the entry key for the biomarkers object type. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol. | [optional]
**variant_biomarker_id** | Option<**String**> | The variant's own biomarker key. | [optional]
**variant_route_key** | Option<**String**> | The variant's own route key (the id its web page uses). | [optional]
**amino_acid_change** | Option<**String**> | The protein change, for example p.V600E. | [optional]
**gdc_ssm_id** | Option<**String**> | The identifier of the somatic mutation in the Genomic Data Commons. | [optional]
**gdc_project_id** | Option<**String**> | The Genomic Data Commons research project (cohort) the counts come from. | [optional]
**gdc_project_name** | Option<**String**> | The name of the research project. | [optional]
**condition_id** | Option<**String**> | The condition key (a MeSH, Medical Subject Headings, descriptor id) the project's disease maps to, where one is mapped. | [optional]
**condition_name** | Option<**String**> | The mapped condition's MeSH heading. | [optional]
**source_disease_types** | Option<**Vec<String>**> | The disease types the source states for the project. | [optional]
**source_primary_sites** | Option<**Vec<String>**> | The primary sites the source states for the project. | [optional]
**affected_case_count** | Option<**i64**> | Cases in the cohort with the exact variant. | [optional]
**ssm_profiled_case_count** | Option<**i64**> | Cases in the cohort profiled for somatic mutations, the denominator of the observed frequency. | [optional]
**all_case_count** | Option<**i64**> | All cases in the project. | [optional]
**observed_frequency** | Option<**f64**> | Affected cases over mutation-profiled cases in this cohort. | [optional]
**frequency_ci95_low** | Option<**f64**> | Lower bound of the 95 percent confidence interval of the observed frequency. | [optional]
**frequency_ci95_high** | Option<**f64**> | Upper bound of the 95 percent confidence interval of the observed frequency. | [optional]
**gdc_data_release** | Option<**String**> | The Genomic Data Commons data release the counts come from. | [optional]
**claim_type** | Option<**String**> | What the figure states; an observed frequency in a named cohort. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



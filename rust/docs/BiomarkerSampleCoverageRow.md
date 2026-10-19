# BiomarkerSampleCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**route_key** | Option<**String**> | The route key of the gene the variants belong to, the entry key for the biomarkers object type. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol. | [optional]
**variant_biomarker_id** | Option<**String**> | The variant's own biomarker key. | [optional]
**variant_route_key** | Option<**String**> | The variant's own route key (the id its web page uses). | [optional]
**amino_acid_change** | Option<**String**> | The protein change, for example p.V600E. | [optional]
**gdc_data_release** | Option<**String**> | The Genomic Data Commons data release the match was made against. | [optional]
**normalization_status** | Option<**String**> | How the variant matched: resolved_single_ssm, no_exact_ssm, multiple_exact_ssms or source_error. | [optional]
**match_count** | Option<**i32**> | Somatic mutations in the catalogue that encode the same protein change. | [optional]
**resolution_reason** | Option<**String**> | The reason for the status, in words. | [optional]
**measured_project_count** | Option<**i64**> | Research projects with an observed frequency for the variant in biomarker_sample_prevalence. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



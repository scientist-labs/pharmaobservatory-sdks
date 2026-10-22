# RareDiseaseGenesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**gene_assertion_id** | Option<**String**> | The identifier of the disease-gene assertion. | [optional]
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol as Orphadata gives it. | [optional]
**gene_name** | Option<**String**> | The gene's full name. | [optional]
**gene_type** | Option<**String**> | The gene type Orphadata records (for example gene_with_protein_product). | [optional]
**association_type** | Option<**String**> | The association Orphadata asserts between the gene and the disease, for example a disease-causing germline mutation or a modifying germline mutation. | [optional]
**association_status** | Option<**String**> | assessed when Orphanet has assessed the association, not_yet_assessed otherwise. | [optional]
**is_known_causal** | Option<**bool**> | True for an assessed disease-causing mutation assertion. | [optional]
**ensembl_id** | Option<**String**> | The Ensembl gene id Orphadata gives for the gene. | [optional]
**ncbi_gene_id** | Option<**String**> | The NCBI Gene id Orphadata gives for the gene. | [optional]
**target_id** | Option<**String**> | The Ensembl gene id of the matching row in the targets dataset; null when the gene has no row there. | [optional]
**gene_resolution_status** | Option<**String**> | How the gene's Ensembl and NCBI identifiers resolved against the gene spine. | [optional]
**gene_identifier_conflict** | Option<**bool**> | True when the identifiers on the assertion disagree with the gene spine. | [optional]
**source_release** | Option<**String**> | The Orphadata Science release the row was read from. | [optional]
**source_url** | Option<**String**> | The URL of the source file the row was read from. | [optional]
**license_name** | Option<**String**> | The licence of the Orphadata Science file. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



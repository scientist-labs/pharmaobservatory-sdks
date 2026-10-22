# PharmaObservatory::RareDiseaseGenesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **gene_assertion_id** | **String** | The identifier of the disease-gene assertion. | [optional] |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **gene_symbol** | **String** | The gene symbol as Orphadata gives it. | [optional] |
| **gene_name** | **String** | The gene&#39;s full name. | [optional] |
| **gene_type** | **String** | The gene type Orphadata records (for example gene_with_protein_product). | [optional] |
| **association_type** | **String** | The association Orphadata asserts between the gene and the disease, for example a disease-causing germline mutation or a modifying germline mutation. | [optional] |
| **association_status** | **String** | assessed when Orphanet has assessed the association, not_yet_assessed otherwise. | [optional] |
| **is_known_causal** | **Boolean** | True for an assessed disease-causing mutation assertion. | [optional] |
| **ensembl_id** | **String** | The Ensembl gene id Orphadata gives for the gene. | [optional] |
| **ncbi_gene_id** | **String** | The NCBI Gene id Orphadata gives for the gene. | [optional] |
| **target_id** | **String** | The Ensembl gene id of the matching row in the targets dataset; null when the gene has no row there. | [optional] |
| **gene_resolution_status** | **String** | How the gene&#39;s Ensembl and NCBI identifiers resolved against the gene spine. | [optional] |
| **gene_identifier_conflict** | **Boolean** | True when the identifiers on the assertion disagree with the gene spine. | [optional] |
| **source_release** | **String** | The Orphadata Science release the row was read from. | [optional] |
| **source_url** | **String** | The URL of the source file the row was read from. | [optional] |
| **license_name** | **String** | The licence of the Orphadata Science file. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseGenesRow.new(
  gene_assertion_id: null,
  orpha_code: null,
  gene_symbol: null,
  gene_name: null,
  gene_type: null,
  association_type: null,
  association_status: null,
  is_known_causal: null,
  ensembl_id: null,
  ncbi_gene_id: null,
  target_id: null,
  gene_resolution_status: null,
  gene_identifier_conflict: null,
  source_release: null,
  source_url: null,
  license_name: null,
  attribution: null
)
```


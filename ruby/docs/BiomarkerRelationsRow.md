# PharmaObservatory::BiomarkerRelationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The biomarker&#39;s own route key, the entry key for the biomarkers object type (the gene symbol for a gene, the variant key for a protein variant). | [optional] |
| **biomarker_id** | **String** | The biomarker&#39;s key. | [optional] |
| **biomarker_kind** | **String** | Whether the row describes a gene or a protein variant. | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **amino_acid_change** | **String** | The protein change, for a variant. | [optional] |
| **related_entity_kind** | **String** | The kind of the related entity (gene, drug or condition). | [optional] |
| **related_entity_id** | **String** | The related entity&#39;s key in its own vocabulary (a gene symbol, a drug id such as unii:..., or a condition key). | [optional] |
| **related_display_name** | **String** | The related entity&#39;s name. | [optional] |
| **relation_type** | **String** | The source&#39;s relation label, as the source states it. | [optional] |
| **pmid** | **String** | The PubMed id of the paper the relation was extracted from. | [optional] |
| **title** | **String** | The paper&#39;s title. | [optional] |
| **journal_nlm_id** | **String** | The journal&#39;s National Library of Medicine (NLM) id. | [optional] |
| **journal_title** | **String** | The journal&#39;s title. | [optional] |
| **pub_year** | **Integer** | The paper&#39;s publication year. | [optional] |
| **n_source_assertions** | **Integer** | Source assertions of this relation in the paper. | [optional] |
| **n_relation_papers** | **Integer** | Papers that carry this biomarker, related entity and relation label. | [optional] |
| **evidence_type** | **String** | The evidence type (a text-mined relation, never a clinical claim). | [optional] |
| **source** | **String** | The source of the evidence. | [optional] |
| **display_order** | **Integer** | The order the web page lists the relation in within its kind. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerRelationsRow.new(
  route_key: null,
  biomarker_id: null,
  biomarker_kind: null,
  gene_symbol: null,
  amino_acid_change: null,
  related_entity_kind: null,
  related_entity_id: null,
  related_display_name: null,
  relation_type: null,
  pmid: null,
  title: null,
  journal_nlm_id: null,
  journal_title: null,
  pub_year: null,
  n_source_assertions: null,
  n_relation_papers: null,
  evidence_type: null,
  source: null,
  display_order: null
)
```


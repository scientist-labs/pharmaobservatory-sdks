# PharmaObservatory::BiomarkerVariantMentionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The route key of the gene the variants belong to, the entry key for the biomarkers object type. | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **variant_biomarker_id** | **String** | The variant&#39;s own biomarker key. | [optional] |
| **variant_route_key** | **String** | The variant&#39;s own route key (the id its web page uses). | [optional] |
| **display_name** | **String** | The variant&#39;s display name. | [optional] |
| **amino_acid_change** | **String** | The protein change, for example V600E. | [optional] |
| **biomarker_kind** | **String** | The kind of biomarker the row describes (a protein variant). | [optional] |
| **evidence_type** | **String** | The evidence type (a literature mention, never an actionability claim). | [optional] |
| **source** | **String** | The source of the evidence. | [optional] |
| **n_papers** | **Integer** | Distinct papers that name the variant. | [optional] |
| **first_publication_year** | **Integer** | Earliest publication year. | [optional] |
| **last_publication_year** | **Integer** | Latest publication year. | [optional] |
| **sample_pmids** | **Array&lt;String&gt;** | Up to five sample PubMed ids, newest first. | [optional] |
| **sample_citations** | **Array&lt;String&gt;** | Up to five sample citations, newest first. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerVariantMentionsRow.new(
  route_key: null,
  gene_symbol: null,
  variant_biomarker_id: null,
  variant_route_key: null,
  display_name: null,
  amino_acid_change: null,
  biomarker_kind: null,
  evidence_type: null,
  source: null,
  n_papers: null,
  first_publication_year: null,
  last_publication_year: null,
  sample_pmids: null,
  sample_citations: null
)
```


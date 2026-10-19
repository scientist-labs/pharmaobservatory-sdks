# PharmaObservatory::BiomarkerLiteratureRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The route key of the gene the paper names, the entry key for the biomarkers object type. | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **pmid** | **String** | The PubMed id. | [optional] |
| **evidence_type** | **String** | The evidence type (a publication mention). | [optional] |
| **title** | **String** | The paper&#39;s title. | [optional] |
| **journal_nlm_id** | **String** | The journal&#39;s NLM id. | [optional] |
| **journal_title** | **String** | The journal&#39;s title. | [optional] |
| **pub_year** | **Integer** | The publication year. | [optional] |
| **citation_count** | **Integer** | Citations, where iCite carries them (null is not measured, never zero). | [optional] |
| **relative_citation_ratio** | **Float** | The iCite relative citation ratio, where measured. | [optional] |
| **display_order** | **Integer** | The page&#39;s fixed order: newest year first, then citations. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerLiteratureRow.new(
  route_key: null,
  gene_symbol: null,
  pmid: null,
  evidence_type: null,
  title: null,
  journal_nlm_id: null,
  journal_title: null,
  pub_year: null,
  citation_count: null,
  relative_citation_ratio: null,
  display_order: null
)
```


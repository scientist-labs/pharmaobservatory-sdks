# PharmaObservatory::BiomarkerEvidenceByYearRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The gene&#39;s route key, the entry key for the biomarkers object type; the gene symbol. | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **evidence_type** | **String** | The evidence arm. trial_mention counts trials that mention the gene (the year is the trial start year); publication_mention counts papers cited by registered trials that mention it (the year is the publication year). | [optional] |
| **evidence_year** | **Integer** | The year of the mentions. Null for the undated class. | [optional] |
| **year_class** | **String** | How the year is treated. plotted: inside the window the web chart draws. pre_window: before 1990. future_dated: after the current year. corpus_incomplete: after the latest year the arm&#39;s corpus covers. undated: the source has no year. | [optional] |
| **n_mentions** | **Integer** | Mentions of the gene in that arm, year and class. A document that mentions the gene more than once counts once per mention. | [optional] |
| **n_documents** | **Integer** | Distinct documents (trials or papers) that mention the gene in that arm, year and class. | [optional] |
| **year_denominator** | **Integer** | Distinct documents in that arm and year that mention any gene, the denominator of the gene&#39;s share of the year. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerEvidenceByYearRow.new(
  route_key: null,
  gene_symbol: null,
  evidence_type: null,
  evidence_year: null,
  year_class: null,
  n_mentions: null,
  n_documents: null,
  year_denominator: null
)
```


# PharmaObservatory::TargetLiteratureRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. | [optional] |
| **target_symbol** | **String** | The target&#39;s HGNC gene symbol. | [optional] |
| **pmid** | **String** | The PubMed id. | [optional] |
| **title** | **String** | The paper&#39;s title. | [optional] |
| **journal_title** | **String** | The journal&#39;s title. | [optional] |
| **pub_year** | **Integer** | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional] |
| **evidence_type** | **String** | The paper&#39;s study-design label, one normalized publication type by study-design priority. | [optional] |
| **is_trial_linked** | **Boolean** | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional] |
| **citation_count** | **Integer** | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional] |
| **relative_citation_ratio** | **Float** | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetLiteratureRow.new(
  target_id: null,
  target_symbol: null,
  pmid: null,
  title: null,
  journal_title: null,
  pub_year: null,
  evidence_type: null,
  is_trial_linked: null,
  citation_count: null,
  relative_citation_ratio: null
)
```


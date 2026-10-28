# PharmaObservatory::BiomarkerAssayEvidenceRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **biomarker_assay_evidence_id** | **String** | The extracted mention&#39;s id. | [optional] |
| **route_key** | **String** | The route key of the gene the mention resolved to, the entry key for the biomarkers object type. Null when the gene did not resolve. | [optional] |
| **biomarker_id** | **String** | The resolved gene&#39;s biomarker key. Null when the gene did not resolve. | [optional] |
| **document_id** | **String** | The id of the source document the facts were extracted from. | [optional] |
| **pmid** | **String** | The PubMed id of the paper. | [optional] |
| **biomarker_ordinal** | **Integer** | The position of the mention in the paper&#39;s extracted biomarker list, from 1. | [optional] |
| **biomarker_name_as_extracted** | **String** | The biomarker name as extracted. | [optional] |
| **gene_symbol_as_extracted** | **String** | The gene symbol as extracted, before resolution. Null when the mention names no gene. | [optional] |
| **biomarker_synonyms** | **Array&lt;String&gt;** | The synonyms the paper gives, a document-level list repeated on each of its rows. | [optional] |
| **gene_resolution_status** | **String** | How the mention resolved to a gene. resolved_direct: exact symbol match. resolved_alias: an unambiguous dictionary alias. unresolved: a symbol was extracted that matches nothing or an ambiguous alias; not guessed. not_applicable: no gene symbol was extracted. | [optional] |
| **gene_resolution_method** | **String** | exact_gene_symbol or dictionary_alias; null when the gene did not resolve. | [optional] |
| **assay_name** | **String** | The assay, trade name or platform the paper names. A document-level value. | [optional] |
| **specimen_type** | **String** | The specimen type (blood, plasma, tissue). A document-level value. | [optional] |
| **comparator_method** | **String** | The reference method the performance was measured against, when reported. | [optional] |
| **reported_performance** | **String** | The performance as the paper states it (sensitivity, specificity, area under the curve), as text and not computed. | [optional] |
| **validation_type_mentioned** | **String** | What the paper claims about the validation type: a fact about the text, not a tier. | [optional] |
| **assay_provider_company** | **Array&lt;String&gt;** | The first extraction&#39;s flat list of organization names, as extracted, with no role. Not resolved to a company, and it names any company a paper mentions. Who makes or runs an assay for a biomarker, with the role, is the biomarker_assay_providers dataset. | [optional] |
| **funding_organizations** | **Array&lt;String&gt;** | Organization names as extracted that funded the study. Never merged with the assay provider. | [optional] |
| **evidence_type** | **String** | Always literature_assay_validation, which tells these rows from the other evidence types. | [optional] |
| **evidence_quote** | **Array&lt;String&gt;** | The sentences of the paper the facts were read from. Null when the extraction gave none. | [optional] |
| **confidence** | **String** | The extraction model&#39;s self-reported confidence. Not a reliability signal: it does not track whether the extraction is correct; never use it as a quality gate. | [optional] |
| **title** | **String** | The paper&#39;s title. | [optional] |
| **published_at** | **Time** | The paper&#39;s publication timestamp, in UTC. Null for about one paper in seven, whose source gives a bare year. | [optional] |
| **extraction_model** | **String** | The model that extracted the facts. | [optional] |
| **extracted_at** | **Time** | When the facts were extracted, in UTC. | [optional] |
| **source_url** | **String** | The address of the paper the facts were extracted from. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerAssayEvidenceRow.new(
  biomarker_assay_evidence_id: null,
  route_key: null,
  biomarker_id: null,
  document_id: null,
  pmid: null,
  biomarker_ordinal: null,
  biomarker_name_as_extracted: null,
  gene_symbol_as_extracted: null,
  biomarker_synonyms: null,
  gene_resolution_status: null,
  gene_resolution_method: null,
  assay_name: null,
  specimen_type: null,
  comparator_method: null,
  reported_performance: null,
  validation_type_mentioned: null,
  assay_provider_company: null,
  funding_organizations: null,
  evidence_type: null,
  evidence_quote: null,
  confidence: null,
  title: null,
  published_at: null,
  extraction_model: null,
  extracted_at: null,
  source_url: null
)
```


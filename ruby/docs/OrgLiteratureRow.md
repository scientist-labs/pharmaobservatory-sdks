# PharmaObservatory::OrgLiteratureRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **org_id** | **String** | Legal-entity level: the resolved parent organization record the byline names (edgar:..., ror:... or sponsor:...), the entry key for the legal_entities object type. | [optional] |
| **pmid** | **String** | The PubMed id. | [optional] |
| **title** | **String** | The paper&#39;s title. | [optional] |
| **journal_title** | **String** | The journal&#39;s title. | [optional] |
| **pub_year** | **Integer** | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional] |
| **evidence_type** | **String** | The paper&#39;s study-design label, one normalized publication type by study-design priority. | [optional] |
| **is_trial_linked** | **Boolean** | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional] |
| **org_arm** | **String** | company when the organization is company-typed (matched on a legal-form token in the affiliation); institution for every other organization. | [optional] |
| **match_tier** | **String** | The strongest tier that attributed a byline to the organization: legal_form or both (company arm), high or medium (institution arm). | [optional] |
| **n_bylines** | **Integer** | The number of distinct author positions on the paper whose affiliation named the organization. | [optional] |
| **citation_count** | **Integer** | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional] |
| **relative_citation_ratio** | **Float** | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgLiteratureRow.new(
  org_id: null,
  pmid: null,
  title: null,
  journal_title: null,
  pub_year: null,
  evidence_type: null,
  is_trial_linked: null,
  org_arm: null,
  match_tier: null,
  n_bylines: null,
  citation_count: null,
  relative_citation_ratio: null
)
```


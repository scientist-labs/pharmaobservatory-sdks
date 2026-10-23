# PharmaObservatory::OwnershipSourceCoverageRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **source_class** | **String** | The ownership source class (for example sec_def_14a for proxy statements, sec_form_adv for Form ADV). | [optional] |
| **subject_source_ref** | **String** | The subject as the source names it: the source&#39;s own reference for the owned company. | [optional] |
| **subject_name** | **String** | The subject&#39;s name as filed. | [optional] |
| **subject_org_id** | **String** | Legal-entity level: the organization record of the subject. | [optional] |
| **subject_company_id** | **String** | Company level: the company the subject belongs to (null when unresolved). | [optional] |
| **subject_company_name** | **String** | Display name of the subject&#39;s company. | [optional] |
| **subject_group_id** | **String** | Group level: the subject&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **subject_group_name** | **String** | Display name of the subject&#39;s group. | [optional] |
| **source_system** | **String** | The system the source class is read from (for example sec_edgar). | [optional] |
| **source_load_state** | **String** | The whole-source state: loaded, served_by_dedicated_mart (13F is served by the institutional holdings datasets), no_documents_ingested or source_not_ingested. | [optional] |
| **source_documents_ingested** | **Integer** | Documents the whole source has ingested. | [optional] |
| **coverage_state** | **String** | What the source read about this subject: reported, no_reportable_owner, extraction_abstained or read_no_owner. | [optional] |
| **documents_ingested** | **Integer** | Documents the source ingested about this subject. | [optional] |
| **observations** | **Integer** | Ownership observations the source recorded about this subject. | [optional] |
| **current_edges** | **Integer** | Ownership relationships that currently stand. | [optional] |
| **owners_distinct** | **Integer** | Distinct owners the source named for this subject. | [optional] |
| **owners_unresolved** | **Integer** | Named owners that did not resolve to an organization or investor. | [optional] |
| **no_reportable_owner_statements** | **Integer** | Filings that positively stated the subject has no reportable owner. | [optional] |
| **extraction_abstained** | **Integer** | Documents on which the extractor abstained. | [optional] |
| **last_filing_date** | **Date** | The latest filing date the source holds for this subject. | [optional] |
| **max_as_of_date** | **Date** | The latest as-of date among the subject&#39;s observations. | [optional] |
| **source_license** | **String** | The licence of the source. | [optional] |
| **attribution** | **String** | The credit line for the source. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OwnershipSourceCoverageRow.new(
  source_class: null,
  subject_source_ref: null,
  subject_name: null,
  subject_org_id: null,
  subject_company_id: null,
  subject_company_name: null,
  subject_group_id: null,
  subject_group_name: null,
  source_system: null,
  source_load_state: null,
  source_documents_ingested: null,
  coverage_state: null,
  documents_ingested: null,
  observations: null,
  current_edges: null,
  owners_distinct: null,
  owners_unresolved: null,
  no_reportable_owner_statements: null,
  extraction_abstained: null,
  last_filing_date: null,
  max_as_of_date: null,
  source_license: null,
  attribution: null
)
```


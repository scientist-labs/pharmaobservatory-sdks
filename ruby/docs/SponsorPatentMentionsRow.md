# PharmaObservatory::SponsorPatentMentionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sponsor_org_id** | **String** | Legal-entity level: the lead sponsor organization record. | [optional] |
| **sponsor_name** | **String** | Display name of the sponsor organization record. | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the sponsor belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **drug_id** | **String** | The drug key. | [optional] |
| **drug_name** | **String** | The drug&#39;s name. | [optional] |
| **patent_number** | **String** | The patent&#39;s publication number. | [optional] |
| **country** | **String** | The patent office&#39;s country code. | [optional] |
| **publication_date** | **Date** | The patent&#39;s publication date. | [optional] |
| **title** | **String** | The patent title. | [optional] |
| **assignee** | **String** | The patent&#39;s applicant as filed, free text. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SponsorPatentMentionsRow.new(
  sponsor_org_id: null,
  sponsor_name: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  drug_id: null,
  drug_name: null,
  patent_number: null,
  country: null,
  publication_date: null,
  title: null,
  assignee: null
)
```


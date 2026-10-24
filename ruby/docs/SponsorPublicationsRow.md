# PharmaObservatory::SponsorPublicationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **pmid** | **String** | The PubMed id. | [optional] |
| **citation** | **String** | The citation text of the paper as the trial registry cites it. | [optional] |
| **title** | **String** | The article title. | [optional] |
| **pub_year** | **Integer** | The publication year. | [optional] |
| **journal_nlm_id** | **String** | The journal&#39;s NLM (National Library of Medicine) catalog id; null when the paper&#39;s journal is unknown. | [optional] |
| **journal_title** | **String** | The journal title; null when the journal is not in the journal directory. | [optional] |
| **publication_trials** | **Integer** | How many trials of any sponsor cite the paper. | [optional] |
| **n_trials** | **Integer** | Distinct trials of this company that cite the paper. | [optional] |
| **n_lead_trials** | **Integer** | Of those, the trials the company leads. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SponsorPublicationsRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  pmid: null,
  citation: null,
  title: null,
  pub_year: null,
  journal_nlm_id: null,
  journal_title: null,
  publication_trials: null,
  n_trials: null,
  n_lead_trials: null
)
```


# PharmaObservatory::InvestigatorPublicationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **pmid** | **String** | The PubMed id of the publication. | [optional] |
| **link_basis** | **String** | authored when the person is linked to the paper as an author; trial_reference when a trial the person is on registers the paper (not authorship). | [optional] |
| **author_position** | **String** | first, middle or last for an authored row; null for a trial_reference row. | [optional] |
| **author_match_method** | **String** | How the author was matched to the person (orcid, name_on_linked_trial and so on); null for a trial_reference row. | [optional] |
| **citation** | **String** | The citation as a registered trial cites the paper (null when no trial cites it). | [optional] |
| **title** | **String** | The paper&#39;s title (null when no registered trial cites it). | [optional] |
| **pub_year** | **Integer** | The publication year (null when no registered trial cites the paper). | [optional] |
| **n_trials_citing** | **Integer** | Registered trials that cite the paper (null when none does). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestigatorPublicationsRow.new(
  investigator_person_id: null,
  pmid: null,
  link_basis: null,
  author_position: null,
  author_match_method: null,
  citation: null,
  title: null,
  pub_year: null,
  n_trials_citing: null
)
```


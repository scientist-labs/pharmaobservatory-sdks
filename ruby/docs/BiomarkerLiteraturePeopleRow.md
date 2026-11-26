# PharmaObservatory::BiomarkerLiteraturePeopleRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The route key of the gene, the entry key for the biomarkers object type. | [optional] |
| **investigator_person_id** | **String** | The author&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **person_name** | **String** | The person&#39;s name, resolved across the person&#39;s merged name fragments. | [optional] |
| **person_affiliation** | **String** | The person&#39;s affiliation as the person record holds it (null when not recorded). | [optional] |
| **orcid** | **String** | The person&#39;s ORCID iD (Open Researcher and Contributor ID) when the person record carries one (null when none). | [optional] |
| **n_papers** | **Integer** | Distinct papers on which a ledger decision links a byline to the person. | [optional] |
| **n_trial_linked_papers** | **Integer** | The part of n_papers a registered clinical trial cites. | [optional] |
| **latest_pub_year** | **Integer** | The newest publication year among those papers. | [optional] |
| **n_first_author** | **Integer** | The part of n_papers on which the person&#39;s linked byline is the first author position. | [optional] |
| **n_last_author** | **Integer** | The part of n_papers on which the person&#39;s linked byline is the last author position. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerLiteraturePeopleRow.new(
  route_key: null,
  investigator_person_id: null,
  person_name: null,
  person_affiliation: null,
  orcid: null,
  n_papers: null,
  n_trial_linked_papers: null,
  latest_pub_year: null,
  n_first_author: null,
  n_last_author: null
)
```


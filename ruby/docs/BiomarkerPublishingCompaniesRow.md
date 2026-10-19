# PharmaObservatory::BiomarkerPublishingCompaniesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The biomarker route key of the gene the papers name (the gene symbol); the entry key for the biomarkers object type. | [optional] |
| **gene_symbol** | **String** | The gene symbol. | [optional] |
| **publisher_company_id** | **String** | Company level: the publisher&#39;s company. | [optional] |
| **publisher_company_name** | **String** | Display name of the publisher&#39;s company. | [optional] |
| **publisher_group_id** | **String** | Group level: the publisher&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **publisher_group_name** | **String** | Display name of the publisher&#39;s group. | [optional] |
| **n_publications** | **Integer** | Papers that name the gene from the company&#39;s organizations (a consortium paper is excluded). | [optional] |
| **n_member_orgs** | **Integer** | The company&#39;s organizations the papers reach. | [optional] |
| **best_match_tier** | **String** | Strongest affiliation match tier. | [optional] |
| **first_pub_year** | **Integer** | Earliest publication year. | [optional] |
| **last_pub_year** | **Integer** | Latest publication year. | [optional] |
| **supporting_pmids** | **Array&lt;String&gt;** | A sample of the supporting PubMed ids. | [optional] |
| **supporting_pmids_truncated** | **Boolean** | True when the sample is shorter than the paper count. | [optional] |
| **n_papers_excluded_as_consortium** | **Integer** | Papers left out because they name too many organizations. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerPublishingCompaniesRow.new(
  route_key: null,
  gene_symbol: null,
  publisher_company_id: null,
  publisher_company_name: null,
  publisher_group_id: null,
  publisher_group_name: null,
  n_publications: null,
  n_member_orgs: null,
  best_match_tier: null,
  first_pub_year: null,
  last_pub_year: null,
  supporting_pmids: null,
  supporting_pmids_truncated: null,
  n_papers_excluded_as_consortium: null
)
```


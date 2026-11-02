# PharmaObservatory::CompanyCosponsorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **row_level** | **String** | root: the co-sponsor&#39;s ultimate parent company, with the distinct trials any company of that root shares with this company. member: a company under that root that itself co-sponsored, with its own count. Filter on one level; the two do not add. | [optional] |
| **cosponsor_company_id** | **String** | Company level: the co-sponsoring company. At row_level root this is the root company; at member it is the member company itself. | [optional] |
| **cosponsor_company_name** | **String** | Display name of the co-sponsoring company. | [optional] |
| **cosponsor_group_id** | **String** | Group level: the co-sponsor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **cosponsor_group_name** | **String** | Display name of the co-sponsor&#39;s group. | [optional] |
| **cosponsor_root_id** | **String** | The ultimate parent company of the co-sponsor, the ladder company&#39;s &#x60;coalesce(ultimate_parent_company_id, company_id)&#x60;. Equal to cosponsor_company_id at row_level root, and the root a member row sits under. | [optional] |
| **cosponsor_rank** | **Integer** | The rank of the ROOT among this company&#39;s co-sponsor roots by distinct shared trials (1 is the most; ties broken by root id). The same on a root row and on each of its member rows. | [optional] |
| **member_rank** | **Integer** | 0 on a root row; 1 and up on the member rows of that root, by the member&#39;s own shared trials. Order by cosponsor_rank then member_rank to list each root followed by its members. | [optional] |
| **n_members** | **Integer** | Member companies of this root that co-sponsored with this company, before the member cut. The same on a root row and its member rows. | [optional] |
| **n_candidates** | **Integer** | Distinct co-sponsor roots for this company before the top-25 cut; the same on every row of a company. | [optional] |
| **n_shared_trials** | **Integer** | Distinct trials that name both. At row_level root, the distinct trials of any company under the root (never a sum over members); at member, that company&#39;s own trials. | [optional] |
| **n_trials_company_lead** | **Integer** | Of the shared trials, those where an organization record of this company holds the ClinicalTrials.gov lead sponsor role. At most n_shared_trials; it may overlap n_trials_cosponsor_lead. | [optional] |
| **n_trials_cosponsor_lead** | **Integer** | Of the shared trials, those where the co-sponsor&#39;s organization record holds the lead sponsor role. At most n_shared_trials. | [optional] |
| **first_shared_start_year** | **Integer** | The earliest registered start year over the shared trials; null when none has a plausible year. | [optional] |
| **last_shared_start_year** | **Integer** | The latest registered start year over the shared trials; null when none has a plausible year. | [optional] |
| **n_shared_conditions** | **Integer** | Distinct specific MeSH disease areas over the shared trials. | [optional] |
| **top_shared_conditions** | **Array&lt;String&gt;** | Up to 5 MeSH disease names by shared-trial count then name. An empty list when none. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyCosponsorsRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  row_level: null,
  cosponsor_company_id: null,
  cosponsor_company_name: null,
  cosponsor_group_id: null,
  cosponsor_group_name: null,
  cosponsor_root_id: null,
  cosponsor_rank: null,
  member_rank: null,
  n_members: null,
  n_candidates: null,
  n_shared_trials: null,
  n_trials_company_lead: null,
  n_trials_cosponsor_lead: null,
  first_shared_start_year: null,
  last_shared_start_year: null,
  n_shared_conditions: null,
  top_shared_conditions: null
)
```


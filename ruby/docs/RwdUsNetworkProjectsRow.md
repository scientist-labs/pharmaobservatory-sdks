# PharmaObservatory::RwdUsNetworkProjectsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **project_id** | **String** | The N3C project id (a Data Use Request uid), the grain. | [optional] |
| **network_key** | **String** | The network key, the entry key of the rwd_us_networks object type. | [optional] |
| **title** | **String** | The project title as N3C publishes it. | [optional] |
| **institution_name** | **String** | The accessing institution as published. | [optional] |
| **institution_name_key** | **String** | The normalized name key; joins to rwd_us_network_memberships.institution_name_key. | [optional] |
| **pi_name** | **String** | Personal data: the principal investigator&#39;s name as N3C&#39;s public dashboard publishes it. Landed whole, tagged, never a gate (#5143). | [optional] |
| **workspace_status** | **String** | The project&#39;s workspace status as N3C publishes it. | [optional] |
| **project_type** | **String** | The project type as N3C publishes it. | [optional] |
| **is_task_team** | **Boolean** | Whether N3C marks the project as a task team. | [optional] |
| **contains_personal_data** | **Boolean** | True wherever pi_name is populated. A tag, never a gate (#5143). | [optional] |
| **personal_data_fields** | **String** | The personal-data columns of the row, as a JSON list: [\&quot;pi_name\&quot;]. | [optional] |
| **as_of** | **Date** | The retrieval date of the roster snapshot. | [optional] |
| **source_url** | **String** | The API endpoint the roster came from. | [optional] |
| **retrieved_at** | **Time** | When the snapshot was retrieved, in UTC. | [optional] |
| **org_id** | **String** | The organization spine id of the accessing institution, or null. Carried, never minted. | [optional] |
| **org_name** | **String** | The spine organization&#39;s name. | [optional] |
| **institution_ror_id** | **String** | The bare ROR id N3C published for the institution where it is a network member; null otherwise. | [optional] |
| **org_match_rung** | **String** | How the institution reached the organization: ror_exact, xref, match_keys, ambiguous or unresolved. | [optional] |
| **org_match_tier** | **String** | The tier of the organization match. | [optional] |
| **org_match_confidence** | **Float** | The organization match confidence. | [optional] |
| **institution_is_network_member** | **Boolean** | Whether the accessing institution itself holds a published agreement. | [optional] |
| **pi_person_id** | **String** | A resolved person on the spine, by name plus organization; see pi_match_method. The entry key of the investigators object type; null when unresolved. | [optional] |
| **pi_match_method** | **String** | How the PI was matched: name_org_spine (name and organization agreed with one spine person), name_org_reporter (a unique NIH RePORTER PI at the same organization), ambiguous (more than one candidate, never picked) or unresolved. | [optional] |
| **pi_corroborated** | **Boolean** | False on every row: no shared trial or paper corroborates the name-plus-organization match. | [optional] |
| **pi_reporter_pi_id** | **String** | NIH RePORTER&#39;s persistent PI_ID, an external key published as a labelled attribute, never as a resolved person; set where a unique RePORTER PI shares the name and organization. | [optional] |
| **candidate_person_count** | **Integer** | On an ambiguous row, the number of distinct spine persons the evidence could not choose between; otherwise 0. | [optional] |
| **also_matched_by_reporter** | **Boolean** | True on a name_org_spine row whose person the RePORTER arm also reached. | [optional] |
| **project_total** | **Integer** | Projects in the roster, the denominator of the resolution rate. | [optional] |
| **projects_with_institution_on_spine** | **Integer** | Projects whose institution resolved to a spine organization. | [optional] |
| **projects_with_resolved_pi** | **Integer** | Projects whose PI resolved to a spine person; over project_total it is the measured resolution rate. | [optional] |
| **projects_with_reporter_pi_id** | **Integer** | Projects carrying a RePORTER PI_ID. | [optional] |
| **attribution** | **String** | The credit line every surface carrying this data prints. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdUsNetworkProjectsRow.new(
  project_id: null,
  network_key: null,
  title: null,
  institution_name: null,
  institution_name_key: null,
  pi_name: null,
  workspace_status: null,
  project_type: null,
  is_task_team: null,
  contains_personal_data: null,
  personal_data_fields: null,
  as_of: null,
  source_url: null,
  retrieved_at: null,
  org_id: null,
  org_name: null,
  institution_ror_id: null,
  org_match_rung: null,
  org_match_tier: null,
  org_match_confidence: null,
  institution_is_network_member: null,
  pi_person_id: null,
  pi_match_method: null,
  pi_corroborated: null,
  pi_reporter_pi_id: null,
  candidate_person_count: null,
  also_matched_by_reporter: null,
  project_total: null,
  projects_with_institution_on_spine: null,
  projects_with_resolved_pi: null,
  projects_with_reporter_pi_id: null,
  attribution: null
)
```


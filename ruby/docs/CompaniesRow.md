# PharmaObservatory::CompaniesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | The company key (company:&lt;key&gt;). | [optional] |
| **company_name** | **String** | The company&#39;s display name: its own legal entity&#39;s name where one is known. | [optional] |
| **name_basis** | **String** | Which source supplied the name: gleif_legal_name, edgar_registrant, ror or sponsor_display_name. | [optional] |
| **country** | **String** | Country of the company&#39;s own legal entity, as the source states it. | [optional] |
| **org_type** | **String** | The organization type of the company&#39;s representative record. | [optional] |
| **org_kinds** | **Array&lt;String&gt;** | The roles the company&#39;s organizations play (Manufacturer, Sponsor and so on). | [optional] |
| **edgar_cik** | **String** | Central index key (CIK) in EDGAR, the US Securities and Exchange Commission (SEC) filing system, of the representative record, when it files. | [optional] |
| **edgar_ticker** | **String** | Exchange ticker of the representative record, when listed. | [optional] |
| **representative_org_id** | **String** | Legal-entity level: the organization record the company&#39;s page shows. | [optional] |
| **parent_company_id** | **String** | The direct owner company (part_of), null for a root. | [optional] |
| **parent_company_name** | **String** | Display name of the direct owner company. | [optional] |
| **part_of_since** | **Date** | Date the ownership link began, when every hop is dated. | [optional] |
| **part_of_since_precision** | **String** | Precision of part_of_since (day, month or year). | [optional] |
| **group_id** | **String** | Group level: the current ultimate parent company; a root company&#39;s group is itself. Current ownership only; dated links are in company_members. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **rollup_depth** | **Integer** | Hops from this company to its ultimate parent (0 for a root). | [optional] |
| **n_legal_entities** | **Integer** | Registered legal entities that are members of the company. | [optional] |
| **n_member_companies** | **Integer** | Companies folded into this one. | [optional] |
| **num_lead_trials** | **Integer** | Trials the company&#39;s organizations lead. | [optional] |
| **num_collaborator_trials** | **Integer** | Trials the company&#39;s organizations collaborate on. | [optional] |
| **num_trials** | **Integer** | Trials the company&#39;s organizations lead or collaborate on. | [optional] |
| **total_lead_enrollment** | **Integer** | Planned enrollment across led trials. | [optional] |
| **num_recent_lead_trials** | **Integer** | Led trials started in the recent window. | [optional] |
| **activity_score** | **Float** | Trial activity score (interventional lead trials and enrollment, weighted toward recent). | [optional] |
| **is_admitted** | **Boolean** | True when the company is admitted to the company roster (lead-trial activity, more than one legal entity or a recorded manufacturer role); false rows carry null firmographic, trial-count and activity columns. | [optional] |
| **has_public_page** | **Boolean** | True when the company has a public organization page. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompaniesRow.new(
  company_id: null,
  company_name: null,
  name_basis: null,
  country: null,
  org_type: null,
  org_kinds: null,
  edgar_cik: null,
  edgar_ticker: null,
  representative_org_id: null,
  parent_company_id: null,
  parent_company_name: null,
  part_of_since: null,
  part_of_since_precision: null,
  group_id: null,
  group_name: null,
  rollup_depth: null,
  n_legal_entities: null,
  n_member_companies: null,
  num_lead_trials: null,
  num_collaborator_trials: null,
  num_trials: null,
  total_lead_enrollment: null,
  num_recent_lead_trials: null,
  activity_score: null,
  is_admitted: null,
  has_public_page: null
)
```


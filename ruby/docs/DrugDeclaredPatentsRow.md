# PharmaObservatory::DrugDeclaredPatentsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **patent_id** | **String** | The patent id in the patent dimension. | [optional] |
| **assignee_sequence** | **Integer** | The assignee&#39;s order on the patent; -1 when there is none. | [optional] |
| **patent_number** | **String** | The patent number. | [optional] |
| **patent_title** | **String** | The patent title. | [optional] |
| **patent_date** | **Date** | The date the patent was granted. | [optional] |
| **declared_expiry** | **Date** | The expiry declared for the drug. | [optional] |
| **appl_nos** | **Array&lt;String&gt;** | The drug application numbers the patent is declared against. | [optional] |
| **link_basis** | **String** | The record that ties the patent to the drug (orange_book or uspto_pte). | [optional] |
| **link_bases** | **Array&lt;String&gt;** | Every record that ties the patent to the drug. | [optional] |
| **has_ped_extension** | **Boolean** | True when the Orange Book lists a pediatric exclusivity extension. | [optional] |
| **assignee_organization** | **String** | The assignee as filed. | [optional] |
| **assignee_org_id** | **String** | Legal-entity level: the assignee&#39;s organization record (null when unresolved). | [optional] |
| **assignee_company_id** | **String** | Company level: the assignee&#39;s company (null when unresolved). | [optional] |
| **assignee_company_name** | **String** | Display name of the assignee&#39;s company. | [optional] |
| **assignee_group_id** | **String** | Group level: the assignee&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **assignee_group_name** | **String** | Display name of the assignee&#39;s group. | [optional] |
| **org_in_spine** | **Boolean** | True when the assignee is an organization record in the spine. | [optional] |
| **resolution_method** | **String** | How the assignee was matched to an organization. | [optional] |
| **match_confidence** | **String** | The match confidence (high, med, low). | [optional] |
| **owner_basis** | **String** | Why the assignee is, or is not, on record. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugDeclaredPatentsRow.new(
  molecule_id: null,
  patent_id: null,
  assignee_sequence: null,
  patent_number: null,
  patent_title: null,
  patent_date: null,
  declared_expiry: null,
  appl_nos: null,
  link_basis: null,
  link_bases: null,
  has_ped_extension: null,
  assignee_organization: null,
  assignee_org_id: null,
  assignee_company_id: null,
  assignee_company_name: null,
  assignee_group_id: null,
  assignee_group_name: null,
  org_in_spine: null,
  resolution_method: null,
  match_confidence: null,
  owner_basis: null
)
```


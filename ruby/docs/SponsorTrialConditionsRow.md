# PharmaObservatory::SponsorTrialConditionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record of the sponsor as filed. | [optional] |
| **sponsor_name** | **String** | Display name of the sponsor organization record. | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the sponsor belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |
| **sponsor_role** | **String** | lead or collaborator. | [optional] |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **condition_id** | **String** | The condition key, a MeSH descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The condition&#39;s MeSH heading. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SponsorTrialConditionsRow.new(
  sponsor_org_id: null,
  sponsor_name: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null,
  sponsor_role: null,
  trial_id: null,
  condition_id: null,
  condition_name: null
)
```


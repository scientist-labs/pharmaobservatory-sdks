# PharmaObservatory::BiomarkerCdxAssertionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **source_assertion_id** | **String** | The assertion id. | [optional] |
| **route_key** | **String** | The biomarker&#39;s route key, the entry key for the biomarkers object type. | [optional] |
| **biomarker_id** | **String** | The biomarker key. | [optional] |
| **biomarker_token** | **String** | The biomarker as the FDA list names it. | [optional] |
| **diagnostic_name** | **String** | The companion diagnostic&#39;s name. | [optional] |
| **indication** | **String** | The indication. | [optional] |
| **sample_type** | **String** | The sample type. | [optional] |
| **submission_cell** | **String** | The submission number cell as listed. | [optional] |
| **latest_stated_decision_date** | **Date** | Latest decision date the list states. | [optional] |
| **source_url** | **String** | Link to the FDA list. | [optional] |
| **manufacturer_name** | **String** | The manufacturer as filed. | [optional] |
| **manufacturer_org_id** | **String** | Legal-entity level: the organization record the manufacturer resolved to (null when unresolved). | [optional] |
| **manufacturer_company_id** | **String** | Company level: the company the manufacturer belongs to (null when unresolved). | [optional] |
| **manufacturer_company_name** | **String** | Display name of the manufacturer&#39;s company. | [optional] |
| **manufacturer_group_id** | **String** | Group level: the manufacturer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **manufacturer_group_name** | **String** | Display name of the manufacturer&#39;s group. | [optional] |
| **manufacturer_match_tier** | **String** | How firmly the manufacturer name matched an organization. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerCdxAssertionsRow.new(
  source_assertion_id: null,
  route_key: null,
  biomarker_id: null,
  biomarker_token: null,
  diagnostic_name: null,
  indication: null,
  sample_type: null,
  submission_cell: null,
  latest_stated_decision_date: null,
  source_url: null,
  manufacturer_name: null,
  manufacturer_org_id: null,
  manufacturer_company_id: null,
  manufacturer_company_name: null,
  manufacturer_group_id: null,
  manufacturer_group_name: null,
  manufacturer_match_tier: null
)
```


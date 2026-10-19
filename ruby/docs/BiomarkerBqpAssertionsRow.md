# PharmaObservatory::BiomarkerBqpAssertionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **source_assertion_id** | **String** | The assertion id. | [optional] |
| **route_key** | **String** | The biomarker&#39;s route key, the entry key for the biomarkers object type. | [optional] |
| **biomarker_id** | **String** | The biomarker key. | [optional] |
| **ddt_number** | **String** | The FDA drug development tool number of the project. | [optional] |
| **ddt_project_id** | **String** | The FDA project id. | [optional] |
| **project_name** | **String** | The project&#39;s name. | [optional] |
| **bqp_signal** | **String** | The source status of the project. Only qualified means FDA-qualified; every other value is a project status. | [optional] |
| **is_fda_qualified** | **Boolean** | True only when the source status is exactly qualified. | [optional] |
| **submission_status** | **String** | The submission status as the source words it. | [optional] |
| **submission_stage** | **String** | The submission stage as the source words it. | [optional] |
| **submitter_name** | **String** | The submitter as the source names it. | [optional] |
| **submitter_org_id** | **String** | Legal-entity level: the organization record the submitter resolved to (null when unresolved). | [optional] |
| **submitter_company_id** | **String** | Company level: the company the submitter belongs to (null when unresolved). | [optional] |
| **submitter_company_name** | **String** | Display name of the submitter&#39;s company. | [optional] |
| **submitter_group_id** | **String** | Group level: the submitter&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **submitter_group_name** | **String** | Display name of the submitter&#39;s group. | [optional] |
| **submitter_match_tier** | **String** | How the submitter name matched an organization record. | [optional] |
| **submitter_role** | **String** | The submitter&#39;s role on the project. | [optional] |
| **biomarker_role** | **String** | The biomarker&#39;s role in the project (for example a diagnostic or a surrogate endpoint). | [optional] |
| **biomarker_role_method** | **String** | How the biomarker role was assigned. | [optional] |
| **is_surrogate_endpoint_claim** | **Boolean** | True when the project claims a surrogate endpoint. | [optional] |
| **source_url** | **String** | Link to the FDA project page. | [optional] |
| **snapshot_date** | **Date** | The date the source list was captured. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerBqpAssertionsRow.new(
  source_assertion_id: null,
  route_key: null,
  biomarker_id: null,
  ddt_number: null,
  ddt_project_id: null,
  project_name: null,
  bqp_signal: null,
  is_fda_qualified: null,
  submission_status: null,
  submission_stage: null,
  submitter_name: null,
  submitter_org_id: null,
  submitter_company_id: null,
  submitter_company_name: null,
  submitter_group_id: null,
  submitter_group_name: null,
  submitter_match_tier: null,
  submitter_role: null,
  biomarker_role: null,
  biomarker_role_method: null,
  is_surrogate_endpoint_claim: null,
  source_url: null,
  snapshot_date: null
)
```


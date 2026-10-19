# PharmaObservatory::SiteTrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_id** | **String** | The site key: one facility spelling from the registry. | [optional] |
| **trial_id** | **String** | The trial&#39;s registry id (NCT number). | [optional] |
| **trial_title** | **String** | The trial&#39;s brief title. | [optional] |
| **study_type** | **String** | INTERVENTIONAL or OBSERVATIONAL. | [optional] |
| **phase** | **String** | The trial phase label. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall status. | [optional] |
| **start_date** | **Date** | Trial start date. | [optional] |
| **start_year** | **Integer** | Year of the trial start. | [optional] |
| **enrollment** | **Integer** | Enrollment as registered for the whole trial. | [optional] |
| **lead_sponsor_name** | **String** | The lead sponsor&#39;s name as filed. | [optional] |
| **lead_sponsor_org_id** | **String** | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional] |
| **lead_sponsor_company_id** | **String** | Company level: the company the lead sponsor belongs to (null when unresolved). | [optional] |
| **lead_sponsor_company_name** | **String** | Display name of the lead sponsor&#39;s company. | [optional] |
| **lead_sponsor_group_id** | **String** | Group level: the lead sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **lead_sponsor_group_name** | **String** | Display name of the lead sponsor&#39;s group. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SiteTrialsRow.new(
  site_id: null,
  trial_id: null,
  trial_title: null,
  study_type: null,
  phase: null,
  overall_status: null,
  start_date: null,
  start_year: null,
  enrollment: null,
  lead_sponsor_name: null,
  lead_sponsor_org_id: null,
  lead_sponsor_company_id: null,
  lead_sponsor_company_name: null,
  lead_sponsor_group_id: null,
  lead_sponsor_group_name: null
)
```


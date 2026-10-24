# PharmaObservatory::InvestigatorTrialsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **roles** | **Array&lt;String&gt;** | The roles the person holds on this trial as registered. | [optional] |
| **is_principal_investigator** | **Boolean** | True when the person is named principal investigator on this trial. | [optional] |
| **trial_title** | **String** | The trial&#39;s brief title. | [optional] |
| **phase_label** | **String** | The trial phase as a display label. | [optional] |
| **overall_status** | **String** | The registry&#39;s overall recruitment status. | [optional] |
| **study_type** | **String** | The registry&#39;s study type. | [optional] |
| **start_date** | **Date** | Trial start date as registered. | [optional] |
| **start_year** | **Integer** | Calendar year of the start date. | [optional] |
| **completion_date** | **Date** | Trial completion date as registered. | [optional] |
| **enrollment** | **Integer** | Enrollment as registered. | [optional] |
| **n_sites** | **Integer** | The trial&#39;s distinct geocoded sites, the count the people page shows (0 when none is geocoded). | [optional] |
| **single_site_id** | **String** | The site key when the trial lists exactly one site (null otherwise); joins the sites dataset. | [optional] |
| **single_site_name** | **String** | The single site&#39;s facility name as the registry lists it. | [optional] |
| **single_site_city** | **String** | The single site&#39;s city. | [optional] |
| **single_site_state** | **String** | The single site&#39;s state or province. | [optional] |
| **single_site_country** | **String** | The single site&#39;s country. | [optional] |
| **sponsor_name** | **String** | The trial&#39;s lead sponsor as filed. | [optional] |
| **sponsor_org_id** | **String** | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional] |
| **sponsor_company_id** | **String** | Company level: the company the lead sponsor belongs to (null when unresolved). | [optional] |
| **sponsor_company_name** | **String** | Display name of the sponsor&#39;s company. | [optional] |
| **sponsor_group_id** | **String** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **sponsor_group_name** | **String** | Display name of the sponsor&#39;s group. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestigatorTrialsRow.new(
  investigator_person_id: null,
  trial_id: null,
  roles: null,
  is_principal_investigator: null,
  trial_title: null,
  phase_label: null,
  overall_status: null,
  study_type: null,
  start_date: null,
  start_year: null,
  completion_date: null,
  enrollment: null,
  n_sites: null,
  single_site_id: null,
  single_site_name: null,
  single_site_city: null,
  single_site_state: null,
  single_site_country: null,
  sponsor_name: null,
  sponsor_org_id: null,
  sponsor_company_id: null,
  sponsor_company_name: null,
  sponsor_group_id: null,
  sponsor_group_name: null
)
```


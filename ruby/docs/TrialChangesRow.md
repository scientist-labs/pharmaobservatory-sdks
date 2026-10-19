# PharmaObservatory::TrialChangesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **observed_at** | **Time** | When the platform observed the change. | [optional] |
| **event_type** | **String** | The kind of event (first_seen, status_change and so on). | [optional] |
| **field** | **String** | The record field that changed (null for a first_seen event). | [optional] |
| **old_value** | **String** | The value before the change. | [optional] |
| **new_value** | **String** | The value after the change. | [optional] |
| **trial_title** | **String** | The trial&#39;s brief title. | [optional] |
| **condition_ids** | **Array&lt;String&gt;** | MeSH (Medical Subject Headings) condition ids the trial is linked to, direct and through the MeSH tree. Filter on one id to scope changes to a condition&#39;s trials. | [optional] |
| **drug_ids** | **Array&lt;String&gt;** | Drug ids the trial&#39;s interventions resolved to. Filter on one id to scope changes to a drug&#39;s trials. | [optional] |
| **sponsor_org_ids** | **Array&lt;String&gt;** | Organization record ids of the sponsors on the trial, lead and collaborator. Filter on one id to scope changes to a sponsor&#39;s trials. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialChangesRow.new(
  trial_id: null,
  observed_at: null,
  event_type: null,
  field: null,
  old_value: null,
  new_value: null,
  trial_title: null,
  condition_ids: null,
  drug_ids: null,
  sponsor_org_ids: null
)
```


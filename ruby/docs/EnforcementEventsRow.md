# PharmaObservatory::EnforcementEventsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **event_id** | **String** | The event id: a hash of the lane and the lane&#39;s own key, unique across lanes. | [optional] |
| **source** | **String** | The lane the event came from: inspection_483 (an FDA inspection), import_refusal, warning_letter, untitled_letter, debarment, recall or aphis_inspection (a USDA APHIS Animal Welfare Act inspection report). | [optional] |
| **event_date** | **Date** | The event date: the inspection date, refusal date, letter issue date (posted date when there is none), debarment effective date or recall initiation date. | [optional] |
| **domain** | **String** | The regulatory domain of the event, derived from the lane&#39;s own classifier: drug, biologic, device, animal_research (the APHIS lane), veterinary, food, cosmetic, tobacco or other. | [optional] |
| **is_biomedical** | **Boolean** | True when domain is drug, biologic or device. The APHIS lane is false here on purpose; the page&#39;s default view adds domain animal_research to this set. | [optional] |
| **headline** | **String** | A one-line description: the inspection product type, the refused product, the letter subject, the debarment type or the recalled product. | [optional] |
| **detail** | **String** | The lane&#39;s supporting text: the cited regulations of an inspection, the refusal charges, the issuing office of a warning letter, the debarment term and citation, or the recall reason. | [optional] |
| **classification** | **String** | The inspection outcome (NAI no action indicated, VAI voluntary action indicated, OAI official action indicated); null for every other lane. | [optional] |
| **recall_class** | **String** | The recall severity class (Class I, II or III); null for every other lane. | [optional] |
| **had_483** | **Boolean** | True when the inspection ended with a Form 483 of observations; null for every other lane. | [optional] |
| **source_url** | **String** | The regulator&#39;s page for the event, where the lane has one. | [optional] |
| **source_ref** | **String** | The lane&#39;s own reference, where it has one: the letter id, debarment id or recall number. | [optional] |
| **device_class** | **String** | The FDA device class of an import-refusal product, from the device classification dimension; null for every other lane. | [optional] |
| **device_specialty** | **String** | The medical specialty of an import-refusal device product; null for every other lane. | [optional] |
| **is_low_signal_device** | **Boolean** | True for a device event the biomedical view hides by default (a low-signal commodity class). | [optional] |
| **awa_outcome** | **String** | The APHIS report&#39;s verdict, the strongest citation kind on it: direct, critical, non_critical, no_citations or no_citations_no_animals; null for every FDA lane. | [optional] |
| **awa_inspection_type** | **String** | The inspection type as printed on the APHIS report (for example ROUTINE INSPECTION); null for every FDA lane. | [optional] |
| **awa_n_citations** | **Integer** | Citations on the APHIS report; null for every FDA lane. Read with awa_species_scope. | [optional] |
| **awa_n_direct** | **Integer** | Of the report&#39;s citations, those cited as direct; null for every FDA lane. | [optional] |
| **awa_n_critical** | **Integer** | Of the report&#39;s citations, those cited as critical; null for every FDA lane. | [optional] |
| **awa_n_non_critical** | **Integer** | Of the report&#39;s citations, those cited as non-critical; null for every FDA lane. | [optional] |
| **awa_n_repeat_citations** | **Integer** | Of the report&#39;s citations, those it marks as repeats; null for every FDA lane. | [optional] |
| **awa_species_summary** | **String** | The species the APHIS report printed with their animal counts, as \&quot;Common name count; ...\&quot; in the report&#39;s order; null on an FDA lane or a report that printed none. | [optional] |
| **awa_species_scope** | **String** | The label AWA-covered species on every APHIS row; null for every FDA lane. The awa_* counts describe the species the Animal Welfare Act covers; read them with this label. | [optional] |
| **firm_name** | **String** | The firm&#39;s name as filed with the regulator (the party named on the record, an individual for an individual debarment). | [optional] |
| **firm_org_id** | **String** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] |
| **firm_company_id** | **String** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] |
| **firm_company_name** | **String** | Display name of the firm&#39;s company. | [optional] |
| **firm_group_id** | **String** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **firm_group_name** | **String** | Display name of the firm&#39;s group. | [optional] |
| **match_tier** | **String** | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous or unresolved. | [optional] |
| **org_resolution_method** | **String** | The route that attached the firm: fei_bridge (the FDA establishment identifier bridge), fei_firm_mint or fei_firm_fold (the operating firm&#39;s registry record), aphis_bridge, aphis_firm_mint or aphis_firm_fold (the APHIS registrant certificate bridge and its operating-firm record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional] |
| **is_org_resolved** | **Boolean** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::EnforcementEventsRow.new(
  event_id: null,
  source: null,
  event_date: null,
  domain: null,
  is_biomedical: null,
  headline: null,
  detail: null,
  classification: null,
  recall_class: null,
  had_483: null,
  source_url: null,
  source_ref: null,
  device_class: null,
  device_specialty: null,
  is_low_signal_device: null,
  awa_outcome: null,
  awa_inspection_type: null,
  awa_n_citations: null,
  awa_n_direct: null,
  awa_n_critical: null,
  awa_n_non_critical: null,
  awa_n_repeat_citations: null,
  awa_species_summary: null,
  awa_species_scope: null,
  firm_name: null,
  firm_org_id: null,
  firm_company_id: null,
  firm_company_name: null,
  firm_group_id: null,
  firm_group_name: null,
  match_tier: null,
  org_resolution_method: null,
  is_org_resolved: null
)
```


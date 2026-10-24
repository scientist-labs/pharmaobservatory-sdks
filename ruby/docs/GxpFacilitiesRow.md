# PharmaObservatory::GxpFacilitiesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **facility_key** | **String** | The canonical facility id. | [optional] |
| **registry** | **String** | The registry of the lead member record. Provenance and a facet, never a join key. | [optional] |
| **registry_native_site_id** | **String** | The lead registry&#39;s own site id (FEI, EudraGMDP site key, Health Canada location id). | [optional] |
| **registry_facility_keys** | **Array&lt;String&gt;** | The registry records the facility is made of, each as registry:native site id (for example fda_fei:3004123456), sorted. | [optional] |
| **facility_name** | **String** | The facility&#39;s name. | [optional] |
| **member_count** | **Integer** | How many registry records are clustered into the facility. | [optional] |
| **address_line** | **String** | The street address. | [optional] |
| **city** | **String** | The city. | [optional] |
| **state_region** | **String** | The state or region. | [optional] |
| **postal** | **String** | The postal code. | [optional] |
| **country** | **String** | The two-letter country code. | [optional] |
| **latitude** | **Float** | Latitude in degrees (WGS84), for the map tool. | [optional] |
| **longitude** | **Float** | Longitude in degrees (WGS84), for the map tool. | [optional] |
| **geocode_precision** | **String** | What a pin locates: for example a street address, a postal code or only a city. | [optional] |
| **coord_source** | **String** | Where the coordinates came from: geonames (GeoNames) or als (Amazon Location Service); it names the attribution to show. | [optional] |
| **firm_name** | **String** | The operating firm&#39;s name where the facility resolved to an organization; the facility name where it did not. | [optional] |
| **firm_org_id** | **String** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] |
| **firm_company_id** | **String** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] |
| **firm_company_name** | **String** | Display name of the firm&#39;s company. | [optional] |
| **firm_group_id** | **String** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **firm_group_name** | **String** | Display name of the firm&#39;s group. | [optional] |
| **is_org_resolved** | **Boolean** | True when the facility was attached to an operating organization. A false row is kept with null firm ids. | [optional] |
| **match_tier** | **String** | How the operating firm was attached: registry (the firm&#39;s own registry record), high, medium or low name-and-geography matches, or null. | [optional] |
| **domains** | **Array&lt;String&gt;** | The regulatory domains of the facility&#39;s signals (drug, biologic, device, veterinary, food, cosmetic, tobacco). | [optional] |
| **primary_domain** | **String** | The domain with the most signals. | [optional] |
| **is_drug_facility** | **Boolean** | True when the registers say the site makes medicines (an FDA drug registration, or an EU or Canadian medicines register entry). | [optional] |
| **is_device_facility** | **Boolean** | True when the registers say the site makes devices. | [optional] |
| **register_domain** | **String** | The display bucket over the register flags: drug, device, drug_and_device, animal_research (a USDA APHIS research facility that is neither a drug nor a device facility) or non_medicine. | [optional] |
| **device_classification_status** | **String** | The registered-product rollup for a device facility: material, unresolved, low_signal or not_applicable. | [optional] |
| **is_low_signal_device** | **Boolean** | True when the facility&#39;s device products are all low-signal commodity classes. | [optional] |
| **regulators** | **Array&lt;String&gt;** | The regulators holding a registry record for the facility (fda, ema, health_canada), sorted. | [optional] |
| **n_regulators** | **Integer** | How many regulators hold a registry record. | [optional] |
| **is_multi_regulator** | **Boolean** | True when more than one regulator holds a record: the corroboration flag. | [optional] |
| **org_regulators** | **Array&lt;String&gt;** | The regulators holding an organization-level record for the operating firm (this carries the UK MHRA leg). | [optional] |
| **severity_rank** | **Integer** | The in-force site-scope severity, 0 (no signal) to 4 (non-compliant). | [optional] |
| **severity_level** | **String** | The in-force site-scope severity: no_signal, clean, advisory, action or non_compliant. | [optional] |
| **severity_scope** | **String** | site when a site-scope verdict set the severity rank, none when nothing did. Never org. | [optional] |
| **severity_source** | **String** | The regulator behind the severity: fda, ema, health_canada or mhra. | [optional] |
| **severity_signal_type** | **String** | The kind of signal behind the severity (for example an inspection classification). | [optional] |
| **severity_raw_verdict** | **String** | The regulator&#39;s own verdict behind the severity (for example OAI). | [optional] |
| **severity_signal_date** | **Date** | The date of the signal behind the severity. | [optional] |
| **severity_source_url** | **String** | The source page of the signal that set the severity. | [optional] |
| **worst_severity_rank_ever** | **Integer** | The highest site-scope severity the facility has ever had, 0 to 4. | [optional] |
| **worst_severity_level_ever** | **String** | The worst severity ever: no_signal, clean, advisory, action or non_compliant. | [optional] |
| **worst_severity_raw_verdict** | **String** | The regulator&#39;s own verdict word behind the worst severity ever. | [optional] |
| **worst_severity_signal_date** | **Date** | The date of the signal behind the worst severity ever. | [optional] |
| **has_quality_signal** | **Boolean** | True when the facility has any quality signal. | [optional] |
| **has_site_scope_signal** | **Boolean** | True when a signal is about this site itself. | [optional] |
| **has_org_scope_signal** | **Boolean** | True when a signal is about the operating firm. | [optional] |
| **latest_signal_date** | **Date** | The date of the latest signal of any scope. | [optional] |
| **first_signal_date** | **Date** | The date of the earliest dated signal of any scope. | [optional] |
| **n_signals_total** | **Integer** | Signals of any scope, all time. | [optional] |
| **n_signals_365d** | **Integer** | Signals in the last 365 days of the build. | [optional] |
| **n_signals_1095d** | **Integer** | Signals in the last 1,095 days of the build. | [optional] |
| **n_signals_1825d** | **Integer** | Signals in the last 1,825 days of the build. | [optional] |
| **n_inspections** | **Integer** | Inspections of the site by any regulator. | [optional] |
| **latest_inspection_date** | **Date** | The date of the latest inspection by any regulator. | [optional] |
| **n_fda_inspections** | **Integer** | FDA inspections of the site. | [optional] |
| **latest_fda_inspection_date** | **Date** | The date of the latest FDA inspection. | [optional] |
| **latest_inspection_classification** | **String** | The latest FDA inspection outcome (NAI, VAI or OAI). | [optional] |
| **n_inspections_nai** | **Integer** | FDA inspections classified NAI. | [optional] |
| **n_inspections_vai** | **Integer** | FDA inspections classified VAI. | [optional] |
| **n_inspections_oai** | **Integer** | FDA inspections classified OAI. | [optional] |
| **n_483** | **Integer** | FDA inspections that ended with a Form 483. | [optional] |
| **n_hc_inspections** | **Integer** | Health Canada inspections of the site. | [optional] |
| **latest_hc_inspection_date** | **Date** | The date of the latest Health Canada inspection. | [optional] |
| **latest_hc_inspection_rating** | **String** | The latest Health Canada rating (for example compliant or non-compliant). | [optional] |
| **n_hc_inspections_non_compliant** | **Integer** | Health Canada inspections rated non-compliant. | [optional] |
| **n_import_refusals** | **Integer** | FDA import refusals at the site. | [optional] |
| **latest_import_refusal_date** | **Date** | The date of the latest import refusal. | [optional] |
| **is_research_facility** | **Boolean** | True when the facility has a USDA APHIS Animal Welfare Act registrant site as a member (the animal_research register); a facility can also be a drug or device facility. | [optional] |
| **n_aphis_inspections** | **Integer** | USDA APHIS inspection reports for the site, one per report and never per citation; zero where the facility has no APHIS member. Read with awa_species_scope. | [optional] |
| **latest_aphis_inspection_date** | **Date** | The date of the latest APHIS inspection report; null when there is none. | [optional] |
| **latest_aphis_inspection_outcome** | **String** | The latest APHIS report&#39;s verdict, the strongest citation kind on it: direct, critical, non_critical, no_citations or no_citations_no_animals; null when there is no APHIS report. | [optional] |
| **n_aphis_inspections_cited** | **Integer** | APHIS reports at the site with at least one citation (direct, critical or non-critical). | [optional] |
| **n_aphis_citations** | **Integer** | Citations across the site&#39;s APHIS reports. Read with awa_species_scope. | [optional] |
| **n_aphis_repeat_citations** | **Integer** | Of those, the citations the report marks as repeats. | [optional] |
| **awa_species_scope** | **String** | The label AWA-covered species on a facility with an APHIS member; null otherwise. The APHIS counts describe the species the Animal Welfare Act covers; read them with this label. | [optional] |
| **org_olaw_assurance_id** | **String** | Organization scope, repeated onto every facility of the organization: the NIH OLAW Public Health Service Animal Welfare Assurance id, or null when the organization has no resolved row. Null is never evidence of no assurance. | [optional] |
| **org_olaw_assurance_status** | **String** | OLAW&#39;s raw two-digit assurance status code as published, never decoded; null with org_olaw_assurance_id. | [optional] |
| **org_olaw_match_tier** | **String** | How well the OLAW institution name resolved to the organization: high, medium or low; the web badge shows high and medium. | [optional] |
| **org_olaw_source_asof** | **Time** | When OLAW&#39;s lookup tool says its list was last updated; the source does not state a timezone. | [optional] |
| **has_org_olaw_assurance** | **Boolean** | True when org_olaw_assurance_id is not null. Organization scope, repeated per facility, not additive. | [optional] |
| **n_gmp_certificates** | **Integer** | EU GMP certificates of the site. | [optional] |
| **latest_gmp_certificate_date** | **Date** | The date of the latest EU GMP certificate. | [optional] |
| **n_noncompliance_statements** | **Integer** | EU GMP non-compliance statements about the site. | [optional] |
| **latest_noncompliance_statement_date** | **Date** | The date of the latest EU non-compliance statement. | [optional] |
| **n_licences** | **Integer** | Health Canada establishment licences of the site. | [optional] |
| **licence_status** | **String** | The licence status (the strictest licence of the site). | [optional] |
| **org_n_licence_actions** | **Integer** | UK MHRA licence suspensions or revocations against the firm (repeated; does not add). | [optional] |
| **org_latest_licence_action_date** | **Date** | The date of the firm&#39;s latest licence action. | [optional] |
| **as_of** | **Date** | The build date the counts are as of. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::GxpFacilitiesRow.new(
  facility_key: null,
  registry: null,
  registry_native_site_id: null,
  registry_facility_keys: null,
  facility_name: null,
  member_count: null,
  address_line: null,
  city: null,
  state_region: null,
  postal: null,
  country: null,
  latitude: null,
  longitude: null,
  geocode_precision: null,
  coord_source: null,
  firm_name: null,
  firm_org_id: null,
  firm_company_id: null,
  firm_company_name: null,
  firm_group_id: null,
  firm_group_name: null,
  is_org_resolved: null,
  match_tier: null,
  domains: null,
  primary_domain: null,
  is_drug_facility: null,
  is_device_facility: null,
  register_domain: null,
  device_classification_status: null,
  is_low_signal_device: null,
  regulators: null,
  n_regulators: null,
  is_multi_regulator: null,
  org_regulators: null,
  severity_rank: null,
  severity_level: null,
  severity_scope: null,
  severity_source: null,
  severity_signal_type: null,
  severity_raw_verdict: null,
  severity_signal_date: null,
  severity_source_url: null,
  worst_severity_rank_ever: null,
  worst_severity_level_ever: null,
  worst_severity_raw_verdict: null,
  worst_severity_signal_date: null,
  has_quality_signal: null,
  has_site_scope_signal: null,
  has_org_scope_signal: null,
  latest_signal_date: null,
  first_signal_date: null,
  n_signals_total: null,
  n_signals_365d: null,
  n_signals_1095d: null,
  n_signals_1825d: null,
  n_inspections: null,
  latest_inspection_date: null,
  n_fda_inspections: null,
  latest_fda_inspection_date: null,
  latest_inspection_classification: null,
  n_inspections_nai: null,
  n_inspections_vai: null,
  n_inspections_oai: null,
  n_483: null,
  n_hc_inspections: null,
  latest_hc_inspection_date: null,
  latest_hc_inspection_rating: null,
  n_hc_inspections_non_compliant: null,
  n_import_refusals: null,
  latest_import_refusal_date: null,
  is_research_facility: null,
  n_aphis_inspections: null,
  latest_aphis_inspection_date: null,
  latest_aphis_inspection_outcome: null,
  n_aphis_inspections_cited: null,
  n_aphis_citations: null,
  n_aphis_repeat_citations: null,
  awa_species_scope: null,
  org_olaw_assurance_id: null,
  org_olaw_assurance_status: null,
  org_olaw_match_tier: null,
  org_olaw_source_asof: null,
  has_org_olaw_assurance: null,
  n_gmp_certificates: null,
  latest_gmp_certificate_date: null,
  n_noncompliance_statements: null,
  latest_noncompliance_statement_date: null,
  n_licences: null,
  licence_status: null,
  org_n_licence_actions: null,
  org_latest_licence_action_date: null,
  as_of: null
)
```


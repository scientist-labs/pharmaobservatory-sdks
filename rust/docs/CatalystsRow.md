# CatalystsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**event_id** | Option<**String**> | The event key: event_type, a colon and source_record_id. Unique. | [optional]
**event_type** | Option<**String**> | What the event is: readout (a trial's primary completion), patent_cliff (loss of exclusivity), pdufa (an FDA decision date or advisory committee) or press_release (a Benzinga press-release event). | [optional]
**event_category** | Option<**String**> | The kind of event within its type. readout: primary_completion. patent_cliff: the cliff basis (curated_loe, orange_book_patent, pte_expiry, bpcia_exclusivity). pdufa: the calendar's own category (pdufa, advisory_committee and others). press_release: the Benzinga event token, lowercased (clinical_trial_result, pdufa_assignment, data_presentation and others). Open vocabulary on the last two. | [optional]
**source** | Option<**String**> | The lane the row came from: ctgov (ClinicalTrials.gov), drug_exclusivity (the exclusivity registers), benzinga, both or sec_extraction (the PDUFA calendar, the same values as pdufa_dates) or benzinga_press_release. | [optional]
**source_record_id** | Option<**String**> | The key of the row in its source, to join back to trials.trial_id, drug_patent_cliff.drug_id, pdufa_dates.pdufa_date_id or the press-release event id. | [optional]
**restriction_tier** | Option<**String**> | none for an open row (ClinicalTrials.gov, the exclusivity registers, the SEC lane of the PDUFA calendar); licensable for a Benzinga row (the calendar or a press release). The api view gates rows on it. | [optional]
**event_date** | Option<[**String**](string.md)> | The date the calendar sorts on: a readout's primary completion date, the cliff date, the decision date, or the end of the press-release event's date window. | [optional]
**event_year** | Option<**i32**> | The calendar year of event_date; the x axis of a by-year count. | [optional]
**date_start** | Option<[**String**](string.md)> | Start of the date interval the source states (the first day of the month for a readout the registry filed by month). | [optional]
**date_end** | Option<[**String**](string.md)> | End of the date interval the source states. | [optional]
**date_precision** | Option<**String**> | Precision of the date (day, month, quarter, half, year, fuzzy, relative). A readout is month where the registry filed only a month, in which case event_date is the last day of that month. | [optional]
**date_phrase** | Option<**String**> | The date as the source words it (an ISO date for a readout and a cliff; the press release's own phrase, such as H2 2026). | [optional]
**date_status** | Option<**String**> | For a readout, whether the registry's completion date is estimated or actual (an anticipated date counts as estimated); null on the other legs. | [optional]
**trial_id** | Option<**String**> | The ClinicalTrials.gov registry id (readout rows). | [optional]
**trial_title** | Option<**String**> | The trial's brief title (readout rows). | [optional]
**phase** | Option<**String**> | The registry's phase code (readout rows). | [optional]
**phase_label** | Option<**String**> | The trial phase as a display label (readout rows). | [optional]
**overall_status** | Option<**String**> | The registry's overall status code (readout rows). | [optional]
**related_trial_ids** | Option<**Vec<String>**> | The registry ids a press-release event names; null on the other legs. | [optional]
**outcome** | Option<**String**> | The outcome a press-release event reports (a decided event stays in the forward book); null on the other legs. | [optional]
**indication_text** | Option<**String**> | The indication as the source words it (PDUFA and press-release rows). | [optional]
**condition_id** | Option<**String**> | The condition key (PDUFA rows). | [optional]
**condition_name** | Option<**String**> | The condition's name (PDUFA rows). | [optional]
**source_type** | Option<**String**> | Kind of source document (SEC filing for the SEC lane, Press release for a press-release event); null on the open legs. | [optional]
**source_url** | Option<**String**> | Link to the source document (PDUFA and press-release rows). | [optional]
**drug_name** | Option<**String**> | The drug name as filed or its display name; null on a readout (its drugs are the list drug_ids). | [optional]
**drug_id** | Option<**String**> | The drug key of a patent-cliff, PDUFA or press-release row; null on a readout and where the drug did not resolve. | [optional]
**molecule_id** | Option<**String**> | The molecule key (PDUFA and press-release rows). | [optional]
**drug_ids** | Option<**Vec<String>**> | Every drug key of the event: a readout carries each drug of its trial, sorted; the other legs carry their one drug, or an empty array. The drug predicate of a watchlist is '<drug>' = ANY(drug_ids). | [optional]
**sponsor_role** | Option<**String**> | What the organization is on the row: lead_sponsor (readout), applicant (patent cliff, the FDA applicant of record), sponsor (PDUFA) or issuer (press release, the company the ticker names). | [optional]
**sponsor_name** | Option<**String**> | The organization's name as filed; null on a press-release row (the source names a ticker, not an organization). | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record (a readout's lead sponsor, a cliff's applicant, a PDUFA row's filing organization; for a press release the identity hub's resolved entity for the ticker). Null when unresolved. | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the organization belongs to (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the organization's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the organization's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the organization's group. | [optional]
**ticker** | Option<**String**> | Exchange ticker (PDUFA and press-release rows). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CompanyFundingEventsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**funding_event_id** | Option<**String**> | The funding event key. | [optional]
**issuer_name** | Option<**String**> | The issuer's name as filed. | [optional]
**issuer_org_id** | Option<**String**> | Legal-entity level: the issuer's organization record. | [optional]
**issuer_company_id** | Option<**String**> | Company level: the issuer's company (null when the issuer did not resolve to a company). | [optional]
**issuer_company_name** | Option<**String**> | Display name of the issuer's company. | [optional]
**issuer_group_id** | Option<**String**> | Group level: the issuer's current ultimate parent company; a root company's group is itself. | [optional]
**issuer_group_name** | Option<**String**> | Display name of the issuer's group. | [optional]
**event_date** | Option<[**String**](string.md)> | The date the source gives for the event; event_stage says what the date means. | [optional]
**fiscal_year** | Option<**i32**> | The NIH fiscal year, for a grant; null for the other sources. | [optional]
**source** | Option<**String**> | The source lane: reporter (NIH RePORTER grants), form_d (SEC Form D Regulation D offerings), edgar_offerings (SEC registered offerings) or sec_extracted (financings extracted from SEC filings). | [optional]
**round_type** | Option<**String**> | The kind of round (for example grant_nih, sbir, reg_d_private, ipo, follow_on, public_offering, atm, registered_direct, debt). | [optional]
**round_subtype** | Option<**String**> | A finer type, such as the grant activity code, the Regulation D exemption or the SEC form. | [optional]
**is_dilutive** | Option<**bool**> | Whether the round dilutes existing equity; null where the instrument is not classified. | [optional]
**amount_usd** | Option<**f64**> | The disclosed amount in whole US dollars; null where the amount is not disclosed. | [optional]
**currency** | Option<**String**> | The currency the source reports the amount in. | [optional]
**amount_is_disclosed** | Option<**bool**> | Whether the source states an amount for the event. | [optional]
**capital_treatment** | Option<**String**> | Whether the event's dollars count as capital raised, and if not why not (for example structured_source, actual_inflow, capacity, committed). | [optional]
**is_actual_issuer_inflow** | Option<**bool**> | True where the event counts as capital the issuer raised. | [optional]
**event_stage** | Option<**String**> | The lifecycle stage the event's date refers to (for example sold_to_date, funded, closed, registered, announced). | [optional]
**administering_ic** | Option<**String**> | The National Institutes of Health institute or center that administers the grant (for example CA for the National Cancer Institute); null for the other sources. | [optional]
**n_awards** | Option<**i64**> | Grant applications folded into the event; 1 for the other sources. | [optional]
**cumulative_raised_usd** | Option<**f64**> | The running total of disclosed amounts over the company's own events up to and including this one, in date order; null where the event has no company. | [optional]
**source_url** | Option<**String**> | Link to the filing the event was read from, where the source has one. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



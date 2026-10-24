# CompanyFirmographicsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | Company level: the company the row is about; the entry key for the companies object type. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**founded_year** | Option<**i32**> | The earliest of the ROR and Wikidata founding years, else our provider profile, else the GLEIF entity creation year last (GLEIF dates re-registrations and reused shells wrongly). | [optional]
**founded_year_source** | Option<**String**> | The source that supplied founded year. | [optional]
**status** | Option<**String**> | public when an organization record of the company carries an SEC filer id (not a Form D-only filer), else private. | [optional]
**hq_city** | Option<**String**> | Headquarters city: GLEIF, then SEC filing business address, then ROR, then Wikidata. | [optional]
**hq_city_source** | Option<**String**> | The source that supplied hq city. | [optional]
**hq_country** | Option<**String**> | Headquarters country: GLEIF, then ROR, then Wikidata. | [optional]
**hq_country_source** | Option<**String**> | The source that supplied hq country. | [optional]
**primary_industry** | Option<**String**> | Primary industry: SEC standard industrial classification description, then ROR organization type, then Wikidata industry. | [optional]
**primary_industry_source** | Option<**String**> | The source that supplied primary industry. | [optional]
**website** | Option<**String**> | Website: ROR, then Wikidata, then our provider profile. | [optional]
**website_source** | Option<**String**> | The source that supplied website. | [optional]
**employees** | Option<**i64**> | Latest employee headcount from Wikidata; a zero or negative value is treated as absent. | [optional]
**employees_source** | Option<**String**> | The source that supplied employees. | [optional]
**employees_as_of** | Option<[**String**](string.md)> | The date Wikidata states for the headcount. | [optional]
**verticals** | Option<**Vec<String>**> | Service areas from our provider profile (the pipe-separated list, split). | [optional]
**verticals_source** | Option<**String**> | The source that supplied verticals. | [optional]
**legal_name** | Option<**String**> | Legal name: GLEIF, then Wikidata, then our provider profile's official name. | [optional]
**legal_name_source** | Option<**String**> | The source that supplied legal name. | [optional]
**corporate_parent** | Option<**String**> | Corporate parent name: GLEIF group name, then Wikidata parent, then our provider profile's parent. This is a stated source value, not the company's ladder group. | [optional]
**corporate_parent_source** | Option<**String**> | The source that supplied corporate parent. | [optional]
**legal_form** | Option<**String**> | Legal form: GLEIF, then Wikidata. | [optional]
**legal_form_source** | Option<**String**> | The source that supplied legal form. | [optional]
**wikidata_qid** | Option<**String**> | The Wikidata item id of the organization record whose profile supplies the five fields below. | [optional]
**wikipedia_url** | Option<**String**> | Wikipedia page: ROR, then Wikidata. | [optional]
**linkedin** | Option<**String**> | LinkedIn organization id from the Wikidata profile. | [optional]
**x_handle** | Option<**String**> | X (Twitter) username from the Wikidata profile. | [optional]
**facebook** | Option<**String**> | Facebook id from the Wikidata profile. | [optional]
**short_description** | Option<**String**> | The Wikidata short description. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InvestigatorPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investigator_person_id** | Option<**String**> | The investigator's resolved person key, the id in the web page address (/people/<id>). The entry key of the investigators object type. | [optional]
**pmid** | Option<**String**> | The PubMed id of the publication. | [optional]
**link_basis** | Option<**String**> | authored when the person is linked to the paper as an author; trial_reference when a trial the person is on registers the paper (not authorship). | [optional]
**author_position** | Option<**String**> | first, middle or last for an authored row; null for a trial_reference row. | [optional]
**author_match_method** | Option<**String**> | How the author was matched to the person (orcid, name_on_linked_trial and so on); null for a trial_reference row. | [optional]
**citation** | Option<**String**> | The citation as a registered trial cites the paper (null when no trial cites it). | [optional]
**title** | Option<**String**> | The paper's title (null when no registered trial cites it). | [optional]
**pub_year** | Option<**i32**> | The publication year (null when no registered trial cites the paper). | [optional]
**n_trials_citing** | Option<**i32**> | Registered trials that cite the paper (null when none does). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



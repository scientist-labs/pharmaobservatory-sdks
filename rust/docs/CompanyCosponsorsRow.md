# CompanyCosponsorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | Company level: the company the row is about; the entry key for the companies object type. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**row_level** | Option<**String**> | root: the co-sponsor's ultimate parent company, with the distinct trials any company of that root shares with this company. member: a company under that root that itself co-sponsored, with its own count. Filter on one level; the two do not add. | [optional]
**cosponsor_company_id** | Option<**String**> | Company level: the co-sponsoring company. At row_level root this is the root company; at member it is the member company itself. | [optional]
**cosponsor_company_name** | Option<**String**> | Display name of the co-sponsoring company. | [optional]
**cosponsor_group_id** | Option<**String**> | Group level: the co-sponsor's current ultimate parent; a root's group is itself. | [optional]
**cosponsor_group_name** | Option<**String**> | Display name of the co-sponsor's group. | [optional]
**cosponsor_root_id** | Option<**String**> | The ultimate parent company of the co-sponsor, the ladder company's `coalesce(ultimate_parent_company_id, company_id)`. Equal to cosponsor_company_id at row_level root, and the root a member row sits under. | [optional]
**cosponsor_rank** | Option<**i32**> | The rank of the ROOT among this company's co-sponsor roots by distinct shared trials (1 is the most; ties broken by root id). The same on a root row and on each of its member rows. | [optional]
**member_rank** | Option<**i32**> | 0 on a root row; 1 and up on the member rows of that root, by the member's own shared trials. Order by cosponsor_rank then member_rank to list each root followed by its members. | [optional]
**n_members** | Option<**i32**> | Member companies of this root that co-sponsored with this company, before the member cut. The same on a root row and its member rows. | [optional]
**n_candidates** | Option<**i32**> | Distinct co-sponsor roots for this company before the top-25 cut; the same on every row of a company. | [optional]
**n_shared_trials** | Option<**i32**> | Distinct trials that name both. At row_level root, the distinct trials of any company under the root (never a sum over members); at member, that company's own trials. | [optional]
**n_trials_company_lead** | Option<**i32**> | Of the shared trials, those where an organization record of this company holds the ClinicalTrials.gov lead sponsor role. At most n_shared_trials; it may overlap n_trials_cosponsor_lead. | [optional]
**n_trials_cosponsor_lead** | Option<**i32**> | Of the shared trials, those where the co-sponsor's organization record holds the lead sponsor role. At most n_shared_trials. | [optional]
**first_shared_start_year** | Option<**i32**> | The earliest registered start year over the shared trials; null when none has a plausible year. | [optional]
**last_shared_start_year** | Option<**i32**> | The latest registered start year over the shared trials; null when none has a plausible year. | [optional]
**n_shared_conditions** | Option<**i32**> | Distinct specific MeSH disease areas over the shared trials. | [optional]
**top_shared_conditions** | Option<**Vec<String>**> | Up to 5 MeSH disease names by shared-trial count then name. An empty list when none. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



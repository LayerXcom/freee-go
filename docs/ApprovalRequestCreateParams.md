# ApprovalRequestCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**ApprovalFlowRouteId** | **int32** | 申請経路ID | 
**FormId** | **int32** | 申請フォームID | 
**ApproverId** | **int32** | 承認者のユーザーID | [optional] 
**Draft** | **bool** | falseの時、in_progress:申請中で作成する。それ以外の時はdraft:下書きで作成する | 
**ParentId** | **int32** | 親申請ID(既存各種申請IDのみ指定可能です。) | [optional] 
**RequestItems** | [**[]ApprovalRequestCreateParamsRequestItems**](approvalRequestCreateParams_request_items.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



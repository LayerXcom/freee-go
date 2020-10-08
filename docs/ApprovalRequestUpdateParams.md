# ApprovalRequestUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**ApprovalFlowRouteId** | **int32** | 申請経路ID | 
**ApproverId** | **int32** | 承認者のユーザーID | [optional] 
**Draft** | **bool** | falseの時、in_progress:申請中で更新する。それ以外の時はdraft:下書きで更新する | 
**RequestItems** | [**[]ApprovalRequestCreateParamsRequestItems**](approvalRequestCreateParams_request_items.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



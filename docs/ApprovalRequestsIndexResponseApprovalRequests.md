# ApprovalRequestsIndexResponseApprovalRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 各種申請ID | 
**CompanyId** | **int32** | 事業所ID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Title** | **string** | 申請タイトル | 
**ApplicantId** | **int32** | 申請者のユーザーID | 
**ApplicationNumber** | **string** | 申請No. | 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**RequestItems** | [**[]ApprovalRequestsIndexResponseRequestItems**](approvalRequestsIndexResponse_request_items.md) | 各種申請の項目一覧（配列） | 
**FormId** | **int32** | 申請フォームID | 
**CurrentStepId** | Pointer to **int32** | 現在承認ステップID | 
**CurrentRound** | **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ApprovalRequestResponseApprovalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 各種申請ID | 
**CompanyId** | **int32** | 事業所ID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Title** | **string** | 申請タイトル | 
**ApplicantId** | **int32** | 申請者のユーザーID | 
**Approvers** | [**[]ExpenseApplicationResponseExpenseApplicationApprovers**](expenseApplicationResponse_expense_application_approvers.md) | 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、    approversはレスポンスに含まれるようになります。    その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。 | 
**ApplicationNumber** | **string** | 申請No. | 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**RequestItems** | [**[]ApprovalRequestsIndexResponseRequestItems**](approvalRequestsIndexResponse_request_items.md) | 各種申請の項目一覧（配列） | 
**FormId** | **int32** | 申請フォームID | 
**ApprovalFlowRouteId** | **int32** | 申請経路ID | 
**Comments** | [**[]ExpenseApplicationResponseExpenseApplicationComments**](expenseApplicationResponse_expense_application_comments.md) | 各種申請のコメント一覧（配列） | 
**ApprovalFlowLogs** | [**[]ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs**](expenseApplicationResponse_expense_application_approval_flow_logs.md) | 各種申請の承認履歴（配列） | 
**CurrentStepId** | Pointer to **int32** | 現在承認ステップID | 
**CurrentRound** | **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 
**ApprovalRequestForm** | [**ApprovalRequestResponseApprovalRequestApprovalRequestForm**](approvalRequestResponse_approval_request_approval_request_form.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



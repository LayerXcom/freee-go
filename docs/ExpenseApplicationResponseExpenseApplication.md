# ExpenseApplicationResponseExpenseApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 経費申請ID | 
**CompanyId** | **int32** | 事業所ID | 
**Title** | **string** | 申請タイトル | 
**IssueDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Description** | Pointer to **string** | 備考 | [optional] 
**EditableOnWeb** | **bool** | 会計freeeのWeb画面から申請内容を編集可能：falseの場合、Web上からの項目行の追加／削除・金額の編集が出来なくなります。APIでの編集は可能です。 | 
**TotalAmount** | **int32** | 合計金額 | [optional] 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**TagIds** | **[]int32** | メモタグID | [optional] 
**ExpenseApplicationLines** | [**[]ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines**](expenseApplicationResponse_expense_application_expense_application_lines.md) | 経費申請の項目行一覧（配列） | 
**DealId** | Pointer to **int32** | 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます) | 
**DealStatus** | Pointer to **string** | 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち) | 
**ApplicantId** | **int32** | 申請者のユーザーID | 
**Approvers** | [**[]ExpenseApplicationResponseExpenseApplicationApprovers**](expenseApplicationResponse_expense_application_approvers.md) | 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、    approversはレスポンスに含まれるようになります。    その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。 | 
**ApplicationNumber** | **string** | 申請No. | 
**ApprovalFlowRouteId** | **int32** | 申請経路ID | 
**Comments** | [**[]ExpenseApplicationResponseExpenseApplicationComments**](expenseApplicationResponse_expense_application_comments.md) | 経費申請のコメント一覧（配列） | 
**ApprovalFlowLogs** | [**[]ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs**](expenseApplicationResponse_expense_application_approval_flow_logs.md) | 経費申請の承認履歴（配列） | 
**CurrentStepId** | Pointer to **int32** | 現在承認ステップID | 
**CurrentRound** | **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 
**Segment1TagId** | Pointer to **int32** | セグメント１ID。セグメント１が使用可能なプランの時のみレスポンスに含まれます。 | [optional] 
**Segment2TagId** | Pointer to **int32** | セグメント２ID。セグメント２が使用可能なプランの時のみレスポンスに含まれます。 | [optional] 
**Segment3TagId** | Pointer to **int32** | セグメント３ID。セグメント３が使用可能なプランの時のみレスポンスに含まれます。 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



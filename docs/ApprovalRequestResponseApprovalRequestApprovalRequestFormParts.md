# ApprovalRequestResponseApprovalRequestApprovalRequestFormParts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 項目ID | 
**Order** | **int32** | 順序 | [optional] 
**Type** | **string** | 項目種別 (title: 申請タイトル, single_line: 自由記述形式 1行, multi_line: 自由記述形式 複数行, select: プルダウン, date: 日付, amount: 金額, receipt: 添付ファイル) | [optional] 
**Label** | **string** | 項目名 | [optional] 
**Annotation** | Pointer to **string** | 追加説明 | [optional] 
**Required** | Pointer to **bool** | 必須かどうか | [optional] 
**Values** | Pointer to [**[]ApprovalRequestResponseApprovalRequestApprovalRequestFormValues**](approvalRequestResponse_approval_request_approval_request_form_values.md) | 選択項目 | [optional] 
**MaxAmount** | Pointer to **int32** | 上限金額 | [optional] 
**MinAmount** | Pointer to **int32** | 下限金額 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



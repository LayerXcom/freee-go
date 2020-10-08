# ApprovalRequestFormResponseApprovalRequestForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 申請フォームID | 
**CompanyId** | **int32** | 事業所ID | 
**Name** | **string** | 申請フォームの名前 | 
**Description** | **string** | 申請フォームの説明 | 
**Status** | **string** | ステータス(draft: 申請で使用しない、active: 申請で使用する、deleted: 削除済み) | 
**CreatedDate** | **string** | 作成日時 | 
**FormOrder** | Pointer to **int32** | 表示順（申請者が選択する申請フォームの表示順を設定できます。小さい数ほど上位に表示されます。（0を除く整数のみ。マイナス不可）未入力の場合、表示順が後ろになります。同じ数字が入力された場合、登録順で表示されます。） | 
**Parts** | [**[]ApprovalRequestResponseApprovalRequestApprovalRequestFormParts**](approvalRequestResponse_approval_request_approval_request_form_parts.md) | 申請フォームの項目 | [optional] 
**RouteSettingCount** | **int32** | 適用された経路数 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



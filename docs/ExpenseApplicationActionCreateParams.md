# ExpenseApplicationActionCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**ApprovalAction** | **string** | 操作(approve: 承認する、force_approve: 代理承認する、cancel: 申請を取り消す、reject: 却下する、feedback: 申請者へ差し戻す、force_feedback: 承認済み・却下済みを取り消す) | 
**TargetStepId** | **int32** | 対象承認ステップID 経費申請の取得APIレスポンス.current_step_idを送信してください。 | 
**TargetRound** | **int32** | 対象round。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。経費申請の取得APIレスポンス.current_roundを送信してください。 | 
**NextApproverId** | Pointer to **int32** | 次ステップの承認者のユーザーID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



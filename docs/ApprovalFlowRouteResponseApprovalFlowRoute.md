# ApprovalFlowRouteResponseApprovalFlowRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 申請経路ID | 
**Name** | **string** | 申請経路名 | [optional] 
**Description** | **string** | 申請経路の説明 | [optional] 
**UserId** | Pointer to **int32** | 更新したユーザーのユーザーID | [optional] 
**DefinitionSystem** | **bool** | システム作成の申請経路かどうか | [optional] 
**FirstStepId** | **int32** | 最初の承認ステップのID | [optional] 
**Usages** | **[]string** | 申請種別（申請経路を使用できる申請種別を示します。例えば、ApprovalRequest の場合は、各種申請で使用できる申請経路です。） * &#x60;TxnApproval&#x60; - 仕訳承認 * &#x60;ExpenseApplication&#x60; - 経費精算 * &#x60;PaymentRequest&#x60; - 支払依頼 * &#x60;ApprovalRequest&#x60; - 各種申請 * &#x60;DocApproval&#x60; - 請求書等 (見積書・納品書・請求書・発注書) | [optional] 
**RequestFormIds** | **[]int32** | 申請経路で利用できる申請フォームID配列 | 
**Steps** | [**[]ApprovalFlowRouteResponseApprovalFlowRouteSteps**](approvalFlowRouteResponse_approval_flow_route_steps.md) | 承認ステップ（配列） | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



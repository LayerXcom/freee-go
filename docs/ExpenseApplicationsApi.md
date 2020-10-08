# \ExpenseApplicationsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExpenseApplication**](ExpenseApplicationsApi.md#CreateExpenseApplication) | **Post** /api/1/expense_applications | 経費申請の作成
[**DestroyExpenseApplication**](ExpenseApplicationsApi.md#DestroyExpenseApplication) | **Delete** /api/1/expense_applications/{id} | 経費申請の削除
[**GetExpenseApplication**](ExpenseApplicationsApi.md#GetExpenseApplication) | **Get** /api/1/expense_applications/{id} | 経費申請詳細の取得
[**GetExpenseApplications**](ExpenseApplicationsApi.md#GetExpenseApplications) | **Get** /api/1/expense_applications | 経費申請一覧の取得
[**UpdateExpenseApplication**](ExpenseApplicationsApi.md#UpdateExpenseApplication) | **Put** /api/1/expense_applications/{id} | 経費申請の更新
[**UpdateExpenseApplicationAction**](ExpenseApplicationsApi.md#UpdateExpenseApplicationAction) | **Post** /api/1/expense_applications/{id}/actions | 経費申請の承認操作



## CreateExpenseApplication

> ExpenseApplicationResponse CreateExpenseApplication(ctx, optional)

経費申請の作成

 <h2 id=\"_1\">概要</h2>  <p>指定した事業所の経費申請を作成する</p>  <p>経費申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-expense-applications\" target=\"_blank\">会計freeeの経費申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>     申請ステータス(下書き、申請中)の指定と変更、及び承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）は以下を参考にして行ってください。     <ul>       <li>         承認操作は申請ステータスが申請中、承認済み、却下のものだけが対象です。         <ul>           <li>             初回申請の場合             <ul><li>申請の作成（POST）</li></ul>           </li>           <li>             作成済みの申請の申請ステータス変更・更新する場合             <ul><li>申請の更新（PUT）</li></ul>           </li>           <li>             申請中、承認済み、却下の申請の承認操作を行う場合             <ul><li>承認操作の実行（POST）</li></ul>           </li>         </ul>       </li>       <li>申請の削除（DELETE）が可能なのは申請ステータスが下書き、差戻しの場合のみです</li>     </ul>   </li>   <li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している経費申請は本API経由で作成ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>申請時には、申請タイトル(title)に加え、申請日(issue_date)、項目行については金額(amount)、日付(transaction_date)、内容(description)が必須項目となります。申請時の業務効率化のため、API入力をお勧めします。</li>   <li>本APIは駅すぱあと連携 (出発駅と到着駅から金額を自動入力する機能)には非対応です。駅すぱあと連携を使用した経費申請は作成できません。</li>   <li>個人アカウントの場合は、プレミアムプランでご利用できます。</li>   <li>法人アカウントの場合は、ベーシックプラン、プロフェッショナルプラン、エンタープライズプランでご利用できます。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateExpenseApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateExpenseApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expenseApplicationCreateParams** | [**optional.Interface of ExpenseApplicationCreateParams**](ExpenseApplicationCreateParams.md)| 経費申請の作成 | 

### Return type

[**ExpenseApplicationResponse**](expenseApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyExpenseApplication

> DestroyExpenseApplication(ctx, id, companyId)

経費申請の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の経費申請を削除する</p>  <p>経費申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-expense-applications\" target=\"_blank\">会計freeeの経費申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>     申請ステータス(下書き、申請中)の指定と変更、及び承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）は以下を参考にして行ってください。     <ul>       <li>         承認操作は申請ステータスが申請中、承認済み、却下のものだけが対象です。         <ul>           <li>             初回申請の場合             <ul><li>申請の作成（POST）</li></ul>           </li>           <li>             作成済みの申請の申請ステータス変更・更新する場合             <ul><li>申請の更新（PUT）</li></ul>           </li>           <li>             申請中、承認済み、却下の申請の承認操作を行う場合             <ul><li>承認操作の実行（POST）</li></ul>           </li>         </ul>       </li>       <li>申請の削除（DELETE）が可能なのは申請ステータスが下書き、差戻しの場合のみです</li>       <li>自分が申請者でない申請の削除が可能なのはユーザーの権限が管理者権限、且つ申請ステータスが差し戻しの場合のみです</li>     </ul>   </li>   <li>本APIは駅すぱあと連携 (出発駅と到着駅から金額を自動入力する機能)には非対応です。駅すぱあと連携を使用した経費申請は削除できません。</li>   <li>個人アカウントの場合は、プレミアムプランでご利用できます。</li>   <li>法人アカウントの場合は、ベーシックプラン、プロフェッショナルプラン、エンタープライズプランでご利用できます。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 経費申請ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseApplication

> ExpenseApplicationResponse GetExpenseApplication(ctx, id, companyId)

経費申請詳細の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の経費申請を取得する</p>  <p>経費申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-expense-applications\" target=\"_blank\">会計freeeの経費申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している経費申請と申請経路はAPI経由で参照は可能ですが、作成と更新、承認ステータスの変更ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>本APIは駅すぱあと連携 (出発駅と到着駅から金額を自動入力する機能)には非対応です。駅すぱあと連携を使用した経費申請は取得できません。</li>   <li>個人アカウントの場合は、プレミアムプランでご利用できます。</li>   <li>法人アカウントの場合は、ベーシックプラン、プロフェッショナルプラン、エンタープライズプランでご利用できます。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 経費申請ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**ExpenseApplicationResponse**](expenseApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseApplications

> ExpenseApplicationsIndexResponse GetExpenseApplications(ctx, companyId, optional)

経費申請一覧の取得

 <h2 id=\"_1\">概要</h2>  <p>指定した事業所の経費申請一覧を取得する</p>  <p>経費申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-expense-applications\" target=\"_blank\">会計freeeの経費申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIでは、経費申請の一覧を取得することができます。</li>   <li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している経費申請と申請経路はAPI経由で参照は可能ですが、作成と更新、承認ステータスの変更ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>個人アカウントの場合は、プレミアムプランでご利用できます。</li>   <li>法人アカウントの場合は、ベーシックプラン、プロフェッショナルプラン、エンタープライズプランでご利用できます。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetExpenseApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExpenseApplicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)、 取引ステータス(unsettled:清算待ち, settled:精算済み) | 
 **payrollAttached** | **optional.Bool**| true:給与連携あり、false:給与連携なし、未指定時:絞り込みなし | 
 **startTransactionDate** | **optional.String**| 発生日(経費申請項目の日付)で絞込：開始日(yyyy-mm-dd) | 
 **endTransactionDate** | **optional.String**| 発生日(経費申請項目の日付)で絞込：終了日(yyyy-mm-dd) | 
 **applicationNumber** | **optional.Int32**| 申請No. | 
 **title** | **optional.String**| 申請タイトル | 
 **startIssueDate** | **optional.String**| 申請日で絞込：開始日(yyyy-mm-dd) | 
 **endIssueDate** | **optional.String**| 申請日で絞込：終了日(yyyy-mm-dd) | 
 **applicantId** | **optional.Int32**| 申請者のユーザーID | 
 **approverId** | **optional.Int32**| 承認者のユーザーID | 
 **minAmount** | **optional.Int32**| 金額で絞込 (下限金額) | 
 **maxAmount** | **optional.Int32**| 金額で絞込 (上限金額) | 
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 500) | 

### Return type

[**ExpenseApplicationsIndexResponse**](expenseApplicationsIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExpenseApplication

> ExpenseApplicationResponse UpdateExpenseApplication(ctx, id, optional)

経費申請の更新

 <h2 id=\"\">概要</h2>  <p>指定した事業所の経費申請を更新する</p>  <p>経費申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-expense-applications\" target=\"_blank\">会計freeeの経費申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIでは、経費申請を更新することができます。</li>   <li>本APIでは、status(申請ステータス): draft:下書き, feedback:差戻しのみ更新可能です。</li>   <li>     申請ステータス(下書き、申請中)の指定と変更、及び承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）は以下を参考にして行ってください。     <ul>       <li>         承認操作は申請ステータスが申請中、承認済み、却下のものだけが対象です。         <ul>           <li>             初回申請の場合             <ul><li>申請の作成（POST）</li></ul>           </li>           <li>             作成済みの申請の申請ステータス変更・更新する場合             <ul><li>申請の更新（PUT）</li></ul>           </li>           <li>             申請中、承認済み、却下の申請の承認操作を行う場合             <ul><li>承認操作の実行（POST）</li></ul>           </li>         </ul>       </li>       <li>申請の削除（DELETE）が可能なのは申請ステータスが下書き、差戻しの場合のみです</li>     </ul>   </li>   <li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している経費申請は本API経由で更新ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>申請時には、申請タイトル(title)に加え、申請日(issue_date)、項目行については金額(amount)、日付(transaction_date)、内容(description)が必須項目となります。申請時の業務効率化のため、API入力をお勧めします。</li>   <li>本APIは駅すぱあと連携 (出発駅と到着駅から金額を自動入力する機能)には非対応です。駅すぱあと連携を使用した経費申請は更新できません。</li>   <li>個人アカウントの場合は、プレミアムプランでご利用できます。</li>   <li>法人アカウントの場合は、ベーシックプラン、プロフェッショナルプラン、エンタープライズプランでご利用できます。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 経費申請ID | 
 **optional** | ***UpdateExpenseApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateExpenseApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expenseApplicationUpdateParams** | [**optional.Interface of ExpenseApplicationUpdateParams**](ExpenseApplicationUpdateParams.md)| 経費申請の更新 | 

### Return type

[**ExpenseApplicationResponse**](expenseApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExpenseApplicationAction

> ExpenseApplicationResponse UpdateExpenseApplicationAction(ctx, id, expenseApplicationActionCreateParams)

経費申請の承認操作

 <h2 id=\"_1\">概要</h2>  <p>指定した事業所の経費申請の承認操作を行う</p>  <p>経費申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-expense-applications\" target=\"_blank\">会計freeeの経費申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIでは、経費申請の承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）を行うことができます。</li>   <li>     申請ステータス(下書き、申請中)の指定と変更、及び承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）は以下を参考にして行ってください。     <ul>       <li>         承認操作は申請ステータスが申請中、承認済み、却下のものだけが対象です。         <ul>           <li>             初回申請の場合             <ul><li>申請の作成（POST）</li></ul>           </li>           <li>             作成済みの申請の申請ステータス変更・更新する場合             <ul><li>申請の更新（PUT）</li></ul>           </li>           <li>             申請中、承認済み、却下の申請の承認操作を行う場合             <ul><li>承認操作の実行（POST）</li></ul>           </li>         </ul>       </li>       <li>申請の削除（DELETE）が可能なのは申請ステータスが下書き、差戻しの場合のみです</li>     </ul>   </li> 　<li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している経費申請はAPI経由で承認ステータスの変更ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>本APIは駅すぱあと連携 (出発駅と到着駅から金額を自動入力する機能)には非対応です。駅すぱあと連携を使用した経費申請は承認操作できません。</li>   <li>個人アカウントの場合は、プレミアムプランでご利用できます。</li>   <li>法人アカウントの場合は、ベーシックプラン、プロフェッショナルプラン、エンタープライズプランでご利用できます。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 経費申請ID | 
**expenseApplicationActionCreateParams** | [**ExpenseApplicationActionCreateParams**](ExpenseApplicationActionCreateParams.md)| 経費申請の承認操作 | 

### Return type

[**ExpenseApplicationResponse**](expenseApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


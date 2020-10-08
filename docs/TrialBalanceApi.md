# \TrialBalanceApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrialBs**](TrialBalanceApi.md#GetTrialBs) | **Get** /api/1/reports/trial_bs | 貸借対照表の取得
[**GetTrialBsThreeYears**](TrialBalanceApi.md#GetTrialBsThreeYears) | **Get** /api/1/reports/trial_bs_three_years | 貸借対照表(３期間比較)の取得
[**GetTrialBsTwoYears**](TrialBalanceApi.md#GetTrialBsTwoYears) | **Get** /api/1/reports/trial_bs_two_years | 貸借対照表(前年比較)の取得
[**GetTrialPl**](TrialBalanceApi.md#GetTrialPl) | **Get** /api/1/reports/trial_pl | 損益計算書の取得
[**GetTrialPlSections**](TrialBalanceApi.md#GetTrialPlSections) | **Get** /api/1/reports/trial_pl_sections | 損益計算書(部門比較)の取得
[**GetTrialPlThreeYears**](TrialBalanceApi.md#GetTrialPlThreeYears) | **Get** /api/1/reports/trial_pl_three_years | 損益計算書(３期間比較)の取得
[**GetTrialPlTwoYears**](TrialBalanceApi.md#GetTrialPlTwoYears) | **Get** /api/1/reports/trial_pl_two_years | 損益計算書(前年比較)の取得



## GetTrialBs

> TrialBsResponse GetTrialBs(ctx, companyId, optional)

貸借対照表の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の貸借対照表を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul>  <li> <p>created_at : 作成日時</p> </li>  <li> <p>account_item_name : 勘定科目名</p> </li>  <li> <p>hierarchy_level: 階層レベル</p> </li>  <li> <p>parent_account_category_name: 上位勘定科目カテゴリー名</p> </li> <li> <p>opening_balance : 期首残高 </p> </li>  <li> <p>debit_amount : 借方金額 </p> </li> <li> <p>credit_amount:  貸方金額 </p> </li> <li> <p>closing_balance : 期末残高 </p> </li> <li> <p>composition_ratio : 構成比</p> </li> <h2 id=\"_3\">注意点</h2> <ul> <li>会計年度が指定されない場合、現在の会計年度がデフォルトとなります。</li> <li>絞り込み条件の日付と、月または年度は同時に指定することはできません。</li> <li>up_to_dateがfalseの場合、残高の集計が完了していません。最新の集計結果を確認したい場合は、時間を空けて再度取得する必要があります。</li>  </ul> <h2 id=\"_4\">レスポンスの例</h2>  <blockquote> <p>GET https://api.freee.co.jp/api/1/reports/trial_bs?company_id=1&amp;fiscal_year=2017&amp;breakdown_display_type=partner</p> </blockquote>  <pre><code>{   &quot;trial_bs&quot; :     {       &quot;company_id&quot; : 1,       &quot;fiscal_year&quot; : 2017,       &quot;breakdown_display_type&quot; : &quot;partner&quot;,       &quot;created_at&quot; : &quot;2018-05-01 12:00:50&quot       &quot;balances&quot; : [{         &quot;account_item_id&quot; : 1000,         &quot;account_item_name&quot; : &quot;現金&quot;,         &quot;hierarchy_level&quot; : 2,         &quot;account_category_name&quot; : &quot;流動資産&quot;,         &quot;opening_balance&quot; : 100000,         &quot;debit_amount&quot; : 50000,         &quot;credit_amount&quot; : 20000,         &quot;closing_balance&quot; : 130000,         &quot;composition_ratio&quot; : 0.25         &quot;partners&quot; : [{           &quot;id&quot; : 123,           &quot;name&quot; : &quot;freee&quot;,           &quot;opening_balance&quot; : 100000,           &quot;debit_amount&quot; : 50000,           &quot;credit_amount&quot; : 20000,           &quot;closing_balance&quot; : 130000,           &quot;composition_ratio&quot; : 0.25           },         ...         ]       },       ...       ]     } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetTrialBsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrialBsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fiscalYear** | **optional.Int32**| 会計年度 | 
 **startMonth** | **optional.Int32**| 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **optional.Int32**| 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **optional.String**| 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **optional.String**| 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **optional.String**| 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **optional.String**| 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **optional.Int32**| 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **optional.String**| 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **optional.Int32**| 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **optional.String**| 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 

### Return type

[**TrialBsResponse**](trialBsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialBsThreeYears

> TrialBsThreeYearsResponse GetTrialBsThreeYears(ctx, companyId, optional)

貸借対照表(３期間比較)の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の貸借対照表(３期間比較)を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul>  <li> <p>created_at : 作成日時</p> </li>  <li> <p>account_item_name : 勘定科目名</p> </li>  <li> <p>hierarchy_level: 階層レベル</p> </li>  <li> <p>parent_account_category_name: 上位勘定科目カテゴリー名</p> </li> <li> <p>two_years_before_closing_balance:  前々年度期末残高 </p> </li> <li> <p>last_year_closing_balance:  前年度期末残高 </p> </li> <li> <p>closing_balance : 期末残高 </p> </li> <li> <p>year_on_year : 前年比</p> </li> <h2 id=\"_3\">注意点</h2> <ul> <li>会計年度が指定されない場合、現在の会計年度がデフォルトとなります。</li> <li>絞り込み条件の日付と、月または年度は同時に指定することはできません。</li> <li>up_to_dateがfalseの場合、残高の集計が完了していません。最新の集計結果を確認したい場合は、時間を空けて再度取得する必要があります。</li>  </ul> <h2 id=\"_4\">レスポンスの例</h2>  <blockquote> <p>GET https://api.freee.co.jp/api/1/reports/trial_bs_three_years?company_id=1&amp;fiscal_year=2017</p> </blockquote>  <pre><code>{   &quot;trial_bs_three_years&quot; :     {       &quot;company_id&quot; : 1,       &quot;fiscal_year&quot; : 2017,       &quot;created_at&quot; : &quot;2018-05-01 12:00:50&quot       &quot;balances&quot; : [{         &quot;account_item_id&quot; : 1000,         &quot;account_item_name&quot; : &quot;現金&quot;,         &quot;hierarchy_level&quot; : 2,         &quot;account_category_name&quot; : &quot;流動資産&quot;,         &quot;two_year_before_closing_balance&quot; : 50000,         &quot;last_year_closing_balance&quot; : 25000,         &quot;closing_balance&quot; : 100000,         &quot;year_on_year&quot; : 0.85       },       ...       ]     } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetTrialBsThreeYearsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrialBsThreeYearsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fiscalYear** | **optional.Int32**| 会計年度 | 
 **startMonth** | **optional.Int32**| 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **optional.Int32**| 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **optional.String**| 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **optional.String**| 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **optional.String**| 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **optional.String**| 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **optional.Int32**| 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **optional.String**| 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **optional.Int32**| 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **optional.String**| 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 

### Return type

[**TrialBsThreeYearsResponse**](trialBsThreeYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialBsTwoYears

> TrialBsTwoYearsResponse GetTrialBsTwoYears(ctx, companyId, optional)

貸借対照表(前年比較)の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の貸借対照表(前年比較)を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul>  <li> <p>created_at : 作成日時</p> </li>  <li> <p>account_item_name : 勘定科目名</p> </li>  <li> <p>hierarchy_level: 階層レベル</p> </li>  <li> <p>parent_account_category_name: 上位勘定科目カテゴリー名</p> </li> <li> <p>last_year_closing_balance:  前年度期末残高 </p> </li> <li> <p>closing_balance : 期末残高 </p> </li> <h2 id=\"_3\">注意点</h2> <ul> <li>会計年度が指定されない場合、現在の会計年度がデフォルトとなります。</li> <li>絞り込み条件の日付と、月または年度は同時に指定することはできません。</li> <li>up_to_dateがfalseの場合、残高の集計が完了していません。最新の集計結果を確認したい場合は、時間を空けて再度取得する必要があります。</li>  </ul>  <h2 id=\"_4\">レスポンスの例</h2>  <blockquote> <p>GET https://api.freee.co.jp/api/1/reports/trial_bs_two_years?company_id=1&amp;fiscal_year=2017</p> </blockquote>  <pre><code>{   &quot;trial_bs_two_years&quot; :     {       &quot;company_id&quot; : 1,       &quot;fiscal_year&quot; : 2017,       &quot;created_at&quot; : &quot;2018-05-01 12:00:50&quot       &quot;balances&quot; : [{         &quot;account_item_id&quot; : 1000,         &quot;account_item_name&quot; : &quot;現金&quot;,         &quot;hierarchy_level&quot; : 2,         &quot;account_category_name&quot; : &quot;流動資産&quot;,         &quot;last_year_closing_balance&quot; : 25000,         &quot;closing_balance&quot; : 100000,         &quot;year_on_year&quot; : 0.85        },       ...       ]     } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetTrialBsTwoYearsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrialBsTwoYearsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fiscalYear** | **optional.Int32**| 会計年度 | 
 **startMonth** | **optional.Int32**| 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **optional.Int32**| 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **optional.String**| 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **optional.String**| 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **optional.String**| 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **optional.String**| 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **optional.Int32**| 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **optional.String**| 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **optional.Int32**| 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **optional.String**| 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 

### Return type

[**TrialBsTwoYearsResponse**](trialBsTwoYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPl

> TrialPlResponse GetTrialPl(ctx, companyId, optional)

損益計算書の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の損益計算書を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul>  <li> <p>created_at : 作成日時</p> </li>  <li> <p>account_item_name : 勘定科目名</p> </li>  <li> <p>hierarchy_level: 階層レベル</p> </li>  <li> <p>parent_account_category_name: 上位勘定科目カテゴリー名</p> </li> <li> <p>opening_balance : 期首残高 </p> </li>  <li> <p>debit_amount : 借方金額 </p> </li> <li> <p>credit_amount:  貸方金額 </p> </li> <li> <p>closing_balance : 期末残高 </p> </li> <li> <p>composition_ratio : 構成比</p> </li> <h2 id=\"_3\">注意点</h2> <ul> <li>会計年度が指定されない場合、現在の会計年度がデフォルトとなります。</li> <li>絞り込み条件の日付と、月または年度は同時に指定することはできません。</li> <li>up_to_dateがfalseの場合、残高の集計が完了していません。最新の集計結果を確認したい場合は、時間を空けて再度取得する必要があります。</li> <li>配賦仕訳の絞り込み（cost_allocation）は法人向けのベーシックプラン以上で利用可能です。</li> </ul> <h2 id=\"_4\">レスポンスの例</h2>  <blockquote> <p>GET https://api.freee.co.jp/api/1/reports/trial_pl?company_id=1&amp;fiscal_year=2017&amp;breakdown_display_type=partner</p> </blockquote>  <pre><code>{   &quot;trial_pl&quot; :     {       &quot;company_id&quot; : 1,       &quot;fiscal_year&quot; : 2017,       &quot;breakdown_display_type&quot; : &quot;partner&quot;,       &quot;created_at&quot; : &quot;2018-05-01 12:00:50&quot       &quot;balances&quot; : [{         &quot;account_item_id&quot; : 1500,         &quot;account_item_name&quot; : &quot;売上高&quot;,         &quot;hierarchy_level&quot; : 2,         &quot;account_category_name&quot; : &quot;営業収益&quot;,         &quot;opening_balance&quot; : 100000,         &quot;debit_amount&quot; : 50000,         &quot;credit_amount&quot; : 20000,         &quot;closing_balance&quot; : 130000,         &quot;composition_ratio&quot; : 0.25         &quot;partners&quot; : [{           &quot;id&quot; : 123,           &quot;name&quot; : &quot;freee&quot;,           &quot;opening_balance&quot; : 100000,           &quot;debit_amount&quot; : 50000,           &quot;credit_amount&quot; : 20000,           &quot;closing_balance&quot; : 130000,           &quot;composition_ratio&quot; : 0.25           },         ...         ]       },       ...       ]     } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetTrialPlOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrialPlOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fiscalYear** | **optional.Int32**| 会計年度 | 
 **startMonth** | **optional.Int32**| 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **optional.Int32**| 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **optional.String**| 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **optional.String**| 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **optional.String**| 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **optional.String**| 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **optional.Int32**| 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **optional.String**| 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **optional.Int32**| 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **optional.Int32**| 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **optional.String**| 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **costAllocation** | **optional.String**| 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） | 

### Return type

[**TrialPlResponse**](trialPlResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlSections

> TrialPlSectionsResponse GetTrialPlSections(ctx, companyId, sectionIds, optional)

損益計算書(部門比較)の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の損益計算書(部門比較)を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul>  <li> <p>created_at : 作成日時</p> </li>  <li> <p>account_item_name : 勘定科目名</p> </li>  <li> <p>hierarchy_level: 階層レベル</p> </li>  <li> <p>parent_account_category_name: 上位勘定科目カテゴリー名</p> </li> <li> <p>closing_balance : 期末残高 </p> </li> <h2 id=\"_3\">注意点</h2> <ul> <li>個人向けのプレミアムプラン、法人向けのビジネスプラン以上で利用可能なAPIです。対象外のプランでは401エラーを返却します。</li> <li>会計年度が指定されない場合、現在の会計年度がデフォルトとなります。</li> <li>絞り込み条件の日付と、月または年度は同時に指定することはできません。</li> <li>up_to_dateがfalseの場合、残高の集計が完了していません。最新の集計結果を確認したい場合は、時間を空けて再度取得する必要があります。</li> <li>配賦仕訳の絞り込み（cost_allocation）は法人向けのベーシックプラン以上で利用可能です。</li> </ul> <h2 id=\"_4\">レスポンスの例</h2>  <blockquote> <p>GET https://api.freee.co.jp/api/1/reports/trial_pl_section?company_id=1&amp;section_ids=1,2,3&amp;fiscal_year=2017</p></p> </blockquote>  <pre><code>{   &quot;trial_pl_sections&quot; :     {       &quot;company_id&quot; : 1,       &quot;section_ids&quot; : &quot;1,2,3&quot;,       &quot;fiscal_year&quot; : 2017,       &quot;created_at&quot; : &quot;2018-05-01 12:00:50&quot       &quot;balances&quot; : [{         &quot;account_item_id&quot; : 1500,         &quot;account_item_name&quot; : &quot;売上高&quot;,         &quot;hierarchy_level&quot; : 2,         &quot;account_category_name&quot; : &quot;営業収益&quot;,         &quot;closing_balance&quot; : 1000000,         &quot;sections&quot; : [{           &quot;id&quot;: 1           &quot;name&quot;: &quot;営業部&quot;,           &quot;closing_balance&quot; : 100000         },         {           &quot;id&quot;: 2           &quot;name&quot;: &quot;広報部&quot;,           &quot;closing_balance&quot; : 200000         },         {           &quot;id&quot;: 3           &quot;name&quot;: &quot;人事部&quot;,           &quot;closing_balance&quot; : 300000         },         ...         ]       },       ...       ]     } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**sectionIds** | **string**| 出力する部門の指定（半角数字のidを半角カンマ区切りスペースなしで指定してください） | 
 **optional** | ***GetTrialPlSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrialPlSectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fiscalYear** | **optional.Int32**| 会計年度 | 
 **startMonth** | **optional.Int32**| 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **optional.Int32**| 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **optional.String**| 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **optional.String**| 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **optional.String**| 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **optional.String**| 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **optional.Int32**| 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **optional.String**| 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **optional.Int32**| 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **optional.String**| 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **costAllocation** | **optional.String**| 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） | 

### Return type

[**TrialPlSectionsResponse**](trialPlSectionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlThreeYears

> TrialPlThreeYearsResponse GetTrialPlThreeYears(ctx, companyId, optional)

損益計算書(３期間比較)の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の損益計算書(３期間比較)を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul>  <li> <p>created_at : 作成日時</p> </li>  <li> <p>account_item_name : 勘定科目名</p> </li>  <li> <p>hierarchy_level: 階層レベル</p> </li>  <li> <p>parent_account_category_name: 上位勘定科目カテゴリー名</p> </li> <li> <p>two_years_before_closing_balance:  前々年度期末残高 </p> </li> <li> <p>last_year_closing_balance:  前年度期末残高 </p> </li> <li> <p>closing_balance : 期末残高 </p> </li> <li> <p>year_on_year : 前年比</p> </li> <h2 id=\"_3\">注意点</h2> <ul> <li>会計年度が指定されない場合、現在の会計年度がデフォルトとなります。</li> <li>絞り込み条件の日付と、月または年度は同時に指定することはできません。</li> <li>up_to_dateがfalseの場合、残高の集計が完了していません。最新の集計結果を確認したい場合は、時間を空けて再度取得する必要があります。</li> <li>配賦仕訳の絞り込み（cost_allocation）は法人向けのベーシックプラン以上で利用可能です。</li> </ul> <h2 id=\"_4\">レスポンスの例</h2>  <blockquote> <p>GET https://api.freee.co.jp/api/1/reports/trial_pl_three_years?company_id=1&fiscal_year=2017</p> </blockquote>  <pre><code>{   &quot;trial_pl_three_years&quot; :     {       &quot;company_id&quot; : 1,       &quot;fiscal_year&quot; : 2017,       &quot;created_at&quot; : &quot;2018-05-01 12:00:50&quot       &quot;balances&quot; : [{         &quot;account_item_id&quot; : 1500,         &quot;account_item_name&quot; : &quot;売上高&quot;,         &quot;hierarchy_level&quot; : 2,         &quot;account_category_name&quot; : &quot;営業収益&quot;,         &quot;two_year_before_closing_balance&quot; : 50000,         &quot;last_year_closing_balance&quot; : 25000,         &quot;closing_balance&quot; : 100000,         &quot;year_on_year&quot; : 0.85       },       ...       ]     } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetTrialPlThreeYearsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrialPlThreeYearsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fiscalYear** | **optional.Int32**| 会計年度 | 
 **startMonth** | **optional.Int32**| 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **optional.Int32**| 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **optional.String**| 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **optional.String**| 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **optional.String**| 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **optional.String**| 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **optional.Int32**| 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **optional.String**| 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **optional.Int32**| 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **optional.Int32**| 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **optional.String**| 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **costAllocation** | **optional.String**| 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） | 

### Return type

[**TrialPlThreeYearsResponse**](trialPlThreeYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlTwoYears

> TrialPlTwoYearsResponse GetTrialPlTwoYears(ctx, companyId, optional)

損益計算書(前年比較)の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の損益計算書(前年比較)を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul>  <li> <p>created_at : 作成日時</p> </li>  <li> <p>account_item_name : 勘定科目名</p> </li>  <li> <p>hierarchy_level: 階層レベル</p> </li>  <li> <p>parent_account_category_name: 上位勘定科目カテゴリー名</p> </li> <li> <p>last_year_closing_balance:  前年度期末残高 </p> </li> <li> <p>closing_balance : 期末残高 </p> </li> <li> <p>year_on_year : 前年比</p> </li> <h2 id=\"_3\">注意点</h2> <ul> <li>会計年度が指定されない場合、現在の会計年度がデフォルトとなります。</li> <li>絞り込み条件の日付と、月または年度は同時に指定することはできません。</li> <li>up_to_dateがfalseの場合、残高の集計が完了していません。最新の集計結果を確認したい場合は、時間を空けて再度取得する必要があります。</li> <li>配賦仕訳の絞り込み（cost_allocation）は法人向けのベーシックプラン以上で利用可能です。</li> </ul>  <h2 id=\"_4\">レスポンスの例</h2>  <blockquote> <p>GET https://api.freee.co.jp/api/1/reports/trial_pl_two_years?company_id=1&amp;fiscal_year=2017</p> </blockquote>  <pre><code>{   &quot;trial_pl_two_years&quot; :     {       &quot;company_id&quot; : 1,       &quot;fiscal_year&quot; : 2017,       &quot;created_at&quot; : &quot;2018-05-01 12:00:50&quot       &quot;balances&quot; : [{         &quot;account_item_id&quot; : 1500,         &quot;account_item_name&quot; : &quot;売上高&quot;,         &quot;hierarchy_level&quot; : 2,         &quot;account_category_name&quot; : &quot;営業収益&quot;,         &quot;last_year_closing_balance&quot; : 25000,         &quot;closing_balance&quot; : 100000,         &quot;year_on_year&quot; : 0.85        },       ...       ]     } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetTrialPlTwoYearsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrialPlTwoYearsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fiscalYear** | **optional.Int32**| 会計年度 | 
 **startMonth** | **optional.Int32**| 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **optional.Int32**| 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **optional.String**| 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **optional.String**| 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **optional.String**| 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **optional.String**| 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **optional.Int32**| 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **optional.String**| 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **optional.Int32**| 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **optional.Int32**| 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **optional.String**| 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **costAllocation** | **optional.String**| 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） | 

### Return type

[**TrialPlTwoYearsResponse**](trialPlTwoYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


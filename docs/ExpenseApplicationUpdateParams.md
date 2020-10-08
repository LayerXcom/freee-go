# ExpenseApplicationUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Title** | **string** | 申請タイトル (250文字以内) | 
**IssueDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Description** | **string** | 備考 (10000文字以内) | [optional] 
**EditableOnWeb** | **bool** | 会計freeeのWeb画面で申請内容を編集可能にするかどうか&lt;br&gt; falseの場合は、Web画面での項目行の追加／削除・金額の編集が出来なくなりますが、APIでの編集は可能です。&lt;br&gt; デフォルトはfalseです。  | [optional] 
**SectionId** | **int32** | 部門ID | [optional] 
**TagIds** | **[]int32** | メモタグID | [optional] 
**ExpenseApplicationLines** | [**[]ExpenseApplicationUpdateParamsExpenseApplicationLines**](expenseApplicationUpdateParams_expense_application_lines.md) |  | 
**ApprovalFlowRouteId** | **int32** | 申請経路ID&lt;br&gt; &lt;ul&gt;     &lt;li&gt;経費申請のステータスを申請中として作成する場合は、必ず指定してください。&lt;/li&gt;     &lt;li&gt;指定する申請経路IDは、申請経路APIを利用して取得してください。&lt;/li&gt;     &lt;li&gt;         未指定の場合は、基本経路を設定している事業所では基本経路が、基本経路を設定していない事業所では利用可能な申請経路の中から最初の申請経路が自動的に使用されます。         &lt;ul&gt;            &lt;li&gt;意図しない申請経路を持った経費申請の作成を防ぐために、使用する申請経路IDを指定することを推奨します。&lt;/li&gt;         &lt;/ul&gt;     &lt;/li&gt;     &lt;li&gt;         ベーシックプランの事業所では以下のデフォルトで用意された申請経路のみ指定できます         &lt;ul&gt;         &lt;li&gt;指定なし&lt;/li&gt;         &lt;li&gt;承認者を指定&lt;/li&gt;         &lt;/ul&gt;     &lt;/li&gt; &lt;/ul&gt;  | [optional] 
**ApproverId** | **int32** | 承認者のユーザーID&lt;br&gt; 指定する承認者のユーザーIDは、申請経路APIを利用して取得してください。  | [optional] 
**Draft** | **bool** | 経費申請のステータス&lt;br&gt; falseを指定した時は申請中（in_progress）で経費申請を更新します。&lt;br&gt; trueを指定した時は下書き（draft）で経費申請を更新します。&lt;br&gt; 未指定の時は下書きとみなして経費申請を更新します。  | [optional] 
**Segment1TagId** | **int32** | セグメント１ID(法人向けプロフェッショナル, 法人向けエンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment2TagId** | **int32** | セグメント２ID(法人向け エンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment3TagId** | **int32** | セグメント３ID(法人向け エンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



package freee

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathCompanies = "companies"
)

type CompanyResponse struct {
	Company Company `json:"company"`
}

type Company struct {
	// 事業所ID
	ID int64 `json:"id"`
	// 事業所の正式名称 (100文字以内)
	Name *string `json:"name"`
	// 正式名称フリガナ (100文字以内)
	NameKana *string `json:"name_kana"`
	// 事業所名
	DisplayName string `json:"display_name"`
	// 源泉徴収税計算（0: 消費税を含める、1: 消費税を含めない）
	TaxAtSourceCalcType int64 `json:"tax_at_source_calc_type"`
	// 担当者名 (50文字以内)
	ContactName *string `json:"contact_name"`
	// 従業員数（0: 経営者のみ、1: 2~5人、2: 6~10人、3: 11~20人、4: 21~30人、5: 31~40人、6: 41~100人、7: 100人以上
	HeadCount *int64 `json:"head_count"`
	// 法人番号 (半角数字13桁、法人のみ)
	CorporateNumber string `json:"corporate_number"`
	// 仕訳番号形式（not_used: 使用しない、digits: 数字（例：5091824）、alnum: 英数字（例：59J0P））
	TxnNumberFormat string `json:"txn_number_format"`
	// 決済口座のデフォルト
	DefaultWalletAccountId *int64 `json:"default_wallet_account_id,omitempty"`
	// プライベート資金/役員資金（false: 使用しない、true: 使用する）
	PrivateSettlement bool `json:"private_settlement"`
	// マイナスの表示方法（0: -、 1: △）
	MinusFormat int64 `json:"minus_format"`
	// ユーザーの権限
	Role string `json:"role"`
	// 電話番号１
	Phone1 string `json:"phone1"`
	// 電話番号２
	Phone2 *string `json:"phone2"`
	// FAX
	Fax *string `json:"fax"`
	// 郵便番号
	Zipcode string `json:"zipcode"`
	// 都道府県コード（-1: 設定しない、0: 北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	PrefectureCode *int64 `json:"prefecture_code"`
	// 市区町村・番地
	StreetName1 string `json:"street_name1"`
	// 建物名・部屋番号など
	StreetName2 string `json:"street_name2"`
	// 請求書レイアウト * `default_classic` - レイアウト１/クラシック (デフォルト)  * `standard_classic` - レイアウト２/クラシック  * `envelope_classic` - 封筒１/クラシック  * `carried_forward_standard_classic` - レイアウト３（繰越金額欄あり）/クラシック  * `carried_forward_envelope_classic` - 封筒２（繰越金額欄あり）/クラシック  * `default_modern` - レイアウト１/モダン  * `standard_modern` - レイアウト２/モダン  * `envelope_modern` - 封筒/モダン
	InvoiceLayout string `json:"invoice_layout"`
	// 金額端数処理方法（0: 切り捨て、1: 切り上げ、2: 四捨五入）
	AmountFraction int64 `json:"amount_fraction"`
	// 種別（agriculture_forestry_fisheries_ore: 農林水産業/鉱業、construction: 建設、manufacturing_processing: 製造/加工、it: IT、transportation_logistics: 運輸/物流、retail_wholesale: 小売/卸売、finance_insurance: 金融/保険、real_estate_rental: 不動産/レンタル、profession: 士業/学術/専門技術サービス、design_production: デザイン/制作、food: 飲食、leisure_entertainment: レジャー/娯楽、lifestyle: 生活関連サービス、education: 教育/学習支援、medical_welfare: 医療/福祉、other_services: その他サービス、other: その他）
	IndustryClass string `json:"industry_class"`
	// 業種（agriculture: 農業, forestry: 林業, fishing_industry: 漁業、水産養殖業, mining: 鉱業、採石業、砂利採取業, civil_contractors: 土木工事業, pavement: 舗装工事業, carpenter: とび、大工、左官等の建設工事業, renovation: リフォーム工事業, electrical_plumbing: 電気、管工事等の設備工事業, grocery: 食料品の製造加工業, machinery_manufacturing: 機械器具の製造加工業, printing: 印刷業, other_manufacturing: その他の製造加工業, software_development: 受託：ソフトウェア、アプリ開発業, system_development: 受託：システム開発業, survey_analysis: 受託：調査、分析等の情報処理業, server_management: 受託：サーバー運営管理, website_production: 受託：ウェブサイト制作, online_service_management: オンラインサービス運営業, online_advertising_agency: オンライン広告代理店業, online_advertising_planning_production: オンライン広告企画・制作業, online_media_management: オンラインメディア運営業, portal_site_management: ポータルサイト運営業, other_it_services: その他、IT サービス業, transport_delivery: 輸送業、配送業, delivery: バイク便等の配達業, other_transportation_logistics: その他の運輸業、物流業, other_wholesale: 卸売業：その他, clothing_wholesale_fiber: 卸売業：衣類卸売／繊維, food_wholesale: 卸売業：飲食料品, entrusted_development_wholesale: 卸売業：機械器具, online_shop: 小売業：無店舗　オンラインショップ, fashion_grocery_store: 小売業：店舗あり　ファッション、雑貨, food_store: 小売業：店舗あり　生鮮食品、飲食料品, entrusted_store: 小売業：店舗あり　機械、器具, other_store: 小売業：店舗あり　その他, financial_instruments_exchange: 金融業：金融商品取引, commodity_futures_investment_advisor: 金融業：商品先物取引、商品投資顧問, other_financial: 金融業：その他, brokerage_insurance: 保険業：仲介、代理, other_insurance: 保険業：その他, real_estate_developer: 不動産業：ディベロッパー, real_estate_brokerage: 不動産業：売買、仲介, rent_coin_parking_management: 不動産業：賃貸、コインパーキング、管理, rental_office_co_working_space: 不動産業：レンタルオフィス、コワーキングスペース, rental_lease: レンタル業、リース業, cpa_tax_accountant: 士業：公認会計士事務所、税理士事務所, law_office: 士業：法律事務所, judicial_and_administrative_scrivener: 士業：司法書士事務所／行政書士事務所, labor_consultant: 士業：社会保険労務士事務所, other_profession: 士業：その他, business_consultant: 経営コンサルタント, academic_research_development: 学術・開発研究機関, advertising_agency: 広告代理店, advertising_planning_production: 広告企画／制作, design_development: ソフトウェア、アプリ開発業（受託）, apparel_industry_design: 服飾デザイン業、工業デザイン業, website_design: ウェブサイト制作（受託）, advertising_planning_design: 広告企画／制作業, other_design: その他、デザイン／制作, restaurants_coffee_shops: レストラン、喫茶店等の飲食店業, sale_of_lunch: 弁当の販売業, bread_confectionery_manufacture_sale: パン、菓子等の製造販売業, delivery_catering_mobile_catering: デリバリー業、ケータリング業、移動販売業, hotel_inn: 宿泊業：ホテル、旅館, homestay: 宿泊業：民泊, travel_agency: 旅行代理店業, leisure_sports_facility_management: レジャー、スポーツ等の施設運営業, show_event_management: ショー、イベント等の興行、イベント運営業, barber: ビューティ、ヘルスケア業：床屋、理容室, beauty_salon: ビューティ、ヘルスケア業：美容室, spa_sand_bath_sauna: ビューティ、ヘルスケア業：スパ、砂風呂、サウナ等, este_ail_salon: ビューティ、ヘルスケア業：その他、エステサロン、ネイルサロン等, bridal_planning_introduce_wedding: 冠婚葬祭業：ブライダルプランニング、結婚式場紹介等, memorial_ceremony_funeral: 冠婚葬祭業：メモリアルセレモニー、葬儀等, moving: 引っ越し業, courier_industry: 宅配業, house_maid_cleaning_agency: 家事代行サービス業：無店舗　ハウスメイド、掃除代行等, re_tailoring_clothes: 家事代行サービス業：店舗あり　衣類修理、衣類仕立て直し等, training_institute_management: 研修所等の施設運営業, tutoring_school: 学習塾、進学塾等の教育・学習支援業, music_calligraphy_abacus_classroom: 音楽教室、書道教室、そろばん教室等のの教育・学習支援業, english_school: 英会話スクール等の語学学習支援業, tennis_yoga_judo_school: テニススクール、ヨガ教室、柔道場等のスポーツ指導、支援業, culture_school: その他、カルチャースクール等の教育・学習支援業, seminar_planning_management: セミナー等の企画、運営業, hospital_clinic: 医療業：病院、一般診療所、クリニック等, dental_clinic: 医療業：歯科診療所, other_medical_services: 医療業：その他、医療サービス等, nursery: 福祉業：保育所等、児童向け施設型サービス, nursing_home: 福祉業：老人ホーム等、老人向け施設型サービス, rehabilitation_support_services: 福祉業：療育支援サービス等、障害者等向け施設型サービス, other_welfare: 福祉業：その他、施設型福祉サービス, visit_welfare_service: 福祉業：訪問型福祉サービス, recruitment_temporary_staffing: 人材紹介業、人材派遣業, life_related_recruitment_temporary_staffing: 生活関連サービスの人材紹介業、人材派遣業, car_maintenance_car_repair: 自動車整備業、自動車修理業, machinery_equipment_maintenance_repair: 機械機器類の整備業、修理業, cleaning_maintenance_building_management: 清掃業、メンテナンス業、建物管理業, security: 警備業, other_services: その他のサービス業, npo: NPO, general_incorporated_association: 一般社団法人, general_incorporated_foundation: 一般財団法人, other_association: その他組織)
	IndustryCode string `json:"industry_code"`
	// 仕訳承認フロー（enable: 有効、 disable: 無効）
	WorkflowSetting string `json:"workflow_setting"`
	// 取引先コードの利用設定（true: 有効、 false: 無効）
	UsePartnerCode bool          `json:"use_partner_code"`
	FiscalYears    []FiscalYears `json:"fiscal_years"`
}

// FiscalYears struct for FiscalYears
type FiscalYears struct {
	// 製造業向け機能（true: 使用する、false: 使用しない）
	UseIndustryTemplate bool `json:"use_industry_template"`
	// 固定資産の控除法（true: 間接控除法、false: 直接控除法）
	IndirectWriteOffMethod bool `json:"indirect_write_off_method"`
	// 期首日
	StartDate *string `json:"start_date,omitempty"`
	// 期末日
	EndDate *string `json:"end_date,omitempty"`
	// 月次償却（0: しない、1: する）
	DepreciationRecordMethod int64 `json:"depreciation_record_method"`
	// 課税区分（0: 免税、1: 簡易課税、2: 本則課税（個別対応方式）、3: 本則課税（一括比例配分方式）、4: 本則課税（全額控除））
	TaxMethod int64 `json:"tax_method"`
	// 簡易課税用事業区分（0: 第一種：卸売業、1: 第二種：小売業、2: 第三種：農林水産業、工業、建設業、製造業など、3: 第四種：飲食店業など、4: 第五種：金融・保険業、運輸通信業、サービス業など、5: 第六種：不動産業など
	SalesTaxBusinessCode int64 `json:"sales_tax_business_code"`
	// 消費税端数処理方法（0: 切り捨て、1: 切り上げ、2: 四捨五入）
	TaxFraction int64 `json:"tax_fraction"`
	// 消費税経理処理方法（0: 税込経理、1: 旧税抜経理、2: 税抜経理）
	TaxAccountMethod int64 `json:"tax_account_method"`
	// 不動産所得使用区分（0: 一般、3: 一般/不動産） ※個人事業主のみ設定可能
	ReturnCode int64 `json:"return_code"`
}

type GetCompanyOpts struct {
	Details      *bool
	AccountItems *bool
	Taxes        *bool
	Items        *bool
	Partners     *bool
	Sections     *bool
	Tags         *bool
	Walletables  *bool
}

func (c *Client) GetCompany(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, opts GetCompanyOpts,
) (*CompanyResponse, *oauth2.Token, error) {
	var result CompanyResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	oauth2Token, err = c.call(ctx, path.Join(APIPathCompanies, fmt.Sprint(companyID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}

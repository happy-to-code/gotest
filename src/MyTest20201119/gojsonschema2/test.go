package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	// jsonStr := `{"header":{"model":{"protocol":"区域性股权市场跨链业务数据模型","version":"2.0.0-alpha3"},"content":{"type":"subject","object_id":"OID123457","version":0,"operation":"create","timestamp":1606271676}},"body":{"subject":{"basic_information_subject":{"general_information_subject":{"subject_create_time":"2020-10-1218:18:18","subject_id":"OID7890123","subject_type":1,"subject_main_administrative_region":1},"subject_qualification_information":[]},"organization_subject_information":{"basic_information_of_enterprise":{"basic_information_description":{"subject_association_articles":[],"subject_business_license":[{"file_number":"1","file_name":"12312312","url":"12312312","hash":"12312312","summary":"12312312","term_of_validity_type":"0","term_of_validity":"yyyy-MM-dd"}],"subject_document_information":[{"type":1,"code":"12312312"},{"type":2,"code":"12312312"}],"subject_business_scope":"scopte","subject_company_business":"软件","subject_company_english_name":"reerea","subject_company_name":"tongji","subject_company_profile":"dddddd","subject_company_short_english_name":"jjkkkk","subject_company_short_name":"erewr","subject_company_type":0,"subject_contact_address":"苏州相城","subject_contact_number":"333","subject_industry":0,"subject_internet_address":"www","subject_mail_box":"zz","subject_office_address":"苏州相城","subject_organization_nature":0,"subject_paid_in_capital":156,"subject_paid_in_capital_currency":"156","subject_postal_code":"22222","subject_registered_address":"苏州","subject_registered_capital":1000000,"subject_registered_capital_currency":"156","subject_shareholders_number":10},"leading_member_information":[]}}},"subject_object_information":{"subject_object_id":"OID123457","subject_type":0}}}`
	jsonStr := `{
    "body": {
        "product_information": {
            "escrow_information": {},
            "essential_information": {
                "basic_product_information": {
                    "product_account_number_max": 200,
                    "product_code": "202555",
                    "product_create_time": "2015-03-18T00:00:00+08:00",
                    "product_customer_browsing_right": 0,
                    "product_info_disclosure_way": 0,
                    "product_issuer_name": "北京博信世纪科技有限公司",
                    "product_issuer_subject_ref": "4B9F6F282EF8FD04443C023C0725420C/0",
                    "product_market_subject_name": "上海股权托管交易中心股份有限公司",
                    "product_market_subject_ref": "gdTokenZ0SETDXj111111111111111111111",
                    "product_name": "北京博信世纪科技有限公司",
                    "product_name_abbreviation": "博信世纪",
                    "product_plate_trading_name": "基础有限",
                    "product_scale": 1000000,
                    "product_scale_unit": 0,
                    "product_trading_market_category": 1,
                    "product_type": 1,
                    "product_type_function": 255
                },
                "product_file_information": [],
                "service_provider_information": []
            },
            "product_subject_information": {
                "business_information": {},
                "fund_use_information": {},
                "portfolio_information": {}
            },
            "release_information": {
                "equity_issue_information": {
                    "product_shares_issued_class": []
                },
                "filing_information": []
            },
            "transaction_information": {
                "delisting_information": {
                    "product_delisting_reason": 255
                },
                "listing_information": {
                    "product_listing_date": "2015-03-18"
                },
                "transaction_status": {}
            }
        }
    },
    "header": {
        "content": {
            "object_id": "A8F88705C3DFB730B4FCE5A6B616CDAA",
            "operation": "create",
            "timestamp": 1425168000,
            "type": "product",
            "version": 0
        },
        "metadata": null,
        "model": {
            "protocol": "区域性股权市场跨链业务数据模型",
            "version": "2.0.0"
        }
    }
}`

	var s = "gdTokenZ0SETDXj11111111111111111111111111111111"
	fmt.Println("s len :", len(s))
	validJsonSchema(jsonStr)

}

var schemaLoader = gojsonschema.NewReferenceLoader("file://E:/20.06.16Project/GoTest/src/MyTest20201119/gojsonschema2/数据对象Schema-v2.0.0.json")

func validJsonSchema(jsonStr string) {
	documentLoader := gojsonschema.NewStringLoader(jsonStr)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid:%+v\n", result.Valid())
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			// fmt.Errorf("%s", desc)
			fmt.Printf("- %s\n", desc)
		}
	}
	fmt.Println("00000000000000000000")
}

package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

var m map[string]gojsonschema.JSONLoader

func main() {
	jsonStr := `{"header":{"model":{"protocol":"区域性股权市场跨链业务数据模型","version":"2.0.0-alpha3"},"content":{"type":"subject","object_id":"OID123457","version":0,"operation":"create","timestamp":1606271676}},"body":{"subject":{"basic_information_subject":{"general_information_subject":{"subject_create_time":"2020-10-1218:18:18","subject_id":"OID7890123","subject_type":1,"subject_main_administrative_region":1},"subject_qualification_information":[]},"organization_subject_information":{"basic_information_of_enterprise":{"basic_information_description":{"subject_association_articles":[],"subject_business_license":[{"file_number":"1","file_name":"12312312","url":"12312312","hash":"12312312","summary":"12312312","term_of_validity_type":"0","term_of_validity":"yyyy-MM-dd"}],"subject_document_information":[{"type":1,"code":"12312312"},{"type":2,"code":"12312312"}],"subject_business_scope":"scopte","subject_company_business":"软件","subject_company_english_name":"reerea","subject_company_name":"tongji","subject_company_profile":"dddddd","subject_company_short_english_name":"jjkkkk","subject_company_short_name":"erewr","subject_company_type":0,"subject_contact_address":"苏州相城","subject_contact_number":"333","subject_industry":0,"subject_internet_address":"www","subject_mail_box":"zz","subject_office_address":"苏州相城","subject_organization_nature":0,"subject_paid_in_capital":156,"subject_paid_in_capital_currency":"156","subject_postal_code":"22222","subject_registered_address":"苏州","subject_registered_capital":1000000,"subject_registered_capital_currency":"156","subject_shareholders_number":10},"leading_member_information":[]}}},"subject_object_information":{"subject_object_id":"OID123457","subject_type":0}}}`

	validJsonSchema(jsonStr)
}

func validJsonSchema(jsonStr string) {

	schemaLoader := gojsonschema.NewReferenceLoader("file://E:/20.06.16Project/GoTest/src/MyTest20201119/gojsonschema/Schema-v2.0.0-Alpha3.1.json")

	// documentLoader := gojsonschema.NewReferenceLoader("file://E:/20.06.16Project/GoTest/src/MyTest20201119/gojsonschema/test.json")

	for i := 0; i < 5; i++ {
		documentLoader := gojsonschema.NewStringLoader(jsonStr)

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			panic(err.Error())
		}

		if result.Valid() {
			fmt.Printf("The document is valid\n")
		} else {
			fmt.Printf("The document is not valid. see errors :\n")
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}
		}
	}
}

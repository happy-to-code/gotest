package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var s = `{
  "body": {
    "subject_information": {
      "basic_information_subject": {
        "general_information_subject": {
          "subject_create_time": "2024-12-16T13:50:00+08:00",
          "subject_id": "gdCmpyId00018KxvJa",
          "subject_main_administrative_region": 0,
          "subject_type": 1
        },
        "subject_qualification_information": [
          {
            "qualification_information": {
              "subject_market_roles_type": 1
            },
            "investor_suitability_information": [
              {
                "subject_investor_qualification": 1,
                "subject_investor_qualification_sub": "aa",
                "subject_investor_qualification_certifier_ref": "ref1/0"
              }
            ]
          },
          {
            "qualification_information": {
              "subject_market_roles_type": 2
            },
            "investor_suitability_information": [
              {
                "subject_investor_qualification": 2,
                "subject_investor_qualification_sub": "bb",
                "subject_investor_qualification_certifier_ref": "ref2/1"
              }
            ]
          }
        ]
      },
      "organization_subject_information": {
        "basic_information_of_enterprise": {
          "basic_information_description": {
            "subject_approval_time": "2020-12-16T12:00:25+08:00",
            "subject_association_articles": [
              {
                "file_name": "1file1.pdf",
                "file_number": "2",
                "hash": "da1234filehash52221",
                "summary": "简述1",
                "term_of_validity": "2020/04/18",
                "term_of_validity_type": "0",
                "url": "http://test.com/file/201/1file1.pdf"
              },
              {
                "file_name": "2file1.pdf",
                "file_number": "2",
                "hash": "da1234filehash52222",
                "summary": "简述2",
                "term_of_validity": "2020/04/18",
                "term_of_validity_type": "0",
                "url": "http://test.com/file/201/2file1.pdf"
              }
            ],
            "subject_business_license": [
              {
                "file_name": "1file1.pdf",
                "file_number": "2",
                "hash": "da1234filehash52221",
                "summary": "简述1",
                "term_of_validity": "2020/04/18",
                "term_of_validity_type": "0",
                "url": "http://test.com/file/201/1file1.pdf"
              },
              {
                "file_name": "2file1.pdf",
                "file_number": "2",
                "hash": "da1234filehash52222",
                "summary": "简述2",
                "term_of_validity": "2020/04/18",
                "term_of_validity_type": "0",
                "url": "http://test.com/file/201/2file1.pdf"
              }
            ],
            "subject_business_scope": "经营范围CHARACTER",
            "subject_city": "注册地所在市CHARACTER",
            "subject_company_business": "主营业务CHARACTER",
            "subject_company_english_name": "英文名称CHARACTERYYYYYYY",
            "subject_company_name": "公司全称CHARACTER",
            "subject_company_profile": "公司简介TEXT",
            "subject_company_short_english_name": "英文简称CHARACTER",
            "subject_company_short_name": "公司简称CHARACTER",
            "subject_company_status": 0,
            "subject_company_status_deregistration": "注销原因CHARACTER",
            "subject_company_status_deregistration_date": "2029-11-10",
            "subject_company_status_windingup": "吊销原因CHARACTER",
            "subject_company_status_windingup_date": "2029-11-10",
            "subject_company_type": 0,
            "subject_contact_address": "联系地址CHARACTER",
            "subject_contact_number": "联系电话CHARACTER",
            "subject_district": "注册地所在区CHARACTER",
            "subject_document_information": [
              {
                "code": "1233333333",
                "type": 1
              }
            ],
            "subject_economic_type": 0,
            "subject_fax": "企业传真CHARACTER",
            "subject_high_technology_enterprise": 0,
            "subject_industry": 1,
            "subject_insured_number": 50,
            "subject_internet_address": "互联网地址CHARACTER",
            "subject_invoice_account_number": "发票账号CHARACTER",
            "subject_invoice_address": "发票地址CHARACTER",
            "subject_invoice_bank": "发票开户行CHARACTER",
            "subject_invoice_telephone_number": "发票电话CHARACTER",
            "subject_legal_type": 0,
            "subject_mail_box": "电子邮箱CHARACTER",
            "subject_name_used_before": [
              "曾用名a",
              "曾用名b",
              "曾用名c"
            ],
            "subject_office_address": "办公地址CHARACTER",
            "subject_organization_nature": 0,
            "subject_paid_in_capital": 6000000,
            "subject_paid_in_capital_currency": "156",
            "subject_personnel_size": "人员规模CHARACTER",
            "subject_postal_code": "邮政编码CHARACTER",
            "subject_province": "注册地所在省份CHARACTER",
            "subject_registered_address": "注册地址CHARACTER",
            "subject_registered_capital": 6000000,
            "subject_registered_capital_currency": "156",
            "subject_registry_date": "2018-10-23",
            "subject_regulator": "主管单位CHARACTER",
            "subject_scale_type": 0,
            "subject_shareholders_number": 10,
            "subject_taxpayer_id_number": "纳税人识别号CHARACTER"
          },
          "leading_member_information": [
            {
              "subject_document_type": 0,
              "subject_key_personnel_address": "证件地址CHARACTER",
              "subject_key_personnel_appointment_end": "2019-12-21",
              "subject_key_personnel_appointment_start": "2019-12-21",
              "subject_key_personnel_contact": "联系方式CHARACTER",
              "subject_key_personnel_id": "证件号码CHARACTER",
              "subject_key_personnel_name": "姓名CHARACTER",
              "subject_key_personnel_nationality": "国籍CHARACTER",
              "subject_key_personnel_position": 0,
              "subject_key_personnel_shareholding": 500,
              "subject_key_personnel_shareholding_ratio": 20,
              "subject_key_personnel_type": 0
            }
          ]
        }
      }
    }
  },
  "header": {
    "content": {
      "object_id": "TCPR001",
      "operation": "update",
      "timestamp": 1607501677,
      "type": "subject",
      "version": 1
    },
    "metadata": null,
    "model": {
      "protocol": "区域性股权市场跨链业务数据模型",
      "version": "2.0.0-alpha4"
    }
  }
}`
	bk := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(s), &bk)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v\n", bk)
	// success := updateMapTree(bk, "header.model.protocol", "--区域性--")
	// fmt.Println(success)
	// fmt.Printf("%v\n", bk)
	valueFromMapTree, b := getValueFromMapTree(bk, "body.subject_information.organization_subject_information.basic_information_of_enterprise.basic_information_description.subject_shareholders_number")
	if b {
		fmt.Println("value:", valueFromMapTree)
	} else {
		fmt.Println("#############")
	}
	success := updateMapTree(bk, "body.subject_information.organization_subject_information.basic_information_of_enterprise.basic_information_description.subject_shareholders_number",
		int(valueFromMapTree.(float64))+2)
	fmt.Println(success)
	// fmt.Printf("%v\n", bk)
	subjectQualificationInfoObject, b := getValueFromMapTree(bk, "body.subject_information.basic_information_subject.subject_qualification_information")
	if b {
		fmt.Println("subjectQualificationInfoObject:", subjectQualificationInfoObject)
		subjectQualificationInfoList, ok := subjectQualificationInfoObject.([]interface{})
		if ok {
			for _, v := range subjectQualificationInfoList {
				fmt.Println("v:", v)
				subjectQualificationInfoMap, ok := v.(map[string]interface{})
				if ok {
					investorSuitabilityInfoObject, ok := subjectQualificationInfoMap["investor_suitability_information"]
					if ok {
						investorSuitabilityInfoList, ok := investorSuitabilityInfoObject.([]interface{})
						if ok {
							for _, investorSuitabilityInfo := range investorSuitabilityInfoList {
								fmt.Println(investorSuitabilityInfo)
								investorSuitabilityInfoMap, ok := investorSuitabilityInfo.(map[string]interface{})
								if ok {
									certifierRef, ok := investorSuitabilityInfoMap["subject_investor_qualification_certifier_ref"]
									if ok {
										certifierRefStr, ok := certifierRef.(string)
										if ok {
											investorSuitabilityInfoMap["subject_investor_qualification_certifier_ref"] = certifierRefStr + "KKKKKKKKKKKKKK"
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("%v\n", bk)
}

// 更新map中的值
func updateMapTree(m map[string]interface{}, path string, value interface{}) bool {
	segs := strings.Split(path, ".")
	for i, seg := range segs {
		if v, ok := m[seg]; ok {
			if i == len(segs)-1 {
				m[seg] = value
				return true
			} else {
				if m, ok = v.(map[string]interface{}); !ok {
					return false
				}
			}
		}
	}
	return false
}

// 从map 中获取值
func getValueFromMapTree(m map[string]interface{}, path string) (interface{}, bool) {
	segs := strings.Split(path, ".")
	for i, seg := range segs {
		if m != nil && len(m) > 0 {
			if v, ok := m[seg]; ok {
				if i == len(segs)-1 {
					return m[seg], true
				} else {
					if m, ok = v.(map[string]interface{}); !ok {
						return nil, false
					}
				}
			}
		}
	}
	return nil, false
}

package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strings"
)

type AccountObject struct {
	account_object_id                        string `json:"account_object_id"`                        // 账户对象标识
	Account_Object_Information_Type          uint32 `json:"account_object_information_type"`          // 对象信息类型
	Account_Holder_Subject_Ref               string `json:"account_holder_subject_ref"`               // 账户所属主体引用
	Account_Depository_Subject_Ref           string `json:"account_depository_subject_ref"`           // 开户机构主体引用
	Account_Number                           string `json:"account_number"`                           // 账号
	Account_Type                             uint32 `json:"account_type"`                             // 账户类型
	Account_Never                            uint32 `json:"account_never"`                            // 账户用途
	Account_Status                           uint32 `json:"account_status"`                           // 账户状态
	Account_Qualification_Certification_File string `json:"account_qualification_certification_file"` // 资质认证文件-
	Account_Certifier                        string `json:"account_certifier"`                        // 资质认证方
	Account_Auditor                          string `json:"account_auditor"`                          // 资质审核方
	Account_Certification_Time               string `json:"account_certification_time"`               // 认证时间
	Account_Audit_Time                       string `json:"account_audit_time"`                       // 审核时间
	Account_Opening_Date                     string `json:"account_opening_date"`                     // 账户开户时间
	Account_Opening_Certificate              string `json:"account_opening_certificate"`              // 账户开户核验凭证-
	Account_Closing_Date                     string `json:"account_closing_date"`                     // 账户销户时间
	Account_Closing_Certificate              string `json:"account_closing_certificate"`              // 账户销户核验凭证-
	Account_Forzen_Cate                      string `json:"account_forzen_date"`                      // 账户冻结时间
	Account_Forzen_Certificate               string `json:"account_forzen_certificate"`               // 账户冻结核验凭证-
	Account_Thaw_Date                        string `json:"account_thaw_date"`                        // 账户解冻时间
	Account_Thaw_Certificate                 string `json:"account_thaw_certificate"`                 // 账户解冻核验凭证-
	Account_Association                      uint32 `json:"account_association"`                      // 关联关系
	Account_Associated_Account_Ref           string `json:"account_associated_account_ref"`           // 关联账户对象引用
	Account_Associated_Acct_Certificates     string `json:"account_associated_acct_certificates"`     // 关联账户开户文件-
}

func main() {
	var m = make(map[string]interface{})
	m["account_object_id"] = "KKKKKKK"
	m["ACCOUNT_OBJECT_INFORMATION_TYPE"] = 1

	fmt.Printf("m：%+v\n", m)
	lowCase := ChangeMapKeyToLowCase(m)
	fmt.Printf("lowCase：%+v\n", lowCase)

	toStruct := MapToStruct(m)
	fmt.Printf("toStruct:%+v\n", toStruct)

	// var a AccountObject
	// bytes, _ := json.Marshal(m)
	// json.Unmarshal(bytes, &a)
	// fmt.Printf("a:%+v\n", a)
}

// map 转 AccountObject对象
func MapToStruct(mapInstance map[string]interface{}) (accountObject AccountObject) {
	var a AccountObject
	err := mapstructure.Decode(mapInstance, &a)
	if err != nil {
		fmt.Print("map转AccountObject对象出错,err:\n", err)
	}
	return a
}

// 将map的key转换成小写
func ChangeMapKeyToLowCase(paramMap map[string]interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})
	for k, v := range paramMap {
		m[strings.ToLower(k)] = v
	}
	return m
}

package main

import (
	"fmt"
	"reflect"
)

type AccountObject struct {
	Account_Object_Id                        string `json:"account_object_id"`                        // 账户对象标识
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
	var account = AccountObject{
		Account_Object_Id:                    "io213",
		Account_Thaw_Certificate:             "kkk",
		Account_Associated_Acct_Certificates: "kksldf",
	}
	fmt.Printf("account:%+v", account)

	fmt.Println("=========================================================================")
	fundAccountInfoMap := StructToMap(account)
	fmt.Printf("fundAccountInfoMap:%+v", fundAccountInfoMap)

}

// 结构体转map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

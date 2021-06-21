package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var minioClient *minio.Client
var ctx context.Context

func init() {
	var err error
	ctx = context.Background()
	// endpoint := "b2904236d6.zicp.vip:8098"
	endpoint := "121.224.59.167:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false
	// 初使化 minio client对象。
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	// 创建一个叫mymusic的存储桶。
	bucketName := "zyf-test"

	// file, err := os.Open("my-testfile")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()
	//
	// fileStat, err := file.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	/*str := `{
	    "header": {
	        "model": {
	            "protocol": "区域性股权市场跨链业务数据模型",
	            "version": "2.0.0-alpha4"
	        },
	        "content": {
	            "type": "subject",
	            "object_id": "OID123456",
	            "version": 4,
	            "operation": "create",
	            "timestamp": 1606984014
	        },
	        "metadata": null
	    },
	    "body": {
	        "subject_information": {
	            "basic_information_subject": {
	                "general_information_subject": {
	                    "subject_create_time": "2020-03-12 18:08:08",
	                    "subject_id": "id123456",
	                    "subject_main_administrative_region": 0,
	                    "subject_type": 0
	                },
	                "subject_qualification_information": [
	                    {
	                        "investor_suitability_information": [
	                            {
	                                "subject_investor_qualification": 0,
	                                "subject_investor_qualification_certificate": {
	                                    "file_name": "12312312",
	                                    "file_number": "1",
	                                    "hash": "12312312",
	                                    "summary": "12312312",
	                                    "term_of_validity": "yyyy/MM/dd",
	                                    "term_of_validity_type": "0",
	                                    "uri": "12312312"
	                                },
	                                "subject_investor_qualification_certificate_time": "2020-03-12 18:08:08",
	                                "subject_investor_qualification_certifier_name": "适当性认证方主体名称",
	                                "subject_investor_qualification_certifier_ref": "PID345678/3",
	                                "subject_investor_qualification_description": "适当性认证描述",
	                                "subject_investor_qualification_status": true,
	                                "subject_investor_qualification_sub": "适当性认证子类"
	                            }
	                        ],
	                        "qualification_authentication_information": [
	                            {
	                                "subject_certification_time": "2020-03-12 18:08:08",
	                                "subject_qualification_authenticator": "资质认证方",
	                                "subject_qualification_code": "code345",
	                                "subject_qualification_reviewer": "资质审核方",
	                                "subject_qualification_status": true,
	                                "subject_review_time": "2020-03-12 18:08:08",
	                                "subject_role_qualification_certification_doc": {
	                                    "file_name": "12312312",
	                                    "file_number": "1",
	                                    "hash": "12312312",
	                                    "summary": "12312312",
	                                    "term_of_validity": "yyyy/MM/dd",
	                                    "term_of_validity_type": "0",
	                                    "uri": "12312312"
	                                }
	                            }
	                        ],
	                        "subject_financial_qualification_type": 0,
	                        "subject_intermediary_qualification": [
	                            1
	                        ],
	                        "subject_market_roles_type": 0,
	                        "subject_qualification_category": 0
	                    }
	                ]
	            },
	            "organization_subject_information": {
	                "basic_information_of_enterprise": {
	                    "basic_information_description": {
	                        "subject_approval_time": "2020-03-12 18:08:08",
	                        "subject_association_articles": {
	                            "file_name": "12312312",
	                            "file_number": "1",
	                            "hash": "12312312",
	                            "summary": "12312312",
	                            "term_of_validity": "yyyy/MM/dd",
	                            "term_of_validity_type": "0",
	                            "uri": "12312312"
	                        },
	                        "subject_business_license": {
	                            "file_name": "12312312",
	                            "file_number": "1",
	                            "hash": "12312312",
	                            "summary": "12312312",
	                            "term_of_validity": "yyyy/MM/dd",
	                            "term_of_validity_type": "0",
	                            "uri": "12312312"
	                        },
	                        "subject_business_scope": "经营范围",
	                        "subject_city": "苏州",
	                        "subject_company_business": "主营业务",
	                        "subject_company_english_name": "shgjs",
	                        "subject_company_name": "上海股交所",
	                        "subject_company_profile": "公司简介",
	                        "subject_company_short_english_name": "英文简称",
	                        "subject_company_short_name": "sh",
	                        "subject_company_status": 0,
	                        "subject_company_status_deregistration": "注销原因",
	                        "subject_company_status_deregistration_date": "2020/08/12",
	                        "subject_company_status_windingup": "吊销原因",
	                        "subject_company_status_windingup_date": "2020/08/12",
	                        "subject_company_type": 0,
	                        "subject_contact_address": "苏州相城青龙港路1188号",
	                        "subject_contact_number": "13255669988",
	                        "subject_district": "相城区",
	                        "subject_document_information": 0,
	                        "subject_economic_type": 1,
	                        "subject_high_technology_enterprise": 0,
	                        "subject_industry": 1,
	                        "subject_insured_number": 3,
	                        "subject_internet_address": "www.tj.com",
	                        "subject_invoice_account_number": "发票账号",
	                        "subject_invoice_address": "发票地址",
	                        "subject_invoice_bank": "发票开户行",
	                        "subject_invoice_telephone_number": "发票电话",
	                        "subject_legal_type": 0,
	                        "subject_mail_box": "256425@yeah.net",
	                        "subject_name_used_before": [
	                            "曾用名1",
	                            "曾用名2",
	                            "曾用名3"
	                        ],
	                        "subject_office_address": "苏州相城青龙港路1188号",
	                        "subject_organization_nature": 0,
	                        "subject_paid_in_capital": 14237850,
	                        "subject_paid_in_capital_currency": "156",
	                        "subject_personnel_size": "45",
	                        "subject_postal_code": "225000",
	                        "subject_province": "江苏省",
	                        "subject_registered_address": "苏州相城青龙港路",
	                        "subject_registered_capital": "注册资本",
	                        "subject_registered_capital_currency": "156",
	                        "subject_registry_date": "2020/03/12",
	                        "subject_regulator": "主管单位",
	                        "subject_scale_type": 0,
	                        "subject_shareholders_number": 12,
	                        "subject_taxpayer_id_number": "纳税人识别号"
	                    },
	                    "leading_member_information": [
	                        {
	                            "subject_document_type": 1,
	                            "subject_key_personnel_address": "江苏苏州相城",
	                            "subject_key_personnel_appointment_end": "2020/08/12",
	                            "subject_key_personnel_appointment_start": "2012/08/12",
	                            "subject_key_personnel_contact": "15688889999",
	                            "subject_key_personnel_id": "311566199002178891",
	                            "subject_key_personnel_name": "姓名",
	                            "subject_key_personnel_nationality": "中国",
	                            "subject_key_personnel_position": 1,
	                            "subject_key_personnel_shareholding": 200000,
	                            "subject_key_personnel_shareholding_ratio": 0.25,
	                            "subject_key_personnel_type": 0
	                        }
	                    ]
	                }
	            }
	        }
	    }
	}`
		// strBytes, _ := json.Marshal(str)
		reader := strings.NewReader(str)
		// newReader := bytes.NewReader([]byte(""))

		n, err := minioClient.PutObject(ctx, bucketName, "zyf88888", reader, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully uploaded bytes: ", n)*/

	// 根据名称获取内容
	// object, err := minioClient.GetObject(ctx, bucketName, "zyf88888", minio.GetObjectOptions{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// localFile, err := os.Create("./MinIO2/local1.json")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if _, err = io.Copy(localFile, object); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 	查询元数据
	// objInfo, err := minioClient.StatObject(ctx, bucketName, "06DzZudY6ileqCyLw9Ph/0", minio.StatObjectOptions{})
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	fmt.Println("err2:", fmt.Sprintf("%s", err) == "The specified key does not exist.")
	// 	return
	// }
	// fmt.Println("-->", objInfo)

	// 删除数据
	err1 := minioClient.RemoveBucket(ctx, bucketName)
	if err1 != nil {
		panic(err1)
	}

}

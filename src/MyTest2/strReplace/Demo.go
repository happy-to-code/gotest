package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `contract ShareContract{

    //平台公钥
    string PubKey = "*****"

    //挂牌企业信息
    struct Enterprise { 
        string EquityCode                    //股权代码 
        string EquityAbbreviation            //股权简称 
        string DataDate                      //数据日期 
        string EquityClass                   //股权类别 
        uint64 TotalEquity                   //总股本 
        string FullNameOfCompany             //公司全称 
        string EnglishName                   //英文名称 
        string TypeOfEnterpriseCertificate   //企业证件号类别         
        string EnterpriseCertificateNumber   //企业证件编号     
        string RegisteredCurrency            //注册币种 
        string Extend                        //附加信息
    }

    //股东信息
    struct ShareHolder {
        string CustomerNumber        //客户号      
        int    ShareHolderNum        //股东号      
        string FundNum               //资金号
        string Address               //账户地址
        string CryptMsg              //加密信息
        string KeyInfo               //密钥信息
        string Extend                //基本信息
    }

    //添加挂牌企业
    public string AddListedCompany(string json){
        Enterprise en = json_to_obj<Enterprise>(json)
        string key = "com_" + en.EnterpriseCertificateNumber
        db_set(key,json)
        return "success"
    }
    //添加股东信息--将账户地址添加再这里结构里
    public string AddInvestor(string json){
        ShareHolder sh = json_to_obj<ShareHolder>(json)
        string cakey = "ca_" + sh.CustomerNumber
        db_set(cakey,json)

        string sakey = "sa_" + sh.ShareHolderNum
        db_set(sakey,json)

        string makey = "ma_" + sh.FundNum
        db_set(makey,json)

        //实例化账户资金
        string addrkey = "addr_" + sh.Address
        uint64 amount = 0
        db_set(addrkey,amount)
        return "success"
    }

    //发行验证

    //转账验证
    //回收验证
    //转板验证

    //获取平台公钥
    public string GetPlatform(){
        return PubKey
    }
}`

	fmt.Println("s:", s)
	fmt.Println("s.replace:", strings.Replace(s, "*****", "12345", 1))
}

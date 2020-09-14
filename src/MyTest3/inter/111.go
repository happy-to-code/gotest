package inter

type InvestorAndAccountAddressList struct {
	ClientNo        string // 客户号
	Extend          string // 扩展字段，json串
	Investor        string // 投资者信息
	SubjectObjectId string
	AccountInfoList []AccountInfo
}
type AccountInfo struct {
	ShareholderAccount ShareholderNoAndAddr // 股东号
	FundAccount        FundNoInfo           // 资金号
}
type ShareholderNoAndAddr struct {
	ShareholderNo string
	Address       string
	PubKey        string // 平台公钥

	AccountInfo     string
	AccountObjectId string
}
type FundNoInfo struct {
	FundNo          string
	AccountInfo     string
	AccountObjectId string
}

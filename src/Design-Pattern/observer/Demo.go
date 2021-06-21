package observer

import "fmt"

type Customer interface {
	update()
}

type CustomerA struct {
}

func (a *CustomerA) update() {
	fmt.Println("我是客户A，我收到报纸了")
}

type CustomerB struct {
}

func (b *CustomerB) update() {
	fmt.Println("我是客户B，我收到报纸了")
}

// 报社  被观察者
type NewsOffice struct {
	customers []Customer
}

func (n *NewsOffice) addCustomer(customer Customer) {
	n.customers = append(n.customers, customer)
}

func (n *NewsOffice) newsPaperCome() {
	// 通知所有客户
	n.notifyAllCustomer()
}

func (n *NewsOffice) notifyAllCustomer() {
	for _, customer := range n.customers {
		customer.update()
	}
}

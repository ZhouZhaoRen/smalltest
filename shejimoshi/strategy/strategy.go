package strategy

import "fmt"

// PaymentStrategy 最关键的接口
type PaymentStrategy interface {
	Pay(string)
}

// Cash 结构体之一
type Cash struct {
	Count int
}

// Pay 结构体实现接口
func (c *Cash) Pay(key string) {
	fmt.Printf("利用现金支付,key==%s\t", key)
	fmt.Printf("count==%d\n", c.Count)
}

// Bank 结构体之一
type Bank struct {
	Name    string
	Address string
}

// Pay 结构体实现接口
func (b *Bank) Pay(key string) {
	fmt.Printf("利用银行支付,key==%s\t", key)
	fmt.Printf("name==%s  address=%s\n", b.Name, b.Address)
}

// Payment 统一的结构体
type Payment struct {
	key      string
	strategy PaymentStrategy
}

// Pay 并不是去实现接口，而是通过接口去调用方法
func (p *Payment) Run() {
	p.strategy.Pay(p.key)
}

// NewPayment 实例化通过结构体
func NewPayment(key string, strategy PaymentStrategy) *Payment {
	return &Payment{
		key:      key,
		strategy: strategy,
	}
}

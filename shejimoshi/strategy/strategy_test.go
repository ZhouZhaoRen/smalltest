package strategy

import "testing"

var userMap=map[int]PaymentStrategy{
	0:&Cash{100},
	1:&Bank{Name: "中国银行",Address: "湛江市"},
}

func TestNewPayment(t *testing.T) {
	payment := NewPayment("zzr", userMap[0])
	payment.Run()
	userMap[0].Pay("zzr")
}

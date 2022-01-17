package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "demo2"
	//msg.Value = sarama.StringEncoder("demo2")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"1.117.76.139:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	for i := 0; i < 10; i++ {
		msg.Value = sarama.StringEncoder("demo"+strconv.Itoa(i))
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err:", err)
			return
		}
		time.Sleep(time.Second*2)
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}

}

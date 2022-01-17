package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	cache "github.com/ZhouZhaoRen/zzr-cache"
	"smalltest/job"
	"time"
)

func main() {
	test06()
}

func test06() {
	c:=make(chan int,3)
	fmt.Println(len(c))
	c<-1
	fmt.Println(len(c))
}

func test05() {
	now := time.Now().UnixNano() / 1e6
	fmt.Println(now)
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now().UnixNano() / 1e6)
}

func test04() {
	job.JobContainerCache.Run()
}

func test03() {
	consumer, err := sarama.NewConsumer([]string{"1.117.76.139:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		for msg := range pc.Messages() {
			fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
		}
		// 异步从每个分区消费信息
		//go func(sarama.PartitionConsumer) {
		//	for msg := range pc.Messages() {
		//		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
		//	}
		//}(pc)
	}
}

func test02() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"1.117.76.139:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

func test01() {
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}
	//
	fmt.Println("count==", c.ItemCount())
	for index, value := range c.Items() {
		fmt.Printf("index==%s   value==%+v\n", index, value)
	}
	n, err := c.DecrementInt("baz", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("n==", n)
	fmt.Println(c.Get("bar"))
}

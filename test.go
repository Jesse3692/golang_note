package main

func consumer(data chan int, done chan bool) {
	/*消费者*/
	for x := range data {
		println("recv:", x)
	}

	done <- true
}

func producer(data chan int) {
	/*生产者*/
	for i := 0; i < 4; i++ {
		data <- i
	}
	close(data)
}

func main() {
	done := make(chan bool) // 用于接收消费结束信号
	data := make(chan int)  // 数据管道
	go consumer(data, done) // 启动消费者
	go producer(data)       // 启动生产者
	<-done                  // 阻塞，直到消费者发回结束信号
}

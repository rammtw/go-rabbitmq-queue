package go_rabbitmq_queue

type Configuration struct {
	AMQPConnectionURL string
}

type AddTask struct {
	Number1 int
	Number2 int
}

var Config = Configuration{
	AMQPConnectionURL: "amqp://guest:guest@localhost:5672/",
}
package queue

import(
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string){
	if err != nil{
		log.Panicf("%s: %s", msg, err)
	}
}

var (
	Queue amqp.Queue
	Channel *amqp.Channel
)
 

func StartQueue(){
	conn, err := amqp.Dial("amqp://user:password@queue:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	
	q, err := ch.QueueDeclare(
		"HotelsQueue", // Name
		false,     // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	failOnError(err, "Failed to declare a queue")
	
	Queue = q
	Channel = ch
	// defer ch.Close()
	// defer conn.Close()
	
}

package queue

import(
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

// func failOnError(err error, msg string){
// 	if err != nil{
// 		log.Panicf("%s: %s", msg, err)
// 	}
// }

func Consumer(){
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
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

    msgs,err := ch.Consume(
		q.Name, // Name
		"",     // Consumer
		true, // Auto-ack
		false,  // Exclusive
		false,  // No-local
		true, // No-wait
		nil, // args
	)
	if err != nil {
		log.Fatalf("Failed to get a message: %s", err)
	}

	var forever chan struct{}

        go func() {
                for d := range msgs {
                        log.Printf(" [x] %s", d.Body)
                }
        }()

        log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
        <-forever
}
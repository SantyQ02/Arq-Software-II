package queue

import(
	log "github.com/sirupsen/logrus"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/google/uuid"
	"fmt"
	"mvc-go/service"
	"encoding/json"
	"os"
)

func failOnError(err error, msg string){
	if err != nil{
		log.Errorf("%s: %s", msg, err)
	}
}

type Message struct{
    HotelID uuid.UUID  `json:"hotel_id"`
    Action  string `json:"action"`
}

func Consumer(){
	conn, err := amqp.Dial(fmt.Sprintf("amqp://user:password@%s:5672/", os.Getenv("RABBITMQ_SERVICE_URL")))
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
		log.Errorf("Failed to get messages: %s", err)
	}

	var forever chan struct{}

        go func() {
                for d := range msgs {
					var message Message
					err = json.Unmarshal(d.Body, &message)
					if err != nil {
						log.Error(err.Error())
					}
					err = service.SolrService.AddOrUpdateHotel(message.HotelID)
					if err != nil {
						log.Error(err.Error())
					}
                    log.Debug(fmt.Sprintf(" [x] %s", message.HotelID))
                }
        }()

        log.Info(" [*] Waiting for logs. To exit press CTRL+C")
        <-forever
}
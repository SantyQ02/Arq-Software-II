# Microservice Hotels

# Notes:

When insert hotels(InsertHotel), Photos should be null. After insert hotel, insert one photo at a time(UploadPhoto).

After inserting the photo, you can access it at this endpoint: 'http://localhost:8090/api/public/images/hotels/name.ext'.

To insert or delete amenities, you should use UpdateHotel.

# Consumer

func failOnError(err error, msg string){
	if err != nil{
		log.Panicf("%s: %s", msg, err)
	}
}

func Consumer(){
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
}

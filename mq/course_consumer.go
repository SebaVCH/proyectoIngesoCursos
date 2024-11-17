package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"proyectoIngesoCursos/models"
	"proyectoIngesoCursos/utils"

	"github.com/streadway/amqp"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Estructura del mensaje recibido desde RabbitMQ
type RabbitMQMessage struct {
	Pattern string `json:"pattern"` // "get_course_price", "get_course_name"
	Data    string `json:"data"`    // CourseID o cualquier otro dato
	ID      string `json:"id"`
}

// Iniciar el consumidor de RabbitMQ para el microservicio de cursos
func StartCourseConsumer() error {
	// Conectar a RabbitMQ
	conn, ch, err := utils.ConnectRabbitMQ()
	if err != nil {
		return fmt.Errorf("error connecting to RabbitMQ: %w", err)
	}
	defer conn.Close()
	defer ch.Close()

	// Conectar a la base de datos
	db, err := gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	// Declarar la cola
	queueName := "courses_queue"
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	// Registrar el consumidor
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	go func() {
		for d := range msgs {
			fmt.Printf("Mensaje recibido: %s\n", string(d.Body))

			// Deserializar el mensaje
			var msg RabbitMQMessage
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Printf("Error al deserializar el mensaje: %s", err)
				continue
			}

			var responseBody []byte

			// Manejar diferentes patrones (price y name)
			switch msg.Pattern {
			case "get_course_price":
				// Buscar el curso en la base de datos
				var curso models.Curso
				result := db.Where("course_id = ?", msg.Data).First(&curso)
				if result.Error != nil {
					log.Printf("No se encontró el curso con ID %s: %s", msg.Data, result.Error)
					continue
				}

				// Responder con el precio y la foto del curso
				response := struct {
					Price    int    `json:"price"`
					ImageURL string `json:"image_url"`
				}{
					Price:    curso.Price,
					ImageURL: curso.ImageURL,
				}

				responseBody, err = json.Marshal(response)
				if err != nil {
					log.Printf("Error al serializar la respuesta: %s", err)
					continue
				}

			case "get_course_name":
				// Buscar el curso en la base de datos
				var curso models.Curso
				result := db.Where("course_id = ?", msg.Data).First(&curso)
				if result.Error != nil {
					log.Printf("No se encontró el curso con ID %s: %s", msg.Data, result.Error)
					continue
				}

				// Responder con el nombre y la foto del curso
				response := struct {
					Name     string `json:"name"`
					ImageURL string `json:"image_url"`
				}{
					Name:     curso.Title,
					ImageURL: curso.ImageURL,
				}

				responseBody, err = json.Marshal(response)
				if err != nil {
					log.Printf("Error al serializar la respuesta: %s", err)
					continue
				}

			default:
				log.Printf("Patrón no soportado: %s", msg.Pattern)
				continue
			}

			// Publicar la respuesta en la cola de respuesta
			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          responseBody,
				})
			if err != nil {
				log.Printf("Error al publicar la respuesta: %s", err)
			}
		}
	}()

	log.Printf("Esperando mensajes. Presiona CTRL+C para salir.")
	<-make(chan bool)

	return nil
}

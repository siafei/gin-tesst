package rabbit

import (
	"github.com/streadway/amqp"
	"log"
)

type Rabbitmq struct {
	name            string
	logger          *log.Logger
	connection      *amqp.Connection
	channel         *amqp.Channel
	done            chan bool
	notifyConnClose chan *amqp.Error
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	isReady         bool
}


func New()  {

}

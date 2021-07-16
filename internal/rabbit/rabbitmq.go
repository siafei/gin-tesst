package rabbit

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
)

type Pool struct {
	name   string
	logger *log.Logger
	cons   []*Conn
	done   chan bool
	mu 		sync.Mutex
	notifyConfirm chan amqp.Confirmation
	isReady       bool
}

type Conn struct {
	con             *amqp.Connection
	chans           []*Chan
	notifyConnClose chan *amqp.Error
}

type Chan struct {
	channel         *amqp.Channel
	notifyChanClose chan *amqp.Error
}

func New() {

}

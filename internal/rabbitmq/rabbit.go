package rabbitmq

type Rabbit struct {
	Exchange string
	Queue    string
	RouteKey string
	Message  string
}

//func Push(rabbit Rabbit) {
//	conn, err := amqp.Dial(addr)
//
//	if err != nil {
//		return nil, err
//	}
//
//	session.changeConnection(conn)
//	log.Println("Connected!")
//	return conn, nil
//}

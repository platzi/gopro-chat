package chat

const archiveSize = 10

var (
	// Envía la respuesta de un canal. Esto enviara el archivo y los nuevos mensajes.
	subscribe = make(chan (chan<- Subscription), 10)
	// Envía un canal para desuscribirse.
	unsubscribe = make(chan (<-chan Event), 10)
	// Envía eventos que se van a publicar.
	publish = make(chan Event, 10)
)

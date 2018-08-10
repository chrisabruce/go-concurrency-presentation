 package Messages

 import "sync"

type Message struct {
	To string,
	From string,
	Body string,
}

type PushNotifQueue struct {
	sync.Mutex
	Msgs   []*Message
}

func (q *PushNotifQueue) Queue(msg *Message) {
	q.Lock()
	q.Msgs = append(q.Msgs, msg)
	q.Unlock()
}

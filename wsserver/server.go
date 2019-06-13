package wsserver

import (
	"sync"
	"time"
	"github.com/gobwas/ws"
	"net"
	"github.com/gobwas/ws/wsutil"
	"net/http"
	"github.com/google/uuid"
	"fmt"
	"country-api/lg"
)

var instance *sync.Map
var once sync.Once

func GetConnections() *sync.Map {

	once.Do(func() {
		instance = &sync.Map{}
	})

	return instance
}

func Run(host, port string, connTimeout, tickerTimeout time.Duration) (err error)  {

	cons := GetConnections()

	// тикер закрывает соединения старые соединения
	runTicker(tickerTimeout, connTimeout)

	err = http.ListenAndServe(host + ":" + port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		code := uuid.New()
		conn, _, _, err := ws.UpgradeHTTP(r, w)

		if err != nil {
			// handle error
			fmt.Println("Error upgrade HTTP")
			lg.Error("Error upgrade HTTP in WebSocketServer")
			return
		}

		cons.Store(code.String(), conn)

		go func(code uuid.UUID) {

			defer closeConnection(code)

			for {

				if v, ok := cons.Load(code.String()); ok == true {

					c := v.(net.Conn)

					msg, op, err := wsutil.ReadClientData(c)

					if err != nil {
						fmt.Println("error", err)
						// handle error
						return
					}

					err = wsutil.WriteServerMessage(c, op, msg)

					if err != nil {
						fmt.Println("error", err)
						// handle error
						return
					}

				} else {
					return
				}
			}

		}(code)

	}))

	return
}

func closeConnection(code uuid.UUID) {

	cons := GetConnections()

	c, ok := cons.Load(code.String())

	if ok {
		c.(net.Conn).Close()
		cons.Delete(code.String())
	}

	fmt.Println("close connection", code.String())
}

func runTicker(timeout, connTimeout time.Duration) {

	cons := GetConnections()

	ticker := time.NewTicker(timeout)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:

				go func() {

					length := 0

					cons.Range(func(code interface{}, _ interface{}) bool {

						length++

						//c, ok := cons.Load(code.(string))
						//
						//if ok {
						//	c.(net.Conn).Close()
						//	cons.Delete(code.(string))
						//}

						v, ok := cons.Load(code)

						if ok {

							op := ws.OpCode(ws.OpText)

							c := v.(net.Conn)
							wsutil.WriteServerMessage(c, op, []byte("hello from server message"))
						}

						return true
					})

					fmt.Println("we have socket connections: ", length)

				}()

			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()
}


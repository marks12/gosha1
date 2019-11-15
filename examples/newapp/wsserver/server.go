package wsserver

import (
    "encoding/json"
    "net"
    "net/http"
    "sync"
    "time"

    "github.com/gobwas/ws"
    "github.com/gobwas/ws/wsutil"
    "github.com/google/uuid"
    "errors"
)

var instance *sync.Map
var once sync.Once

func GetConnections() *sync.Map {
    once.Do(func() {
        instance = &sync.Map{}
    })

    return instance
}

type UserMessage struct {
    Type string
    Data interface{}
}

type UserResponse struct {
    Type    string
    Result  interface{}
    Success bool
}

type UserConnection struct {
    Con            net.Conn
    WsToken        string
    HttpToken      string
    AdminSessionId string
    UpdatedAt      time.Time
}

func Run(host, port string) (err error) {
    setHandlers()
    err = http.ListenAndServe(host+":"+port, http.HandlerFunc(wssHandler))

    return
}

func setHandlers() {
    SetMessageHandler("SetToken", SetTokenHandler)
    SetMessageHandler("hello", HelloHandler)
    SetMessageHandler("Event", EventHandler)
}

func wssHandler(w http.ResponseWriter, r *http.Request) {

    code := uuid.New()

    conn, _, _, err := ws.UpgradeHTTP(r, w)
    if err != nil {
        // handle error
        return
    }

    cons := GetConnections()
    cons.Store(code, &UserConnection{conn, code.String(), "", "", time.Now()})

    go runWssReader(cons, code)

}

func runWssReader(cons *sync.Map, code uuid.UUID) {

    defer closeConnection(code)

    for {

        if v, ok := cons.Load(code); ok == true {

            c := v.(*UserConnection)

            msg, op, err := wsutil.ReadClientData(c.Con)

            if err != nil {
                // handle error
                return
            }

            if len(string(msg)) > 0 {

                cons.Store(code, c)

                message, err := ParseMessage(msg)
                if err != nil {
                    // handle error
                    continue
                }

                answer := GetMessageHandler(message.Type)(message, c)

                answer.Type = answer.Type+"Response"

                bAnswer, err := json.Marshal(answer)
                if err != nil {
                    // handle error
                    continue
                }
                err = wsutil.WriteServerMessage(c.Con, op, bAnswer)

                if err != nil {
                    // handle error
                    return
                }
            }

        } else {
            return
        }
    }
}

func ParseMessage(msg []byte) (UserMessage, error) {
    message := UserMessage{}
    err := json.Unmarshal(msg, &message)
    return message, err
}

func SendToUserByToken(token string, msg []byte) (err error) {
    cons := GetConnections()

    sent := false

    cons.Range(func(code interface{}, _ interface{}) bool {
        if v, ok := cons.Load(code); ok == true {
            op := ws.OpCode(ws.OpText)

            if c, ok := v.(*UserConnection); ok == true {
                if c.HttpToken == token {
                    wsutil.WriteServerMessage(c.Con, op, msg)
                    sent = true
                }
            }
        }

        return true
    })

    if sent {
        return nil
    }

    return errors.New("User not found with token:" + token)
}

func closeConnection(code uuid.UUID) {

    cons := GetConnections()

    c, ok := cons.Load(code)

    if ok {

        if ucon, ok := c.(UserConnection); ok == true {
            ucon.Con.Close()
            cons.Delete(code)
        }
    }
}


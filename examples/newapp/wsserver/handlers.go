package wsserver

import (
    "newapp/google"
    "newapp/types"
    "encoding/json"
)

type WsHandler func(msg UserMessage, userConn *UserConnection) (answer UserResponse)

type RpcMessage struct {
    Message        UserMessage
    AdminSessionId string
}

var defaultHandler = func(msg UserMessage, userConn *UserConnection) (answer UserResponse) { return }

var handlers = map[string]WsHandler{}

func SetMessageHandler(messageType string, h WsHandler) {
    handlers[messageType] = h
}

func GetMessageHandler(messageType string) (h WsHandler) {
    h, ok := handlers[messageType]
    if !ok {
        h = defaultHandler
    }
    return
}

func SetTokenHandler(msg UserMessage, userConn *UserConnection) (answer UserResponse) {
    if msg.Data != nil {
        token := msg.Data.(string)

        auth := types.Authenticator{Token: token}
        if auth.IsAuthorized() {
            userConn.HttpToken = token
            answer = UserResponse{
                Type:    msg.Type,
                Result:  "token was set",
                Success: true,
            }
            return
        }
    }
    answer = UserResponse{
        Type:    msg.Type,
        Result:  "invalid token",
        Success: false,
    }
    return
}

func HelloHandler(msg UserMessage, userConn *UserConnection) (answer UserResponse) {
    name := ""
    if msg.Data == nil {
        name = "Anonymous"
    } else {
        name = msg.Data.(string)
    }

    answer = UserResponse{
        Type:    msg.Type,
        Result:  "Hello " + name,
        Success: true,
    }
    return
}

type SiteStat struct {
    Event       string
    Url         string
    ElementName string
}

func EventHandler(msg UserMessage, userConn *UserConnection) (answer UserResponse) {
    if msg.Data != nil {
        b, err := json.Marshal(msg.Data)
        if err != nil {
            answer = UserResponse{
                Type:    msg.Type,
                Result:  err.Error(),
                Success: false,
            }
            return
        }
        stats := SiteStat{}
        err = json.Unmarshal(b, &stats)
        if err != nil {
            answer = UserResponse{
                Type:    msg.Type,
                Result:  err.Error(),
                Success: false,
            }
            return
        }
        GaMessage := google.NewGMessage()
        GaMessage.EventAction = stats.Event
        //GaMessage.EventValue = stats.Url
        GaMessage.EventCategory = stats.ElementName
        GaMessage.ClientId = userConn.WsToken

        if err := GaMessage.Send(); err != nil {
            answer = UserResponse{
                Type:    msg.Type,
                Result:  err.Error(),
                Success: false,
            }
            return
        }

        answer = UserResponse{
            Type:    msg.Type,
            Result:  "",
            Success: true,
        }
        return
    }

    answer = UserResponse{
        Type:    msg.Type,
        Result:  "Data not send",
        Success: false,
    }
    return
}


package cmd

const usualWsserver = `package wsserver

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

`

const usualWssHandlers = `package wsserver

import (
    "{ms-name}/google"
    "{ms-name}/types"
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

package wsserver

import "biz/view"

type BuildMode string

type BuildModeInterface interface {
	GetBuildMode() BuildMode
}

func (mode BuildMode) GetBuildMode() BuildMode {
	return mode
}

const ModeReplace BuildMode = "replace"
const ModeAppend BuildMode = "append"

func (BuildMode BuildMode) ToString() string {
	return string(BuildMode)
}

type Instruction struct {
	EventName       string            ` + "`" + `json:"EventName,omitempty"` + "`" + `
	CollectFields   []FormFields      ` + "`" + `json:"CollectFields,omitempty"` + "`" + `
	SetLocalStorage []KeyVal          ` + "`" + `json:"SetLocalStorage,omitempty"` + "`" + `
	SetCookie       []KeyVal          ` + "`" + `json:"SetCookie,omitempty"` + "`" + `
	StoreVar        []KeyVal          ` + "`" + `json:"StoreVar,omitempty"` + "`" + `
	SetInnerHtml    []KeyVal          ` + "`" + `json:"SetInnerHtml,omitempty"` + "`" + `
	AddClass        []KeyVal          ` + "`" + `json:"AddClass,omitempty"` + "`" + `
	ReplaceClasses  []KeyVal          ` + "`" + `json:"ReplaceClasses,omitempty"` + "`" + `
	RemoveClass     []KeyVal          ` + "`" + `json:"RemoveClass,omitempty"` + "`" + `
	Redirect        string            ` + "`" + `json:"Redirect,omitempty"` + "`" + `
	RestApiRequests []ApiRequest      ` + "`" + `json:"RestApiRequests,omitempty"` + "`" + `
	ConsoleLog      []KeyVal          ` + "`" + `json:"ConsoleLog,omitempty"` + "`" + `
	BuildHtmlBlock  []HtmlBuildConfig ` + "`" + `json:"BuildHtmlBlock,omitempty"` + "`" + `
	ClearHtmlBlock  ClearHtmlBlock    ` + "`" + `json:"ClearHtmlBlock,omitempty"` + "`" + `
}

type KeyVal struct {
	Key string ` + "`" + `json:"Key,omitempty"` + "`" + `
	Val string ` + "`" + `json:"Val,omitempty"` + "`" + `
}

type HtmlBuildConfig struct {
	TemplateStoreKey string                  ` + "`" + `json:"TemplateStoreKey,omitempty"` + "`" + `
	DataStoreKey     string                  ` + "`" + `json:"DataStoreKey,omitempty"` + "`" + `
	Mode             BuildModeInterface      ` + "`" + `json:"Mode,omitempty"` + "`" + `
	ContentSlot      view.NamespaceInterface ` + "`" + `json:"ContentSlot,omitempty"` + "`" + `
}

type ClearHtmlBlock struct {
	ContentSlot view.NamespaceInterface ` + "`" + `json:"ContentSlot,omitempty"` + "`" + `
}

type ApiRequest struct {
	DataVar             string        ` + "`" + `json:"DataVar,omitempty"` + "`" + `
	FailInstructions    []Instruction ` + "`" + `json:"FailInstructions,omitempty"` + "`" + `
	SuccessRespStatus   int           ` + "`" + `json:"SuccessRespStatus,omitempty"` + "`" + `
	SuccessInstructions []Instruction ` + "`" + `json:"SuccessInstructions,omitempty"` + "`" + `
	Type                string        ` + "`" + `json:"Type,omitempty"` + "`" + `
	Url                 string        ` + "`" + `json:"Url,omitempty"` + "`" + `
}

type FormFields struct {
	CollectorKey string ` + "`" + `json:"CollectorKey,omitempty"` + "`" + `
	Id string ` + "`" + `json:"Id,omitempty"` + "`" + `
	Validators []FieldValidator ` + "`" + `json:"Validators,omitempty"` + "`" + `
}

type FieldValidator struct {
	Error      FieldValidatorError ` + "`" + `json:"Error,omitempty"` + "`" + `
	Equal      []string            ` + "`" + `json:"Equal,omitempty"` + "`" + `
	NotEqual   []string            ` + "`" + `json:"NotEqual,omitempty"` + "`" + `
	MaxLen     int                 ` + "`" + `json:"MaxLen,omitempty"` + "`" + `
	MinLen     int                 ` + "`" + `json:"MinLen,omitempty"` + "`" + `
	Regular    string              ` + "`" + `json:"Regular,omitempty"` + "`" + `
	NotRegular string              ` + "`" + `json:"NotRegular,omitempty"` + "`" + `
	Min        float64             ` + "`" + `json:"Min,omitempty"` + "`" + `
	Max        float64             ` + "`" + `json:"Max,omitempty"` + "`" + `
}

type FieldValidatorError struct {
	MinLen     string   ` + "`" + `json:"MinLen,omitempty"` + "`" + `
	MaxLen     string   ` + "`" + `json:"MaxLen,omitempty"` + "`" + `
	Regular    string   ` + "`" + `json:"Regular,omitempty"` + "`" + `
	NotRegular string   ` + "`" + `json:"NotRegular,omitempty"` + "`" + `
	Equal      []string ` + "`" + `json:"Equal,omitempty"` + "`" + `
	NotEqual   []string ` + "`" + `json:"NotEqual,omitempty"` + "`" + `
	Min        string   ` + "`" + `json:"Min,omitempty"` + "`" + `
	Max        string   ` + "`" + `json:"Max,omitempty"` + "`" + `
}


`

var usualTemplateWsserver = template{
	Path:    "./wsserver/server.go",
	Content: assignCurrentDateTime(assignMsName(usualWsserver)),
}

var usualTemplateWssHandlers = template{
	Path:    "./wsserver/handlers.go",
	Content: assignCurrentDateTime(assignMsName(usualWssHandlers)),
}

package cmd

const usualServicesCaller = `package services

import "runtime"

func GetCallerFunction() string {

    pc := make([]uintptr, 10)
    runtime.Callers(3, pc)
    f := runtime.FuncForPC(pc[0])
    return f.Name()
}

`

var usualTemplateServicesCaller = template{
    Path:    "./services/caller.go",
    Content: usualServicesCaller,
}

const usualServicesTicket = `package services

type TicketInterface interface {

    GetId()             int
}

type Ticket struct {
    id int
    TicketInterface
}

func (ticket Ticket) GetId() int {
    return ticket.id
}

`

var usualTemplateServicesTicket = template{
    Path:    "./services/ticket.go",
    Content: usualServicesTicket,
}

package services

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


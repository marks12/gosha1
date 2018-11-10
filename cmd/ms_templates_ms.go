package cmd

const msMsTicket = `package ms

import "github.com/google/uuid"

// add some tickets
`
const msMsEntityTicket = `

const Ticket{entity-name}Find            = "{new-guid1}" // список
const Ticket{entity-name}Create          = "{new-guid2}" // Создать
const Ticket{entity-name}Read            = "{new-guid3}" // получить по ID
const Ticket{entity-name}Update          = "{new-guid4}" // Обновить
const Ticket{entity-name}Delete          = "{new-guid5}" // удалить

var TicketGuid{entity-name}Find, _           = uuid.Parse(Ticket{entity-name}Find)
var TicketGuid{entity-name}Create, _         = uuid.Parse(Ticket{entity-name}Create)
var TicketGuid{entity-name}Read, _           = uuid.Parse(Ticket{entity-name}Read)
var TicketGuid{entity-name}Update, _         = uuid.Parse(Ticket{entity-name}Update)
var TicketGuid{entity-name}Delete, _         = uuid.Parse(Ticket{entity-name}Delete)

`

var msTemplateMsTicket = template{
    Path:    "./ms/ticket.go",
    Content: msMsTicket,
}

var msTemplateMsTicketEntity = template{
    Path:    "./ms/ticket.go",
    Content: assignGuid(msMsEntityTicket),
}

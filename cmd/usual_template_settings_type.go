package cmd

const usualSettingsType = `package settings

import (
    "encoding/binary"
    "github.com/google/uuid"
)

type ConfigId uint64

func (con ConfigId) Uuid() uuid.UUID {
    i := uint64(con)
    b := make([]byte, 16)
    binary.LittleEndian.PutUint64(b, i)
    id, _ := uuid.FromBytes(b)
    return id
}

func (con ConfigId) Int() int {
    return int(con)
}
`

var usualTemplateSettingsType = template{
	Path:    "./settings/type.go",
	Content: usualSettingsType,
}

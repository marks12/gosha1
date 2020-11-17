package cmd

import "strings"

const usual_go_mod = `module {ms-name}

go {go-version}
`

func GetGoModContent() string {
    raw := strings.Replace(usual_go_mod, "{ms-name}", GetCurrentAppName(), -1)
    raw = strings.Replace(raw, "{go-version}", GetGoVersion(), -1)
    return raw
}


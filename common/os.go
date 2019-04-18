package common

import (
	"os"
	"regexp"
	"errors"
)

type RegularFind struct {
	BoolResult bool
	StringResult string
	ArrayResult []string
}

func GetOsArgument(arg string) (RegularFind, error) {

	for _, a := range os.Args {

		if a == arg {
			return RegularFind{BoolResult:true}, nil
		}

		re := regexp.MustCompile(arg + "=([a-zA-Z0-9-_]*)")
		as := re.FindSubmatch([]byte(a))

		if len(as) > 1 {

			arrs := []string{}

			for _, s := range as {
				arrs = append(arrs, string(s))
			}

			return RegularFind{
				StringResult: string(as[1]),
				ArrayResult: arrs,
			}, nil
		}
	}

	return RegularFind{}, errors.New("Not found argument " + arg)
}

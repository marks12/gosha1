package mode

import (
	"fmt"
	"gosha/common"
	"os"
)

func CheckReadOnly() {
	if common.InArray("--readonly", os.Args) || common.InArray("readonly", os.Args) {
		SetReadOnlyMode()

		if IsReadOnlyMode() {
			fmt.Printf("\nREADONLY MODE\n\n")
		}
	}
}

func CheckSelfUpdate() {
	if common.InArray("--update", os.Args) || common.InArray("update", os.Args) {
		SetSelfUpdateMode()

		if IsSelfUpdate() {
			fmt.Printf("\nSELF UPDATE MODE ENABLED\n\n")
		}
	}
}




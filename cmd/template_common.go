package cmd

const removeLineComment = `remove this line for disable generator functionality`

func getRemoveLine(name string) string {
	return "//" + name + " " + removeLineComment
}
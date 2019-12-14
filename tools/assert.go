package tools

//Assert for test
func Assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

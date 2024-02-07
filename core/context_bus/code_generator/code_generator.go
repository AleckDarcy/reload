package code_generator

// Code generator parses user observation and re-generate the code to get:
// 1. enriched observation: add and manage more information written by other parties;
// 2. higher performance: avoid string parsing (parser.go), runtime information fetching (runtime.go)

// init initialize storages that support generated code
func init() {
	CodeInfoStorage.infos = make([]CodeInfo, CodeInfo_End)
}
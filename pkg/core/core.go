package core

type DiffOperation int

const (
	DeleteOperation DiffOperation = -1
	Equal DiffOperation= 0
	InsertOperation DiffOperation = 1
)

type DiffSegment struct {
	Operation DiffOperation
	Text string
}

type DiffMatchPatch interface {
	DiffMain(a string, b string) []DiffSegment
	DiffCleanupSemantic(diffs []DiffSegment) []DiffSegment
	DiffCleanupEfficiency(diffs []DiffSegment) []DiffSegment
	DiffLevenshtein(diffs []DiffSegment) int
	DiffPrettyHtml(diffs []DiffSegment) string
	MatchMain(text string, pattern string, location int) int
	PatchMake(a string, b string)
	PatchMakeFromDiff(diffs []DiffSegment)
	PatchMakeFromTextAndDiff(a string, diffs []DiffSegment)
	PatchToText()
	PatchFromText()
	PatchApply()
}

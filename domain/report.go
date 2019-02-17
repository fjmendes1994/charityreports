package domain

type ProjectReport struct {
	ProjectPath string
	Branch      string
	Commits     []Commit
}

type Commit struct {
	Index   int
	Id      string
	Reports CommitReports
}

type CommitReports struct {
	Coverage Coverage
}

type Coverage struct {
	Percentage float64
}

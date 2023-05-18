package processors

type BeginTransaction struct {
	Name string
	IProcessor
}

type CommitTransaction struct {
	Name string
	IProcessor
}

type Commit struct {
	IProcessor
}

type Rollback struct {
	IProcessor
}

func NewBeginTransaction(name string) *BeginTransaction {
	return &BeginTransaction{
		Name: name,
	}
}

func NewCommitTransaction(name string) *CommitTransaction {
	return &CommitTransaction{
		Name: name,
	}
}

func NewCommit() *Commit {
	return &Commit{}
}

func NewRollback() *Rollback {
	return &Rollback{}
}

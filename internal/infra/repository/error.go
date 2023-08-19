package repository

type RepoError struct{}

func (e *RepoError) Error() string {
	return "Error from repository."
}

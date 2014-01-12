package models

type AlreadyExistsError struct {}
type SaveErr struct{}
type NotFoundErr struct{}
type DBErr struct{}

func (err AlreadyExistsError) Error() string {
  return "Already exists"
}

func (err SaveErr) Error() string {
  return "Error during save"
}

func (err NotFoundErr) Error() string {
  return "Not found"
}

func (err DBErr) Error() string {
  return "Database Error"
}

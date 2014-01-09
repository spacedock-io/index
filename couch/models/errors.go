package models

type AlreadyExistsError struct {
}

func (err AlreadyExistsError) Error() string {
  return "Already exists"
}

type ConflictError struct {
}

func (err ConflictError) Error() string {
  return "Conflict"
}

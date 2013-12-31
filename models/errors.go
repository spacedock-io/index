package models

type AlreadyExistsError struct {
}

func (err AlreadyExistsError) Error() string {
  return "Already exists"
}

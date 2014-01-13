package models

type Error struct {
  Message string
  Type string
}

var (
  AlreadyExistsError = &Error{
    Message: "Already exists.",
    Type: "exists",
  }
  SaveErr = &Error{
    Message: "Error during save.",
    Type: "save",
  }
  NotFoundErr = &Error{
    Message: "Not found.",
    Type: "not_found",
  }
  DBErr = &Error{
    Message: "Database error.",
    Type: "db",
  }
  TokenErr = &Error{
    Message: "Error generating token",
    Type: "token",
  }
)

func (err *Error) Error() string {
  return err.Message
}

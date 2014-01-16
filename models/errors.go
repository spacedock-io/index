package models

type Error struct {
  Message string
  Type string
}

var (
  AccessSetError = &Error{
    Message: "Unable to set access.",
  }
  AlreadyExistsError = &Error{
    Message: "Already exists.",
  }
  SaveErr = &Error{
    Message: "Error during save.",
  }
  NotFoundErr = &Error{
    Message: "Not found.",
  }
  DBErr = &Error{
    Message: "Database error.",
  }
  TokenErr = &Error{
    Message: "Error generating token",
  }
  TokenNotFound = &Error{
    Message: "Token could not be found.",
  }
)

func (err *Error) Error() string {
  return err.Message
}

package models

var SpacedockUser = &User{
  Username: "spacedock",
  Emails: []Email{
    Email{
      Email: "hello@spacedock.io",
    },
  },
}

var SpacedockUserPassword = "4321"

var SpacedockFooRepo = &Repo{
  Namespace: "spacedock",
  Name: "foo",
}

var SpacedockFooImages = make([]interface{}, 0)

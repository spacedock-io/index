package models



type Image struct {
  Id int64  `json:"-"`
  Uuid string   `json:"id"`
  Json []byte `json:"-"`
  Checksum string `json:"checksum,omitempty"`
  Size int64 `json:"size,omitempty"`
  RepoId int64  `json:"-"`
}

func (img *Image) Create(id string) {
  img.Uuid = id
}

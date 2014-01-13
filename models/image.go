package models



type Image struct {
  Id int64
  Uuid string
  Json []byte
  Checksum string
  Size int64
  Ancestry []string
  RepoId int64
}

func (img *Image) Create(id string) {
  img.Uuid = id
}
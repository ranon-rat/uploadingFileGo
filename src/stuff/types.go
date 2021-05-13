package stuff

type Api struct {
	Files []File `json:"Files"`
	Size  int    `json:"size"`
}
type File struct {
	Id   int    `json:"id"`
	Name string `json:"filename"`
}
type HttpCodeError struct {
	Err  string
	Code int
}

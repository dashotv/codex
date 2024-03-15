package definitions

// File represents a file in the file system.
// model: true
type File struct {
	Path       string `json:"path" bson:"path" grimoire:"index"`
	Size       int64  `json:"size" bson:"size"`
	ModifiedAt int64  `json:"modified_at" bson:"modified_at" grimoire:"index"`
}

type FileService interface {
	Index(IndexRequest) FilesResponse
	Walk(KeyRequest) EmptyResponse
}

type FilesResponse struct {
	Count  int64   `json:"count"`
	Result []*File `json:"files"`
}

type EmptyResponse struct{}

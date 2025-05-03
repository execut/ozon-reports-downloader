package file

type File struct {
    content  []byte
    fileType string
}

func NewFile(content []byte, fileType string) *File {
    return &File{content: content, fileType: fileType}
}

func (f File) Content() []byte {
    return f.content
}

func (f File) FileType() string {
    return f.fileType
}

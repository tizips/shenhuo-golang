package basic

// DoUploadOfFile 文件上传响应
type DoUploadOfFile struct {
	Name string `json:"name"` // 文件名
	Uri  string `json:"uri"`  // 相对存储路径
	Url  string `json:"url"`  // 完整访问地址
}

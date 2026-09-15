package basic

// DoUploadOfFile 文件上传请求
type DoUploadOfFile struct {
	Dir string `json:"dir" form:"dir" validate:"required,max=100,dirs" label:"目录"` // 上传目标目录；仅允许合法目录标识
}

package model

type RequestMeta struct {
	AlgorithmName string `json:"algorithm_name"` // 表示算法名
	StoreLoc      string `json:"store_loc"`      // 表示临时存储位置
	FileName      string `json:"file_name"`      // 表示上传的文件名称
}

package model

type File struct {
	Name string
	Ext  string
	Path string
}

type FileConvert struct {
	Name   string
	ExtOr  string
	ExtDst string
	Path   string
}

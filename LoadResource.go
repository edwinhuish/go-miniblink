package GoMiniblink

import (
	"io"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
)

type LoadResource interface {
	Domain() string
	ByUri(uri *url2.URL) []byte
}

type FileLoader struct {
	domain    string
	val       any // dir or http.FileSystem
	isBinData bool
}

func NewFileLoaderStatic(dir string, domain string) *FileLoader {
	return &FileLoader{
		domain:    strings.ToLower(strings.TrimRight(domain, "/")),
		val:       strings.TrimRight(dir, string(os.PathSeparator)),
		isBinData: false,
	}
}

func NewFileLoaderBin(fs http.FileSystem, domain string) *FileLoader {
	return &FileLoader{
		domain:    strings.ToLower(strings.TrimRight(domain, "/")),
		val:       fs,
		isBinData: true,
	}
}

func (_this *FileLoader) Domain() string {
	return _this.domain
}

func (_this *FileLoader) ByUri(uri *url2.URL) []byte {
	if _this.isBinData {
		return _this.readBinDataByUri(uri)
	}

	return _this.readStaticFileByUri(uri)
}

func (_this *FileLoader) readStaticFileByUri(uri *url2.URL) []byte {
	path := strings.Join([]string{_this.val.(string), uri.Path}, "")
	path = strings.ReplaceAll(path, "/", string(os.PathSeparator))
	if data, err := os.ReadFile(path); err == nil {
		return data
	}
	return nil
}

func (_this *FileLoader) readBinDataByUri(uri *url2.URL) []byte {
	path := strings.ReplaceAll(uri.Path, "/", string(os.PathSeparator))
	//获取文件
	f, err := _this.val.(http.FileSystem).Open(path)
	if err != nil {
		return nil
	}

	defer f.Close()

	//读取二进制数据
	bt, err := io.ReadAll(f)
	if err != nil {
		return nil
	}

	return bt
}

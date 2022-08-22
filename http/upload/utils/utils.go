package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// 解压 tar.gz
func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		filename := filepath.Join(dest, hdr.Name)

		switch hdr.Typeflag {
		case tar.TypeDir:
			// 是目录，则创建
			if err = os.MkdirAll(filename, 0755); err != nil {
				return err
			}

		case tar.TypeReg:
			// 是普通文件，创建并将内容写入
			// 是普通文件，创建并将内容写入
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(file, tr)
			// 循环内不能用defer，先关闭文件句柄
			if err2 := file.Close(); err2 != nil {
				return err2
			}
			// 这里再对文件copy的结果做判断
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func IsWindows() bool {
	sysType := runtime.GOOS
	return sysType == "windows"
}

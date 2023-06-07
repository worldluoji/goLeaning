package localcopyworker

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"

	colorUtils "clitemplate/utils/colorUtils"
	fileutils "clitemplate/utils/fileutils"
)

//go:embed pc
var templates embed.FS

func CopyEmbededFiles(dest string) error {
	var src = "pc"
	if err := fileutils.MkDir(dest); err != nil {
		return err
	}
	return copy(src, dest)
}

func copy(src, dest string) error {
	list, err := templates.ReadDir(src)
	if err != nil {
		return fmt.Errorf("fatal error template file: %s", err)
	}
	for _, item := range list {
		fileName := item.Name()
		srcNew := src + "/" + fileName //filepath.Join(src, fileName)
		destNew := filepath.Join(dest, fileName)
		// fmt.Println(src, srcNew, destNew)
		if item.IsDir() {
			if err := fileutils.MkDir(destNew); err != nil {
				return fmt.Errorf("mkdir failed: %s", err)
			}
			copy(srcNew, destNew)
		} else {
			var f fs.File
			if f, err = templates.Open(srcNew); err != nil {
				return fmt.Errorf("open file failed: %s", err)
			}

			if _, err := fileutils.CopyFileFS(f, destNew); err != nil {
				return fmt.Errorf("copy file failed: %s", err)
			}
		}
	}
	return nil
}

type LocalCopyWorker struct {
}

func (worker *LocalCopyWorker) Do(dest string) bool {
	fmt.Println(colorUtils.White("Begin to copy template..."))
	if err := CopyEmbededFiles(dest); err != nil {
		fmt.Println(colorUtils.Red("copy template error...."), err)
		return false
	}

	fmt.Println(colorUtils.Green("Copy template successed!!!"))
	return true
}

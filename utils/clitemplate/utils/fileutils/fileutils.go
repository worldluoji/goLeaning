package fileutils

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	//desFile, err := os.Create(des)  //无法复制源文件的所有权限
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

func MkDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil { //os.ModePerm
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}

func CopyDir(srcPath, desPath string) error {
	//检查目录是否正确
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("源路径不是一个正确的目录！")
		}
	}

	if desInfo, err := os.Stat(desPath); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			return errors.New("目标路径不是一个正确的目录！")
		}
	}

	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return errors.New("源路径与目标路径不能相同！")
	}

	/** Wlak方法，从srcPath目录开始遍历，每一个子目录中的文件、文件夹都会遍历到，
		 * func中，path就是具体文件或文件夹得路径，f是对应文件的信息
		 The err argument reports an error related to path, signaling that Walk
	     will not walk into that directory. The function can decide how to
	     handle that error; as described earlier, returning the error will
	     cause Walk to stop walking the entire tree.
	*/
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		//复制目录是将源目录中的子目录复制到目标路径中，不包含源目录本身
		if path == srcPath {
			return nil
		}

		//生成新路径
		destNewPath := strings.Replace(path, srcPath, desPath, -1)

		if !f.IsDir() {
			CopyFile(path, destNewPath)
		} else {
			if !FileIsExisted(destNewPath) {
				return MkDir(destNewPath)
			}
		}

		return nil
	})

	return err
}

func GetCurrentJointDir(name string) string {
	// os.Getwd()这个方法获取的是进程在OS系统所在的目录
	current, _ := os.Getwd()
	current = filepath.Join(current, name)
	return current
}

func CopyFileFS(srcFile fs.File, des string) (written int64, err error) {
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	//desFile, err := os.Create(des)  //无法复制源文件的所有权限
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)

}

func GetDestPath(project string) string {
	var destPath string
	if IsDir(project) {
		destPath = project
		MkDir(destPath)
	} else {
		destPath = GetCurrentJointDir(project)
		MkDir(destPath)
		if !IsDir(destPath) {
			fmt.Println("dest path error....", destPath)
			return ""
		}
	}
	return destPath
}

package gitoperation

import (
	"errors"
	"fmt"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

/**
  bare: 是否开启bare选项
  directory: 需要指定directory，明确把代码clone到哪里
  url: 要clone的repo url
  referenceName: 可以是分支名称，也可以是tag，也可以是commit id
  auth: 这个是可选的，就要看你有没有授权校验
*/
func GitClone(directory, url, branch string) (*git.Repository, error) {
	if url == "" {
		return nil, errors.New("git url can not be null")
	}

	var cloneDir string
	if directory == "" {
		if dir, err := getWorkDir(); err != nil {
			return nil, err
		} else {
			cloneDir = dir
		}
	} else {
		cloneDir = directory
	}

	fmt.Println("git clone", url, cloneDir, branch)
	return git.PlainClone(cloneDir, false,
		&git.CloneOptions{
			URL: url,
			// Auth: &http.BasicAuth{
			// 	Username: "worldluoji",
			// 	Password: "Mldncsdn3",
			// },
			Progress:      os.Stdout,
			SingleBranch:  true,
			ReferenceName: plumbing.ReferenceName("refs/heads/" + branch),
		})
}

func getWorkDir() (string, error) {
	if str, err := os.Getwd(); err != nil {
		return "", err
	} else {
		return str, nil
	}
}

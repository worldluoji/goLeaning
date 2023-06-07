package gitworker

import (
	"errors"
	"fmt"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"

	colorUtils "clitemplate/utils/colorUtils"
)

/**
  bare: 是否开启bare选项
  dest: 需要指定directory，明确把代码clone到哪里
  url: 要clone的repo url
  referenceName: 可以是分支名称，也可以是tag，也可以是commit id
  auth: 这个是可选的，就要看你有没有授权校验
*/
func GitClone(dest, url, branch string) (*git.Repository, error) {
	if url == "" {
		return nil, errors.New("git url can not be null")
	}

	var cloneDir string
	if dest == "" {
		if dir, err := getWorkDir(); err != nil {
			return nil, err
		} else {
			cloneDir = dir
		}
	} else {
		cloneDir = dest
	}

	// fmt.Println("git clone", url, cloneDir, branch)
	return git.PlainClone(cloneDir, false,
		&git.CloneOptions{
			URL: url,
			// Auth: &http.BasicAuth{
			// 	Username: "worldluoji",
			// 	Password: "xxxxxxxx",
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

type GitWorker struct {
	Dest   string
	Url    string
	Branch string
}

func (worker *GitWorker) Do(dest string) bool {
	fmt.Println(colorUtils.White("Begin to get template from remote..."))
	if _, err := GitClone(dest, worker.Url, worker.Branch); err != nil {
		fmt.Println(colorUtils.Red("Get template from remote failed "), err)
		return false
	}
	fmt.Println(colorUtils.Green("Get template from remote successed!!!"))
	return true
}

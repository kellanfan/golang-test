package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //蹇界暐鍚庣紑鍖归厤鐨勫ぇ灏忓啓

	for _, fi := range dir {
		if fi.IsDir() { // 蹇界暐鐩綍
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //鍖归厤鏂囦欢
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, nil
}

// 鑾峰彇鎸囧畾鐩綍鍙婃墍鏈夊瓙鐩綍涓嬬殑鎵€鏈夋枃浠讹紝鍙互鍖归厤鍚庣紑杩囨护銆�
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //蹇界暐鍚庣紑鍖归厤鐨勫ぇ灏忓啓

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //閬嶅巻鐩綍
		//if err != nil { //蹇界暐閿欒
		// return err
		//}

		if fi.IsDir() { // 蹇界暐鐩綍
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

func FindDir(dirPth string, suffix string) (dirs []string, err error) {
	dirs = make([]string, 0, 10)
	suffix = strings.ToUpper(suffix)

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //閬嶅巻鐩綍
		//if err != nil { //蹇界暐閿欒
		// return err
		//}

		if !fi.IsDir() { // 蹇界暐鏂囦欢
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			dirs = append(dirs, filename)
		}

		return nil
	})

	return dirs, err

}

func DeleteDir(targetPath string) {
	dir, err := ioutil.ReadDir(targetPath)
	if err != nil {
		fmt.Println(err)
	}
	for _, d := range dir {
		fmt.Println("Remove %s", d)
		os.RemoveAll(path.Join([]string{targetPath, d.Name()}...))
	}
}

func main() {
	items := [3]string{"File", "Video", "Image"}
	dirs := [2]string{"C:\\Users\\Administrator\\Documents\\WeChat Files", "C:\\Users\\Administrator\\Documents\\WXWork"}
	for _, dir := range dirs {
		for _, item := range items {
			files, err := FindDir(dir, item)
			fmt.Println(files, err)
			for _, dir := range files {
				DeleteDir(dir)
			}
		}
	}

}

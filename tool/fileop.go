package tool

import "os"

func IsFileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//delete的时候删除本地文件
//未实现摸鱼中
func RemoveFile(fileName string) error {
	return os.Remove(fileName)
}

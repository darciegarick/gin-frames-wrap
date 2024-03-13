package pkg

import "os"

// PathExists 判断给定路径是否存在。
// 参数 path 是要检查的路径。
// 如果路径存在，返回 true；如果路径不存在，返回 false；如果发生错误，返回错误信息。
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

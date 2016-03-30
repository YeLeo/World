package wordcount

import (
	"fmt"
	//"github.com/yeleo/world/sortable"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var m map[interface{}]interface{}

func Do() {
	m = make(map[interface{}]interface{}, 100000)
	err := filepath.Walk("C:/Users/M/Desktop/BigText", walkFn)
	if err != nil {
		fmt.Println(err)
	}
	data := make([]string, 0, 100000)
	//for _, item := range sortable.SortableMap(m) {
	//	data = append(data, item.Key+"\t"+item.Value)
	//}
	err = ioutil.WriteFile("C:/Users/M/Desktop/BigText/Result-Go.txt", []byte(strings.Join(data, "\n")), os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

}

func walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	//文本处理可能存在顺序问题
	text := strings.Split(strings.TrimSpace(string(data)), " ")
	for _, v := range text {
		if m[v] != nil {
			m[v] = m[v].(int) + 1
		} else {
			m[v] = 1
		}
	}
	return nil
}

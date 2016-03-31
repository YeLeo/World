package wordcount

import (
	"fmt"
	//"github.com/yeleo/world/sortable"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var m map[string]int

func Do() {
	m = make(map[string]int, 100000)
	err := filepath.Walk("C:/Users/M/Desktop/BigText/", walkFn)
	if err != nil {
		fmt.Println(err)
	}
	data := make([]string, 0, 100000)
	//for _, item := range sortable.SortableMap(m) {
	//	data = append(data, item.Key+"\t"+item.Value)
	//}
	for k, v := range m {
		data = append(data, k+"\t"+strconv.Itoa(v))
	}
	err = ioutil.WriteFile("C:/Users/M/Desktop/BigText/Result-go.txt", []byte(strings.Join(data, "\n")), os.ModeAppend)
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
	text := string(data)
	text = strings.Replace(text, "\r", " ", -1)
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, "\t", " ", -1)
	text = strings.ToLower(text)
	textArray := strings.Split(text, " ")
	for _, v := range textArray {
		if m[v] != 0 {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return nil
}

package wordcount

import (
	"fmt"
	"github.com/yeleo/world/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var m map[interface{}]interface{}

func Do(root, filename string) {
	m = make(map[interface{}]interface{}, 100000)
	err := filepath.Walk(root, walkFn)
	if err != nil {
		fmt.Println(err)
	}
	data := make([]string, 0, 100000)
	sm := util.SortableMap(m, util.SortByValue, util.DESC)
	sort.Sort(sm)
	for _, item := range sm.Map {
		data = append(data, item.Key.(string)+"\t"+strconv.Itoa(item.Value.(int)))
	}
	err = ioutil.WriteFile(filename, []byte(strings.Join(data, "\n")), os.ModeAppend)
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
	err = textParse(path)
	if err != nil {
		return err
	}
	return nil
}

func textParse(path string) error {
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
		if m[v] != nil {
			m[v] = m[v].(int) + 1
		} else {
			m[v] = 1
		}
	}
	return nil
}

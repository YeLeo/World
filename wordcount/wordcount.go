package wordcount

import (
	"errors"
	"fmt"
	"github.com/yeleo/world/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const numDigesters = 10
const resultSize int = 100

type result struct {
	words []string
	err   error
}

func Do(root, filename string) {
	done := make(chan struct{})
	defer close(done)
	paths, errChan := getFilePath(done, root)
	resultChan := make(chan result, resultSize)
	var wg sync.WaitGroup
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			readFileAndParse(done, paths, resultChan)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	m := make(map[interface{}]interface{}, 100000)
	for result := range resultChan {
		if result.err != nil {
			fmt.Println(result.err)
			return
		}
		for _, word := range result.words {
			if m[word] != nil {
				m[word] = m[word].(int) + 1
			} else {
				m[word] = 1
			}
		}
	}
	if err := <-errChan; err != nil {
		fmt.Println(err)
		return
	}
	data := make([]string, 0, 100000)
	sm := util.InitMap(m, util.SortByValue, util.DESC)
	sort.Sort(sm)
	for _, item := range sm.KeyValuePairs {
		data = append(data, item.Key.(string)+"\t"+strconv.Itoa(item.Value.(int)))
	}
	err := ioutil.WriteFile(filename, []byte(strings.Join(data, "\n")), os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
}

func getFilePath(done <-chan struct{}, root string) (<-chan string, chan error) {
	paths := make(chan string)
	errChan := make(chan error, 1)
	go func() {
		defer close(paths)
		errChan <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("getFilePath:canceled!")
			}
			return nil
		})
	}()

	return paths, errChan
}

func readFileAndParse(done chan struct{}, paths <-chan string, bufferPool chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		text := string(data)
		text = strings.Replace(text, "\r", " ", -1)
		text = strings.Replace(text, "\n", " ", -1)
		text = strings.Replace(text, "\t", " ", -1)
		text = strings.ToLower(text)
		words := strings.Split(text, " ")
		select {
		case bufferPool <- result{words: words, err: err}:
		case <-done:
			return
		}
	}
}

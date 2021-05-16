package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	g := &Group{}
	g.MaxGroups(10)

	file, err := os.OpenFile("test.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	_, err = file.Stat()
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(file)
	urls := make([]string, 0, 1000)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			break
		}
		urls = append(urls, line)
	}

	for i := 0; i < 10000; i++ {
		g.Go(func(idx int) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				// Get the data
				resp, err := http.Get(urls[idx%len(urls)])
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				err = ioutil.WriteFile(fmt.Sprintf("%d.jpg", idx), data, 0644)
				return err
			}
		}(i))
	}
	err = g.Wait()
	fmt.Println("err is ", err)
}

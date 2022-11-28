package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bytesRead, _ := ioutil.ReadFile(os.Args[1])
	file_content := string(bytesRead)
	lines := strings.Split(file_content, "\n")
	lines = lines[1 : len(lines)-1]
	itemLike, itemDislike := getLikeDislikeItem(&lines)
	//remove diklike item
	for key, _ := range itemDislike {
		delete(itemLike, key)
	}
	fmt.Print(len(itemLike), " ")
	for key, _ := range itemLike {
		fmt.Print(key, " ")
	}
}
func getLikeDislikeItem(lines *[]string) (map[string]int, map[string]int) {
	itemLike := make(map[string]int)
	itemDislike := make(map[string]int)
	for index, line := range *lines {
		itemList := strings.Split(line, " ")
		if index%2 == 0 {
			for item := 1; item < len(itemList); item++ {
				if value, found := itemLike[itemList[item]]; found {
					itemLike[itemList[item]] = value + 1
				} else {
					itemLike[itemList[item]] = 1
				}
			}
		} else {
			for item := 1; item < len(itemList); item++ {
				if value, found := itemDislike[itemList[item]]; found {
					itemDislike[itemList[item]] = value + 1
				} else {
					itemDislike[itemList[item]] = 1
				}
			}

		}
	}
	return itemLike, itemDislike
}

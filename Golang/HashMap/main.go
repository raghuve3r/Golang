package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}
	//lst, _ := ioutil.ReadAll(res.Body)
	//str := string(lst)
	//fmt.Println(str)

	words := make(map[string][]string)

	d := bufio.NewScanner(res.Body)
	d.Split(bufio.ScanWords)

	for d.Scan() {
		//fmt.Println(string(d.Text()[0]))
		words[string(d.Text()[0])] = append(words[string(d.Text()[0])], d.Text())
	}
	defer res.Body.Close()
	if er := d.Err(); er != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", er)
	}
	i := 0
	for k, v := range words {
		fmt.Println("Key : ", k, "\n", "Value : ", v)
		if i == 200 {
			break
		}
		i++
	}
}

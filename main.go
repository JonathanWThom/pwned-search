package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	password := strings.Join(args, " ")
	hasher := sha1.New()
	hasher.Write([]byte(password))
	sha := hex.EncodeToString(hasher.Sum(nil))
	prefix := sha[:5]
	suffix := sha[5:]

	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + prefix)
	if err != nil {
		fmt.Printf("Error occurred while checking password: %v\n", err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), ":")
		candidate := result[0]

		if candidate == strings.ToUpper(suffix) {
			occurences := result[1]
			fmt.Printf("%s was found with %v occurences (hash: %s)\n", password, occurences, sha)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error occurred while checking password: %v\n", err)
		return
	}

	fmt.Printf("%s was not found\n", password)
}

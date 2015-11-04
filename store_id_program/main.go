package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "regexp"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    re     := regexp.MustCompile("[^/]+.rpm")

    for true {
        line, _  := reader.ReadString('\n')
        parts    := strings.Split(line, " ")
        url      := parts[0]
        filename := re.FindString(url)

        if len(filename) > 0 {
            fmt.Printf("OK store-id=%v\n", filename);
        } else {
            fmt.Printf("OK store-id=%v\n", url);
        }
    }
}

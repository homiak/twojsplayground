package main

import (
    "bufio"
    "bytes"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

var dicEncode = map[rune]rune{
    'q': 'й',
    'w': 'ц',
    'e': 'у',
    'r': 'к',
    't': 'н',
    'y': 'г',
    'u': 'ш',
    'i': 'щ',
    'o': 'з',
    'p': 'х',
    'a': 'ъ',
    's': 'ё',
    'd': 'ф',
    'f': 'ы',
    'g': 'в',
    'h': 'п',
    'j': 'р',
    'k': 'о',
    'l': 'л',
    'z': 'д',
    'x': 'ж',
    'c': 'э',
    'v': 'я',
    'b': 'ч',
    'n': 'м',
    'm': 'ю',
    //митьб
}

var dicDecode = map[rune]rune{
    'й': 'q',
    'ц': 'w',
    'у': 'e',
    'к': 'r',
    'н': 't',
    'г': 'y',
    'ш': 'u',
    'щ': 'i',
    'з': 'o',
    'х': 'p',
    'ъ': 'a',
    'ё': 's',
    'ф': 'd',
    'ы': 'f',
    'в': 'g',
    'п': 'h',
    'р': 'j',
    'о': 'k',
    'л': 'l',
    'д': 'z',
    'ж': 'x',
    'э': 'c',
    'я': 'v',
    'ч': 'b',
    'м': 'n',
    'ю': 'm',
    //митьб
}

func main() {
    var text string
    fmt.Print("Enter text: ")
    _, _ = fmt.Scan(&text)
    fmt.Print("Operation: ")
    var operation string
    _, _ = fmt.Scan(&operation)
    dic := dicEncode
    if strings.TrimSpace(operation) == "decode" {
        dic = dicDecode
    }
    buf := bytes.Buffer{}
    for _, c := range strings.TrimSpace(strings.ToLower(text)) {
        enc, ok := dic[c]
        if !ok {
            buf.WriteRune(c)
 //           fmt.Printf("unknown char in the input %v\n", c)
        } else {
            buf.WriteRune(enc)
        }
    }
    fmt.Printf("Output: %s ", buf.String())

}

func greet() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Print("Enter how many times to repeat: ")
    times, _ := reader.ReadString('\n')
    time, err := strconv.Atoi(strings.TrimSpace(times))
    if err != nil {
        log.Fatalf("Wrong number of times - %s", err)
    }
    for i := 0; i < time; i = i + 1 {
        fmt.Printf("%d Hello %s! ", i, strings.TrimSpace(text))
    }
}

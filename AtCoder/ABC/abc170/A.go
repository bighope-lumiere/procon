package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var sc = bufio.NewReaderSize(os.Stdin, 1024*1024*10)

func main() {
    x := getStdinIntArr()
    ans := 1
    for x[ans-1] != 0 {
        ans++
    }
    fmt.Println(ans)
}

// slice[start:end+1] = val
func fill(slice []int, val, start, end int) ([]int, error) {
    if len(slice) < start || len(slice) < end {
        return nil, fmt.Errorf("Error")
    }
    for i := start; i < end; i++ {
        slice[i] = val
    }
    return slice, nil
}

func getStdin() string {
    return readLine()
}
func getStdinInt() int {
    str := getStdin()
    res, _ := strconv.Atoi(str)
    return res
}
func getStdinIntArr() []int {
    str := getStdin()
    list := strings.Split(str, " ")
    res := make([]int, len(list))
    for idx, v := range list {
        res[idx], _ = strconv.Atoi(v)
    }
    return res
}
func getStdinArr() []string {
    str := getStdin()
    list := strings.Split(str, " ")
    return list
}
func getStdinIntArr64() []int64 {
    str := getStdin()
    list := strings.Split(str, " ")
    res := make([]int64, len(list))
    for idx, v := range list {
        res[idx], _ = strconv.ParseInt(v, 10, 64)
    }
    return res
}
func readLine() string {
    buf := make([]byte, 0, 0)
    for {
        l, p, e := sc.ReadLine()
        if e != nil {
            panic(e)
        }
        buf = append(buf, l...)
        if !p {
            break
        }
    }
    return string(buf)
}
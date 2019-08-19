package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Achillesxu/goProgramLanguage/util"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type currency int

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type tree struct {
	value       int
	left, right *tree
}

const (
	USD currency = iota
	EUR
	GBP
	RMB
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Value string
}

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"programming languages": {"data structures", "computer organization"},
}

type address struct {
	street     string
	city       string
	state      string
	postalCode string
}

type subscriber struct {
	name        string
	rate        string
	active      string
	HomeAddress address
}

func main() {
	//replaceStrMain()
	//arraySlice()
	//unicodePrint()
	//strSortPrint()
	//mapPrint()
	//structPrint()
	//jsonPrint()
	//f := squares()
	//fmt.Println(f())
	//fmt.Println(f())
	//topoSortPrint()
	//testPoint()
	//clockServer()
	//pipeline()
	//ch := make(chan string)
	//go sendFunc(ch)
	//go getFunc(ch)
	//time.Sleep(2e9)
	//fmt.Println("Enter a grade")
	//grade, err := getFloat()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var status string
	//if grade >= 60 {
	//	status = "passing"
	//} else {
	//	status = "failing"
	//}
	//fmt.Println("A grade of ", grade, "is ", status)
	//var aaa = [3]int{1, 2, 3}
	//fmt.Println(aaa)
	//slice := []string{"a", "b"}
	//fmt.Println(slice, len(slice), cap(slice))
	//slice = append(slice, "d")
	//fmt.Println(slice, len(slice), cap(slice))
	//slice = append(slice, "d")
	//fmt.Println(slice, len(slice), cap(slice))
	//var intSlice []int
	//var strSlice []string
	//fmt.Printf("%#v, %#v", intSlice, strSlice)
	//files, err := ioutil.ReadDir("/Users/achilles_xushy/GoLandProjects/src/github.com/Achillesxu/goProgramLanguage")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, file := range files {
	//	if file.IsDir() {
	//		fmt.Println("dir:", file.Name())
	//	} else {
	//		fmt.Println("file: ", file.Name())
	//	}
	//}
	resp, err := http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(body))
	}
}

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func sendFunc(chSend chan string) {
	chSend <- "hello world1"
	chSend <- "hello world2"
	chSend <- "hello world3"
	chSend <- "hello world4"
	chSend <- "hello world5"
}

func getFunc(chGet chan string) {
	var getVal string
	for {
		getVal = <-chGet
		fmt.Println(getVal)

	}
}

func pipeline() {
	nat := make(chan int)
	squ := make(chan int)

	go func() {
		for x := 0; ; x++ {
			nat <- x
		}
	}()

	go func() {
		for {
			x := <-nat
			squ <- x * x
		}
	}()

	for {
		fmt.Println(<-squ)
	}
}

func clockServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func clientFunc() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func spFib() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func testPoint() {
	p := util.PointF{1.0, 2}
	q := util.PointF{4.0, 6}
	fmt.Println(util.Distance(p, q))
	fmt.Println(p.Distance(q))
}

func topoSortPrint() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func parseHtml() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func findLinksPrint() {
	for _, url := range os.Args[:1] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML :%v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, n)
	}
	return links
}

func jsonPrint() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func replaceStrMain() {
	searchIn := "John 2567.34 William 3564.43 Steve 2345.12"
	pat := "[0-9]+.[0-9]+"
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 3, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found")
	}
	re, _ := regexp.Compile(pat)

	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}

func arraySlice() {
	months := []string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2, len(Q2), cap(Q2), summer, cap(summer))

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	//endlessSummer := summer[:5]
	//fmt.Println(endlessSummer)
	//a := []int{1,2,3,4,5}
	//reverse(a)
	//fmt.Println(a)
}

func unicodePrint() {
	var runes []rune

	for _, r := range "hello, 徐世寅" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	var z []int
	fmt.Println(len(z), cap(z))
}

func strSortPrint() {
	names := []string{"achilles", "xushiyin", "workaholic"}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name)
	}
}

func mapPrint() {
	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)
	fmt.Println(mapEqual(map[string]int{"B": 42}, map[string]int{"B": 42}))
}

func structPrint() {
	p := Point{2, 1}
	q := Point{1, 2}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)

	var wh Wheel
	wh = Wheel{
		Circle: Circle{
			Point: Point{
				X: 8, Y: 9,
			},
			Radius: 5,
		},
		Spokes: 30,
	}
	fmt.Printf("%#v\n", wh)
	wh.X = 42
	fmt.Printf("%#v\n", wh)

	wh.Circle.Point.Y = 89
	fmt.Printf("%v\n", wh)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	fmt.Println(len(z), cap(z))
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func nonEmpty(strings []string) []string {
	var out []string
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func treeSort(values []int) {
	var root *tree
	for _, v := range values {
		root = addNode(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func addNode(r *tree, v int) *tree {
	if r == nil {
		r = new(tree)
		r.value = v
	}
	if v < r.value {
		r.left = addNode(r.left, v)
	} else {
		r.right = addNode(r.right, v)
	}
	return r
}

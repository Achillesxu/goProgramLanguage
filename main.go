package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
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

func main() {
	//replaceStrMain()
	//arraySlice()
	//unicodePrint()
	//strSortPrint()
	//mapPrint()
	//structPrint()
	jsonPrint()
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


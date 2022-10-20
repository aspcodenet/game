package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"main/concurrency"
	"math/rand"
	"os"
)

type Human struct {
	Namn  string
	Level int
}

type Fly struct {
	Namn  string
	Level int
}

type IActionable interface {
	act()
}

type ILevelCalculator interface {
	updateLevel()
}

type IActionableLevelCalculator interface {
	IActionable
	ILevelCalculator
}

func (t *Human) updateLevel() {
	t.Level++
}

func (t *Human) act() {
	var actions = [...]string{"rapar", "smaskar", "tänker"}
	var action = actions[rand.Intn(len(actions))]
	fmt.Printf("%s %s\n", t.Namn, action)
}

func (t *Fly) act() {
	var actions = [...]string{"surrar", "flyger", "är irriterande"}
	var action = actions[rand.Intn(len(actions))]
	fmt.Printf("%s %s\n", t.Namn, action)
}

func (t *Fly) updateLevel() {
	t.Level++
}

func gameLoop(t interface{}) {
	v := t
	switch v.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v.(int)*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v.(string)))
	case Human:
		fmt.Printf("%q is %v bytes long\n", v, len(v.(string)))
	case *Human:
		fmt.Printf("%q is %v bytes long\n", v, len(v.(string)))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}

// func Withdraw(kontoNr string, belopp int) *VADVIVIL!  {
// }

type alphaReader string

func (a alphaReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] >= 'A' && a[i] <= 'Z' {
			p[count] = a[i]
			count++
		}
	}
	return len(a), nil
}

func main() {
	concurrency.Run3()
	return
	enumtest()
	json11()
	var s alphaReader
	s = "Hello, World!"
	p := make([]byte, len(s))
	if _, err := io.ReadFull(s, p); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", string(p))
	}

	// var allGuests []interface{}
	// p := &Human{Namn: "Stefan"}
	// p2 := &Human{Namn: "Kerstin"}
	// flugan := &Fly{Namn: "Flugan"}
	// allGuests = append(allGuests, p, p2, flugan, 13)
	// for {
	// 	for _, guest := range allGuests {
	// 		gameLoop(guest)
	// 	}
	// }

}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}
type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}

func readLines(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)

}

func json11() {
	// response, err := http.Get("https://fakestoreapi.com/products")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	responseData, _ := readLines("./a.txt")

	// responseData, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var listan []Product
	json.Unmarshal(responseData, &listan)
	for _, p := range listan {
		fmt.Println(p.Title)
	}

}

type Season int64

const (
	Summer Season = iota
	Autumn
	Winter
	Spring
)

func PrintWeather(season Season) {
	switch season {
	case Summer:
		fmt.Println("Varmt")
	case Autumn:
		fmt.Println("Kallt")
	case Winter:
		fmt.Println("Svinkallt")
	case Spring:
		fmt.Println("Bättre igen")
	}
}

func enumtest() {
	PrintWeather(Autumn)

	PrintWeather(1222)

}

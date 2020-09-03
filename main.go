package main

import (
    "fmt"
    "os"
    "bufio"
    "github.com/antchfx/htmlquery"
    "strconv"
    "strings"
    "log"
    "time"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Ticker : ")
    ticker, err := reader.ReadString('\n')
    if err != nil {
      panic(err)
    }
    tickerWithoutNewLine := strings.TrimSuffix(ticker, "\n")
    fmt.Println(GetTickerPrice(tickerWithoutNewLine))
    time.Sleep(5 * time.Second)
}

func GetTickerPrice(ticker string) float64 {
  doc, err := htmlquery.LoadURL("https://finance.yahoo.com/quote/" + ticker)
	if err != nil {
		log.Fatal("Ticker not found")
	}
	// Find all news item.
	list, err := htmlquery.Query(doc, "//span[@class='Trsdu(0.3s) Fw(b) Fz(36px) Mb(-4px) D(ib)']")
  if err != nil {
    log.Fatal("Wrong path, contact admin")
  }
  if list == nil {
    log.Fatal("No ticker found")
  }
	price := htmlquery.InnerText(list)

  priceFloat, err := strconv.ParseFloat(price, 64)
  if err != nil {
    fmt.Println("Couldn't convert to float")
  }
  return priceFloat
}

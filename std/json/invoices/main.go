package main

import (
  "fmt"
  // "io"
  "encoding/json"
)

type Invoice struct {
  Id int
  Name string
  Items []Item
}

type Item struct {
  Id int
  Name string
  Price float64
}

// func MarshalInvoices(writer io.Writer, invoices []Invoice) error {
//   encoder := json.NewEncoder(writer)
//   return encoder.Encode(invoices)
// }

func main() {
  invoices := make([]Invoice, 0)
  invoice := Invoice{Id: 0, Name: "BootsInvoice"}
  invoice.Items = []Item{
    Item{0, "Nike", 4500},
    Item{1, "Addidas", 2100},
  }
  invoices = append(invoices, invoice)

  b, _ := json.Marshal(invoices)

  fmt.Println(string(b))
}

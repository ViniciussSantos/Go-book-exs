package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var itemsTable = template.Must(template.New("itemsTable").Parse(`<table>
<tr style='text-align: left'>
<th>Item</th>
<th>Price</th>
</tr>

{{range $item, $price := .}}
<tr>
<td>{{$item}}</td>
<td>{{$price}}
</tr>
{{end}}
</table>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	itemsTable.Execute(w, db)
}

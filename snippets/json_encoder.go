package snippets

import (
	"encoding/json"
	"fmt"
)

func JsonEncoder() {
	r := request{
		Name: "test name",
		Headers: []header{
			{
				Key:          "key1",
				Value:        []byte("value 1"),
				Json:         []byte(`{"key":"name","value":"sakthi"}`),
				Names:        []string{"n1", "n2"},
				Ptr:          nil,
				AlwaysIgnore: 123,
				Magic:        100,
			},
		},
	}

	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println(">>>", err)
		return
	}
	fmt.Printf(">>request=%+v", string(b))
}

type request struct {
	Name    string   `json:"name"`
	Headers []header `json:"headers"`
}

type header struct {
	Key          string   `json:"key"`
	Value        []byte   `json:"value"`
	Json         []byte   `json:"json"`
	Names        []string `json:"names"`
	Ptr          *string  `json:"ptr"`
	AlwaysIgnore int      `json:"-"`
	Magic        int      `json:"-,"`
}

package fetcher

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Status struct {
	Verified  bool `json:"verified"`
	SentCount int  `json:"sentCount"`
}

type Fact struct {
	Id        string `json:"_id"`
	User      string `json:"user"`
	Text      string `json:"text"`
	Version   int    `json:"__v"`
	Source    string `json:"source"`
	Type      string `json:"type"`
	Deleted   bool   `json:"deleted"`
	Used      bool   `json:"used"`
	Status    Status `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (fact *Fact) PrintText() {
	fmt.Println(fact.Text)
}

func GetCatFacts() []Fact {
	resp, err := http.Get("https://cat-fact.herokuapp.com/facts")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	//Create a variable of the same type as our model
	var factResponse []Fact
	//Decode the data
	err = json.NewDecoder(resp.Body).Decode(&factResponse)
	if err != nil {
		log.Fatalf("Oops! an error occurred, please try again %v", err)
	}

	return factResponse
}

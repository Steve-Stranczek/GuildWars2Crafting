package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func responseGet(w http.ResponseWriter, r *http.Request){

	var resp Item

	err := json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	fmt.Fprintf(w,"Repsponse: %+v", resp)
}

func main() {

	bearer := os.Args[1];
	item := os.Args[2];

	requestUrl := "https://api.guildwars2.com/v2/items/" + item;
	fmt.Println("Request url = " + requestUrl);
    req, err := http.NewRequest("GET", requestUrl, nil)
	req.Header.Add("Authorization", bearer);

	client := &http.Client{}
    resp, err := client.Do(req)

	var respDeserialized Item

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(resp.Body)

	responseAsString := string(responseData)

	empResp := []byte(responseAsString)

	json.Unmarshal(empResp, &respDeserialized)

    if err != nil {
        log.Fatal(err)
    }

	fmt.Println((respDeserialized.Rarity))

}

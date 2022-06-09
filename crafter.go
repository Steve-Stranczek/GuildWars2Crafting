package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getItemCount(w http.ResponseWriter, req *http.Request) {

	if req.Body == http.NoBody {
		http.Error(w, "Please send a request body", 400)
		return
	} 
	
	var itemreq itemRequest

	err := json.NewDecoder(req.Body).Decode(&itemreq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("bearer: %s\n", itemreq.Bearer)
	fmt.Printf("item: %d\n", itemreq.Item)

	countOfItem := findItem("Bearer " + itemreq.Bearer, itemreq.Item);
	fmt.Fprintf(w, "%d", countOfItem)

}

func findItem(bearer string, item int) int {
	fmt.Println("Inside findItem")

	requestUrl := "https://api.guildwars2.com/v2/account/materials";
	fmt.Println("Request url = " + requestUrl);

    req, err := http.NewRequest("GET", requestUrl, nil)
	req.Header.Add("Authorization", bearer);

	client := &http.Client{}
    resp, err := client.Do(req)

	var respDeserialized materialStorage

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

	itemCount := 0
	
	for i:= 0; i < len(respDeserialized); i++{
		if(respDeserialized[i].ID == item){
			itemCount = respDeserialized[i].Count
		}
	}

	return itemCount
}

func main() {

	http.HandleFunc("/getItemCount", getItemCount)
	err := http.ListenAndServe("localhost:12345", nil)
	 
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}

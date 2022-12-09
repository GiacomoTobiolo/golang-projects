package main

import ( 
	"encoding/json" // for json
	"fmt"           // for printing
	"io/ioutil"     // for reading the website
	"log"           //logging
	"net/http"      //importing http
	"strconv"       //for converting string to int
	"strings"       //for strings.Join
	"text/template" // for the html
)

type StringArtists struct { 
	ID       string              
	Image    string              
	Name     string              
	Members  string              
	Creation string              
	First    string              
	Relation map[string][]string 
} 


type Artists struct { 
	ID       int      `json:"id"`          
	Image    string   `json:"image"`       
	Name     string   `json:"name"`         
	Members  []string `json:"members"`      
	Creation int      `json:"creationDate"` 
	First    string   `json:"firstAlbum"`
	Relation map[string][]string
} 

type Relations struct {
	Relation map[string][]string `json:"datesLocations"` // 
} 


var artistid []Artists 


var m []string 
var nbr int    


func Unmarshal() error { 

	names, err := http.Get("https://groupietrackers.herokuapp.com/api/artists") 
	if err != nil {            
		log.Fatal("Error happened in http.Get. Err: ", err)                                                  
	
	} 
	bytes, err := ioutil.ReadAll(names.Body) 
	if err != nil {                          
		log.Fatal("error")
	} 

	err = json.Unmarshal(bytes, &artistid) 
	if err != nil {            
		log.Fatal("Error  ", err)            
	} 
	return nil 
} 

func FindByName(input string) error { 
	x := 0 
	for x <= len(artistid) { 
		if input == artistid[x].Name { 
			nbr = x 
			break   
		}
		x++ 
	}
	nbr = x 
	return nil
}



func Page(w http.ResponseWriter, r *http.Request) { 
	if r.URL.Path != "/" { 
		http.Error(w, "404 not found.", http.StatusNotFound) 
		return                                               
	} 

	for x := range artistid { 
		m = append(m, artistid[x].Name ) 
	}
	JoinArtis := Artists{ 
		Name: (strings.Join(m, ", ")), 
	} 
	t, err := template.ParseFiles("templates/main.html") 
	if err != nil {                                      
		log.Fatalf("Error happened in template.ParseFiles. Err: %s", err) 
	} 
	err = t.Execute(w, JoinArtis) 
	if err != nil {         
		log.Fatalf("Error happened in t.Execute. Err: %s", err) 
	} 

} 



func SoloPage(w http.ResponseWriter, r *http.Request) { 
	if r.FormValue("NAME") == "" {
		http.Error(w, "400 error.", http.StatusNotFound)
		return
	}

	switch r.Method { 
	case "GET": 
		http.ServeFile(w, r, "templates/solo.html") 
	case "POST": 
		if r.FormValue("NAME") != "" { 
			err := FindByName(r.FormValue("NAME")) 
			if err != nil {
				http.Error(w, "not found.", http.StatusNotFound)
				return
			}

		}

		JoinArtis := StringArtists{ // 
			               
			Image:    (artistid[nbr].Image),                                  
			Name:     ("Name: " + artistid[nbr].Name),   
			Members:  ("Members: " + strings.Join(artistid[nbr].Members, ", ")), 
			Creation: ("Creation: " + strconv.Itoa(artistid[nbr].Creation)),
			First:    ("First album: " + artistid[nbr].First),
	        Relation: artistid[nbr].Relation,
		}

		t, err := template.ParseFiles("templates/solo.html") 
		if err != nil {                                          
			log.Fatalf("Error happened in template.ParseFiles. Err: %s", err) 
		}
		if JoinArtis.Relation == nil { 
			JoinArtis.Relation = make(map[string][]string) 
		} 
		err = t.Execute(w, JoinArtis) 
		if err != nil {   
			log.Fatalf("Error happened in t.Execute. Err: %s", err) 
		}
	}
} 

func main() { 

	err := Unmarshal() 

	if err != nil { 
		log.Fatal("Error happened in Unmarshal.", err) 
	} 

	port := ":8070" 

	http.HandleFunc("/", Page)                     
	http.HandleFunc("/solo", SoloPage)

	fmt.Print("Server running at localhost:8070", port) 

	err = http.ListenAndServe(port, nil) 
	if err != nil {                      
		log.Fatal("Error happened in http.ListenAndServe", err) 
	} 
}





package router_test
import (
	"fmt"
	"gopkg.in/orivil/router.v0"
	"log"
)

func ExampleRouter() {
	// 1. new router
	r := router.NewRouter()
	id := 1

	// 2. add route
	returnID, err := r.Add("GET/user/::country/::age", id)
	if err != nil {
		log.Fatal(err)
	}

	// check if the route path was registered
	if returnID != 0 {
		// old id will be replaced with new id
		fmt.Printf("route %d was covered by %d\n", returnID, id)
	}

	// 3. match route
	if _id, params, ok := r.Match("GET/user/china/18"); ok {
		country := params["country"]
		age := params["age"]
		fmt.Println(_id == id && country == "china" && age == "18")
	}

	// Output:
	// true
}



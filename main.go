package main

import (
  "os"
  "github.com/go-martini/martini"
  "github.com/yanatan16/golang-instagram/instagram"
  //"./instagram"
  "encoding/json"
  //"fmt"
  "net/url"
)

func main() {
  /*
     A martini.Classic() instance automatically
     serves static files from the "public"
     directory in the root of your server.
  */
  m := martini.Classic()

  m.Get("/hello", func() string {
    return "world!"
  })

  m.Get("/panic", func() string {
    panic("Hey I am in PANIC! NOW!")
  })


  m.Get("/checkInsta", func() string {
    instaClienKey := "ec769078ea8a4872b2dbf0c403168939";
    // unauthorized api client pass only clientId
    // for auth need pass second param - token
    insta := instagram.New(instaClienKey, "")

    response, err := insta.GetMediaPopular(nil);
    if err != nil {
  		panic(err)
  	}
  	//fmt.Println("Successfully created instagram.Api without user credentials")
    res1B, _ := json.Marshal(response)
    return string(res1B)
      //return json.Marshal(response);
  })

  m.Get("/checkMyInsta", func() string {
    // insta
    instaClienKey := "ec769078ea8a4872b2dbf0c403168939"
    // unauthorized api client pass only clientId
    // for auth need pass second param - token
    insta := instagram.New(instaClienKey, "")

    // user id not the same as username
    // lookup user id by name here http://jelled.com/instagram/lookup-user-id#
    myId := "263020745"
    response, err := insta.GetUserRecentMedia(myId, nil);
    /*if err != nil {
  		panic(err)
  	}*/
  	//fmt.Println("Successfully created instagram.Api without user credentials")
    res1B, _ := json.Marshal(response)
    return string(res1B)
      //return json.Marshal(response);
  })

  m.Get("/checkPhilippinesInsta", func() string {
    // insta
    instaClienKey := "ec769078ea8a4872b2dbf0c403168939"
    // unauthorized api client pass only clientId
    // for auth need pass second param - token
    insta := instagram.New(instaClienKey, "")

    // user id not the same as username
    // lookup user id by name here http://jelled.com/instagram/lookup-user-id#
    id := "1651025335"
    response, err := insta.GetUserRecentMedia(id, nil);
    if err != nil {
      panic(err)
    }
    //fmt.Println("Successfully created instagram.Api without user credentials")
    res1B, _ := json.Marshal(response)
    return string(res1B)
      //return json.Marshal(response);
  })

  m.Get("/instaSearch/:name", func(params martini.Params) string {
    searchName := params["name"]

    iparams := url.Values{}
  	iparams.Set("count", "5") // Get 5 users
  	iparams.Set("q", searchName)  // Search for user

    // insta
    instaClienKey := "ec769078ea8a4872b2dbf0c403168939"
    // unauthorized api client pass only clientId
    // for auth need pass second param - token
    insta := instagram.New(instaClienKey, "")

    response, err := insta.GetUserSearch(iparams);
    if err != nil {
  		panic(err)
  	}
  	//fmt.Println("Successfully created instagram.Api without user credentials")
    res1B, _ := json.Marshal(response)
    return string(res1B)
      //return json.Marshal(response);
  })


  /* Martini's Run function looks for
     the PORT and HOST environment variables
     and uses those. Otherwise Martini will
     default to localhost:3000.*/
  // m.Run()

  // Config
	// ----------------------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

  m.RunOnAddr(":" + port)
}

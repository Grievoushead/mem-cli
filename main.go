package main

import (
  "os"
  "github.com/go-martini/martini"
  "github.com/yanatan16/golang-instagram/instagram"
  "encoding/json"
  //"fmt"
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

  m.Get("/panic", func() string {
    panic("Hey I am in PANIC! NOW!")
  })

  m.Get("/checkMyInsta", func() string {
    instaClienKey := "ec769078ea8a4872b2dbf0c403168939";
    // unauthorized api client pass only clientId
    // for auth need pass second param - token
    insta := instagram.New(instaClienKey, "")

    response, err := insta.GetUserRecentMedia("comeseephilippines", nil);
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

package main

import (
  "os"
  "github.com/go-martini/martini"
  "github.com/yanatan16/golang-instagram/instagram"
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


  // Instagram API
  instaClienKey := "ec769078ea8a4872b2dbf0c403168939";
  // unauthorized api client pass only clientId
  // for auth need pass second param - token
  instaApi := instagram.New(instaClienKey, "")

  m.Get("/insta/pop", func() string {
    res, err := instaApi.GetMediaPopular(nil);
    if err != nil {
  		panic(err)
  	}

    resJson, _ := json.Marshal(res)

    return string(resJson)
  })

  m.Get("/insta/my", func() string {
    // user id not the same as username
    // saerch user id by name
    myId := "263020745"
    res, err := instaApi.GetUserRecentMedia(myId, nil);
    if err != nil {
  		panic(err)
  	}

    resJson, _ := json.Marshal(res)

    return string(resJson)
  })

  m.Get("/insta/philippines", func() string {
    uid := "1651025335"
    res, err := instaApi.GetUserRecentMedia(uid, nil);
    if err != nil {
      panic(err)
    }
    resJson, _ := json.Marshal(res)

    return string(resJson)
  })

  m.Get("/insta/search/:name", func(params martini.Params) string {
    searchName := params["name"]

    instaParams := url.Values{}
    instaParams.Set("count", "5") // Get 5 users
    instaParams.Set("q", searchName)  // Search for user

    res, err := instaApi.GetUserSearch(instaParams);
    if err != nil {
  		panic(err)
  	}

    resJson, _ := json.Marshal(res)

    return string(resJson)
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

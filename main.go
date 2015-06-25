package main

import (
  "os"
  "github.com/go-martini/martini"
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

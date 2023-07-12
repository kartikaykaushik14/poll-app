package main

import (
	"fmt"
	"net/http"
	handler "poll-app/handlers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func Routes(router *httprouter.Router) {

	// router.GET("/api/v1/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 	fmt.Fprint(w, "Hello, World!")
	// })

	router.POST("/api/v1/user", handler.CreateUser)
	router.GET("/api/v1/user/:username/:password", handler.GetUser)
	router.GET("/api/v1/polls", handler.GetPolls)
	router.POST("/api/v1/polls", handler.CreatePolls)
	router.GET("/api/v1/votes/:pollOptionId", handler.GetVotes)
	router.POST("/api/v1/votes/", handler.CreateVotes)
	router.GET("/api/v1/users/:userId", handler.GetUserById)

	// router.GET("/api/v1/options/:pollOptionId", GetVotesForOption)
	// router.GET("/api/v1/polls/:pollId", handler.GetPollForID)
}

func main() {
	router := httprouter.New()

	// Create a CORS middleware handler
	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set the Access-Control-Allow-Origin header to allow requests from localhost:3000
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Handle preflight OPTIONS request
			if r.Method == http.MethodOptions {
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	Routes(router)

	err := http.ListenAndServe(":8080", corsHandler(router))
	if err != nil {
		fmt.Println("Unable to start server")
	}
}

// func initDatabase() {
// 	client, client_err := ent.Open("postgres", "host=localhost port=5432 user=kartikay.kaushik dbname=kartikay.kaushik sslmode=disable")
// 	if client_err != nil {
// 		log.Fatalf("Failed connecting to postgres: %v", client_err)
// 	} else {
// 		fmt.Println("Connection Succesful!")
// 	}
// 	defer client.Close()
// 	return client, client_err
// }

// func main() {
// 	connectDatabase()
// }

// func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
// 	u, err := client.User.
// 		Create().
// 		Save(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed creating user: %w", err)
// 	}
// 	log.Println("user was created: ", u)
// 	return u, nil
// }

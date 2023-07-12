package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"poll-app/database"
	"poll-app/ent/user"
	"poll-app/ent/vote"
	"strconv"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Parse the request body
	user := new(struct {
		FirstName string
		LastName  string
		Username  string
		Password  string
	})

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error Parsing Input")
		return
	}

	//Save to the database
	createdUser, err := database.EntClient.User.
		Create().
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetUsername(user.Username).
		SetPassword(user.Password).
		Save(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Unable to create user")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("username")
	password := ps.ByName("password")

	// Fetch the votes from the database
	user, err := database.EntClient.User.
		Query().
		Where(user.UsernameEQ(username), user.PasswordEQ(password)).
		All(r.Context())

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	//Handle encoding errors. - static check for linting and static code analysis.
}

func GetUserById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userIdStr := ps.ByName("userId")
	userId, err := strconv.Atoi(userIdStr)
	// Fetch the votes from the database
	user, err := database.EntClient.User.
		Query().
		Where(user.IDEQ(userId)).
		All(r.Context())

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func GetPolls(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	getpolls, err := database.EntClient.Poll.
		Query().
		WithPollOptions().
		All(context.Background())
	if err != nil {
		http.Error(w, "No Polls Found", http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getpolls)
}

func CreateVotes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Parse the request body
	vote := new(struct {
		PollOptionId string
		UserId       string
	})

	if err := json.NewDecoder(r.Body).Decode(&vote); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error Parsing Input")
		return
	}

	//Save to the database
	createdUser, err := database.EntClient.Vote.
		Create().
		SetPollOptionId(vote.PollOptionId).
		SetUserId(vote.UserId).
		Save(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Unable to create user")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func GetVotes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pollOptionId := ps.ByName("pollOptionId")

	// Fetch the votes from the database
	votes, err := database.EntClient.Vote.
		Query().
		Where(vote.PollOptionId(pollOptionId)).
		All(r.Context())

	if err != nil {
		http.Error(w, "Votes not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(votes)
}

func CreatePolls(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Parse the request body
	poll := new(struct {
		Question string   `json:"question"`
		Options  []string `json:"options"`
	})

	if err := json.NewDecoder(r.Body).Decode(&poll); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error Parsing Input")
		return
	}

	//Save to the database
	createdPoll, err := database.EntClient.Poll.
		Create().
		SetQuestion(poll.Question).
		Save(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Unable to create poll")
		return
	}

	for _, option := range poll.Options {
		createdOption, err := database.EntClient.PollOption.
			Create().
			SetOption(option).
			Save(r.Context())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Unable to create option")
			return
		}

		// Associate the created option with the poll
		_, err = createdPoll.Update().
			AddPollOptions(createdOption).
			Save(r.Context())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Unable to associate option with poll")
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPoll)
}

// func VoteForOption(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	optionID := ps.ByName("optionID")
// 	userID := ps.ByName("userID")

// 	// Check if the option exists
// 	option, err := database.EntClient.PollOption.
// 		Query().
// 		Where(polloption.ID(optionID)).
// 		Only(r.Context())

// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Option not found")
// 		return
// 	}

// 	// Check if the user exists
// 	user, err := database.EntClient.User.
// 		Query().
// 		Where(user.ID(userID)).
// 		Only(r.Context())

// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("User not found")
// 		return
// 	}

// 	// Create a new entry in the user_votes table
// 	_, err = database.EntClient.UserVotes.
// 		Create().
// 		SetOption(option).
// 		SetUser(user).
// 		Save(r.Context())

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode("Unable to vote")
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("Vote recorded")
// }

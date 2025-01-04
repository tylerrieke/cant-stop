package main

import (
	"cant-stop/models"
	"cant-stop/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"strconv"
)

var games = map[string]models.Game{}

func main() {
	r := mux.NewRouter()

	games["game"] = models.Game{
		CurrentPlayerIndex: 0,
		Players:            makePlayers(1),
		CurrentTurnColumns: make(map[int]int),
		TurnPhase:          models.GAME_OVER,
		Dice:               []int{},
		Options:            make([][][]int, 0),
	}

	r.HandleFunc("/hello-world", helloWorld)
	r.HandleFunc("/game", getGame)
	r.HandleFunc("/game/start", startGame).Methods("POST")
	r.HandleFunc("/game/roll", roll).Methods("POST")
	r.HandleFunc("/game/select", selectOption).Methods("POST")
	r.HandleFunc("/game/stop", stop).Methods("POST")
	r.HandleFunc("/game/player", switchPlayers).Methods("POST")

	// Solves Cross Origin Access Issue
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	handler := c.Handler(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + os.Getenv("PORT"),
	}

	log.Fatal(srv.ListenAndServe())
}

func getGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, game := range games {
		jsonBytes, err := utils.StructToJSON(game)
		if err != nil {
			fmt.Print(err)
		}
		w.Write(jsonBytes)
		return // Exit the loop after the first iteration
	}
	http.Error(w, "Game not found", http.StatusNotFound)
	return
}

func makePlayers(count int) []models.Player {
	var players = make([]models.Player, count)
	for i := 0; i < count; i++ {
		players[i] = models.Player{
			Color:        models.PlayerColors[i],
			ColumnValues: map[int]int{},
		}
	}
	return players
}

func startGame(w http.ResponseWriter, r *http.Request) {
	playerCount, err := strconv.Atoi(r.URL.Query().Get("numPlayers"))
	if err != nil {
		// ... handle error
		panic(err)
	}
	games["game"] = models.Game{
		CurrentPlayerIndex: 0,
		Players:            makePlayers(playerCount),
		CurrentTurnColumns: make(map[int]int),
		TurnPhase:          models.TURN_START,
		Dice:               []int{6, 6, 6, 6},
		Options:            make([][][]int, 0),
	}
	w.Header().Set("Content-Type", "application/json")
	for _, game := range games {
		jsonBytes, err := utils.StructToJSON(game)
		if err != nil {
			fmt.Print(err)
		}
		w.Write(jsonBytes)
		return
	}
}

func roll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for i, game := range games {
		games[i] = *game.Roll()
		jsonBytes, err := utils.StructToJSON(games[i])
		if err != nil {
			fmt.Print(err)
		}
		w.Write(jsonBytes)
		return // Exit the loop after the first iteration
	}
	http.Error(w, "Game not found", http.StatusNotFound)
	return
}

func selectOption(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var option [][]int

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&option)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, game := range games {
		games[i] = *game.SelectOption(option)
		jsonBytes, err := utils.StructToJSON(games[i])
		if err != nil {
			fmt.Print(err)
		}
		w.Write(jsonBytes)
		return // Exit the loop after the first iteration
	}

	http.Error(w, "Game not found", http.StatusNotFound)
	return
}

func stop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for i, game := range games {
		games[i] = *game.Stop()
		jsonBytes, err := utils.StructToJSON(games[i])
		if err != nil {
			fmt.Print(err)
		}
		w.Write(jsonBytes)
		return // Exit the loop after the first iteration
	}
	http.Error(w, "Game not found", http.StatusNotFound)
	return
}

func switchPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for i, game := range games {
		games[i] = *game.NextPlayerTurn()
		jsonBytes, err := utils.StructToJSON(games[i])
		if err != nil {
			fmt.Print(err)
		}
		w.Write(jsonBytes)
		return // Exit the loop after the first iteration
	}
	http.Error(w, "Game not found", http.StatusNotFound)
	return
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Can't Stop",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}

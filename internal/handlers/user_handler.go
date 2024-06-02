package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nafnaufal/roadmap-forum/internal/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Name     string `json:"name"`
		Bio      string `json:"bio"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	session := db.GetDriver().NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err = session.Run("CREATE (n:User {name: $name, bio: $bio, email: $email, password: $password, time: $time}) RETURN n", map[string]interface{}{
		"name":     user.Name,
		"bio":      user.Bio,
		"email":    user.Email,
		"password": user.Password,
		"time":     time.Now().String(),
	})
	if err != nil {
		http.Error(w, "Failed to create User", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s created", user.Name)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	session := db.GetDriver().NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run("MATCH (n:User) RETURN n.name", nil)
	if err != nil {
		http.Error(w, "Failed to retrieve Users", http.StatusInternalServerError)
		return
	}

	var users []string
	for result.Next() {
		record := result.Record()
		name, _ := record.Get("n.name")
		users = append(users, name.(string))
	}

	if err = result.Err(); err != nil {
		http.Error(w, "Result error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

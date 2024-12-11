package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

var numeroGenerato int
var count int

func main() {
	// Inizializza il numero casuale
	rand.Seed(int64(rand.Intn(100)))
	numeroGenerato = rand.Intn(100)
	count = 0

	// Imposta il routing
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tentativo", tentativoHandler)

	// Avvia il server
	fmt.Println("Il gioco è avviato! Vai su http://localhost:8080 per iniziare.")
	http.ListenAndServe(":8080", nil)
}

// Home Handler che mostra il form iniziale
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("home").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Gioco del Numero</title>
	</head>
	<body>
		<h1>Indovina il Numero!</h1>
		<p>Prova a indovinare il numero tra 0 e 99.</p>
		<form action="/tentativo" method="post">
			<label for="tentativo">Inserisci un numero:</label>
			<input type="number" id="tentativo" name="tentativo" required>
			<button type="submit">Invia</button>
		</form>
	</body>
	</html>
	`)
	tmpl.Execute(w, nil)
}

// Tentativo Handler che gestisce il gioco
func tentativoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tentativo, err := strconv.Atoi(r.FormValue("tentativo"))
	if err != nil {
		http.Error(w, "Inserisci un valore valido.", http.StatusBadRequest)
		return
	}

	count++

	var messaggio string
	if tentativo < numeroGenerato {
		messaggio = "Troppo basso! Riprova!"
	} else if tentativo > numeroGenerato {
		messaggio = "Troppo alto! Riprova!"
	} else {
		messaggio = fmt.Sprintf("Hai indovinato in %d tentativi!", count)
		numeroGenerato = rand.Intn(100) // Nuovo numero generato per una nuova partita
		count = 0                       // Resetta il contatore dei tentativi
	}

	// Risposta al cliente con il risultato
	tmpl, _ := template.New("result").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Gioco del Numero</title>
	</head>
	<body>
		<h1>Gioco del Numero</h1>
		<p>{{.}}</p>
		<a href="/">Prova un altro numero!</a>
	</body>
	</html>
	`)
	tmpl.Execute(w, messaggio)
}

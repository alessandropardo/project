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
	<html lang="it">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Gioco del Numero</title>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				background: url('https://www.w3schools.com/w3images/forest.jpg') no-repeat center center fixed;
				background-size: cover;
				color: white;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
				margin: 0;
				font-size: 1.2em;
			}
			.container {
				text-align: center;
				background-color: rgba(0, 0, 0, 0.6);
				padding: 30px;
				border-radius: 10px;
				width: 100%;
				max-width: 450px;
			}
			h1 {
				font-size: 2.5em;
				color: #ffeb3b;
			}
			form {
				margin-top: 20px;
			}
			input[type="number"] {
				padding: 12px;
				font-size: 1.5em;
				width: 80%;
				border: 1px solid #fff;
				border-radius: 8px;
				margin-top: 10px;
			}
			button {
				padding: 12px 20px;
				font-size: 1.2em;
				border: none;
				border-radius: 8px;
				background-color: #4CAF50;
				color: white;
				cursor: pointer;
				margin-top: 15px;
			}
			button:hover {
				background-color: #45a049;
			}
			a {
				display: inline-block;
				margin-top: 20px;
				color: #ffeb3b;
				text-decoration: none;
				font-size: 1.2em;
			}
			a:hover {
				text-decoration: underline;
			}
		</style>
		<script>
			function playSound(file) {
				var audio = new Audio(file);
				audio.play();
			}
		</script>
	</head>
	<body>
		<div class="container">
			<h1>Indovina il Numero!</h1>
			<p>Prova a indovinare il numero tra 0 e 99.</p>
			<form action="/tentativo" method="post">
				<input type="number" id="tentativo" name="tentativo" required>
				<button type="submit" onclick="playSound('click-sound.mp3')">Invia</button>
			</form>
		</div>
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
	var soundFile string
	if tentativo < numeroGenerato {
		messaggio = "Troppo basso! Riprova!"
		soundFile = "error-sound.mp3"
	} else if tentativo > numeroGenerato {
		messaggio = "Troppo alto! Riprova!"
		soundFile = "error-sound.mp3"
	} else {
		messaggio = fmt.Sprintf("Hai indovinato in %d tentativi!", count)
		soundFile = "success-sound.mp3"
		numeroGenerato = rand.Intn(100) // Nuovo numero generato per una nuova partita
		count = 0                       // Resetta il contatore dei tentativi
	}

	// Risposta al cliente con il risultato
	tmpl, _ := template.New("result").Parse(`
	<!DOCTYPE html>
	<html lang="it">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Gioco del Numero</title>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				background: url('https://www.w3schools.com/w3images/forest.jpg') no-repeat center center fixed;
				background-size: cover;
				color: white;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
				margin: 0;
				font-size: 1.2em;
			}
			.container {
				text-align: center;
				background-color: rgba(0, 0, 0, 0.6);
				padding: 30px;
				border-radius: 10px;
				width: 100%;
				max-width: 450px;
			}
			h1 {
				font-size: 2.5em;
				color: #ffeb3b;
			}
			p {
				font-size: 1.3em;
				margin-top: 20px;
			}
			a {
				display: inline-block;
				margin-top: 20px;
				color: #ffeb3b;
				text-decoration: none;
				font-size: 1.2em;
			}
			a:hover {
				text-decoration: underline;
			}
		</style>
		<script>
			function playSound(file) {
				var audio = new Audio(file);
				audio.play();
			}
		</script>
	</head>
	<body onload="playSound('{{.SoundFile}}')">
		<div class="container">
			<h1>Gioco del Numero</h1>
			<p>{{.Message}}</p>
			<a href="/">Prova un altro numero!</a>
		</div>
	</body>
	</html>
	`)
	tmpl.Execute(w, map[string]interface{}{
		"Message":   messaggio,
		"SoundFile": soundFile,
	})
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

const commonStyle = `
<style>
    body {
        font-family: Arial, sans-serif;
        background-color: #f4f4f9;
        color: #333;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }
    .container {
        text-align: center;
        padding: 20px;
        background: white;
        border-radius: 10px;
        box-shadow: 0 4px 6px rgba(0,0,0,0.1);
    }
    h1 { color: #2c3e50; }
    p { font-size: 1.2em; }
    .number {
        font-size: 2em;
        font-weight: bold;
        color: #e74c3c;
        margin-top: 10px;
    }
		.content-image {
        max-width: 100%;
        height: auto;
        border-radius: 8px;
        margin: 20px 0;
        box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    }
    a {
        display: inline-block;
        margin-top: 20px;
        text-decoration: none;
        color: #3498db;
    }
    a:hover { text-decoration: underline; }
</style>
`

func main() {

	fs := http.FileServer(http.Dir("./image"))
	http.Handle("/image/", http.StripPrefix("/image/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<html><head>"+commonStyle+"</head><body><div class='container'><h1>Hello, World!</h1><a href='/info'>К инфо</a> | <a href='/treshak'>Лютая Жесть</a> | <a href='/obnal'>ООО Обнал</a></div></body></html>")
	})

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<html>
			<head>%s</head>
			<body>
				<div class="container">
					<h1>Информация</h1>
					<p class="number">%d</p>
					<a href="/">На главную</a>
				</div>
			</body>
			</html>`, commonStyle, rand.Intn(1000))
	})

	http.HandleFunc("/obnal", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		const usd = 83.92
		const eur = 94.04

		var resultHTML string

		if r.Method == "POST" {
			r.ParseForm()
			rublesStr := r.FormValue("rubles")

			var rubles float64
			_, err := fmt.Sscanf(rublesStr, "%f", &rubles)

			if err != nil || rubles <= 0 {
				resultHTML = `<p style="color: red;">Ошибка: введите корректное число!</p>`
			} else {
				dollars := rubles / usd
				euros := rubles / eur
				resultHTML = fmt.Sprintf(`
                <div class="result">
                    <p class="number">%.2f ₽</p>
                    <p>Доллары: <strong>%.2f $</strong></p>
                    <p>Евро: <strong>%.2f €</strong></p>
                    <p style="font-size: 0.9em; color: #7f8c8d;">
                        Курс: 1$ = %.2f₽, 1€ = %.2f₽
                    </p>
                </div>`, rubles, dollars, euros, usd, eur)
			}
		}

		fmt.Fprintf(w, `
        <html>
        <head>%s</head>
        <body>
            <div class="container">
                <h1>ООО "Обнал"</h1>
                <h2>Конвертер валют</h2>
                
                <form method="POST" action="/obnal">
                    <label for="rubles">Введите сумму в рублях:</label><br>
                    <input type="number" id="rubles" name="rubles" 
                           step="0.01" min="0" placeholder="0.00" 
                           style="padding: 10px; font-size: 1.1em; 
                                  border: 2px solid #3498db; 
                                  border-radius: 5px; 
                                  margin: 10px 0;
                                  width: 200px;"><br>
                    <button type="submit" 
                            style="padding: 10px 20px; 
                                   font-size: 1.1em; 
                                   background: #3498db; 
                                   color: white; 
                                   border: none; 
                                   border-radius: 5px; 
                                   cursor: pointer;">
                        Конвертировать
                    </button>
                </form>
                
                %s
                
                <br>
                <a href="/">На главную</a>
            </div>
        </body>
        </html>`, commonStyle, resultHTML)
	})

	http.HandleFunc("/treshak", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<html>
			<head>%s</head>
			<body>
				<div class="container">
					<h1>Я предупреждал</h1>
					<img src="/image/aaa.jpeg" alt="Картинка" class="content-image">
					<br>
					<a href="/">На главную</a>
				</div>
			</body>
			</html>`, commonStyle)
	})

	fmt.Print("Страница работает: http://localhost:8080 ")
	fmt.Print(" http://localhost:8080/treshak ")
	fmt.Print(" http://localhost:8080/info ")
	fmt.Print(" http://localhost:8080/obnal ")

	http.ListenAndServe(":8080", nil)
}

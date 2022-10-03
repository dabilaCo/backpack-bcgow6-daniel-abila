Muy buenos ejemplos de implementaciones diferentes en http: https://www.alexedwards.net/blog/golang-response-snippets

```browser
    http://localhost:8080/bootcampers
```

```zsh
curl -X GET http://localhost:8080/bootcampers
```

```js
fetch('http://localhost:8080/bootcampers')
  .then(response => response.json())
  .then(data => console.log(data));
```


```go
http.HandleFunc("/peliculas", GetPeliculas)

	fmt.Println("Inicializando servidor en puerto 8080 🚀")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("No se pudo iniciar el servidor: error ", err)
	}
```
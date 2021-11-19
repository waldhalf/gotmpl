# Go webser template

- Built in Go version 1.15
- Uses [Chi router](https://github.com/go-chi/chi) router
- Uses [SCS](https://github.com/alexedwards/scs) session management by Alex Edwrads
- Uses [nosurf](https://github.com/justinas/nosurf)

# Git branches

## 01 - Enable static files and create template/js/css

- create a folder named static at root level of the application
- in the rooter file

```
fileServer := http.FileServer(http.Dir("./static/"))
```

```
mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
```

## 02 - Create a logger system

## 03 - Create migrations with soda (part of gobuffalo)
## 04 - Create repository db + cleanup code
## 05 - Create db function for rooms, reservation, restrictions
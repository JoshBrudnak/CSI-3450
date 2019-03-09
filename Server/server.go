package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    _ "github.com/lib/pq"
    "html/template"
    "log"
    "math/rand"
    "net/http"
    "os"
    "time"
)

var db *sql.DB
var templates *template.Template

const (
    createUserT = "create table if not exists budget_user(id serial primary key, fname text, lname text, password text, email text, admin boolean, date date, t_money float);"
    createBudgetCategoryT = "create table if not exists budget_category(id serial primary key, user_id int REFERENCES budget_user (id), b_value float);"
    createBudgetEntryT = "create table if not exists budget_entry(id serial primary key, category_id int REFERENCES budget_category (id), date date, value float);"
    createSessionT = "create table if not exists Session(userId text, sessionKey text, time timestamp);"
)

type config struct {
    URL      string
    Username string
    Password string
    Dbname   string
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func logIfErr(err error) {
    if err != nil {
        log.Println(err)
    }
}

func logServerErr(w http.ResponseWriter, err error) {
    if err != nil {
        log.Println(err)
        http.Error(w, "Server error", http.StatusInternalServerError)
    }
}

func compileTemplates() {
    // TODO: get real path
    t, err := template.ParseFiles("../index.html")
    templates = template.Must(t, err)
}

func createHash() string {
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, 20)

    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }

    return string(b)
}

func init() {
    var c config
    file, err := os.Open("database.json")
    checkErr(err)
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&c)
    checkErr(err)
    dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Username, c.Password, c.URL, c.Dbname)
    db, err = sql.Open("postgres", dbURL)
    checkErr(err)
    db.SetMaxOpenConns(80)

    query(createUserT)
    query(createBudgetCategoryT)
    query(createBudgetEntryT)
    query(createSessionT)

    f, err := os.OpenFile("budget_app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    log.SetOutput(f)
    fmt.Println("Server started ...")

    rand.Seed(time.Now().UnixNano())
}

func home(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "index.html", nil)
}

func main() {
    //compileTemplates()

    //http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../react/dueto/build/static/js"))))
    //http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("../react/dueto/src/resources"))))
    http.HandleFunc("/api/login", login)
    http.HandleFunc("/api/logout", logout)
    http.HandleFunc("/api/profile", login)
    //http.HandleFunc("/", home)

    http.ListenAndServe(":8080", nil)
}

package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	AddArtist  = "insert into budget_user(fname, lname, password, email, admin, date, t_money) VALUES($1, $2, $3 $4, $5, NOW()::date, $6);"
	AddBudgetCategory = " insert into budget_category(user_id, category_name, b_value) VALUES($1, $2, $3);"
	AddBudgetEntry = "insert into budget_entry(category_id, date, value) VALUES($1, $2, $3);"
	AddSession = "insert into Session(userId, sessionKey, time) VALUES($1, $2, now()::timestamp);"

	RemoveSession     = "delete from Session where sessionkey = $1;"
	RemoveEntry     = "delete from budget_entry where id = $1;"
	RemoveCategory     = "delete from budget_category where id = $1;"
	RemoveOldSessions = "delete from session where age(now(), time) > '5 hour';"

	SelectUserData        = "select fname, lname, email, admin, date, t_money from budget_user where id = $1;"
	SelectUserCatagories        = "select id, category_name, b_value from budger_category where user_id = $1 order by category_name;"

	SelectUserAuth        = "select id, password from budget_user where email = $1;"
	SelectSession         = "select count(userId) from session where sessionkey = $1;"
	SelectAuthId          = "select userId from session where sessionKey = $1;"
)

type Authentication struct {
	Username string
	Password string
}

type NewUser struct {
	FirstName       string
	lastName       string
	Password   string
	Repassword string
	Email      string
	TotalMoney      string
}

type Category struct {
  Name string
  BValue string
}

type Profile struct {
	FirstName       string
	LastName       string
	Admin      bool
	Email      string
	Date      bool
	TotalMoney      string
	CategoryList []Category
}

func (p *Profile) SetCategoryList(c []Category) {
	p.CategoryList = c
}

func query(sql string) {
	_, err := db.Query(sql)
	logIfErr(err)
}

func authenticate(cookie *http.Cookie) bool {
	var sessionCount int

	if cookie.String() != "" {
		sessionId := cookie.Value

		rows, err := db.Query(SelectSession, sessionId)
		checkErr(err)

		rows.Next()
		err = rows.Scan(&sessionCount)
		logIfErr(err)
		rows.Close()

		if sessionCount > 0 {
			return true
		}
	}

	return false
}

func getUserId(sessionId string) string {
	var id string
	rows, err := db.Query(SelectAuthId, sessionId)
	logIfErr(err)

	rows.Next()
	err = rows.Scan(&id)
	logIfErr(err)
	rows.Close()

	return id
}

func login(w http.ResponseWriter, r *http.Request) {
	var id int
	var hashPassword string
	var auth Authentication

	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&auth)
	logIfErr(err)
	defer r.Body.Close()

	if auth.Username == "" || auth.Password == "" {
		return
	}

	rows, err := db.Query(RemoveOldSessions)
	logIfErr(err)
	rows.Close()

	rows, err = db.Query(SelectUserAuth, auth.Username)
	logIfErr(err)

	rows.Next()
	err = rows.Scan(&id, &hashPassword)
	logIfErr(err)
	rows.Close()

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(auth.Password))

	if err == nil {
		sessionId := createHash()
		rows, err = db.Query(AddSession, id, sessionId)
		logIfErr(err)
		rows.Close()

		exp := time.Now().Add(5 * time.Hour)
		cookie := http.Cookie{Name: "SESSIONID", Value: sessionId, Path: "/", Expires: exp, HttpOnly: true}
		http.SetCookie(w, &cookie)
	} else {
		http.Error(w, "Incorrect username or password", 401)
	}
}

func profile(w http.ResponseWriter, r *http.Request) {
	var p Profile
	var categories []Category

	w.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, _ := r.Cookie("SESSIONID")

	if !authenticate(cookie) {
		http.Error(w, "Authentication failed", http.StatusForbidden)
		return
	}

	userId := getUserId(cookie.Value)

	rows, err := db.Query(SelectUserData, userId)
	checkErr(err)

	rows.Next()
	err = rows.Scan(&p.FirstName, &p.LastName, &p.Email, &p.Date, &p.Admin, &p.TotalMoney)
	logIfErr(err)
	rows.Close()

	categoryRows, viderr := db.Query(SelectUserCatagories, userId)
	logIfErr(viderr)

	for categoryRows.Next() {
		var c Category
		err = categoryRows.Scan(&c.Name, &c.BValue)
		logIfErr(err)

	    categories = append(categories, c)
	}
	categoryRows.Close()

	p.SetCategoryList(categories)

	err = json.NewEncoder(w).Encode(p)
	logServerErr(w, err)
}

func logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, err := r.Cookie("SESSIONID")
	sessionId := cookie.Value

	rows, err := db.Query(RemoveSession, sessionId)
	logIfErr(err)
	rows.Close()
}

package main

import (
	"encoding/json"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	AddUser           = "insert into budget_user(fname, lname, password, email, admin, date, t_money) VALUES($1, $2, $3 $4, $5, NOW()::date, $6);"
	AddBudgetCategory = " insert into budget_category(user_id, category_name, b_value) VALUES($1, $2, $3);"
	AddBudgetEntry    = "insert into budget_entry(category_id, date, value) VALUES($1, now()::timestamp, $3);"
	AddSession        = "insert into Session(userId, sessionKey, time) VALUES($1, $2, now()::timestamp);"

	UpdateUser     = " update budget_user set fname = $2, lname = $3, email = $4, t_money = $5 where id = $1;"
	UpdatePassword = " update budget_user set password = $2 where id = $1;"

	RemoveSession       = "delete from Session where sessionkey = $1;"
	RemoveUser          = "delete from budget_user where user_id = $1;"
	RemoveEntry         = "delete from budget_entry where id = $1;"
	RemoveAllEntries    = "delete from budget_entry where category_id = (select id from budget_category where user_id = $1)"
	RemoveCategory      = "delete from budget_category where id = $1;"
	RemoveAllCategories = "delete from budget_category where user_id = $1;"
	RemoveOldSessions   = "delete from session where age(now(), time) > '5 hour';"

	SelectUserData       = "select fname, lname, email, admin, date, t_money from budget_user where id = $1;"
	SelectDashboardData  = "select category_name, budget_entry.date, budget_entry.value from budget_entry, budget_category where category_id = budget_category.id and budget_category.user_id = $1;"
	SelectUserCatagories = "select id, category_name, b_value from budger_category where user_id = $1 order by category_name;"

	SelectUserAuth = "select id, password from budget_user where email = $1;"
	SelectSession  = "select count(userId) from session where sessionkey = $1;"
	SelectAuthId   = "select userId from session where sessionKey = $1;"
)

type Authentication struct {
	Username string
	Password string
}

type NewUser struct {
	FirstName  string
	LastName   string
	Password   string
	Repassword string
	Email      string
	TotalMoney string
}

type Category struct {
	Name   string
	BValue string
}

type CategoryEntry struct {
	CategoryName string
	Date         string
	Value        string
}

type Entry struct {
	CatagoryId string
	Value      string
}

type Profile struct {
	FirstName    string
	LastName     string
	Admin        bool
	Email        string
	Date         bool
	TotalMoney   string
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

func dashboard(w http.ResponseWriter, r *http.Request) {
	var dashboard []CategoryEntry

	w.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, _ := r.Cookie("SESSIONID")

	if !authenticate(cookie) {
		http.Error(w, "Authentication failed", http.StatusForbidden)
		return
	}

	userId := getUserId(cookie.Value)

	rows, err := db.Query(SelectDashboardData, userId)
	checkErr(err)

	for rows.Next() {
		var c CategoryEntry
		err = rows.Scan(&c.CategoryName, &c.Date, &c.Value)
		logIfErr(err)

		dashboard = append(dashboard, c)
	}
	rows.Close()

	err = json.NewEncoder(w).Encode(dashboard)
	logServerErr(w, err)
}

func makeTransaction(w http.ResponseWriter, r *http.Request) {
	var data Entry

	w.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, _ := r.Cookie("SESSIONID")

	if !authenticate(cookie) {
		http.Error(w, "Authentication failed", http.StatusForbidden)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	_, err := db.Query(AddBudgetEntry, data.CatagoryId, data.Value)
	logIfErr(err)
}

func makeCategory(w http.ResponseWriter, r *http.Request) {
	var data Category

	w.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, _ := r.Cookie("SESSIONID")

	if !authenticate(cookie) {
		http.Error(w, "Authentication failed", http.StatusForbidden)
		return
	}

	userId := getUserId(cookie.Value)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	_, err := db.Query(AddBudgetCategory, userId, data.Name, data.BValue)
	logIfErr(err)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, _ := r.Cookie("SESSIONID")

	if !authenticate(cookie) {
		http.Error(w, "Authentication failed", http.StatusForbidden)
	}

	userId := getUserId(cookie.Value)

	_, err := db.Query(RemoveAllEntries, userId)
	logIfErr(err)
	_, err = db.Query(RemoveAllCategories, userId)
	logIfErr(err)
	_, err = db.Query(RemoveUser, userId)
	logIfErr(err)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var id string
	var hashPassword string
	var data NewUser

	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	if data.Password != data.Repassword {
		http.Error(w, "Passwords do not match", http.StatusNotAcceptable)
		return
	}

	err := checkmail.ValidateFormat(data.Email)
	if err != nil {
		http.Error(w, "Enter a valid email", http.StatusNotAcceptable)
		return
	}

	bHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 1)
	hash := string(bHash)

	rows, err := db.Query(AddUser, data.FirstName, data.LastName, hash, data.Email, data.TotalMoney)
	logIfErr(err)
	rows.Close()

	authRows, err := db.Query(SelectUserAuth, data.Email)
	logServerErr(w, err)

	if authRows.Next() {
		err = authRows.Scan(&id, &hashPassword)
		logIfErr(err)
	} else {
		http.Error(w, "Unable to create user", http.StatusForbidden)
	}
	authRows.Close()
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	var data NewUser

	cookie, _ := r.Cookie("SESSIONID")

	if !authenticate(cookie) {
		http.Error(w, "Authentication failed", http.StatusForbidden)
	}

	userId := getUserId(cookie.Value)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	err := checkmail.ValidateFormat(data.Email)
	if err != nil {
		http.Error(w, "Enter a valid email", http.StatusNotAcceptable)
		return
	}

	rows, err := db.Query(UpdateUser, userId, data.FirstName, data.LastName, data.Email, data.TotalMoney)
	logIfErr(err)
	rows.Close()
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
	var data NewUser

	cookie, _ := r.Cookie("SESSIONID")

	if !authenticate(cookie) {
		http.Error(w, "Authentication failed", http.StatusForbidden)
	}

	userId := getUserId(cookie.Value)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	if data.Password != data.Repassword {
		http.Error(w, "Passwords do not match", http.StatusNotAcceptable)
		return
	}

	err := checkmail.ValidateFormat(data.Email)
	if err != nil {
		http.Error(w, "Enter a valid email", http.StatusNotAcceptable)
		return
	}

	bHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 1)
	hash := string(bHash)

	rows, err := db.Query(UpdatePassword, userId, hash)
	logIfErr(err)
	rows.Close()
}

func logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, err := r.Cookie("SESSIONID")
	sessionId := cookie.Value

	rows, err := db.Query(RemoveSession, sessionId)
	logIfErr(err)
	rows.Close()
}

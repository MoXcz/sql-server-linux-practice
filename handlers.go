package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MoXcz/sql-server-linux-practice/ui/html"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	component := html.Base()
	component.Render(r.Context(), w)
}

func (app *application) handlePostAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		Logger.Error("could not parse form", "err", err)
		return
	}

	name := r.FormValue("name")
	balance := r.FormValue("balance")

	balanceF, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		return
	}

	id, err := app.account.Insert(name, balanceF)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	account, err := app.account.Get(id)

	http.Redirect(w, r, fmt.Sprintf("/account/%d", id), http.StatusSeeOther)
	component := html.Account(account)
	component.Render(r.Context(), w)
}

func (app *application) handleGetNewAccount(w http.ResponseWriter, r *http.Request) {
	component := html.AccountForm()
	component.Render(r.Context(), w)
}

func (app *application) handleGetAccounts(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("accountID"))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	acc, err := app.account.Get(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	component := html.Account(acc)
	component.Render(r.Context(), w)
}

func (app *application) handleGetAccount(w http.ResponseWriter, r *http.Request) {
	accs, err := app.account.Latest()
	if err != nil {
		return
	}
	component := html.Accounts(accs)
	component.Render(r.Context(), w)
}

package handlers

import (
    "context"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "log"
    "net/http"
    "regexp"
)

type AuthorizationHandler struct {
    router *mux.Router
}

func NewAuthorizationHandler(router *mux.Router) *AuthorizationHandler {
    return &AuthorizationHandler{router: router}
}

func (auth *AuthorizationHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    
    path := req.URL.Path

    loginRegexp := regexp.MustCompile(`^(/login\.html)|(/api/login)$`)

    if loginRegexp.MatchString(path) {
        auth.router.ServeHTTP(w, req)
        return
    }

    var store = sessions.NewCookieStore([]byte("2Macri0Gato1Ignorante9"))
        
    session, err := store.Get(req, "mechanic-session")
    if err != nil {
        log.Println(err.Error())
        return
    }

    if session.IsNew {
        log.Println("User not authenticated. Redirecting to login page.")
        http.Redirect(w, req, "/login.html", http.StatusSeeOther)
        return
    }
    
    userId, exist := session.Values["user_id"]

    if !exist {
        errStr := "Missing user_id in user session"
        log.Println(errStr)
        http.Error(w, errStr, http.StatusInternalServerError)
        return
    }

    prevContext := req.Context()
    newContext := context.WithValue(prevContext, "user_id", userId)

    auth.router.ServeHTTP(w, req.WithContext(newContext))
    return
}
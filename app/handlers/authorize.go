package handlers

import(
    "regexp"
    "log"
    "net/http"
    "context"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
)

type AuthorizationHandler struct {
    router *mux.Router
}

func NewAuthorizationHandler(router *mux.Router) *AuthorizationHandler {
    return &AuthorizationHandler{router: router}
}

func (auth *AuthorizationHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    
    path := req.URL.Path
    
    staticLoginRE := regexp.MustCompile(`^/static/login/$`)
    loginRE := regexp.MustCompile(`^/login$`)

    if staticLoginRE.MatchString(path) || loginRE.MatchString(path) {
        auth.router.ServeHTTP(w, req)
        return
    }

    var store = sessions.NewCookieStore([]byte("2Macri0Gato1Ignorante9"))
        
    session, err := store.Get(req, "mechanic-session")
    if err != nil {
        log.Println(err.Error());
        return
    }

    if session.IsNew {
        http.Redirect(w, req, "/static/login/", http.StatusSeeOther);
        return
    }
    
    user_id, exist := session.Values["user_id"];
    
    if !exist {
        errStr := "Missing user_id in user session"
        log.Println(errStr)
        http.Error(w, errStr, http.StatusInternalServerError)
        return
    }

    prevContext := req.Context()
    newContext := context.WithValue(prevContext, "user_id", user_id)

    auth.router.ServeHTTP(w, req.WithContext(newContext))
    return
}
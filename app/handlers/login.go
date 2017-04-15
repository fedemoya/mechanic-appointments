package handlers

import(
    "log"
    "net/http"
    "github.com/gorilla/sessions"
    "mechanics-backend/app/persistance"
    "mechanics-backend/app/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

  var err error

  var login string
  login, err = Validate(r.FormValue("Login")).Required().AsLogin()
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  var password string
  password, err = Validate(r.FormValue("Password")).Required().AsPassword()
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  repository := persistance.NewRepository("mechanics.db")

  users := []models.User{}
  err = repository.Search(&models.User{}, &users, "Login = ?", login)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  if len(users) > 0 && users[0].Password == password {

    log.Println("El usuario " + users[0].Login + " a sido logueado satisfactoriamente.");
    
    var store = sessions.NewCookieStore([]byte("2Macri0Gato1Ignorante9"))
    session, err := store.Get(r, "mechanic-session")
    if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
    }

    session.Values["user_id"] = users[0].Id
    session.Save(r, w)

  } else {
    
    http.Error(w, "Usuario o contraseña icorrecto", http.StatusUnauthorized)
    return

  }
  
}

package handlers

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil{
			fmt.Fprintf(w, 'erro ao fazer parse do form: %v", err)
			return
		}
	err := controllers.Create(r.PostForm.Get("name"), r.PostForm("email"))
	if err != nil {
		fmt.Fprint(w, "erro ao salvar os dados: %v", err)
		return
	}
}

	http.ServeFile(w, r, "handlers/templates/subscription_form.html")
}

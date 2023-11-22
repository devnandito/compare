package structs

type EmailStore struct {
	Email string
	Name string
	Module string
	Content map[string]string
	Status string
	Past string
}

func (es *EmailStore) InsertData(data *EmailStore) (EmailStore) {
	emailStore := EmailStore {
		Email: data.Email,
		Name: data.Name,
		Module: data.Module,
		Content: data.Content,
	}

	return emailStore
}

func (es *EmailStore) InsetArray(content, module string) (data map[string]string) {
	data[module] = data[module]
	return data
}

func (es *EmailStore) Verifity(email string, data map[string]EmailStore) (e bool) {
	for _, v := range data {
		if email == v.Email {
			e = true
		} else {
			e = false
		}
	}

	return e
}

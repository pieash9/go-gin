package services

type NotesServices struct {
}

type Note struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (n *NotesServices) GetNotes() []Note {
	data := []Note{
		{
			Id:   1,
			Name: "Note 1",
		},
		{
			Id:   2,
			Name: "Note 2",
		},
	}
	return data
}

func (n *NotesServices) CreateNote() string {
	return "POST request notes."
}

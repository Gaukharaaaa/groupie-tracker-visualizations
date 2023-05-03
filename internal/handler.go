package internal

import (
	"encoding/json"
	"groupie-tracker/api"
	"io"
	"net/http"
	"strconv"
	"text/template"
)

type asd struct {
	ArtistPage api.Artist
	Rel        api.Relation
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		CheckError("Ops, Page Not Found 404", 404, w)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var temp []api.Artists
	res, err := http.Get(api.AllLinks)
	if err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}
	file, err := io.ReadAll(res.Body)
	if err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}
	if err := res.Body.Close(); err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}
	if err := json.Unmarshal(file, &temp); err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}

	templ, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}
	if err := templ.Execute(w, temp); err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}
}

func ArtistsPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		CheckError("Ops, Bad Request 400", 400, w)
		return
	}
	if id == 0 || id < 0 || id > 52 {
		CheckError("Ops, Page Not Found 404", 404, w)
		return
	}

	idString := strconv.Itoa(id)
	allartists, err := UnmarshalStruct(api.AllLinks + "/" + idString)
	if err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}
	allrelation, err := UnmarshalStructRelation(api.Relations + idString)
	if err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}
	AllData := asd{
		ArtistPage: allartists,
		Rel:        allrelation,
	}
	templ, err := template.ParseFiles("./ui/html/artist.html")
	if err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}

	if err = templ.Execute(w, AllData); err != nil {
		CheckError("Ops, Internal Server 500", 500, w)
		return
	}
}

func UnmarshalStruct(url string) (api.Artist, error) {
	var artist api.Artist
	res, err := http.Get(url)
	if err != nil {
		return artist, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return artist, err
	}
	if err := res.Body.Close(); err != nil {
		return artist, err
	}
	if err := json.Unmarshal(data, &artist); err != nil {
		return artist, err
	}
	return artist, nil
}

func UnmarshalStructRelation(url string) (api.Relation, error) {
	var rel api.Relation
	res, err := http.Get(url)
	if err != nil {
		return rel, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return rel, err
	}
	if err := res.Body.Close(); err != nil {
		return rel, err
	}

	if err := json.Unmarshal(data, &rel); err != nil {
		return rel, err
	}
	return rel, nil
}

func CheckError(s string, num int, w http.ResponseWriter) {
	errorPage, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(num)
	errorPage.Execute(w, s)
}

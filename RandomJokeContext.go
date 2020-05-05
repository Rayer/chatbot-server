package ChatbotAPIs

import (
	"encoding/json"
	"net/http"
)

/*
{"id":102,"type":"general","setup":"Did you hear the one about the guy with the broken hearing aid?","punchline":"Neither did he."}
*/
type RandomJoke struct {
	Id        int
	Type      string
	Setup     string
	Punchline string
}

func fetchRandomJoke() (*RandomJoke, error) {
	res, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	ret := RandomJoke{}
	err = json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

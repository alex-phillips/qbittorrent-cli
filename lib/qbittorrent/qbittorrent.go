package qbittorrent

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/alex-phillips/qbittorrent/lib/log"
)

type Api struct {
	Client *http.Client
	Host   string
}

func GetApi(host string, username string, password string) *Api {
	jar, _ := cookiejar.New(nil)
	api := Api{
		Client: &http.Client{Jar: jar},
		Host:   host,
	}

	api.Login(username, password)

	return &api
}

func (api *Api) Delete(hash string) (*string, error) {
	resp, err := api.Client.PostForm(api.Host+"/command/delete", url.Values{
		"hashes": {hash},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	retval := string(body)
	if retval != "" {
		return nil, errors.New(retval)
	}

	return &retval, nil
}

func (api *Api) DeletePermanently(hash string) (*string, error) {
	resp, err := api.Client.PostForm(api.Host+"/command/deletePerm", url.Values{
		"hashes": {hash},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	retval := string(body)
	if retval != "" {
		return nil, errors.New(retval)
	}

	return &retval, nil
}

func (api *Api) GetPreferences() string {
	resp, err := api.Client.Get(api.Host + "/query/preferences")
	if err != nil {
		log.Error.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (api *Api) GetTorrents(filters map[string]string) string {
	req, err := http.NewRequest("GET", api.Host+"/query/torrents", nil)
	if err != nil {
		log.Error.Fatalln(err)
	}

	q := req.URL.Query()
	for key, val := range filters {
		q.Add(key, val)
	}

	req.URL.RawQuery = q.Encode()

	// var torrents []Torrent

	resp, err := api.Client.Do(req)
	if err != nil {
		return "[]"
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)

	// err = json.Unmarshal(body, &torrents)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// return torrents
}

func (api *Api) Login(username string, password string) (*string, error) {
	resp, err := api.Client.PostForm(api.Host+"/login", url.Values{
		"username": {username},
		"password": {password},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	retval := string(body)
	if retval != "" {
		return nil, errors.New(retval)
	}

	return &retval, nil
}

func (api *Api) Pause(hash string) (*string, error) {
	resp, err := api.Client.PostForm(api.Host+"/command/pause", url.Values{
		"hash": {hash},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	retval := string(body)
	if retval != "" {
		return nil, errors.New(retval)
	}

	return &retval, nil
}

// func (api *Api) SetCategory(hash string, category string) (*string, error) {
// 	resp, err := api.Client.PostForm(api.Host+"/command/setCategory", url.Values{
// 		"hash":     {hash},
// 		"category": {category},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)

// 	retval := string(body)

// 	return &retval, nil
// }

func (api *Api) SetPreference(key string, value string) (*string, error) {
	preferences := map[string]string{
		key: value,
	}
	data, _ := json.Marshal(preferences)
	resp, err := api.Client.PostForm(api.Host+"/command/setPreferences", url.Values{
		"json": {string(data)},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	retval := string(body)
	if retval != "" {
		return nil, errors.New(retval)
	}

	return &retval, nil
}

func (api *Api) SetSavePath(hash string, savepath string) (*string, error) {
	resp, err := api.Client.PostForm(api.Host+"/command/setLocation", url.Values{
		"hashes":   {hash},
		"location": {savepath},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	retval := string(body)
	if retval != "" {
		return nil, errors.New(retval)
	}

	return &retval, nil
}

func (api *Api) Resume(hash string) (*string, error) {
	resp, err := api.Client.PostForm(api.Host+"/command/resume", url.Values{
		"hash": {hash},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	retval := string(body)
	if retval != "" {
		return nil, errors.New(retval)
	}

	return &retval, nil
}

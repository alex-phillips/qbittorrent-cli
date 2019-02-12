package qbittorrent

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"

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

func (api *Api) UploadFile(path string, params map[string]string) (*string, error) {
	file, _ := os.Open(path)
	defer file.Close()

	formBody := &bytes.Buffer{}
	writer := multipart.NewWriter(formBody)

	for k, v := range params {
		_ = writer.WriteField(k, v)
	}

	part, _ := writer.CreateFormFile("torrents", filepath.Base(file.Name()))
	io.Copy(part, file)

	writer.Close()

	req, err := http.NewRequest("POST", api.Host+"/command/upload", formBody)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		log.Error.Fatalln(err)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		log.Error.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error.Fatalln(err)
	}

	fmt.Println(string(body))

	retval := string(body)

	return &retval, nil
}

func (api *Api) UploadLink(link string, params map[string]string) (*string, error) {
	urlVals := url.Values{
		"urls": {link},
	}

	for k, v := range params {
		urlVals.Add(k, v)
	}

	resp, err := api.Client.PostForm(api.Host+"/command/download", urlVals)
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

package http

import (
        "bytes"
        "crypto/tls"
        "io/ioutil"
        "net/http"
        "time"
)

var tr = &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

var httpClient = &http.Client{
        //Transport: &http.Transport{
        //      MaxIdleConnsPerHost: 5,
        //},
        Timeout: 5 * time.Second,
        Transport: tr,
}

func Send(req *http.Request) ([]byte, error) {
        resp, err := httpClient.Do(req)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        return body, err
}

func Get(url string, contentType, token string) ([]byte, error) {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                return nil, err
        }
        req.Header.Set("Content-Type", contentType)
        if token != "" {
                req.Header.Set("Authorization", token)
        }

        return Send(req)
}

func Post(url string, body []byte, contentType, token string) ([]byte, error) {
        req, err := http.NewRequest("POST", url, bytes.NewReader(body))
        if err != nil {
                return nil, err
        }
        req.Header.Set("Content-Type", contentType)
        if token != "" {
                req.Header.Set("Authorization", token)
        }
        return Send(req)
}

func Put(url string, body []byte, contentType, token string) ([]byte, error) {
        req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
        if err != nil {
                return nil, err
        }
        req.Header.Set("Content-Type", contentType)
        if token != "" {
                req.Header.Set("Authorization", token)
        }
        return Send(req)
}

func Delete(url string, body []byte, contentType, token string) ([]byte, error) {
        req, err := http.NewRequest("DELETE", url, bytes.NewReader(body))
        if err != nil {
                return nil, err
        }
        req.Header.Set("Content-Type", contentType)
        if token != "" {
                req.Header.Set("Authorization", token)
        }
        return Send(req)
}

func PostWithJson(url string, body []byte, token string) ([]byte, error) {
        //var header http.Header
        return Post(url, body, "application/json;charset=UTF-8", token)
}

package main

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func MakeHttpRequest(method, url string, reqData io.Reader) (string, int) {

	var body string
	var status int
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		// Sometimes responses take FOREVER.
		TLSHandshakeTimeout: 60 * time.Second,
	}
	c := &http.Client{
		Transport: t,
	}

	if method == "GET" {
		resp, _ := c.Get(url)
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error: Could not convert response body to byte array - ", err)
		}

		bodyString := string(bodyBytes)
		body = bodyString
		status = resp.StatusCode
	} else if method == "POST" {
		resp, _ := c.Post(url, "", reqData)
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error: Could not convert response body to byte array - ", err)
		}

		bodyString := string(bodyBytes)
		body = bodyString
		status = resp.StatusCode
	}
	return body, status
}

package core

import (
	"crypto/sha256"
	"github.com/gabriel-vasile/mimetype"
	"io/ioutil"
	"net/http"
)

type DownloadInfo struct {
	Status      int
	Body        []byte
	ContentType string
}

type WebsiteInfo struct {
	StatusCode   int
	Body         []byte
	ContentType  string
	Size         int
	ContentTypeG string
	Hash         [32]byte
}

func Download(url string) (error, DownloadInfo) {
	resp, err := http.Get(url)
	if err != nil {
		return err, DownloadInfo{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, DownloadInfo{}
	}

	return nil, DownloadInfo{Status: resp.StatusCode, Body: body, ContentType: resp.Header.Get("Content-Type")}
}

func GetInfo(url string) (error, WebsiteInfo) {
	err, d := Download(url)
	if err != nil {
		return err, WebsiteInfo{}
	}

	return nil, WebsiteInfo{
		StatusCode:   d.Status,
		Body:         d.Body,
		Size:         len(d.Body),
		ContentType:  d.ContentType,
		ContentTypeG: mimetype.Detect(d.Body).String(),
		Hash:         hash(d.Body),
	}
}

func hash(b []byte) [32]byte {
	return sha256.Sum256(b)
}

package engine

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/shivakishore14/govm/domain"
	"log"
)

var URL = "https://golang.org/dl/"

func RemoteList(os string, arch string) domain.Versions {
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}
	remoteVersions := []domain.Version{}
	doc.Find(".toggle, .toggleVisible").Each(func(i int, s *goquery.Selection) {
		s.Find(".codetable").Each(func(i int, s *goquery.Selection) {
			v := getFileByConf(s, os, arch)
			if !v.IsEmpty() {
				remoteVersions = append(remoteVersions, v)
			}
		})
	})
	return remoteVersions
}

func getFileByConf(tableSelection *goquery.Selection, os string, arch string) domain.Version {
	version := domain.Version{}
	name, _ := tableSelection.Parent().Parent().Attr("id")
	tableSelection.Find("tr").Each(func(i int, s *goquery.Selection) {
		data := s.Find("td")
		link, _ := data.Eq(0).Find("a").Attr("href")

		v := domain.Version{
			Name:         name,
			DownloadLink: link,
			FileName:     data.Eq(0).Text(),
			Kind:         data.Eq(1).Text(),
			Os:           data.Eq(2).Text(),
			Arch:         data.Eq(3).Text(),
			Size:         data.Eq(4).Text(),
			SHA1:         data.Eq(5).Text(),
		}
		if v.Kind == "Archive" && v.Os == getRemoteOs(os) && v.Arch == getRemoteArch(arch) {
			version = v
		}
	})
	return version
}

func getRemoteOs(hostOs string) string {
	switch hostOs {
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	}
	return ""
}

func getRemoteArch(hostArch string) string {
	switch hostArch {
	case "amd64":
		return "x86-64"
	}
	return "x86"
}

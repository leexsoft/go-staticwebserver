package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var myServer = LoadAppConfig()

type server struct {
	XMLName xml.Name    `xml:"server"`
	App     application `xml:"app"`
}

type application struct {
	XMLName      xml.Name `xml:"app"`
	IP           string   `xml:"ip,attr"`
	Port         string   `xml:"port,attr"`
	Root         string   `xml:"root,attr"`
	StaticFile   string   `xml:"staticfile"`
	StaticFolder string   `xml:"staticfolder"`
}

func (this *server) GetStaticFileExtensionds() []string {
	strs := []string{}
	if len(this.App.StaticFile) > 0 {
		strs = strings.Split(this.App.StaticFile, "|")
	}
	return strs
}

func (this *server) GetStaticFolders() []string {
	strs := []string{}
	if len(this.App.StaticFolder) > 0 {
		strs = strings.Split(this.App.StaticFolder, "|")
	}
	return strs
}

func LoadAppConfig() *server {
	file, err := os.Open("app.config")
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	sve := new(server)
	err = xml.Unmarshal(data, sve)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	} else {
		return sve
	}
}

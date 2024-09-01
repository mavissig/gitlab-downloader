package loader

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Loader struct {
	cfg      *Config
	projects []*proj
}

type proj struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Links struct {
		Self     string `json:"self"`
		Branches string `json:"repo_branches"`
	} `json:"_links"`
	Branches []string
}

func New() *Loader {
	return &Loader{
		cfg: LoadConfig(),
	}
}

func (l *Loader) Run() {

	err := l.getAllProj()
	if err != nil {
		log.Fatalln(err)
	}

	err = l.downloadAllProj("Download")
	if err != nil {
		log.Fatalln(err)
	}
}

func (l *Loader) getAllProj() error {

	for i := 1; ; i++ {
		end, err := l.reqGetProj(i)
		if err != nil {
			return err
		}
		if end {
			break
		}
	}

	b, err := json.Marshal(l.projects)
	if err != nil {
		return err
	}

	err = write("res", b)
	if err != nil {
		return err
	}

	return nil
}

func (l *Loader) downloadAllProj(resFolderName string) error {
	err := os.Mkdir(resFolderName, 0777)
	if err != nil {
		return err
	}

	for _, p := range l.projects {
		branches, err := l.reqGetAllBranches(p.Links.Branches)
		if err != nil {
			log.Println("[DOWNLOAD][GET BRANCHES][ERROR]: ", err)
			continue
		}

		resFolderNameNew := fmt.Sprintf("%s/%s", resFolderName, p.Name)

		err = os.Mkdir(resFolderNameNew, 0777)
		if err != nil {
			return err
		}

		for _, branch := range branches {
			err = l.reqDownloadProj(resFolderNameNew, p.Name, branch, p.Links.Self)
			if err != nil {
				log.Println("[DOWNLOAD][ERROR]: ", err)
				continue
			}
		}
	}

	return nil
}

package usecase

import (
	"fmt"
	"loader/internal/domain/entity"
	"loader/internal/infrastructure/formating"
	"loader/internal/infrastructure/gitlab"
	"loader/internal/infrastructure/utils"
	"log"
	"os"
	"sync"
)

func DownloadValidate() error {
	if entity.CFG["addr"] == nil {
		return fmt.Errorf("адрес GitLab не указан")
	}

	if utils.IsEmpty(entity.TOKEN) {
		if utils.IsEmpty(entity.CFG["token"].(string)) {
			return fmt.Errorf(formating.LogError("токен не указан"))
		}
		entity.TOKEN = entity.CFG["token"].(string)
	}

	return nil
}

func DownloadProjects() {
	inProj, outProj := utils.CreateChannelPair[*entity.Proj]()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; ; i++ {
			end, err := gitlab.ReqGetProj(inProj, entity.CFG["addr"].(string), entity.TOKEN, i)
			if err != nil {
				log.Printf("[UC][DOWNLOAD_PROJECTS]%s", err)
				return
			}
			if end {
				break
			}
		}
	}()

	go func() {
		wg.Wait()
		close(inProj)
	}()

	err := os.Mkdir(utils.NormalizePath(entity.CFG["output-dir"].(string)), 0777)
	if err != nil {
		log.Printf("[UC][DOWNLOAD_PROJECTS] %s", formating.LogError(err.Error()))
		return
	}

	wgWorkers := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wgWorkers.Add(1)
		go downloadWorker(wgWorkers, outProj)
	}

	wgWorkers.Wait()
}

func downloadWorker(wg *sync.WaitGroup, outProj <-chan *entity.Proj) {
	defer wg.Done()

	for p := range outProj {
		branches, err := gitlab.ReqGetAllBranches(p.Links.Branches, entity.TOKEN)
		if err != nil {
			log.Println("[DOWNLOAD][GET BRANCHES][ERROR]: ", err)
			continue
		}

		resFolderNameNew := fmt.Sprintf("%s/%s", utils.NormalizePath(entity.CFG["output-dir"].(string)), p.Name)

		err = os.Mkdir(resFolderNameNew, 0777)
		if err != nil {

			return
		}

		for _, branch := range branches {
			err = gitlab.ReqDownloadProj(entity.TOKEN, resFolderNameNew, p.Name, branch, p.Links.Self)
			if err != nil {
				log.Println("[DOWNLOAD][ERROR]: ", err)
				continue
			}
		}
	}
}

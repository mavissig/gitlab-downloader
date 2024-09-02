package loader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (l *Loader) reqGetProj(page int) (bool, error) {
	baseUrl := fmt.Sprintf("%s/projects", l.cfg.Url)

	u, err := url.Parse(baseUrl)
	if err != nil {
		return false, err
	}

	params := url.Values{}
	params.Add("page", fmt.Sprintf("%d", page))

	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", l.cfg.Token))

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if string(b) == "[]" {
		return true, nil
	}

	p := []*proj{}

	err = json.Unmarshal(b, &p)
	if err != nil {
		return false, err
	}

	l.projects = append(l.projects, p...)

	return false, nil
}

func (l *Loader) reqDownloadProj(resFolderName, projName, branch, self string) error {

	baseUrl := fmt.Sprintf("%s/repository/archive.zip", self)

	u, err := url.Parse(baseUrl)
	if err != nil {
		return err
	}

	params := url.Values{}
	if branch != "" {
		params.Add("sha", branch)
	}

	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", l.cfg.Token))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dstFileName := fmt.Sprintf("%s/%s", resFolderName, projName)

	validBranchName := strings.Replace(branch, "/", "-", -1)

	if branch != "" {
		dstFileName = fmt.Sprintf("%s-%s.zip", dstFileName, validBranchName)
	} else {
		dstFileName = fmt.Sprintf("%s.zip", dstFileName)
	}

	f, err := os.Create(dstFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (l *Loader) reqGetAllBranches(baseUrl string) ([]string, error) {

	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", l.cfg.Token))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var branches []struct {
		BranchName string `json:"name"`
	}

	err = json.Unmarshal(b, &branches)
	if err != nil {
		return nil, err
	}

	var res []string

	for _, s := range branches {
		res = append(res, s.BranchName)
	}

	return res, nil
}

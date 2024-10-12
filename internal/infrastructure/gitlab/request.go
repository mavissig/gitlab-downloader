package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"loader/internal/infrastructure/formating"
	"net/http"
	"net/url"
	"os"
	"strings"

	"loader/internal/domain/entity"
)

func ReqGetProj(inProj chan<- *entity.Proj, addr, token string, page int) (bool, error) {
	baseUrl := strings.Join([]string{addr, "projects"}, "/")

	// todo: debug
	fmt.Println(baseUrl)

	u, err := url.Parse(baseUrl)
	if err != nil {
		return false, fmt.Errorf("[REQUEST] %s", formating.LogError("ReqGetProj parsing url:", err.Error()))
	}

	params := url.Values{}
	params.Add("page", fmt.Sprintf("%d", page))

	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return false, fmt.Errorf("[REQUEST] %s", formating.LogError("ReqGetProj creating request:", err.Error()))
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("[REQUEST] %s", formating.LogError("ReqGetProj sending request:", err.Error()))
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("[REQUEST] %s", formating.LogError("ReqGetProj reading response body:", err.Error()))
	}

	if string(b) == "[]" {
		return true, nil
	}

	p := []*entity.Proj{}

	err = json.Unmarshal(b, &p)
	if err != nil {
		return false, fmt.Errorf("[REQUEST] %s", formating.LogError("ReqGetProj unmarshalling response body:", err.Error()))
	}

	for _, proj := range p {
		inProj <- proj
	}

	return false, nil
}

func ReqDownloadProj(token, resFolderName, projName, branch, self string) error {

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

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

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

func ReqGetAllBranches(baseUrl, token string) ([]string, error) {

	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

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

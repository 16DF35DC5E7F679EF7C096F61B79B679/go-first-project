package googleSearch

import (
	"encoding/json"
	"fmt"
	"go-first-project/dependencies"
	"io/ioutil"
	"net/http"
)


func Search(searchQuery string, limit string, googleSearchCredentials *dependencies.GoogleSearchCredentials) (*ResultDTO, error) {
	fmt.Printf(googleSearchCredentials.ToString())
	searchUrl := "https://www.googleapis.com/customsearch/v1"
	httpMethod := "GET"
	httpClient := &http.Client{}
	req, err := http.NewRequest(httpMethod, searchUrl, nil)
	if err != nil {
		fmt.Printf("Error in creating new request for searching artist %e", err)
		return nil, err
	}
	query := req.URL.Query()
	query.Set("q", searchQuery)
	query.Set("key", googleSearchCredentials.ApiKey)
	query.Set("cx", googleSearchCredentials.SearchEngineId)
	query.Set("num", limit)
	req.URL.RawQuery = query.Encode()
	resp, err := httpClient.Do(req)
	fmt.Printf("Searching google for %s ", searchQuery)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in reading response body %e", err)
		return nil, err
	}
	bodyStr := string(bodyText)
	fmt.Printf("Response is %s ", bodyStr)
	var searchResultAPIResponse SearchResultAPIResponseDTO
	jsonError := json.Unmarshal(bodyText, &searchResultAPIResponse)
	//jsonError := json.NewDecoder(resp.Body).Decode(searchResultAPIResponse)
	if jsonError != nil {
		fmt.Printf("Error in unmarshalling response body text %e ", jsonError)
		return nil, err
	}
	fmt.Println("API response is ", searchResultAPIResponse)
	return searchResultAPIResponse.toResultDTO()

}

type ItemResultDTO struct {
	Kind string
	Title string
	HtmlTitle string
	Link string
	DisplayLink string
	Snippet string
}

func (itemResultDTO *ItemResultDTO) ToString() string {
	return "ItemResultDTO: [Kind: "+ itemResultDTO.Kind + " Title: " + itemResultDTO.Title + "]"
}

type ResultDTO struct {
	SearchQuery string
	Items []ItemResultDTO
}

func (resultDTO *ResultDTO) ToString() string{
	toString :=  "ResultDTO=[\n"
	for index, itemResultDTO := range resultDTO.Items {
		toString = toString + string(index) + ". " + itemResultDTO.ToString() + "\n"
	}
	toString = toString + " SearchQuery: "+ resultDTO.SearchQuery + "]"
	return toString
}

func (searchResultResponse *SearchResultAPIResponseDTO) toResultDTO() (*ResultDTO, error) {
	items := searchResultResponse.Items
	itemResults := make([]ItemResultDTO, len(items))
	for index, item := range items {
		itemResults[index] = ItemResultDTO{
			Kind: item.Kind,
			Title: item.Title,
			HtmlTitle: item.HtmlSnippet,
			Link: item.Link,
			DisplayLink: item.DisplayLink,
			Snippet: item.Snippet,
		}
	}
	return &ResultDTO{
		SearchQuery: searchResultResponse.Queries.Request[0].SearchTerms,
		Items: itemResults,
	}, nil
}
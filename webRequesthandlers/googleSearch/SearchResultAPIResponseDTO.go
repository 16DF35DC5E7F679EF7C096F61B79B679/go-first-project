package googleSearch

type URLSearchResultAPIResponseDTO struct {

}

type QueriesSearchResultAPIResponseDTO struct {
	Request []RequestSearchResultAPIResponseDTO `json:"request"`
	NextPage []NextPageSearchResultAPIResponseDTO `json:"nextPage"`
}

type RequestSearchResultAPIResponseDTO struct {
	Title string `json:"title"`
	TotalResults string `json:"totalResults"`
	SearchTerms string `json:"searchTerms"`
	Count int `json:"count"`
	StartIndex int `json:"startIndex"`
	Safe string `json:"safe"`
}

type NextPageSearchResultAPIResponseDTO struct {
	Title string `json:"title"`
	TotalResults string `json:"totalResults"`
	SearchTerms string `json:"searchTerms"`
	Count int `json:"count"`
	StartIndex int `json:"startIndex"`
	Safe string `json:"safe"`
}

type ContextSearchAPIResponseDTO struct {
	Title string `json:"title"`
}

type SearchInformationSearchAPIResponseDTO struct {

}

type ResultItemSearchAPIResponseDTO struct {
	Kind string `json:"kind"`
	Title string `json:"title"`
	HtmlTitle string `json:"htmlTitle"`
	Link string`json:"link"`
	DisplayLink string `json:"displayLink"`
	Snippet string `json:"snippet"`
	HtmlSnippet string `json:"htmlSnippet"`
}

type SearchResultAPIResponseDTO struct {
	Items []ResultItemSearchAPIResponseDTO `json:"items"`
	Context ContextSearchAPIResponseDTO `json:"context"`
	Queries QueriesSearchResultAPIResponseDTO `json:"queries"`
}

//func (searchResultResponse *SearchResultAPIResponseDTO) ToString() string {
//
//}
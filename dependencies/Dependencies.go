package dependencies

type GoogleSearchCredentials struct {
	ApiKey string
	SearchEngineId string
}

func (googleSearchCredentials *GoogleSearchCredentials) ToString() string {
	if googleSearchCredentials == nil {
		return "[GoogleSearchCredentials=null]"
	}
	return "GoogleSearchCredentials=[apiKey: " + googleSearchCredentials.ApiKey + " searchEngineId: " + googleSearchCredentials.SearchEngineId + "]"
}

type App struct{
	GoogleSearchCredentials *GoogleSearchCredentials
}

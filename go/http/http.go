package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func Get[T any](path string, responseBodyKey string, result *T) error {
    url := fmt.Sprintf("https://wttcmsapigateway-new.azure-api.net/%s", path)
	
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var tempMap map[string]interface{}
    err = json.Unmarshal(body, &tempMap)
    if err != nil {
        return err
    }

    keyValue, ok := tempMap[responseBodyKey]
    if !ok {
        return fmt.Errorf("key %s not found in response body", responseBodyKey)
    }

    keyValueBytes, err := json.Marshal(keyValue)
    if err != nil {
        return err
    }

    err = json.Unmarshal(keyValueBytes, result)
    if err != nil {
        return err
    }

    return nil
}
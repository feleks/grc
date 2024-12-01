package main

import (
	"encoding/base64"
	"fmt"
)

var preparedAssets map[string][]byte = nil

func prepareAssets() error {
	if preparedAssets != nil {
		return nil
	}

	preparedAssets = make(map[string][]byte, len(assetsContent))
	for k, v := range assetsContent {
		content, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return fmt.Errorf("failed to decode asset %s: %v", k, err)
		}
		preparedAssets[k] = content
	}
	return nil
}

func getAsset(path string) ([]byte, error) {
	err := prepareAssets()
	if err != nil {
		return nil, fmt.Errorf("failed to get asset %s: %v", path, err)
	}

	content, ok := preparedAssets[path]
	if !ok {
		return nil, fmt.Errorf("failed to get asset %s: it is not registered", path)
	}

	return content, nil
}

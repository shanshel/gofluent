package gofluent

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Lang struct {
	dirPath     string
	defaultLang string
}

var languages map[string]map[interface{}]interface{} = make(map[string]map[interface{}]interface{})
var activeLang string

func (l *Lang) Setup(dirPath, defaultLang string, preload bool) error {
	_, err := os.Stat(dirPath)
	if err != nil {
		return err
	}
	l.dirPath = dirPath
	l.defaultLang = defaultLang

	if preload {
		files, err := ioutil.ReadDir(dirPath)
		if err != nil {
			return err
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}
			fileName := file.Name()
			if filepath.Ext(fileName) == ".yaml" {
				l.Switch(fileName[:len(fileName)-len(filepath.Ext(fileName))])
			}
		}

	}

	err = l.Switch(defaultLang)

	if err != nil {
		return err
	}

	return nil

}

func (l *Lang) Switch(lang string) error {

	if _, isLoaded := languages[lang]; isLoaded {
		activeLang = lang
		return nil
	}

	langFilePath := path.Clean(l.dirPath + "/" + lang + ".yaml")
	if _, err := os.Stat(langFilePath); err != nil {
		return err
	}

	filename, err := filepath.Abs(langFilePath)
	if err != nil {
		return err
	}

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	languages[lang] = make(map[interface{}]interface{})

	err = yaml.Unmarshal(fileContent, languages[lang])
	if err != nil {
		return err
	}

	activeLang = lang

	return nil
}

func (l *Lang) Get(section, key string, params ...interface{}) string {

	//1.Use Default if Active Lang is empty
	if activeLang == "" {
		activeLang = l.defaultLang
	}

	//2. Return nothing if key is ""
	if key == "" {
		return ""
	}

	//3. If the active lang not loaded return ""
	if _, exist := languages[activeLang]; !exist {
		return ""
	}

	//4. When Section provided
	if section != "" {

		if sectionInterface, exist := languages[activeLang][section]; exist {

			if sectionMap, ok := sectionInterface.(map[string]interface{}); ok {
				if sentance, exist := sectionMap[key]; exist {
					if sentanceString, ok := sentance.(string); ok {
						return fmt.Sprintf(sentanceString, params...)

					}
				}
			}
		}
	} else {
		if sentanceInterface, exist := languages[activeLang][key]; exist {
			if sentance, ok := sentanceInterface.(string); ok {
				return fmt.Sprintf(sentance, params...)
			}
		}
	}

	return key

}

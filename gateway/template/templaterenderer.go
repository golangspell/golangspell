package template

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/danilovalente/golangspell/config"
	"github.com/danilovalente/golangspell/domain"
)

//Renderer is reponsible for executing and rendering the code templates
type Renderer struct {
	rootTemplatePath  string
	currentPath       string
	globalVariables   map[string]interface{}
	specificVariables map[string]map[string]interface{}
}

//mergeVariables - specific overrides global if the name is the same
func (renderer *Renderer) mergeVariables(fileName string) map[string]interface{} {
	allVariables := make(map[string]interface{}, 0)
	for key, val := range renderer.globalVariables {
		allVariables[key] = val
	}
	for key, val := range renderer.specificVariables {
		allVariables[key] = val
	}
	return allVariables
}

func (renderer *Renderer) backupExistingCode(sourcePath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()
	destination, err := os.Create(fmt.Sprintf("%s_backup_%d", sourcePath, time.Now().Unix()))
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

//RenderFile renders a template file
func (renderer *Renderer) RenderFile(sourcePath string, info os.FileInfo) error {
	fileName := filepath.Base(sourcePath)
	fmt.Printf("Rendering template: %s\n", fileName)
	allVariables := renderer.mergeVariables(fileName)
	destinationPath := strings.ReplaceAll(strings.ReplaceAll(sourcePath, renderer.rootTemplatePath, renderer.currentPath), ".got", ".go")
	if destFileInfo, err := os.Stat(destinationPath); err == nil && !destFileInfo.IsDir() {
		err := renderer.backupExistingCode(destinationPath)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer file.Close()
	tmpl, err := template.New(fileName).ParseFiles(sourcePath)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	err = tmpl.Execute(w, allVariables)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

func (renderer *Renderer) createDirectory(sourcePath string) error {
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return os.Mkdir(sourcePath, 0700)
	}
	return nil
}

//RenderPath renders an object (file os directory) in the templates directory
func (renderer *Renderer) RenderPath(sourcePath string, info os.FileInfo, err error) error {
	if err != nil {
		log.Printf("An error occurred while trying to analyze path %s\n", err)
		return nil
	}
	if info.IsDir() {
		if sourcePath == renderer.rootTemplatePath {
			return nil
		}
		destinationPath := strings.ReplaceAll(sourcePath, renderer.rootTemplatePath, renderer.currentPath)
		err := renderer.createDirectory(destinationPath)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	return renderer.RenderFile(sourcePath, info)
}

//RenderTemplate renders all templates in the template directory providing the respective variables
//spell: specifies the spell which contains the command
//commandName: specifies the name of the command for which the template will be rendered
//variables: specifies the list of variables (value) which should be provided for rendering each file (key)
func (renderer *Renderer) RenderTemplate(spell domain.Spell, commandName string, globalVariables map[string]interface{}, specificVariables map[string]map[string]interface{}) error {
	spellInstallation := domain.GolangLibrary{Name: spell.Name, URL: spell.URL}
	renderer.rootTemplatePath = fmt.Sprintf("%s%stemplates%s%s", spellInstallation.SrcPath(), config.PlatformSeparator, config.PlatformSeparator, commandName)
	if nil == globalVariables {
		renderer.globalVariables = make(map[string]interface{}, 0)
	} else {
		renderer.globalVariables = globalVariables
	}
	if nil == specificVariables {
		renderer.specificVariables = make(map[string]map[string]interface{}, 0)
	} else {
		renderer.specificVariables = specificVariables
	}
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}
	renderer.currentPath = currentPath
	err = filepath.Walk(renderer.rootTemplatePath, renderer.RenderPath)
	return err
}

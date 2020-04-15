package template

import (
	"bufio"
	"bytes"
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

const stringtemplatesdirectory = "notrenderedstringtemplates"

//Renderer is reponsible for executing and rendering the code templates
type Renderer struct {
	rootTemplatePath   string
	stringTemplatePath string
	currentPath        string
	globalVariables    map[string]interface{}
	specificVariables  map[string]map[string]interface{}
}

//mergeVariables - specific overrides global if the name is the same
func (renderer *Renderer) mergeVariables(fileName string) map[string]interface{} {
	allVariables := make(map[string]interface{})
	for key, val := range renderer.globalVariables {
		allVariables[key] = val
	}
	for key, val := range renderer.specificVariables {
		allVariables[key] = val
	}
	return allVariables
}

//BackupExistingCode make a copy of the changed file
func (renderer *Renderer) BackupExistingCode(sourcePath string) error {
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

//RenderString processing the provided template source file, using the provided variables
func (renderer *Renderer) RenderString(spell domain.Spell, commandName string, stringTemplateFileName string, variables map[string]interface{}) (string, error) {
	spellInstallation := domain.GolangLibrary{Name: spell.Name, URL: spell.URL}
	renderer.rootTemplatePath = fmt.Sprintf("%s%stemplates%s%s", spellInstallation.SrcPath(), config.PlatformSeparator, config.PlatformSeparator, commandName)
	renderer.stringTemplatePath = fmt.Sprintf("%s%s%s", renderer.rootTemplatePath, config.PlatformSeparator, stringtemplatesdirectory)
	tmpl, err := template.New(stringTemplateFileName).ParseFiles(fmt.Sprintf("%s%s%s", renderer.stringTemplatePath, config.PlatformSeparator, stringTemplateFileName))
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, variables)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

//RenderFile renders a template file
func (renderer *Renderer) RenderFile(sourcePath string, info os.FileInfo) error {
	fileName := filepath.Base(sourcePath)
	fmt.Printf("Rendering template: %s\n", fileName)
	allVariables := renderer.mergeVariables(fileName)
	rootTemplatePath := renderer.rootTemplatePath
	currentPath := renderer.currentPath
	sourcePathEscaped := sourcePath
	if config.PlatformIsWindows {
		rootTemplatePath = strings.ReplaceAll(rootTemplatePath, "\\", "/")
		currentPath = strings.ReplaceAll(currentPath, "\\", "/")
		sourcePathEscaped = strings.ReplaceAll(sourcePathEscaped, "\\", "/")
	}
	destinationPath := strings.ReplaceAll(strings.ReplaceAll(sourcePathEscaped, rootTemplatePath, currentPath), ".got", ".go")
	if config.PlatformIsWindows {
		destinationPath = strings.ReplaceAll(destinationPath, "/", "\\")
	}
	if destFileInfo, err := os.Stat(destinationPath); err == nil && !destFileInfo.IsDir() {
		err := renderer.BackupExistingCode(destinationPath)
		if err != nil {
			return err
		}
	}
	directory := filepath.Dir(destinationPath)
	if strings.Contains(directory, stringtemplatesdirectory) {
		return nil
	}
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
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
	if strings.Contains(sourcePath, stringtemplatesdirectory) {
		return nil
	}
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
	renderer.stringTemplatePath = fmt.Sprintf("%s%s%s", renderer.rootTemplatePath, config.PlatformSeparator, stringtemplatesdirectory)
	if nil == globalVariables {
		renderer.globalVariables = make(map[string]interface{})
	} else {
		renderer.globalVariables = globalVariables
	}
	if nil == specificVariables {
		renderer.specificVariables = make(map[string]map[string]interface{})
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

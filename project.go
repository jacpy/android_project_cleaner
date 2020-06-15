package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func findProject(p string, fs []os.FileInfo) error {
	oldProject, newProject := isProject(fs)
	if oldProject && newProject {
		if newProject {
			for _, f := range fs {
				name := f.Name()
				if f.IsDir() {
					filePath := filepath.Join(p, name)
					if dirs, err := ioutil.ReadDir(filePath); err != nil {
						log.Println(err)
					} else if isModule(dirs) {
						fp := filepath.Join(p, name, "build")
						if isExists(fp) {
							if err := os.RemoveAll(fp); err != nil {
								log.Println("remove project path ", fp, " failure")
								log.Println(err)
							}
						}
					}
				}
			}
		} else if oldProject {
			fp := filepath.Join(p, "bin")
			if isExists(fp) {
				if err := os.RemoveAll(fp); err != nil {
					log.Println("remove project path ", fp, " failure")
					log.Println(err)
				}
			}
		}

		capturePath := filepath.Join(p, "captures")
		if isExists(capturePath) {
			err := os.RemoveAll(capturePath)
			if err != nil {
				log.Println("remove path ", capturePath, " failure")
				log.Println(err)
			}
		}

		screenshotPath := filepath.Join(p, "screenshots")
		if isExists(screenshotPath) {
			err := os.RemoveAll(screenshotPath)
			if err != nil {
				log.Println("remove path ", screenshotPath, " failure")
				log.Println(err)
			}
		}
	} else {
		for _, f := range fs {
			name := f.Name()
			if f.IsDir() {
				s := filepath.Join(p, name)
				if children, err := ioutil.ReadDir(s); err != nil {
					log.Println("read dir ", s, " failure")
					log.Println(err)
				} else if err := findProject(s, children); err != nil {
					log.Println("find project error:")
					log.Println(err)
				}
			}
		}
	}

	return nil
}

func isProject(fs []os.FileInfo) (bool, bool) {
	hasGradleProperties := false
	hasBuildGradle := false
	hasGradlew := false

	hasSrc := false
	hasManifest := false

	for _, dir := range fs {
		name := dir.Name()
		if dir.IsDir() {
			if name == "src" {
				hasSrc = true
			}
		} else {
			if name == "gradle.properties" {
				hasGradleProperties = true
			} else if name == "build.gradle" {
				hasBuildGradle = true
			} else if name == "gradlew" {
				hasGradlew = true
			} else if name == "AndroidManifest.xml" {
				hasManifest = true
			}
		}
	}

	return hasSrc && hasManifest, hasGradleProperties && hasBuildGradle && hasGradlew
}

func isModule(fs []os.FileInfo) bool {
	hasSrc := false
	hasBuildGradle := false
	for _, f := range fs {
		name := f.Name()
		if name == "src" && f.IsDir() {
			hasSrc = true
		} else if name == "build.gradle" && !f.IsDir() {
			hasBuildGradle = true
		}
	}

	return hasSrc && hasBuildGradle
}

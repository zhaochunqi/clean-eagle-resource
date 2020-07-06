package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func FixImagePath(EagleRootPath string, dryRun bool) {

	imageFolder := path.Join(EagleRootPath, "images")
	files, err := ioutil.ReadDir(imageFolder)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Printf("index is %v \n", i)
		if file.IsDir() {
			dirPath := path.Join(imageFolder, file.Name())
			imageFiles, err := ioutil.ReadDir(dirPath)

			if err != nil {
				log.Fatal(err)
			}

			fixImageFiles(&imageFiles, dirPath, dryRun)
		} else {
			fmt.Println("Skipping " + file.Name())
		}
	}

}

func fixImageFiles(files *[]os.FileInfo, dirPath string, dryRun bool) {

	metadataPath := path.Join(dirPath, "metadata.json")

	//如果 metadata 不存在，删除文件夹
	_, err := os.Stat(metadataPath)
	if os.IsNotExist(err) || len(*files) <= 1{
		fmt.Printf("Metadata not exist or only contains metadata.\n")
		fmt.Printf("Removing %s \n", dirPath)

		if !dryRun {
			err := os.RemoveAll(dirPath)

			if err != nil {
				log.Fatal(err)
			}
		}
		return
	}

	jsonFile, err := os.Open(metadataPath)

	defer jsonFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	var metadata Metadata

	json.Unmarshal(bytes, &metadata)

	for _, file := range *files {
		//fmt.Printf("FileName is %s \n", file.Name())

		//正确的应该是 xxx.ext
		//以及 xxx_thumbnail.png
		//和 metadata
		if file.Name() == "metadata.json" {
			continue
		}

		ext := filepath.Ext(file.Name())
		nameWithoutExt := strings.TrimSuffix(file.Name(), ext)

		if ext == ".png" || ext == ".jpg" {
			if strings.HasSuffix(nameWithoutExt, "_thumbnail") {
				fileNameWithoutSuffix := strings.TrimSuffix(nameWithoutExt, "_thumbnail")

				if metadata.Name == fileNameWithoutSuffix && ext != ".jpg"{
					continue
				}

				oldName := path.Join(dirPath, file.Name())

				fmt.Printf("ext is %s \n", ext)
				if metadata.NoThumbnail || ext == ".jpg" {
					fmt.Printf("Removing %s .\n", oldName)
					if !dryRun {
						err := os.Remove(oldName)
						if err != nil {
							log.Fatal(err)
						}
					}
					continue
				}

				newName := path.Join(dirPath, metadata.Name + "_thumbnail." + metadata.Ext)
				fmt.Printf("Renaming %s to %s. \n", oldName,newName)

				if !dryRun {
					err := os.Rename(oldName, newName)
					if err != nil {
						log.Fatal(err)
					}
				}
				continue
				//删除缩略图
				//filePath := path.Join(dirPath, file.Name())
				//fmt.Printf("Deleting %s \n", filePath)
			}
		}

		if ext == "."+metadata.Ext {
			if file.Name() == metadata.Name+"."+metadata.Ext {
				continue
			}

			oldName := path.Join(dirPath, file.Name())
			newName := path.Join(dirPath, metadata.Name + "." + metadata.Ext)
			fmt.Printf("Renaming %s to %s. \n", oldName,newName)
			if !dryRun {
				err := os.Rename(oldName, newName)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			tempFilePath := path.Join(dirPath, file.Name())
			fmt.Printf("Deleting file %s \n", tempFilePath)
			if !dryRun {
				err := os.Remove(tempFilePath)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

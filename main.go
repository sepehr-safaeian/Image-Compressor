package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/moemoe89/go-helpers/pkg/imagecompressor/tinify"
)

func main() {
	var targetFolder string
	asciiArt := `
	┳           ┏┓                 
	┃┏┳┓┏┓┏┓┏┓  ┃ ┏┓┏┳┓┏┓┏┓┏┓┏┏┏┓┏┓
	┻┛┗┗┗┻┗┫┗   ┗┛┗┛┛┗┗┣┛┛ ┗ ┛┛┗┛┛ 
		   ┛           ┛           
`
	displayASCIIArt(asciiArt)

	fmt.Println("Location Folder Address:")
	fmt.Scan(&targetFolder)

	compressor(targetFolder)

}

// Compression images
func compressor(targetFolder string) {
	// Create Tinify clients using an API Key
	client, err := tinify.New(
		tinify.WithAPIKey("<YOUR_API>"),
		tinify.WithHTTPClient(
			&http.Client{
				Transport:     nil,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       0,
			},
		),
	)
	if err != nil {
		panic(err)
	}

	// Open the folder
	folder, err := os.Open(targetFolder)
	if err != nil {
		panic(err)
	}
	defer folder.Close()

	// Read all files in the folder
	fileInfos, err := folder.Readdir(-1)
	if err != nil {
		panic(err)
	}

	compressedFolder := filepath.Join(targetFolder, "Compressed")
	err = os.Mkdir(compressedFolder, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Iterate through each file in the folder
	for _, fileInfo := range fileInfos {
		// Construct the full path to the file
		filePath := filepath.Join(targetFolder, fileInfo.Name())

		// Open the image file for compression
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Compress the image
		result, err := client.Upload(context.Background(), file, "")
		if err != nil {
			panic(err)
		}

		// Print the URL of the compressed file
		fmt.Println("✅", fileInfo.Name())

		// Create the path for the compressed file in the "compressed" folder
		compressedFilePath := filepath.Join(compressedFolder, fileInfo.Name())

		// Download the compressed image
		response, err := http.Get(result.URL)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		// Create a new file in the "compressed" folder
		compressedFileOnDisk, err := os.Create(compressedFilePath)
		if err != nil {
			panic(err)
		}
		defer compressedFileOnDisk.Close()

		// Copy the compressed image data to the new file
		_, err = io.Copy(compressedFileOnDisk, response.Body)
		if err != nil {
			panic(err)
		}
	}
}

func displayASCIIArt(asciiArt string) {
	fmt.Println(asciiArt)
}

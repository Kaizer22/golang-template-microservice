package utils

import (
	"encoding/json"
	"github.com/golang/snappy"
	"github.com/gorilla/schema"
	"io/ioutil"
	"main/logging"
	"net/http"
	"os"
	"strconv"
)

var (
	SupportedFiletypes = []string{"csv"}
)

// Decoder: request parameters decoder
var Decoder = schema.NewDecoder()

func RespondWithOutputImage(w http.ResponseWriter, status int, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Compressed", "false")
	w.Write(fileBytes)

	err = file.Close()
	if err != nil {
		logging.InfoFormat("cannot close file %s: %s", filename, err)
	}
	err = os.Remove(filename)
	if err != nil {
		logging.InfoFormat("cannot delete file %s: %s", filename, err)
	}
}
func RespondWithMultipart(w http.ResponseWriter, status int, filenames []string) {
	rawFiles := make([][]byte, 0, len(filenames))
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		rawFiles = append(rawFiles, fileBytes)
		file.Close()
	}
	var contentDescriptor = ""
	var totalContent []byte
	for i, value := range rawFiles {
		totalContent = append(totalContent, value...)
		err := os.Remove(filenames[i])
		if err != nil {
			logging.InfoFormat("cannot delete file %s: %s", filenames[i], err)
		}
		l := strconv.Itoa(len(value)) + "_"
		contentDescriptor += l
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("My-Content-Descriptor", contentDescriptor)
	//w.WriteHeader(status)
	if _, err := w.Write(totalContent); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func RespondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logging.ErrorFormat("Unexpected error while marshalling:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Compressed", "false")
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

func RespondWithCompressedJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logging.ErrorFormat("Unexpected error while marshalling:", err)
	}
	response = snappy.Encode(response, response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Compressed", "true")
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

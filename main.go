package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"os/exec"
)

func init() {
	mime.AddExtensionType(".mpd", "application/dash+xml")
}

func convertMP4ToDASH(inputFile, outputDir string) error {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Command to convert MP4 to DASH using Bento4's mp4dash tool
	cmd := exec.Command("mp4dash", "--output-dir", outputDir, inputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func dashHandler(w http.ResponseWriter, r *http.Request) {
	inputFile := "abc.mkv" // Path to the input MP4 file
	outputDir := "output"  // Directory to store the DASH files

	err := convertMP4ToDASH(inputFile, outputDir)
	if err != nil {
		http.Error(w, "Failed to convert MP4 to DASH", http.StatusInternalServerError)
		log.Println("Conversion error:", err)
		return
	}

	// Serve the DASH files
	http.StripPrefix("/dash/", http.FileServer(http.Dir(outputDir))).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/convert", dashHandler)
	http.Handle("/dash/", http.StripPrefix("/dash/", http.FileServer(http.Dir("output"))))

	// 提供HTML文件服务
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

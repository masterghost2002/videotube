package main

import (
	"log"
	"os"

	ffmpeg "github.com/masterghost2002/videotube/cmd/transcoder/ffpeg"
)

func main() {
	// Input MP4 file path
	inputFilePath := "/home/rakesh/data/codespace/videotube/transcoder_output/input.mp4"

	// Output directory for HLS files
	outputDir := "/home/rakesh/data/codespace/videotube/transcoder_output/output/"

	// HLS options
	hlsOptions := []string{
		"-hls_time", "10", // Segment duration (in seconds)
		"-hls_playlist_type", "vod", // Playlist type (default is event)
		"-hls_segment_filename", outputDir + "video_%v_%03d.ts", // Segment filename pattern
		"-f", "hls", // Output format is HLS
	}

	// Quality levels and resolutions
	resolutions := map[string]string{
		"720p": "1280x720",
		"480p": "854x480",
		"360p": "640x360",
	}

	// Bitrates for each resolution (in Mbps)
	bitrates := map[string][2]int{
		"720p": {5000, 4500}, // {peak bitrate, average bitrate} in kbps
		"480p": {2500, 2250}, // {peak bitrate, average bitrate} in kbps
		"360p": {1000, 900},  // {peak bitrate, average bitrate} in kbps
	}

	// Create output directory if not exists
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Check if input file exists
	if _, err := os.Stat(inputFilePath); os.IsNotExist(err) {
		log.Fatalf("Input file %s does not exist", inputFilePath)
	}
	masterPlaylistPath := outputDir + "master_playlist.m3u8"
	if err := ffmpeg.Convert(masterPlaylistPath, inputFilePath, outputDir, hlsOptions, resolutions, bitrates); err != nil {

	}

}

package ffmpeg

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Convert(masterPlaylistPath string, inputFilePath string, outputDir string, hlsOptions []string, resolutions map[string]string, bitrates map[string][2]int) error {

	// Generate HLS playlists for each resolution
	for name, resolution := range resolutions {
		playlistPath := outputDir + "playlist_" + name + ".m3u8"
		cmdArgs := []string{
			"-i", inputFilePath,
			"-vf", "scale=" + resolution,
		}
		cmdArgs = append(cmdArgs, hlsOptions...)
		cmdArgs = append(cmdArgs, playlistPath)

		cmd := exec.Command("ffmpeg", cmdArgs...)

		// Execute ffmpeg command
		if _, err := cmd.CombinedOutput(); err != nil {
			return err
		}

		log.Printf("Generated HLS playlist for %s", name)
	}

	// Generate master playlist manually
	masterPlaylistContent := "#EXTM3U\n"
	for name, bitrate := range bitrates {
		peakBandwidth := bitrate[0] * 1000    // Convert to bps
		averageBandwidth := bitrate[1] * 1000 // Convert to bps
		playlistPath := "playlist_" + name + ".m3u8"
		masterPlaylistContent += fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,AVERAGE-BANDWIDTH=%d,RESOLUTION=%s\n%s\n", peakBandwidth, averageBandwidth, resolutions[name], playlistPath)
	}

	if err := os.WriteFile(masterPlaylistPath, []byte(masterPlaylistContent), 0644); err != nil {
		return err
	}

	log.Printf("Generated master playlist: %s", masterPlaylistPath)
	return nil
}

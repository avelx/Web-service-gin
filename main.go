package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"Title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type track struct {
	TrackId string `json:"TrackId"`
	AlbumId string `json:"AlbumId"`
	Title   string `json:"Title"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var allTracks = []track{
	{TrackId: "1", AlbumId: "23", Title: "test"},
	{TrackId: "2", AlbumId: "23", Title: "test"},
	{TrackId: "3", AlbumId: "23", Title: "test"},
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func processCsvFile(records *[][]string) []track {
	rec := *records
	var recs = make([]track, len(rec))
	for i := 0; i < len(rec); i++ {
		//id, _ := strconv.Atoi(rec[i][0])
		var r = track{
			TrackId: rec[i][0],
			AlbumId: rec[i][1],
			Title:   rec[i][2],
		}
		recs[i] = r
	}
	return recs
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.GET("/tracks", getTracks)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getTracks(c *gin.Context) {
	records := readCsvFile("data/tracks.csv")
	recs := processCsvFile(&records)
	fmt.Println("Record: ", recs)
	c.IndentedJSON(http.StatusOK, recs)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

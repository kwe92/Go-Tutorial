package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

var image_urls []string = []string{
	"https://images.unsplash.com/photo-1516117172878-fd2c41f4a759",
	"https://images.unsplash.com/photo-1532009324734-20a7a5813719",
	"https://images.unsplash.com/photo-1524429656589-6633a470097c",
	"https://images.unsplash.com/photo-1530224264768-7ff8c1789d79",
	"https://images.unsplash.com/photo-1564135624576-c5c88640f235",
	"https://images.unsplash.com/photo-1541698444083-023c97d3f4b6",
	"https://images.unsplash.com/photo-1522364723953-452d3431c267",
	// "https://images.unsplash.com/photo-1513938709626-033611b8cc03",
	// "https://images.unsplash.com/photo-1507143550189-fed454f93097",
	"https://images.unsplash.com/photo-1493976040374-85c8e12f0c0e",
	"https://images.unsplash.com/photo-1504198453319-5ce911bafcde",
	"https://images.unsplash.com/photo-1530122037265-a5f1f91d3b99",
	"https://images.unsplash.com/photo-1516972810927-80185027ca84",
	"https://images.unsplash.com/photo-1550439062-609e1531270e",
	"https://images.unsplash.com/photo-1549692520-acc6669e2f0c",
}

var imgPath string = "/Users/kwe/GoLang/go-Tutorial/"

func main() {

	// declare and initialize a wait group to synchronize the completion of goroutines
	// each goroutine should be passed a pointer to the wait group to ensure the same counter is decremented
	waitGroup := &sync.WaitGroup{}

	// set the counter to the number of goroutines to wait for
	waitGroup.Add(len(image_urls))

	start := time.Now()

	client := &http.Client{Timeout: 30 * time.Minute}

	// execute each download as a goroutine `lightweight thread of execution`
	for index, url := range image_urls {
		go getImage(index+1, url, client, waitGroup)

	}

	// wait for all goroutines to complete
	waitGroup.Wait()

	elapsed := time.Since(start)

	fmt.Println("Download time taken:", elapsed)

}

// getImage: fetches the image of a valid url
func getImage(id int, url string, client *http.Client, wg *sync.WaitGroup) {

	// split url string into an array
	strArray := strings.Split(url, "-")

	fmt.Println(strArray)

	fileNamePart00 := strArray[len(strArray)-2]

	fileNamePart01 := strArray[len(strArray)-1]

	fileName := fmt.Sprintf("photo_%s_%s.jpg", fileNamePart00, fileNamePart01)

	// request image | receive response object
	response, err := client.Get(url)

	checkErr(err)

	defer response.Body.Close()

	// create file if it does not exist otherwise truncate file
	file, err := os.Create(fileName)

	checkErr(err)

	defer file.Close()

	// copy the response body into a file
	_, err = io.Copy(file, response.Body)

	checkErr(err)

	fmt.Println("Download successful | gorutine number:", id)

	err = resizeImg(imgPath + fileName)

	checkErr(err)

	// decrement the wait group counter
	wg.Done()

}

// resizeImg: resizes an image
func resizeImg(imgPath string) error {

	src, err := imaging.Open(imgPath)

	// checkErr(err)

	src = imaging.Resize(src, 480, 800, imaging.Lanczos)

	imaging.Save(src, imgPath)

	fmt.Println("resized image:", imgPath)

	return err
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

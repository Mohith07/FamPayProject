package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	"FamPayProject/clients"
	"FamPayProject/model"
	"FamPayProject/service"
)


func Init() {
	log.Infof("running the cron job")

	searchConfig := service.GetVideoSearchConfig()

	cronConfig := service.GetCronConfig()

	apiKeys := service.GetAPIKeys()

	for _, apiKey := range apiKeys {

		ytService, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey.APIKey))
		if err != nil {
			fmt.Println("Error creating new YouTube client: " + err.Error())
			return
		}

		call := youtube.NewSearchService(ytService).
						List([]string{"id,snippet"}).
						Q(searchConfig.Query).
						MaxResults(searchConfig.MaxResults).
						Order(searchConfig.OrderBy).
						Type(searchConfig.SearchType).
						PublishedAfter(cronConfig.CreatedAt.Format(time.RFC3339))

		response, err := call.Do()
		if err != nil {
			fmt.Println("error while making a call " + err.Error())
			continue
		}

		if response.HTTPStatusCode > 400 && response.HTTPStatusCode < 500 {
			log.Info("quota for the given developer key has exhausted")
			continue
		}

		for _, item := range response.Items {
			id := item.Id.VideoId
			description := item.Snippet.Description
			title := item.Snippet.Title
			url := item.Snippet.Thumbnails.Default.Url
			publishedAt := item.Snippet.PublishedAt
			log.Infof("adding a new video - " + title)
			currVideo := &model.Video{}
			clients.GetDB().Find(currVideo, "name = ?", id)
			if currVideo.Title != "" {
				log.Info("this video has already been added")
				break
			}
			clients.GetDB().Save(&model.Video{Title: title, Description: description, URL: url, PublishedAt: publishedAt, Name: id})
		}
		//add entry in cron job table
		clients.GetDB().Save(&model.YoutubeCronJobRun{
			PublishedAfter: cronConfig.CreatedAt.Format(time.RFC3339),
			APIKey:         apiKey.APIKey,
		})
		break
	}
}
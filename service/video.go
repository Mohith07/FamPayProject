package service

import (
	"math"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	g "gorm.io/gorm"

	"FamPayProject/clients"
	"FamPayProject/config"
	"FamPayProject/model"
)

func GetAllVideos(queryParams url.Values) ([]model.Video, int, error) {
	var videos []model.Video
	var totalCount int64
	query := clients.GetDB().Table(model.Video{}.TableName())
	if err := query.Error; err != nil {
		log.Infof("error fetching data from db " + err.Error())
		return nil, 0, err
	}
	query.Count(&totalCount)
	pageNumber := getPageNumber(queryParams)
	query = query.Order("published_at desc").
		Offset(calculateDbOffset(pageNumber, 20, int(totalCount))).Limit(20)
	if err := query.Find(&videos).Error; err != nil {
		return videos, 0, err
	}
	return videos, pageNumber, nil
}

func SearchVideos(queryParams url.Values) ([]model.Video, int, error) {
	var videos []model.Video
	var totalCount int64
	query := clients.GetDB().Table(model.Video{}.TableName())
	if err := query.Error; err != nil {
		log.Infof("error fetching data from db " + err.Error())
		return nil, 0, err
	}
	getSearchableQuery("title", queryParams.Get("title"), query)
	getSearchableQuery("description", queryParams.Get("desc"), query)
	query.Count(&totalCount)
	pageNumber := getPageNumber(queryParams)
	query = getPaginatedQuery(query, pageNumber, totalCount)
	if err := query.Find(&videos).Error; err != nil {
		return videos,0, err
	}
	return videos, pageNumber, nil
}

func getPaginatedQuery(query *g.DB, pageNumber int, totalCount int64) *g.DB {
	query = query.Order("published_at desc").
		Offset(calculateDbOffset(pageNumber, 20, int(totalCount))).Limit(20)
	return query
}

func getPageNumber(queryParams url.Values) int {
	pageNumber := queryParams.Get("page")
	if pageNumber == "" {
		pageNumber = "1"
	}
	number, _ := strconv.Atoi(pageNumber)
	return number
}

func getSearchableQuery(fieldName, fieldVal string, query *g.DB) {
	if fieldVal != "" {
		title := strings.Split(fieldVal, " ")
		searchString := "%"
		for _,str := range title {
			searchString = searchString + str + "%"
		}
		query.Where(fieldName + " like ?", searchString)
	}
}

func GetCronConfig() *model.YoutubeCronJobRun {
	var cronConfig []model.YoutubeCronJobRun
	clients.GetDB().Find(&cronConfig).Order("id desc").Limit(1)
	if len(cronConfig) == 0 {
		ytConfig := model.YoutubeCronJobRun{
			BaseModel: model.BaseModel{CreatedAt: time.Now().AddDate(0, 0, -1)},
			PublishedAfter: time.Now().AddDate(0, 0, -1).Format(time.RFC3339),
			APIKey:         "",
		}
		clients.GetDB().Save(&ytConfig)
		cronConfig = append(cronConfig, ytConfig)
	}
	return &cronConfig[0]
}

func GetAPIKeys() []model.APIConfig {
	var apiKeys []model.APIConfig
	clients.GetDB().Find(&apiKeys, "is_active = ?", true)
	if len(apiKeys) == 0 {
		defaultKey := config.GetConfig().DefaultAPIKey
		//todo refactor and move all this to config
		if os.Getenv("api") != "" {
			defaultKey = os.Getenv("api")
		}
		apiKey := model.APIConfig{
			APIKey:    defaultKey,
			IsActive:  true,
		}
		clients.GetDB().Save(&apiKey)
		apiKeys = append(apiKeys, apiKey)
		return apiKeys
	}
	return apiKeys
}

func GetVideoSearchConfig() *model.VideoSearchConfig {
	var searchConfig []model.VideoSearchConfig
	clients.GetDB().Find(&searchConfig, "is_active = ?", true)
	if len(searchConfig) == 0 {
		//get default values
		ytConfig := &model.VideoSearchConfig{
			BaseModel:  model.BaseModel{},
			Query:      "football",
			MaxResults: 20,
			OrderBy:    "date",
			SearchType:       "video",
			IsActive:   false,
		}
		clients.GetDB().Save(ytConfig)
		return ytConfig
	}
	return &searchConfig[0]
}

func calculateDbOffset(pageNumber, numberOfRecordsPerPage, totalCount int) int {
	if totalCount <= numberOfRecordsPerPage {
		return 0
	}
	maxPageNumber := int(math.Ceil(float64(totalCount) / float64(numberOfRecordsPerPage)))
	if pageNumber >= maxPageNumber {
		return numberOfRecordsPerPage * (maxPageNumber - 1)
	}
	return numberOfRecordsPerPage * (pageNumber - 1)
}

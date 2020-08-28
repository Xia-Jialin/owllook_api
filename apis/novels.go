package apis

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/howie6879/owllook_api/common"
	"github.com/howie6879/owllook_api/config"
)

var (
	NovelsRulesMap = config.NovelsRulesMap
)

// SearchAuthors returns all novels resource that you serached
func SearchAuthors(c *gin.Context) {
	novelName := c.Param("name")
	if novelName == "" {
		c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Parameter name can't be empty"})
		return
	}

	novelSource := c.Param("source")
	if novelSource == "" {
		c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Parameter name can't be empty"})
		return
	}
	
	var ok bool
	var currentRule config.NovelRule

	currentRule, ok = NovelsRulesMap[novelSource+"_1"]
	if ok == false {
		currentRule, ok = NovelsRulesMap[novelSource]
		if !ok {
			c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Parameter error"})
			return
		}
	}

	resultData, err := common.FetchHtml(novelName, currentRule)
	if err != nil {
		log.Println("Request URL error", err)
		c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Request error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "info": resultData})
}

// SearchNovels returns all novels resource that you serached
func SearchNovels(c *gin.Context) {
	novelName := c.Param("name")
	if novelName == "" {
		c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Parameter name can't be empty"})
		return
	}

	novelSource := c.Param("source")
	currentRule, ok := NovelsRulesMap[novelSource]
	if !ok {
		c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Parameter error"})
		return
	}

	resultData, err := common.FetchHtml(novelName, currentRule)
	if err != nil {
		log.Println("Request URL error", err)
		c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Request error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "info": resultData})
}

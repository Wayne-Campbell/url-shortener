package handler

import(
	"github.com/Wayne-Campbell/url-shortener/shortener"
	"github.com/Wayne-Campbell/url-shortener/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

//request model definition
type UrlCreationRequest struct{
	LongUrl string 'json:"long_url" binding:"required'
	UserId string 'json:"user_id" binding:"required'
}

func CreateShortUrl(c *gin.Context){
	var creationRequest UrlCreationRequest
	if err:= c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//generates shortened hash
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http:localhost:9808/"
	c.JSON(200, gin.H{
		"message": "short url created successfully"
		"short_url": host +shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context){
	shortUrl := c.Param("shortUrl")
	//retrieve initial url from the provided short url
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	//apply http redirection functions
	c.Redirect(302, initialUrl)
}
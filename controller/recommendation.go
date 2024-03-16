package controller

import (
	"Glossika_interview/database/models"
	"Glossika_interview/mycache"
	"Glossika_interview/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type RecommendationController struct{}

func (rc RecommendationController) GetRecommendation(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	// get from cache
	key := mycache.GetRecommendationKey(user.ID)
	cache := mycache.NewRedisCache(c)
	res, err := cache.Get(c, key)
	if err != nil && errors.Is(err, redis.Nil) {
		generateRecommendationCache(cache, c, user)
		return
	}

	// if found, response to user
	response.Success(c, res)
}

func generateRecommendationCache(cache *mycache.RedisCache, c *gin.Context, user models.User) {
	// add mutex lock
	mu := &sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()

	// if not found, get from db
	// check query flag
	sfKey := mycache.GetRecommendationQueryFlagKey(user.ID)
	_, sferr := cache.Get(c, sfKey)
	// if query flag not exists, get from db and set to cache
	if sferr != nil && errors.Is(sferr, redis.Nil) {
		// set query flag
		cache.Set(c, sfKey, true, 300)
		go func() {
			// query from db
			time.Sleep(time.Second * 300)

			// set to cache
			cache.Set(c, mycache.GetRecommendationKey(user.ID), "recommendation data", 600)
		}()
	}
	// if query flag exists, response to user for waiting
	response.Success(c, gin.H{"message": "recommendation is generating, please wait"})
	return
}

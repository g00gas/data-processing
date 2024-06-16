package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log-client-go/elasticClient"
	"net/http"
)

type AggregationResult struct {
	Key      string `json:"key"`
	DocCount int64  `json:"doc_count"`
}

func GetTopUserAgents(c *gin.Context) {
	client, err := elasticClient.GetElasticsearchClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aggs := elastic.NewTermsAggregation().Field("user_agent.name.keyword").Size(5)
	searchResult, err := client.Search().
		Index("apache_logs").
		Aggregation("top_user_agents", aggs).
		Size(0).
		Do(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aggRes, found := searchResult.Aggregations.Terms("top_user_agents")
	if !found {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Aggregation not found"})
		return
	}

	var results []AggregationResult
	for _, bucket := range aggRes.Buckets {
		results = append(results, AggregationResult{
			Key:      bucket.Key.(string),
			DocCount: bucket.DocCount,
		})
	}

	c.JSON(http.StatusOK, results)
}

func GetHTTPMethodsCount(c *gin.Context) {
	client, err := elasticClient.GetElasticsearchClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aggs := elastic.NewTermsAggregation().Field("http.request.method.keyword").Size(10)
	searchResult, err := client.Search().
		Index("apache_logs").
		Aggregation("http_methods_count", aggs).
		Size(0).
		Do(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aggRes, found := searchResult.Aggregations.Terms("http_methods_count")
	if !found {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Aggregation not found"})
		return
	}

	var results []AggregationResult
	for _, bucket := range aggRes.Buckets {
		results = append(results, AggregationResult{
			Key:      bucket.Key.(string),
			DocCount: bucket.DocCount,
		})
	}

	c.JSON(http.StatusOK, results)
}

func GetStatusCodeDistribution(c *gin.Context) {
	client, err := elasticClient.GetElasticsearchClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aggs := elastic.NewTermsAggregation().Field("http.response.status_code").Size(10)
	searchResult, err := client.Search().
		Index("apache_logs").
		Aggregation("status_code_distribution", aggs).
		Size(0).
		Do(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aggRes, found := searchResult.Aggregations.Terms("status_code_distribution")
	if !found {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Aggregation not found"})
		return
	}

	var results []AggregationResult
	for _, bucket := range aggRes.Buckets {
		results = append(results, AggregationResult{
			Key:      fmt.Sprintf("%v", bucket.Key),
			DocCount: bucket.DocCount,
		})
	}

	c.JSON(http.StatusOK, results)
}

func GetDeviceNames(c *gin.Context) {
	client, err := elasticClient.GetElasticsearchClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aggs := elastic.NewTermsAggregation().Field("user_agent.device.name.keyword").Size(10)
	searchResult, err := client.Search().
		Index("apache_logs").
		Aggregation("device_names", aggs).
		Size(0).
		Do(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aggRes, found := searchResult.Aggregations.Terms("device_names")
	if !found {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Aggregation not found"})
		return
	}

	var results []AggregationResult
	for _, bucket := range aggRes.Buckets {
		results = append(results, AggregationResult{
			Key:      bucket.Key.(string),
			DocCount: bucket.DocCount,
		})
	}

	c.JSON(http.StatusOK, results)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"sort"
	"tdidf/internal/models"
)

func (h *Handler) TfIdf(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	files := form.File["txtFile"]
	var txt string
	for i := range files {
		f, err := files[i].Open()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		txt = string(data)
	}

	tf := h.Service.Tf(txt)
	idf := h.Service.Idf(txt)

	var results []models.Response
	for word, tfValue := range tf {
		if idfValue, ok := idf[word]; ok {
			results = append(results, models.Response{
				Key:   word,
				Tf:    tfValue,
				Idf:   idfValue,
				TfIdf: tfValue * idfValue,
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Idf > results[j].Idf
	})

	if len(results) > 50 {
		results = results[:50]
	}

	c.JSON(200, gin.H{
		"results": results,
	})
}

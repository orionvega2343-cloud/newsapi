package internal

import "github.com/gin-gonic/gin"

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) Register(c *gin.Context) {
	var req req
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	err = h.s.Register(req.Email, req.Password)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(201, gin.H{"message": "ok"})
}

func (h *Handler) Login(c *gin.Context) {
	var req req
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	res, err := h.s.Login(req.Email, req.Password)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, gin.H{"token": res})

}

func (h *Handler) GetArticle(c *gin.Context) {
	res, err := h.s.GetArticles()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, gin.H{"article": res})

}

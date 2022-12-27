package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
)
type Cat struct {
	Name string `json:"name"`
	TailLength int `json:"tail_length"`
	IsStripe bool `json:"is_stripe"`
}
func main() {
	var cat Cat
	cat.Name="Мурзик"
	cat.TailLength=31
	cat.IsStripe=true
	fmt.Print(cat)
	r:=gin.Default()
	r.GET("/cat", func(c *gin.Context) {
		c.JSON(200, cat)
	})
	r.Run("0.0.0.0:8888")
}
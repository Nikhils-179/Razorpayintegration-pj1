package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	razorpay "github.com/razorpay/razorpay-go"
)

type PageVariable struct {
	OrderId string
	Email   string
	Name    string
	Amount  string
	Contact string
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("*.html")
	router.GET("/", App)
	router.GET("/payment-success", PaymentSuccess)
	router.Run(":8089")
}

func App(c *gin.Context) {

	page := &PageVariable{}
	page.Amount = "100"
	page.Email = " ns@gmail.com"
	page.Name = "NS"
	page.Contact = "9876554431"
	client := razorpay.NewClient("rzp_test_c7ykIqTKLN8wRO", "MvPUXJgnb2PJhDxfDF0LHa7h")

	data := map[string]interface{}{
		"amount":   page.Amount,
		"currency": "INR",
		"receipt":  "Xy123end",
	}

	body, err := client.Order.Create(data, nil)
	fmt.Println("///////receipt", body)

	if err != nil {
		fmt.Println("Problem getting the respository information", err)
		os.Exit(1)
	}

	value := body["id"]

	str := value.(string)
	fmt.Println("str////////", str)

	HomePageVars := PageVariable{
		OrderId: str,
		Amount:  page.Amount,
		Email:   page.Email,
		Name:    page.Name,
		Contact: page.Contact,
	}

	c.HTML(http.StatusOK, "app.html", HomePageVars)

}

func PaymentSuccess(c *gin.Context) {
	paymentid := c.Query("paymentid")
	orderid := c.Query("orderid")
	signature := c.Query("signature")

	fmt.Println(paymentid, "paymentid")
	fmt.Println(orderid, "orderid")
	fmt.Println(signature, "signature")
}

func PaymentFailure(c *gin.Context) {}

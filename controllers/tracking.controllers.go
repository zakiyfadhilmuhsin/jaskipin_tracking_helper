package controllers

import (
	"encoding/json"
	"io/ioutil"
	"jaskipin-front/db"
	"jaskipin-front/models"
	"net/http"
	"strings"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// type DataManifest struct {
// 	Data interface{} `json:"data"`
// }

func GetTracking(c echo.Context) error {
	id_order := c.Param("id_order")

	var res models.Response
	var tracking models.Tracking

	// Read
	// db.Database.First(&tracking, id_order) // find product with id 1
	db.Database.Where("id_order = ?", id_order).Find(&tracking) // find product with code l1212
	//db.Database.First(&page, "title = ?", "L1212") // find product with code l1212

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = tracking

	return c.JSON(http.StatusOK, res)
}

// func GetManifest(c echo.Context) error {
// 	carrier := c.Param("carrier")
// 	trackingNumber := c.Param("trackingNumber")

// 	//var res models.Response

// 	req.Debug = true

// 	header := make(http.Header)
// 	header.Set("Content-Type", "application/json")
// 	header.Set("Trackingmore-Api-Key", "0adc6568-bff3-429c-a3cb-cac4a96f6450")
// 	r, err := req.Get("https://api.trackingmore.com/v2/trackings/"+carrier+"/"+trackingNumber, header)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resp := r.Response()
// 	defer resp.Body.Close()

// 	// bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// bodyString := string(bodyBytes)

// 	var result interface{}

// 	err = json.NewDecoder(resp.Body).Decode(result)
// 	if err == io.EOF {
// 		return nil
// 	}
// 	return err

// 	// res.Status = http.StatusOK
// 	// res.Message = "success"
// 	// res.Data = data

// 	return c.JSON(http.StatusOK, result)
// 	//return json.NewEncoder(resp.Body).Encode(data)
// }

func GetManifest(c echo.Context) error {
	//3 create a tracking number
	//var url string ="http://api.trackingmore.com/v2/trackings/post"
	//var postData string ="{\"tracking_number\": \"BYS006086078\",\"carrier_code\":\"yanwen\",\"title\":\"chase chen\",\"customer_name\":\"chase\",\"customer_email\":\"abc@qq.com\",\"order_id\":\"#123\",\"order_create_time\":\"2018-05-20 12:00\",\"destination_code\":\"IL\",\"tracking_ship_date\":\"1521314361\",\"tracking_postal_code\":\"13ES20\",\"lang\":\"en\",\"logistics_channel\":\"4PX page\"}"
	//httpDo(url,postData,"POST")

	var res models.Response

	carrier := c.Param("carrier")
	trackingNumber := c.Param("trackingNumber")

	//5 Get tracking results of a single tracking
	var url string = "http://api.trackingmore.com/v2/trackings/" + carrier + "/" + trackingNumber
	var postData string = ""
	var method string = "GET"
	// data := httpDo(url, postData, "GET")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(postData))
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "application/json'")
	req.Header.Set("Trackingmore-Api-Key", "d9636e04-1068-44db-8a55-cfa12742d89e")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	// Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(body), &result)

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = result

	return c.JSON(http.StatusOK, res)
}

func CreateManifest(c echo.Context) error {
	//3 create a tracking number
	//var url string ="http://api.trackingmore.com/v2/trackings/post"
	//var postData string ="{\"tracking_number\": \"BYS006086078\",\"carrier_code\":\"yanwen\",\"title\":\"chase chen\",\"customer_name\":\"chase\",\"customer_email\":\"abc@qq.com\",\"order_id\":\"#123\",\"order_create_time\":\"2018-05-20 12:00\",\"destination_code\":\"IL\",\"tracking_ship_date\":\"1521314361\",\"tracking_postal_code\":\"13ES20\",\"lang\":\"en\",\"logistics_channel\":\"4PX page\"}"
	//httpDo(url,postData,"POST")

	var res models.Response

	carrier_code := c.FormValue("carrier_code")
	tracking_number := c.FormValue("tracking_number")

	//5 Get tracking results of a single tracking
	var url string = "http://api.trackingmore.com/v2/trackings/post"
	var postData string = "{\"tracking_number\":\"" + tracking_number + "\",\"carrier_code\":\"" + carrier_code + "\"}"
	var method string = "POST"
	// data := httpDo(url, postData, "GET")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(postData))
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "application/json'")
	req.Header.Set("Trackingmore-Api-Key", "0adc6568-bff3-429c-a3cb-cac4a96f6450")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	// Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(body), &result)

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = result

	return c.JSON(http.StatusOK, res)
}

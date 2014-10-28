package pentabot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	updatingStage bool = false
)

func VerifyGoal() {
	var m map[string]string
	url := "https://api.spark.io/v1/devices/53ff72065067544846101187/result?access_token=b8df700c30b45c4e46679cc91a702eb7d21842c8"
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Get(url)
	if err != nil {
		fmt.Println("Error getting url")
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("Error reading body")
		}

		json.Unmarshal(contents, &m)

		if response.Status != "200 OK" {
			fmt.Println("Error: ", m["error"])
			return
		}

		distance, err := strconv.ParseFloat(m["result"], 64)

		if err != nil {
			fmt.Println("Error parsing result")
			return
		}

		if distance < 10 && !updatingStage {
			updatingStage = true
			NextStage()
			time.AfterFunc(5*time.Second, func() {
				updatingStage = false
			})
		}
	}
}

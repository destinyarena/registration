package registration

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *handler) badIP(ip string) (bool, error) {
	url := fmt.Sprintf("http://check.getipintel.net/check.php?ip=%s&contact=ar2roguerra@gmail.com", ip)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, fmt.Errorf("Error Code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	rawdata := string(body)
	percent, err := strconv.ParseFloat(rawdata, 64)
	if err != nil {
		return false, err
	}

	h.Logger.Infof("IP: %s Percent: %v", ip, percent)

	if percent >= 1.0 {
		return true, nil
	}

	return false, nil
}

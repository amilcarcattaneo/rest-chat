package check

import (
	checkClient "rest-chat/src/api/clients/check"
)

func Check() (bool, error) {
	return checkClient.Check()
}

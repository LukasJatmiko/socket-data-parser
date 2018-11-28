package responseparser

import (
	"regexp"
)

func Parse(databuf []byte, datalength int) []string {
	rgx := regexp.MustCompile(`(?i)}{`)
	messages := rgx.Split(string(databuf[:datalength]), -1)

	for msgIdx, _ := range messages {
		//Restore splitted message
		if len(messages) != 1 {
			if msgIdx == 1 {
				messages[msgIdx] = messages[msgIdx] + "}"
			} else {
				if msgIdx == (len(messages) - 1) {
					messages[msgIdx] = "{" + messages[msgIdx]
				} else {
					messages[msgIdx] = "{" + messages[msgIdx] + "}"
				}
			}
		}
	}
	return messages
}

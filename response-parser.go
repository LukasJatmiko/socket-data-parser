package responseparser

import (
	"regexp"
)

//Parse : Parse data into JSON
func Parse(databuf []byte, datalength int) []string {
	rgx := regexp.MustCompile(`(?i)}[\n\s\t\r]*{`)
	messages := rgx.Split(string(databuf[:datalength]), -1)

	for msgIdx := range messages {
		//Restore splitted message
		if len(messages) != 1 {
			if msgIdx == 0 {
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

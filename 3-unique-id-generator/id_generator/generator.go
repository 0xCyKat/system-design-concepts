package id_generator

import (
	"errors"
	"time"
)

var (
	lastTimeStamp int64
	sequenceID    int64
)

const (
	customEpoch  = int64(1767225600000)
	machineID    = int64(1) << 12
	dataCenterID = int64(1) << 17
)

/*
Implementation of Twitter Snowflake ID generation Algorithm

It's a 64 bit numerical ID, divided into 5 sections

1. Sign Bit (1)
2. Timestamp (41)
3. Datacenter ID (5)
4. Machine ID (5)
5. Sequence Number (12)

It means a snowflake cluster can generate

2^5 * 2^5 * 2^12 = 4,194,304

~4 million unique IDs per millisecond
~4 billon unique IDs per second
*/

func NextID() (int64, error) {
	currentTime := time.Now().UnixMilli()

	if currentTime < lastTimeStamp {
		return 0, errors.New("clock moved backwards")
	}

	if currentTime == lastTimeStamp {
		sequenceID = (sequenceID + 1) & 4095
		if sequenceID == 0 {
			for currentTime <= lastTimeStamp {
				currentTime = time.Now().UnixMilli()
			}
		}
	} else {
		sequenceID = 0
	}

	lastTimeStamp = currentTime

	id := ((currentTime - customEpoch) << 22) | dataCenterID | machineID | sequenceID

	return id, nil
}

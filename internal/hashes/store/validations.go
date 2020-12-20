package store

import (
	"log"
	"strconv"
)

func (s *Store) ValidateMsisdnLen(msisdn string) bool {
	l := len(msisdn)
	return 12 <= l && l <= 21
}

var ccs = map[string]int{"380": 380}

func (s *Store) ValidateCC(msisdn string) (string, bool) {
	cc := msisdn[:3]
	_, ok := ccs[cc]
	return cc, ok
}

func (s *Store) ValidateNDC(msisdn string) (string, bool) {
	ndcStr := msisdn[3:5]

	ndc, err := strconv.Atoi(ndcStr)
	if err != nil {
		log.Println(err)
	}

	for _, n := range s.ndcs {
		if ndc == n {
			return ndcStr, true
		}
	}

	return ndcStr, false
}

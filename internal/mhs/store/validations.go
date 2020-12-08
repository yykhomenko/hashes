package store

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
	ndc := msisdn[3:5]
	_, ok := ndcs[ndc]
	return ndc, ok
}

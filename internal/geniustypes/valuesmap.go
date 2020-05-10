package geniustypes

import (
	"fmt"
	"strings"
)

type ValuesMap map[string]string

func (l ValuesMap) String() string {
	res := make([]string, 0, len(l))
	for k, v := range l {
		res = append(res, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(res, ", ")
}

func (l ValuesMap) Set(v string) error {
	nv := strings.Split(v, "=")
	if len(nv) != 2 {
		return fmt.Errorf("expected NAME=VALUE, got %s", v)
	}
	l[nv[0]] = nv[1]
	return nil
}

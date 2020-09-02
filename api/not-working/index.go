package handler

import (
    "net/http"
    "strconv"
    "github.com/pkg/math"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(strconv.Itoa(math.Max(1, 2)))
}

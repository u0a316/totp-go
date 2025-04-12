
package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pquerna/otp/totp"
)

const secret = "JBSWY3DPEHPK3PXP" // Simpan ini aman, bisa baca dari file/env juga

func main() {
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		ok := totp.Validate(code, secret)

		if ok {
			json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		} else {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		}
	})

	http.ListenAndServe(":0.0.0.0:8080", nil) // hanya bisa diakses lokal
}

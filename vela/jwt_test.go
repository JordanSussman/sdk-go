// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	TestTokenGood    = makeSampleToken(jwt.MapClaims{"exp": float64(time.Now().Unix() + 100)})
	TestTokenExpired = makeSampleToken(jwt.MapClaims{"exp": float64(time.Now().Unix() - 100)})
)

func TestIsTokenExpired(t *testing.T) {
	// run tests
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "expired token",
			args: args{
				token: TestTokenExpired,
			},
			want: true,
		},
		{
			name: "good token",
			args: args{
				token: TestTokenGood,
			},
			want: false,
		},
		{
			name: "empty token",
			args: args{
				token: "",
			},
			want: true,
		},
		{
			name: "bad token",
			args: args{
				token: "/65",
			},
			want: true,
		},
		{
			name: "no exp",
			args: args{
				token: makeSampleToken(jwt.MapClaims{}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTokenExpired(tt.args.token); got != tt.want {
				t.Errorf("IsTokenExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

// makeSampleToken is a helper to create test tokens
// with the given claims.
func makeSampleToken(c jwt.Claims) string {
	// create a new token
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)

	// get the signing string (header + claims)
	s, e := t.SigningString()

	if e != nil {
		return ""
	}

	// add bogus signature
	s = fmt.Sprintf("%s.abcdef", s)

	return s
}

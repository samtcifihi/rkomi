package main

import (
	_ "flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var ir1 int
	var ir2 int
	err := false

	// parse arguments and flags

	// convert strengths to internal points rating
	if len(os.Args) >= 3 {
		ir1, _ = xrtoir(os.Args[1], "ogs", 14) // assumes stoneWidth=14
		ir2, _ = xrtoir(os.Args[2], "ogs", 14) // assumes stoneWidth=14
	} else {
		fmt.Println("Insufficient Arguments")
		err = true
	}

	// calculate reverse komi and/or handicap stones
	rkomi := ir2 - ir1  // assumes order=white; random=false; rkomi=-1,-1
	komi := rkomi + 7.0 // assumes baseKomi=7.0

	// output
	// assumes many many flags
	if !err && len(os.Args) >= 3 {
		fmt.Printf("White [%v] plays Black [%v] with no handicap stones and a komi of %v\n", os.Args[1], os.Args[2], komi)
	}
}

// convertRating converts an external rating xr to an internal rating ir.
// xro (external rating origin) defines the behavior of the conversion.
// invalid inputs result in an error flag being returned.
func xrtoir(xr string, xro string, stoneWidth int) (int, bool) {
	ir := 0
	err := false

	switch xro {
	case "ir":
		r, e := strconv.Atoi(xr)
		if e != nil {
			err = true
		}
		ir = r
	case "kd":
		// [0, stoneWidth) == 1d
		// validate for 1 or more digits, followed by one of {k, d}
		// \d+[kd]
		valid, e := regexp.Match(`\d+[kd]`, []byte(xr))
		if (e != nil) || !valid {
			err = true
		} else {
			if xr[len(xr)-1] == byte('k') {
				// conversion for kyus
				r, e := strconv.Atoi(xr[:len(xr)-1])
				if e != nil {
					err = true
				}
				// place in middle of range, rounded towards 0
				ir = (r * stoneWidth * -1) + int(float64(stoneWidth)*0.5)
			} else {
				// conversion for dans
				r, e := strconv.Atoi(xr[:len(xr)-1])
				if e != nil {
					err = true
				}
				// place in middle of range, rounded towards 0
				ir = (r * stoneWidth) - int(float64(stoneWidth)*0.5)
			}
		}
	case "ogs":
		// run as kyu-dan if applicable
		validkd, e := regexp.Match(`\d+[kd]`, []byte(xr))
		if e != nil {
			err = true
		} else if validkd {
			ir, err = xrtoir(xr, "kd", stoneWidth)
		} else {
			// [0, stoneWidth) == 1d
			r, e := strconv.ParseFloat(xr, 64)
			if e != nil {
				err = true
			}
			ir = int(math.Round(((math.Log(r/525.0) * 23.15) - 30.0) * float64(stoneWidth)))
		}
	default:
		err = true
	}

	return ir, err
}

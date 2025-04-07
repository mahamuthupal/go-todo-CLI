package todo

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "os"

)

// lowercase struct bc we donot export it
type item struct {
	Task        string
	Done        bool
}

// list of to-do items -- upercase casue we want it to export
type List []item
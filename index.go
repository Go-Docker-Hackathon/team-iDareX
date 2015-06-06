package main

import (
	"net/http" 
	"html/template" 
	"fmt"
	"log"
	"strings"
	"os" 
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)


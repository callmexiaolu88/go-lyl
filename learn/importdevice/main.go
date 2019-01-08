package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"honeywell.com/foxconn/fire-platform-common-pkg/mongodb"

	"honeywell.com/foxconn/fire-protocol-stack-side/mock/model"
)

var firedir, waterdir, hostname string

func init() {
	flag.StringVar(&firedir, "firedir", "firedevice", "fire device dir")
	flag.StringVar(&waterdir, "waterdir", "waterdevice", "water device dir")
	flag.StringVar(&hostname, "hostname", "127.0.0.1", "host name")

}

func main() {
	flag.Parse()
	fs, err := ioutil.ReadDir(firedir)
	if err != nil {
		fmt.Println(err)
		return
	}
	dbname := "firemock"
	colName := "firedevice"
	firedevices := []model.FileToDBDevice{}
	for _, f := range fs {
		if !f.IsDir() {
			filename := filepath.Join(firedir, f.Name())
			bs, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Println(err)
				continue
			}
			bs = bytes.TrimPrefix(bs, []byte("\xef\xbb\xbf"))
			ts := []model.FileToDBDevice{}
			err = json.Unmarshal(bs, &ts)
			if err != nil {
				fmt.Println(err)
				continue
			}
			firedevices = append(firedevices, ts...)
		}
	}
	if len(firedevices) > 0 {
		fmt.Println(firedevices[len(firedevices)-1])
		err = mongodb.MongodbDao.InitSession(hostname, dbname)
		if err != nil {
			fmt.Println(err)
			return
		}
		adds := make([]interface{}, 0)
		for _, v := range firedevices {
			query := bson.M{
				"id": v.ID,
			}
			n, err := mongodb.MongodbDao.QueryCount(colName, query)
			if err == mgo.ErrNotFound || n == 0 {
				adds = append(adds, v)
				continue
			} else if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("device repeated.info:", v)
		}
		if len(adds) > 0 {
			err = mongodb.MongodbDao.Insert("firedevice", adds...)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

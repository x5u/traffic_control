// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Parameter struct {
	Id          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	ConfigFile  string         `db:"config_file" json:"configFile"`
	Value       string         `db:"value" json:"value"`
	LastUpdated time.Time      `db:"last_updated" json:"lastUpdated"`
	Links       ParameterLinks `json:"_links" db:-`
}

type ParameterLinks struct {
	Self string `db:"self" json:"_self"`
}

type ParameterLink struct {
	ID  int64  `db:"parameter" json:"id"`
	Ref string `db:"parameter_id_ref" json:"_ref"`
}

// @Title getParameterById
// @Description retrieves the parameter information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Parameter
// @Resource /api/2.0
// @Router /api/2.0/parameter/{id} [get]
func getParameterById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []Parameter{}
	arg := Parameter{}
	arg.Id = int64(id)
	queryStr := "select *, concat('" + API_PATH + "parameter/', id) as self "
	queryStr += " from parameter where id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getParameters
// @Description retrieves the parameter
// @Accept  application/json
// @Success 200 {array}    Parameter
// @Resource /api/2.0
// @Router /api/2.0/parameter [get]
func getParameters(db *sqlx.DB) (interface{}, error) {
	ret := []Parameter{}
	queryStr := "select *, concat('" + API_PATH + "parameter/', id) as self "
	queryStr += " from parameter"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postParameter
// @Description enter a new parameter
// @Accept  application/json
// @Param                 Body body     Parameter   true "Parameter object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/parameter [post]
func postParameter(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Parameter
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO parameter("
	sqlString += "name"
	sqlString += ",config_file"
	sqlString += ",value"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:config_file"
	sqlString += ",:value"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putParameter
// @Description modify an existing parameterentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Parameter   true "Parameter object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/parameter/{id}  [put]
func putParameter(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Parameter
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE parameter SET "
	sqlString += "name = :name"
	sqlString += ",config_file = :config_file"
	sqlString += ",value = :value"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delParameterById
// @Description deletes parameter information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Parameter
// @Resource /api/2.0
// @Router /api/2.0/parameter/{id} [delete]
func delParameter(id int, db *sqlx.DB) (interface{}, error) {
	arg := Parameter{}
	arg.Id = int64(id)
	result, err := db.NamedExec("DELETE FROM parameter WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
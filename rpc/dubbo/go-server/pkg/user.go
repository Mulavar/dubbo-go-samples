/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pkg

import (
	"fmt"
	"strconv"
	"time"
)

import (
	hessian "github.com/apache/dubbo-go-hessian2"
)

type Gender hessian.JavaEnum

const (
	MAN hessian.JavaEnum = iota
	WOMAN
)

var genderName = map[hessian.JavaEnum]string{
	MAN:   "MAN",
	WOMAN: "WOMAN",
}

var genderValue = map[string]hessian.JavaEnum{
	"MAN":   MAN,
	"WOMAN": WOMAN,
}

func (g Gender) JavaClassName() string {
	return "org.apache.dubbo.Gender"
}

func (g Gender) String() string {
	s, ok := genderName[hessian.JavaEnum(g)]
	if ok {
		return s
	}

	return strconv.Itoa(int(g))
}

func (g Gender) EnumValue(s string) hessian.JavaEnum {
	v, ok := genderValue[s]
	if ok {
		return v
	}

	return hessian.InvalidJavaEnum
}

type (
	User struct {
		// !!! Cannot define lowercase names of variable
		ID   string
		Name string
		Age  int32
		Time time.Time
		Sex  Gender // notice: java enum Object <--> go string
	}
)

var (
	DefaultUser = User{
		ID: "000", Name: "Alex Stocks", Age: 31,
		Sex: Gender(MAN),
	}

	userMap = make(map[string]User)
)

func init() {
	userMap["000"] = DefaultUser
	userMap["001"] = User{ID: "001", Name: "ZhangSheng", Age: 18, Sex: Gender(MAN)}
	userMap["002"] = User{ID: "002", Name: "Lily", Age: 20, Sex: Gender(WOMAN)}
	userMap["003"] = User{ID: "113", Name: "Moorse", Age: 30, Sex: Gender(WOMAN)}
	for k, v := range userMap {
		v.Time = time.Now()
		userMap[k] = v
	}
}

func (u User) String() string {
	return fmt.Sprintf(
		"User{ID:%s, Name:%s, Age:%d, Time:%s, Sex:%s}",
		u.ID, u.Name, u.Age, u.Time, u.Sex,
	)
}

func (u User) JavaClassName() string {
	return "org.apache.dubbo.User"
}

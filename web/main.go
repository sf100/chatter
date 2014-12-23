// Copyright © 2014 Terry Mao, LiuDing All rights reserved.
// This file is part of gopush-cluster.

// gopush-cluster is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// gopush-cluster is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with gopush-cluster.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	log "code.google.com/p/log4go"
	"flag"
	"github.com/sf100/chatter/app"
	"github.com/sf100/chatter/db"
	"github.com/sf100/chatter/perf"
	"github.com/sf100/chatter/process"
	"github.com/sf100/chatter/ver"
	"runtime"
)

func main() {
	var err error
	// Parse cmd-line arguments
	flag.Parse()
	log.Info("web ver: \"%s\" start", ver.Version)
	if err = InitConfig(); err != nil {
		panic(err)
	}

	/*初始化我的配置信息*/
	//init appconfig
	if err = app.InitConfig(); err != nil {
		log.Error("InitConfig() error(%v)", err)
		return
	}
	//init db config
	if err = db.InitConfig(); err != nil {
		log.Error("db-InitConfig() error(%v)", err)
		return
	}

	// Set max routine
	runtime.GOMAXPROCS(Conf.MaxProc)
	// init log
	log.LoadConfiguration(Conf.Log)
	defer log.Close()
	// init zookeeper
	zkConn, err := InitZK()
	if err != nil {
		if zkConn != nil {
			zkConn.Close()
		}
		panic(err)
	}

	/*初始化数据库，和app的redis*/
	db.InitDB()
	defer db.CloseDB()
	app.InitRedisStorage()

	// start pprof http
	perf.Init(Conf.PprofBind)
	// start http listen.
	StartHTTP()
	// process init
	if err = process.Init(Conf.User, Conf.Dir, Conf.PidFile); err != nil {
		panic(err)
	}
	// init signals, block wait signals
	signalCH := InitSignal()
	HandleSignal(signalCH)
	log.Info("web stop")
}

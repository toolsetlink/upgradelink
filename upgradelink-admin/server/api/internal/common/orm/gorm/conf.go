// Copyright 2023 The Ryan SU Authors (https://github.com/suyuan32). All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gorm

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Conf is the configuration structure for GORM.
type Conf struct {
	Type        string `json:",default=mysql,options=[mysql,postgres]"` // type of database: mysql, postgres
	Host        string `json:",default=localhost"`                      // address
	Port        int    `json:",default=3306"`                           // port
	Config      string `json:",optional"`                               // extra config such as charset=utf8mb4&parseTime=True
	DBName      string `json:",default=simple_admin"`                   // database name
	Username    string `json:",default=root"`                           // username
	Password    string `json:",optional"`                               // password
	MaxIdleConn int    `json:",default=10"`                             // the maximum number of connections in the idle connection pool
	MaxOpenConn int    `json:",default=100"`                            // the maximum number of open connections to the database
	LogMode     string `json:",default=error"`                          // open gorm's global logger
}

// MysqlDSN returns the MySQL DSN link from the configuration.
func (g Conf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", g.Username, g.Password, g.Host, g.Port, g.DBName, g.Config)
}

// PostgreSqlDSN returns the PostgreSQL DSN link from the configuration.
func (g Conf) PostgreSqlDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s", g.Host, g.Username, g.Password,
		g.DBName, g.Port, g.Config)
}

func (g Conf) NewGORM() (*gorm.DB, error) {
	switch g.Type {
	case "mysql":
		return MysqlClient(g)
	case "postgres":
		return PgSqlClient(g)
	default:
		return MysqlClient(g)
	}
}

func MysqlClient(c Conf) (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		DSN:                       c.MysqlDSN(),
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // autoconfiguration based on currently MySQL version
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.New(writer{}, logger.Config{
			SlowThreshold:             1 * time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  getLevel(c.LogMode),
		}),
	}); err != nil {
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(c.MaxIdleConn)
		sqlDB.SetMaxOpenConns(c.MaxOpenConn)
		return db, nil
	}
}

func PgSqlClient(c Conf) (*gorm.DB, error) {
	pgsqlConfig := postgres.Config{
		DSN:                  c.PostgreSqlDSN(),
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}

	if db, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{
		Logger: logger.New(writer{}, logger.Config{
			SlowThreshold:             1 * time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  getLevel(c.LogMode),
		}),
	}); err != nil {
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(c.MaxIdleConn)
		sqlDB.SetMaxOpenConns(c.MaxOpenConn)
		return db, nil
	}
}

// getLevel returns the gorm level from the level in go zero.
func getLevel(logMode string) logger.LogLevel {
	var level logger.LogLevel
	switch logMode {
	case "info":
		level = logger.Info
	case "warn":
		level = logger.Warn
	case "error":
		level = logger.Error
	case "silent":
		level = logger.Silent
	default:
		level = logger.Error
	}
	return level
}

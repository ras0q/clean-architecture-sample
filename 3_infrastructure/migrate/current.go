package migrate

import "github.com/ras0q/clean-architecture-sample/2_interface/repository/model"

func AllTables() []interface{} {
	return []interface{}{
		// Note: テーブルを新たに作成した際はここに追記する
		&model.User{},
	}
}

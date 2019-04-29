package di

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	domain_task "taskforce/domain/task"
	"taskforce/task"
	"taskforce/task/delivery"
	"taskforce/task/gateway"
	"taskforce/task/repository"
	"taskforce/task/usecase"
)

func InjectStateFactory() domain_task.StateFactory {
	return domain_task.NewStateFactory()
}

func InjectTaskFactory() domain_task.Factory {
	return domain_task.NewTaskFactory(InjectStateFactory())
}

func InjectSqliteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// テーブル作成
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "TASKS" ("ID" INTEGER PRIMARY KEY, "TITLE" VARCHAR(255), "DESCRIPTION" VARCHAR(255), "STATE" VARCHAR(255))`,
	)
	if err != nil {
		panic(err)
	}
	return db
}

func InjectRepository() task.Repository {
	return repository.NewRepository(InjectSqliteDB())
}

func InjectGateway() task.Gateway {
	return gateway.NewSqlTaskGateway(InjectRepository(), InjectTaskFactory())
}

func InjectUsecase() task.Usecase {
	return usecase.NewTaskUsecase(InjectGateway(), InjectTaskFactory())
}

func InjectCommandLineDelivery() *delivery.CommandLineDelivery {
	return delivery.NewCommandLineDelivery(InjectUsecase(), InjectViewUsecase())
}

func InjectViewGateway() task.ViewGateway {
	return gateway.NewViewTaskGateway(InjectRepository(), InjectTaskFactory())
}

func InjectViewUsecase() task.ViewTaskUsecase {
	return usecase.NewViewTaskUsecase(InjectViewGateway())
}

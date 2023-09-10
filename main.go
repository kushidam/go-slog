package main

import (
	"errors"
	"log"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	// 各レベル出力メソッド
	slog.Info("hello", "count", 1 )
	slog.Debug("hello", "count", 2 )
	slog.Warn("hello", "count", 3)

	err := errors.New("this is error")
	slog.Error("message", err, "count", 4)

	// 構造化レコードをテキスト形式ロガー作成
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Info("hello", "count", 5)

	// logパッケージへのdefault設定
	log.Println("log package SetDefault before")
	slog.SetDefault(logger)
	log.Println("log package SetDefault after")

	// JSON形式出力
	logger_json := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger_json.Info("hello json", "count", 6)

	// レベル設定：ログ イベントの重要性または重大度を表す整数
	var programLevel = new(slog.LevelVar) // デフォルトの情報
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)

	//グループ
	logger.Info(
		"nest",
		slog.Int("count", 7),
		slog.Group("Group1",
			slog.Int("key1", 1000),
			slog.Int("key2", 2000),
		),
	)
	type Users struct {
		name string
		age int
	}
	user1 := Users{
		name : "hogehoge",
		age : 20,
	}
	logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Info("users","data", user1)
	
}
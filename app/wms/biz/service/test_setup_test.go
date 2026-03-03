package service

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/MoScenix/wms/app/wms/biz/dal"
	"github.com/joho/godotenv"
)

func setDefaultEnv(key, value string) {
	if os.Getenv(key) == "" {
		_ = os.Setenv(key, value)
	}
}

func TestMain(m *testing.M) {
	_, thisFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(thisFile)
	target := filepath.Clean(filepath.Join(baseDir, "../../"))
	_ = os.Chdir(target)
	_ = godotenv.Load(filepath.Join(target, ".env"))

	setDefaultEnv("GO_ENV", "test")
	setDefaultEnv("MYSQL_USER", "gorm")
	setDefaultEnv("MYSQL_PASSWORD", "gorm")
	setDefaultEnv("MYSQL_DATABASE", "gorm")
	setDefaultEnv("MYSQL_HOST", "127.0.0.1")

	dal.Init()
	os.Exit(m.Run())
}

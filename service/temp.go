package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func flashtime(a *App) {
	go func() {
		for {
			now := time.Now().Format(time.DateTime)
			runtime.EventsEmit(a.ctx, "time", now)
			time.Sleep(time.Second)
		}
	}()
}

func getJSONString() string {
	datas := map[string]interface{}{
		"code":    200,
		"message": "请求成功",
		"data": map[string]interface{}{
			"dir": "/",
			"list": []map[string]interface{}{
				{
					"category":        6,
					"extent_int2":     0,
					"extent_int8":     0,
					"extent_tinyint7": 0,
					"from_type":       0,
					"fs_id":           1050555082167035,
					"is_scene":        0,
					"isdir":           1,
					"local_ctime":     1582705583,
					"local_mtime":     1582705583,
					"oper_id":         0,
					"owner_id":        0,
					"owner_type":      0,
					"path":            "/apps",
					"pl":              0,
					"real_category":   "",
					"server_atime":    0,
					"server_ctime":    1582705583,
					"server_filename": "apps",
					"server_mtime":    1665826588,
					"share":           0,
					"size":            0,
					"tkbind_id":       0,
					"unlist":          0,
					"wpfile":          0,
				},
				{
					"category":        6,
					"extent_int2":     0,
					"extent_int8":     0,
					"extent_tinyint7": 0,
					"from_type":       0,
					"fs_id":           1045474540778144,
					"is_scene":        0,
					"isdir":           1,
					"local_ctime":     1724234634,
					"local_mtime":     1724234634,
					"oper_id":         4503599727530676,
					"owner_id":        0,
					"owner_type":      0,
					"path":            "/访问须知",
					"pl":              2,
					"real_category":   "",
					"server_atime":    0,
					"server_ctime":    1724234634,
					"server_filename": "访问须知",
					"server_mtime":    1724234664,
					"share":           0,
					"size":            0,
					"tkbind_id":       0,
					"unlist":          0,
					"wpfile":          0,
				},
			},
		},
	}

	jsonBytes, err := json.Marshal(datas)
	if err != nil {
		return "1"
	}

	return string(jsonBytes)
}

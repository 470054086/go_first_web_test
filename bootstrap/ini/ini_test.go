package ini

import "testing"

func TestIniCfg_GetKey(t *testing.T) {
	New(fileName)
	key := Cfg.GetSelect("Mysql").GetKey("port")
	if key != "80" {
		t.Fatal("读取错误")
	}
}

func BenchmarkIniCfg_GetKey(b *testing.B) {
	New(fileName)
	for i:=0 ;i<b.N;i++ {
		Cfg.GetSelect("Mysql").GetKey("port")
	}
}

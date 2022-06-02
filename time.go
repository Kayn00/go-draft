package main

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/text/gstr"
	"runtime"
)

func main() {
	UpdateSystemDate("2020-03-20 15:02:41.372")
}

func UpdateSystemDate(dateTime string) bool {
	system := runtime.GOOS
	switch system {
	case "windows":
		{
			_, err1 := gproc.ShellExec(`date  ` + gstr.Split(dateTime, " ")[0])
			_, err2 := gproc.ShellExec(`time  ` + gstr.Split(dateTime, " ")[1])
			if err1 != nil && err2 != nil {
				glog.Info("更新系统时间错误:请用管理员身份启动程序!")
				return false
			}
			return true
		}
	case "linux":
		{
			glog.Info("111111")
			_, err1 := gproc.ShellExec(`date -s  "` + dateTime + `"`)
			if err1 != nil {
				glog.Info("更新系统时间错误:", err1.Error())
				return false
			}
			return true
		}
	case "darwin":
		{
			glog.Info("2222")
			_, err1 := gproc.ShellExec(`sudo systemsetup -setusingnetworktime off`)
			if err1 != nil {
				glog.Info("更新系统时间错误1:", err1.Error())
				return false
			}
			_, err2 := gproc.ShellExec(`sudo systemsetup -setdate 09/02/22`)
			if err2 != nil {
				glog.Info("更新系统时间错误2:", err2.Error())
				return false
			}
			return true
		}
	}
	return false
}
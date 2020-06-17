package utils

import (
  "net"
)

// 获取本机IP地址及mac地址
func GetIP() (string, string) {
  ifs, err := net.Interfaces()
  if err != nil {
    panic(err)
  }

  for _, v := range ifs {
    if (v.Flags & net.FlagLoopback) != 0 {
      continue
    }
    if (v.Flags & net.FlagUp) == 0 {
      continue
    }
    addrs, err := v.Addrs()
    if err != nil {
      panic(err)
    }

    for _, address := range addrs {
      // 检查ip地址判断是否回环地址
      if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
        if ipnet.IP.To4() != nil {
          return ipnet.IP.String(), v.HardwareAddr.String()
        }
      }
    }
  }

  panic("can't find valid IP!")
}

# SunloginRCE

向日葵11.x版本RCE扫描利用程序，基于fscan能力，添加了主机存活探测、网段扫描和并发优化等功能，方便安全人员在内网做漏洞检测和修复

# 说明

代码魔改自 https://github.com/Mr-xn/sunlogin_rce 项目，致谢  

考虑到原项目没有写网段扫描的功能，于是我复用和适配了[fscan](https://github.com/shadow1ng/fscan)的主机存活探测、端口扫描以及IP地址解析等相关代码

#### 代码仅做内部安全检测使用，请勿用于非法渗透活动，请遵纪守法

# 使用

```
Usage of sunRce.exe:
  -c string
        Input Cmd Command
  -h string
        IP Address: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12
  -n int
        Set Scan Threads (default 600)
  -p string
        Ports Range: 50000 | 40000-65535 | 49440,51731,52841 (default "50000")
  -t string
        Choose Attack Type: scan | rce
```
### 扫描单个IP:  
![scanone](https://github.com/ce-automne/SunloginRCE/blob/main/scanone.png)
### 扫描IP段：  
![scanmore](https://github.com/ce-automne/SunloginRCE/blob/main/scanmore.png) 
### 命令执行：   
![exp](https://github.com/ce-automne/SunloginRCE/blob/main/exp.png)

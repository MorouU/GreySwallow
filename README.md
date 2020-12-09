# GreySwallow 1.05
一个简单的流量加密和转发工具
## 原理
  * 在进行端口转发的同时将流量加密/解密
## 功能
  * 支持 **TCP/UDP** 端口转发
  * 支持 **AES128/AES192/AES256/DES/RSA** 加密方式
  * 支持 **CBC/ECB/CFB** 加密模式(仅限*AES*和*DES*)
  * 支持 **当前输入/文件/URL/TCP** 方式获取密钥源
  * 支持 **A->B** 以及 **B->A** 两个方向的加密/解密
  * 支持使用 **GET** 和 **POST** 请求方法、自定义 **GET** 或 **POST** 请求参数(在使用URL获取密钥源时)
## 使用方法（for Linux or Windows）
  * __参数设定__
    * -l, --listen string           Listen -> [TCP/UDP:IP:PORT]
    * -c, --connect string          Connect -> [TCP/UDP:IP:PORT]
    * -m, --encrypt-method string   Encrypt-Method -> [AES128/AES192/AES256/DES/RSA/NULL] (default "NULL")
    * -M, --encrypt-mode string     Encrypt-Mode -> [CBC/ECB/CFB] (default "CBC")
    * -s, --encrypt-source string   Encrypt-Source -> [KEY/Filename/URL/(IP):(PORT)]
    * -f, --encrypt-from string     Encrypt-From -> [String/File/URL/TCP] (default "String")
    * -b, --source-b64-de           Source-Base64-Decode -> [True/False]
    * -t, --turn-method string      Turn-Method -> [EN/DE/NULL] (default "NULL")
    * -a, --accept-method string    Accept-Method -> [EN/DE/NULL] (default "NULL")
    * -r, --url-method string       Url-Method -> [GET/POST] (default "GET")
    * -p, --url-params string       Url-Params -> [QueryString]
    * -d, --url-data string         Url-Data -> [BodyString]
 
  * __简单示例__
    * 端口转发:
    ```
     ./gs -l tcp:127.0.0.1:3339 -c tcp:127.0.0.1:3340 
    ```
    * 单向加密：
    ```
     # -> 使用AES192进行加密
     # -> 加密模式为ECB
     # -> 密钥为 BMV587BMV587BMV587BMV587
     ./gs -l tcp:127.0.0.1:3339 -c tcp:127.0.0.1:3340 -m AES192 -s Qk1WNTg3Qk1WNTg3Qk1WNTg3Qk1WNTg3Cg== -b -M ECB -t EN
     
     # -> 使用AES256进行加密
     # -> 加密模式为CBC
     # -> 密钥从 ./serect.txt 文件获取
     ./gs -l tcp:127.0.0.1:3339 -c tcp:127.0.0.1:3340 -m AES256 -s ./serect.txt -f file -M CBC -t EN
    ```
    * 双向加密：
    ```
    # -> 使用AES256进行加密
    # -> 加密模式为CBC
    # -> 密钥从 ./serect.txt 文件获取
     ./gs -l tcp:127.0.0.1:3339 -c tcp:127.0.0.1:3340 -m AES256 -s ./serect.txt -f file -t EN -a DE
     ./gs -l tcp:127.0.0.1:3340 -c tcp:127.0.0.1:3341 -m AES256 -s ./serect.txt -f file -t DE -a EN
    
    # -> 使用DES进行加密
    # -> 加密模式为ECB
    # -> 密钥为 BMV587BM
     ./gs -l tcp:127.0.0.1:3339 -c tcp:127.0.0.1:3340 -m DES -s BMV587BM -M ECB -t EN -a DE
     ./gs -l tcp:127.0.0.1:3340 -c tcp:127.0.0.1:3341 -m DES -s BMV587BM -M ECB -t DE -a EN
     
    ```
# 后记
*该玩意儿仅仅是用来学习和练习Go之用！！！*

### go 常用的小工具

+  compress 基于 Gzip
+  convert  主要是转换一些数字与字符
+  password 基于 "golang.org/x/crypto/bcrypt"  生成密码与校验密码
+  rand     提供随机生成的数字或字符串
+  str      一些字符串处理操作
+  struct   把结构体转为map
+  time     提供时间类型， 重写了UnmarshalJSON方法，解析 null,"" 为时间时不会报错
+  uuid     提供生成UUID方法，包括（UUID, UUID去掉'-', UUID 转19位数字
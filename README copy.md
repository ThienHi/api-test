# GO-REST API 
Project này sử dụng echo framework và GORM 

Echo framework là framework khá phổ biến trong golang được sử dụng để làm API. Nó có các module trung gian khá hữu ích trong api chẳng hạn như compression, proxying, loggingv ...

# Cấu trúc source code
`tree -d`
```
├── api => Định nghĩa logic xử lý trong API
├── config => Định nghĩa các cài đặt biến môi trường ( Database connection, server v.v.)
├── models => Định nghĩa table trong database qua  GORM
├── router => Chứa route API
├── server.go => File main để run app
```
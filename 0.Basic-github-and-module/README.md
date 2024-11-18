# Intern-Golang

## 1. Module
- Module: tập hợp tất cả các package , mỗi 1 module sẽ có 1 version và sẽ được quản lí ở file "go.mod"
- File go.mod: quản lí version của module và quản lí các package mình tải về và quản lí version của các package
- File go.sum: Golang sẽ hash(SHA256) tất cả các package mình tải về , khi đó Go sẽ so sánh hash của thư viện đó khi mình cập nhật. 
- Ví dụ: 4.2.1 
Major version: 4 khi mà thay đổi lớn nó không tương thích được với các code cũ
Minor version: 2 thay đổi mà vẫn tương thích được code cũ
Fix bug version: 1 khi mà fix bug version.

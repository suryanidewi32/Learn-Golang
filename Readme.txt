How to run Mongo API :

MongoDB Configuration =

Host = mongodb://localhost:27017
Database = atopdatadb
Collection  = user_auth

*API saat ini berdasarkan struktur MongoDB saya, Teman2 boleh coba-coba ganti.

How to run =

Require = gin-gonic API Framework (https://github.com/gin-gonic/gin)

1. Buat directory/folder github.com di folder root Golang (C:\Go\src)-> Path tergantung instalasi awal.
2. Extract folder Mongo-API-Test
3. Buka command prompt dan pindah ke dir/folder Mongo-API-Test
4. Install gin-gonic, run command "go get github.com/gin-gonic/gin"
5. run command go build
6. run command go run main.go


Main Crud Route :
1. Get = /todos
2. Post = /todo

*Route lain dapat dicek di code : "route.go" di dalam route folder



## 生成模型
``go run main.go make model category``
## 生成迁移文件
``go run main.go make migration add_categories_table``
## 执行迁移
``go run main.go migrate up``
## 生成request文件
``go run main.go make request category``
## 生成controller文件
``go run main.go make apicontroller v1/category``
## 创建factory文件
``go run main.go make factory category``
## 创建seeder文件
``go run main.go make seeder category``
## 填充指定 seeder
``go run main.go seed SeedCategoriesTable``
## 填充全部 seeder
``go run main.go seed``
## 生成policy文件
``go run main.go make policy category``

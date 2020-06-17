webapp starter project
---

**Feature**:
1. Use gin as router;
2. Use gorm to abstract db operation;
3. Package assets into executable, easy to deploy;

**QA**:
1. How to embed assets?

    Put your assets(html/js/css/images/mp3) into assets directory, and run `go run assets_generate.go`.<br>
    The generated file `assets_vfsdata.go` should contains all the contents within the directory.

2. How to add new data model?

    Define your model's structure in model package, make sure it implements the `migratable` interface,<br>
    finally modify `dbTables` variable, tables will be automatically created/updated on server's startup.
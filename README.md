webapp starter project
---

**Feature**:
1. Use gin as router;
2. Use gorm to abstract db operation;
3. Package assets into executable, easy to deploy;

**QA**:
1. How to embed assets?

    Put your assets(html/js/css/images/mp3) into directory `assets`, everthing will be loaded<br>
    into `embed.FS` during compilation.

2. How to add new data model?

    Define your model's structure in model package, make sure it implements the `migratable` interface,<br>
    finally modify `dbTables` variable, tables will be automatically created/updated on server's startup.
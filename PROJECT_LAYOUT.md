# Go project layout

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

```
- projectName     // folder with a project inside
|- .git
|- 📁 cmd        // place for project entry point "main.go"
|- 📁 internal   // code that is used only in the project
  |- 📁 handler
  |- 📁 model
  |- 📁 repository
  |- 📁 service
|- 📁 pkg        // code that you can import from an external project
|- 📁 api        // OpenAPI and protocol definitions
```

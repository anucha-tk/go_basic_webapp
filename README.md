# GO BASIC WEBAPP

http server-side render page

<!--toc:start-->

- [GO BASIC WEBAPP](#go-basic-webapp)
  - [Techs and Tools](#techs-and-tools)
  - [How to run](#how-to-run)
  - [Tips](#tips) - [port error](#port-error)
  <!--toc:end-->

## Techs and Tools

- http
- templates
  - html
  - cache the templates
- css
  - tailwindcss
- infrastructure
  - commitlint and linter

## How to run

```bash
make dev
```

## Tips

### port error

```bash
# find open port
lsof -i :[port number]
# kill
kill -9 [PID]
```

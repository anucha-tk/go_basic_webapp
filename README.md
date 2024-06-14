# GO BASIC WEBAPP

http server-side render page and HTMX

<!--toc:start-->

- [GO BASIC WEBAPP](#go-basic-webapp)
  - [Techs and Tools](#techs-and-tools)
  - [How to run dev](#how-to-run-dev)
  - [Port error](#port-error)
  <!--toc:end-->

## Techs and Tools

- http
- router
  - chi
- middlewares
  - CSRF
  - logger
  - session
- templates
  - html(htmx)
  - cache the templates
- css
  - tailwindcss(CDN)
- infrastructure
  - commitlint and linter
- simple htmx
  - get time every 5 sec

## How to run dev

```bash
make dev
```

## Port error

```bash
# find open port
lsof -i :[port number]
# kill
kill -9 [PID]
```

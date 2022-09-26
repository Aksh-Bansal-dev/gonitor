# gonitor

A simple tool to monitor linux servers.

## Features

It provides endpoints to:

- `/stats` gets CPU, RAM, and Disk usage
- `/healthcheck` pings all endpoints mentioned in the config
- more coming soon...

## Config

`gonitor.json` must be placed in home folder.

```json
{
  "endpoints": [
    {
      "name": "Google",
      "url": "https://www.google.com"
    },
    {
      "name": "booo",
      "url": "http://localhost:3000"
    }
  ]
}
```

## Todo

- [ ] ping all apps on the server to check if they are running every 30min.
- [ ] send message on discord when cpu, ram or disk usage is high.
- [ ] send message on discord if any app crashes.

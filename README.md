# stay-alive

A Go application that keeps serverless instances alive by periodically pinging Redis servers.

## Configuration

The application uses a JSON configuration file (`config.json`) to specify Redis servers and their ping schedules:

```json
{
  "servers": [
    {
      "url": "redis://localhost:6379",
      "schedule": "*/5 * * * *"
    }
  ]
}
```

- `url`: Redis server URL
- `schedule`: Cron expression for ping frequency

### Cron Expression Format

```
┌────────────── minute (0 - 59)
│ ┌──────────── hour (0 - 23)
│ │ ┌────────── day of month (1 - 31)
│ │ │ ┌──────── month (1 - 12)
│ │ │ │ ┌────── day of week (0 - 6) (Sunday to Saturday)
│ │ │ │ │
* * * * *
```

Common cron expressions:

- `*/5 * * * *` - Every 5 minutes
- `0 * * * *` - Every hour
- `0 */2 * * *` - Every 2 hours
- `0 0 * * *` - Once a day at midnight
- `*/15 * * * *` - Every 15 minutes

## Usage

1. Configure your Redis servers in `config.json`
2. Run the application:
   ```bash
   go run main.go
   ```

## Requirements

- Go 1.16 or higher
- Redis server(s)

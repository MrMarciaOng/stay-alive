# stay-alive

A Go application that keeps serverless instances alive by periodically pinging Redis servers.

![Stay-Alive Banner](https://img.shields.io/badge/STAY--ALIVE-Go_Redis_Pinger-1E1E1E?style=for-the-badge&logo=go&logoColor=00ADD8&labelColor=DC382D)

![Go Version](https://img.shields.io/badge/go-1.18+-00ADD8?logo=go)
![Redis](https://img.shields.io/badge/redis-%23DC382D.svg?logo=redis&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-blue)

## 📋 Table of Contents
- ⚙️ [Configuration](#configuration)
- 💻 [Installation](#installation--running)  
- 🧪 [Running Locally](#running-locally-for-development)
- ✅ [Verification](#verification)
- 🤝 [Contributing](#contributing)
- 👥 [Contributors](#contributors)

## Requirements

- Go 1.18 or higher (tested with 1.20, compatible with 1.21+)
- Redis server(s)

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

## Installation & Running

### Development
```bash
# Clone the repository
git clone https://github.com/<replace-github-username>/stay-alive.git
cd stay-alive

# Build and run
go run main.go
```

### Production Deployment
1. Build the binary:
```bash
go build -o stay-alive  # Compiles Go code into standalone executable
```

2. Run with config (background mode):
```bash
nohup ./stay-alive > stay-alive.log 2>&1 &
# Breakdown:
# nohup      - Prevents termination when terminal closes
# >          - Redirects stdout to stay-alive.log
# 2>&1       - Redirects stderr to stdout (same log file)
# &          - Runs process in background
```

3. Verify:
```bash
tail -f stay-alive.log  # Shows last 10 lines + follows new logs
# Ctrl+C to exit
```

## Running Locally (for development)

1. Configure your Redis servers in `config.json`
2. Run the application:
   ```bash
   go run main.go
   ```

## Verification

### Application Output (Successful)
```
2025/04/04 19:19:01 Redis server ping successful - Name: Local Redis, Response: PONG, Time: 2025-04-04T19:19:01+08:00
2025/04/04 19:19:01 Added cron job for Redis server Local Redis with schedule: */1 * * * *
2025/04/04 19:19:01 Started Redis ping service
2025/04/04 19:20:00 Redis server ping successful - Name: Local Redis, Response: PONG, Time: 2025-04-04T19:20:00+08:00
```

### Redis CLI Monitor Output
```
1743765600.002450 [0 [::1]:52971] "ping"
1743765660.001749 [0 [::1]:53035] "ping"
```

Expected behavior:
- Initial immediate ping on startup
- Subsequent pings every Nth minute (based on config.json schedule)
- Matching timestamps between application and Redis logs

## Contributors

### Developers
- [@MrMarciaOng](https://github.com/MrMarciaOng) - Project creator  
- [@sayyidkhan](https://github.com/sayyidkhan) - Contributor

### How to Contribute
1. 🍴 Fork the repository  
2. 🌱 Create a feature branch (`git checkout -b feature/your-feature`)  
3. 💻 Make your changes  
4. 📝 Commit your changes (`git commit -m 'Add some feature'`)  
5. 🔀 Push to the branch (`git push origin feature/your-feature`)  
6. ✨ Open a Pull Request

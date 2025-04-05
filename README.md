# stay-alive

A Go application that keeps serverless instances alive by periodically pinging Redis servers.

## Inspiration
> "Been learning Go on the side, so I decided to build something to fix a recurring pain point as a dev. My serverless Redis instances from [Upstash](https://upstash.com/) kept shutting down due to inactivity. (I don't blame them - I have 20+ instances!)"  
> — @MrMarciaOng

![Stay-Alive Banner](https://img.shields.io/badge/STAY--ALIVE-Go_Redis_Pinger-1E1E1E?style=for-the-badge&logo=go&logoColor=00ADD8&labelColor=DC382D)

![Go Version](https://img.shields.io/badge/go-1.18+-00ADD8?logo=go)
![Redis](https://img.shields.io/badge/redis-%23DC382D.svg?logo=redis&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-blue)

## 📋 Table of Contents
- ⚙️ [Configuration](#configuration)
- 💻 [Installation](#installation--running)  
- 🧪 [Running Locally](#running-locally-for-development)
- ✅ [Verification](#verification)
- 📦 [Release Process](#release-process-3-steps)
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

## Release Process (3 Steps)

1. **Verify**  
```bash
# Check existing tags (sorted by version):
git tag --list -n | sort -V
```
- Visit `https://github.com/sayyidkhan/stay-alive/releases`  
- Download the binary under "Assets" 

2. **Tag Your Release**  
```bash
git tag -a v0.1.0 -m "Test release"
git push origin v0.1.0
```

3. **Wait 1 Minute**  
- GitHub will automatically:  
  ✅ Build the binary  
  ✅ Create a release  
  ✅ Upload `stay-alive`   

⚠️ **Troubleshooting**:  
- If the release doesn't appear:  
```bash
# Check workflow runs:
gh run list -w "Build and Release"
```

## Contributors

### Developers
<table>
  <tr>
    <td align="center">
      <a href="https://github.com/MrMarciaOng">
        <img src="https://avatars.githubusercontent.com/u/24979131" width="100px;" alt="MrMarciaOng"/>
        <br />
        <sub><b>MrMarciaOng</b></sub>
      </a>
      <br />Project Creator
    </td>
    <td align="center">
      <a href="https://github.com/sayyidkhan">
        <img src="https://avatars.githubusercontent.com/u/22993048" width="100px;" alt="sayyidkhan"/>
        <br />
        <sub><b>sayyidkhan</b></sub>
      </a>
      <br />Contributor
    </td>
  </tr>
</table>

### How to Contribute
1. 🍴 Fork the repository  
2. 🌱 Create a feature branch (`git checkout -b feature/your-feature`)  
3. 💻 Make your changes  
4. 📝 Commit your changes (`git commit -m 'Add some feature'`)  
5. 🔀 Push to the branch (`git push origin feature/your-feature`)  
6. ✨ Open a Pull Request

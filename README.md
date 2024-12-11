# ⏰ cronik

```text

                          _  _
  ___  _ __   ___   _ __  (_)| | __
 / __|| '__| / _ \ | '_ \ | || |/ /
| (__ | |   | (_) || | | || ||   <
 \___||_|    \___/ |_| |_||_||_|\_\

```

⏰ cronik: A Golang-based scheduler for periodic, delayed, and retryable jobs with robust error handling.

## ✨ Features

- **Periodic Jobs**: Schedule jobs to run at regular intervals (e.g., every minute, hour, day).
- **Delayed Jobs**: Schedule jobs to run after a specified delay (e.g., 10 minutes from now).
- **Retryable Jobs**: Automatically retry jobs that fail, with configurable retry policies.
- **Robust Error Handling**: Handle errors gracefully and provide detailed error logs.
- **Concurrency Control**: Manage the number of concurrent jobs to optimize resource usage.
- **Job Prioritization**: Assign priorities to jobs to ensure critical tasks are executed first.
- **Flexible Scheduling**: Support for cron expressions and other flexible scheduling options.
- **Persistence**: Save job states to a database to ensure jobs are not lost on application restart.
- **Monitoring and Alerts**: Monitor job execution and send alerts on failures or other important events.
- **Extensible**: Easily extend the scheduler with custom job types and scheduling strategies.

## 🚀 Installation

To install `cronik`, follow these steps:

1. **Install Go**: Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

2. **Get the `cronik` package**: Use `go get` to download the `cronik` package.

   ```sh
   go get github.com/trinhminhtriet/cronik
   ```

3. **Import `cronik` in your project**: Add the following import statement to your Go code.

   ```go
   import "github.com/trinhminhtriet/cronik"
   ```

Now you are ready to use `cronik` in your project!

## 💡 Usage

Here's a quick example of how to use `cronik` to schedule jobs:

1. **Import the package**:

```go
import (
    "fmt"
    "time"
    "github.com/trinhminhtriet/cronik"
)
```

2. **Start the scheduler**:

   ```go
   scheduler.Start()
   ```

3. **Cancel the scheduler** (when your application is shutting down):

   ```go
   scheduler.Cancel()
   ```

4. **Schedule a job**:

```go
// Every 2 minutes
t.When{Each: "2m"}

// Every 100 milliseconds
t.When{Every: t.Every(100).Milliseconds()}

// Every hour at :30
t.When{Every: t.Every(1).Hours(), At: "**:30"}

// Every day at the next beginning of an hour **:00
t.When{Every: t.Every(1).Days(), At: "**:00"}

// Every 2 weeks on Saturdays at 10:00
t.When{Every: &t.Every(2).Weeks(), On: t.Sat, At: "10:00"}

// Saturday at 15:00, not repeated
t.When{Day: t.Sat, At: "15:00"}

// Every week on Sun at 11:00, last run was explicitly given.
// If your process shuts down at 10:00 on Sunday, it allows scheduler
// to schedule the job to run in a hour on an immediate restart.
t.When{LastRun: lastRun, Every: &t.Every(1).Weeks(), On: t.Sun, At: "10:00"}
```

## 🤝 How to contribute

We welcome contributions!

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m "feat: my new feature"`;
- Push to your branch: `git push origin my-feature`.

Once your pull request has been merged, you can delete your branch.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

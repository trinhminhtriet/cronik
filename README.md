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

4. **Run your project**: Use `go run` to run your project or `go build` to build it.
   ```sh
   go run main.go
   ```

Now you are ready to use `cronik` in your project!

## 💡 Usage

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

2. **Create a new scheduler**:

   ```go
   scheduler := cronik.NewScheduler()
   ```

3. **Define a job**:

   ```go
   job := func() {
       fmt.Println("Job executed at", time.Now())
   }
   ```

4. **Schedule the job**:

   - **Periodic Job**: Schedule the job to run every minute.

     ```go
     scheduler.SchedulePeriodicJob(job, time.Minute)
     ```

   - **Delayed Job**: Schedule the job to run after a 10-minute delay.

     ```go
     scheduler.ScheduleDelayedJob(job, 10*time.Minute)
     ```

   - **Retryable Job**: Schedule the job with retry policy (e.g., retry up to 3 times with a 1-minute interval).

     ```go
     scheduler.ScheduleRetryableJob(job, 3, time.Minute)
     ```

   - **Interval Job**: Schedule the job to run at specific intervals (e.g., every 5 minutes).
     ```go
     scheduler.ScheduleIntervalJob(job, 5*time.Minute)
     ```

5. **Start the scheduler**:

   ```go
   scheduler.Start()
   ```

6. **Stop the scheduler** (when your application is shutting down):
   ```go
   scheduler.Stop()
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

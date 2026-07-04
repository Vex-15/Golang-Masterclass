# Golang Masterclass

This repository is where I document my journey of learning Go by building real projects.

Instead of keeping notes that I'll probably never revisit, I wanted one place where I could experiment, make mistakes, fix them, and gradually build a collection of projects that showcase different parts of the language.

Most folders focus on a single concept or feature, while some combine multiple ideas into small practical projects.

If you're learning Go too, I hope you find something useful here.

---

## What You'll Find

Each folder covers a different topic as I progress through Go.

Current topics include:

### Concurrency

Projects exploring Go's concurrency model using:

- Goroutines
- WaitGroups
- Channels
- Buffered Channels
- Semaphore Pattern
- Context Cancellation
- Synchronization

Projects:

- Concurrent File Downloader
- Ping-Pong with Context Cancellation
- Goroutine Basics

More topics will be added as I continue learning.

---

## Repository Structure

```text
Golang-Masterclass
│
├── goroutines/
│   ├── fileDownloader/
│   ├── ...
│
├── README.md
```

The structure will continue to grow as I cover more Go concepts.

---

## Why This Repository?

I learn best by building.

Rather than writing dozens of isolated examples, I prefer creating projects that naturally introduce new concepts.

Every project teaches me something different—whether it's concurrency, networking, file handling, or backend development.

This repository is essentially my public learning journal.

---

## Featured Project

### Concurrent File Downloader

A small project that downloads multiple files concurrently while exploring some of Go's most useful concurrency primitives.

Concepts covered:

- Goroutines
- WaitGroups
- Channels
- Semaphore Pattern
- Concurrent HTTP Requests
- Download Metrics

I also wrote a detailed article explaining how the project works:

📖 I also wrote a detailed walkthrough explaining how the project works on Medium:

**[Building a Concurrent File Downloader in Go: A Practical Guide to Goroutines](https://medium.com/@vexstack/building-a-concurrent-file-downloader-in-go-a-practical-guide-to-goroutines-4071dd01f946)**

---

## Running a Project

Clone the repository:

```bash
git clone https://github.com/Vex-15/Golang-Masterclass.git
```

Move into the repository:

```bash
cd Golang-Masterclass
```

Run any project:

```bash
cd goroutines
go run .
```

---

## Contributing

This is primarily a personal learning repository, but if you notice something that can be improved, feel free to open an issue or submit a pull request.

I'm always open to learning better ways of writing Go.

---

## Connect

GitHub: https://github.com/Vex-15

X: https://x.com/vexstack

Medium: https://medium.com/@vexstack

---

If you find this repository useful, consider giving it a ⭐.

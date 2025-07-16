# 🚀 gitinit – One Command to Git Init, Commit & Push

[![Go Version](https://img.shields.io/badge/go-1.21+-brightgreen)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

`gitinit` is a powerful CLI tool written in Go that automates Git repository setup. Instead of manually typing `git init`, `git add .`, `git commit`, etc., run one command: `gitinit`.

---

## ✨ Features

- ✅ Initializes a Git repo in the current directory
- ✅ Adds and commits all files with a custom message
- ✅ Renames the default branch (e.g. to `main`)
- ✅ Prompts for and sets remote origin URL
- ✅ Pushes the initial commit to remote
- ✅ Adds `.gitignore` for Go, Node.js, Python, or none
- ✅ User-friendly CLI prompts (like VS Code)

---

## 🛠 Installation

### 🔁 1. Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/gitinit.git
cd gitinit
```

## 🛠 Build

```bash
go build -o gitinit
```

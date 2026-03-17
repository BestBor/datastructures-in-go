# Data Structures in Go

Implementations of classic data structures in Go, built from scratch with an educational and reusable approach.

This repository exists for two main reasons:
1. To understand how data structures really work in Go.
2. To build clear, well-tested, and reusable utility packages.

---

## 📦 Repository Structure

Each data structure lives in its **own package**:

```
datastructures-in-go/
├── go.mod
├── README.md
├── linkedlist/
│   └── linkedlist.go
├── stack/
│   └── stack.go
├── queue/
│   └── queue.go
├── bst/
│   └── bst.go
└── trie/
    └── trie.go
```

This organization allows:
- Reusing data structures independently
- Importing only what is necessary
- Keeping tests simple and localized

---

## 🧠 Principles

- No external dependencies
- Idiomatic Go code
- Small, well-defined packages
- Tests from the start
- Clarity > premature optimization

---

## 📚 Available Structures

| Structure | Package | Type |
|---|---|---|
| Linked List | `linkedlist` | Linear |
| Stack | `stack` | Linear |
| Queue | `queue` | Linear |
| Binary Search Tree | `bst` | Tree |
| Trie | `trie` | Tree |

## 🧪 Tests

```bash
go test ./...
```

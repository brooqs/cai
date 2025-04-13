
# cai – Command + AI 🚀

**cai** is a minimalist and practical command-line tool written in **Go**, allowing users to quickly fetch Linux commands or configuration snippets directly from the terminal using OpenAI's GPT models. It streamlines productivity by providing precise commands instantly, without explanations or unnecessary formatting.

## 🌟 Features

- **Instant commands**: Quickly get Linux commands or configuration lines by simply typing your query.
- **Configurable**: Easily customizable through a `config.ini` file for API keys, GPT models, and token limits.
- **Clean Output**: Automatically removes markdown formatting, code blocks, and unnecessary terms like `plaintext` for clean, copy-paste friendly results.
- **Static Build**: Compiled as a static Go binary, making deployment and distribution effortless.

## ⚡ Installation

### Clone Repository

```bash
git clone https://github.com/brooqs/cai.git
cd cai
```

### Dependencies

```bash
go mod tidy
```

### Build Static Binary

```bash
CGO_ENABLED=0 go build -ldflags "-s -w" -o cai ./cmd/cai/
```

### Move Binary to PATH

```bash
sudo mv cai /usr/local/bin/
```

## ⚙ Configuration

Create a configuration file at `/etc/cai/config.ini` with your OpenAI credentials:

```ini
[openai]
api_key = YOUR_API_KEY
model = gpt-4-turbo
max_tokens = 150
```

Set secure permissions:

```bash
sudo chmod 600 /etc/cai/config.ini
```

## 🚀 Usage

Simply invoke `cai` followed by your prompt in quotes:

```bash
cai "add new disk partition to fstab"
```

### Example Output:

```bash
UUID=xxxxx-xxxxx-xxxxx /mnt/disk ext4 defaults 0 2
```

## 📁 Project Structure

```
cai/
├── cmd/
│   └── cai/
│       └── main.go
├── config/
│   └── config.go
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## 💻 Contributions

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/brooqs/cai/issues).

## 📄 License

This project is licensed under the MIT License.

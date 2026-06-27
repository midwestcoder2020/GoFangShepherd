# 🐺 FangShepherd

**Author:** midwestcoder2020
**Version:** 1.0  
**Date:** June, 2026

---

## 🗂️ Overview

**FangShepherd** is a lightweight, terminal-based tool for **extracting**, **defanging**, and **refanging** Indicators of Compromise (IOCs). Built with threat analysts and digital forensic investigators in mind, FangShepherd helps sanitize malicious URLs, IPs, emails, and hashes for safe sharing and analysis.

---

## ✨ Features

- 🔍 **IOC Extraction**: URLs, IP addresses, email addresses, MD5, SHA1, SHA256 hashes
- 🦷 **Defang/Refang**: Safe IOCs for sharing or revert to original form for analysis
- 📄 **File I/O Support**: Read from and save to text files
- 🎨 **CLI Art Banner**: Terminal ASCII art for style using `pyfiglet`

---

## 📦 Installation

### Requirements
- Python 3.x
- [pyfiglet](https://pypi.org/project/pyfiglet/)

### Install `pyfiglet`
```bash
pip install pyfiglet
```

---

## 🚀 Usage



Run the script via terminal:

### Python

```bash
python3 FangShepherd.py
```

### golang
Run the script via terminal:

```bash
go run .
```

```bash
go run main.go
```

### Main Menu Options
| Option | Description                          |
|--------|--------------------------------------|
|   1    | Extract IOCs + Defang                |
|   2    | Extract IOCs + Refang                |
|   3    | Just Defang (line-by-line)           |
|   4    | Just Refang (line-by-line)           |

---

## 🧪 Example Workflow

1. Choose action: e.g., `1` for Extract + Defang
2. Select input method:  
   - `1` Paste text  
   - `2` Read from file
3. Paste your content or provide a file path  
4. Processed IOCs are printed to screen  
5. Optionally save to an output file

---

## ⚙️ Defang / Refang Rules

| Original | Defanged   |
|----------|------------|
| `http`   | `hxxp`     |
| `.`      | `[.]`      |
| `@`      | `[@]`      |

---

## 🗃️ Sample IOC Extraction

**Input Text:**
```
http://malicious.com/path
192.168.1.1
user@example.com
d41d8cd98f00b204e9800998ecf8427e
```

**Defanged Output:**
```
hxxp://malicious[.]com/path
192[.]168[.]1[.]1
user[@]example[.]com
d41d8cd98f00b204e9800998ecf8427e
```

---

## 💾 Save Results

After processing, choose to save results:
```text
Save results to a file? (y/n): y
Enter output file name (e.g., output.txt):
```

---

## 📜 License

This project is licensed under the MIT License.  
© 2025 Alayna Ferdarko. All rights reserved.

---

## ⚠️ Disclaimer

This tool is intended for educational and professional use only. While FangShepherd helps reduce the risk of accidental interaction with live IOCs, always handle threat data in a secure and controlled environment.

---

## 🤝 Contributions

Contributions, issues, and feature requests are welcome!  
Feel free to fork this repo and submit a pull request.

---

## 📬 Contact

For inquiries or feedback:  
📧 alaynavendetta@proton.me

# LetterBytes

**LetterBytes**: Converting paper letters into digital bytes!

## Overview

LetterBytes is a handy tool designed to help you organize and digitalize all German letters that arrive in your mailbox. This program scans, translates, summarizes, and archives your letters, turning your paper piles into a well-organized digital archive.

## Features

- [x] **Scan**: Instantly digitizes your letters.
- [ ] **Translate**: Converts German text to English.
- [ ] **Summarize**: Provides a concise summary and highlights important action items.
- [ ] **Archive**: Stores your scanned letters in a database, searchable by OCR.
- [ ] **Fun**: Makes your life easier and adds a bit of humor to your day!

## Prerequisites

- Go
- Environment variable `LETTERBYTES_SCANNER` set to your scanner's name
- [ ] Google Cloud Translation API
- [ ] Tesseract OCR
- [ ] PostgreSQL (for storing BLOBs)

## Usage

1. **Set the environment variable**:
   
First, run the program to find available scanners:
   
```sh
❯ ./letterbytes 
❌ environment variable LETTERBYTES_SCANNER not set
🚀 finding scanners...
possible values for env variable LETTERBYTES_SCANNER:

LETTERBYTES_SCANNER='escl:https://192.168.0.123:443'
LETTERBYTES_SCANNER='airscan:e0:Canon TS5300 series'
```

```sh
❯ ./letterbytes doc
Image scanned and saved to doc.png
Image scanned and saved to doc.tiff
```
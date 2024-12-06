# Email Marketing API

## Overview

A Go API to manage an email marketing campaign with the following endpoints:
1. Upload a CSV of emails.
2. Add a single email to the in-memory list.
3. Remove an email from the list.
4. Upload a JPG or PDF to send as an attachment to the emails.

## Prerequisites

- [Go](https://golang.org/) 1.18+
- [AWS SDK for Go](https://github.com/aws/aws-sdk-go)
- AWS SES configured and verified email addresses

## Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd email-marketing-api
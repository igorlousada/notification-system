# Project Name
Notification Sending System
## Description

This project is backend software responsible for notification system for purchase. The objective of this application
is to respond to a HTTP request from an client to notify that a purchase has been made. The notification should be one-to-multiple
channels (slack, email, text).

The system has the following constraints:
 - Horizontally scalable.
 - At least once for sending message
 - Easy extendable for new channels

### Solution
With the above description, the following solution has been create:

[Apache Kafka](https://kafka.apache.org/) as message broker. When the backend receives a valid message,
it sends the message to Kafka. After the message is correctly persisted in Kafka, we return to the client
a HTTP status 200. 

## Requirements

- Go 1.x or higher
- Kafka (if applicable for this project)
- [Other dependencies or external services]

## Installation

### 1. Clone the repository:
```bash
git clone https://github.com/yourusername/projectname.git
cd projectname
# PackChat

A minimal, fast, and responsive real-time chat application using WebSockets and MessagePack for efficient data transfer.

## Features
- Lightweight and fast communication with binary MessagePack encoding
- Real-time messaging with WebSockets
- Multiple chat rooms support


## Technology Stack
- Backend: Go
- Frontend: HTML, CSS, JavaScript
- Data Transport: WebSockets
- Data Format: MessagePack
- Styling: PureCSS (minimal dependency)

## Getting Started

### Prerequisites

- Go 1.18 or later

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/isErenG/packchat.git
   cd packchat
   ```

2. Run the application:
   ```bash
   go run cmd/main.go
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

## How It Works

PackChat uses WebSockets for real-time communication between clients and the server. All messages are serialized using MessagePack, a binary format that's more efficient than JSON, resulting in faster transmission and lower bandwidth usage.

When you send a message, it's compressed using MessagePack and sent to the server, which then broadcasts it to all clients in the same room. Messages include the sender's information, content, and timestamp.

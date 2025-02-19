# WatchParty

WatchParty is a demo web application that allows users to have a **live chat while watching a video**. Inspired by **YouTube** and **Discord**, it provides a shared viewing experience with real-time interaction. The project is built using **Go** for both backend and frontend rendering, ensuring a fast and lightweight experience.

## Features
- **Video Streaming**: Watch pre-recorded videos seamlessly.
- **Live Chat**: Engage with others through real-time chat while watching a video.

## Technologies Used

### **Go Packages**
- [`html/template`](https://pkg.go.dev/html/template) - For server-side HTML rendering.
- [`github.com/gorilla/websocket`](https://pkg.go.dev/github.com/gorilla/websocket) - For real-time live chat functionality.
- [`net/http`](https://pkg.go.dev/net/http) - For handling HTTP requests and serving static files.

### **CSS Styling Choices**
The app follows a clean and modern UI with the following design choices:
- **Background Color**: Light Sky Blue (`#87CEEB`) – A friendly and inviting color.
- **Secondary Color**: Soft Lavender (`#D8BFD8`) – Used for headers and subtle contrast.
- **Accent Color**: Vibrant Coral (`#FF6F61`) – Used for buttons and interactive elements.

```

## Running WatchParty
1. **Clone the repository**
   ```sh
   git clone https://github.com/yourusername/watchparty.git
   cd watchparty
   ```
2. **Install dependencies**
   ```sh
   go mod tidy
   ```
3. **Run the server**
   ```sh
   go run cmd/server/main.go
   ```
4. **Open the app in a browser**
   - Visit: [http://localhost:8080](http://localhost:8080)
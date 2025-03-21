# Chat System Using WebSocket

## Guide to Clone the App

To clone this application, follow these steps:

1. Ensure you have Git installed on your machine. You can download it from [git-scm.com](https://git-scm.com/).

2. Open your terminal or command prompt.

3. Navigate to the directory where you want to clone the repository.

4. Run the following command to clone the repository:

   ```bash
   git clone https://github.com/lcv-back/chat-system.git
   ```

5. Change into the project directory:

   ```bash
   cd chat-system
   ```

### Frontend Setup (ReactJS)

1. Change into the frontend project directory:

   ```bash
   cd frontend
   ```

2. Install the necessary dependencies:

   ```bash
   npm install
   ```

3. Start the frontend application:

   ```bash
   npm start
   ```

### Backend Setup (Golang)

1. Navigate to the backend folder:

   ```bash
   cd ../backend
   ```

2. Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

3. Install the necessary Go dependencies:

   ```bash
   go mod tidy
   ```

4. Start the backend server:

   ```bash
   go run cmd/server/main.go
   ```

Now you should be able to access the chat system using WebSocket!

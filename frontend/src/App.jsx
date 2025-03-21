import React, { useState, useEffect } from 'react';
import './App.css';
import ChatComponent from './Components/ChatComponent';

// URL WebSocket server của bạn
const SOCKET_URL = 'ws://192.168.1.41:1234/ws';

function App() {
    const [socket, setSocket] = useState(null);

    useEffect(() => {
        // Tạo kết nối WebSocket khi ứng dụng được mount
        const socketConnection = new WebSocket(SOCKET_URL);

        // Đặt sự kiện khi kết nối WebSocket mở
        socketConnection.onopen = () => {
            console.log('Connected to WebSocket server');
        };

        // Đặt sự kiện khi WebSocket đóng
        socketConnection.onclose = () => {
            console.log('Disconnected from WebSocket server');
        };

        // Lưu socket vào state
        setSocket(socketConnection);

        // Đóng kết nối WebSocket khi component unmount
        return () => {
            socketConnection.close();
        };
    }, []);

    return (
        <div className="App">
            <h1>Real - Time Chat</h1>
            {socket && <ChatComponent socket={socket} />}
        </div>
    );
}

export default App;
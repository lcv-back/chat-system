import React, { useState, useEffect } from 'react';

function ChatComponent({ socket, userName }) {
  const [message, setMessage] = useState('');
  const [messages, setMessages] = useState([]);
  const [typing, setTyping] = useState(false); // Trạng thái đang nhập
  const [typingUser, setTypingUser] = useState(null); // Người dùng đang nhập

  // Lắng nghe các tin nhắn từ server qua WebSocket
  useEffect(() => {
    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.type === 'message') {
        setMessages((prevMessages) => [...prevMessages, { user: data.user, message: data.message }]);
      } else if (data.type === 'typing') {
        setTypingUser(data.user);
      }
    };

    // Cleanup on unmount
    return () => {
      socket.onmessage = null;
    };
  }, [socket]);

  // Hàm gửi tin nhắn
  const sendMessage = () => {
    if (message.trim()) {
      socket.send(JSON.stringify({ type: 'message', message, user: userName }));
      setMessages((prevMessages) => [...prevMessages, { user: 'Me', message }]);
      setMessage(''); // Xóa input sau khi gửi
      setTyping(false); // Đặt trạng thái "typing" về false khi gửi tin nhắn
    }
  };

  // Hàm xử lý sự kiện khi nhấn Enter
  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      sendMessage();
    } else {
      // Gửi thông báo "typing" khi người dùng nhập
      if (!typing) {
        setTyping(true);
        socket.send(JSON.stringify({ type: 'typing', user: userName }));
      }
    }
  };

  return (
    <div>
      <div id="messages" style={{ border: '1px solid #ccc', padding: '10px', height: '300px', overflowY: 'scroll' }}>
        {messages.map((msg, index) => (
          <div key={index} style={{ textAlign: msg.user === 'Me' ? 'right' : 'left' }}>
            <div style={{ display: 'inline-block', padding: '5px', background: msg.user === 'Me' ? '#4CAF50' : '#ccc', borderRadius: '20px', color: '#fff' }}>
              {msg.user}: {msg.message}
            </div>
          </div>
        ))}
      </div>
      {typingUser && typingUser !== userName && (
        <div>{typingUser} is typing...</div> // Hiển thị người đang nhập
      )}
      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        onKeyDown={handleKeyPress} // Thêm sự kiện onKeyDown
        placeholder="Type a message..."
        style={{ width: '80%' }}
      />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
}

export default ChatComponent;

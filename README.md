## Amarath

Math Assistant AI
A lightweight AI chat application designed to help users solve and understand mathematical problems through natural language interaction.
The backend is built using __ElysiaJS__, providing a high-performance and minimal API layer. The system integrates Ollama to run the __Qwen3 1.7B__ language model locally, enabling private and efficient inference without relying on external AI APIs.
To support real-time interaction, the application uses __Redis Pub/Sub__ for streaming AI responses to clients while messages are being generated. Chat messages and conversation metadata are persisted in __PostgreSQL__, allowing the assistant to maintain conversational context and history.

- Real-time AI response streaming using Redis Pub/Sub
- Chatroom-based conversation management
- Context-aware responses using conversation history
- Web-assisted context retrieval for enhanced answers
- Secure chatroom access with user authorization

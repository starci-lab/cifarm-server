# CiFarm Server  

## Related Repositories  
Here are the related repositories for different components of the CiFarm ecosystem:

- **Game (Client - Unity)**: [CiFarm Client](https://github.com/starci-lab/cifarm-client)
- **Game Server (Backend)**: [CiFarm Server](https://github.com/starci-lab/cifarm-server)
- **Wallet Integration for Game**: [CiWallet](https://github.com/starci-lab/ciwallet)
- **Telegram Bot**: [CiWallet Bots](https://github.com/starci-lab/ciwallet-bots)
- **Blockchain Backend**: [CiFarm Periphery](https://github.com/starci-lab/cifarm-periphery)

## Overview  
CiFarm Server is the backbone of the CiFarm game, designed to handle all client requests seamlessly. Built using **Go** and the **Nakama** framework, it provides a robust communication layer leveraging both **gRPC** and **WebSocket** protocols to deliver a real-time and responsive gameplay experience.  

## Key Features  

- **Multi-Protocol Communication**  
  - Supports both **gRPC** for efficient request handling and **WebSocket** for real-time updates, ensuring a smooth and immersive gaming experience.  

- **Scalable Architecture**  
  - Optimized to handle a high number of concurrent users, making it suitable for games with growing player bases.  

- **Powered by Nakama**  
  - Integrates with the Nakama framework to offer essential game server functionalities, including:  
    - User authentication and management  
    - Real-time multiplayer support  
    - Social features (friends, leaderboards, etc.)  

## Technologies  

- **Go**: A powerful, efficient programming language for high-performance server-side development.  
- **Nakama**: A scalable, open-source game server framework for multiplayer and social games.  

## Notes  
This is an off-chain game server designed to handle game-related interactions and user management outside the blockchain network.  

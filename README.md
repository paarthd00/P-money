# P-money Blockchain

## Overview

P-money is a blockchain-based cryptocurrency implementation with a basic structure for managing transactions and wallets. It features a straightforward blockchain architecture, blocks, and transactions, along with a wallet component for handling public and private keys.

## Project Structure

- **main.go:** The main entry point of the P-money program, initializing the blockchain, creating a wallet, and adding blocks to the blockchain with transaction data.

- **blockchain/blockchain.go:** Houses the blockchain-related logic, including definitions for blocks and the blockchain structure, block creation, hashing, and functions for managing the blockchain.

- **wallet/wallet.go:** Implements a basic wallet with functions for generating public and private key pairs, creating Bitcoin-like addresses, and initializing a new wallet.

- **client/client.go:** Contains the client-side code responsible for sending transactions to the P-money blockchain server.

- **server/server.go:** The server-side code for handling incoming transactions, updating the blockchain, and managing connections.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/paarthd00/p-money.git
   ```

2. Navigate to the project directory:

   ```bash
   cd p-money
   ```

3. Run the server:

   ```bash
   go run server/server.go
   ```

4. Open a new terminal window and run the client with arguments for the username, recipient name, amount, and coin type:

   ```bash
   go run client/client.go <username> <recipient> <amount> <coin>
   ```

## Dependencies

- This project uses the Go programming language. Ensure you have Go installed on your machine.

## Customization

- Tailor the P-money blockchain and wallet by modifying the `blockchain` and `wallet` packages. Explore additional features, tweak the wallet, or enhance the block creation process.

## License

This P-money implementation is open-source and released under the [MIT License](LICENSE).

## Acknowledgments

- Inspired by the principles of blockchain and cryptocurrency technologies.

---

### Project Structure Tree

```plaintext
.
├── blockchain
│   └── block.go
├── client
│   ├── client
│   └── client.go
├── go.mod
├── main
├── main.go
├── README.md
├── server
│   ├── server
│   └── server.go
└── wallet
    └── wallet.go
```

Feel free to adapt the README to include any additional details or instructions specific to your project. Good luck with your P-money blockchain implementation! 🚀💰

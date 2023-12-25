# P-money Blockchain

## Overview

P-money is a simple implementation of a blockchain-based cryptocurrency. It includes a basic blockchain structure with blocks and transactions, as well as a wallet for managing public and private keys.

## Files

- **main.go:** The main entry point for the P-money program. It initializes a new blockchain, creates a wallet, and adds blocks to the blockchain with transaction data.

- **blockchain/blockchain.go:** Contains the blockchain-related logic, including block and blockchain struct definitions, block creation, hashing, and blockchain management functions.

- **wallet/wallet.go:** Implements a basic wallet with functions for generating public and private key pairs, creating a Bitcoin-like address, and initializing a new wallet.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/paarthd00/p-money.git
   ```

2. Navigate to the project directory:

   ```bash
   cd p-money
   ```

3. Run the main program:

   ```bash
   go run main.go
   ```

   This will create a new blockchain, generate a wallet, and add blocks to the blockchain with transaction data.

## Dependencies

- This project uses the Go programming language. Make sure you have Go installed on your machine.

## Customization

- You can customize the P-money blockchain and wallet by modifying the `blockchain` and `wallet` packages. For example, you may want to add more features to the wallet or modify the block creation process.

## License

This P-money implementation is open-source and released under the [MIT License](LICENSE).

## Acknowledgments

- This project is inspired by the principles of blockchain and cryptocurrency technologies.

---

Feel free to tailor the README file to include additional details, instructions, or information specific to your project. Adjust the license information, acknowledgments, and other sections as needed. Good luck with your P-money implementation! 🚀💰

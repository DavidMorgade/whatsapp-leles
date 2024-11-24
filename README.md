# WhatsApp Leles

WhatsApp Leles is a Go-based application that interacts with WhatsApp using the WhatsMeow library. It includes functionalities for database management, event handling, and more.

## Features

- Generates a QR code for WhatsApp login
- Retrieves cryptocurrency prices
- Fetches weather information
- Implements AI with OpenAI API
- Personal assistants for various tasks

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/whatsapp-leles.git
    cd whatsapp-leles
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Build the project:
    ```sh
    go build -o whatsapp-leles main.go
    ```

## Usage

1. Set the required API keys in your environment:
    ```sh
    export API_KEY=your_api_key
    export RAPID_API_KEY=your_rapid_api_key
    export OPEN_AI_KEY=your_open_ai_key
    export ASSISTANT_LELE=your_assistant_lele_key
    export ASSISTANT_JAYN=your_assistant_jayn_key
    export ASSISTANT_TOTI=your_assistant_toti_key
    export ASSISTANT_MANU=your_assistant_manu_key
    export ASSISTANT_ROXA=your_assistant_roxa_key
    export ASSISTANT_MARIA=your_assistant_maria_key
    export COIN_API_KEY=your_coin_api_key
    ```

2. Run the application:
    ```sh
    ./whatsapp-leles
    ```

3. Scan the generated QR code with your WhatsApp mobile app to log in.

## Endpoints

- **Assistant Routes**: Handles mentions and commands for personal assistants.
- **Regular IA Routes**: Processes general AI-related commands.
- **Audio Routes**: Manages audio-related commands.
- **Image Routes**: Handles image-related commands.
- **Help Routes**: Provides help and support commands.
- **Crypto Routes**: Fetches cryptocurrency prices.
- **Weather Routes**: Retrieves weather information.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

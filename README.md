# SHA1 Password Cracker

This project is a simple SHA1 password cracker written in Go. It attempts to crack SHA1 hashed passwords by comparing them against a list of known passwords and salts.

## Features

- Crack SHA1 hashed passwords
- Support for salted passwords
- Easy to use

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Sumit0x00/SHA1-Password-Cracker.git
    cd SHA1-Password-Cracker
    ```

2. Build the project:
    ```sh
    go build main.go
    ```

3. Run
    ```
    ./main
    ```

4. Testing
    ```
    go test
    ```
    
## Usage

To crack a SHA1 hashed password, you can use the `CrackSHA1Hash` function from the `pass` package.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

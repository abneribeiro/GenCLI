# GenCli

GenCli is a versatile Command-Line Interface (CLI) tool designed to streamline project generation. Quickly create projects with customized names, types, and explore advanced features for efficient and flexible development workflows.

## Features

- **Simple CLI Interface**: Easy-to-use command-line interface for swift project creation.
- **Customization**: Specify project names and types with flexibility.
- **Advanced Options**: Explore advanced options for tailoring projects to your needs.
- **Versatility**: Suitable for various languages and project structures.

## Getting Started

1. **Installation**:
   ```bash
   go get -u github.com/abneribeiro/gen
   ```

2. **Usage**:
   ```bash
   gen proj -n projectName.js
   ```

3. **Options**:
   - `-n, --name`: Specify the project name.

## Examples

- Create a Go project named "MyApp":
   ```bash
    gen proj -n MyApp.go
   ```

## Future Enhancements

- Project Templates
- Git Integration
- Additional Advanced Options

## Contributions

Contributions are welcome! Feel free to open issues and submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

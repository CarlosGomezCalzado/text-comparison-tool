# Text Comparison Tool

This project is a text comparison tool implemented in Go. It allows users to compare two strings and identify added, deleted, and modified content between them. The tool utilizes a rolling hash algorithm for efficient text search.

## Features

- **Text Comparison**: Compare two strings and identify added, deleted, and modified content.
- **Efficient Search**: Utilizes a rolling hash algorithm for efficient text search.
- **User-Friendly Interface**: Simple command-line interface for easy interaction.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/CarlosGomezCalzado/text-comparison-tool.git
    ```

2. Navigate to the project directory:

    ```bash
    cd text-comparison-tool
    ```

3. Build the project:

    ```bash
    go build
    ```

## Usage

To use the text comparison tool, follow these steps:

1. Run the built executable:

    ```bash
    ./text-comparison-tool
    ```

2. Follow the prompts to enter the old text, updated text, and window size for comparison.

3. The tool will display the comparison result, highlighting added, deleted, and modified content between the two texts.

## Example

Here's an example of using the text comparison tool:

```bash
./text-comparison-tool
Enter the old text:
Lorem ipsum dolor sit amet.
Enter the updated text:
Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Enter the window size for comparison: 5
_______________________________________
Old text: Lorem ipsum dolor sit amet.
Updated text: Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Comparison result:
Start character: 23 [+++ , consectetur adipiscing elit.]
```

In this example, the tool identifies the added content `, consectetur adipiscing elit.` starting from character 23 in the updated text.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](https://www.mit.edu/~amini/LICENSE.md) file for details.

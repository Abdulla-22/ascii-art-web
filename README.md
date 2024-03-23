# Ascii-Art-Web (The code will be uploded soon)

## Description

Ascii-Art-Web is a web application that allows users to generate ASCII art using different banner styles. The application runs a server and provides a web GUI for users to input text and select a banner style. Upon submission, the server generates the ASCII art based on the user's input and selected banner style.

## Usage: How to Run

To run the Ascii-Art-Web application, follow these steps:

1. Clone the project repository.
2. Ensure you have Go installed on your machine.
3. Open a terminal and navigate to the project's root directory.
4. Run the following command to start the server and make sure you are in the `/App` directory:

   ```shell
   go run .
   ```

5. Open a web browser and visit `http://localhost:8080` to access the application.

## Implementation Details: Algorithm

1. The server listens for incoming HTTP requests on port 8080.
2. When a GET request is received at the root URL (`/`), the server responds with the main HTML page, which contains the text input, banner style selection, and submit button.
3. When a POST request is received at `/submit`, the server extracts the text and banner style values from the request.
4. The server then generates the ASCII art based on the provided text and banner style.
5. The generated ASCII art is included in the response, either as a separate page or appended to the main page, depending on the chosen implementation.
6. The appropriate HTTP status codes are returned based on the success or failure of the request.


## Currently Supported Banner names

- [x] standard banner
- [x] Thinkertoy banner
- [x] shadow banner

## Future updates 

- [ ] reverse text
- [ ] colored art
- [ ] and many more...

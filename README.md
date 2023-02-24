# 🤖 GiveAdvice - AI Advice Generator
![logo](https://cdn.jsdelivr.net/gh/gtoxlili/give-advice/frontend/src/assets/logo.png)

GiveAdvice is a project that aims to explore the capabilities of OpenAI's GPT-3 API in the field of AIGC (Artificial Intelligence Generated Content). The project is built using Go and React, and provides users with two main features:

1. **Advice Generation:** Users can input a description of something they want to create, and the program will generate a sentence or paragraph that describes what they're looking for. Alternatively, users can input an article, and then engage in a Q&A session with the program based on the content of the article.

2. **User Record Keeping:** User interactions are locally saved, so users can resume asking questions about a particular topic or view past consultation information at any time.

In addition, the project supports internationalization (i18n), which includes Chinese, Japanese, and English languages. The program also allows users to input their own token to avoid any traffic restrictions. Otherwise, the program controls the number of visits from a particular IP address within a specific time frame.

The project will also include the following features in the future:

1. **User Authentication:** To provide better access control, user authentication will be implemented based on JWT (JSON Web Tokens).

2. **Cloud Syncing:** User records will be synchronized to the cloud to allow for seamless interactions from multiple devices.

3. **Translation:** The program will use Deepl to localize the user input into the language of their choice.

## Demo
Check out the [demo](https://ai.gtio.work/) of GiveAdvice to see it in action!

## Installation
To install and run the program, please follow the instructions below:

1. Clone the repository using git clone https://github.com/gtoxlili/give-advice.git.
2. Install Go, React, and Node.js on your local machine.
3. Fill in `OPENAI_TOKEN` in Taskfile.yml with your OpenAI API key.
4. Run go run main.go to start the server.
5. In a new terminal window, navigate to the client directory and run `task default && ./dist/give-advice`.
6. Open a web browser and go to http://localhost:16806 to view the program.

## Contributing
If you're interested in contributing to GiveAdvice, please fork the repository and submit a pull request with your changes. We welcome all contributions!

## License
This project is licensed under the MIT License. See the LICENSE file for details.

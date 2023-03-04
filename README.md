# 🤖 GiveAdvice - AI Advice Generator
![logo](https://cdn.jsdelivr.net/gh/gtoxlili/advice-hub/frontend/src/assets/logo.png)

GiveAdvice is a project that aims to explore the capabilities of OpenAI's GPT-3 API in the area of AIGC (Artificial Intelligence Generated Content). The project is built in Go and React and provides users with two main features:

1. **Advice generation:** Users can enter a description of something they want to create, and the program will generate a sentence or paragraph that describes what they're looking for. Alternatively, users can enter an article and then engage in a Q&A session with the program based on the content of the article.

2. **User Record Keeping:** User interactions are stored locally, so users can resume questions on a particular topic at any time, or view past consultation information.

In addition, the project supports internationalization (i18n), which includes Chinese, Japanese, and English languages. The program also allows users to enter their own token to avoid any traffic restrictions. Otherwise, the program controls the number of visits from a certain IP address within a certain time frame.

The project will also include the following features in the future:

1. **User Authentication:** To provide better access control, user authentication based on JWT (JSON Web Tokens) will be implemented.

2. **Cloud Sync:** User records will be synchronized to the cloud to enable seamless interactions from multiple devices.

3. **Translation:** The program will use Deepl to localize user input into the language of their choice.

## Demo
Check out the [demo](https://ai.gtio.work/) of GiveAdvice to see it in action!

## Installation
To install and run the program, please follow the instructions below:

1. Clone the repository with git clone https://github.com/gtoxlili/advice-hub.git.
2. Install Go, Node.js, pnpm and task on your local machine.
3. Fill in `OPENAI_TOKEN` in Taskfile.yml with your OpenAI API key.
4. In a new terminal window, navigate to the Client directory and run `task default && ./dist/advice-hub`.
5. Open a web browser and go to http://localhost:7458 to view the program.

## Contributing
If you're interested in contributing to GiveAdvice, please fork the repository and submit a pull request with your changes. We welcome all contributions!

## License
This project is released under the GPL-3.0 License. See the LICENSE file for details.

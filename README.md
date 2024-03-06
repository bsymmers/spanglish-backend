## **Cognate Translator Backend**
[Frontend](https://github.com/bsymmers/cognate-translator-frontend)

This project serves as a very slim backend designed to work with the Cognate Translator frontend to facilitate the translation process. 

The server and API is implemented in Go with integration to both an ML model in Python used for language detection and Deepl’s API for machine translation.

The main server communicates with the Python ML model through [grpc](https://grpc.io/), with the go side acting as the “client” and the python side acting as the “server”.

## **Integration with the Deepl API**

This backend interacts directly with Deepl’s API, specifically for the purpose of translating input text. You can read more about the API [here](https://www.deepl.com/pro-api?cta=header-pro-api).

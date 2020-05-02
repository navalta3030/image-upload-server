# Summary ![Production](https://github.com/navalta3030/image-upload-server/workflows/Production/badge.svg?branch=master) ![Staging](https://github.com/navalta3030/image-upload-server/workflows/Staging/badge.svg)

# Routes
  - /upload (POST)
    - Requires Form
      - file (image)
      - meta_data (json)
        ```
        {
          "name": "mark",
          "email": "mark@gmail.com"
        }
        ```
      - returns (json) - all string
        ```
          {
            "data":[ 
              [<filename>, <sickness>, <accuracy>]
            ]
          }
        ```
    - Passes image to sanic server for prediction.
    - Upload image to digital ocean space.
    - Insert user information in database, bypass if already exist.
  - /getRecord (POST)
    - Requires Json
      - input
        ```
        {
          "name": "mark",
          "email": "mark@gmail.com"
        }
        ```
      - returns (json) - all string
        ```
        [
          {
            "link": <link_of_image>,
            "assistant_prediction": "Pneumonia",
            "doctor_prediction": ""
          }
        ]
        ```

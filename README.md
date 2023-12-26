# Take home Challenge - Full Stack Engineer

## Supported Languages
Challenge can be attempted in either of the following languages:
- Python
- Golang
- Javascript

## Description
A data entry clerk is entering in city metadata for cities across the planet and this data is being transpiled into batches represented as json files in the `tmp/` folder.

The engineering team needs to build an application to process this data and validate its authenticity. The only metadata that is validated to be correct during entry at the data clerk's office is the `city` property.

## Data
All files under `tmp/` folder needs to be processed.

The city metadata to verify authenticity is stored inside `cities.json`

## Acceptance Criteria
* Semantics Note: A "element" in this context is defined as a single object in each batch.
* If anything inside a element is determined to be inauthentic, it is considered "unsuccessfully validated".
* If everything inside a element is determined to be authentic, it is considered "successfully validated".
* Build a pretend serverless function to process data. This function should be production-ready.
* Function consumes city metadata in batches on a per file basis.
* Function returns a list of successfully validated elements.
* Function returns a list of unsuccessfully validated elements.
* Function returns a list of any elements/files that cannot be processed.


## Bonus Points:
* Containerize the solution in docker.

## App Description:
* My iteration of this take home challenge reads all provided files within the [data package](data/)
* A CitiesMap is created and stored in the client with the cities located inside [cites.json](data/cities.json)
* Tmp cities located [here](data/tmp), are compared and validated against the map
* Files for both valid elements, invalid elements and unprocessable files are then generated inside of the [results package](results)
* This app can be run using either `go run` or a docker container.
* Due to time constraints, there is only partial testing of some of the functions

## App Directions:
* Install Go
  - For Linux:
    - Run `sudo apt-get update
sudo apt-get install golang
`
  - For Mac (Homebrew)
    - Run `brew install go
`
* From the IONTRA_TH directory, run `go mod tidy` to install external dependencies
* From the root of the directory run `go run Iontra_th.go`
* Check the results package for the generated files

## Docker:
* Download Docker Desktop and make sure it is running
* From the root of the app run `docker-compose up --build`
* Check the results package for the generated files



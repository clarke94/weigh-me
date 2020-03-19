[![Netlify Status](https://api.netlify.com/api/v1/badges/39244a45-e184-486a-a115-fc808067a13c/deploy-status)](https://app.netlify.com/sites/weigh-me/deploys)

# WeighMe

This project was created for learning purposes. An app that allows a user to create an account and track their weight. Allowing a user to set goals, view friends and gain achievements.

**This project is still under development**

## Quick links

* [Netlify](https://app.netlify.com/sites/weigh-me/overview)
* [Live site](https://weigh-me.netlify.com/)
* [Trello board](https://trello.com/b/lHuaDmdf/weigh-me)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

[Node v10+](https://nodejs.org/en/) is needed, check your version with the below command in a terminal

```
node -v
```

[Angular CLI v9](https://www.npmjs.com/package/@angular/cli) is needed, check your version with the below command in a terminal

```
ng --version
```

[Golang](https://golang.org/doc/install) is needed, check your version with the below command in a terminal

```
go version
```

### Installing

A step by step series of examples that tell you how to get a development env running

Clone this repo and install the dependencies.

```
git clone git@github.com:clarke94/weigh-me.git
cd weigh-me
npm i
```

### Database

Get in touch to request access to the database and add the following file to the root of the project

```
MONGO_USER=<USERNAME>
MONGO_PASSWORD=<PASSWORD>
MONGO_CLUSTER=<CLUSTER>
```

### Development server

Run the server for the front-end

```
npm start
```

Run the server for the back-end

```
npm run server
```

## Running the tests

### Unit tests

Run `npm test` to execute the unit tests via [Karma](https://karma-runner.github.io).

### End-to-end tests

Run `npm run e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

### Linting

Run `npm run lint` to execute the code lints via [TSLint](https://palantir.github.io/tslint/) and [Stylelint](https://stylelint.io/).

## Deployment

Continuous deployment through [Netlify](https://www.netlify.com/) on the master branch with automatic deploy previews on every pull request.

## Contributing

1. Clone the repo locally
2. Contribute
3. Create a pull request

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/clarke94/weigh-me/releases).

## Authors

* **Liam Clarke** - [Github](https://github.com/clarke94)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

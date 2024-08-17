<a name="readme-top"></a>

<div align="center">
  <a href="https://github.com/harbanery/be-blanja-marketplace-app">
    <img src="./public/brandicon.png" alt="Logo Blanja" width="250">
  </a>

  <h1 align="center">Blanja</h1>

  <p align="center">
    Marketplace Implementation
    <br />
    <br />
    <a href="https://api-blanja-marketplace-app.onrender.com/" target="_blank">View Demo</a>
    ·
    <a href="https://blanja-website-project.netlify.app/" target="_blank">View Front-End Demo</a>
    ·
    <a href="https://github.com/harbanery/blanja-marketplace-app" target="_blank">View Front-End Repo</a>
  </p>
</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Setup Environment Variables](#setup-environment-variables)
  - [Running the Application](#running-the-application)
- [Usage](#usage)
  - [Features](#features)
  - [Project Structure](#project-structure)
  - [Documentation](#documentation)
- [Contributing](#contributing)
  - [Meet the Contributors](#meet-the-contributors)
  - [How to Contribute](#how-to-contribute)
- [License](#license)
- [Contacts](#contacts)
- [Acknowledgements](#acknowledgements)

<!-- <details>
  <summary>Table of Contents</summary>
  <ul>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#start-development-server">Start Development Server</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a>
    <ul>
        <li><a href="#project-structure">Project Structure</a></li>
        <li><a href="#documentation">Documentation</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ul>
</details> -->

## About The Project

**Blanja** is a web-based marketplace application that provides a platform for sellers and buyers to conduct online transactions for various products. This API allows developers to access and manage the various features and functions available on **Blanja**, including users, products, categories, carts, and orders.

### Built With

[![Go][Go]][Go-url]
[![Fiber][Fiber]][Fiber-url]
[![PostgreSQL][Postgre]][Postgre-url]
[![Cloudinary][Cloudinary]][Cloudinary-url]

<!-- [![Midtrans][Midtrans]][Midtrans-url] -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

Ensure you have the following installed on your local machine:

- [Golang][Go-url]
- [PostgreSQL for Local Database][Postgre-url]
- [Ngrok for Public Localhost](https://ngrok.com/)

### Installation

1. Clone Repo

   ```sh
   git clone https://github.com/harbanery/be-blanja-marketplace-app.git
   ```

2. Go to folder directory

   ```bash
   cd be-blanja-marketplace-app
   ```

3. Install Go packages

   ```sh
   go mod tidy
   ```

### Setup Environment Variables

1. Create a `.env` or `.env.local` file in your local root directory.

2. Add the following variables to the `.env` or `.env.local` file:

   ```sh
      # DATABASE
      URL=YOUR_POSTGRESQL_CONNECTION_STRING_URL

      # JWT
      SECRETKEY=YOUR_SECRETKEY

      # CLOUDINARY
      CLOUDINARY_URL=YOUR_CLOUDINARY_LINK_ACCESS

      # SMTP
      BASE_URL=YOUR_BACKEND_HOST_URL
      SMTP_EMAIL_PASS=YOUR_EMAIL_TOKEN
      SMTP_EMAIL_USER=YOUR_EMAIL_SENDER

      # MIDTRANS
      CLIENT_KEY=YOUR_MIDTRANS_CLIENT_KEY
      SERVER_KEY=YOUR_MIDTRANS_CLIENT_KEY

      # REDIRECT_URL
      REDIRECT_URL=YOUR_FRONTEND_HOST_URL
   ```

   For more value information with format, you can go to [`.env.example`](https://github.com/harbanery/be-blanja-marketplace-app/blob/main/.env.example)

### Running the Application

1. To start the development server, you can start with this:

   ```sh
   go run main.go
   ```

   Or this:

   ```sh
   air
   ```

2. After that, you can run browser with this url:

   ```sh
   http://localhost:3000
   ```

> The local server always start on port `3000` by default.

_If you want to interact with the endpoints, go to [Documentation](#documentation)._

## Usage

### Features

- **Registration (and Reset Password) Authentication** with [SMTP by Gomail](https://github.com/go-gomail/gomail)
- **Upload multiple images** with **Cloudinary**
- **Support for multiple images, colors, and sizes** for products, including **filtering**
- **ORM support for PostgreSQL** using [GORM](https://gorm.io/)
- **XSS Middleware and HTML Sanitizer** using [Bluemonday](https://github.com/microcosm-cc/bluemonday)
- **Model Struct Validation** using [Validator](https://github.com/go-playground/validator)
- **Payment gateway** using [Midtrans](https://github.com/veritrans/go-midtrans)
- **Token authentication and authorization** using [JWT](https://github.com/golang-jwt/jwt)

### Project Structure

Sure, here's the project structure in shell mode with descriptions:

```sh
be-blanja-marketplace-app/
├── public/                           # Static or public assets
│   └── brandicon.png                 # Blanja icon (purposely to this README.md)
├── src/                              # Main application source code
│   ├── configs/                      # Configuration files
│   │   └── db.go                     # Database configuration
│   ├── controllers/                  # Controllers handling HTTP requests
│   │   ├── AddressController.go      # Address-related operations
│   │   ├── CartController.go         # Cart-related operations
│   │   ├── CartProductController.go  # Cart product-related operations
│   │   ├── CategoryController.go     # Category-related operations
│   │   ├── CheckoutController.go     # Checkout process
│   │   ├── ColorController.go        # Color-related operations
│   │   ├── CustomerController.go     # Customer-related operations
│   │   ├── OrderController.go        # Order-related operations
│   │   ├── ProductController.go      # Product-related operations
│   │   ├── SellerController.go       # Seller-related operations
│   │   ├── SizeController.go         # Size-related operations
│   │   ├── UploadController.go       # File upload operations
│   │   └── UserController.go         # User-related operations
│   ├── helpers/                      # Helper functions and utilities
│   │   ├── jwt.go                    # JWT handling
│   │   ├── migration.go              # Database migrations
│   │   ├── params.go                 # Query parameter handling
│   │   ├── transaction.go            # Transaction utilities
│   │   ├── url.go                    # URL utilities
│   │   └── validation.go             # Data validation
│   ├── middlewares/                  # Middleware functions
│   │   ├── jwt.go                    # JWT middleware
│   │   └── xssClean.go               # XSS protection middleware
│   ├── models/                       # Data models
│   │   ├── Address.go                # Address model
│   │   ├── Cart.go                   # Cart model
│   │   ├── CartProduct.go            # Model for Cart to Product
│   │   ├── Category.go               # Category model
│   │   ├── Checkout.go               # Checkout model
│   │   ├── Color.go                  # Color model for Product
│   │   ├── Customer.go               # Customer model
│   │   ├── Image.go                  # Image model for Product
│   │   ├── Order.go                  # Order model
│   │   ├── Product.go                # Product model
│   │   ├── Seller.go                 # Seller model
│   │   ├── Size.go                   # Size model for Product
│   │   └── User.go                   # User model
│   ├── routes/                       # Route definitions
│   │   └── main.go                   # Main routing file
│   └── utils/                        # Utility functions
│       ├── cloudinary.go             # Cloudinary integration
│       ├── midtrans.go               # Midtrans integration
│       └── smtp.go                   # SMTP integration
├── tmp/                              # Temporary files
│   ├── build-errors.log              # Build errors log
│   ├── main.exe                      # Compiled main executable
│   └── ngrok.exe                     # Ngrok executable
├── .air.toml                         # Air live reload configuration
├── .env                              # Environment variables
├── .gitignore                        # Git ignore rules
├── .golangci.yml                     # GolangCI-Lint configuration
├── go.mod                            # Go module dependencies
├── go.sum                            # Go module checksums
├── LICENSE                           # License file
├── main.go                           # Main application entry point
└── README.md                         # Project documentation
```

### Documentation

Access the API documentation for the **Blanja** project. Use this documentation to test endpoints and understand the structure and functionality of the available API`s in this project.

[![Blanja API Postman Collection](https://run.pstmn.io/button.svg)](https://documenter.getpostman.com/view/36623326/2sA3duEsiT#e0a6ecab-217d-4075-854e-93131d42d842)

Or you can just import the collection manually:

[`v2`](https://github.com/harbanery/be-blanja-marketplace-app/blob/main/docs/blanja-api-codecraft.postman_collection_v2.json) [`v2.1`](https://github.com/harbanery/be-blanja-marketplace-app/blob/main/docs/blanja-api-codecraft.postman_collection_v2.1.json)

And other documentations:

- [`Database Diagram`](https://github.com/harbanery/be-blanja-marketplace-app/blob/main/docs/blanja-api-codecraft.diagram.png)

- [`Database Query`](https://github.com/harbanery/be-blanja-marketplace-app/blob/main/docs/blanja-api-codecraft.query.sql)

## Contributing

### Meet The Contributors

<table align="center">
  <tr>
    <th colspan="4">Codecraft</th>
  </tr>
  <tr align="center">
    <th>Project Manager / Front-End Developer</th>
    <th>Front-End Developer</th>
    <th>Fullstack Developer</th>
    <th>Back-End Developer</th>
  </tr>
  <tr align="center">
    <td><img width="200" src="https://avatars.githubusercontent.com/u/80629118?v=4"/></td>
    <td><img width="200" src="https://avatars.githubusercontent.com/u/74017000?v=4"/></td>
    <td><img width="200" src="https://avatars.githubusercontent.com/u/137192782?v=4"/></td>
    <td><img width="200" src="https://avatars.githubusercontent.com/u/89146375?v=4"/></td>
  </tr>
  <tr align="center">
    <td><a href="https://github.com/echestratus" target="_blank"><b>Farhan Nur Hakim</b></a></td>
    <td><a href="https://github.com/wafash08" target="_blank"><b>M. Wafa Saeful Haq</b></a></td>
    <td><a href="https://github.com/nizuma666" target="_blank"><b>Syaifulloh Ismail</b></a></td>
    <td><a href="https://github.com/harbanery" target="_blank"><b>Raihan Yusuf</b></a></td>
  </tr>
</table>

### Contribution List

1. [**Raihan Yusuf**](https://github.com/harbanery)

   - Completed all APIs except for Cart and Checkout

   - Responsible for features, including payment gateway, email verification, and product filtering

   - Relational database design

   - Postman API documentation

2. [**Syaifulloh Ismail**](https://github.com/nizuma666)

   - Completed the Cart and Checkout APIs, including the ability to add and reduce the quantity of products in the cart

### How to Contribute

Contributing project to github is pretty straight forward.

1. **Fork the Repository**: Click the "Fork" button at the top-right corner of this page to create your own copy of the repository.
2. **Installation**: Do the installation process right [here](#installation).
3. **Create a new branch**: Create a new branch to work on your changes.
   ```sh
   git branch -M feature/your-feature-name
   ```
4. **Make changes**: Implement your changes or new features in your branch.
5. **Commit your changes**: After making your changes, commit them with a descriptive message.
   ```sh
   git add .
   git commit -m "Add description of your changes"
   ```
6. **Push to GitHub**: Push your changes to your forked repository.
   ```sh
   git push origin feature/your-feature-name
   ```
7. **Create Pull Request**: Go to the original repository and open a pull request to merge your changes.
   - Navigate to your fork on GitHub.
   - Click on the "New Pull Request" button.
   - Select the base repository and branch you want to merge into.
   - Provide a detailed description of your changes and click "Create Pull Request".

## License

Distributed under the MIT License. See [`LICENSE`](https://github.com/harbanery/be-blanja-marketplace-app/blob/main/LICENSE) for more information.

## Contacts

If you have any questions or inquiries regarding this back-end project, feel free to contact at:

[![Raihan Yusuf's Email](https://img.shields.io/badge/Raihan%20Yusuf-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:ryusuf05@gmail.com)
[![Raihan Yusuf's LinkedIn](https://img.shields.io/badge/Raihan%20Yusuf-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/raihan-yusuf/)

[![Syaifulloh Ismail's Email](https://img.shields.io/badge/Syaifulloh%20Ismail-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:syaifullohismail123@gmail.com)
[![Syaifulloh Ismail's LinkedIn](https://img.shields.io/badge/Syaifulloh%20Ismail-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/syaifulloh-ismail/)

## Acknowledgements

Feel free to check it out:

- [Img Shields](https://shields.io)
- [Choose an Open Source License](https://choosealicense.com/)
- [GitHub Pages](https://pages.github.com/)

<!-- MARKDOWN LINKS & IMAGES -->

[Postgre]: https://img.shields.io/badge/postgresql-336791?style=for-the-badge&logo=postgresql&logoColor=white
[Postgre-url]: https://www.postgresql.org/
[Go]: https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://golang.org/
[Fiber]: https://img.shields.io/badge/fiber-057A7A?style=for-the-badge&logo=go&logoColor=white
[Fiber-url]: https://gofiber.io/
[Cloudinary]: https://img.shields.io/badge/cloudinary-3448C5?style=for-the-badge&logo=cloudinary&logoColor=white
[Cloudinary-url]: https://cloudinary.com/

<!-- [GORM]: https://img.shields.io/badge/gorm-000000?style=for-the-badge&logo=go&logoColor=white
[GORM-url]: https://gorm.io/ -->
<!-- [Midtrans]: https://img.shields.io/badge/midtrans-0C9CB4?style=for-the-badge&logo=go&logoColor=white
[Midtrans-url]: https://github.com/veritrans/go-midtrans -->
<!-- ![Go](https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/fiber-057A7A?style=for-the-badge&logo=fiber&logoColor=white)
![GORM](https://img.shields.io/badge/gorm-000000?style=for-the-badge&logo=go&logoColor=white)
![Cloudinary](https://img.shields.io/badge/cloudinary-3448C5?style=for-the-badge&logo=cloudinary&logoColor=white)
![Midtrans](https://img.shields.io/badge/midtrans-0C9CB4?style=for-the-badge&logo=midtrans&logoColor=white) -->

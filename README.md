# ⭐ Welcome to rhyfil!
- A point-of-sale web app built for tablet use.

## 💡 Motivation
- Working in the service industry for the last ten years and coffee shops for most of them, I've had a lot of experience working with different POS systems, and wanted to make a more streamlined "product to modifiers to transaction" organization that is focused on the improving the workflow of a busy barista without sacrificing quality for the customer. A lot of POS sale applications are built with an over abundance of features that clog up the cycle because they are built to be universally designed for other industries. Rhyfil is designed from the ground up with the intention to simplify the process for the user by having simple, powerful actions available and avoiding unnecesary tools and required actions.

## 🔦 Tech Stack
- **Backend** Go 1.23, net/http library
- **Database** PostgreSQL with goose migrations
- **DB Query Generation** sqlc
- **Frontend:** Vanilla JavaScript, HTML, CSS
- **Drivers:** lib/pq, google/uuid

## 💿 Usage
- Dynamic menu loading from database
- Modifier options tailored to each product (e.g. size, milk, syrups)
- Cart system with running total and per-item modifier tracking
- Order persistence - transactions saved to PostgreSQL database
- Login screen with local storage session management

## 🚀 Quick Start
1. Clone the repo
2. Create a PostgresSQL database and set the connection URL in your config
3. Run migrations: `goose up`
4. Seed products: `./rhyfil newproduct "Latte" 4.50 100`
5. register users: `./rhyfil register "John"`
6. Optional: Add Modifiers and Modifier Groups with other implemented commands: `addmodifiergroup`, `addgroupoption`
7. Start the server: `./rhyfil serve`
8. Visit `http://localhost:8080/login.html`

## 🛠️ Planned Features
- Password authorization and session/idling expiration
- In app Admin page to manage products and add modifiers
- QOL features such as delete from cart
- Printable receipts. exportable order history

## 🤝 Contributing
### Clone the repo
```bash
go clone github.com/Rhyster42/rhyfil
cd rhyfil
```
### Build the compiled library
```bash
go build
```
### Submit a pull request
If you would like to contribute, please fork the repository and make a pull request to the `main` branch.

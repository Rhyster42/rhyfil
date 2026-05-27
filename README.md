# ⭐ Welcome to rhyfil!
- A point-of-sale web app build for tablet use.
- Designed as a learning project to demonstrate full stack development with Go and PostgreSQL.

## 🔦 Tech Stack
- **Backend** Go 1.23, net/http library
- **Database** PostgreSQL with goose migrations
- **DB Query Generation** sqlc
- **Frontend:** Vanilla JavaScript, HTML, CSS
- **Drivers:** lib/pq, google/uuid

## 💿 Features
- Dynamic menu loading from database
- Modifier options tailored to each product (e.g. size, milk, syrups)
- Cart system with running total and per-item modifier tracking
- Order persistence - transactions saved to PostgreSQL database
- Login screen with local storage session management

## 🚀 How to Run
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

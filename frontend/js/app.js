const API_BASE = 'http://localhost:8080'; // Your Go server
let cart = [];

// Load products on page load
async function loadProducts() {
    try {
        const response = await fetch(`${API_BASE}/products`);
        const products = await response.json();
        displayProducts(products);
    } catch (error) {
        console.error('Failed to load products:', error);
    }
}

function displayProducts(products) {
    const grid = document.getElementById('product-grid');
    grid.innerHTML = products.map(product => `
        <div class="product-card" onclick="addToCart(${product.id}, '${product.name}', ${product.base_price})">
            <div class="product-name">${product.name}</div>
            <div class="product-price">$${product.base_price.toFixed(2)}</div>
        </div>
    `).join('');
}

function addToCart(id, name, price) {
    cart.push({ id, name, price, quantity: 1 });
    updateCartDisplay();
}

function updateCartDisplay() {
    const cartDiv = document.getElementById('cart-items');
    const total = cart.reduce((sum, item) => sum + item.price, 0);
    
    cartDiv.innerHTML = cart.map((item, index) => `
        <div class="cart-item">
            ${item.name} - $${item.price.toFixed(2)}
            <button onclick="removeFromCart(${index})">×</button>
        </div>
    `).join('');
    
    document.getElementById('total').textContent = total.toFixed(2);
}

function removeFromCart(index) {
    cart.splice(index, 1);
    updateCartDisplay();
}

async function checkout() {
    if (cart.length === 0) return;
    
    const total = cart.reduce((sum, item) => sum + item.price, 0);
    
    try {
        const response = await fetch(`${API_BASE}/transactions`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                items: cart,
                total_amount: total,
                payment_method: 'cash' // Could add payment method selector
            })
        });
        
        if (response.ok) {
            alert('Sale completed!');
            cart = [];
            updateCartDisplay();
        }
    } catch (error) {
        console.error('Checkout failed:', error);
        alert('Transaction failed');
    }
}

document.getElementById('checkout-btn').addEventListener('click', checkout);

// Load products when page loads
loadProducts();

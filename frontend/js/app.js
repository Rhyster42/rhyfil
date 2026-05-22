let pendingItem = null;
let pendingQuantity = 0;
let cart = [];

async function loadMenu() {
    try {
        const response = await fetch('http://localhost:8080/items');
        if (!response.ok) throw new Error("Network fetch response failed");

        const menuItems = await response.json();
        console.log(menuItems);

        const tableMenu = document.getElementById("menu-table");
        let menuString = ""

        menuItems.forEach(item => {
            menuString += `<tr data-id="${item.ID}" data-name="${item.Name}" data-price="${item.Price}"><td>${item.Name}</td></tr>`
        });

        tableMenu.innerHTML = menuString; // whats visible in menu table

        const rows = tableMenu.querySelectorAll("tr");
        rows.forEach(row => {
            row.addEventListener("click", () => {
                const productId = row.dataset.id;
                loadModifiers(productId);
                if (row.classList.contains('selected')) {
                    pendingQuantity++;
                    document.getElementById("add_button").textContent = `Add to Cart (${pendingQuantity})`;
                } else {
                    rows.forEach(r => r.classList.remove("selected"));
                    row.classList.add("selected");
                    pendingQuantity = 1;
                    document.getElementById("add_button").textContent = `Add to Cart (${pendingQuantity})`;
                }
                pendingItem = {
                    id: row.dataset.id,
                    name: row.dataset.name,
                    price: row.dataset.price,
                    modifiers: []
                };           
            });
        });

        const addToCartButton = document.getElementById('add_button');
        addToCartButton.addEventListener("click", addToCart);
        const checkOutButton = document.getElementById('check-out');
        checkOutButton.addEventListener("click", checkOut)

    } catch (error) {
        console.error('Fetch error:', error);
    }
}

async function loadModifiers(productId) {
    console.log("click worked: " + productId);
    const response = await fetch(`http://localhost:8080/modifiers?product_id=${productId}`);
    if (!response.ok) throw new Error("Network fetch response failed");

    const modifierGroups = await response.json();
    console.log(modifierGroups)

    const modifierTable = document.getElementById("modifier-table")
    let modGroupsString = "";

    for (let i = 0; i < modifierGroups.length; i++) {

        let modGroupId = modifierGroups[i].ID;
        const response = await fetch(`http://localhost:8080/modifier_options?group_id=${modGroupId}`);
        if (!response.ok) throw new Error("Network fetch response failed");
        let modOptions = await response.json();

        modGroupsString += `<tr><td>${modifierGroups[i].Name}</td></tr>`;
        let modOptionsString = "<tr>";
        for (let j = 0; j < modOptions.length; j++) {
            modOptionsString += `<td class=mod_options data-id="${modOptions[j].ID}" data-name="${modOptions[j].Name}" data-price="${modOptions[j].PriceAdjustment}">${modOptions[j].Name}</td>`
        };
        modGroupsString += modOptionsString + '</tr>';
    }

    modifierTable.innerHTML = modGroupsString; //whats visible in modifier table

    const modRows = modifierTable.querySelectorAll("td.mod_options")

    modRows.forEach(row => {
        row.addEventListener("click", () => {
            row.classList.toggle('selected');
            let modData = {
                id: row.dataset.id,
                name: row.dataset.name,
                price: row.dataset.price
            };
            if (row.classList.contains('selected')) {
                pendingItem.modifiers.push(modData)
            } else {
                pendingItem.modifiers = pendingItem.modifiers.filter(m => m.id !== modData.id)
            }
        })
    })
}

async function addToCart() {
    if (!pendingItem) return;

    const cartItems = document.getElementById("cart-table");
    let currentOrderString = cartItems.innerHTML;
    

    currentOrderString += `<tr><td item-id="${pendingItem.id}" class="cart-item">${pendingItem.name} X ${pendingQuantity}-${pendingItem.price}</td></tr>`
    let itemTotal = Number(pendingItem.price);

    pendingItem.modifiers.forEach(modifier => {
        currentOrderString += `<tr><td class="cart-mod">${modifier.name}-${modifier.price}</td></tr>`;
        itemTotal += Number(modifier.price);
    });

    currentOrderString += `<tr><td class="item-total">${itemTotal}</td></tr>`;

    let itemData = {
        id: pendingItem.id,
        name: pendingItem.name,
        price: pendingItem.price,
        modifiers: pendingItem.modifiers,
        itemTotal: itemTotal,
        quantity: pendingQuantity
    }
    cart.push(itemData)
    

    calculateTotal()

    cartItems.innerHTML = currentOrderString;

    pendingQuantity = 0
    pendingItem = null;
    document.getElementById("add_button").textContent = "Add to Cart";
};

function calculateTotal() {
    let cartTotal = document.getElementById('cart-total')  
    
    let total = cart.reduce((total, item) => (total + (item.itemTotal * item.quantity)), 0)
   
    const formattedTotal = new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(total);
    cartTotal.innerHTML = formattedTotal;
}

async function checkOut() {
    
    let jsonCartItems = []

    cart.forEach(item => {
        let jsonMods = []
        item.modifiers.forEach(modifier => {
            jsonMods.push({
                modId: modifier.id,
                priceAdjustment: parseFloat(modifier.price)
            });
        })
        jsonCartItems.push({
            productId: item.id,
            itemTotal: parseFloat(item.itemTotal),
            itemMods: jsonMods,
            quantity: item.quantity
        });
    })

    let orderTotal = cart.reduce((sum, item) => sum + (item.itemTotal * item.quantity), 0);

    let jsonOrder = {
        total: orderTotal,
        items: jsonCartItems
    };

    const response = await fetch('http://localhost:8080/checkout', {
        method: 'POST',
        headers: {
        'Content-Type': 'application/json'
        },
        body: JSON.stringify(jsonOrder)
    });
    return response.json(); 
}

loadMenu();

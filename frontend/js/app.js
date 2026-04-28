async function loadMenu() {
    try {
        const response = await fetch('http:localhost:8080/items');
        if (!response.ok) throw new Error("Network fetch response failed");

        const menuItems = await response.json();
        console.log(items);
    } catch (error) {
        console.error('Fetch error:', error);
    }
}


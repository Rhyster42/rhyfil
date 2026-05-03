async function loadMenu() {
    try {
        const response = await fetch('http://localhost:8080/items');
        if (!response.ok) throw new Error("Network fetch response failed");

        const menuItems = await response.json();
        console.log(menuItems);

        const tableMenu = document.getElementById("menu-table");
        let menuString = ""

        menuItems.forEach(item => {
            menuString += "<tr><td>" + item.name + "</td></tr>"
        });

        tableMenu.innerHTML = menuString;

    } catch (error) {
        console.error('Fetch error:', error);
    }

}

loadMenu();
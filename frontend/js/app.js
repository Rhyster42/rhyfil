async function loadMenu() {
    try {
        const response = await fetch('http://localhost:8080/items');
        if (!response.ok) throw new Error("Network fetch response failed");

        const menuItems = await response.json();
        console.log(menuItems);

        const tableMenu = document.getElementById("menu-table");
        let menuString = ""

        menuItems.forEach(item => {
            menuString += `<tr data-id="${item.ID}"><td>${item.Name}</td></tr>`
        });

        tableMenu.innerHTML = menuString;

        const rows = tableMenu.querySelectorAll("tr");
        rows.forEach(row => {
            row.addEventListener("click", () => {
                const productId = row.dataset.id;
                loadModifiers(productId);
            });
        });

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
            modOptionsString += `<td>${modOptions[j].Name}</td>`
        };
        modGroupsString += modOptionsString + '</tr>';
    }

    modifierTable.innerHTML = modGroupsString;
    
}

loadMenu();

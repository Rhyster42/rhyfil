
document.getElementById("login-form").addEventListener("submit", async (e) => {
    e.preventDefault();
    const name = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    

    const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 
            name: name,
            password: password,
         })
    });

    if (response.ok) {
        window.location.href = '/index.html';
        localStorage.setItem('user', name)
    } else {
        console.error('Login failed:', response.statusText);
    };
})

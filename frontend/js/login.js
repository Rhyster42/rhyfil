
document.getElementById("login-form").addEventListener("submit", async (e) => {
    e.preventDefault();
    const name = document.getElementById("username").value;
    

    const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name: name })
    });

    if (response.ok) {
        window.location.href = '/index.html';
    } else {
        console.error('Login failed:', response.statusText);
    };
    localStorage.setItem('user', name);
    window.location.href = '/index.html';
})

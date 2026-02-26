const API = "http://localhost:8080";

// CREATE TICKET
async function createTicket() {
    const name = document.getElementById("name").value;
    const description = document.getElementById("desc").value;
    const category = document.getElementById("category").value;

    const res = await fetch(`${API}/api/create`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ name, description, category })
    });

    if (res.ok) {
        alert("Ticket created!");
        loadTickets();
    } else {
        alert("Error creating ticket");
    }
    document.getElementById("name").value = "";
    document.getElementById("desc").value = "";
}

// LOAD TICKETS
async function loadTickets() {
    const res = await fetch(`${API}/api/tickets`);
    const data = await res.json();

    const container = document.getElementById("tickets");
    container.innerHTML = "";

    if (data.length === 0) {
        container.innerHTML = "<p>No tickets yet 🚀</p>";
        return;
    }

    data.forEach(t => {
        container.innerHTML += `
            <div class="ticket">
                <b>${t.Name}</b> (${t.Category})<br>
                ${t.Description}<br><br>
                👤 Assigned: ${t.AssignedTo}<br>
                📌 Status: ${t.Status}<br>
                🕒 ${t.CreatedAt}
            </div>
        `;

        /* container.innerHTML += `
    <div class="ticket">
        <h3>${t.Name}</h3>
        <p><strong>Category:</strong> ${t.Category}</p>
        <p><strong>Description:</strong> ${t.Description}</p>
        <p><strong>Assigned:</strong> ${t.AssignedTo}</p>
        <p><strong>Status:</strong> ${t.Status}</p>
    </div>
`; */
    });
}

// Load on start
loadTickets();
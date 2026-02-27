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
        body: JSON.stringify({
            Name: name,
            Description: description,
            Category: category
        })
    });

    if (res.ok) {
        const data = await res.json();

        alert(`🎉 Ticket Created Successfully!

Ticket ID: ${data.TicketID}
Ticket Number: ${data.TicketNumber}`);

    } else {
        alert("❌ Error creating ticket");
    }

    document.getElementById("name").value = "";
    document.getElementById("desc").value = "";
    document.getElementById("category").value = "";
}


// TRACK TICKET BY ID
async function trackTicket() {
    const id = document.getElementById("trackId").value;
    const container = document.getElementById("trackResult");

    if (!id) {
        container.innerHTML = "<p>⚠️ Please enter ticket ID</p>";
        return;
    }

    const res = await fetch(`${API}/api/ticket/id/${id}`);

    container.innerHTML = "";

    if (!res.ok) {
        container.innerHTML = "<p>❌ Ticket not found</p>";
        return;
    }

    const t = await res.json();

    container.innerHTML = `
        <div class="ticket">
            <b>${t.Name}</b> (${t.Category})<br>
            ${t.Description}<br><br>
            👤 Assigned: ${t.AssignedTo}<br>
            📌 Status: ${t.Status}<br>
            🕒 ${t.CreatedAt}
        </div>
    `;
}
function clearTracking() {
    document.getElementById("trackId").value = "";
    document.getElementById("trackResult").innerHTML = "";
}
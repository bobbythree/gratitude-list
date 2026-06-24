import './style.css'

// html elements
const userInput = document.getElementById("user-input");
const form = document.getElementById("input-form")
const listUL = document.getElementById("list-ul")

window.addEventListener("load", () => {
  loadItems();
})

// get req
async function loadItems() {
  const response = await fetch("/api/list");
  const items = await response.json();

  listUL.innerHTML = "";

  items.forEach(item => {
    const li = document.createElement("li");
    li.textContent = item.item;

    const deleteBtn = document.createElement("button");
    deleteBtn.textContent = "delete";
    deleteBtn.style.padding = "0.25rem 0.5rem"
    deleteBtn.addEventListener("click", async () => {
      await fetch(`/api/list?id=${item.id}`, {
        method: "DELETE",
      });

      await loadItems();
    });

    li.appendChild(deleteBtn);

    listUL.appendChild(li);
  });
}

form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const value = userInput.value;
  console.log(value);

  // post req
  const url = "/api/list"
  const payload = { item: value }
  try {
    const response = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload)
    });

    const result = await response.json();
    console.log("form submitted successfully", result);
    await loadItems();
  } catch (err) {
    console.error("submission failed", err);
  }

  // clear input
  userInput.value = "";

})


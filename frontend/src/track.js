document.addEventListener('DOMContentLoaded', async () => {
  const list = document.getElementById('expense-list');

  const res = await fetch('/api/expenses');
  if (!res.ok) {
    list.innerText = 'Failed to load expenses';
    return;
  }

  const expenses = await res.json();

  if (expenses.length === 0) {
    list.innerText = 'No expenses found';
    return;
  }

  list.innerHTML = '';

  expenses.forEach((e) => {
    const div = document.createElement('div');
    div.className = 'expense-item';

    div.innerHTML = `
      <div>
        <strong>${e.category}</strong>
        <div class="desc">${e.description || ''}</div>
      </div>
      <div class="right">
        ₹${(e.amount / 100).toFixed(2)}
        <div class="date">${e.date}</div>
      </div>
    `;

    list.appendChild(div);
  });
});

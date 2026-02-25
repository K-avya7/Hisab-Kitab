document.addEventListener('DOMContentLoaded', () => {
  const modal = document.getElementById('modal');
  const addExpenseCard = document.getElementById('add-expense-card');
  const trackExpenseCard = document.getElementById('track-expense-card');
  const closeBtn = document.getElementById('close-modal');
  const form = document.getElementById('expense-form');

  addExpenseCard.addEventListener('click', () => {
    modal.classList.remove('hidden');
  });

    trackExpenseCard.addEventListener('click', () => {
    window.location.href = '/track.html';
    });


  closeBtn.addEventListener('click', () => {
    modal.classList.add('hidden');
  });

  modal.addEventListener('click', (e) => {
    if (e.target === modal) {
      modal.classList.add('hidden');
    }
  });

  form.addEventListener('submit', async (e) => {
    e.preventDefault();

    const payload = {
      amount: Math.round(Number(amount.value) * 100),
      category: category.value,
      description: description.value,
      date: date.value
    };

    const res = await fetch('/api/expenses', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });

    if (!res.ok) {
      alert('Failed to save expense');
      return;
    }

    modal.classList.add('hidden');
    form.reset();
  });
});

const create = (data) => () =>
  fetch("/api", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data()),
  });

const update = (id, data) => () =>
  fetch(`/api/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data()),
  });

(() => {
  const addBtn = document.querySelector("#add");
  const addTitle = document.querySelector("#title");
  const checkboxes = document.querySelectorAll('input[type="checkbox"]');

  addBtn.addEventListener(
    "click",
    create(() => ({ title: addTitle.value }))
  );

  checkboxes.forEach((checkbox) => {
    checkbox.addEventListener(
      "change",
      update(checkbox.id, () => ({ done: checkbox.checked }))
    );
  });
})();

const tableBody = document.querySelector(".table-body");
const setContent = content => {
  tableBody.innerHTML = content;
};

export class UI {
  constructor() {}

  listTodos(todos) {
    let todoContent = "";

    todos.forEach(todo => {
      todoContent += `<tr class="table-row">
        <td class="table-data">
          <i class="far fa-check-circle"></i>
          <i class="far fa-times-circle"></i>
        </td>
        <td class="table-data">
          <a href="" class="task-title">${todo.title}</a>
          <input type="hidden" value="${todo.description}">
          <input type="hidden" value="${todo.ID}">
        </td>
        <td class="table-data">08 May 2019</td>
        <td class="table-data">
          <i class="fas fa-pencil-alt icon"></i>
          <i
            class="far fa-trash-alt icon"
            data-target="#delete-task-modal"
            data-toggle="modal"
          ></i>
        </td>
      </tr>`;
    });

    setContent(todoContent);
  }
}

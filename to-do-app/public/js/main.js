import { HTTP } from "./http.js";

const http = new HTTP();

const host = "http://localhost:8080";

const tableBody = document.querySelector(".table-body");

const initialize = async () => {
  try {
    const todos = await getTodoList();

    setTodoContent(todos);
  } catch (err) {
    console.log("Sorry, something went wrong: " + err.message);
  }
};

const getTodoList = async () => {
  try {
    const response = await http.get(`${host}/todos`);
    return response;
  } catch (err) {
    console.log("Sorry, something went wrong: " + err.message);
  }
};

const setTodoContent = todos => {
  let todoContent = "";

  todos.forEach(todo => {
    todoContent += `<tr class="table-row">
        <td class="table-data">
          <i class="far fa-check-circle"></i>
          <i class="far fa-times-circle"></i>
        </td>
        <td class="table-data">
          <a href="" class="task-title">${todo.Title}</a>
          <input type="hidden" value="${todo.Description}">
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

  tableBody.innerHTML = todoContent;
};

initialize();

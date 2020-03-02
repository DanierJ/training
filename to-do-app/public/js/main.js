import { TodoController } from "./controllers/todoController.js";

const todoCtrl = new TodoController();

const initialize = () => {
  todoCtrl.listTodos();
};

initialize();

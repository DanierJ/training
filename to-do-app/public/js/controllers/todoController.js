import { Todo } from "../models/todoModel.js";
import { UI } from "./uiController.js";

const TodoModel = new Todo(),
  uiCtrl = new UI();

export class TodoController {
  async listTodos() {
    try {
      const todos = await TodoModel.getAll();
      uiCtrl.listTodos(todos);
    } catch (err) {
      console.log("Sorry, something went wrong: " + err.message);
    }
  }
}

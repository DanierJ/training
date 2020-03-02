import { HTTP } from "../controllers/httpController.js";

const host = "http://localhost:8080";

const http = new HTTP();

export class Todo {
  async getAll() {
    try {
      const response = await http.get(`${host}/todos`);
      return response;
    } catch (err) {
      console.log("Sorry, something went wrong: " + err.message);
    }
  }
}

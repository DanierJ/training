import { UI } from "./ui.js";

const uiCtrl = new UI();

export class Util {
  constructor() {}

  validateClock({ hours, minutes }, alertBox) {
    // Validating
    if (hours > 12 || minutes > 60 || isNaN(hours) || isNaN(minutes)) {
      uiCtrl.showAlert(alertBox, 2000);
      return false;
    }

    return true;
  }
}

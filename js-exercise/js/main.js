import { UI } from "./controllers/ui.js";
import { Converter } from "./controllers/converter.js";

const uiCtrl = new UI(),
  converterCtrl = new Converter();

const alertBox = document.getElementById("alert"),
  submitBtn = document.getElementById("submit-btn"),
  changeLayoutBtn = document.getElementById("change-ly-btn");

uiCtrl.cleanAlert(alertBox);

changeLayoutBtn.addEventListener("click", uiCtrl.changeLayout);
submitBtn.addEventListener("click", converterCtrl.convertAngle(alertBox));

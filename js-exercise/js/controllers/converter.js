import { UI } from "./ui.js";
import { Util } from "./util.js";

const uiCtrl = new UI(),
  utilCtrl = new Util();

const toAngleConversion = ({ hours, minutes }) => {
  if (hours === 12) {
    hours = 0;
  }

  return Math.abs(hours * 30 + minutes * 0.5 - minutes * 6);
};

export class Converter {
  constructor() {}

  convertAngle(alertBox) {
    return e => {
      const clock = uiCtrl.getInputInfo(e),
        angle = toAngleConversion(clock);

      if (utilCtrl.validateClock(clock, alertBox)) {
        uiCtrl.showInfo({ clock, angle });
      }
    };
  }
}

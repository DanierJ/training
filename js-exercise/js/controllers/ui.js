const hoursInput = document.getElementById("hours-input"),
  minInput = document.getElementById("min-input"),
  boxInfo = document.getElementById("box-info"),
  angleBox = document.getElementById("angle-info"),
  formContainer = document.getElementById("form"),
  formGroups = document.querySelectorAll(".form-group");

export class UI {
  constructor() {}

  showAlert(el, time) {
    el.style.display = "block";

    setTimeout(() => {
      this.cleanAlert(el);
    }, time);
  }

  cleanAlert(el) {
    el.style.display = "none";
  }

  changeLayout() {
    if (
      formContainer.style.flexDirection === "" ||
      formContainer.style.flexDirection === "row"
    ) {
      formGroups.forEach(el => {
        el.style.paddingTop = "2rem";
      });

      formContainer.style.flexDirection = "column";

      return;
    }

    formGroups.forEach(el => {
      el.style.paddingTop = "0";
    });

    formContainer.style.flexDirection = "row";
  }

  showInfo({ clock, angle }) {
    boxInfo.innerHTML = `<p class="info-box-text" id="box-info"> 
                              At <span class="hour">${clock.hours} Hours</span> with <span class="min">${clock.minutes} Minutes</span> there's an angle of:
                          </p>`;

    angleBox.innerHTML = `<p id="angle-info">${angle}º</p>`;
  }

  getInputInfo(e) {
    e.preventDefault();

    const clock = {
      hours: parseInt(hoursInput.value),
      minutes: parseInt(minInput.value)
    };

    return clock;
  }
}

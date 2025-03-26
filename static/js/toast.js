class Toast {
  constructor(level, message) {
    this.level = level;
    this.message = message;
    this.levelClasses = {
      info: "bg-blue-700",
      success: "bg-green-700",
      warning: "bg-yellow-700",
      error: "bg-red-700"
    };
  }

  makeContainer() {
    const button = document.createElement("button");
    button.className = `flex flex-col p-4 rounded-lg min-h-[80px] w-full max-w-md ${this.getLevelClasses()}`;
    button.setAttribute("role", "alert");
    button.setAttribute("aria-label", "Close");
    button.addEventListener("click", () => button.remove());
    
    // Level element - oben, fett und groß
    const levelEl = document.createElement("div");
    levelEl.className = "font-bold text-lg uppercase mb-1 text-white";
    levelEl.textContent = this.level;
    
    // Message element - darunter
    const messageEl = document.createElement("div");
    messageEl.className = "text-white";
    messageEl.textContent = this.message;
    
    button.appendChild(levelEl);
    button.appendChild(messageEl);
    
    return button;
  }

  getLevelClasses() {
    return this.levelClasses[this.level] || this.levelClasses.info;
  }

  show(containerSelector = "#toast-container", duration = 10000) {
    const container = document.querySelector(containerSelector);
    if (!container) {
      console.error(`Toast container not found: ${containerSelector}`);
      return;
    }

    const toast = this.makeContainer();
    
    // Flex-Container für gleichmäßige Abstände
    const toastWrapper = document.createElement("div");
    toastWrapper.className = "mb-2 w-full";
    toastWrapper.appendChild(toast);
    
    container.appendChild(toastWrapper);

    if (duration > 0) {
      setTimeout(() => toastWrapper.remove(), duration);
    }
  }
}

// Event listener bleibt gleich
document.addEventListener("makeToast", function(e) {
  try {
    const toastData = Array.isArray(e.detail) ? e.detail : (e.detail?.value || []);
    
    toastData.forEach(item => {
      if (!item) return;
      const level = (item.level || item.Level || "").toLowerCase();
      const message = item.message || item.Message || "";
      if (level && message) {
        new Toast(level, message).show();
      }
    });
  } catch (error) {
    console.error("Error processing toast event:", error);
  }
});

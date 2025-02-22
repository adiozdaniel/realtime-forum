class ToastNotification {
	constructor() {
		this.notifications = document.querySelector(".toast-notifications");
		this.timer = 5000;
	}

	createToast(type, message) {
		const icons = {
			success: "fa-circle-check",
			error: "fa-circle-xmark",
			warning: "fa-circle-exclamation",
			info: "fa-circle-info",
		};

		const toast = document.createElement("li");
		toast.className = `toast ${type}`;
		toast.innerHTML = `
			<div class="column">
				<i class="fa-solid ${icons[type] || icons.info}"></i>
				<span>${message}</span>
			</div>
			<i class="fa-solid fa-xmark" onclick="this.parentElement.remove()"></i>
		`;

		this.notifications.appendChild(toast);
		setTimeout(() => toast.remove(), this.timer);
	}
}

// Usage example
const toast = new ToastNotification();

export { toast };

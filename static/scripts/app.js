import { navigateTo, router } from "./router.js";

const initApp = () => {
	console.log("opened page");

	document.body.addEventListener("click", (event) => {
		if (event.target.matches("[data-link]")) {
			event.preventDefault();
			navigateTo(event.target.href);
		}
	});

	router();
};

document.addEventListener("DOMContentLoaded", initApp);

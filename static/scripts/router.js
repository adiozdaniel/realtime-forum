import { views } from "./views.js";

export const navigateTo = (url) => {
	history.pushState(null, null, url);
	router();
};

const routes = {
	"/": views.home,
	"/about": views.about,
	"/contact": views.contact,
};

export const router = () => {
	const view = routes[location.pathname] || views.notFound;
	document.getElementById("MainApp").innerHTML = view();
};

// Listen for browser navigation (Back/Forward)
window.addEventListener("popstate", router);

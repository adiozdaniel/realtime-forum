import { AuthService } from "./authservice.js";

class Authmiddleware {
	constructor() {
		this.authService = new AuthService();
		this.authButton = document.querySelector(".sign-in-button");
	}

	// Authchecker method to check if the user is authenticated
	async authChecker() {
		const isAuthenticated = await this.authService.isAuthenticated();

		if (isAuthenticated.error) {
			console.log(isAuthenticated.message);
			return null;
		}

		localStorage.removeItem("userdata");

		if (isAuthenticated?.data) {
			localStorage.setItem("userdata", JSON.stringify(isAuthenticated.data));
			this.authButton.textContent = "Sign Out";
			return isAuthenticated.data;
		}

		this.authButton.textContent = "Sign In";
		return null;
	}
}

const authmiddleware = new Authmiddleware();

// Export a function to get user data instead of a resolved variable
async function getUserData() {
	return await authmiddleware.authChecker();
}

export { getUserData };
